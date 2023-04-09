// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: DenseInfo.proto

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

type DenseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   []int32 `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	Timestamp []int64 `protobuf:"zigzag64,2,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	Changeset []int64 `protobuf:"zigzag64,3,rep,packed,name=changeset" json:"changeset,omitempty"`
	Uid       []int32 `protobuf:"zigzag32,4,rep,packed,name=uid" json:"uid,omitempty"`
	UserSid   []int32 `protobuf:"zigzag32,5,rep,packed,name=user_sid,json=userSid" json:"user_sid,omitempty"`
	Visible   []bool  `protobuf:"varint,6,rep,packed,name=visible" json:"visible,omitempty"`
}

func (x *DenseInfo) Reset() {
	*x = DenseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DenseInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenseInfo) ProtoMessage() {}

func (x *DenseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_DenseInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenseInfo.ProtoReflect.Descriptor instead.
func (*DenseInfo) Descriptor() ([]byte, []int) {
	return file_DenseInfo_proto_rawDescGZIP(), []int{0}
}

func (x *DenseInfo) GetVersion() []int32 {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *DenseInfo) GetTimestamp() []int64 {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DenseInfo) GetChangeset() []int64 {
	if x != nil {
		return x.Changeset
	}
	return nil
}

func (x *DenseInfo) GetUid() []int32 {
	if x != nil {
		return x.Uid
	}
	return nil
}

func (x *DenseInfo) GetUserSid() []int32 {
	if x != nil {
		return x.UserSid
	}
	return nil
}

func (x *DenseInfo) GetVisible() []bool {
	if x != nil {
		return x.Visible
	}
	return nil
}

var File_DenseInfo_proto protoreflect.FileDescriptor

var file_DenseInfo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x6e,
	0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02, 0x10, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02, 0x10, 0x01, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x11, 0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x11,
	0x42, 0x02, 0x10, 0x01, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x08, 0x42, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x6c, 0x73,
	0x68, 0x69, 0x6b, 0x2f, 0x6f, 0x73, 0x6d, 0x69, 0x6d, 0x70, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f,
}

var (
	file_DenseInfo_proto_rawDescOnce sync.Once
	file_DenseInfo_proto_rawDescData = file_DenseInfo_proto_rawDesc
)

func file_DenseInfo_proto_rawDescGZIP() []byte {
	file_DenseInfo_proto_rawDescOnce.Do(func() {
		file_DenseInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_DenseInfo_proto_rawDescData)
	})
	return file_DenseInfo_proto_rawDescData
}

var file_DenseInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DenseInfo_proto_goTypes = []interface{}{
	(*DenseInfo)(nil), // 0: proto.DenseInfo
}
var file_DenseInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DenseInfo_proto_init() }
func file_DenseInfo_proto_init() {
	if File_DenseInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DenseInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenseInfo); i {
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
			RawDescriptor: file_DenseInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DenseInfo_proto_goTypes,
		DependencyIndexes: file_DenseInfo_proto_depIdxs,
		MessageInfos:      file_DenseInfo_proto_msgTypes,
	}.Build()
	File_DenseInfo_proto = out.File
	file_DenseInfo_proto_rawDesc = nil
	file_DenseInfo_proto_goTypes = nil
	file_DenseInfo_proto_depIdxs = nil
}