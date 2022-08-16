// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: zzzz-customer-solana.proto

package customer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "github.com/SolmateDev/go-client/proto/base"
	job "github.com/SolmateDev/go-client/proto/job"
	solana "github.com/SolmateDev/go-client/proto/solana"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_zzzz_customer_solana_proto protoreflect.FileDescriptor

var file_zzzz_customer_solana_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x7a, 0x7a, 0x7a, 0x7a, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d,
	0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdb, 0x03, 0x0a, 0x06,
	0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x12, 0x3a, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x1e, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x08, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0b, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x73, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78,
	0x12, 0x18, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x78, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x6f, 0x6c, 0x61,
	0x6e, 0x61, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2e, 0x54, 0x78, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x70, 0x61, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_zzzz_customer_solana_proto_goTypes = []interface{}{
	(*solana.LinkRequest)(nil),              // 0: solana.LinkRequest
	(*solana.AccountSubscribeRequest)(nil),  // 1: solana.AccountSubscribeRequest
	(*solana.ValidatorCreateRequest)(nil),   // 2: solana.ValidatorCreateRequest
	(*base.Empty)(nil),                      // 3: base.Empty
	(*solana.ValidatorDeleteRequest)(nil),   // 4: solana.ValidatorDeleteRequest
	(*solana.SendBatchRequest)(nil),         // 5: solana.SendBatchRequest
	(*solana.LinkResponse)(nil),             // 6: solana.LinkResponse
	(*solana.AccountSubscribeResponse)(nil), // 7: solana.AccountSubscribeResponse
	(*job.Job)(nil),                         // 8: job.Job
	(*solana.ValidatorList)(nil),            // 9: solana.ValidatorList
	(*solana.SendBatchResponse)(nil),        // 10: solana.SendBatchResponse
	(*solana.TxUpdate)(nil),                 // 11: solana.TxUpdate
}
var file_zzzz_customer_solana_proto_depIdxs = []int32{
	0,  // 0: customer.Solana.LinkAccount:input_type -> solana.LinkRequest
	1,  // 1: customer.Solana.SubscribeAccount:input_type -> solana.AccountSubscribeRequest
	2,  // 2: customer.Solana.CreateValidator:input_type -> solana.ValidatorCreateRequest
	3,  // 3: customer.Solana.ListValidator:input_type -> base.Empty
	4,  // 4: customer.Solana.DestroyValidator:input_type -> solana.ValidatorDeleteRequest
	5,  // 5: customer.Solana.SendTx:input_type -> solana.SendBatchRequest
	5,  // 6: customer.Solana.SendTxComplete:input_type -> solana.SendBatchRequest
	6,  // 7: customer.Solana.LinkAccount:output_type -> solana.LinkResponse
	7,  // 8: customer.Solana.SubscribeAccount:output_type -> solana.AccountSubscribeResponse
	8,  // 9: customer.Solana.CreateValidator:output_type -> job.Job
	9,  // 10: customer.Solana.ListValidator:output_type -> solana.ValidatorList
	3,  // 11: customer.Solana.DestroyValidator:output_type -> base.Empty
	10, // 12: customer.Solana.SendTx:output_type -> solana.SendBatchResponse
	11, // 13: customer.Solana.SendTxComplete:output_type -> solana.TxUpdate
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_zzzz_customer_solana_proto_init() }
func file_zzzz_customer_solana_proto_init() {
	if File_zzzz_customer_solana_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zzzz_customer_solana_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zzzz_customer_solana_proto_goTypes,
		DependencyIndexes: file_zzzz_customer_solana_proto_depIdxs,
	}.Build()
	File_zzzz_customer_solana_proto = out.File
	file_zzzz_customer_solana_proto_rawDesc = nil
	file_zzzz_customer_solana_proto_goTypes = nil
	file_zzzz_customer_solana_proto_depIdxs = nil
}
