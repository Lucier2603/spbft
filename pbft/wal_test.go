package pbft

import (
	"csic/pbft/pbftpb"
	"fmt"
	"google.golang.org/protobuf/proto"
	"testing"
)


func TestMarshal(t *testing.T) {
	e := &pbftpb.PbftMsg{
		SequenceId:1,
		From:2,
		Digest:"xxxx",
	}
	fmt.Println(e.String())

	bytes, _ := proto.Marshal(e)
	e2 := &pbftpb.PbftMsg{}
	proto.Unmarshal(bytes, e2)

	fmt.Println("1")
	fmt.Println(e2.String())
}

func TestStorage(t *testing.T) {
	st := NewMemoryStorage("E:\\1.txt")

	d := []byte("11223344")

	e := &pbftpb.Entry{
		ViewId:1,
		Index:1,
		Type:pbftpb.EntryType_EntryNormal,
		Data:d,
	}

	st.Append([]pbftpb.Entry{*e})

	e2, _ := st.Entry(1)
	fmt.Print(string(e2.Data))
}
