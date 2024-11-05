// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: tcp.proto

package protodefine

import (
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

type CMDID_Tcp int32

const (
	CMDID_Tcp_IDUnknowTp        CMDID_Tcp = 0 //0保留
	CMDID_Tcp_IDTCPTransferMsg  CMDID_Tcp = 1 //TCP消息传输报文
	CMDID_Tcp_IDTCPSessionCome  CMDID_Tcp = 2 //新建了一个TCP会话
	CMDID_Tcp_IDTCPSessionClose CMDID_Tcp = 3 //关闭了一个TCP会话
	CMDID_Tcp_IDTCPSessionKick  CMDID_Tcp = 4 //主动要求关闭一个TCP会话
	CMDID_Tcp_IDLastTp          CMDID_Tcp = 100
)

// Enum value maps for CMDID_Tcp.
var (
	CMDID_Tcp_name = map[int32]string{
		0:   "IDUnknowTp",
		1:   "IDTCPTransferMsg",
		2:   "IDTCPSessionCome",
		3:   "IDTCPSessionClose",
		4:   "IDTCPSessionKick",
		100: "IDLastTp",
	}
	CMDID_Tcp_value = map[string]int32{
		"IDUnknowTp":        0,
		"IDTCPTransferMsg":  1,
		"IDTCPSessionCome":  2,
		"IDTCPSessionClose": 3,
		"IDTCPSessionKick":  4,
		"IDLastTp":          100,
	}
)

func (x CMDID_Tcp) Enum() *CMDID_Tcp {
	p := new(CMDID_Tcp)
	*p = x
	return p
}

func (x CMDID_Tcp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDID_Tcp) Descriptor() protoreflect.EnumDescriptor {
	return file_tcp_proto_enumTypes[0].Descriptor()
}

func (CMDID_Tcp) Type() protoreflect.EnumType {
	return &file_tcp_proto_enumTypes[0]
}

func (x CMDID_Tcp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMDID_Tcp.Descriptor instead.
func (CMDID_Tcp) EnumDescriptor() ([]byte, []int) {
	return file_tcp_proto_rawDescGZIP(), []int{0}
}

type TCPTransferMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base       *BaseInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DataKindId uint32    `protobuf:"varint,2,opt,name=data_kind_id,json=dataKindId,proto3" json:"data_kind_id,omitempty"`
	DataSubId  uint32    `protobuf:"varint,3,opt,name=data_sub_id,json=dataSubId,proto3" json:"data_sub_id,omitempty"`
	Data       []byte    `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TCPTransferMsg) Reset() {
	*x = TCPTransferMsg{}
	mi := &file_tcp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TCPTransferMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPTransferMsg) ProtoMessage() {}

func (x *TCPTransferMsg) ProtoReflect() protoreflect.Message {
	mi := &file_tcp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPTransferMsg.ProtoReflect.Descriptor instead.
func (*TCPTransferMsg) Descriptor() ([]byte, []int) {
	return file_tcp_proto_rawDescGZIP(), []int{0}
}

func (x *TCPTransferMsg) GetBase() *BaseInfo {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TCPTransferMsg) GetDataKindId() uint32 {
	if x != nil {
		return x.DataKindId
	}
	return 0
}

func (x *TCPTransferMsg) GetDataSubId() uint32 {
	if x != nil {
		return x.DataSubId
	}
	return 0
}

func (x *TCPTransferMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type TCPSessionCome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *BaseInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *TCPSessionCome) Reset() {
	*x = TCPSessionCome{}
	mi := &file_tcp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TCPSessionCome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPSessionCome) ProtoMessage() {}

func (x *TCPSessionCome) ProtoReflect() protoreflect.Message {
	mi := &file_tcp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPSessionCome.ProtoReflect.Descriptor instead.
func (*TCPSessionCome) Descriptor() ([]byte, []int) {
	return file_tcp_proto_rawDescGZIP(), []int{1}
}

func (x *TCPSessionCome) GetBase() *BaseInfo {
	if x != nil {
		return x.Base
	}
	return nil
}

type TCPSessionClose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *BaseInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *TCPSessionClose) Reset() {
	*x = TCPSessionClose{}
	mi := &file_tcp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TCPSessionClose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPSessionClose) ProtoMessage() {}

func (x *TCPSessionClose) ProtoReflect() protoreflect.Message {
	mi := &file_tcp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPSessionClose.ProtoReflect.Descriptor instead.
func (*TCPSessionClose) Descriptor() ([]byte, []int) {
	return file_tcp_proto_rawDescGZIP(), []int{2}
}

func (x *TCPSessionClose) GetBase() *BaseInfo {
	if x != nil {
		return x.Base
	}
	return nil
}

type TCPSessionKick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *BaseInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *TCPSessionKick) Reset() {
	*x = TCPSessionKick{}
	mi := &file_tcp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TCPSessionKick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPSessionKick) ProtoMessage() {}

func (x *TCPSessionKick) ProtoReflect() protoreflect.Message {
	mi := &file_tcp_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPSessionKick.ProtoReflect.Descriptor instead.
func (*TCPSessionKick) Descriptor() ([]byte, []int) {
	return file_tcp_proto_rawDescGZIP(), []int{3}
}

func (x *TCPSessionKick) GetBase() *BaseInfo {
	if x != nil {
		return x.Base
	}
	return nil
}

var File_tcp_proto protoreflect.FileDescriptor

var file_tcp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x54, 0x43, 0x50, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x29, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4b,
	0x69, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x75,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x75, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x0e, 0x54, 0x43, 0x50,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x0f, 0x54, 0x43, 0x50, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x0e, 0x54, 0x43, 0x50, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4b, 0x69, 0x63, 0x6b, 0x12, 0x29, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x2a, 0x82, 0x01, 0x0a, 0x09, 0x43, 0x4d, 0x44, 0x49, 0x44, 0x5f, 0x54, 0x63, 0x70, 0x12,
	0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x54, 0x70, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x49, 0x44, 0x54, 0x43, 0x50, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x54, 0x43, 0x50, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x65, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49,
	0x44, 0x54, 0x43, 0x50, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x54, 0x43, 0x50, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4b, 0x69, 0x63, 0x6b, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x44, 0x4c, 0x61,
	0x73, 0x74, 0x54, 0x70, 0x10, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x33, 0x7a, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x69, 0x6c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tcp_proto_rawDescOnce sync.Once
	file_tcp_proto_rawDescData = file_tcp_proto_rawDesc
)

func file_tcp_proto_rawDescGZIP() []byte {
	file_tcp_proto_rawDescOnce.Do(func() {
		file_tcp_proto_rawDescData = protoimpl.X.CompressGZIP(file_tcp_proto_rawDescData)
	})
	return file_tcp_proto_rawDescData
}

var file_tcp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tcp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tcp_proto_goTypes = []any{
	(CMDID_Tcp)(0),          // 0: protodefine.CMDID_Tcp
	(*TCPTransferMsg)(nil),  // 1: protodefine.TCPTransferMsg
	(*TCPSessionCome)(nil),  // 2: protodefine.TCPSessionCome
	(*TCPSessionClose)(nil), // 3: protodefine.TCPSessionClose
	(*TCPSessionKick)(nil),  // 4: protodefine.TCPSessionKick
	(*BaseInfo)(nil),        // 5: protodefine.BaseInfo
}
var file_tcp_proto_depIdxs = []int32{
	5, // 0: protodefine.TCPTransferMsg.base:type_name -> protodefine.BaseInfo
	5, // 1: protodefine.TCPSessionCome.base:type_name -> protodefine.BaseInfo
	5, // 2: protodefine.TCPSessionClose.base:type_name -> protodefine.BaseInfo
	5, // 3: protodefine.TCPSessionKick.base:type_name -> protodefine.BaseInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tcp_proto_init() }
func file_tcp_proto_init() {
	if File_tcp_proto != nil {
		return
	}
	file_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tcp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tcp_proto_goTypes,
		DependencyIndexes: file_tcp_proto_depIdxs,
		EnumInfos:         file_tcp_proto_enumTypes,
		MessageInfos:      file_tcp_proto_msgTypes,
	}.Build()
	File_tcp_proto = out.File
	file_tcp_proto_rawDesc = nil
	file_tcp_proto_goTypes = nil
	file_tcp_proto_depIdxs = nil
}