// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: friend_cluster_apply_service.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateFriendClusterApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverId int32  `protobuf:"varint,1,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	ApplyDesc  string `protobuf:"bytes,2,opt,name=apply_desc,json=applyDesc,proto3" json:"apply_desc,omitempty"`
	Flag       int32  `protobuf:"varint,3,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *CreateFriendClusterApplyRequest) Reset() {
	*x = CreateFriendClusterApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_cluster_apply_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFriendClusterApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFriendClusterApplyRequest) ProtoMessage() {}

func (x *CreateFriendClusterApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_cluster_apply_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFriendClusterApplyRequest.ProtoReflect.Descriptor instead.
func (*CreateFriendClusterApplyRequest) Descriptor() ([]byte, []int) {
	return file_friend_cluster_apply_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFriendClusterApplyRequest) GetReceiverId() int32 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *CreateFriendClusterApplyRequest) GetApplyDesc() string {
	if x != nil {
		return x.ApplyDesc
	}
	return ""
}

func (x *CreateFriendClusterApplyRequest) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type CreateFriendClusterApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateFriendClusterApplyResponse) Reset() {
	*x = CreateFriendClusterApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_cluster_apply_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFriendClusterApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFriendClusterApplyResponse) ProtoMessage() {}

func (x *CreateFriendClusterApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_cluster_apply_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFriendClusterApplyResponse.ProtoReflect.Descriptor instead.
func (*CreateFriendClusterApplyResponse) Descriptor() ([]byte, []int) {
	return file_friend_cluster_apply_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFriendClusterApplyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListFriendClusterApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyId   int32                  `protobuf:"varint,1,opt,name=apply_id,json=applyId,proto3" json:"apply_id,omitempty"`
	Nickname  string                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	ApplyDesc string                 `protobuf:"bytes,3,opt,name=apply_desc,json=applyDesc,proto3" json:"apply_desc,omitempty"`
	Flag      int32                  `protobuf:"varint,4,opt,name=flag,proto3" json:"flag,omitempty"`
	ApplyTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
}

func (x *ListFriendClusterApplyResponse) Reset() {
	*x = ListFriendClusterApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_cluster_apply_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendClusterApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendClusterApplyResponse) ProtoMessage() {}

func (x *ListFriendClusterApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_cluster_apply_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendClusterApplyResponse.ProtoReflect.Descriptor instead.
func (*ListFriendClusterApplyResponse) Descriptor() ([]byte, []int) {
	return file_friend_cluster_apply_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListFriendClusterApplyResponse) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

func (x *ListFriendClusterApplyResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ListFriendClusterApplyResponse) GetApplyDesc() string {
	if x != nil {
		return x.ApplyDesc
	}
	return ""
}

func (x *ListFriendClusterApplyResponse) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *ListFriendClusterApplyResponse) GetApplyTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ApplyTime
	}
	return nil
}

var File_friend_cluster_apply_service_proto protoreflect.FileDescriptor

var file_friend_cluster_apply_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d,
	0x43, 0x68, 0x61, 0x74, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x75, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x44, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22, 0x3c, 0x0a, 0x20, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x65, 0x73, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x32,
	0xbd, 0x02, 0x0a, 0x19, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9a, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x2e, 0x71, 0x69, 0x61,
	0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x71, 0x69,
	0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e,
	0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d,
	0x43, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42,
	0x0b, 0x5a, 0x09, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_cluster_apply_service_proto_rawDescOnce sync.Once
	file_friend_cluster_apply_service_proto_rawDescData = file_friend_cluster_apply_service_proto_rawDesc
)

func file_friend_cluster_apply_service_proto_rawDescGZIP() []byte {
	file_friend_cluster_apply_service_proto_rawDescOnce.Do(func() {
		file_friend_cluster_apply_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_cluster_apply_service_proto_rawDescData)
	})
	return file_friend_cluster_apply_service_proto_rawDescData
}

var file_friend_cluster_apply_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_friend_cluster_apply_service_proto_goTypes = []interface{}{
	(*CreateFriendClusterApplyRequest)(nil),  // 0: qianxia.IMChat.CreateFriendClusterApplyRequest
	(*CreateFriendClusterApplyResponse)(nil), // 1: qianxia.IMChat.CreateFriendClusterApplyResponse
	(*ListFriendClusterApplyResponse)(nil),   // 2: qianxia.IMChat.ListFriendClusterApplyResponse
	(*timestamppb.Timestamp)(nil),            // 3: google.protobuf.Timestamp
	(*EmptyRequest)(nil),                     // 4: qianxia.IMChat.EmptyRequest
}
var file_friend_cluster_apply_service_proto_depIdxs = []int32{
	3, // 0: qianxia.IMChat.ListFriendClusterApplyResponse.apply_time:type_name -> google.protobuf.Timestamp
	0, // 1: qianxia.IMChat.FriendClusterApplyService.CreateFriendClusterApply:input_type -> qianxia.IMChat.CreateFriendClusterApplyRequest
	4, // 2: qianxia.IMChat.FriendClusterApplyService.ListFriendClusterApply:input_type -> qianxia.IMChat.EmptyRequest
	1, // 3: qianxia.IMChat.FriendClusterApplyService.CreateFriendClusterApply:output_type -> qianxia.IMChat.CreateFriendClusterApplyResponse
	2, // 4: qianxia.IMChat.FriendClusterApplyService.ListFriendClusterApply:output_type -> qianxia.IMChat.ListFriendClusterApplyResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friend_cluster_apply_service_proto_init() }
func file_friend_cluster_apply_service_proto_init() {
	if File_friend_cluster_apply_service_proto != nil {
		return
	}
	file_command_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_friend_cluster_apply_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFriendClusterApplyRequest); i {
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
		file_friend_cluster_apply_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFriendClusterApplyResponse); i {
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
		file_friend_cluster_apply_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendClusterApplyResponse); i {
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
			RawDescriptor: file_friend_cluster_apply_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_cluster_apply_service_proto_goTypes,
		DependencyIndexes: file_friend_cluster_apply_service_proto_depIdxs,
		MessageInfos:      file_friend_cluster_apply_service_proto_msgTypes,
	}.Build()
	File_friend_cluster_apply_service_proto = out.File
	file_friend_cluster_apply_service_proto_rawDesc = nil
	file_friend_cluster_apply_service_proto_goTypes = nil
	file_friend_cluster_apply_service_proto_depIdxs = nil
}
