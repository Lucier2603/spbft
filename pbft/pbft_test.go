package pbft

import (
	"bytes"
	"csic/pbft/pbftpb"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestPbft(t *testing.T) {

	cfg1 := &Config{
		ID:1,
		IP:"127.0.0.1",
		PBFT_PORT:"8081",
		HTTP_PORT:"8091",
		Peers:[]uint64{1,2,3,4},
		Routers: map[uint64]string{1:"127.0.0.1:8081", 2:"127.0.0.1:8082", 3:"127.0.0.1:8083",4:"127.0.0.1:8084"},
		HeartbeatTick:1000,
		StorageDir:"E:\\\\db1.txt",
	}

	cfg2 := &Config{
		ID:2,
		IP:"127.0.0.1",
		PBFT_PORT:"8082",
		HTTP_PORT:"8092",
		Peers:[]uint64{1,2,3,4},
		Routers: map[uint64]string{1:"127.0.0.1:8081", 2:"127.0.0.1:8082", 3:"127.0.0.1:8083",4:"127.0.0.1:8084"},
		HeartbeatTick:1000,
		StorageDir:"E:\\db2.txt",
	}

	cfg3 := &Config{
		ID:3,
		IP:"127.0.0.1",
		PBFT_PORT:"8083",
		HTTP_PORT:"8093",
		Peers:[]uint64{1,2,3,4},
		Routers: map[uint64]string{1:"127.0.0.1:8081", 2:"127.0.0.1:8082", 3:"127.0.0.1:8083",4:"127.0.0.1:8084"},
		HeartbeatTick:1000,
		StorageDir:"E:\\db3.txt",
	}

	cfg4 := &Config{
		ID:4,
		IP:"127.0.0.1",
		PBFT_PORT:"8084",
		HTTP_PORT:"8094",
		Peers:[]uint64{1,2,3,4},
		Routers: map[uint64]string{1:"127.0.0.1:8081", 2:"127.0.0.1:8082", 3:"127.0.0.1:8083",4:"127.0.0.1:8084"},
		HeartbeatTick:1000,
		StorageDir:"E:\\db4.txt",
	}

	p1 := NewPbftNode(cfg1)
	p2 := NewPbftNode(cfg2)
	p3 := NewPbftNode(cfg3)
	p4 := NewPbftNode(cfg4)

	p1.Start()
	p2.Start()
	p3.Start()
	p4.Start()


	msg := &pbftpb.PbftMsg{
		SequenceId: 99,
		From: 999,
		To: 1,
		ViewId: 1,
		MsgType: pbftpb.MessageType_NewTrxReq,
		Digest: "xxxx",
	}


	// client
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
	}
	buff := bytes.NewBuffer(jsonMsg)
	http.Post("http://127.0.0.1:8081" + Url(msg.MsgType), "application/json", buff)

	//time.Sleep(time.Second * 2)

	msg.Digest = "yyyy"
	jsonMsg, err = json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
	}
	buff = bytes.NewBuffer(jsonMsg)
	http.Post("http://127.0.0.1:8081" + Url(msg.MsgType), "application/json", buff)

	time.Sleep(time.Second * 5)

	es3 := p3.Entries(0, 20, 20)
	fmt.Println("p3 ")
	fmt.Println(len(es3))
	if len(es3) != 0 {
		for _, e := range es3 {
			fmt.Println(" p3 entry  --  " + e.String())
			fmt.Println(string(e.Data))
		}
	} else {
		fmt.Println("p3 empty")
	}

	time.Sleep(time.Second * 5)
	es3 = p3.Entries(0, 20, 20)
	fmt.Println("p3 ")
	fmt.Println(len(es3))
	if len(es3) != 0 {
		for _, e := range es3 {
			fmt.Println(" p3 entry  --  " + e.String())
			fmt.Println(string(e.Data))
		}
	} else {
		fmt.Println("p3 empty")
	}

	time.Sleep(time.Second * 30)
}
