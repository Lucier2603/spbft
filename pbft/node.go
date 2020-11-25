package pbft

import (
	"context"
	"csic/pbft/pbftpb"
	"encoding/json"
	"fmt"
	"time"
)

type SoftState struct {
	Lead      	uint64 // must use atomic operations to access; keep 64-bit aligned.
	PbftState 	StateType
}

type HardState struct {
	Term		uint64
	Vote		uint64
	Commit		uint64
}

type Ready struct {
	// The current volatile state of a Node.
	// SoftState will be nil if there is no update.
	// It is not required to consume or store SoftState.
	*SoftState

	// The current state of a Node to be saved to stable storage BEFORE
	// Messages are sent.
	// HardState will be equal to empty state if there is no update.
	* HardState

	// Entries specifies entries to be saved to stable storage BEFORE
	// Messages are sent.
	Entries []pbftpb.Entry

	// CommittedEntries specifies entries to be committed to a
	// store/state-machine. These have previously been committed to stable
	// store.
	CommittedEntries []pbftpb.Entry

	// Messages specifies outbound messages to be sent AFTER Entries are
	// committed to stable storage.
	// If it contains a MsgSnap message, the application MUST report back to raft
	// when the snapshot has been received or has failed by calling ReportSnapshot.
	Messages []*pbftpb.PbftMsg
}

type Node interface {
	Start() error
	Stop() error

	Status() *State

	// start elect
	Tick()
	Campaign(ctx context.Context) error
	Step(ctx context.Context, msg *pbftpb.PbftMsg) error
	// 应用层在接收到Ready后，应当处理Ready中的每一个有效字段，处理完毕后，调用Advance()通知raft Ready已处理完毕。
	Ready() <-chan Ready
	Advance() error

	Propose(ctx context.Context, data []byte) error

	ApplyConfChange(cc pbftpb.ConfChange) *pbftpb.ConfState

	ProposeConfChange(ctx context.Context, cc pbftpb.ConfChange) error

	Entries(low uint64, high uint64, max uint64) []pbftpb.Entry
}

type PbftNode struct {
	ID			uint64
	cfg			*Config

	pbftServer *PbftServer
	httpServer *HttpServer

	electTick	int

	stepCh		chan *pbftpb.PbftMsg
	receiveCh	chan *HttpRequest
	sendCh		chan *HttpResponse

	core		*Core
}

type Peer struct {
	ID      	uint64
	Context 	[]byte
}


func NewPbftNode(cfg *Config) Node {

	c, err := NewCore(cfg)
	if err != nil {
		panic(err)
	}

	n := &PbftNode{
		ID: cfg.ID,
		cfg: cfg,
		electTick: cfg.ElectionTick,
		stepCh: make(chan *pbftpb.PbftMsg),
		receiveCh: make(chan *HttpRequest),
		sendCh: make(chan *HttpResponse),
		core:c,
	}

	return n
}



func (n *PbftNode) Start() error {

	DefaultLogger().Info("Start Pbft Node ...")

	// init pbftServer
	n.pbftServer = NewPbftServer(n.cfg, n.core.HttpCh, n.stepCh)
	go n.pbftServer.Start()

	// init httpServer
	n.httpServer = NewHttpServer(n.cfg, n.receiveCh, n.sendCh)
	go n.httpServer.Start()

	// 维持一个计时器，定时发送心跳
	go func() {
		for {
			// 定时器每1秒工作一次
			time.Sleep(1000000000)

			// 给core发送心跳信息
			n.Tick()
		}
	}()

	// channel接收
	go func() {
		for {
			select {
			case msg := <- n.stepCh:
				n.Step(nil, msg)
			case req := <- n.receiveCh:
				n.sendCh <- n.handleHttp(req)
			}
		}
	}()

	// 启动核心
	n.core.Start()

	// todo refresh state according to other peers

	return nil
}

func (p *PbftNode) Stop() error {
	return nil
}

func (n *PbftNode) Status() *State {
	return n.core.state
}

func (p *PbftNode) Tick() {
	p.core.tickCh <- 1
}

func (p *PbftNode) Campaign(ctx context.Context) error {
	return nil
}

func (p *PbftNode) Step(ctx context.Context, msg *pbftpb.PbftMsg) error {
	fmt.Println("Node receive msg " + msg.MsgType.String())
	p.core.msgCh <- *msg
	return nil
}

func (p *PbftNode) Ready() <-chan Ready {
	// todo raft 的 readyCh 是怎么用的？
	return p.core.readyCh
}

func (p *PbftNode) Advance() error {
	return nil
}

func (p *PbftNode) Propose(ctx context.Context, data []byte) error {
	return nil
}

func (p *PbftNode) ApplyConfChange(cc pbftpb.ConfChange) *pbftpb.ConfState {
	return nil
}

func (p *PbftNode) ProposeConfChange(ctx context.Context, cc pbftpb.ConfChange) error {
	return nil
}

func (p *PbftNode) Entries(low uint64, high uint64, max uint64) []pbftpb.Entry {
	// todo test
	es, _ := p.core.st.Entries(low, high, max)
	return es
}

// 接收普通的http请求
func (p *PbftNode) handleHttp(req *HttpRequest) *HttpResponse {
	switch req.reqType {
	case ReqType_QueryEntries:
		request := QueryEntryRequest{}
		response := QueryEntryResponse{}
		err := json.Unmarshal([]byte(req.data), request)
		if err != nil {
			return nil
		}
		response.Entries = p.Entries(request.Start, request.End, request.Max)

		resp := HttpResponse{}
		b, err := json.Marshal(response)
		if err != nil {
			return nil
		}
		resp.data = string(b)
		return &resp
	}
	return nil
}