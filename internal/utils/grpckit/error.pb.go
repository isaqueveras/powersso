// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: protos/error.proto

package grpckit

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *string `protobuf:"bytes,1,opt,name=Location,proto3,oneof" json:"Location,omitempty"`
	RawError *string `protobuf:"bytes,2,opt,name=RawError,proto3,oneof" json:"RawError,omitempty"`
	Error    *string `protobuf:"bytes,3,opt,name=Error,proto3,oneof" json:"Error,omitempty"`
	Code     *uint64 `protobuf:"varint,4,opt,name=Code,proto3,oneof" json:"Code,omitempty"`
}

func (x *ErrorGRPC) Reset() {
	*x = ErrorGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorGRPC) ProtoMessage() {}

func (x *ErrorGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_protos_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorGRPC.ProtoReflect.Descriptor instead.
func (*ErrorGRPC) Descriptor() ([]byte, []int) {
	return file_protos_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorGRPC) GetLocation() string {
	if x != nil && x.Location != nil {
		return *x.Location
	}
	return ""
}

func (x *ErrorGRPC) GetRawError() string {
	if x != nil && x.RawError != nil {
		return *x.RawError
	}
	return ""
}

func (x *ErrorGRPC) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *ErrorGRPC) GetCode() uint64 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

var File_protos_error_proto protoreflect.FileDescriptor

var file_protos_error_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x6b, 0x69, 0x74, 0x22, 0xae, 0x01,
	0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x47, 0x52, 0x50, 0x43, 0x12, 0x1f, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x52, 0x61, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x08, 0x52, 0x61, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x52, 0x61, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x61,
	0x71, 0x75, 0x65, 0x76, 0x65, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x2d, 0x73,
	0x73, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6b, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_error_proto_rawDescOnce sync.Once
	file_protos_error_proto_rawDescData = file_protos_error_proto_rawDesc
)

func file_protos_error_proto_rawDescGZIP() []byte {
	file_protos_error_proto_rawDescOnce.Do(func() {
		file_protos_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_error_proto_rawDescData)
	})
	return file_protos_error_proto_rawDescData
}

var file_protos_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_error_proto_goTypes = []interface{}{
	(*ErrorGRPC)(nil), // 0: grpckit.ErrorGRPC
}
var file_protos_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_error_proto_init() }
func file_protos_error_proto_init() {
	if File_protos_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorGRPC); i {
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
	file_protos_error_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_error_proto_goTypes,
		DependencyIndexes: file_protos_error_proto_depIdxs,
		MessageInfos:      file_protos_error_proto_msgTypes,
	}.Build()
	File_protos_error_proto = out.File
	file_protos_error_proto_rawDesc = nil
	file_protos_error_proto_goTypes = nil
	file_protos_error_proto_depIdxs = nil
}
