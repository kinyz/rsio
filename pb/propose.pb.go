// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: propose.proto

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

type ProposeType int32

const (
	ProposeType_SyncMessage           ProposeType = 0 //发送消息
	ProposeType_SyncMemoryData        ProposeType = 1 //同步内存数据
	ProposeType_SyncDiskData          ProposeType = 2 //同步硬盘数据
	ProposeType_SyncMemoryAndDiskData ProposeType = 3 //同步内存和硬盘数据
	ProposeType_SyncRemoveMemoryData  ProposeType = 4 //删除内存数据
	ProposeType_SyncRemoveDiskData    ProposeType = 5 //删除硬盘数据
	ProposeType_SyncLock              ProposeType = 6 //同步锁
	ProposeType_TopicAddConsumer      ProposeType = 7 //增加消费者
)

// Enum value maps for ProposeType.
var (
	ProposeType_name = map[int32]string{
		0: "SyncMessage",
		1: "SyncMemoryData",
		2: "SyncDiskData",
		3: "SyncMemoryAndDiskData",
		4: "SyncRemoveMemoryData",
		5: "SyncRemoveDiskData",
		6: "SyncLock",
		7: "TopicAddConsumer",
	}
	ProposeType_value = map[string]int32{
		"SyncMessage":           0,
		"SyncMemoryData":        1,
		"SyncDiskData":          2,
		"SyncMemoryAndDiskData": 3,
		"SyncRemoveMemoryData":  4,
		"SyncRemoveDiskData":    5,
		"SyncLock":              6,
		"TopicAddConsumer":      7,
	}
)

func (x ProposeType) Enum() *ProposeType {
	p := new(ProposeType)
	*p = x
	return p
}

func (x ProposeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProposeType) Descriptor() protoreflect.EnumDescriptor {
	return file_propose_proto_enumTypes[0].Descriptor()
}

func (ProposeType) Type() protoreflect.EnumType {
	return &file_propose_proto_enumTypes[0]
}

func (x ProposeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProposeType.Descriptor instead.
func (ProposeType) EnumDescriptor() ([]byte, []int) {
	return file_propose_proto_rawDescGZIP(), []int{0}
}

type Propose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposeType ProposeType `protobuf:"varint,1,opt,name=ProposeType,proto3,enum=pb.ProposeType" json:"ProposeType,omitempty"`
	Topic       string      `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"` //Topic
	Key         string      `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`     //key
	Value       []byte      `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"` //同步内容
	TimesTamp   int64       `protobuf:"varint,5,opt,name=TimesTamp,proto3" json:"TimesTamp,omitempty"`
}

func (x *Propose) Reset() {
	*x = Propose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_propose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propose) ProtoMessage() {}

func (x *Propose) ProtoReflect() protoreflect.Message {
	mi := &file_propose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propose.ProtoReflect.Descriptor instead.
func (*Propose) Descriptor() ([]byte, []int) {
	return file_propose_proto_rawDescGZIP(), []int{0}
}

func (x *Propose) GetProposeType() ProposeType {
	if x != nil {
		return x.ProposeType
	}
	return ProposeType_SyncMessage
}

func (x *Propose) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Propose) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Propose) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Propose) GetTimesTamp() int64 {
	if x != nil {
		return x.TimesTamp
	}
	return 0
}

var File_propose_proto protoreflect.FileDescriptor

var file_propose_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x98, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x54, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x54, 0x61, 0x6d, 0x70, 0x2a, 0xb5,
	0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x69, 0x73, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x41, 0x6e, 0x64, 0x44, 0x69, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x06,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x10, 0x07, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_propose_proto_rawDescOnce sync.Once
	file_propose_proto_rawDescData = file_propose_proto_rawDesc
)

func file_propose_proto_rawDescGZIP() []byte {
	file_propose_proto_rawDescOnce.Do(func() {
		file_propose_proto_rawDescData = protoimpl.X.CompressGZIP(file_propose_proto_rawDescData)
	})
	return file_propose_proto_rawDescData
}

var file_propose_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_propose_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_propose_proto_goTypes = []interface{}{
	(ProposeType)(0), // 0: pb.ProposeType
	(*Propose)(nil),  // 1: pb.Propose
}
var file_propose_proto_depIdxs = []int32{
	0, // 0: pb.Propose.ProposeType:type_name -> pb.ProposeType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_propose_proto_init() }
func file_propose_proto_init() {
	if File_propose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_propose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propose); i {
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
			RawDescriptor: file_propose_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_propose_proto_goTypes,
		DependencyIndexes: file_propose_proto_depIdxs,
		EnumInfos:         file_propose_proto_enumTypes,
		MessageInfos:      file_propose_proto_msgTypes,
	}.Build()
	File_propose_proto = out.File
	file_propose_proto_rawDesc = nil
	file_propose_proto_goTypes = nil
	file_propose_proto_depIdxs = nil
}