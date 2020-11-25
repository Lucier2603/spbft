package pbft

import (
	"bytes"
	"csic/pbft/pbftpb"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


type PbftServer struct {
	pbft_address	string

	// 消息发送队列
	send_msg_q		*Queue
	// 需要发送pbft msg的channel
	pbftCh			chan *PbftMsgWrapper

	// pbft msg 输出channel
	stepCh 			chan *pbftpb.PbftMsg

}

func NewPbftServer(cfg *Config, hCh chan *PbftMsgWrapper, sCh chan *pbftpb.PbftMsg) *PbftServer {
	return &PbftServer{
		pbft_address:cfg.IP + ":" + cfg.PBFT_PORT,
		stepCh:sCh,
		pbftCh:hCh,
		send_msg_q:NewQueue(),
	}
}

func (server *PbftServer) Start() error {

	// start channel
	go func() {
		for {
			select {
			case w := <- server.pbftCh:
				server.send_msg_q.Push(w)
			}
		}
	}()

	// start send
	go func() {
		for {
			w := server.send_msg_q.Pop()
			if w != nil {
				go doSend(w.(*PbftMsgWrapper))
			} else {
				time.Sleep(time.Duration(100000000))
			}
		}
	}()

	// start listen
	go func() {
		if err := http.ListenAndServe(server.pbft_address, server); err != nil {
			DefaultLogger().Error("PbftServer start error!  {}", err)
		}
	}()

	DefaultLogger().Info("PbftServer started ...")

	return nil
}

func (server *PbftServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var msg *pbftpb.PbftMsg
	err := json.NewDecoder(request.Body).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.stepCh <- msg

	return
}

func (server *PbftServer) Send(w *PbftMsgWrapper) {
	server.send_msg_q.Push(w)
}

func (server *HttpServer) SendNow(w *PbftMsgWrapper) {
	doSend(w)
}

func doSend(w *PbftMsgWrapper) error {

	jsonMsg, err := json.Marshal(w.msg)
	if err != nil {
		return ErrMarshalJson
	}
	buff := bytes.NewBuffer(jsonMsg)

	http.Post("http://" + w.url, "application/json", buff)

	return nil
}

type PbftMsgWrapper struct {
	msg			*pbftpb.PbftMsg
	url			string
}









type HttpServer struct {
	http_address	string

	// 普通 msg 输出channel
	receiveCh		chan *HttpRequest
	sendCh			chan *HttpResponse

}

func NewHttpServer(cfg *Config, rCh chan *HttpRequest, sCh chan *HttpResponse) *HttpServer {
	return &HttpServer{
		http_address: cfg.IP + ":" + cfg.HTTP_PORT,
		receiveCh:rCh,
		sendCh:sCh,
	}
}

func (server *HttpServer) Start() error {

	// start listen
	go func() {
		if err := http.ListenAndServe(server.http_address, server); err != nil {
			DefaultLogger().Error("PbftServer start error!  {}", err)
		}
	}()

	DefaultLogger().Info("HttpServer started ...")

	return nil
}

func (server *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var req *HttpRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.receiveCh <- req

	select {
	case resp := <- server.sendCh:
		b, err := json.Marshal(resp)
		if err != nil {
			DefaultLogger().Error("Error in marshal JSON http response: {}", resp.data)
		}
		writer.Write(b)
	}
	return
}



type HttpRequest struct {
	reqType		HttpReqType
	data		string
}

type HttpResponse struct {
	data		string
}

