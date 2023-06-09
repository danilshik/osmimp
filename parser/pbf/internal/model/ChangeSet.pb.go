// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: ChangeSet.proto

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

type ChangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// // Parallel arrays.
	Keys           []uint32    `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"` // String IDs.
	Vals           []uint32    `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"` // String IDs.
	Info           *Info       `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	CreatedAt      *int64      `protobuf:"varint,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	ClosetimeDelta *int64      `protobuf:"varint,9,opt,name=closetime_delta,json=closetimeDelta" json:"closetime_delta,omitempty"`
	Open           *bool       `protobuf:"varint,10,opt,name=open" json:"open,omitempty"`
	Bbox           *HeaderBBox `protobuf:"bytes,11,opt,name=bbox" json:"bbox,omitempty"`
}

func (x *ChangeSet) Reset() {
	*x = ChangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChangeSet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSet) ProtoMessage() {}

func (x *ChangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_ChangeSet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSet.ProtoReflect.Descriptor instead.
func (*ChangeSet) Descriptor() ([]byte, []int) {
	return file_ChangeSet_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeSet) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ChangeSet) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *ChangeSet) GetVals() []uint32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

func (x *ChangeSet) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ChangeSet) GetCreatedAt() int64 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

func (x *ChangeSet) GetClosetimeDelta() int64 {
	if x != nil && x.ClosetimeDelta != nil {
		return *x.ClosetimeDelta
	}
	return 0
}

func (x *ChangeSet) GetOpen() bool {
	if x != nil && x.Open != nil {
		return *x.Open
	}
	return false
}

func (x *ChangeSet) GetBbox() *HeaderBBox {
	if x != nil {
		return x.Bbox
	}
	return nil
}

var File_ChangeSet_proto protoreflect.FileDescriptor

var file_ChangeSet_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x42, 0x42, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x04,
	0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x76, 0x61, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x70, 0x65,
	0x6e, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x62, 0x6f, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x42,
	0x6f, 0x78, 0x52, 0x04, 0x62, 0x62, 0x6f, 0x78, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x6c, 0x73, 0x68, 0x69, 0x6b,
	0x2f, 0x6f, 0x73, 0x6d, 0x69, 0x6d, 0x70, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f,
}

var (
	file_ChangeSet_proto_rawDescOnce sync.Once
	file_ChangeSet_proto_rawDescData = file_ChangeSet_proto_rawDesc
)

func file_ChangeSet_proto_rawDescGZIP() []byte {
	file_ChangeSet_proto_rawDescOnce.Do(func() {
		file_ChangeSet_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChangeSet_proto_rawDescData)
	})
	return file_ChangeSet_proto_rawDescData
}

var file_ChangeSet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChangeSet_proto_goTypes = []interface{}{
	(*ChangeSet)(nil),  // 0: proto.ChangeSet
	(*Info)(nil),       // 1: proto.Info
	(*HeaderBBox)(nil), // 2: proto.HeaderBBox
}
var file_ChangeSet_proto_depIdxs = []int32{
	1, // 0: proto.ChangeSet.info:type_name -> proto.Info
	2, // 1: proto.ChangeSet.bbox:type_name -> proto.HeaderBBox
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChangeSet_proto_init() }
func file_ChangeSet_proto_init() {
	if File_ChangeSet_proto != nil {
		return
	}
	file_HeaderBBox_proto_init()
	file_Info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChangeSet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSet); i {
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
			RawDescriptor: file_ChangeSet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChangeSet_proto_goTypes,
		DependencyIndexes: file_ChangeSet_proto_depIdxs,
		MessageInfos:      file_ChangeSet_proto_msgTypes,
	}.Build()
	File_ChangeSet_proto = out.File
	file_ChangeSet_proto_rawDesc = nil
	file_ChangeSet_proto_goTypes = nil
	file_ChangeSet_proto_depIdxs = nil
}
