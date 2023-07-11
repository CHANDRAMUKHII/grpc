// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: pb/sample.proto

package pb

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

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters string `protobuf:"bytes,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_pb_sample_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetParameters() string {
	if x != nil {
		return x.Parameters
	}
	return ""
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FetchedData string `protobuf:"bytes,1,opt,name=fetched_data,json=fetchedData,proto3" json:"fetched_data,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_pb_sample_proto_rawDescGZIP(), []int{1}
}

func (x *FetchResponse) GetFetchedData() string {
	if x != nil {
		return x.FetchedData
	}
	return ""
}

type BatchFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientIds []string `protobuf:"bytes,1,rep,name=patient_ids,json=patientIds,proto3" json:"patient_ids,omitempty"`
}

func (x *BatchFetchRequest) Reset() {
	*x = BatchFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchFetchRequest) ProtoMessage() {}

func (x *BatchFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchFetchRequest.ProtoReflect.Descriptor instead.
func (*BatchFetchRequest) Descriptor() ([]byte, []int) {
	return file_pb_sample_proto_rawDescGZIP(), []int{2}
}

func (x *BatchFetchRequest) GetPatientIds() []string {
	if x != nil {
		return x.PatientIds
	}
	return nil
}

type BatchFetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FetchedData []string `protobuf:"bytes,1,rep,name=fetched_data,json=fetchedData,proto3" json:"fetched_data,omitempty"`
}

func (x *BatchFetchResponse) Reset() {
	*x = BatchFetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sample_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchFetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchFetchResponse) ProtoMessage() {}

func (x *BatchFetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sample_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchFetchResponse.ProtoReflect.Descriptor instead.
func (*BatchFetchResponse) Descriptor() ([]byte, []int) {
	return file_pb_sample_proto_rawDescGZIP(), []int{3}
}

func (x *BatchFetchResponse) GetFetchedData() []string {
	if x != nil {
		return x.FetchedData
	}
	return nil
}

var File_pb_sample_proto protoreflect.FileDescriptor

var file_pb_sample_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2e, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x32, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x34, 0x0a, 0x11, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22,
	0x37, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x32, 0x99, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e,
	0x67, 0x6f, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x14, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x44, 0x42, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x74, 0x63, 0x68, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f,
	0x6e, 0x67, 0x6f, 0x44, 0x42, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_sample_proto_rawDescOnce sync.Once
	file_pb_sample_proto_rawDescData = file_pb_sample_proto_rawDesc
)

func file_pb_sample_proto_rawDescGZIP() []byte {
	file_pb_sample_proto_rawDescOnce.Do(func() {
		file_pb_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_sample_proto_rawDescData)
	})
	return file_pb_sample_proto_rawDescData
}

var file_pb_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_sample_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),       // 0: pb.FetchRequest
	(*FetchResponse)(nil),      // 1: pb.FetchResponse
	(*BatchFetchRequest)(nil),  // 2: pb.BatchFetchRequest
	(*BatchFetchResponse)(nil), // 3: pb.BatchFetchResponse
}
var file_pb_sample_proto_depIdxs = []int32{
	0, // 0: pb.MongoDBService.FetchDataFromMongoDB:input_type -> pb.FetchRequest
	2, // 1: pb.MongoDBService.FetchDataBatchFromMongoDB:input_type -> pb.BatchFetchRequest
	1, // 2: pb.MongoDBService.FetchDataFromMongoDB:output_type -> pb.FetchResponse
	3, // 3: pb.MongoDBService.FetchDataBatchFromMongoDB:output_type -> pb.BatchFetchResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_sample_proto_init() }
func file_pb_sample_proto_init() {
	if File_pb_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_pb_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_pb_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchFetchRequest); i {
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
		file_pb_sample_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchFetchResponse); i {
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
			RawDescriptor: file_pb_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_sample_proto_goTypes,
		DependencyIndexes: file_pb_sample_proto_depIdxs,
		MessageInfos:      file_pb_sample_proto_msgTypes,
	}.Build()
	File_pb_sample_proto = out.File
	file_pb_sample_proto_rawDesc = nil
	file_pb_sample_proto_goTypes = nil
	file_pb_sample_proto_depIdxs = nil
}
