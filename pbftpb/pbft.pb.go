// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pbft.proto

package pbftpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MessageType int32

const (
	MessageType_MsgHup            MessageType = 0
	MessageType_MsgBeat           MessageType = 1
	MessageType_MsgProp           MessageType = 2
	MessageType_MsgApp            MessageType = 3
	MessageType_MsgAppResp        MessageType = 4
	MessageType_MsgVote           MessageType = 5
	MessageType_MsgVoteResp       MessageType = 6
	MessageType_MsgSnap           MessageType = 7
	MessageType_MsgHeartbeat      MessageType = 8
	MessageType_MsgHeartbeatResp  MessageType = 9
	MessageType_MsgUnreachable    MessageType = 10
	MessageType_MsgSnapStatus     MessageType = 11
	MessageType_MsgCheckQuorum    MessageType = 12
	MessageType_MsgTransferLeader MessageType = 13
	MessageType_MsgTimeoutNow     MessageType = 14
	MessageType_MsgReadIndex      MessageType = 15
	MessageType_MsgReadIndexResp  MessageType = 16
	MessageType_MsgPreVote        MessageType = 17
	MessageType_MsgPreVoteResp    MessageType = 18
	MessageType_NewTrxReq         MessageType = 19
	MessageType_NewTrxResp        MessageType = 20
	MessageType_PreprepareReq     MessageType = 21
	MessageType_PreprepareResp    MessageType = 22
	MessageType_PrepareReq        MessageType = 23
	MessageType_PrepareResp       MessageType = 24
	MessageType_CommitReq         MessageType = 25
	MessageType_CommitResp        MessageType = 26
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:  "MsgHup",
		1:  "MsgBeat",
		2:  "MsgProp",
		3:  "MsgApp",
		4:  "MsgAppResp",
		5:  "MsgVote",
		6:  "MsgVoteResp",
		7:  "MsgSnap",
		8:  "MsgHeartbeat",
		9:  "MsgHeartbeatResp",
		10: "MsgUnreachable",
		11: "MsgSnapStatus",
		12: "MsgCheckQuorum",
		13: "MsgTransferLeader",
		14: "MsgTimeoutNow",
		15: "MsgReadIndex",
		16: "MsgReadIndexResp",
		17: "MsgPreVote",
		18: "MsgPreVoteResp",
		19: "NewTrxReq",
		20: "NewTrxResp",
		21: "PreprepareReq",
		22: "PreprepareResp",
		23: "PrepareReq",
		24: "PrepareResp",
		25: "CommitReq",
		26: "CommitResp",
	}
	MessageType_value = map[string]int32{
		"MsgHup":            0,
		"MsgBeat":           1,
		"MsgProp":           2,
		"MsgApp":            3,
		"MsgAppResp":        4,
		"MsgVote":           5,
		"MsgVoteResp":       6,
		"MsgSnap":           7,
		"MsgHeartbeat":      8,
		"MsgHeartbeatResp":  9,
		"MsgUnreachable":    10,
		"MsgSnapStatus":     11,
		"MsgCheckQuorum":    12,
		"MsgTransferLeader": 13,
		"MsgTimeoutNow":     14,
		"MsgReadIndex":      15,
		"MsgReadIndexResp":  16,
		"MsgPreVote":        17,
		"MsgPreVoteResp":    18,
		"NewTrxReq":         19,
		"NewTrxResp":        20,
		"PreprepareReq":     21,
		"PreprepareResp":    22,
		"PrepareReq":        23,
		"PrepareResp":       24,
		"CommitReq":         25,
		"CommitResp":        26,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_pbft_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_pbft_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{0}
}

type EntryType int32

const (
	EntryType_EntryNormal     EntryType = 0
	EntryType_EntryConfChange EntryType = 1
)

// Enum value maps for EntryType.
var (
	EntryType_name = map[int32]string{
		0: "EntryNormal",
		1: "EntryConfChange",
	}
	EntryType_value = map[string]int32{
		"EntryNormal":     0,
		"EntryConfChange": 1,
	}
)

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_pbft_proto_enumTypes[1].Descriptor()
}

func (EntryType) Type() protoreflect.EnumType {
	return &file_pbft_proto_enumTypes[1]
}

func (x EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryType.Descriptor instead.
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{1}
}

type ConfChangeType int32

const (
	ConfChangeType_ConfChangeAddNode        ConfChangeType = 0
	ConfChangeType_ConfChangeRemoveNode     ConfChangeType = 1
	ConfChangeType_ConfChangeUpdateNode     ConfChangeType = 2
	ConfChangeType_ConfChangeAddLearnerNode ConfChangeType = 3
)

