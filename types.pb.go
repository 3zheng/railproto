// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: types.proto

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

type CMDKindId int32

const (
	CMDKindId_NullType              CMDKindId = 0  //保留
	CMDKindId_IDKindLogger          CMDKindId = 8  ///日志
	CMDKindId_IDKindNetTCP          CMDKindId = 9  //对应的tcp.proto,因为怕取名net.proto和golang的标准库冲突了就麻烦了
	CMDKindId_IDKindRouter          CMDKindId = 10 ///路由类
	CMDKindId_IDKindAppFrame        CMDKindId = 11 ///App之间交互的一些通用报文，简化设计
	CMDKindId_IDKindGate            CMDKindId = 12 ///对应gate.proto
	CMDKindId_IDKindClient          CMDKindId = 13 ///客户端类
	CMDKindId_IDKindMatchClient     CMDKindId = 14 ///match客户端类
	CMDKindId_IDKindMCMS            CMDKindId = 15 /// mc与ms之间的交互
	CMDKindId_IDKindTableLogic      CMDKindId = 16 /// mp与TableLogic之间的交互
	CMDKindId_IDKindFund            CMDKindId = 17 //财富类
	CMDKindId_IDKindHallClient      CMDKindId = 18 //大厅客户端类
	CMDKindId_IDKindMatchPhase      CMDKindId = 19 //比赛过程大类
	CMDKindId_IDKindRankList        CMDKindId = 20 //比赛排名大类
	CMDKindId_IDKindList            CMDKindId = 21
	CMDKindId_IDKindMatchDB         CMDKindId = 22
	CMDKindId_IDKindPrivateInternal CMDKindId = 256  // 内部私有
	CMDKindId_IDKindGameStart       CMDKindId = 288  // 开始具体游戏
	CMDKindId_IDFinal               CMDKindId = 4094 //最后一个4095
)

// Enum value maps for CMDKindId.
var (
	CMDKindId_name = map[int32]string{
		0:    "NullType",
		8:    "IDKindLogger",
		9:    "IDKindNetTCP",
		10:   "IDKindRouter",
		11:   "IDKindAppFrame",
		12:   "IDKindGate",
		13:   "IDKindClient",
		14:   "IDKindMatchClient",
		15:   "IDKindMCMS",
		16:   "IDKindTableLogic",
		17:   "IDKindFund",
		18:   "IDKindHallClient",
		19:   "IDKindMatchPhase",
		20:   "IDKindRankList",
		21:   "IDKindList",
		22:   "IDKindMatchDB",
		256:  "IDKindPrivateInternal",
		288:  "IDKindGameStart",
		4094: "IDFinal",
	}
	CMDKindId_value = map[string]int32{
		"NullType":              0,
		"IDKindLogger":          8,
		"IDKindNetTCP":          9,
		"IDKindRouter":          10,
		"IDKindAppFrame":        11,
		"IDKindGate":            12,
		"IDKindClient":          13,
		"IDKindMatchClient":     14,
		"IDKindMCMS":            15,
		"IDKindTableLogic":      16,
		"IDKindFund":            17,
		"IDKindHallClient":      18,
		"IDKindMatchPhase":      19,
		"IDKindRankList":        20,
		"IDKindList":            21,
		"IDKindMatchDB":         22,
		"IDKindPrivateInternal": 256,
		"IDKindGameStart":       288,
		"IDFinal":               4094,
	}
)

func (x CMDKindId) Enum() *CMDKindId {
	p := new(CMDKindId)
	*p = x
	return p
}

func (x CMDKindId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDKindId) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[0].Descriptor()
}

func (CMDKindId) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[0]
}

func (x CMDKindId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMDKindId.Descriptor instead.
func (CMDKindId) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

type EnumAppId int32

const (
	EnumAppId_UnknowId    EnumAppId = 0
	EnumAppId_Send2All    EnumAppId = 1 //发送到所有
	EnumAppId_Send2AnyOne EnumAppId = 2 //发送到任意一个
)

// Enum value maps for EnumAppId.
var (
	EnumAppId_name = map[int32]string{
		0: "UnknowId",
		1: "Send2All",
		2: "Send2AnyOne",
	}
	EnumAppId_value = map[string]int32{
		"UnknowId":    0,
		"Send2All":    1,
		"Send2AnyOne": 2,
	}
)

func (x EnumAppId) Enum() *EnumAppId {
	p := new(EnumAppId)
	*p = x
	return p
}

func (x EnumAppId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumAppId) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[1].Descriptor()
}

