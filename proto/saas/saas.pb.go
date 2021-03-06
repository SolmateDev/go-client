// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: saas.proto

package saas

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

type Service int32

const (
	Service_VALIDATOR_RPC      Service = 0
	Service_VALIDATOR_INSTANCE Service = 1
	Service_TX_PROCESSING      Service = 2
	Service_SOLANA_SANDBOX     Service = 3
)

// Enum value maps for Service.
var (
	Service_name = map[int32]string{
		0: "VALIDATOR_RPC",
		1: "VALIDATOR_INSTANCE",
		2: "TX_PROCESSING",
		3: "SOLANA_SANDBOX",
	}
	Service_value = map[string]int32{
		"VALIDATOR_RPC":      0,
		"VALIDATOR_INSTANCE": 1,
		"TX_PROCESSING":      2,
		"SOLANA_SANDBOX":     3,
	}
)

func (x Service) Enum() *Service {
	p := new(Service)
	*p = x
	return p
}

func (x Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service) Descriptor() protoreflect.EnumDescriptor {
	return file_saas_proto_enumTypes[0].Descriptor()
}

func (Service) Type() protoreflect.EnumType {
	return &file_saas_proto_enumTypes[0]
}

func (x Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service.Descriptor instead.
func (Service) EnumDescriptor() ([]byte, []int) {
	return file_saas_proto_rawDescGZIP(), []int{0}
}

type Unit int32

const (
	Unit_TIME_HOUR  Unit = 0
	Unit_TIME_MONTH Unit = 1
	Unit_DATA_KB    Unit = 2
	Unit_DATA_MB    Unit = 3
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "TIME_HOUR",
		1: "TIME_MONTH",
		2: "DATA_KB",
		3: "DATA_MB",
	}
	Unit_value = map[string]int32{
		"TIME_HOUR":  0,
		"TIME_MONTH": 1,
		"DATA_KB":    2,
		"DATA_MB":    3,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_saas_proto_enumTypes[1].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_saas_proto_enumTypes[1]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_saas_proto_rawDescGZIP(), []int{1}
}

type Tier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpTo  int64      `protobuf:"varint,1,opt,name=up_to,json=upTo,proto3" json:"up_to,omitempty"`
	Fixed *base.Rate `protobuf:"bytes,2,opt,name=fixed,proto3" json:"fixed,omitempty"`
	Float *base.Rate `protobuf:"bytes,3,opt,name=float,proto3" json:"float,omitempty"`
}

func (x *Tier) Reset() {
	*x = Tier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tier) ProtoMessage() {}

