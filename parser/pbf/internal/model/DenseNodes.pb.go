// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: DenseNodes.proto

package model

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

type DenseNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []int64    `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"` // DELTA coded
	Info      []*Info    `protobuf:"bytes,4,rep,name=info" json:"info,omitempty"`
	Denseinfo *DenseInfo `protobuf:"bytes,5,opt,name=denseinfo" json:"denseinfo,omitempty"`
	Lat       []int64    `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"` // DELTA coded
	Lon       []int64    `protobuf:"zigzag64,9,rep,packed,name=lon" json:"lon,omitempty"` // DELTA coded
	// Special packing of keys and vals into one array. May be empty if all nodes in this block are tagless.
	KeysVals []int32 `protobuf:"varint,10,rep,packed,name=keys_vals,json=keysVals" json:"keys_vals,omitempty"`
}

func (x *DenseNodes) Reset() {
	*x = DenseNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DenseNodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenseNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenseNodes) ProtoMessage() {}

func (x *DenseNodes) ProtoReflect() protoreflect.Message {
	mi := &file_DenseNodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenseNodes.ProtoReflect.Descriptor instead.
func (*DenseNodes) Descriptor() ([]byte, []int) {
	return file_DenseNodes_proto_rawDescGZIP(), []int{0}
}

func (x *DenseNodes) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DenseNodes) GetInfo() []*Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *DenseNodes) GetDenseinfo() *DenseInfo {
	if x != nil {
		return x.Denseinfo
	}
	return nil
}

func (x *DenseNodes) GetLat() []int64 {
	if x != nil {
		return x.Lat
	}
	return nil
}

func (x *DenseNodes) GetLon() []int64 {
	if x != nil {
		return x.Lon
	}
	return nil
}

func (x *DenseNodes) GetKeysVals() []int32 {
	if x != nil {
		return x.KeysVals
	}
	return nil
}

var File_DenseNodes_proto protoreflect.FileDescriptor

var file_DenseNodes_proto_rawDesc = []byte{
	0x0a, 0x10, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6e, 0x73, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x12, 0x42, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x09, 0x64, 0x65,
	0x6e, 0x73, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x14, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x76,
	0x61, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x73, 0x56, 0x61, 0x6c, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x6c, 0x73, 0x68, 0x69, 0x6b, 0x2f,
	0x6f, 0x73, 0x6d, 0x69, 0x6d, 0x70, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f,
}

var (
	file_DenseNodes_proto_rawDescOnce sync.Once
	file_DenseNodes_proto_rawDescData = file_DenseNodes_proto_rawDesc
)

func file_DenseNodes_proto_rawDescGZIP() []byte {
	file_DenseNodes_proto_rawDescOnce.Do(func() {
		file_DenseNodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_DenseNodes_proto_rawDescData)
	})
	return file_DenseNodes_proto_rawDescData
}

var file_DenseNodes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DenseNodes_proto_goTypes = []interface{}{
	(*DenseNodes)(nil), // 0: proto.DenseNodes
	(*Info)(nil),       // 1: proto.Info
	(*DenseInfo)(nil),  // 2: proto.DenseInfo
}
var file_DenseNodes_proto_depIdxs = []int32{
	1, // 0: proto.DenseNodes.info:type_name -> proto.Info
	2, // 1: proto.DenseNodes.denseinfo:type_name -> proto.DenseInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_DenseNodes_proto_init() }
func file_DenseNodes_proto_init() {
	if File_DenseNodes_proto != nil {
		return
	}
	file_Info_proto_init()
	file_DenseInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DenseNodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenseNodes); i {
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
			RawDescriptor: file_DenseNodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DenseNodes_proto_goTypes,
		DependencyIndexes: file_DenseNodes_proto_depIdxs,
		MessageInfos:      file_DenseNodes_proto_msgTypes,
	}.Build()
	File_DenseNodes_proto = out.File
	file_DenseNodes_proto_rawDesc = nil
	file_DenseNodes_proto_goTypes = nil
	file_DenseNodes_proto_depIdxs = nil
}