// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: notify/v1/notify.proto

package notifyv1

import (
	v1 "github.com/x0y14/jackal/gen/types/v1"
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

type FetchMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastMessageId int64 `protobuf:"varint,1,opt,name=last_message_id,json=lastMessageId,proto3" json:"last_message_id,omitempty"`
}

func (x *FetchMessageRequest) Reset() {
	*x = FetchMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_v1_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessageRequest) ProtoMessage() {}

func (x *FetchMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notify_v1_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessageRequest.ProtoReflect.Descriptor instead.
func (*FetchMessageRequest) Descriptor() ([]byte, []int) {
	return file_notify_v1_notify_proto_rawDescGZIP(), []int{0}
}

func (x *FetchMessageRequest) GetLastMessageId() int64 {
	if x != nil {
		return x.LastMessageId
	}
	return 0
}

type FetchMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *v1.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FetchMessageResponse) Reset() {
	*x = FetchMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_v1_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessageResponse) ProtoMessage() {}

func (x *FetchMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notify_v1_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessageResponse.ProtoReflect.Descriptor instead.
func (*FetchMessageResponse) Descriptor() ([]byte, []int) {
	return file_notify_v1_notify_proto_rawDescGZIP(), []int{1}
}

func (x *FetchMessageResponse) GetMessage() *v1.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_notify_v1_notify_proto protoreflect.FileDescriptor

var file_notify_v1_notify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x13, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x64, 0x0a,
	0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x91, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x30, 0x79, 0x31, 0x34, 0x2f, 0x6a, 0x61, 0x63, 0x6b, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notify_v1_notify_proto_rawDescOnce sync.Once
	file_notify_v1_notify_proto_rawDescData = file_notify_v1_notify_proto_rawDesc
)

func file_notify_v1_notify_proto_rawDescGZIP() []byte {
	file_notify_v1_notify_proto_rawDescOnce.Do(func() {
		file_notify_v1_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_notify_v1_notify_proto_rawDescData)
	})
	return file_notify_v1_notify_proto_rawDescData
}

var file_notify_v1_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notify_v1_notify_proto_goTypes = []interface{}{
	(*FetchMessageRequest)(nil),  // 0: notify.v1.FetchMessageRequest
	(*FetchMessageResponse)(nil), // 1: notify.v1.FetchMessageResponse
	(*v1.Message)(nil),           // 2: types.v1.Message
}
var file_notify_v1_notify_proto_depIdxs = []int32{
	2, // 0: notify.v1.FetchMessageResponse.message:type_name -> types.v1.Message
	0, // 1: notify.v1.NotifyService.FetchMessage:input_type -> notify.v1.FetchMessageRequest
	1, // 2: notify.v1.NotifyService.FetchMessage:output_type -> notify.v1.FetchMessageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notify_v1_notify_proto_init() }
func file_notify_v1_notify_proto_init() {
	if File_notify_v1_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notify_v1_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchMessageRequest); i {
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
		file_notify_v1_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchMessageResponse); i {
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
			RawDescriptor: file_notify_v1_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notify_v1_notify_proto_goTypes,
		DependencyIndexes: file_notify_v1_notify_proto_depIdxs,
		MessageInfos:      file_notify_v1_notify_proto_msgTypes,
	}.Build()
	File_notify_v1_notify_proto = out.File
	file_notify_v1_notify_proto_rawDesc = nil
	file_notify_v1_notify_proto_goTypes = nil
	file_notify_v1_notify_proto_depIdxs = nil
}