func (EnumAppId) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[1]
}

func (x EnumAppId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumAppId.Descriptor instead.
func (EnumAppId) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

type EnumAppType int32

const (
	EnumAppType_UnknowType     EnumAppType = 0
	EnumAppType_Gate           EnumAppType = 5   //网关
	EnumAppType_Router         EnumAppType = 6   //路由中继器
	EnumAppType_Login          EnumAppType = 7   //登录管理
	EnumAppType_Online         EnumAppType = 8   //在线管理
	EnumAppType_Fund           EnumAppType = 9   //财富
	EnumAppType_List           EnumAppType = 10  //列表
	EnumAppType_FreeMatch      EnumAppType = 12  //
	EnumAppType_Match          EnumAppType = 13  //
	EnumAppType_TableLogic     EnumAppType = 14  //
	EnumAppType_MatchPhase     EnumAppType = 15  //
	EnumAppType_RankList       EnumAppType = 16  //
	EnumAppType_MatchDB        EnumAppType = 18  //
	EnumAppType_Conect_To_Gold EnumAppType = 19  //
	EnumAppType_Tool           EnumAppType = 224 //工具
	EnumAppType_Last           EnumAppType = 254 //
)

// Enum value maps for EnumAppType.
var (
	EnumAppType_name = map[int32]string{
		0:   "UnknowType",
		5:   "Gate",
		6:   "Router",
		7:   "Login",
		8:   "Online",
		9:   "Fund",
		10:  "List",
		12:  "FreeMatch",
		13:  "Match",
		14:  "TableLogic",
		15:  "MatchPhase",
		16:  "RankList",
		18:  "MatchDB",
		19:  "Conect_To_Gold",
		224: "Tool",
		254: "Last",
	}
	EnumAppType_value = map[string]int32{
		"UnknowType":     0,
		"Gate":           5,
		"Router":         6,
		"Login":          7,
		"Online":         8,
		"Fund":           9,
		"List":           10,
		"FreeMatch":      12,
		"Match":          13,
		"TableLogic":     14,
		"MatchPhase":     15,
		"RankList":       16,
		"MatchDB":        18,
		"Conect_To_Gold": 19,
		"Tool":           224,
		"Last":           254,
	}
)

func (x EnumAppType) Enum() *EnumAppType {
	p := new(EnumAppType)
	*p = x
	return p
}

func (x EnumAppType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumAppType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[2].Descriptor()
}

func (EnumAppType) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[2]
}

func (x EnumAppType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumAppType.Descriptor instead.
func (EnumAppType) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

type BaseUserInfo_UserType int32

const (
	BaseUserInfo_UNKNOW BaseUserInfo_UserType = 0  //未知
	BaseUserInfo_Normal BaseUserInfo_UserType = 1  //正常类型
	BaseUserInfo_Robot  BaseUserInfo_UserType = 10 //机器人
)

// Enum value maps for BaseUserInfo_UserType.
var (
	BaseUserInfo_UserType_name = map[int32]string{
		0:  "UNKNOW",
		1:  "Normal",
		10: "Robot",
	}
	BaseUserInfo_UserType_value = map[string]int32{
		"UNKNOW": 0,
		"Normal": 1,
		"Robot":  10,
	}
)

func (x BaseUserInfo_UserType) Enum() *BaseUserInfo_UserType {
	p := new(BaseUserInfo_UserType)
	*p = x
	return p
}

func (x BaseUserInfo_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseUserInfo_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[3].Descriptor()
}

func (BaseUserInfo_UserType) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[3]
}

func (x BaseUserInfo_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseUserInfo_UserType.Descriptor instead.
func (BaseUserInfo_UserType) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1, 0}
}

type BaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KindId     uint32 `protobuf:"varint,1,opt,name=kind_id,json=kindId,proto3" json:"kind_id,omitempty"`
	SubId      uint32 `protobuf:"varint,2,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	ConnId     uint64 `protobuf:"varint,3,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`               //conn_id
	GateConnId uint64 `protobuf:"varint,4,opt,name=gate_conn_id,json=gateConnId,proto3" json:"gate_conn_id,omitempty"` //如果这个报文来自客户端会自带gate_conn_id 和 conn_id相等
	RemoteAdd  string `protobuf:"bytes,5,opt,name=remote_add,json=remoteAdd,proto3" json:"remote_add,omitempty"`       //对端IP地址和端口
	AttApptype uint32 `protobuf:"varint,6,opt,name=att_apptype,json=attApptype,proto3" json:"att_apptype,omitempty"`   //报文关联的apptype,在收到的时候代表来源的apptype,在发出去的时候代表目标的apptype
	AttAppid   uint32 `protobuf:"varint,7,opt,name=att_appid,json=attAppid,proto3" json:"att_appid,omitempty"`         //报文关联的appid,在收到的时候代表来源的appid,在发出去的时候代表目标的appid
}

