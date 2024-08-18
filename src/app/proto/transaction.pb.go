// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.5
// source: src/app/proto/transaction.proto

package transaction

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number           string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type             string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Stock            string `protobuf:"bytes,3,opt,name=stock,proto3" json:"stock,omitempty"`
	OrderBook        string `protobuf:"bytes,4,opt,name=order_book,json=orderBook,proto3" json:"order_book,omitempty"`
	Quantity         string `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price            string `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	ExecutedQuantity string `protobuf:"bytes,7,opt,name=executed_quantity,json=executedQuantity,proto3" json:"executed_quantity,omitempty"`
	ExecutionPrice   string `protobuf:"bytes,8,opt,name=execution_price,json=executionPrice,proto3" json:"execution_price,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_app_proto_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_src_app_proto_transaction_proto_msgTypes[0]
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
	return file_src_app_proto_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Transaction) GetStock() string {
	if x != nil {
		return x.Stock
	}
	return ""
}

func (x *Transaction) GetOrderBook() string {
	if x != nil {
		return x.OrderBook
	}
	return ""
}

func (x *Transaction) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *Transaction) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Transaction) GetExecutedQuantity() string {
	if x != nil {
		return x.ExecutedQuantity
	}
	return ""
}

func (x *Transaction) GetExecutionPrice() string {
	if x != nil {
		return x.ExecutionPrice
	}
	return ""
}

type TransactionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction []*Transaction `protobuf:"bytes,1,rep,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *TransactionReq) Reset() {
	*x = TransactionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_app_proto_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionReq) ProtoMessage() {}

func (x *TransactionReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_app_proto_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionReq.ProtoReflect.Descriptor instead.
func (*TransactionReq) Descriptor() ([]byte, []int) {
	return file_src_app_proto_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionReq) GetTransaction() []*Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type TransactionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousPrice int64 `protobuf:"varint,1,opt,name=previous_price,json=previousPrice,proto3" json:"previous_price,omitempty"`
	OpenPrice     int64 `protobuf:"varint,2,opt,name=open_price,json=openPrice,proto3" json:"open_price,omitempty"`
	HighestPrice  int64 `protobuf:"varint,3,opt,name=highest_price,json=highestPrice,proto3" json:"highest_price,omitempty"`
	LowestPrice   int64 `protobuf:"varint,4,opt,name=lowest_price,json=lowestPrice,proto3" json:"lowest_price,omitempty"`
	ClosePrice    int64 `protobuf:"varint,5,opt,name=close_price,json=closePrice,proto3" json:"close_price,omitempty"`
	Volume        int64 `protobuf:"varint,6,opt,name=volume,proto3" json:"volume,omitempty"`
	Value         int64 `protobuf:"varint,7,opt,name=value,proto3" json:"value,omitempty"`
	AveragePrice  int64 `protobuf:"varint,8,opt,name=average_price,json=averagePrice,proto3" json:"average_price,omitempty"`
}

func (x *TransactionResp) Reset() {
	*x = TransactionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_app_proto_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResp) ProtoMessage() {}

func (x *TransactionResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_app_proto_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResp.ProtoReflect.Descriptor instead.
func (*TransactionResp) Descriptor() ([]byte, []int) {
	return file_src_app_proto_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionResp) GetPreviousPrice() int64 {
	if x != nil {
		return x.PreviousPrice
	}
	return 0
}

func (x *TransactionResp) GetOpenPrice() int64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *TransactionResp) GetHighestPrice() int64 {
	if x != nil {
		return x.HighestPrice
	}
	return 0
}

func (x *TransactionResp) GetLowestPrice() int64 {
	if x != nil {
		return x.LowestPrice
	}
	return 0
}

func (x *TransactionResp) GetClosePrice() int64 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *TransactionResp) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *TransactionResp) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TransactionResp) GetAveragePrice() int64 {
	if x != nil {
		return x.AveragePrice
	}
	return 0
}

var File_src_app_proto_transaction_proto protoreflect.FileDescriptor

var file_src_app_proto_transaction_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf6,
	0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x02, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x6f, 0x77, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0x57, 0x0a, 0x0b, 0x4f,
	0x48, 0x4c, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_src_app_proto_transaction_proto_rawDescOnce sync.Once
	file_src_app_proto_transaction_proto_rawDescData = file_src_app_proto_transaction_proto_rawDesc
)

func file_src_app_proto_transaction_proto_rawDescGZIP() []byte {
	file_src_app_proto_transaction_proto_rawDescOnce.Do(func() {
		file_src_app_proto_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_app_proto_transaction_proto_rawDescData)
	})
	return file_src_app_proto_transaction_proto_rawDescData
}

var file_src_app_proto_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_app_proto_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),     // 0: transaction.Transaction
	(*TransactionReq)(nil),  // 1: transaction.TransactionReq
	(*TransactionResp)(nil), // 2: transaction.TransactionResp
}
var file_src_app_proto_transaction_proto_depIdxs = []int32{
	0, // 0: transaction.TransactionReq.transaction:type_name -> transaction.Transaction
	1, // 1: transaction.OHLCService.Calculate:input_type -> transaction.TransactionReq
	2, // 2: transaction.OHLCService.Calculate:output_type -> transaction.TransactionResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_app_proto_transaction_proto_init() }
func file_src_app_proto_transaction_proto_init() {
	if File_src_app_proto_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_app_proto_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_src_app_proto_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionReq); i {
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
		file_src_app_proto_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResp); i {
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
			RawDescriptor: file_src_app_proto_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_app_proto_transaction_proto_goTypes,
		DependencyIndexes: file_src_app_proto_transaction_proto_depIdxs,
		MessageInfos:      file_src_app_proto_transaction_proto_msgTypes,
	}.Build()
	File_src_app_proto_transaction_proto = out.File
	file_src_app_proto_transaction_proto_rawDesc = nil
	file_src_app_proto_transaction_proto_goTypes = nil
	file_src_app_proto_transaction_proto_depIdxs = nil
}