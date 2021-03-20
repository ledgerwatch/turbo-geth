// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: remote/ethbackend.proto

package remote

import (
	types "github.com/ledgerwatch/turbo-geth/gointerfaces/types"
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

type TxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signedtx []byte `protobuf:"bytes,1,opt,name=signedtx,proto3" json:"signedtx,omitempty"`
}

func (x *TxRequest) Reset() {
	*x = TxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxRequest) ProtoMessage() {}

func (x *TxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxRequest.ProtoReflect.Descriptor instead.
func (*TxRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{0}
}

func (x *TxRequest) GetSignedtx() []byte {
	if x != nil {
		return x.Signedtx
	}
	return nil
}

type AddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash *types.H256 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *AddReply) Reset() {
	*x = AddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReply) ProtoMessage() {}

func (x *AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReply.ProtoReflect.Descriptor instead.
func (*AddReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{1}
}

func (x *AddReply) GetHash() *types.H256 {
	if x != nil {
		return x.Hash
	}
	return nil
}

type EtherbaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EtherbaseRequest) Reset() {
	*x = EtherbaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtherbaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtherbaseRequest) ProtoMessage() {}

func (x *EtherbaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtherbaseRequest.ProtoReflect.Descriptor instead.
func (*EtherbaseRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{2}
}

type EtherbaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *types.H160 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *EtherbaseReply) Reset() {
	*x = EtherbaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtherbaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtherbaseReply) ProtoMessage() {}

func (x *EtherbaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtherbaseReply.ProtoReflect.Descriptor instead.
func (*EtherbaseReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{3}
}

func (x *EtherbaseReply) GetAddress() *types.H160 {
	if x != nil {
		return x.Address
	}
	return nil
}

type NetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetVersionRequest) Reset() {
	*x = NetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetVersionRequest) ProtoMessage() {}

func (x *NetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetVersionRequest.ProtoReflect.Descriptor instead.
func (*NetVersionRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{4}
}

type NetVersionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NetVersionReply) Reset() {
	*x = NetVersionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetVersionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetVersionReply) ProtoMessage() {}

func (x *NetVersionReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetVersionReply.ProtoReflect.Descriptor instead.
func (*NetVersionReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{5}
}

func (x *NetVersionReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{6}
}

type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type uint64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"` // type (only header at that moment)
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`  //  serialized data
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReply.ProtoReflect.Descriptor instead.
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{7}
}

func (x *SubscribeReply) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SubscribeReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWorkRequest) Reset() {
	*x = GetWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkRequest) ProtoMessage() {}

func (x *GetWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkRequest.ProtoReflect.Descriptor instead.
func (*GetWorkRequest) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{8}
}

type GetWorkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderHash  string `protobuf:"bytes,1,opt,name=headerHash,proto3" json:"headerHash,omitempty"`   // 32 bytes hex encoded current block header pow-hash
	SeedHash    string `protobuf:"bytes,2,opt,name=seedHash,proto3" json:"seedHash,omitempty"`       // 32 bytes hex encoded seed hash used for DAG
	Target      string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`           // 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
	BlockNumber string `protobuf:"bytes,4,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"` // hex encoded block number
}

func (x *GetWorkReply) Reset() {
	*x = GetWorkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_ethbackend_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkReply) ProtoMessage() {}

func (x *GetWorkReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_ethbackend_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkReply.ProtoReflect.Descriptor instead.
func (*GetWorkReply) Descriptor() ([]byte, []int) {
	return file_remote_ethbackend_proto_rawDescGZIP(), []int{9}
}

func (x *GetWorkReply) GetHeaderHash() string {
	if x != nil {
		return x.HeaderHash
	}
	return ""
}

func (x *GetWorkReply) GetSeedHash() string {
	if x != nil {
		return x.SeedHash
	}
	return ""
}

func (x *GetWorkReply) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *GetWorkReply) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

var File_remote_ethbackend_proto protoreflect.FileDescriptor

