// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: chat_service.proto

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

type ChatMeg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId    int32    `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`                              //  发送者Id
	ReceiverId  int32    `protobuf:"varint,2,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`                        // 接收者Id
	MessageType int32    `protobuf:"varint,3,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`                     // 消息类型
	Content     string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                                                 // 发送的内容
	SendType    SendType `protobuf:"varint,5,opt,name=send_type,json=sendType,proto3,enum=qianxia.IMChat.SendType" json:"send_type,omitempty"` // 发送类型
}

func (x *ChatMeg) Reset() {
	*x = ChatMeg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMeg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMeg) ProtoMessage() {}

func (x *ChatMeg) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMeg.ProtoReflect.Descriptor instead.
func (*ChatMeg) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMeg) GetSenderId() int32 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *ChatMeg) GetReceiverId() int32 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *ChatMeg) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *ChatMeg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMeg) GetSendType() SendType {
	if x != nil {
		return x.SendType
	}
	return SendType_USER
}

var File_chat_service_proto protoreflect.FileDescriptor

var file_chat_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d,
	0x43, 0x68, 0x61, 0x74, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x73, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x32, 0x54, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x45, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x17, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x67, 0x1a, 0x17, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78,
	0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x49, 0x4d, 0x43, 0x68, 0x61,
	0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_service_proto_rawDescOnce sync.Once
	file_chat_service_proto_rawDescData = file_chat_service_proto_rawDesc
)

func file_chat_service_proto_rawDescGZIP() []byte {
	file_chat_service_proto_rawDescOnce.Do(func() {
		file_chat_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_service_proto_rawDescData)
	})
	return file_chat_service_proto_rawDescData
}

var file_chat_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chat_service_proto_goTypes = []interface{}{
	(*ChatMeg)(nil), // 0: qianxia.IMChat.ChatMeg
	(SendType)(0),   // 1: qianxia.IMChat.SendType
}
var file_chat_service_proto_depIdxs = []int32{
	1, // 0: qianxia.IMChat.ChatMeg.send_type:type_name -> qianxia.IMChat.SendType
	0, // 1: qianxia.IMChat.ChatService.ChatMessage:input_type -> qianxia.IMChat.ChatMeg
	0, // 2: qianxia.IMChat.ChatService.ChatMessage:output_type -> qianxia.IMChat.ChatMeg
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_service_proto_init() }
func file_chat_service_proto_init() {
	if File_chat_service_proto != nil {
		return
	}
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMeg); i {
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
			RawDescriptor: file_chat_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_service_proto_goTypes,
		DependencyIndexes: file_chat_service_proto_depIdxs,
		MessageInfos:      file_chat_service_proto_msgTypes,
	}.Build()
	File_chat_service_proto = out.File
	file_chat_service_proto_rawDesc = nil
	file_chat_service_proto_goTypes = nil
	file_chat_service_proto_depIdxs = nil
}
