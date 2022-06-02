// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: password.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "github.com/SolmateDev/go-client/proto/base"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmailDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *EmailDomain) Reset() {
	*x = EmailDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailDomain) ProtoMessage() {}

func (x *EmailDomain) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailDomain.ProtoReflect.Descriptor instead.
func (*EmailDomain) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{0}
}

func (x *EmailDomain) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailDomain) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type ParseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *ParseRequest) Reset() {
	*x = ParseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseRequest) ProtoMessage() {}

func (x *ParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseRequest.ProtoReflect.Descriptor instead.
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{1}
}

func (x *ParseRequest) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type ParseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Expire int64    `protobuf:"varint,2,opt,name=expire,proto3" json:"expire,omitempty"`
	Rights []string `protobuf:"bytes,3,rep,name=rights,proto3" json:"rights,omitempty"`
}

func (x *ParseResponse) Reset() {
	*x = ParseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseResponse) ProtoMessage() {}

func (x *ParseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseResponse.ProtoReflect.Descriptor instead.
func (*ParseResponse) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{2}
}

func (x *ParseResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ParseResponse) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *ParseResponse) GetRights() []string {
	if x != nil {
		return x.Rights
	}
	return nil
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     *EmailDomain `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password  string       `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	SmsNumber string       `protobuf:"bytes,3,opt,name=sms_number,json=smsNumber,proto3" json:"sms_number,omitempty"`
	// send the signed artifact, not the plain token
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRequest) GetEmail() *EmailDomain {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetSmsNumber() string {
	if x != nil {
		return x.SmsNumber
	}
	return ""
}

func (x *RegisterRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SecondFactorLocalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *EmailDomain `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SecondFactorLocalRequest) Reset() {
	*x = SecondFactorLocalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondFactorLocalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondFactorLocalRequest) ProtoMessage() {}

func (x *SecondFactorLocalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondFactorLocalRequest.ProtoReflect.Descriptor instead.
func (*SecondFactorLocalRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{4}
}

func (x *SecondFactorLocalRequest) GetEmail() *EmailDomain {
	if x != nil {
		return x.Email
	}
	return nil
}

type SecondFactorLocalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// just feed this value directly into the registration process
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// user scans this into Google Authenticator
	AuthenticatorToken string `protobuf:"bytes,2,opt,name=authenticator_token,json=authenticatorToken,proto3" json:"authenticator_token,omitempty"`
}

func (x *SecondFactorLocalResponse) Reset() {
	*x = SecondFactorLocalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondFactorLocalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondFactorLocalResponse) ProtoMessage() {}

func (x *SecondFactorLocalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondFactorLocalResponse.ProtoReflect.Descriptor instead.
func (*SecondFactorLocalResponse) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{5}
}

func (x *SecondFactorLocalResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SecondFactorLocalResponse) GetAuthenticatorToken() string {
	if x != nil {
		return x.AuthenticatorToken
	}
	return ""
}

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     *EmailDomain `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	EmailName string       `protobuf:"bytes,2,opt,name=email_name,json=emailName,proto3" json:"email_name,omitempty"`
	// CODE, COMPANY_NAME are replaced with values internally set
	SubjectTemplate  string `protobuf:"bytes,3,opt,name=subject_template,json=subjectTemplate,proto3" json:"subject_template,omitempty"`
	BodyTemplate     string `protobuf:"bytes,4,opt,name=body_template,json=bodyTemplate,proto3" json:"body_template,omitempty"`
	BodyHtmlTemplate string `protobuf:"bytes,5,opt,name=body_html_template,json=bodyHtmlTemplate,proto3" json:"body_html_template,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{6}
}

func (x *SendEmailRequest) GetEmail() *EmailDomain {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *SendEmailRequest) GetEmailName() string {
	if x != nil {
		return x.EmailName
	}
	return ""
}

func (x *SendEmailRequest) GetSubjectTemplate() string {
	if x != nil {
		return x.SubjectTemplate
	}
	return ""
}

func (x *SendEmailRequest) GetBodyTemplate() string {
	if x != nil {
		return x.BodyTemplate
	}
	return ""
}

func (x *SendEmailRequest) GetBodyHtmlTemplate() string {
	if x != nil {
		return x.BodyHtmlTemplate
	}
	return ""
}

type SendSMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*SendSMSRequest_SmsNumber
	//	*SendSMSRequest_Email
	Data isSendSMSRequest_Data `protobuf_oneof:"data"`
	// CODE, COMPANY_NAME are replaced with values internally set
	BodyTemplate string `protobuf:"bytes,3,opt,name=body_template,json=bodyTemplate,proto3" json:"body_template,omitempty"`
}

func (x *SendSMSRequest) Reset() {
	*x = SendSMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSRequest) ProtoMessage() {}

func (x *SendSMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSRequest.ProtoReflect.Descriptor instead.
func (*SendSMSRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{7}
}

func (m *SendSMSRequest) GetData() isSendSMSRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *SendSMSRequest) GetSmsNumber() string {
	if x, ok := x.GetData().(*SendSMSRequest_SmsNumber); ok {
		return x.SmsNumber
	}
	return ""
}

func (x *SendSMSRequest) GetEmail() *EmailDomain {
	if x, ok := x.GetData().(*SendSMSRequest_Email); ok {
		return x.Email
	}
	return nil
}

func (x *SendSMSRequest) GetBodyTemplate() string {
	if x != nil {
		return x.BodyTemplate
	}
	return ""
}

type isSendSMSRequest_Data interface {
	isSendSMSRequest_Data()
}

type SendSMSRequest_SmsNumber struct {
	// use this when the user is first registering
	SmsNumber string `protobuf:"bytes,1,opt,name=sms_number,json=smsNumber,proto3,oneof"`
}

type SendSMSRequest_Email struct {
	// use this when the user is attempting to log in
	Email *EmailDomain `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

func (*SendSMSRequest_SmsNumber) isSendSMSRequest_Data() {}

func (*SendSMSRequest_Email) isSendSMSRequest_Data() {}

type OneTimeCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpireTime int64 `protobuf:"varint,1,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *OneTimeCodeResponse) Reset() {
	*x = OneTimeCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneTimeCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneTimeCodeResponse) ProtoMessage() {}

func (x *OneTimeCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneTimeCodeResponse.ProtoReflect.Descriptor instead.
func (*OneTimeCodeResponse) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{8}
}

func (x *OneTimeCodeResponse) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

type VerifySMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *SendSMSRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Code    string          `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifySMSRequest) Reset() {
	*x = VerifySMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifySMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifySMSRequest) ProtoMessage() {}

func (x *VerifySMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifySMSRequest.ProtoReflect.Descriptor instead.
func (*VerifySMSRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{9}
}

func (x *VerifySMSRequest) GetRequest() *SendSMSRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *VerifySMSRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifyEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *EmailDomain `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code  string       `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifyEmailRequest) Reset() {
	*x = VerifyEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailRequest) ProtoMessage() {}

func (x *VerifyEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailRequest.ProtoReflect.Descriptor instead.
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{10}
}

func (x *VerifyEmailRequest) GetEmail() *EmailDomain {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *VerifyEmailRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    *EmailDomain `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string       `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{11}
}

func (x *LoginRequest) GetEmail() *EmailDomain {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{12}
}

func (x *LoginResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

var File_password_proto protoreflect.FileDescriptor

var file_password_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0x20, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77,
	0x74, 0x22, 0x58, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6d, 0x73, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x18, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x62,
	0x0a, 0x19, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xd8, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f,
	0x64, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x6f, 0x64,
	0x79, 0x48, 0x74, 0x6d, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x89, 0x01,
	0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x13, 0x4f, 0x6e, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x56, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x4d, 0x53, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x51, 0x0a, 0x12, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x53, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x21, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6a, 0x77, 0x74, 0x32, 0xf4, 0x03, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x30, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x53, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x32, 0x46, 0x41, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4f,
	0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x73, 0x65, 0x4a, 0x57, 0x54,
	0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x70, 0x61, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_password_proto_rawDescOnce sync.Once
	file_password_proto_rawDescData = file_password_proto_rawDesc
)

func file_password_proto_rawDescGZIP() []byte {
	file_password_proto_rawDescOnce.Do(func() {
		file_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_password_proto_rawDescData)
	})
	return file_password_proto_rawDescData
}

