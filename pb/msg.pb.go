// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: msg.proto

package pb

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

type MsgId int32

const (
	MsgId_LoginRequest          MsgId = 0
	MsgId_LoginResponse         MsgId = 1
	MsgId_RegisterRequest       MsgId = 2
	MsgId_RegisterAckResponse   MsgId = 3
	MsgId_GetPlayerInfoRequest  MsgId = 4
	MsgId_GetPlayerInfoResponse MsgId = 5
	MsgId_UpdateCoinRequest     MsgId = 6
	MsgId_UpdateCoinResponse    MsgId = 7
)

// Enum value maps for MsgId.
var (
	MsgId_name = map[int32]string{
		0: "LoginRequest",
		1: "LoginResponse",
		2: "RegisterRequest",
		3: "RegisterAckResponse",
		4: "GetPlayerInfoRequest",
		5: "GetPlayerInfoResponse",
		6: "UpdateCoinRequest",
		7: "UpdateCoinResponse",
	}
	MsgId_value = map[string]int32{
		"LoginRequest":          0,
		"LoginResponse":         1,
		"RegisterRequest":       2,
		"RegisterAckResponse":   3,
		"GetPlayerInfoRequest":  4,
		"GetPlayerInfoResponse": 5,
		"UpdateCoinRequest":     6,
		"UpdateCoinResponse":    7,
	}
)

func (x MsgId) Enum() *MsgId {
	p := new(MsgId)
	*p = x
	return p
}

func (x MsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MsgId) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MsgId) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MsgId(num)
	return nil
}

// Deprecated: Use MsgId.Descriptor instead.
func (MsgId) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type MsgPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId    *MsgId `protobuf:"varint,1,req,name=MsgId,enum=MsgId" json:"MsgId,omitempty"`
	PlayerId *int32 `protobuf:"varint,2,req,name=PlayerId" json:"PlayerId,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=Data" json:"Data,omitempty"`
}

func (x *MsgPacket) Reset() {
	*x = MsgPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPacket) ProtoMessage() {}

func (x *MsgPacket) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPacket.ProtoReflect.Descriptor instead.
func (*MsgPacket) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *MsgPacket) GetMsgId() MsgId {
	if x != nil && x.MsgId != nil {
		return *x.MsgId
	}
	return MsgId_LoginRequest
}

func (x *MsgPacket) GetPlayerId() int32 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

func (x *MsgPacket) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x09, 0x4d,
	0x73, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x52,
	0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x2a, 0xbe, 0x01, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x04, 0x12, 0x19, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x06, 0x12,
	0x16, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x07, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_msg_proto_goTypes = []interface{}{
	(MsgId)(0),        // 0: MsgId
	(*MsgPacket)(nil), // 1: MsgPacket
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: MsgPacket.MsgId:type_name -> MsgId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPacket); i {
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
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
