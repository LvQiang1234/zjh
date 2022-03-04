// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: account.proto

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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	PassWord *string `protobuf:"bytes,2,req,name=PassWord" json:"PassWord,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *LoginReq) GetPassWord() string {
	if x != nil && x.PassWord != nil {
		return *x.PassWord
	}
	return ""
}

type LoginAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId *int32     `protobuf:"varint,1,opt,name=PlayerId" json:"PlayerId,omitempty"`
	Error    *ErrorType `protobuf:"varint,5,req,name=Error,enum=ErrorType" json:"Error,omitempty"`
}

func (x *LoginAck) Reset() {
	*x = LoginAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAck) ProtoMessage() {}

func (x *LoginAck) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAck.ProtoReflect.Descriptor instead.
func (*LoginAck) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *LoginAck) GetPlayerId() int32 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

func (x *LoginAck) GetError() ErrorType {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ErrorType_Success
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	PassWord *string `protobuf:"bytes,2,req,name=PassWord" json:"PassWord,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *RegisterReq) GetPassWord() string {
	if x != nil && x.PassWord != nil {
		return *x.PassWord
	}
	return ""
}

type RegisterAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *ErrorType `protobuf:"varint,1,req,name=Error,enum=ErrorType" json:"Error,omitempty"`
}

func (x *RegisterAck) Reset() {
	*x = RegisterAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAck) ProtoMessage() {}

func (x *RegisterAck) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAck.ProtoReflect.Descriptor instead.
func (*RegisterAck) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterAck) GetError() ErrorType {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ErrorType_Success
}

type PlayerInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayerInfoReq) Reset() {
	*x = PlayerInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfoReq) ProtoMessage() {}

func (x *PlayerInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfoReq.ProtoReflect.Descriptor instead.
func (*PlayerInfoReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

type PlayerInfoAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      *ErrorType `protobuf:"varint,1,req,name=Error,enum=ErrorType" json:"Error,omitempty"`
	PlayerName *string    `protobuf:"bytes,2,opt,name=PlayerName" json:"PlayerName,omitempty"`
	Coins      *int32     `protobuf:"varint,3,opt,name=coins" json:"coins,omitempty"`
	ImageId    *int32     `protobuf:"varint,4,opt,name=ImageId" json:"ImageId,omitempty"`
}

func (x *PlayerInfoAck) Reset() {
	*x = PlayerInfoAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfoAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfoAck) ProtoMessage() {}

func (x *PlayerInfoAck) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfoAck.ProtoReflect.Descriptor instead.
func (*PlayerInfoAck) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerInfoAck) GetError() ErrorType {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ErrorType_Success
}

func (x *PlayerInfoAck) GetPlayerName() string {
	if x != nil && x.PlayerName != nil {
		return *x.PlayerName
	}
	return ""
}

func (x *PlayerInfoAck) GetCoins() int32 {
	if x != nil && x.Coins != nil {
		return *x.Coins
	}
	return 0
}

func (x *PlayerInfoAck) GetImageId() int32 {
	if x != nil && x.ImageId != nil {
		return *x.ImageId
	}
	return 0
}

type UpdateCoinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num *int32 `protobuf:"varint,1,req,name=Num" json:"Num,omitempty"`
}

func (x *UpdateCoinReq) Reset() {
	*x = UpdateCoinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoinReq) ProtoMessage() {}

func (x *UpdateCoinReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoinReq.ProtoReflect.Descriptor instead.
func (*UpdateCoinReq) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCoinReq) GetNum() int32 {
	if x != nil && x.Num != nil {
		return *x.Num
	}
	return 0
}

type UpdateCoinAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   *ErrorType `protobuf:"varint,1,req,name=Error,enum=ErrorType" json:"Error,omitempty"`
	CoinNum *int32     `protobuf:"varint,2,req,name=CoinNum" json:"CoinNum,omitempty"`
}

func (x *UpdateCoinAck) Reset() {
	*x = UpdateCoinAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoinAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoinAck) ProtoMessage() {}

func (x *UpdateCoinAck) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoinAck.ProtoReflect.Descriptor instead.
func (*UpdateCoinAck) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCoinAck) GetError() ErrorType {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ErrorType_Success
}

func (x *UpdateCoinAck) GetCoinNum() int32 {
	if x != nil && x.CoinNum != nil {
		return *x.CoinNum
	}
	return 0
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x48, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72,
	0x64, 0x22, 0x2f, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x6b,
	0x12, 0x20, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x22, 0x4b, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07,
	0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_account_proto_goTypes = []interface{}{
	(*LoginReq)(nil),      // 0: LoginReq
	(*LoginAck)(nil),      // 1: LoginAck
	(*RegisterReq)(nil),   // 2: RegisterReq
	(*RegisterAck)(nil),   // 3: RegisterAck
	(*PlayerInfoReq)(nil), // 4: PlayerInfoReq
	(*PlayerInfoAck)(nil), // 5: PlayerInfoAck
	(*UpdateCoinReq)(nil), // 6: UpdateCoinReq
	(*UpdateCoinAck)(nil), // 7: UpdateCoinAck
	(ErrorType)(0),        // 8: ErrorType
}
var file_account_proto_depIdxs = []int32{
	8, // 0: LoginAck.Error:type_name -> ErrorType
	8, // 1: RegisterAck.Error:type_name -> ErrorType
	8, // 2: PlayerInfoAck.Error:type_name -> ErrorType
	8, // 3: UpdateCoinAck.Error:type_name -> ErrorType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	file_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAck); i {
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
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAck); i {
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
		file_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfoReq); i {
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
		file_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfoAck); i {
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
		file_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoinReq); i {
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
		file_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoinAck); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
