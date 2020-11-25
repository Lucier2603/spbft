package pbft

import (
	"csic/pbft/pbftpb"
)

type StateType uint64
const (
	StatePrimary StateType = iota
	StateReplica
	StateCandidate
)

type HttpReqType uint64
const (
	ReqType_QueryEntries HttpReqType = iota
)

type MsgPhase uint64
const (
	NewRequest MsgPhase = iota
	PrePrepare_Start
	PrePrepare_Done
	Prepare_Start
	Prepare_Done
	Commit_Start
	Commit_Done
	All_Done
)


type State struct {
	Role			StateType
	Leader			uint64
	ViewId			uint64
	CurrentSequence	uint64
}

type MsgStatus struct {
	SequenceId		uint64
	msg				pbftpb.PbftMsg

	Phase			MsgPhase
	// 已经确认prepare/commit的nodeId集合
	PrepareSet		*LongSet
	CommitSet		*LongSet

	// 已经得到reply的nodeId集合
	replyMap		map[pbftpb.MessageType]*LongSet
	resendTimes		int

	NewRequestLogId	uint64
	PreprepareLogId	uint64
	PrepareLogId	uint64
	CommitLogId		uint64
	FinalReplyLogId	uint64
}


func NewMsgStatus(msg pbftpb.PbftMsg, status *State) *MsgStatus {
	ms := &MsgStatus{
		SequenceId: msg.SequenceId,
		msg: msg,
		PrepareSet: NewLongSet(),
		CommitSet: NewLongSet(),
		replyMap: make(map[pbftpb.MessageType]*LongSet),
		resendTimes: 0,
	}

	ms.replyMap[pbftpb.MessageType_NewTrxReq] = NewLongSet()
	ms.replyMap[pbftpb.MessageType_PreprepareResp] = NewLongSet()
	ms.replyMap[pbftpb.MessageType_PrepareReq] = NewLongSet()
	ms.replyMap[pbftpb.MessageType_CommitReq] = NewLongSet()

	if status.Role == StatePrimary {
		ms.Phase = NewRequest
	} else {
		ms.Phase = PrePrepare_Start
	}

	return ms
}

func (s *MsgStatus) OnSend(mt pbftpb.MessageType, to uint64) {
	v, ok := s.replyMap[mt]
	if ok {
		v.put(to)
	}
}

func (s *MsgStatus) OnReply(mt pbftpb.MessageType, from uint64) {
	v, ok := s.replyMap[mt]
	if ok {
		v.remove(from)
	}
}