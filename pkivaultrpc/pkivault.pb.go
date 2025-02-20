// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.27.1
// source: pkivault.proto

package pkivaultrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type UserRole int32

const (
	UserRole_ADMIN UserRole = 0
	UserRole_USER  UserRole = 1
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "ADMIN",
		1: "USER",
	}
	UserRole_value = map[string]int32{
		"ADMIN": 0,
		"USER":  1,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_pkivault_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_pkivault_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{0}
}

type PagingQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Search        string                 `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Order         string                 `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	Sort          string                 `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Limit         int64                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64                  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PagingQuery) Reset() {
	*x = PagingQuery{}
	mi := &file_pkivault_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PagingQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingQuery) ProtoMessage() {}

func (x *PagingQuery) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingQuery.ProtoReflect.Descriptor instead.
func (*PagingQuery) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{0}
}

func (x *PagingQuery) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *PagingQuery) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *PagingQuery) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *PagingQuery) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PagingQuery) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// request messages
type UserDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserDeleteRequest) Reset() {
	*x = UserDeleteRequest{}
	mi := &file_pkivault_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRequest) ProtoMessage() {}

func (x *UserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRequest.ProtoReflect.Descriptor instead.
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{1}
}

func (x *UserDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserUpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Role          UserRole               `protobuf:"varint,5,opt,name=role,proto3,enum=pkivaultrpc.UserRole" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserUpdateRequest) Reset() {
	*x = UserUpdateRequest{}
	mi := &file_pkivault_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRequest) ProtoMessage() {}

func (x *UserUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{2}
}

func (x *UserUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserUpdateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserUpdateRequest) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_ADMIN
}

type UserCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	IsOAuth       bool                   `protobuf:"varint,4,opt,name=isOAuth,proto3" json:"isOAuth,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	mi := &file_pkivault_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{3}
}

func (x *UserCreateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCreateRequest) GetIsOAuth() bool {
	if x != nil {
		return x.IsOAuth
	}
	return false
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_pkivault_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PagingQuery   *PagingQuery           `protobuf:"bytes,1,opt,name=pagingQuery,proto3" json:"pagingQuery,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_pkivault_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{5}
}

func (x *UserRequest) GetPagingQuery() *PagingQuery {
	if x != nil {
		return x.PagingQuery
	}
	return nil
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// response messages
type UserList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserInList          `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Total         int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserList) Reset() {
	*x = UserList{}
	mi := &file_pkivault_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{6}
}

func (x *UserList) GetUsers() []*UserInList {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UserList) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UserInList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Languages     []*UserLanguageProfile `protobuf:"bytes,5,rep,name=languages,proto3" json:"languages,omitempty"`
	Role          string                 `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	IsOAuth       bool                   `protobuf:"varint,7,opt,name=isOAuth,proto3" json:"isOAuth,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInList) Reset() {
	*x = UserInList{}
	mi := &file_pkivault_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInList) ProtoMessage() {}

func (x *UserInList) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInList.ProtoReflect.Descriptor instead.
func (*UserInList) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{7}
}

func (x *UserInList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInList) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInList) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInList) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInList) GetLanguages() []*UserLanguageProfile {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *UserInList) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserInList) GetIsOAuth() bool {
	if x != nil {
		return x.IsOAuth
	}
	return false
}

type UserLanguageProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Language      string                 `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Actions       []string               `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLanguageProfile) Reset() {
	*x = UserLanguageProfile{}
	mi := &file_pkivault_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLanguageProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLanguageProfile) ProtoMessage() {}

func (x *UserLanguageProfile) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLanguageProfile.ProtoReflect.Descriptor instead.
func (*UserLanguageProfile) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{8}
}

func (x *UserLanguageProfile) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *UserLanguageProfile) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	CreatedId     int64                  `protobuf:"varint,2,opt,name=createdId,proto3" json:"createdId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_pkivault_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkivault_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkivault_proto_rawDescGZIP(), []int{9}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetCreatedId() int64 {
	if x != nil {
		return x.CreatedId
	}
	return 0
}

var File_pkivault_proto protoreflect.FileDescriptor

