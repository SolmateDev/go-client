// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: filesv1.proto

package files

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "github.com/SolmateDev/go-client/proto/base"
	kyc "github.com/SolmateDev/go-client/proto/kyc"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResidencyProofType int32

const (
	ResidencyProofType_UTILITY_BILL   ResidencyProofType = 0
	ResidencyProofType_PHONE_BILL     ResidencyProofType = 1
	ResidencyProofType_BANK_STATEMENT ResidencyProofType = 2
)

// Enum value maps for ResidencyProofType.
var (
	ResidencyProofType_name = map[int32]string{
		0: "UTILITY_BILL",
		1: "PHONE_BILL",
		2: "BANK_STATEMENT",
	}
	ResidencyProofType_value = map[string]int32{
		"UTILITY_BILL":   0,
		"PHONE_BILL":     1,
		"BANK_STATEMENT": 2,
	}
)

func (x ResidencyProofType) Enum() *ResidencyProofType {
	p := new(ResidencyProofType)
	*p = x
	return p
}

func (x ResidencyProofType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResidencyProofType) Descriptor() protoreflect.EnumDescriptor {
	return file_filesv1_proto_enumTypes[0].Descriptor()
}

func (ResidencyProofType) Type() protoreflect.EnumType {
	return &file_filesv1_proto_enumTypes[0]
}

func (x ResidencyProofType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResidencyProofType.Descriptor instead.
func (ResidencyProofType) EnumDescriptor() ([]byte, []int) {
	return file_filesv1_proto_rawDescGZIP(), []int{0}
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commit    *base.Commit `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	Valid     bool         `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Hash      []byte       `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	RequestId int64        `protobuf:"varint,4,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesv1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_filesv1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_filesv1_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetCommit() *base.Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *Info) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *Info) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Info) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

// this returns a list of documents (valid and invalid) that the user has uploaded; the request_id is in Info
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identification []*Status_IdentificationPair `protobuf:"bytes,1,rep,name=identification,proto3" json:"identification,omitempty"`
	Residency      []*Status_ResidencyProof     `protobuf:"bytes,2,rep,name=residency,proto3" json:"residency,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesv1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_filesv1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_filesv1_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetIdentification() []*Status_IdentificationPair {
	if x != nil {
		return x.Identification
	}
	return nil
}

func (x *Status) GetResidency() []*Status_ResidencyProof {
	if x != nil {
		return x.Residency
	}
	return nil
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*FileRequest_NationalId
	//	*FileRequest_ResidencyProof
	Request isFileRequest_Request `protobuf_oneof:"request"`
	Note    string                `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesv1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesv1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_filesv1_proto_rawDescGZIP(), []int{2}
}

func (m *FileRequest) GetRequest() isFileRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *FileRequest) GetNationalId() kyc.NationalIdentifierType {
	if x, ok := x.GetRequest().(*FileRequest_NationalId); ok {
		return x.NationalId
	}
	return kyc.NationalIdentifierType(0)
}

func (x *FileRequest) GetResidencyProof() ResidencyProofType {
	if x, ok := x.GetRequest().(*FileRequest_ResidencyProof); ok {
		return x.ResidencyProof
	}
	return ResidencyProofType_UTILITY_BILL
}

func (x *FileRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type isFileRequest_Request interface {
	isFileRequest_Request()
}

type FileRequest_NationalId struct {
	NationalId kyc.NationalIdentifierType `protobuf:"varint,1,opt,name=national_id,json=nationalId,proto3,enum=kyc.NationalIdentifierType,oneof"`
}

type FileRequest_ResidencyProof struct {
	ResidencyProof ResidencyProofType `protobuf:"varint,2,opt,name=residency_proof,json=residencyProof,proto3,enum=filesv1.ResidencyProofType,oneof"`
}

func (*FileRequest_NationalId) isFileRequest_Request() {}

func (*FileRequest_ResidencyProof) isFileRequest_Request() {}

type Status_IdentificationPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info                      `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Type kyc.NationalIdentifierType `protobuf:"varint,2,opt,name=type,proto3,enum=kyc.NationalIdentifierType" json:"type,omitempty"`
}

func (x *Status_IdentificationPair) Reset() {
	*x = Status_IdentificationPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesv1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_IdentificationPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_IdentificationPair) ProtoMessage() {}

func (x *Status_IdentificationPair) ProtoReflect() protoreflect.Message {
	mi := &file_filesv1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_IdentificationPair.ProtoReflect.Descriptor instead.
func (*Status_IdentificationPair) Descriptor() ([]byte, []int) {
	return file_filesv1_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Status_IdentificationPair) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Status_IdentificationPair) GetType() kyc.NationalIdentifierType {
	if x != nil {
		return x.Type
	}
	return kyc.NationalIdentifierType(0)
}

type Status_ResidencyProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info              `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Type ResidencyProofType `protobuf:"varint,2,opt,name=type,proto3,enum=filesv1.ResidencyProofType" json:"type,omitempty"`
}

