// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: p2psentry/control.proto

package core

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type InboundMessageId int32

const (
	InboundMessageId_NewBlockHashes InboundMessageId = 0
	InboundMessageId_BlockHeaders   InboundMessageId = 1
	InboundMessageId_BlockBodies    InboundMessageId = 2
	InboundMessageId_NewBlock       InboundMessageId = 3
	InboundMessageId_NodeData       InboundMessageId = 4
)

// Enum value maps for InboundMessageId.
var (
	InboundMessageId_name = map[int32]string{
		0: "NewBlockHashes",
		1: "BlockHeaders",
		2: "BlockBodies",
		3: "NewBlock",
		4: "NodeData",
	}
	InboundMessageId_value = map[string]int32{
		"NewBlockHashes": 0,
		"BlockHeaders":   1,
		"BlockBodies":    2,
		"NewBlock":       3,
		"NodeData":       4,
	}
)

func (x InboundMessageId) Enum() *InboundMessageId {
	p := new(InboundMessageId)
	*p = x
	return p
}

func (x InboundMessageId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InboundMessageId) Descriptor() protoreflect.EnumDescriptor {
	return file_p2psentry_control_proto_enumTypes[0].Descriptor()
}

func (InboundMessageId) Type() protoreflect.EnumType {
	return &file_p2psentry_control_proto_enumTypes[0]
}

func (x InboundMessageId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InboundMessageId.Descriptor instead.
func (InboundMessageId) EnumDescriptor() ([]byte, []int) {
	return file_p2psentry_control_proto_rawDescGZIP(), []int{0}
}

type InboundMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     InboundMessageId `protobuf:"varint,1,opt,name=id,proto3,enum=control.InboundMessageId" json:"id,omitempty"`
	Data   []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	PeerId []byte           `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (x *InboundMessage) Reset() {
	*x = InboundMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2psentry_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InboundMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InboundMessage) ProtoMessage() {}

func (x *InboundMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2psentry_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InboundMessage.ProtoReflect.Descriptor instead.
func (*InboundMessage) Descriptor() ([]byte, []int) {
	return file_p2psentry_control_proto_rawDescGZIP(), []int{0}
}

func (x *InboundMessage) GetId() InboundMessageId {
	if x != nil {
		return x.Id
	}
	return InboundMessageId_NewBlockHashes
}

func (x *InboundMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InboundMessage) GetPeerId() []byte {
	if x != nil {
		return x.PeerId
	}
	return nil
}

type Forks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genesis []byte   `protobuf:"bytes,1,opt,name=genesis,proto3" json:"genesis,omitempty"`
	Forks   []uint64 `protobuf:"varint,2,rep,packed,name=forks,proto3" json:"forks,omitempty"`
}

func (x *Forks) Reset() {
	*x = Forks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2psentry_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forks) ProtoMessage() {}

func (x *Forks) ProtoReflect() protoreflect.Message {
	mi := &file_p2psentry_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forks.ProtoReflect.Descriptor instead.
func (*Forks) Descriptor() ([]byte, []int) {
	return file_p2psentry_control_proto_rawDescGZIP(), []int{1}
}

func (x *Forks) GetGenesis() []byte {
	if x != nil {
		return x.Genesis
	}
	return nil
}

func (x *Forks) GetForks() []uint64 {
	if x != nil {
		return x.Forks
	}
	return nil
}

type StatusData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId       uint64 `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	TotalDifficulty []byte `protobuf:"bytes,2,opt,name=total_difficulty,json=totalDifficulty,proto3" json:"total_difficulty,omitempty"`
	BestHash        []byte `protobuf:"bytes,3,opt,name=best_hash,json=bestHash,proto3" json:"best_hash,omitempty"`
	ForkData        *Forks `protobuf:"bytes,4,opt,name=fork_data,json=forkData,proto3" json:"fork_data,omitempty"`
}

func (x *StatusData) Reset() {
	*x = StatusData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2psentry_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusData) ProtoMessage() {}

func (x *StatusData) ProtoReflect() protoreflect.Message {
	mi := &file_p2psentry_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusData.ProtoReflect.Descriptor instead.
func (*StatusData) Descriptor() ([]byte, []int) {
	return file_p2psentry_control_proto_rawDescGZIP(), []int{2}
}

func (x *StatusData) GetNetworkId() uint64 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *StatusData) GetTotalDifficulty() []byte {
	if x != nil {
		return x.TotalDifficulty
	}
	return nil
}

func (x *StatusData) GetBestHash() []byte {
	if x != nil {
		return x.BestHash
	}
	return nil
}

func (x *StatusData) GetForkData() *Forks {
	if x != nil {
		return x.ForkData
	}
	return nil
}

var File_p2psentry_control_proto protoreflect.FileDescriptor

var file_p2psentry_control_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x32, 0x70, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x0e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x05, 0x46, 0x6f, 0x72,
	0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x66, 0x6f, 0x72,
	0x6b, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x65, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x62, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6b,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x08, 0x66, 0x6f, 0x72,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x65, 0x0a, 0x10, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x65, 0x77,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x10, 0x04, 0x32, 0x8d, 0x01, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x48, 0x0a, 0x15, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x49, 0x6e, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_p2psentry_control_proto_rawDescOnce sync.Once
	file_p2psentry_control_proto_rawDescData = file_p2psentry_control_proto_rawDesc
)

func file_p2psentry_control_proto_rawDescGZIP() []byte {
	file_p2psentry_control_proto_rawDescOnce.Do(func() {
		file_p2psentry_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2psentry_control_proto_rawDescData)
	})
	return file_p2psentry_control_proto_rawDescData
}

var file_p2psentry_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_p2psentry_control_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_p2psentry_control_proto_goTypes = []interface{}{
	(InboundMessageId)(0),  // 0: control.InboundMessageId
	(*InboundMessage)(nil), // 1: control.InboundMessage
	(*Forks)(nil),          // 2: control.Forks
	(*StatusData)(nil),     // 3: control.StatusData
	(*empty.Empty)(nil),    // 4: google.protobuf.Empty
}
var file_p2psentry_control_proto_depIdxs = []int32{
	0, // 0: control.InboundMessage.id:type_name -> control.InboundMessageId
	2, // 1: control.StatusData.fork_data:type_name -> control.Forks
	1, // 2: control.Control.ForwardInboundMessage:input_type -> control.InboundMessage
	4, // 3: control.Control.GetStatus:input_type -> google.protobuf.Empty
	4, // 4: control.Control.ForwardInboundMessage:output_type -> google.protobuf.Empty
	3, // 5: control.Control.GetStatus:output_type -> control.StatusData
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_p2psentry_control_proto_init() }
func file_p2psentry_control_proto_init() {
	if File_p2psentry_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2psentry_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InboundMessage); i {
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
		file_p2psentry_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forks); i {
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
		file_p2psentry_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusData); i {
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
			RawDescriptor: file_p2psentry_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_p2psentry_control_proto_goTypes,
		DependencyIndexes: file_p2psentry_control_proto_depIdxs,
		EnumInfos:         file_p2psentry_control_proto_enumTypes,
		MessageInfos:      file_p2psentry_control_proto_msgTypes,
	}.Build()
	File_p2psentry_control_proto = out.File
	file_p2psentry_control_proto_rawDesc = nil
	file_p2psentry_control_proto_goTypes = nil
	file_p2psentry_control_proto_depIdxs = nil
}
