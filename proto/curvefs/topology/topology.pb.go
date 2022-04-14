// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: topology.proto

package topology

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

type StatMetadataUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatMetadataUsageRequest) Reset() {
	*x = StatMetadataUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topology_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatMetadataUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatMetadataUsageRequest) ProtoMessage() {}

func (x *StatMetadataUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_topology_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatMetadataUsageRequest.ProtoReflect.Descriptor instead.
func (*StatMetadataUsageRequest) Descriptor() ([]byte, []int) {
	return file_topology_proto_rawDescGZIP(), []int{0}
}

type MetadataUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KB
	MetaserverAddr *string `protobuf:"bytes,1,req,name=metaserverAddr" json:"metaserverAddr,omitempty"`
	Total          *uint64 `protobuf:"varint,2,req,name=total" json:"total,omitempty"`
	Used           *uint64 `protobuf:"varint,3,req,name=used" json:"used,omitempty"`
}

func (x *MetadataUsage) Reset() {
	*x = MetadataUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topology_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataUsage) ProtoMessage() {}

func (x *MetadataUsage) ProtoReflect() protoreflect.Message {
	mi := &file_topology_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataUsage.ProtoReflect.Descriptor instead.
func (*MetadataUsage) Descriptor() ([]byte, []int) {
	return file_topology_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataUsage) GetMetaserverAddr() string {
	if x != nil && x.MetaserverAddr != nil {
		return *x.MetaserverAddr
	}
	return ""
}

func (x *MetadataUsage) GetTotal() uint64 {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return 0
}

func (x *MetadataUsage) GetUsed() uint64 {
	if x != nil && x.Used != nil {
		return *x.Used
	}
	return 0
}

type StatMetadataUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataUsages []*MetadataUsage `protobuf:"bytes,1,rep,name=metadataUsages" json:"metadataUsages,omitempty"`
}

func (x *StatMetadataUsageResponse) Reset() {
	*x = StatMetadataUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topology_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatMetadataUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatMetadataUsageResponse) ProtoMessage() {}

func (x *StatMetadataUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_topology_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatMetadataUsageResponse.ProtoReflect.Descriptor instead.
func (*StatMetadataUsageResponse) Descriptor() ([]byte, []int) {
	return file_topology_proto_rawDescGZIP(), []int{2}
}

func (x *StatMetadataUsageResponse) GetMetadataUsages() []*MetadataUsage {
	if x != nil {
		return x.MetadataUsages
	}
	return nil
}

var File_topology_proto protoreflect.FileDescriptor

var file_topology_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x63, 0x75, 0x72, 0x76, 0x65, 0x66, 0x73, 0x2e, 0x6d, 0x64, 0x73, 0x2e, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x61, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x64, 0x22, 0x68, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x75, 0x72,
	0x76, 0x65, 0x66, 0x73, 0x2e, 0x6d, 0x64, 0x73, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32,
	0x87, 0x01, 0x0a, 0x0f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x2e, 0x63, 0x75, 0x72, 0x76, 0x65,
	0x66, 0x73, 0x2e, 0x6d, 0x64, 0x73, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x75, 0x72, 0x76, 0x65,
	0x66, 0x73, 0x2e, 0x6d, 0x64, 0x73, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x10, 0x63, 0x75, 0x72,
	0x76, 0x65, 0x66, 0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x80, 0x01, 0x01,
}

var (
	file_topology_proto_rawDescOnce sync.Once
	file_topology_proto_rawDescData = file_topology_proto_rawDesc
)

func file_topology_proto_rawDescGZIP() []byte {
	file_topology_proto_rawDescOnce.Do(func() {
		file_topology_proto_rawDescData = protoimpl.X.CompressGZIP(file_topology_proto_rawDescData)
	})
	return file_topology_proto_rawDescData
}

var file_topology_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_topology_proto_goTypes = []interface{}{
	(*StatMetadataUsageRequest)(nil),  // 0: curvefs.mds.topology.StatMetadataUsageRequest
	(*MetadataUsage)(nil),             // 1: curvefs.mds.topology.MetadataUsage
	(*StatMetadataUsageResponse)(nil), // 2: curvefs.mds.topology.StatMetadataUsageResponse
}
var file_topology_proto_depIdxs = []int32{
	1, // 0: curvefs.mds.topology.StatMetadataUsageResponse.metadataUsages:type_name -> curvefs.mds.topology.MetadataUsage
	0, // 1: curvefs.mds.topology.TopologyService.StatMetadataUsage:input_type -> curvefs.mds.topology.StatMetadataUsageRequest
	2, // 2: curvefs.mds.topology.TopologyService.StatMetadataUsage:output_type -> curvefs.mds.topology.StatMetadataUsageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_topology_proto_init() }
func file_topology_proto_init() {
	if File_topology_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_topology_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatMetadataUsageRequest); i {
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
		file_topology_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataUsage); i {
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
		file_topology_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatMetadataUsageResponse); i {
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
			RawDescriptor: file_topology_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_topology_proto_goTypes,
		DependencyIndexes: file_topology_proto_depIdxs,
		MessageInfos:      file_topology_proto_msgTypes,
	}.Build()
	File_topology_proto = out.File
	file_topology_proto_rawDesc = nil
	file_topology_proto_goTypes = nil
	file_topology_proto_depIdxs = nil
}