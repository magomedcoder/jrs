// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.1
// source: v1/user.proto

package v1

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserItem) Reset() {
	*x = UserItem{}
	mi := &file_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserItem) ProtoMessage() {}

func (x *UserItem) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserItem.ProtoReflect.Descriptor instead.
func (*UserItem) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserItem) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*UserItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	mi := &file_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUsersResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetUsersResponse) GetItems() []*UserItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_v1_user_proto protoreflect.FileDescriptor

var file_v1_user_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x67, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x6a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x4e, 0x9a, 0x84, 0x9e, 0x03, 0x49, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6d, 0x69, 0x6e,
	0x3d, 0x33, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x4e, 0x9a, 0x84, 0x9e, 0x03, 0x49, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6d, 0x69, 0x6e, 0x3d, 0x38, 0x22, 0x20,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x56, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x42, 0x9a, 0x84, 0x9e, 0x03, 0x3d, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6d, 0x69, 0x6e, 0x3d, 0x32, 0x22, 0x20,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x82,
	0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x9a,
	0x84, 0x9e, 0x03, 0x0f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e,
	0x03, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x49, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x67, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x67, 0x6f, 0x6d,
	0x65, 0x64, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_user_proto_rawDescOnce sync.Once
	file_v1_user_proto_rawDescData = file_v1_user_proto_rawDesc
)

func file_v1_user_proto_rawDescGZIP() []byte {
	file_v1_user_proto_rawDescOnce.Do(func() {
		file_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_user_proto_rawDescData)
	})
	return file_v1_user_proto_rawDescData
}

var file_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil), // 0: gskeleton.http.v1.user.CreateUserRequest
	(*GetUserResponse)(nil),   // 1: gskeleton.http.v1.user.GetUserResponse
	(*UserItem)(nil),          // 2: gskeleton.http.v1.user.UserItem
	(*GetUsersResponse)(nil),  // 3: gskeleton.http.v1.user.GetUsersResponse
}
var file_v1_user_proto_depIdxs = []int32{
	2, // 0: gskeleton.http.v1.user.GetUsersResponse.items:type_name -> gskeleton.http.v1.user.UserItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_user_proto_init() }
func file_v1_user_proto_init() {
	if File_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_user_proto_goTypes,
		DependencyIndexes: file_v1_user_proto_depIdxs,
		MessageInfos:      file_v1_user_proto_msgTypes,
	}.Build()
	File_v1_user_proto = out.File
	file_v1_user_proto_rawDesc = nil
	file_v1_user_proto_goTypes = nil
	file_v1_user_proto_depIdxs = nil
}