// Enum value maps for ConfChangeType.
var (
	ConfChangeType_name = map[int32]string{
		0: "ConfChangeAddNode",
		1: "ConfChangeRemoveNode",
		2: "ConfChangeUpdateNode",
		3: "ConfChangeAddLearnerNode",
	}
	ConfChangeType_value = map[string]int32{
		"ConfChangeAddNode":        0,
		"ConfChangeRemoveNode":     1,
		"ConfChangeUpdateNode":     2,
		"ConfChangeAddLearnerNode": 3,
	}
)

func (x ConfChangeType) Enum() *ConfChangeType {
	p := new(ConfChangeType)
	*p = x
	return p
}

func (x ConfChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_pbft_proto_enumTypes[2].Descriptor()
}

func (ConfChangeType) Type() protoreflect.EnumType {
	return &file_pbft_proto_enumTypes[2]
}

func (x ConfChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfChangeType.Descriptor instead.
func (ConfChangeType) EnumDescriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{2}
}

type PbftMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceId uint64      `protobuf:"varint,1,opt,name=SequenceId,proto3" json:"SequenceId,omitempty"`
	From       uint64      `protobuf:"varint,2,opt,name=From,proto3" json:"From,omitempty"`
	To         uint64      `protobuf:"varint,3,opt,name=To,proto3" json:"To,omitempty"`
	ViewId     uint64      `protobuf:"varint,4,opt,name=ViewId,proto3" json:"ViewId,omitempty"`
	MsgType    MessageType `protobuf:"varint,5,opt,name=MsgType,proto3,enum=MessageType" json:"MsgType,omitempty"`
	Digest     string      `protobuf:"bytes,6,opt,name=Digest,proto3" json:"Digest,omitempty"`
}

func (x *PbftMsg) Reset() {
	*x = PbftMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbftMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbftMsg) ProtoMessage() {}

func (x *PbftMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbftMsg.ProtoReflect.Descriptor instead.
func (*PbftMsg) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{0}
}

func (x *PbftMsg) GetSequenceId() uint64 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

func (x *PbftMsg) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *PbftMsg) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *PbftMsg) GetViewId() uint64 {
	if x != nil {
		return x.ViewId
	}
	return 0
}

func (x *PbftMsg) GetMsgType() MessageType {
	if x != nil {
		return x.MsgType
	}
	return MessageType_MsgHup
}

func (x *PbftMsg) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewId uint64    `protobuf:"varint,2,opt,name=ViewId,proto3" json:"ViewId,omitempty"`
	Index  uint64    `protobuf:"varint,3,opt,name=Index,proto3" json:"Index,omitempty"`
	Type   EntryType `protobuf:"varint,1,opt,name=Type,proto3,enum=EntryType" json:"Type,omitempty"`
	Data   []byte    `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetViewId() uint64 {
	if x != nil {
		return x.ViewId
	}
	return 0
}

func (x *Entry) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Entry) GetType() EntryType {
	if x != nil {
		return x.Type
	}
	return EntryType_EntryNormal
}

func (x *Entry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type HardState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View   uint64 `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	Vote   uint64 `protobuf:"varint,2,opt,name=Vote,proto3" json:"Vote,omitempty"`
	Commit uint64 `protobuf:"varint,3,opt,name=Commit,proto3" json:"Commit,omitempty"`
}

func (x *HardState) Reset() {
	*x = HardState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardState) ProtoMessage() {}

func (x *HardState) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardState.ProtoReflect.Descriptor instead.
func (*HardState) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{2}
}

func (x *HardState) GetView() uint64 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *HardState) GetVote() uint64 {
	if x != nil {
		return x.Vote
	}
	return 0
}

func (x *HardState) GetCommit() uint64 {
	if x != nil {
		return x.Commit
	}
	return 0
}

type ConfState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes    []uint64 `protobuf:"varint,1,rep,packed,name=nodes,proto3" json:"nodes,omitempty"`
	Learners []uint64 `protobuf:"varint,2,rep,packed,name=learners,proto3" json:"learners,omitempty"`
}

func (x *ConfState) Reset() {
	*x = ConfState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfState) ProtoMessage() {}

func (x *ConfState) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfState.ProtoReflect.Descriptor instead.
func (*ConfState) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{3}
}

