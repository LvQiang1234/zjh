// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: error.proto

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

type ErrorType int32

const (
	ErrorType_Success ErrorType = 0
	ErrorType_Fail    ErrorType = 1
)

// Enum value maps for ErrorType.
var (
	ErrorType_name = map[int32]string{
		0: "Success",
		1: "Fail",
	}
	ErrorType_value = map[string]int32{
		"Success": 0,
		"Fail":    1,
	}
)

func (x ErrorType) Enum() *ErrorType {
	p := new(ErrorType)
	*p = x
	return p
}

func (x ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[0].Descriptor()
}

func (ErrorType) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[0]
}

func (x ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ErrorType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ErrorType(num)
	return nil
}

// Deprecated: Use ErrorType.Descriptor instead.
func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x22, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_proto_goTypes = []interface{}{
	(ErrorType)(0), // 0: ErrorType
}
var file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}