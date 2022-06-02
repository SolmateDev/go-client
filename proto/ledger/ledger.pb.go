// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: ledger.proto

package ledger

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

type BalanceSheet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId    string             `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Asset     *GroupAccount      `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Liability *GroupAccount      `protobuf:"bytes,3,opt,name=liability,proto3" json:"liability,omitempty"`
	Equity    *Equity            `protobuf:"bytes,4,opt,name=equity,proto3" json:"equity,omitempty"`
	Master    []*CommodityMaster `protobuf:"bytes,5,rep,name=master,proto3" json:"master,omitempty"`
}

func (x *BalanceSheet) Reset() {
	*x = BalanceSheet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceSheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceSheet) ProtoMessage() {}

func (x *BalanceSheet) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceSheet.ProtoReflect.Descriptor instead.
func (*BalanceSheet) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceSheet) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BalanceSheet) GetAsset() *GroupAccount {
	if x != nil {
		return x.Asset
	}
	return nil
}

func (x *BalanceSheet) GetLiability() *GroupAccount {
	if x != nil {
		return x.Liability
	}
	return nil
}

func (x *BalanceSheet) GetEquity() *Equity {
	if x != nil {
		return x.Equity
	}
	return nil
}

func (x *BalanceSheet) GetMaster() []*CommodityMaster {
	if x != nil {
		return x.Master
	}
	return nil
}

type CommodityMaster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commodity *base.ShortCommodity `protobuf:"bytes,1,opt,name=commodity,proto3" json:"commodity,omitempty"`
	// Types that are assignable to Master:
	//	*CommodityMaster_Currency
	//	*CommodityMaster_Crypto
	//	*CommodityMaster_Product
	Master isCommodityMaster_Master `protobuf_oneof:"master"`
}

func (x *CommodityMaster) Reset() {
	*x = CommodityMaster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommodityMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommodityMaster) ProtoMessage() {}

func (x *CommodityMaster) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommodityMaster.ProtoReflect.Descriptor instead.
func (*CommodityMaster) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{1}
}

func (x *CommodityMaster) GetCommodity() *base.ShortCommodity {
	if x != nil {
		return x.Commodity
	}
	return nil
}

func (m *CommodityMaster) GetMaster() isCommodityMaster_Master {
	if m != nil {
		return m.Master
	}
	return nil
}

func (x *CommodityMaster) GetCurrency() *CurrencyMaster {
	if x, ok := x.GetMaster().(*CommodityMaster_Currency); ok {
		return x.Currency
	}
	return nil
}

func (x *CommodityMaster) GetCrypto() *CryptoMaster {
	if x, ok := x.GetMaster().(*CommodityMaster_Crypto); ok {
		return x.Crypto
	}
	return nil
}

func (x *CommodityMaster) GetProduct() *ProductMaster {
	if x, ok := x.GetMaster().(*CommodityMaster_Product); ok {
		return x.Product
	}
	return nil
}

type isCommodityMaster_Master interface {
	isCommodityMaster_Master()
}

type CommodityMaster_Currency struct {
	Currency *CurrencyMaster `protobuf:"bytes,2,opt,name=currency,proto3,oneof"`
}

type CommodityMaster_Crypto struct {
	Crypto *CryptoMaster `protobuf:"bytes,3,opt,name=crypto,proto3,oneof"`
}

type CommodityMaster_Product struct {
	Product *ProductMaster `protobuf:"bytes,4,opt,name=product,proto3,oneof"`
}

func (*CommodityMaster_Currency) isCommodityMaster_Master() {}

func (*CommodityMaster_Crypto) isCommodityMaster_Master() {}

func (*CommodityMaster_Product) isCommodityMaster_Master() {}

type CurrencyMaster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency base.Currency `protobuf:"varint,1,opt,name=currency,proto3,enum=base.Currency" json:"currency,omitempty"`
}

func (x *CurrencyMaster) Reset() {
	*x = CurrencyMaster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyMaster) ProtoMessage() {}

func (x *CurrencyMaster) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyMaster.ProtoReflect.Descriptor instead.
func (*CurrencyMaster) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyMaster) GetCurrency() base.Currency {
	if x != nil {
		return x.Currency
	}
	return base.Currency(0)
}

