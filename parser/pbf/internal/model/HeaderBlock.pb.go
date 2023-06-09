// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: HeaderBlock.proto

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

type HeaderBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bbox *HeaderBBox `protobuf:"bytes,1,opt,name=bbox" json:"bbox,omitempty"`
	// Additional tags to aid in parsing this dataset
	RequiredFeatures []string `protobuf:"bytes,4,rep,name=required_features,json=requiredFeatures" json:"required_features,omitempty"`
	OptionalFeatures []string `protobuf:"bytes,5,rep,name=optional_features,json=optionalFeatures" json:"optional_features,omitempty"`
	Writingprogram   *string  `protobuf:"bytes,16,opt,name=writingprogram" json:"writingprogram,omitempty"`
	Source           *string  `protobuf:"bytes,17,opt,name=source" json:"source,omitempty"` // From the bbox field.
}

func (x *HeaderBlock) Reset() {
	*x = HeaderBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HeaderBlock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderBlock) ProtoMessage() {}

func (x *HeaderBlock) ProtoReflect() protoreflect.Message {
	mi := &file_HeaderBlock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderBlock.ProtoReflect.Descriptor instead.
func (*HeaderBlock) Descriptor() ([]byte, []int) {
	return file_HeaderBlock_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderBlock) GetBbox() *HeaderBBox {
	if x != nil {
		return x.Bbox
	}
	return nil
}

func (x *HeaderBlock) GetRequiredFeatures() []string {
	if x != nil {
		return x.RequiredFeatures
	}
	return nil
}

func (x *HeaderBlock) GetOptionalFeatures() []string {
	if x != nil {
		return x.OptionalFeatures
	}
	return nil
}

func (x *HeaderBlock) GetWritingprogram() string {
	if x != nil && x.Writingprogram != nil {
		return *x.Writingprogram
	}
	return ""
}

func (x *HeaderBlock) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

var File_HeaderBlock_proto protoreflect.FileDescriptor

var file_HeaderBlock_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x42, 0x42, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x04,
	0x62, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x04, 0x62,
	0x62, 0x6f, 0x78, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69,
	0x6c, 0x73, 0x68, 0x69, 0x6b, 0x2f, 0x6f, 0x73, 0x6d, 0x69, 0x6d, 0x70, 0x2f, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
}

var (
	file_HeaderBlock_proto_rawDescOnce sync.Once
	file_HeaderBlock_proto_rawDescData = file_HeaderBlock_proto_rawDesc
)

func file_HeaderBlock_proto_rawDescGZIP() []byte {
	file_HeaderBlock_proto_rawDescOnce.Do(func() {
		file_HeaderBlock_proto_rawDescData = protoimpl.X.CompressGZIP(file_HeaderBlock_proto_rawDescData)
	})
	return file_HeaderBlock_proto_rawDescData
}

var file_HeaderBlock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HeaderBlock_proto_goTypes = []interface{}{
	(*HeaderBlock)(nil), // 0: proto.HeaderBlock
	(*HeaderBBox)(nil),  // 1: proto.HeaderBBox
}
var file_HeaderBlock_proto_depIdxs = []int32{
	1, // 0: proto.HeaderBlock.bbox:type_name -> proto.HeaderBBox
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HeaderBlock_proto_init() }
func file_HeaderBlock_proto_init() {
	if File_HeaderBlock_proto != nil {
		return
	}
	file_HeaderBBox_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HeaderBlock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderBlock); i {
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
			RawDescriptor: file_HeaderBlock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HeaderBlock_proto_goTypes,
		DependencyIndexes: file_HeaderBlock_proto_depIdxs,
		MessageInfos:      file_HeaderBlock_proto_msgTypes,
	}.Build()
	File_HeaderBlock_proto = out.File
	file_HeaderBlock_proto_rawDesc = nil
	file_HeaderBlock_proto_goTypes = nil
	file_HeaderBlock_proto_depIdxs = nil
}