var file_remote_ethbackend_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x65, 0x74, 0x68, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x09, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x74, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x74, 0x78, 0x22, 0x2b, 0x0a,
	0x08, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x48, 0x32, 0x35, 0x36, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x12, 0x0a, 0x10, 0x45, 0x74,
	0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37,
	0x0a, 0x0e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x25, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x31, 0x36, 0x30, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x4e, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x0f,
	0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x12, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x10, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x84, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xb3, 0x02, 0x0a, 0x0a, 0x45, 0x54, 0x48, 0x42, 0x41,
	0x43, 0x4b, 0x45, 0x4e, 0x44, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x3d, 0x0a, 0x09, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x12, 0x18,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x40, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3f, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x18, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x16,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x31, 0x0a, 0x10,
	0x69, 0x6f, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x2d, 0x67, 0x65, 0x74, 0x68, 0x2e, 0x64, 0x62,
	0x42, 0x0a, 0x45, 0x54, 0x48, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x4e, 0x44, 0x50, 0x01, 0x5a, 0x0f,
	0x2e, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x3b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_ethbackend_proto_rawDescOnce sync.Once
	file_remote_ethbackend_proto_rawDescData = file_remote_ethbackend_proto_rawDesc
)

func file_remote_ethbackend_proto_rawDescGZIP() []byte {
	file_remote_ethbackend_proto_rawDescOnce.Do(func() {
		file_remote_ethbackend_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_ethbackend_proto_rawDescData)
	})
	return file_remote_ethbackend_proto_rawDescData
}

var file_remote_ethbackend_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_remote_ethbackend_proto_goTypes = []interface{}{
	(*TxRequest)(nil),         // 0: remote.TxRequest
	(*AddReply)(nil),          // 1: remote.AddReply
	(*EtherbaseRequest)(nil),  // 2: remote.EtherbaseRequest
	(*EtherbaseReply)(nil),    // 3: remote.EtherbaseReply
	(*NetVersionRequest)(nil), // 4: remote.NetVersionRequest
	(*NetVersionReply)(nil),   // 5: remote.NetVersionReply
	(*SubscribeRequest)(nil),  // 6: remote.SubscribeRequest
	(*SubscribeReply)(nil),    // 7: remote.SubscribeReply
	(*GetWorkRequest)(nil),    // 8: remote.GetWorkRequest
	(*GetWorkReply)(nil),      // 9: remote.GetWorkReply
	(*types.H256)(nil),        // 10: types.H256
	(*types.H160)(nil),        // 11: types.H160
}
var file_remote_ethbackend_proto_depIdxs = []int32{
	10, // 0: remote.AddReply.hash:type_name -> types.H256
	11, // 1: remote.EtherbaseReply.address:type_name -> types.H160
	0,  // 2: remote.ETHBACKEND.Add:input_type -> remote.TxRequest
	2,  // 3: remote.ETHBACKEND.Etherbase:input_type -> remote.EtherbaseRequest
	4,  // 4: remote.ETHBACKEND.NetVersion:input_type -> remote.NetVersionRequest
	6,  // 5: remote.ETHBACKEND.Subscribe:input_type -> remote.SubscribeRequest
	8,  // 6: remote.ETHBACKEND.GetWork:input_type -> remote.GetWorkRequest
	1,  // 7: remote.ETHBACKEND.Add:output_type -> remote.AddReply
	3,  // 8: remote.ETHBACKEND.Etherbase:output_type -> remote.EtherbaseReply
	5,  // 9: remote.ETHBACKEND.NetVersion:output_type -> remote.NetVersionReply
	7,  // 10: remote.ETHBACKEND.Subscribe:output_type -> remote.SubscribeReply
	9,  // 11: remote.ETHBACKEND.GetWork:output_type -> remote.GetWorkReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_remote_ethbackend_proto_init() }
func file_remote_ethbackend_proto_init() {
	if File_remote_ethbackend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_ethbackend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReply); i {
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
		file_remote_ethbackend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtherbaseRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtherbaseReply); i {
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
		file_remote_ethbackend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetVersionRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetVersionReply); i {
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
		file_remote_ethbackend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReply); i {
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
		file_remote_ethbackend_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkRequest); i {
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
		file_remote_ethbackend_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkReply); i {
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
			RawDescriptor: file_remote_ethbackend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_ethbackend_proto_goTypes,
		DependencyIndexes: file_remote_ethbackend_proto_depIdxs,
		MessageInfos:      file_remote_ethbackend_proto_msgTypes,
	}.Build()
	File_remote_ethbackend_proto = out.File
	file_remote_ethbackend_proto_rawDesc = nil
	file_remote_ethbackend_proto_goTypes = nil
	file_remote_ethbackend_proto_depIdxs = nil
}
