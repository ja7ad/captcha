// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: captcha.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Captcha struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// check_challenge check captcha challenge in rpc method
	CheckChallenge bool `protobuf:"varint,1,opt,name=check_challenge,json=checkChallenge,proto3" json:"check_challenge,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Captcha) Reset() {
	*x = Captcha{}
	mi := &file_captcha_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Captcha) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Captcha) ProtoMessage() {}

func (x *Captcha) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Captcha.ProtoReflect.Descriptor instead.
func (*Captcha) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *Captcha) GetCheckChallenge() bool {
	if x != nil {
		return x.CheckChallenge
	}
	return false
}

var file_captcha_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Captcha)(nil),
		Field:         7564324,
		Name:          "captcha.captcha",
		Tag:           "bytes,7564324,opt,name=captcha",
		Filename:      "captcha.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional captcha.Captcha captcha = 7564324;
	E_Captcha = &file_captcha_proto_extTypes[0]
)

var File_captcha_proto protoreflect.FileDescriptor

var file_captcha_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x07, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x3a, 0x4d,
	0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa4, 0xd8, 0xcd, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x37, 0x61,
	0x64, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_captcha_proto_rawDescOnce sync.Once
	file_captcha_proto_rawDescData []byte
)

func file_captcha_proto_rawDescGZIP() []byte {
	file_captcha_proto_rawDescOnce.Do(func() {
		file_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_captcha_proto_rawDesc), len(file_captcha_proto_rawDesc)))
	})
	return file_captcha_proto_rawDescData
}

var file_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_captcha_proto_goTypes = []any{
	(*Captcha)(nil),                    // 0: captcha.Captcha
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_captcha_proto_depIdxs = []int32{
	1, // 0: captcha.captcha:extendee -> google.protobuf.MethodOptions
	0, // 1: captcha.captcha:type_name -> captcha.Captcha
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_captcha_proto_init() }
func file_captcha_proto_init() {
	if File_captcha_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_captcha_proto_rawDesc), len(file_captcha_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_captcha_proto_goTypes,
		DependencyIndexes: file_captcha_proto_depIdxs,
		MessageInfos:      file_captcha_proto_msgTypes,
		ExtensionInfos:    file_captcha_proto_extTypes,
	}.Build()
	File_captcha_proto = out.File
	file_captcha_proto_goTypes = nil
	file_captcha_proto_depIdxs = nil
}