var file_pkivault_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x22, 0x7d, 0x0a,
	0x0b, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x23, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x9c, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x7b, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x22, 0x20, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x7b, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0b, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4f, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xd8, 0x01,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3e, 0x0a, 0x09, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x73, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x22, 0x4b, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x42, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x2a, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x32, 0xef, 0x02, 0x0a, 0x0f, 0x70,
	0x6b, 0x69, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x6b, 0x69,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x1b, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x70,
	0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x6b, 0x69, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x6f, 0x6c, 0x61, 0x72,
	0x61, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x69, 0x2d, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x69, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pkivault_proto_rawDescOnce sync.Once
	file_pkivault_proto_rawDescData []byte
)

func file_pkivault_proto_rawDescGZIP() []byte {
	file_pkivault_proto_rawDescOnce.Do(func() {
		file_pkivault_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkivault_proto_rawDesc), len(file_pkivault_proto_rawDesc)))
	})
	return file_pkivault_proto_rawDescData
}

var file_pkivault_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkivault_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkivault_proto_goTypes = []any{
	(UserRole)(0),               // 0: pkivaultrpc.UserRole
	(*PagingQuery)(nil),         // 1: pkivaultrpc.PagingQuery
	(*UserDeleteRequest)(nil),   // 2: pkivaultrpc.UserDeleteRequest
	(*UserUpdateRequest)(nil),   // 3: pkivaultrpc.UserUpdateRequest
	(*UserCreateRequest)(nil),   // 4: pkivaultrpc.UserCreateRequest
	(*GetUserRequest)(nil),      // 5: pkivaultrpc.GetUserRequest
	(*UserRequest)(nil),         // 6: pkivaultrpc.UserRequest
	(*UserList)(nil),            // 7: pkivaultrpc.UserList
	(*UserInList)(nil),          // 8: pkivaultrpc.UserInList
	(*UserLanguageProfile)(nil), // 9: pkivaultrpc.UserLanguageProfile
	(*Response)(nil),            // 10: pkivaultrpc.Response
}
var file_pkivault_proto_depIdxs = []int32{
	0,  // 0: pkivaultrpc.UserUpdateRequest.role:type_name -> pkivaultrpc.UserRole
	1,  // 1: pkivaultrpc.UserRequest.pagingQuery:type_name -> pkivaultrpc.PagingQuery
	8,  // 2: pkivaultrpc.UserList.users:type_name -> pkivaultrpc.UserInList
	9,  // 3: pkivaultrpc.UserInList.languages:type_name -> pkivaultrpc.UserLanguageProfile
	6,  // 4: pkivaultrpc.pkiVaultService.GetUsers:input_type -> pkivaultrpc.UserRequest
	5,  // 5: pkivaultrpc.pkiVaultService.GetUserProfile:input_type -> pkivaultrpc.GetUserRequest
	4,  // 6: pkivaultrpc.pkiVaultService.CreateUser:input_type -> pkivaultrpc.UserCreateRequest
	3,  // 7: pkivaultrpc.pkiVaultService.UpdateUser:input_type -> pkivaultrpc.UserUpdateRequest
	2,  // 8: pkivaultrpc.pkiVaultService.DeleteUser:input_type -> pkivaultrpc.UserDeleteRequest
	7,  // 9: pkivaultrpc.pkiVaultService.GetUsers:output_type -> pkivaultrpc.UserList
	8,  // 10: pkivaultrpc.pkiVaultService.GetUserProfile:output_type -> pkivaultrpc.UserInList
	10, // 11: pkivaultrpc.pkiVaultService.CreateUser:output_type -> pkivaultrpc.Response
	10, // 12: pkivaultrpc.pkiVaultService.UpdateUser:output_type -> pkivaultrpc.Response
	10, // 13: pkivaultrpc.pkiVaultService.DeleteUser:output_type -> pkivaultrpc.Response
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkivault_proto_init() }
func file_pkivault_proto_init() {
	if File_pkivault_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkivault_proto_rawDesc), len(file_pkivault_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkivault_proto_goTypes,
		DependencyIndexes: file_pkivault_proto_depIdxs,
		EnumInfos:         file_pkivault_proto_enumTypes,
		MessageInfos:      file_pkivault_proto_msgTypes,
	}.Build()
	File_pkivault_proto = out.File
	file_pkivault_proto_goTypes = nil
	file_pkivault_proto_depIdxs = nil
}
