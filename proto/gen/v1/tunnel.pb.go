// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: def/v1/tunnel.proto

package v1

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

type PostTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *PostTweetRequest) Reset() {
	*x = PostTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_v1_tunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTweetRequest) ProtoMessage() {}

func (x *PostTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_def_v1_tunnel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTweetRequest.ProtoReflect.Descriptor instead.
func (*PostTweetRequest) Descriptor() ([]byte, []int) {
	return file_def_v1_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *PostTweetRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type PostTweetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PostTweetResponse) Reset() {
	*x = PostTweetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_v1_tunnel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTweetResponse) ProtoMessage() {}

func (x *PostTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_def_v1_tunnel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTweetResponse.ProtoReflect.Descriptor instead.
func (*PostTweetResponse) Descriptor() ([]byte, []int) {
	return file_def_v1_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *PostTweetResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_def_v1_tunnel_proto protoreflect.FileDescriptor

var file_def_v1_tunnel_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x65, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x2f, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x43, 0x0a, 0x0d, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x11,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_def_v1_tunnel_proto_rawDescOnce sync.Once
	file_def_v1_tunnel_proto_rawDescData = file_def_v1_tunnel_proto_rawDesc
)

func file_def_v1_tunnel_proto_rawDescGZIP() []byte {
	file_def_v1_tunnel_proto_rawDescOnce.Do(func() {
		file_def_v1_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(file_def_v1_tunnel_proto_rawDescData)
	})
	return file_def_v1_tunnel_proto_rawDescData
}

var file_def_v1_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_def_v1_tunnel_proto_goTypes = []interface{}{
	(*PostTweetRequest)(nil),  // 0: PostTweetRequest
	(*PostTweetResponse)(nil), // 1: PostTweetResponse
}
var file_def_v1_tunnel_proto_depIdxs = []int32{
	0, // 0: TunnelService.PostTweet:input_type -> PostTweetRequest
	1, // 1: TunnelService.PostTweet:output_type -> PostTweetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_def_v1_tunnel_proto_init() }
func file_def_v1_tunnel_proto_init() {
	if File_def_v1_tunnel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_def_v1_tunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTweetRequest); i {
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
		file_def_v1_tunnel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTweetResponse); i {
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
			RawDescriptor: file_def_v1_tunnel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_def_v1_tunnel_proto_goTypes,
		DependencyIndexes: file_def_v1_tunnel_proto_depIdxs,
		MessageInfos:      file_def_v1_tunnel_proto_msgTypes,
	}.Build()
	File_def_v1_tunnel_proto = out.File
	file_def_v1_tunnel_proto_rawDesc = nil
	file_def_v1_tunnel_proto_goTypes = nil
	file_def_v1_tunnel_proto_depIdxs = nil
}