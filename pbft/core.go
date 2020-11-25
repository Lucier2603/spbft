package pbft

import (
	"csic/pbft/pbftpb"
	"errors"
	"sync"
)

var ErrNotLeader = errors.New("current role is not leader")
var ErrMsgValidFail = errors.New("msg valid fail")
var ErrMarshalJson = errors.New("msg marshal json fail")
var ErrUnMarshalJson = errors.New("msg unmarshal json fail")
var ErrDigestUnmatch = errors.New("already exist different digest for this sequence")

type Core struct {
	ID			uint64

	peers		[]uint64
	routers		map[uint64]string

	msgStatuses map[uint64]*MsgStatus
	state       *State
	lock        *sync.Mutex

	st			Storage

	mh			msgHandler

	stopCh		chan int
	msgCh		chan pbftpb.PbftMsg
	tickCh		chan int
	readyCh		chan Ready
	HttpCh		chan *PbftMsgWrapper
}

type msgHandler func(c *Core, msg *pbftpb.PbftMsg) error

func NewCore(cfg *Config) (*Core, error) {
	//w, err := CreateWal("", nil)
	//if err != nil {
	//	return nil, err
	//}

	c := &Core{
		ID:				cfg.ID,
		peers:			cfg.Peers,
		routers: 		cfg.Routers,
		msgStatuses:	make(map[uint64]*MsgStatus),
		lock:			&sync.Mutex{},
		st:				NewMemoryStorage(cfg.StorageDir),
		stopCh:			make(chan int),
		msgCh: 			make(chan pbftpb.PbftMsg),
		tickCh:			make(chan int),
		readyCh: 		make(chan Ready),
		HttpCh:			make(chan *PbftMsgWrapper),
	}

	return c, nil
}

func (core *Core) Start() error {
	if err := core.loadState(); err != nil {
		return err
	}

	go core.run()

	return nil
}

func (c *Core) loadState() error {
	// load state from storage
	// todo 暂时写死，供测试用
	c.state = &State{
		Role: StateReplica,
		Leader: 1,
		ViewId: 1,
		CurrentSequence: 0,
	}
	if c.ID == 1 {
		c.state.Role = StatePrimary
		c.mh = msgHandler_primary
	} else {
		c.mh = msgHandler_replica
	}
	return nil;
}

func (c *Core) saveState() error {
	// save state to local file
	return nil
}

func (core *Core) Stop() {
	core.stopCh <- 1
}

func (core *Core) stop() error {
	// todo wait for msg to be done
	// todo save state to storage
	// todo finally we stop

	return nil
}

func (core *Core) Step(msg *pbftpb.PbftMsg) {
	core.msgCh <- *msg
}



func (core *Core) run() {
	for {
		select {
		case _ = <- core.stopCh:
			core.stop()

		case _ = <- core.tickCh:
			core.tick()
			// 运行定时重发
			core.resend()

		case msg := <- core.msgCh:
			if !core.valid(&msg) {
				continue
			}
			core.mh(core, &msg)
		}
	}
}

func (core *Core) tick() error {
	// todo do something

	return nil
}

func (core *Core) valid(msg *pbftpb.PbftMsg) bool {
	return true
}

func msgHandler_primary(c *Core, msg *pbftpb.PbftMsg) error {
	if !msgHandler_common(c, msg) {
		return nil
	}

	switch msg.MsgType {
	case pbftpb.MessageType_NewTrxReq:
		handle_newtrx(c, msg)
	case pbftpb.MessageType_PreprepareReq:
		handle_preprepare(c, msg)
	case pbftpb.MessageType_PrepareReq:
		handle_prepare(c, msg)
	case pbftpb.MessageType_CommitReq:
		handle_commit(c, msg)
	}
	return nil
}

func msgHandler_replica(c *Core, msg *pbftpb.PbftMsg) error {

	if !msgHandler_common(c, msg) {
		return nil
	}
	if msgHandler_response(c, msg) {
		return nil
	} 

	switch msg.MsgType {
	case pbftpb.MessageType_PreprepareReq:
		handle_preprepare(c, msg)
	case pbftpb.MessageType_PrepareReq:
		handle_prepare(c, msg)
	case pbftpb.MessageType_CommitReq:
		handle_commit(c, msg)
	}
	return nil
}

func msgHandler_common(c *Core, msg *pbftpb.PbftMsg) bool {

	// check view
	if msg.ViewId == c.state.ViewId {
		return true
	} else if msg.ViewId > c.state.ViewId {
		// todo 触发同步数据
		// send sync msg
	}
	// todo check role
	return false
}

func msgHandler_response(c *Core, msg *pbftpb.PbftMsg) bool {
	if !IsRespMsg(msg.MsgType) {
		return false
	}

	ms, ok := c.msgStatuses[msg.SequenceId]
	// 不存在的情况，有可能该resp延迟很久才到达，此时已经commit.
	if ok {
		ms.OnReply(msg.MsgType, msg.From)
	}

	return true
}

func handle_newtrx(c *Core, msg *pbftpb.PbftMsg) error {

	c.lock.Lock()
	defer c.lock.Unlock()

	c.state.CurrentSequence++
	msg.SequenceId = c.state.CurrentSequence
	//msg.Digest = Digest()
	msg.ViewId = c.state.ViewId

	// 暂时存到本地
	msgStatus := NewMsgStatus(*msg, c.state)
	msgStatus.NewRequestLogId = msg.SequenceId
	c.msgStatuses[msg.SequenceId] = msgStatus

	// broadcast
	msg.MsgType = pbftpb.MessageType_PreprepareReq
	c.broadcast(*msg)

	return nil
}