type CryptoMaster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto base.Crypto `protobuf:"varint,2,opt,name=crypto,proto3,enum=base.Crypto" json:"crypto,omitempty"`
}

func (x *CryptoMaster) Reset() {
	*x = CryptoMaster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoMaster) ProtoMessage() {}

func (x *CryptoMaster) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoMaster.ProtoReflect.Descriptor instead.
func (*CryptoMaster) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{3}
}

func (x *CryptoMaster) GetCrypto() base.Crypto {
	if x != nil {
		return x.Crypto
	}
	return base.Crypto(0)
}

type ProductMaster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityId *base.SecurityID `protobuf:"bytes,1,opt,name=security_id,json=securityId,proto3" json:"security_id,omitempty"`
}

func (x *ProductMaster) Reset() {
	*x = ProductMaster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductMaster) ProtoMessage() {}

func (x *ProductMaster) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductMaster.ProtoReflect.Descriptor instead.
func (*ProductMaster) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{4}
}

func (x *ProductMaster) GetSecurityId() *base.SecurityID {
	if x != nil {
		return x.SecurityId
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent  int64       `protobuf:"varint,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Balance *base.Value `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{5}
}

func (x *Account) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetParent() int64 {
	if x != nil {
		return x.Parent
	}
	return 0
}

func (x *Account) GetBalance() *base.Value {
	if x != nil {
		return x.Balance
	}
	return nil
}

type GroupAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top    *Account            `protobuf:"bytes,1,opt,name=top,proto3" json:"top,omitempty"`
	Source map[string]*Account `protobuf:"bytes,2,rep,name=source,proto3" json:"source,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GroupAccount) Reset() {
	*x = GroupAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupAccount) ProtoMessage() {}

func (x *GroupAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupAccount.ProtoReflect.Descriptor instead.
func (*GroupAccount) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{6}
}

func (x *GroupAccount) GetTop() *Account {
	if x != nil {
		return x.Top
	}
	return nil
}

func (x *GroupAccount) GetSource() map[string]*Account {
	if x != nil {
		return x.Source
	}
	return nil
}

type Equity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top    *GroupAccount     `protobuf:"bytes,1,opt,name=top,proto3" json:"top,omitempty"`
	Income map[int64]*Income `protobuf:"bytes,2,rep,name=income,proto3" json:"income,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map year to income
}

func (x *Equity) Reset() {
	*x = Equity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equity) ProtoMessage() {}

func (x *Equity) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equity.ProtoReflect.Descriptor instead.
func (*Equity) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{7}
}

func (x *Equity) GetTop() *GroupAccount {
	if x != nil {
		return x.Top
	}
	return nil
}

func (x *Equity) GetIncome() map[int64]*Income {
	if x != nil {
		return x.Income
	}
	return nil
}

type Income struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revenue *GroupAccount `protobuf:"bytes,1,opt,name=revenue,proto3" json:"revenue,omitempty"`
	Expense *GroupAccount `protobuf:"bytes,2,opt,name=expense,proto3" json:"expense,omitempty"`
}

func (x *Income) Reset() {
	*x = Income{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Income) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Income) ProtoMessage() {}

func (x *Income) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Income.ProtoReflect.Descriptor instead.
func (*Income) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{8}
}

func (x *Income) GetRevenue() *GroupAccount {
	if x != nil {
		return x.Revenue
	}
	return nil
}

func (x *Income) GetExpense() *GroupAccount {
	if x != nil {
		return x.Expense
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commit        *base.Commit `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	TransactionId int64        `protobuf:"varint,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{9}
}

func (x *Transaction) GetCommit() *base.Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *Transaction) GetTransactionId() int64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

type Split struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int64 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Amount    int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Split) Reset() {
	*x = Split{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Split) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Split) ProtoMessage() {}

func (x *Split) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Split.ProtoReflect.Descriptor instead.
func (*Split) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{10}
}

func (x *Split) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Split) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ListMasterSingle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Single:
	//	*ListMasterSingle_Crypto
	//	*ListMasterSingle_Currency
	Single isListMasterSingle_Single `protobuf_oneof:"single"`
}

func (x *ListMasterSingle) Reset() {
	*x = ListMasterSingle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMasterSingle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMasterSingle) ProtoMessage() {}