func (x *BaseInfo) Reset() {
	*x = BaseInfo{}
	mi := &file_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseInfo) ProtoMessage() {}

func (x *BaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseInfo.ProtoReflect.Descriptor instead.
func (*BaseInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *BaseInfo) GetKindId() uint32 {
	if x != nil {
		return x.KindId
	}
	return 0
}

func (x *BaseInfo) GetSubId() uint32 {
	if x != nil {
		return x.SubId
	}
	return 0
}

func (x *BaseInfo) GetConnId() uint64 {
	if x != nil {
		return x.ConnId
	}
	return 0
}

func (x *BaseInfo) GetGateConnId() uint64 {
	if x != nil {
		return x.GateConnId
	}
	return 0
}

func (x *BaseInfo) GetRemoteAdd() string {
	if x != nil {
		return x.RemoteAdd
	}
	return ""
}

func (x *BaseInfo) GetAttApptype() uint32 {
	if x != nil {
		return x.AttApptype
	}
	return 0
}

func (x *BaseInfo) GetAttAppid() uint32 {
	if x != nil {
		return x.AttAppid
	}
	return 0
}

// 基础用户信息
type BaseUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       uint64                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                               //用户ID
	GameId       uint64                `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`                                               //数字ID
	Gender       uint32                `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`                                                             //性别
	FaceId       uint32                `protobuf:"varint,7,opt,name=face_id,json=faceId,proto3" json:"face_id,omitempty"`                                               //头像id
	CustomFace   string                `protobuf:"bytes,8,opt,name=custom_face,json=customFace,proto3" json:"custom_face,omitempty"`                                    //自定义的图像地址
	NickName     string                `protobuf:"bytes,9,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`                                          //昵称
	ManagerRight uint32                `protobuf:"varint,10,opt,name=manager_right,json=managerRight,proto3" json:"manager_right,omitempty"`                            //管理权限
	UserRight    uint32                `protobuf:"varint,11,opt,name=user_right,json=userRight,proto3" json:"user_right,omitempty"`                                     //用户权限
	UserType     BaseUserInfo_UserType `protobuf:"varint,12,opt,name=user_type,json=userType,proto3,enum=protodefine.BaseUserInfo_UserType" json:"user_type,omitempty"` //用户类别
}

func (x *BaseUserInfo) Reset() {
	*x = BaseUserInfo{}
	mi := &file_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseUserInfo) ProtoMessage() {}

func (x *BaseUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseUserInfo.ProtoReflect.Descriptor instead.
func (*BaseUserInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *BaseUserInfo) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BaseUserInfo) GetGameId() uint64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *BaseUserInfo) GetGender() uint32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *BaseUserInfo) GetFaceId() uint32 {
	if x != nil {
		return x.FaceId
	}
	return 0
}

func (x *BaseUserInfo) GetCustomFace() string {
	if x != nil {
		return x.CustomFace
	}
	return ""
}

func (x *BaseUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *BaseUserInfo) GetManagerRight() uint32 {
	if x != nil {
		return x.ManagerRight
	}
	return 0
}

func (x *BaseUserInfo) GetUserRight() uint32 {
	if x != nil {
		return x.UserRight
	}
	return 0
}

func (x *BaseUserInfo) GetUserType() BaseUserInfo_UserType {
	if x != nil {
		return x.UserType
	}
	return BaseUserInfo_UNKNOW
}

// 用户当前连接信息
type UserSessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GateId     uint32 `protobuf:"varint,1,opt,name=gate_id,json=gateId,proto3" json:"gate_id,omitempty"`               //来自哪个gate
	GateConnId uint64 `protobuf:"varint,2,opt,name=gate_conn_id,json=gateConnId,proto3" json:"gate_conn_id,omitempty"` //gate 相关连接
	Client_IP  string `protobuf:"bytes,9,opt,name=client_IP,json=clientIP,proto3" json:"client_IP,omitempty"`          //客户端的IP
}

func (x *UserSessionInfo) Reset() {
	*x = UserSessionInfo{}
	mi := &file_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSessionInfo) ProtoMessage() {}

