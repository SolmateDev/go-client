// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: session.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "github.com/SolmateDev/go-client/proto/base"
	files "github.com/SolmateDev/go-client/proto/v1/files"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expire int64 `protobuf:"varint,1,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *ApiKeyRequest) Reset() {
	*x = ApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyRequest) ProtoMessage() {}

func (x *ApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyRequest.ProtoReflect.Descriptor instead.
func (*ApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKeyRequest) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type ApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Expire int64  `protobuf:"varint,2,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *ApiKeyResponse) Reset() {
	*x = ApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyResponse) ProtoMessage() {}

func (x *ApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyResponse.ProtoReflect.Descriptor instead.
func (*ApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{1}
}

func (x *ApiKeyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ApiKeyResponse) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *base.UserInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Summary string         `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Avatar  *files.Info    `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetInfo() *base.UserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Profile) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Profile) GetAvatar() *files.Info {
	if x != nil {
		return x.Avatar
	}
	return nil
}

type ApiKeySessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ApiKeySessionRequest) Reset() {
	*x = ApiKeySessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeySessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeySessionRequest) ProtoMessage() {}

func (x *ApiKeySessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeySessionRequest.ProtoReflect.Descriptor instead.
func (*ApiKeySessionRequest) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{3}
}

func (x *ApiKeySessionRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ApiKeySessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt    string           `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Detail *SessionResponse `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *ApiKeySessionResponse) Reset() {
	*x = ApiKeySessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeySessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeySessionResponse) ProtoMessage() {}

func (x *ApiKeySessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeySessionResponse.ProtoReflect.Descriptor instead.
func (*ApiKeySessionResponse) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{4}
}

func (x *ApiKeySessionResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *ApiKeySessionResponse) GetDetail() *SessionResponse {
	if x != nil {
		return x.Detail
	}
	return nil
}

type SessionCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt      string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Domain   string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	SourceIp string `protobuf:"bytes,3,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
}

func (x *SessionCheck) Reset() {
	*x = SessionCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCheck) ProtoMessage() {}

func (x *SessionCheck) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCheck.ProtoReflect.Descriptor instead.
func (*SessionCheck) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{5}
}

func (x *SessionCheck) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *SessionCheck) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SessionCheck) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

type SessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rights []string `protobuf:"bytes,2,rep,name=rights,proto3" json:"rights,omitempty"`
	Expire int64    `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *SessionResponse) Reset() {
	*x = SessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionResponse) ProtoMessage() {}

func (x *SessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionResponse.ProtoReflect.Descriptor instead.
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{6}
}

func (x *SessionResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SessionResponse) GetRights() []string {
	if x != nil {
		return x.Rights
	}
	return nil
}

func (x *SessionResponse) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

var File_session_proto protoreflect.FileDescriptor

var file_session_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x27, 0x0a, 0x0d, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x3a, 0x0a, 0x0e, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x6e, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x25,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x28, 0x0a, 0x14, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x58, 0x0a, 0x15, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x55, 0x0a, 0x0c, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70,
	0x22, 0x5a, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x32, 0x8c, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x12,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xa3, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12,
	0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x10,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x61, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_proto_rawDescOnce sync.Once
	file_session_proto_rawDescData = file_session_proto_rawDesc
)

func file_session_proto_rawDescGZIP() []byte {
	file_session_proto_rawDescOnce.Do(func() {
		file_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_proto_rawDescData)
	})
	return file_session_proto_rawDescData
}

var file_session_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_session_proto_goTypes = []interface{}{
	(*ApiKeyRequest)(nil),         // 0: auth.ApiKeyRequest
	(*ApiKeyResponse)(nil),        // 1: auth.ApiKeyResponse
	(*Profile)(nil),               // 2: auth.Profile
	(*ApiKeySessionRequest)(nil),  // 3: auth.ApiKeySessionRequest
	(*ApiKeySessionResponse)(nil), // 4: auth.ApiKeySessionResponse
	(*SessionCheck)(nil),          // 5: auth.SessionCheck
	(*SessionResponse)(nil),       // 6: auth.SessionResponse
	(*base.UserInfo)(nil),         // 7: base.UserInfo
	(*files.Info)(nil),            // 8: filesv1.Info
	(*base.Empty)(nil),            // 9: base.Empty
	(*base.UserLookup)(nil),       // 10: base.UserLookup
}
var file_session_proto_depIdxs = []int32{
	7,  // 0: auth.Profile.info:type_name -> base.UserInfo
	8,  // 1: auth.Profile.avatar:type_name -> filesv1.Info
	6,  // 2: auth.ApiKeySessionResponse.detail:type_name -> auth.SessionResponse
	3,  // 3: auth.Session.CreateByApiKey:input_type -> auth.ApiKeySessionRequest
	5,  // 4: auth.Session.Check:input_type -> auth.SessionCheck
	0,  // 5: auth.User.CreateApiKey:input_type -> auth.ApiKeyRequest
	9,  // 6: auth.User.ViewApiKey:input_type -> base.Empty
	10, // 7: auth.User.Lookup:input_type -> base.UserLookup
	4,  // 8: auth.Session.CreateByApiKey:output_type -> auth.ApiKeySessionResponse
	6,  // 9: auth.Session.Check:output_type -> auth.SessionResponse
	1,  // 10: auth.User.CreateApiKey:output_type -> auth.ApiKeyResponse
	1,  // 11: auth.User.ViewApiKey:output_type -> auth.ApiKeyResponse
	2,  // 12: auth.User.Lookup:output_type -> auth.Profile
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_session_proto_init() }
func file_session_proto_init() {
	if File_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyRequest); i {
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
		file_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyResponse); i {
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
		file_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeySessionRequest); i {
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
		file_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeySessionResponse); i {
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
		file_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCheck); i {
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
		file_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionResponse); i {
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
			RawDescriptor: file_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_session_proto_goTypes,
		DependencyIndexes: file_session_proto_depIdxs,
		MessageInfos:      file_session_proto_msgTypes,
	}.Build()
	File_session_proto = out.File
	file_session_proto_rawDesc = nil
	file_session_proto_goTypes = nil
	file_session_proto_depIdxs = nil
}