func handle_preprepare(c *Core, msg *pbftpb.PbftMsg) error {

	// 检查是否本地已有msg
	ms := c.msgStatuses[msg.SequenceId]

	if ms != nil {
		if ms.msg.Digest != msg.Digest {
			// do nothing
			return ErrDigestUnmatch
		} else {
			// 本地已有，则判断状态
			if ms.Phase >= PrePrepare_Done {
				c.reply(*msg)
				return nil
			}
		}
	}


	c.lock.Lock()
	defer c.lock.Unlock()

	if c.state.CurrentSequence < msg.SequenceId {
		c.state.CurrentSequence = msg.SequenceId
	}

	msgStatus := NewMsgStatus(*msg, c.state)
	msgStatus.Phase = PrePrepare_Start
	c.msgStatuses[msg.SequenceId] = msgStatus

	c.reply(*msg)

	// 开始prepare流程

	msg.MsgType = pbftpb.MessageType_PrepareReq

	msgStatus.Phase = PrePrepare_Done
	msgStatus.resendTimes = 0
	c.broadcast(*msg)

	msgStatus.Phase = Prepare_Start

	return nil
}


func handle_prepare(c *Core, msg *pbftpb.PbftMsg) error {

	if !ValidMsg(msg) {
		return ErrMsgValidFail
	}

	ms := c.msgStatuses[msg.SequenceId]
	if ms == nil {
		// 先收到prepare消息，则不予理会
		return nil
	}

	if ms.msg.Digest != msg.Digest {
		// do nothing
		return ErrDigestUnmatch
	}

	// 状态判断
	if ms.Phase < Prepare_Start {
		return nil
	}
	if ms.Phase >= Prepare_Done {
		c.reply(*msg)
		return nil
	}

	// 仅当 prepare_start，才实际处理
	prepares := c.msgStatuses[msg.SequenceId].PrepareSet
	prepares.put(msg.From)

	c.reply(*msg)

	// 当收到的prepare消息数量到达阈值时，进入commit阶段
	if prepares.size() < c.threshold() {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	ms.Phase = Prepare_Done

	// broadcast commit message
	msg.MsgType = pbftpb.MessageType_CommitReq

	ms.resendTimes = 0
	c.broadcast(*msg)

	ms.Phase = Commit_Start

	return nil
}

func handle_commit(c *Core, msg *pbftpb.PbftMsg) error {
	if !ValidMsg(msg) {
		return ErrMsgValidFail
	}

	ms := c.msgStatuses[msg.SequenceId]
	if ms == nil {
		// 先收到prepare消息，则不予理会
		return nil
	}

	if ms.msg.Digest != msg.Digest {
		// do nothing
		return ErrDigestUnmatch
	}

	// 状态判断
	if ms.Phase < Commit_Start {
		return nil
	}
	if ms.Phase >= Commit_Done {
		c.reply(*msg)
		return nil
	}


	commits := c.msgStatuses[msg.SequenceId].CommitSet
	commits.put(msg.From)

	c.reply(*msg)

	// 当收到的prepare消息数量到达阈值时，广播commit
	// todo 如果已经广播过了，那么就不再commit.
	if commits.size() < c.threshold() {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	ms.Phase = Commit_Done

	// save msg
	e := &pbftpb.Entry{
		ViewId:msg.ViewId,
		Index:msg.SequenceId,
		Type:pbftpb.EntryType_EntryNormal,
		Data:[]byte(msg.Digest),
	}

	if err := c.st.Append([]pbftpb.Entry{*e}); err != nil {
		// todo 清空中间态数据
		return err
	}

	// todo reply to client
	//msg.MsgType = pbftpb.MessageT
	// save reply msg (final commit)
	//logId, _ := c.logMgr.AppendLog(msg)
	//c.msgStatuses[msg.SequenceId].FinalReplyLogId = logId

	ms.Phase = All_Done

	return nil
}

func (c *Core) threshold() int {
	return (2 * len(c.peers)) / 3 + 1
}

// 定时重发
func (c *Core) resend() error {
	for _, ms := range c.msgStatuses {

		ms.resendTimes++
		if ms.resendTimes > 5 {
			continue
		}
		for msgType, peers := range ms.replyMap{
			msg := ms.msg
			msg.MsgType = msgType
			for peer, _ := range peers.map_long {
				c.send(msg, peer)
			}
		}
	}
	return nil
}











func (c *Core) broadcast(msg pbftpb.PbftMsg) error {
	for peerID, _ := range c.routers {
		c.send(msg, peerID)
	}
	return nil
}

func (c *Core) broadcastNoSelf(msg pbftpb.PbftMsg) error {
	for peerID, _ := range c.routers {
		if peerID == c.ID {
			continue
		}
		c.send(msg, peerID)
	}
	return nil
}

func (c *Core) send(msg pbftpb.PbftMsg, peerID uint64) error {

	msg.To = peerID
	msg.From = c.ID

	// 加入set，以日后重发request
	if !IsRespMsg(msg.MsgType) {
		c.msgStatuses[msg.SequenceId].OnSend(msg.MsgType, msg.To)
	}

	w := &PbftMsgWrapper{
		msg: &msg,
		url: c.routers[peerID] + Url(msg.MsgType),
	}

	c.HttpCh <- w

	return nil
}

func (c *Core) reply(msg pbftpb.PbftMsg) error {
	switch msg.MsgType {
	case pbftpb.MessageType_NewTrxReq: msg.MsgType = pbftpb.MessageType_NewTrxResp
	case pbftpb.MessageType_PreprepareReq: msg.MsgType = pbftpb.MessageType_PreprepareResp
	case pbftpb.MessageType_PrepareReq: msg.MsgType = pbftpb.MessageType_PrepareResp
	case pbftpb.MessageType_CommitReq: msg.MsgType = pbftpb.MessageType_CommitResp
	}

	c.send(msg, msg.From)
	return nil
}
