// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: api/grpc/proto/auth.proto

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

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_auth_proto_rawDescGZIP(), []int{0}
}

type Login_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Login_Request) Reset() {
	*x = Login_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login_Request) ProtoMessage() {}

func (x *Login_Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login_Request.ProtoReflect.Descriptor instead.
func (*Login_Request) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_auth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Login_Request) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Login_Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Login_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (x *Login_Response) Reset() {
	*x = Login_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login_Response) ProtoMessage() {}

func (x *Login_Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login_Response.ProtoReflect.Descriptor instead.
func (*Login_Response) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_auth_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Login_Response) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_api_grpc_proto_auth_proto protoreflect.FileDescriptor

var file_api_grpc_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x22, 0x78, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x41, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x2c, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x43, 0x0a, 0x0b, 0x41,
	0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_proto_auth_proto_rawDescOnce sync.Once
	file_api_grpc_proto_auth_proto_rawDescData = file_api_grpc_proto_auth_proto_rawDesc
)

func file_api_grpc_proto_auth_proto_rawDescGZIP() []byte {
	file_api_grpc_proto_auth_proto_rawDescOnce.Do(func() {
		file_api_grpc_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_proto_auth_proto_rawDescData)
	})
	return file_api_grpc_proto_auth_proto_rawDescData
}

var file_api_grpc_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_grpc_proto_auth_proto_goTypes = []interface{}{
	(*Login)(nil),          // 0: auth.Login
	(*Login_Request)(nil),  // 1: auth.Login.Request
	(*Login_Response)(nil), // 2: auth.Login.Response
}
var file_api_grpc_proto_auth_proto_depIdxs = []int32{
	1, // 0: auth.AuthService.Login:input_type -> auth.Login.Request
	2, // 1: auth.AuthService.Login:output_type -> auth.Login.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_proto_auth_proto_init() }
func file_api_grpc_proto_auth_proto_init() {
	if File_api_grpc_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
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
		file_api_grpc_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login_Request); i {
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
		file_api_grpc_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login_Response); i {
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
			RawDescriptor: file_api_grpc_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_proto_auth_proto_goTypes,
		DependencyIndexes: file_api_grpc_proto_auth_proto_depIdxs,
		MessageInfos:      file_api_grpc_proto_auth_proto_msgTypes,
	}.Build()
	File_api_grpc_proto_auth_proto = out.File
	file_api_grpc_proto_auth_proto_rawDesc = nil
	file_api_grpc_proto_auth_proto_goTypes = nil
	file_api_grpc_proto_auth_proto_depIdxs = nil
}