func (x *UserSessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSessionInfo.ProtoReflect.Descriptor instead.
func (*UserSessionInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *UserSessionInfo) GetGateId() uint32 {
	if x != nil {
		return x.GateId
	}
	return 0
}

func (x *UserSessionInfo) GetGateConnId() uint64 {
	if x != nil {
		return x.GateConnId
	}
	return 0
}

func (x *UserSessionInfo) GetClient_IP() string {
	if x != nil {
		return x.Client_IP
	}
	return ""
}

// 财富类型项
type FundItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Count uint64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FundItem) Reset() {
	*x = FundItem{}
	mi := &file_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FundItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundItem) ProtoMessage() {}

func (x *FundItem) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundItem.ProtoReflect.Descriptor instead.
func (*FundItem) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *FundItem) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FundItem) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FundItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x08, 0x42,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x69, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6b, 0x69, 0x6e, 0x64, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x73, 0x75, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x41, 0x70, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x74, 0x74, 0x41, 0x70, 0x70, 0x69, 0x64, 0x22,
	0xe3, 0x02, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x61,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x46, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x10, 0x0a, 0x22, 0x69, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x49, 0x50,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50,
	0x22, 0x44, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0xf5, 0x02, 0x0a, 0x09, 0x43, 0x4d, 0x44, 0x4b, 0x69,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x4e, 0x65,
	0x74, 0x54, 0x43, 0x50, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44, 0x4b, 0x69,
	0x6e, 0x64, 0x41, 0x70, 0x70, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x47, 0x61, 0x74, 0x65, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c,
	0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x0d, 0x12, 0x15,
	0x0a, 0x11, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x4d,
	0x43, 0x4d, 0x53, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x10, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x49,
	0x44, 0x4b, 0x69, 0x6e, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x49,
	0x44, 0x4b, 0x69, 0x6e, 0x64, 0x48, 0x61, 0x6c, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10,
	0x12, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x10, 0x13, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44, 0x4b, 0x69, 0x6e,
	0x64, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x14, 0x12, 0x0e, 0x0a, 0x0a, 0x49,
	0x44, 0x4b, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x15, 0x12, 0x11, 0x0a, 0x0d, 0x49,
	0x44, 0x4b, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x42, 0x10, 0x16, 0x12, 0x1a,
	0x0a, 0x15, 0x49, 0x44, 0x4b, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x80, 0x02, 0x12, 0x14, 0x0a, 0x0f, 0x49, 0x44,
	0x4b, 0x69, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0xa0, 0x02,
	0x12, 0x0c, 0x0a, 0x07, 0x49, 0x44, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x10, 0xfe, 0x1f, 0x2a, 0x38,
	0x0a, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x49, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x65, 0x6e,
	0x64, 0x32, 0x41, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x32,
	0x41, 0x6e, 0x79, 0x4f, 0x6e, 0x65, 0x10, 0x02, 0x2a, 0xdd, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x75,
	0x6d, 0x41, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x61, 0x74, 0x65,
	0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x10, 0x06, 0x12, 0x09,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x64, 0x10, 0x09, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x72, 0x65,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69,
	0x63, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x10,
	0x10, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x42, 0x10, 0x12, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x54, 0x6f, 0x5f, 0x47, 0x6f, 0x6c, 0x64,
	0x10, 0x13, 0x12, 0x09, 0x0a, 0x04, 0x54, 0x6f, 0x6f, 0x6c, 0x10, 0xe0, 0x01, 0x12, 0x09, 0x0a,
	0x04, 0x4c, 0x61, 0x73, 0x74, 0x10, 0xfe, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x33, 0x7a, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x72, 0x61,
	0x69, 0x6c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_proto_goTypes = []any{
	(CMDKindId)(0),             // 0: protodefine.CMDKindId
	(EnumAppId)(0),             // 1: protodefine.EnumAppId
	(EnumAppType)(0),           // 2: protodefine.EnumAppType
	(BaseUserInfo_UserType)(0), // 3: protodefine.BaseUserInfo.UserType
	(*BaseInfo)(nil),           // 4: protodefine.BaseInfo
	(*BaseUserInfo)(nil),       // 5: protodefine.BaseUserInfo
	(*UserSessionInfo)(nil),    // 6: protodefine.UserSessionInfo
	(*FundItem)(nil),           // 7: protodefine.FundItem
}
var file_types_proto_depIdxs = []int32{
	3, // 0: protodefine.BaseUserInfo.user_type:type_name -> protodefine.BaseUserInfo.UserType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		EnumInfos:         file_types_proto_enumTypes,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
