// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: account.proto

package proto

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

	Error *ErrorType `protobuf:"varint,1,req,name=Error,enum=ErrorType" json:"Error,omitempty"`
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

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x2c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x2f, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f,
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

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_account_proto_goTypes = []interface{}{
	(*LoginReq)(nil),    // 0: LoginReq
	(*LoginAck)(nil),    // 1: LoginAck
	(*RegisterReq)(nil), // 2: RegisterReq
	(*RegisterAck)(nil), // 3: RegisterAck
	(ErrorType)(0),      // 4: ErrorType
}
var file_account_proto_depIdxs = []int32{
	4, // 0: LoginAck.Error:type_name -> ErrorType
	4, // 1: RegisterAck.Error:type_name -> ErrorType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
