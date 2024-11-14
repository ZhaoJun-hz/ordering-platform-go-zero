// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: rpc/sys/sys.proto

package sysclient

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

type InitApiRouteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handle string `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"` // 方法名
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`   // 方法描述
	Path   string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`     // 请求路径
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`     // 类型 1 系统 2 业务
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"` // 请求方式
}

func (x *InitApiRouteData) Reset() {
	*x = InitApiRouteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitApiRouteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitApiRouteData) ProtoMessage() {}

func (x *InitApiRouteData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitApiRouteData.ProtoReflect.Descriptor instead.
func (*InitApiRouteData) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{0}
}

func (x *InitApiRouteData) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *InitApiRouteData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InitApiRouteData) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *InitApiRouteData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InitApiRouteData) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type InitApiReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*InitApiRouteData `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *InitApiReq) Reset() {
	*x = InitApiReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitApiReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitApiReq) ProtoMessage() {}

func (x *InitApiReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitApiReq.ProtoReflect.Descriptor instead.
func (*InitApiReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{1}
}

func (x *InitApiReq) GetList() []*InitApiRouteData {
	if x != nil {
		return x.List
	}
	return nil
}

type InitApiResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitApiResp) Reset() {
	*x = InitApiResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitApiResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitApiResp) ProtoMessage() {}

func (x *InitApiResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitApiResp.ProtoReflect.Descriptor instead.
func (*InitApiResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{2}
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[3]
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
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReq.ProtoReflect.Descriptor instead.
func (*InfoReq) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{5}
}

func (x *InfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type InfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar   string   `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Username string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Roles    []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	UserId   string   `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Desc     string   `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	HomePath string   `protobuf:"bytes,6,opt,name=homePath,proto3" json:"homePath,omitempty"`
}

func (x *InfoResp) Reset() {
	*x = InfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sys_sys_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sys_sys_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResp.ProtoReflect.Descriptor instead.
func (*InfoResp) Descriptor() ([]byte, []int) {
	return file_rpc_sys_sys_proto_rawDescGZIP(), []int{6}
}

func (x *InfoResp) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *InfoResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *InfoResp) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *InfoResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InfoResp) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *InfoResp) GetHomePath() string {
	if x != nil {
		return x.HomePath
	}
	return ""
}

var File_rpc_sys_sys_proto protoreflect.FileDescriptor

var file_rpc_sys_sys_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x79, 0x73, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x80,
	0x01, 0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x70, 0x69, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3d, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x12,
	0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x79, 0x73, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x70,
	0x69, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x3d, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x21, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x32, 0x47, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x70, 0x69, 0x12, 0x15,
	0x2e, 0x73, 0x79, 0x73, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x73, 0x79, 0x73, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x32, 0x76, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x73, 0x79, 0x73, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x79, 0x73,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x33, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x73,
	0x79, 0x73, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x13, 0x2e, 0x73, 0x79, 0x73, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x73, 0x79, 0x73, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sys_sys_proto_rawDescOnce sync.Once
	file_rpc_sys_sys_proto_rawDescData = file_rpc_sys_sys_proto_rawDesc
)

func file_rpc_sys_sys_proto_rawDescGZIP() []byte {
	file_rpc_sys_sys_proto_rawDescOnce.Do(func() {
		file_rpc_sys_sys_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sys_sys_proto_rawDescData)
	})
	return file_rpc_sys_sys_proto_rawDescData
}

var file_rpc_sys_sys_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_sys_sys_proto_goTypes = []any{
	(*InitApiRouteData)(nil), // 0: sysclient.InitApiRouteData
	(*InitApiReq)(nil),       // 1: sysclient.InitApiReq
	(*InitApiResp)(nil),      // 2: sysclient.InitApiResp
	(*LoginReq)(nil),         // 3: sysclient.LoginReq
	(*LoginResp)(nil),        // 4: sysclient.loginResp
	(*InfoReq)(nil),          // 5: sysclient.InfoReq
	(*InfoResp)(nil),         // 6: sysclient.InfoResp
}
var file_rpc_sys_sys_proto_depIdxs = []int32{
	0, // 0: sysclient.InitApiReq.list:type_name -> sysclient.InitApiRouteData
	1, // 1: sysclient.BaseService.InitApi:input_type -> sysclient.InitApiReq
	3, // 2: sysclient.UserService.Login:input_type -> sysclient.LoginReq
	5, // 3: sysclient.UserService.UserInfo:input_type -> sysclient.InfoReq
	2, // 4: sysclient.BaseService.InitApi:output_type -> sysclient.InitApiResp
	4, // 5: sysclient.UserService.Login:output_type -> sysclient.loginResp
	6, // 6: sysclient.UserService.UserInfo:output_type -> sysclient.InfoResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_sys_sys_proto_init() }
func file_rpc_sys_sys_proto_init() {
	if File_rpc_sys_sys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sys_sys_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InitApiRouteData); i {
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
		file_rpc_sys_sys_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InitApiReq); i {
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
		file_rpc_sys_sys_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InitApiResp); i {
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
		file_rpc_sys_sys_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
		file_rpc_sys_sys_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResp); i {
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
		file_rpc_sys_sys_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*InfoReq); i {
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
		file_rpc_sys_sys_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*InfoResp); i {
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
			RawDescriptor: file_rpc_sys_sys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rpc_sys_sys_proto_goTypes,
		DependencyIndexes: file_rpc_sys_sys_proto_depIdxs,
		MessageInfos:      file_rpc_sys_sys_proto_msgTypes,
	}.Build()
	File_rpc_sys_sys_proto = out.File
	file_rpc_sys_sys_proto_rawDesc = nil
	file_rpc_sys_sys_proto_goTypes = nil
	file_rpc_sys_sys_proto_depIdxs = nil
}