func (x *Tier) ProtoReflect() protoreflect.Message {
	mi := &file_saas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tier.ProtoReflect.Descriptor instead.
func (*Tier) Descriptor() ([]byte, []int) {
	return file_saas_proto_rawDescGZIP(), []int{0}
}

func (x *Tier) GetUpTo() int64 {
	if x != nil {
		return x.UpTo
	}
	return 0
}

func (x *Tier) GetFixed() *base.Rate {
	if x != nil {
		return x.Fixed
	}
	return nil
}

func (x *Tier) GetFloat() *base.Rate {
	if x != nil {
		return x.Float
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service     Service                   `protobuf:"varint,1,opt,name=service,proto3,enum=saas.Service" json:"service,omitempty"`
	ProductId   string                    `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PriceId     string                    `protobuf:"bytes,3,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
	Name        *base.LanguageTranslation `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description *base.LanguageTranslation `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Tier        []*Tier                   `protobuf:"bytes,6,rep,name=tier,proto3" json:"tier,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_saas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_saas_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetService() Service {
	if x != nil {
		return x.Service
	}
	return Service_VALIDATOR_RPC
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetPriceId() string {
	if x != nil {
		return x.PriceId
	}
	return ""
}

func (x *Product) GetName() *base.LanguageTranslation {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Product) GetDescription() *base.LanguageTranslation {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Product) GetTier() []*Tier {
	if x != nil {
		return x.Tier
	}
	return nil
}

type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId     string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	SubscriptionId string `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	Start          int64  `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	Finish         int64  `protobuf:"varint,4,opt,name=finish,proto3" json:"finish,omitempty"`
	DataA          int64  `protobuf:"varint,5,opt,name=data_a,json=dataA,proto3" json:"data_a,omitempty"`
	DataB          int64  `protobuf:"varint,6,opt,name=data_b,json=dataB,proto3" json:"data_b,omitempty"`
	Count          int64  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_saas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_saas_proto_rawDescGZIP(), []int{2}
}

func (x *Usage) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Usage) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *Usage) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Usage) GetFinish() int64 {
	if x != nil {
		return x.Finish
	}
	return 0
}

func (x *Usage) GetDataA() int64 {
	if x != nil {
		return x.DataA
	}
	return 0
}

func (x *Usage) GetDataB() int64 {
	if x != nil {
		return x.DataB
	}
	return 0
}

func (x *Usage) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UsageHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage []*Usage `protobuf:"bytes,1,rep,name=usage,proto3" json:"usage,omitempty"`
}

func (x *UsageHistory) Reset() {
	*x = UsageHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageHistory) ProtoMessage() {}

func (x *UsageHistory) ProtoReflect() protoreflect.Message {
	mi := &file_saas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageHistory.ProtoReflect.Descriptor instead.
func (*UsageHistory) Descriptor() ([]byte, []int) {
	return file_saas_proto_rawDescGZIP(), []int{3}
}

func (x *UsageHistory) GetUsage() []*Usage {
	if x != nil {
		return x.Usage
	}
	return nil
}

type UsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      *base.TimeWindow `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	ProductId string           `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *UsageRequest) Reset() {
	*x = UsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageRequest) ProtoMessage() {}

func (x *UsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saas_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageRequest.ProtoReflect.Descriptor instead.
func (*UsageRequest) Descriptor() ([]byte, []int) {
	return file_saas_proto_rawDescGZIP(), []int{4}
}

func (x *UsageRequest) GetTime() *base.TimeWindow {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *UsageRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

var File_saas_proto protoreflect.FileDescriptor

var file_saas_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x61,
	0x61, 0x73, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f,
	0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x05, 0x75, 0x70, 0x5f, 0x74, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x70, 0x54, 0x6f, 0x12, 0x20, 0x0a, 0x05, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x12, 0x20, 0x0a,
	0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x22,
	0xf8, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73,
	0x61, 0x61, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2d,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x69,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e,
	0x54, 0x69, 0x65, 0x72, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x22, 0xc3, 0x01, 0x0a, 0x05, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x15, 0x0a, 0x06,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x61,
	0x74, 0x61, 0x41, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x42, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x31, 0x0a, 0x0c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x53, 0x0a, 0x0c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x2a, 0x5b, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x4f, 0x52,
	0x5f, 0x52, 0x50, 0x43, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41,
	0x54, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x54, 0x58, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x4f, 0x4c, 0x41, 0x4e, 0x41, 0x5f, 0x53, 0x41, 0x4e, 0x44,
	0x42, 0x4f, 0x58, 0x10, 0x03, 0x2a, 0x3f, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x4b, 0x42, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x4d, 0x42, 0x10, 0x03, 0x42, 0x1d, 0x5a, 0x1b, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x61, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x61, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saas_proto_rawDescOnce sync.Once
	file_saas_proto_rawDescData = file_saas_proto_rawDesc
)

func file_saas_proto_rawDescGZIP() []byte {
	file_saas_proto_rawDescOnce.Do(func() {
		file_saas_proto_rawDescData = protoimpl.X.CompressGZIP(file_saas_proto_rawDescData)
	})
	return file_saas_proto_rawDescData
}

var file_saas_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_saas_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_saas_proto_goTypes = []interface{}{
	(Service)(0),                     // 0: saas.Service
	(Unit)(0),                        // 1: saas.Unit
	(*Tier)(nil),                     // 2: saas.Tier
	(*Product)(nil),                  // 3: saas.Product
	(*Usage)(nil),                    // 4: saas.Usage
	(*UsageHistory)(nil),             // 5: saas.UsageHistory
	(*UsageRequest)(nil),             // 6: saas.UsageRequest
	(*base.Rate)(nil),                // 7: base.Rate
	(*base.LanguageTranslation)(nil), // 8: base.LanguageTranslation
	(*base.TimeWindow)(nil),          // 9: base.TimeWindow
}
var file_saas_proto_depIdxs = []int32{
	7, // 0: saas.Tier.fixed:type_name -> base.Rate
	7, // 1: saas.Tier.float:type_name -> base.Rate
	0, // 2: saas.Product.service:type_name -> saas.Service
	8, // 3: saas.Product.name:type_name -> base.LanguageTranslation
	8, // 4: saas.Product.description:type_name -> base.LanguageTranslation
	2, // 5: saas.Product.tier:type_name -> saas.Tier
	4, // 6: saas.UsageHistory.usage:type_name -> saas.Usage
	9, // 7: saas.UsageRequest.time:type_name -> base.TimeWindow
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_saas_proto_init() }
func file_saas_proto_init() {
	if File_saas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_saas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tier); i {
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
		file_saas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_saas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
		file_saas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageHistory); i {
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
		file_saas_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageRequest); i {
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
			RawDescriptor: file_saas_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saas_proto_goTypes,
		DependencyIndexes: file_saas_proto_depIdxs,
		EnumInfos:         file_saas_proto_enumTypes,
		MessageInfos:      file_saas_proto_msgTypes,
	}.Build()
	File_saas_proto = out.File
	file_saas_proto_rawDesc = nil
	file_saas_proto_goTypes = nil
	file_saas_proto_depIdxs = nil
}