func (x *ListMasterSingle) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMasterSingle.ProtoReflect.Descriptor instead.
func (*ListMasterSingle) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{11}
}

func (m *ListMasterSingle) GetSingle() isListMasterSingle_Single {
	if m != nil {
		return m.Single
	}
	return nil
}

func (x *ListMasterSingle) GetCrypto() *CryptoMaster {
	if x, ok := x.GetSingle().(*ListMasterSingle_Crypto); ok {
		return x.Crypto
	}
	return nil
}

func (x *ListMasterSingle) GetCurrency() *CurrencyMaster {
	if x, ok := x.GetSingle().(*ListMasterSingle_Currency); ok {
		return x.Currency
	}
	return nil
}

type isListMasterSingle_Single interface {
	isListMasterSingle_Single()
}

type ListMasterSingle_Crypto struct {
	Crypto *CryptoMaster `protobuf:"bytes,1,opt,name=crypto,proto3,oneof"`
}

type ListMasterSingle_Currency struct {
	Currency *CurrencyMaster `protobuf:"bytes,2,opt,name=currency,proto3,oneof"`
}

func (*ListMasterSingle_Crypto) isListMasterSingle_Single() {}

func (*ListMasterSingle_Currency) isListMasterSingle_Single() {}

type ListMasterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Single []*ListMasterSingle `protobuf:"bytes,1,rep,name=single,proto3" json:"single,omitempty"`
}

func (x *ListMasterResponse) Reset() {
	*x = ListMasterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMasterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMasterResponse) ProtoMessage() {}

func (x *ListMasterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMasterResponse.ProtoReflect.Descriptor instead.
func (*ListMasterResponse) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{12}
}

func (x *ListMasterResponse) GetSingle() []*ListMasterSingle {
	if x != nil {
		return x.Single
	}
	return nil
}

type ListMasterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMasterRequest) Reset() {
	*x = ListMasterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMasterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMasterRequest) ProtoMessage() {}

func (x *ListMasterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMasterRequest.ProtoReflect.Descriptor instead.
func (*ListMasterRequest) Descriptor() ([]byte, []int) {
	return file_ledger_proto_rawDescGZIP(), []int{13}
}

var File_ledger_proto protoreflect.FileDescriptor

var file_ledger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x68,
	0x65, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x05,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x6c, 0x69, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x09, 0x6c, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x06,
	0x65, 0x71, 0x75, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x71,
	0x75, 0x69, 0x74, 0x79, 0x12, 0x2f, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x06, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x22, 0xe8, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69,
	0x74, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x34,
	0x0a, 0x0c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x0a, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x1a, 0x4a, 0x0a, 0x0b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xaf, 0x01, 0x0a,
	0x06, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12,
	0x32, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x74, 0x79, 0x2e,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x63,
	0x6f, 0x6d, 0x65, 0x1a, 0x49, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x68,
	0x0a, 0x06, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65,
	0x6e, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x07, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x05, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x61, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ledger_proto_rawDescOnce sync.Once
	file_ledger_proto_rawDescData = file_ledger_proto_rawDesc
)

func file_ledger_proto_rawDescGZIP() []byte {
	file_ledger_proto_rawDescOnce.Do(func() {
		file_ledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledger_proto_rawDescData)
	})
	return file_ledger_proto_rawDescData
}

