package pbft

import "csic/pbft/pbftpb"

func ValidMsg(msg *pbftpb.PbftMsg) bool {
	return true
}

func Url(msgType pbftpb.MessageType) string {
	switch msgType {
	case pbftpb.MessageType_NewTrxReq:
		return "/newTrx"
	case pbftpb.MessageType_NewTrxResp:
		return "/newTrxReply"
	case pbftpb.MessageType_PreprepareReq:
		return "/prePrepare"
	case pbftpb.MessageType_PreprepareResp:
		return "/prePrepareReply"
	case pbftpb.MessageType_PrepareReq:
		return "/prepare"
	case pbftpb.MessageType_PrepareResp:
		return "/prepareReply"
	case pbftpb.MessageType_CommitReq:
		return "/commit"
	case pbftpb.MessageType_CommitResp:
		return "/commitReply"
	default:
		return ""
	}
}

func Digest() string {
	return ""
}


// utils
func IsRespMsg(t pbftpb.MessageType) bool {
	return t % 2 == 0
}



// set
var void_v struct{}
type LongSet struct {
	map_long		map[uint64]struct{}
}

func NewLongSet() *LongSet {
	return &LongSet{
		map_long: make(map[uint64]struct{}),
	}
}

func (s *LongSet) put(k uint64) {
	s.map_long[k] = void_v
}

func (s *LongSet) size() int {
	return len(s.map_long)
}

func (s *LongSet) remove(k uint64) {
	delete(s.map_long, k)
}




// queue
type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
	}
)

// Create a new queue
func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}
//获取队列长度
func (this *Queue) Len() int {
	return this.length
}
//返回true队列不为空
func (this *Queue) Any() bool {
	return this.length > 0
}
//返回队列顶端元素
func (this *Queue) Peek() interface{} {
	if this.top == nil {
		return nil
	}
	return this.top.value
}
//入队操作
func (this *Queue) Push(v interface{}) {
	n := &node{nil, nil, v}
	if this.length == 0 {
		this.top = n
		this.rear = this.top
	} else {
		n.pre = this.rear
		this.rear.next = n
		this.rear = n
	}
	this.length++
}
//出队操作
func (this *Queue) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	if this.top.next == nil {
		this.top = nil
	} else {
		this.top = this.top.next
		this.top.pre.next = nil
		this.top.pre = nil
	}
	this.length--
	return n.value
}

