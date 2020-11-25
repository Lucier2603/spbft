package pbft

import "csic/pbft/pbftpb"

const None uint64 = 0


//type PbftMsg struct {
//	SequenceId			uint64
//
//	From 				uint64
//	To   				uint64
//
//	Client				uint64
//
//	ViewId				uint64
//	MsgType				PbftMsgType
//	Digest				string
//}



type Snapshot struct {
	ViewId				uint64
	CommittedLogId		uint64
}



// 普通http请求 -- 查询entry
type QueryEntryRequest struct {
	Start			uint64
	End				uint64
	Max				uint64
}
type QueryEntryResponse struct {
	Entries			[]pbftpb.Entry
}
