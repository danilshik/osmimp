// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: Node.proto

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"zigzag64,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"` // String IDs.
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"` // String IDs.
	Info *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`         // May be omitted in omitmeta
	Lat  *int64   `protobuf:"zigzag64,8,req,name=lat" json:"lat,omitempty"`
	Lon  *int64   `protobuf:"zigzag64,9,req,name=lon" json:"lon,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_Node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_Node_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Node) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Node) GetVals() []uint32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

func (x *Node) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Node) GetLat() int64 {
	if x != nil && x.Lat != nil {
		return *x.Lat
	}
	return 0
}

func (x *Node) GetLon() int64 {
	if x != nil && x.Lon != nil {
		return *x.Lon
	}
	return 0
}

var File_Node_proto protoreflect.FileDescriptor

var file_Node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8b, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x12, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x12, 0x16, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x02, 0x28, 0x12, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x02, 0x28, 0x12, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69,
	0x6c, 0x73, 0x68, 0x69, 0x6b, 0x2f, 0x6f, 0x73, 0x6d, 0x69, 0x6d, 0x70, 0x2f, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
}

var (
	file_Node_proto_rawDescOnce sync.Once
	file_Node_proto_rawDescData = file_Node_proto_rawDesc
)

func file_Node_proto_rawDescGZIP() []byte {
	file_Node_proto_rawDescOnce.Do(func() {
		file_Node_proto_rawDescData = protoimpl.X.CompressGZIP(file_Node_proto_rawDescData)
	})
	return file_Node_proto_rawDescData
}

var file_Node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Node_proto_goTypes = []interface{}{
	(*Node)(nil), // 0: proto.Node
	(*Info)(nil), // 1: proto.Info
}
var file_Node_proto_depIdxs = []int32{
	1, // 0: proto.Node.info:type_name -> proto.Info
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Node_proto_init() }
func file_Node_proto_init() {
	if File_Node_proto != nil {
		return
	}
	file_Info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_Node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Node_proto_goTypes,
		DependencyIndexes: file_Node_proto_depIdxs,
		MessageInfos:      file_Node_proto_msgTypes,
	}.Build()
	File_Node_proto = out.File
	file_Node_proto_rawDesc = nil
	file_Node_proto_goTypes = nil
	file_Node_proto_depIdxs = nil
}
