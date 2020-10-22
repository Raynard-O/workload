// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: person.proto

package grpc_from0

import (
	proto "github.com/golang/protobuf/proto"
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

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CPUUtilization    int32   `protobuf:"varint,1,opt,name=CPU_utilization,json=CPUUtilization,proto3" json:"CPU_utilization,omitempty"`
	NetworkIN         int32   `protobuf:"varint,2,opt,name=NetworkIN,proto3" json:"NetworkIN,omitempty"`
	NetworkOUT        int32   `protobuf:"varint,3,opt,name=NetworkOUT,proto3" json:"NetworkOUT,omitempty"`
	MemoryUtilization float32 `protobuf:"fixed32,4,opt,name=Memory_utilization,json=MemoryUtilization,proto3" json:"Memory_utilization,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetCPUUtilization() int32 {
	if x != nil {
		return x.CPUUtilization
	}
	return 0
}

func (x *Sample) GetNetworkIN() int32 {
	if x != nil {
		return x.NetworkIN
	}
	return 0
}

func (x *Sample) GetNetworkOUT() int32 {
	if x != nil {
		return x.NetworkOUT
	}
	return 0
}

func (x *Sample) GetMemoryUtilization() float32 {
	if x != nil {
		return x.MemoryUtilization
	}
	return 0
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batch_ID int32     `protobuf:"varint,1,opt,name=Batch_ID,json=BatchID,proto3" json:"Batch_ID,omitempty"`
	Samples  []*Sample `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *Batch) GetBatch_ID() int32 {
	if x != nil {
		return x.Batch_ID
	}
	return 0
}

func (x *Batch) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type RFD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RFWID       string   `protobuf:"bytes,1,opt,name=RFWID,proto3" json:"RFWID,omitempty"`
	LastBatchId int32    `protobuf:"varint,2,opt,name=last_batch_id,json=lastBatchId,proto3" json:"last_batch_id,omitempty"`
	Batches     []*Batch `protobuf:"bytes,3,rep,name=batches,proto3" json:"batches,omitempty"`
}

func (x *RFD) Reset() {
	*x = RFD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RFD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RFD) ProtoMessage() {}

func (x *RFD) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RFD.ProtoReflect.Descriptor instead.
func (*RFD) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{2}
}

func (x *RFD) GetRFWID() string {
	if x != nil {
		return x.RFWID
	}
	return ""
}

func (x *RFD) GetLastBatchId() int32 {
	if x != nil {
		return x.LastBatchId
	}
	return 0
}

func (x *RFD) GetBatches() []*Batch {
	if x != nil {
		return x.Batches
	}
	return nil
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x43, 0x50, 0x55, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x43, 0x50, 0x55, 0x55, 0x74, 0x69,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x4e, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4f, 0x55, 0x54, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x4f, 0x55, 0x54, 0x12, 0x2d, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x11, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19,
	0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x07, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x22, 0x66, 0x0a, 0x03, 0x52, 0x46, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x46, 0x57, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x46, 0x57, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x6b, 0x72, 0x75, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x66, 0x72, 0x6f, 0x6d, 0x30, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_person_proto_goTypes = []interface{}{
	(*Sample)(nil), // 0: main.sample
	(*Batch)(nil),  // 1: main.batch
	(*RFD)(nil),    // 2: main.RFD
}
var file_person_proto_depIdxs = []int32{
	0, // 0: main.batch.samples:type_name -> main.sample
	1, // 1: main.RFD.batches:type_name -> main.batch
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
		file_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RFD); i {
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
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}