var file_ledger_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_ledger_proto_goTypes = []interface{}{
	(*BalanceSheet)(nil),        // 0: ledger.BalanceSheet
	(*CommodityMaster)(nil),     // 1: ledger.CommodityMaster
	(*CurrencyMaster)(nil),      // 2: ledger.CurrencyMaster
	(*CryptoMaster)(nil),        // 3: ledger.CryptoMaster
	(*ProductMaster)(nil),       // 4: ledger.ProductMaster
	(*Account)(nil),             // 5: ledger.Account
	(*GroupAccount)(nil),        // 6: ledger.GroupAccount
	(*Equity)(nil),              // 7: ledger.Equity
	(*Income)(nil),              // 8: ledger.Income
	(*Transaction)(nil),         // 9: ledger.Transaction
	(*Split)(nil),               // 10: ledger.Split
	(*ListMasterSingle)(nil),    // 11: ledger.ListMasterSingle
	(*ListMasterResponse)(nil),  // 12: ledger.ListMasterResponse
	(*ListMasterRequest)(nil),   // 13: ledger.ListMasterRequest
	nil,                         // 14: ledger.GroupAccount.SourceEntry
	nil,                         // 15: ledger.Equity.IncomeEntry
	(*base.ShortCommodity)(nil), // 16: base.ShortCommodity
	(base.Currency)(0),          // 17: base.Currency
	(base.Crypto)(0),            // 18: base.Crypto
	(*base.SecurityID)(nil),     // 19: base.SecurityID
	(*base.Value)(nil),          // 20: base.Value
	(*base.Commit)(nil),         // 21: base.Commit
}
var file_ledger_proto_depIdxs = []int32{
	6,  // 0: ledger.BalanceSheet.asset:type_name -> ledger.GroupAccount
	6,  // 1: ledger.BalanceSheet.liability:type_name -> ledger.GroupAccount
	7,  // 2: ledger.BalanceSheet.equity:type_name -> ledger.Equity
	1,  // 3: ledger.BalanceSheet.master:type_name -> ledger.CommodityMaster
	16, // 4: ledger.CommodityMaster.commodity:type_name -> base.ShortCommodity
	2,  // 5: ledger.CommodityMaster.currency:type_name -> ledger.CurrencyMaster
	3,  // 6: ledger.CommodityMaster.crypto:type_name -> ledger.CryptoMaster
	4,  // 7: ledger.CommodityMaster.product:type_name -> ledger.ProductMaster
	17, // 8: ledger.CurrencyMaster.currency:type_name -> base.Currency
	18, // 9: ledger.CryptoMaster.crypto:type_name -> base.Crypto
	19, // 10: ledger.ProductMaster.security_id:type_name -> base.SecurityID
	20, // 11: ledger.Account.balance:type_name -> base.Value
	5,  // 12: ledger.GroupAccount.top:type_name -> ledger.Account
	14, // 13: ledger.GroupAccount.source:type_name -> ledger.GroupAccount.SourceEntry
	6,  // 14: ledger.Equity.top:type_name -> ledger.GroupAccount
	15, // 15: ledger.Equity.income:type_name -> ledger.Equity.IncomeEntry
	6,  // 16: ledger.Income.revenue:type_name -> ledger.GroupAccount
	6,  // 17: ledger.Income.expense:type_name -> ledger.GroupAccount
	21, // 18: ledger.Transaction.commit:type_name -> base.Commit
	3,  // 19: ledger.ListMasterSingle.crypto:type_name -> ledger.CryptoMaster
	2,  // 20: ledger.ListMasterSingle.currency:type_name -> ledger.CurrencyMaster
	11, // 21: ledger.ListMasterResponse.single:type_name -> ledger.ListMasterSingle
	5,  // 22: ledger.GroupAccount.SourceEntry.value:type_name -> ledger.Account
	8,  // 23: ledger.Equity.IncomeEntry.value:type_name -> ledger.Income
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_ledger_proto_init() }
func file_ledger_proto_init() {
	if File_ledger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ledger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceSheet); i {
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
		file_ledger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommodityMaster); i {
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
		file_ledger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyMaster); i {
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
		file_ledger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoMaster); i {
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
		file_ledger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductMaster); i {
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
		file_ledger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_ledger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupAccount); i {
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
		file_ledger_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equity); i {
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
		file_ledger_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Income); i {
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
		file_ledger_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_ledger_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Split); i {
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
		file_ledger_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMasterSingle); i {
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
		file_ledger_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMasterResponse); i {
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
		file_ledger_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMasterRequest); i {
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
	file_ledger_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CommodityMaster_Currency)(nil),
		(*CommodityMaster_Crypto)(nil),
		(*CommodityMaster_Product)(nil),
	}
	file_ledger_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*ListMasterSingle_Crypto)(nil),
		(*ListMasterSingle_Currency)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ledger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ledger_proto_goTypes,
		DependencyIndexes: file_ledger_proto_depIdxs,
		MessageInfos:      file_ledger_proto_msgTypes,
	}.Build()
	File_ledger_proto = out.File
	file_ledger_proto_rawDesc = nil
	file_ledger_proto_goTypes = nil
	file_ledger_proto_depIdxs = nil
}
