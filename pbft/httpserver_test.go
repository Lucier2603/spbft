package pbft

import (
	"csic/pbft/pbftpb"
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {
	cfg1 := &Config{
		ID:1,
		Peers:[]uint64{1,2,3,4},
		Routers: map[uint64]string{1:"127.0.0.1:8081", 2:"127.0.0.1:8082", 3:"127.0.0.1:8083",4:"127.0.0.1:8084"},
		HeartbeatTick:1000,
	}

	ch := make(chan *PbftMsgWrapper)
	server := NewPbftServer(cfg1, ch, nil)

	server.Start()

	var i uint64
	for i = 0; i < 20; i++ {
		ch <- &PbftMsgWrapper{
			url: "192.168.0.1:8080/test/",
			msg: &pbftpb.PbftMsg{
				SequenceId: 1 + i,
			},
		}
	}

	time.Sleep(time.Duration(4000000000))
}