var file_password_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_password_proto_goTypes = []interface{}{
	(*EmailDomain)(nil),               // 0: auth.EmailDomain
	(*ParseRequest)(nil),              // 1: auth.ParseRequest
	(*ParseResponse)(nil),             // 2: auth.ParseResponse
	(*RegisterRequest)(nil),           // 3: auth.RegisterRequest
	(*SecondFactorLocalRequest)(nil),  // 4: auth.SecondFactorLocalRequest
	(*SecondFactorLocalResponse)(nil), // 5: auth.SecondFactorLocalResponse
	(*SendEmailRequest)(nil),          // 6: auth.SendEmailRequest
	(*SendSMSRequest)(nil),            // 7: auth.SendSMSRequest
	(*OneTimeCodeResponse)(nil),       // 8: auth.OneTimeCodeResponse
	(*VerifySMSRequest)(nil),          // 9: auth.VerifySMSRequest
	(*VerifyEmailRequest)(nil),        // 10: auth.VerifyEmailRequest
	(*LoginRequest)(nil),              // 11: auth.LoginRequest
	(*LoginResponse)(nil),             // 12: auth.LoginResponse
	(*base.Empty)(nil),                // 13: base.Empty
}
var file_password_proto_depIdxs = []int32{
	0,  // 0: auth.RegisterRequest.email:type_name -> auth.EmailDomain
	0,  // 1: auth.SecondFactorLocalRequest.email:type_name -> auth.EmailDomain
	0,  // 2: auth.SendEmailRequest.email:type_name -> auth.EmailDomain
	0,  // 3: auth.SendSMSRequest.email:type_name -> auth.EmailDomain
	7,  // 4: auth.VerifySMSRequest.request:type_name -> auth.SendSMSRequest
	0,  // 5: auth.VerifyEmailRequest.email:type_name -> auth.EmailDomain
	0,  // 6: auth.LoginRequest.email:type_name -> auth.EmailDomain
	3,  // 7: auth.Auth.Register:input_type -> auth.RegisterRequest
	4,  // 8: auth.Auth.Create2FALocal:input_type -> auth.SecondFactorLocalRequest
	7,  // 9: auth.Auth.SendSmsCode:input_type -> auth.SendSMSRequest
	9,  // 10: auth.Auth.VerifySmsCode:input_type -> auth.VerifySMSRequest
	6,  // 11: auth.Auth.SendEmailCode:input_type -> auth.SendEmailRequest
	10, // 12: auth.Auth.VerifyEmailCode:input_type -> auth.VerifyEmailRequest
	11, // 13: auth.Auth.Login:input_type -> auth.LoginRequest
	1,  // 14: auth.Auth.ParseJWT:input_type -> auth.ParseRequest
	13, // 15: auth.Auth.Register:output_type -> base.Empty
	5,  // 16: auth.Auth.Create2FALocal:output_type -> auth.SecondFactorLocalResponse
	8,  // 17: auth.Auth.SendSmsCode:output_type -> auth.OneTimeCodeResponse
	13, // 18: auth.Auth.VerifySmsCode:output_type -> base.Empty
	8,  // 19: auth.Auth.SendEmailCode:output_type -> auth.OneTimeCodeResponse
	13, // 20: auth.Auth.VerifyEmailCode:output_type -> base.Empty
	12, // 21: auth.Auth.Login:output_type -> auth.LoginResponse
	2,  // 22: auth.Auth.ParseJWT:output_type -> auth.ParseResponse
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_password_proto_init() }
func file_password_proto_init() {
	if File_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailDomain); i {
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
		file_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseRequest); i {
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
		file_password_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseResponse); i {
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
		file_password_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_password_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondFactorLocalRequest); i {
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
		file_password_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondFactorLocalResponse); i {
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
		file_password_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_password_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSRequest); i {
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
		file_password_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneTimeCodeResponse); i {
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
		file_password_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifySMSRequest); i {
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
		file_password_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailRequest); i {
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
		file_password_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_password_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
	file_password_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*SendSMSRequest_SmsNumber)(nil),
		(*SendSMSRequest_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_password_proto_goTypes,
		DependencyIndexes: file_password_proto_depIdxs,
		MessageInfos:      file_password_proto_msgTypes,
	}.Build()
	File_password_proto = out.File
	file_password_proto_rawDesc = nil
	file_password_proto_goTypes = nil
	file_password_proto_depIdxs = nil
}