func (x *Status_ResidencyProof) Reset() {
	*x = Status_ResidencyProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesv1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status_ResidencyProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status_ResidencyProof) ProtoMessage() {}

func (x *Status_ResidencyProof) ProtoReflect() protoreflect.Message {
	mi := &file_filesv1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status_ResidencyProof.ProtoReflect.Descriptor instead.
func (*Status_ResidencyProof) Descriptor() ([]byte, []int) {
	return file_filesv1_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Status_ResidencyProof) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Status_ResidencyProof) GetType() ResidencyProofType {
	if x != nil {
		return x.Type
	}
	return ResidencyProofType_UTILITY_BILL
}

var File_filesv1_proto protoreflect.FileDescriptor

var file_filesv1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6b, 0x79, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x75, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xe2, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4a, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x1a, 0x68, 0x0a, 0x12, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x64, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52,
	0x0a, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x0f, 0x72,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2a, 0x4a, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x54, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x48,
	0x4f, 0x4e, 0x45, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x41,
	0x4e, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x21,
	0x5a, 0x1f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x61, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filesv1_proto_rawDescOnce sync.Once
	file_filesv1_proto_rawDescData = file_filesv1_proto_rawDesc
)

func file_filesv1_proto_rawDescGZIP() []byte {
	file_filesv1_proto_rawDescOnce.Do(func() {
		file_filesv1_proto_rawDescData = protoimpl.X.CompressGZIP(file_filesv1_proto_rawDescData)
	})
	return file_filesv1_proto_rawDescData
}

var file_filesv1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_filesv1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_filesv1_proto_goTypes = []interface{}{
	(ResidencyProofType)(0),           // 0: filesv1.ResidencyProofType
	(*Info)(nil),                      // 1: filesv1.Info
	(*Status)(nil),                    // 2: filesv1.Status
	(*FileRequest)(nil),               // 3: filesv1.FileRequest
	(*Status_IdentificationPair)(nil), // 4: filesv1.Status.IdentificationPair
	(*Status_ResidencyProof)(nil),     // 5: filesv1.Status.ResidencyProof
	(*base.Commit)(nil),               // 6: base.Commit
	(kyc.NationalIdentifierType)(0),   // 7: kyc.NationalIdentifierType
}
var file_filesv1_proto_depIdxs = []int32{
	6, // 0: filesv1.Info.commit:type_name -> base.Commit
	4, // 1: filesv1.Status.identification:type_name -> filesv1.Status.IdentificationPair
	5, // 2: filesv1.Status.residency:type_name -> filesv1.Status.ResidencyProof
	7, // 3: filesv1.FileRequest.national_id:type_name -> kyc.NationalIdentifierType
	0, // 4: filesv1.FileRequest.residency_proof:type_name -> filesv1.ResidencyProofType
	1, // 5: filesv1.Status.IdentificationPair.info:type_name -> filesv1.Info
	7, // 6: filesv1.Status.IdentificationPair.type:type_name -> kyc.NationalIdentifierType
	1, // 7: filesv1.Status.ResidencyProof.info:type_name -> filesv1.Info
	0, // 8: filesv1.Status.ResidencyProof.type:type_name -> filesv1.ResidencyProofType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_filesv1_proto_init() }
func file_filesv1_proto_init() {
	if File_filesv1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filesv1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_filesv1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_filesv1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_filesv1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_IdentificationPair); i {
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
		file_filesv1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status_ResidencyProof); i {
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
	file_filesv1_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FileRequest_NationalId)(nil),
		(*FileRequest_ResidencyProof)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filesv1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filesv1_proto_goTypes,
		DependencyIndexes: file_filesv1_proto_depIdxs,
		EnumInfos:         file_filesv1_proto_enumTypes,
		MessageInfos:      file_filesv1_proto_msgTypes,
	}.Build()
	File_filesv1_proto = out.File
	file_filesv1_proto_rawDesc = nil
	file_filesv1_proto_goTypes = nil
	file_filesv1_proto_depIdxs = nil
}