func (x *ConfState) GetNodes() []uint64 {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *ConfState) GetLearners() []uint64 {
	if x != nil {
		return x.Learners
	}
	return nil
}

type ConfChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type    ConfChangeType `protobuf:"varint,2,opt,name=Type,proto3,enum=ConfChangeType" json:"Type,omitempty"`
	NodeID  uint64         `protobuf:"varint,3,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	Context []byte         `protobuf:"bytes,4,opt,name=Context,proto3" json:"Context,omitempty"`
}

func (x *ConfChange) Reset() {
	*x = ConfChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfChange) ProtoMessage() {}

func (x *ConfChange) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfChange.ProtoReflect.Descriptor instead.
func (*ConfChange) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{4}
}

func (x *ConfChange) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ConfChange) GetType() ConfChangeType {
	if x != nil {
		return x.Type
	}
	return ConfChangeType_ConfChangeAddNode
}

func (x *ConfChange) GetNodeID() uint64 {
	if x != nil {
		return x.NodeID
	}
	return 0
}

func (x *ConfChange) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_pbft_proto protoreflect.FileDescriptor

var file_pbft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a,
	0x07, 0x50, 0x62, 0x66, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x56, 0x69, 0x65, 0x77, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x56, 0x69,
	0x65, 0x77, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x22, 0x69, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x56, 0x69, 0x65, 0x77, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x56,
	0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x4b, 0x0a, 0x09, 0x48, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x56, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x12, 0x12, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x56, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x3d, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x66, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x73, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2a, 0xd9, 0x03, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x48, 0x75, 0x70, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x4d, 0x73, 0x67, 0x42, 0x65, 0x61, 0x74, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x73, 0x67,
	0x50, 0x72, 0x6f, 0x70, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x41, 0x70, 0x70,
	0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x56, 0x6f, 0x74, 0x65, 0x10, 0x05, 0x12,
	0x0f, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x10, 0x06,
	0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x10, 0x07, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x10, 0x08, 0x12,
	0x14, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x55, 0x6e, 0x72, 0x65,
	0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x73, 0x67,
	0x53, 0x6e, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e,
	0x4d, 0x73, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x10, 0x0c,
	0x12, 0x15, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x10, 0x0d, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x73, 0x67, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4e, 0x6f, 0x77, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x10, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x50, 0x72, 0x65, 0x56, 0x6f, 0x74, 0x65,
	0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x50, 0x72, 0x65, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x10, 0x12, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x78,
	0x52, 0x65, 0x71, 0x10, 0x13, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x70, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x10, 0x15, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x10, 0x16, 0x12, 0x0e, 0x0a, 0x0a,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x10, 0x17, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x10, 0x18, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x10, 0x19, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x10, 0x1a, 0x2a, 0x31, 0x0a, 0x09,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x01, 0x2a,
	0x79, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41,
	0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18,
	0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x64, 0x64, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pbft_proto_rawDescOnce sync.Once
	file_pbft_proto_rawDescData = file_pbft_proto_rawDesc
)

func file_pbft_proto_rawDescGZIP() []byte {
	file_pbft_proto_rawDescOnce.Do(func() {
		file_pbft_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbft_proto_rawDescData)
	})
	return file_pbft_proto_rawDescData
}

var file_pbft_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pbft_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pbft_proto_goTypes = []interface{}{
	(MessageType)(0),    // 0: MessageType
	(EntryType)(0),      // 1: EntryType
	(ConfChangeType)(0), // 2: ConfChangeType
	(*PbftMsg)(nil),     // 3: PbftMsg
	(*Entry)(nil),       // 4: Entry
	(*HardState)(nil),   // 5: HardState
	(*ConfState)(nil),   // 6: ConfState
	(*ConfChange)(nil),  // 7: ConfChange
}
var file_pbft_proto_depIdxs = []int32{
	0, // 0: PbftMsg.MsgType:type_name -> MessageType
	1, // 1: Entry.Type:type_name -> EntryType
	2, // 2: ConfChange.Type:type_name -> ConfChangeType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pbft_proto_init() }
func file_pbft_proto_init() {
	if File_pbft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbftMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfChange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbft_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbft_proto_goTypes,
		DependencyIndexes: file_pbft_proto_depIdxs,
		EnumInfos:         file_pbft_proto_enumTypes,
		MessageInfos:      file_pbft_proto_msgTypes,
	}.Build()
	File_pbft_proto = out.File
	file_pbft_proto_rawDesc = nil
	file_pbft_proto_goTypes = nil
	file_pbft_proto_depIdxs = nil
}