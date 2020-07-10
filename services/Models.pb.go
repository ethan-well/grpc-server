// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: Models.proto

package services

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProdModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId    int32   `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdName  string  `protobuf:"bytes,2,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	ProdPrice float32 `protobuf:"fixed32,3,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
}

func (x *ProdModel) Reset() {
	*x = ProdModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdModel) ProtoMessage() {}

func (x *ProdModel) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdModel.ProtoReflect.Descriptor instead.
func (*ProdModel) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{0}
}

func (x *ProdModel) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *ProdModel) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

func (x *ProdModel) GetProdPrice() float32 {
	if x != nil {
		return x.ProdPrice
	}
	return 0
}

type OrderMain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int32                `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo    string               `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId     int32                `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderMoney float32              `protobuf:"fixed32,4,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`
	OrderTime  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=order_time,json=orderTime,proto3" json:"order_time,omitempty"`
	SubOrders  []*SubOrder          `protobuf:"bytes,6,rep,name=sub_orders,json=subOrders,proto3" json:"sub_orders,omitempty"`
}

func (x *OrderMain) Reset() {
	*x = OrderMain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderMain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderMain) ProtoMessage() {}

func (x *OrderMain) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderMain.ProtoReflect.Descriptor instead.
func (*OrderMain) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{1}
}

func (x *OrderMain) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderMain) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *OrderMain) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderMain) GetOrderMoney() float32 {
	if x != nil {
		return x.OrderMoney
	}
	return 0
}

func (x *OrderMain) GetOrderTime() *timestamp.Timestamp {
	if x != nil {
		return x.OrderTime
	}
	return nil
}

func (x *OrderMain) GetSubOrders() []*SubOrder {
	if x != nil {
		return x.SubOrders
	}
	return nil
}

// 子订单
type SubOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo string  `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId  int32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Price   float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Count   int32   `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SubOrder) Reset() {
	*x = SubOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubOrder) ProtoMessage() {}

func (x *SubOrder) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubOrder.ProtoReflect.Descriptor instead.
func (*SubOrder) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{2}
}

func (x *SubOrder) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SubOrder) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *SubOrder) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SubOrder) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SubOrder) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Score int32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_Models_proto protoreflect.FileDescriptor

var file_Models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xf5, 0x01, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x0a, 0x05, 0x25, 0x00, 0x00, 0x80, 0x3f, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x31, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x75, 0x62, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09, 0x73, 0x75, 0x62, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Models_proto_rawDescOnce sync.Once
	file_Models_proto_rawDescData = file_Models_proto_rawDesc
)

func file_Models_proto_rawDescGZIP() []byte {
	file_Models_proto_rawDescOnce.Do(func() {
		file_Models_proto_rawDescData = protoimpl.X.CompressGZIP(file_Models_proto_rawDescData)
	})
	return file_Models_proto_rawDescData
}

var file_Models_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Models_proto_goTypes = []interface{}{
	(*ProdModel)(nil),           // 0: services.ProdModel
	(*OrderMain)(nil),           // 1: services.OrderMain
	(*SubOrder)(nil),            // 2: services.SubOrder
	(*User)(nil),                // 3: services.User
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_Models_proto_depIdxs = []int32{
	4, // 0: services.OrderMain.order_time:type_name -> google.protobuf.Timestamp
	2, // 1: services.OrderMain.sub_orders:type_name -> services.SubOrder
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Models_proto_init() }
func file_Models_proto_init() {
	if File_Models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdModel); i {
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
		file_Models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderMain); i {
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
		file_Models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubOrder); i {
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
		file_Models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_Models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Models_proto_goTypes,
		DependencyIndexes: file_Models_proto_depIdxs,
		MessageInfos:      file_Models_proto_msgTypes,
	}.Build()
	File_Models_proto = out.File
	file_Models_proto_rawDesc = nil
	file_Models_proto_goTypes = nil
	file_Models_proto_depIdxs = nil
}
