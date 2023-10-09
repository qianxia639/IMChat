// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: friend_group_apply_service.proto

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

type ApplyType int32

const (
	ApplyType_FRIEND  ApplyType = 0 // 好友
	ApplyType_CLUSTER ApplyType = 1 // 群组
)

// Enum value maps for ApplyType.
var (
	ApplyType_name = map[int32]string{
		0: "FRIEND",
		1: "CLUSTER",
	}
	ApplyType_value = map[string]int32{
		"FRIEND":  0,
		"CLUSTER": 1,
	}
)

func (x ApplyType) Enum() *ApplyType {
	p := new(ApplyType)
	*p = x
	return p
}

func (x ApplyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplyType) Descriptor() protoreflect.EnumDescriptor {
	return file_friend_group_apply_service_proto_enumTypes[0].Descriptor()
}

func (ApplyType) Type() protoreflect.EnumType {
	return &file_friend_group_apply_service_proto_enumTypes[0]
}

func (x ApplyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplyType.Descriptor instead.
func (ApplyType) EnumDescriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_PENDING  Status = 0 // 等待中
	Status_ACCEPTEd Status = 1 // 接受
	Status_REJECTED Status = 2 // 拒绝
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PENDING",
		1: "ACCEPTEd",
		2: "REJECTED",
	}
	Status_value = map[string]int32{
		"PENDING":  0,
		"ACCEPTEd": 1,
		"REJECTED": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_friend_group_apply_service_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_friend_group_apply_service_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{1}
}

type CreateFriendGroupApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverId int32     `protobuf:"varint,1,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	ApplyDesc  string    `protobuf:"bytes,2,opt,name=apply_desc,json=applyDesc,proto3" json:"apply_desc,omitempty"`
	ApplyType  ApplyType `protobuf:"varint,3,opt,name=apply_type,json=applyType,proto3,enum=qianxia.IMChat.ApplyType" json:"apply_type,omitempty"`
}

func (x *CreateFriendGroupApplyRequest) Reset() {
	*x = CreateFriendGroupApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_group_apply_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFriendGroupApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFriendGroupApplyRequest) ProtoMessage() {}

func (x *CreateFriendGroupApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_group_apply_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFriendGroupApplyRequest.ProtoReflect.Descriptor instead.
func (*CreateFriendGroupApplyRequest) Descriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFriendGroupApplyRequest) GetReceiverId() int32 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *CreateFriendGroupApplyRequest) GetApplyDesc() string {
	if x != nil {
		return x.ApplyDesc
	}
	return ""
}

func (x *CreateFriendGroupApplyRequest) GetApplyType() ApplyType {
	if x != nil {
		return x.ApplyType
	}
	return ApplyType_FRIEND
}

type CreateFriendGroupApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateFriendGroupApplyResponse) Reset() {
	*x = CreateFriendGroupApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_group_apply_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFriendGroupApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFriendGroupApplyResponse) ProtoMessage() {}

func (x *CreateFriendGroupApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_group_apply_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFriendGroupApplyResponse.ProtoReflect.Descriptor instead.
func (*CreateFriendGroupApplyResponse) Descriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFriendGroupApplyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReplyFriendGroupApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId  int32     `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Status    Status    `protobuf:"varint,2,opt,name=status,proto3,enum=qianxia.IMChat.Status" json:"status,omitempty"`
	ApplyType ApplyType `protobuf:"varint,3,opt,name=apply_type,json=applyType,proto3,enum=qianxia.IMChat.ApplyType" json:"apply_type,omitempty"`
	Note      string    `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *ReplyFriendGroupApplyRequest) Reset() {
	*x = ReplyFriendGroupApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_group_apply_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyFriendGroupApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyFriendGroupApplyRequest) ProtoMessage() {}

func (x *ReplyFriendGroupApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friend_group_apply_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyFriendGroupApplyRequest.ProtoReflect.Descriptor instead.
func (*ReplyFriendGroupApplyRequest) Descriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyFriendGroupApplyRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *ReplyFriendGroupApplyRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDING
}

func (x *ReplyFriendGroupApplyRequest) GetApplyType() ApplyType {
	if x != nil {
		return x.ApplyType
	}
	return ApplyType_FRIEND
}

func (x *ReplyFriendGroupApplyRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type ReplyFriendGroupApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReplyFriendGroupApplyResponse) Reset() {
	*x = ReplyFriendGroupApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_group_apply_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyFriendGroupApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyFriendGroupApplyResponse) ProtoMessage() {}

func (x *ReplyFriendGroupApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_group_apply_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyFriendGroupApplyResponse.ProtoReflect.Descriptor instead.
func (*ReplyFriendGroupApplyResponse) Descriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyFriendGroupApplyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListFriendGroupApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyId   int32                  `protobuf:"varint,1,opt,name=apply_id,json=applyId,proto3" json:"apply_id,omitempty"`
	Nickname  string                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	ApplyDesc string                 `protobuf:"bytes,3,opt,name=apply_desc,json=applyDesc,proto3" json:"apply_desc,omitempty"`
	ApplyType int32                  `protobuf:"varint,4,opt,name=apply_type,json=applyType,proto3" json:"apply_type,omitempty"`
	ApplyTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
}

func (x *ListFriendGroupApplyResponse) Reset() {
	*x = ListFriendGroupApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_group_apply_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendGroupApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendGroupApplyResponse) ProtoMessage() {}

func (x *ListFriendGroupApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_group_apply_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendGroupApplyResponse.ProtoReflect.Descriptor instead.
func (*ListFriendGroupApplyResponse) Descriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListFriendGroupApplyResponse) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

func (x *ListFriendGroupApplyResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ListFriendGroupApplyResponse) GetApplyDesc() string {
	if x != nil {
		return x.ApplyDesc
	}
	return ""
}

func (x *ListFriendGroupApplyResponse) GetApplyType() int32 {
	if x != nil {
		return x.ApplyType
	}
	return 0
}

func (x *ListFriendGroupApplyResponse) GetApplyTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ApplyTime
	}
	return nil
}

var File_friend_group_apply_service_proto protoreflect.FileDescriptor

var file_friend_group_apply_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68,
	0x61, 0x74, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x99, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61,
	0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x1e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61,
	0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x71, 0x69, 0x61,
	0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x22, 0x39, 0x0a, 0x1d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xce, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x2a, 0x24, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x55,
	0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xc4, 0x03, 0x0a, 0x17, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x12, 0x2d, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x91, 0x01, 0x0a,
	0x15, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61,
	0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49,
	0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x3a, 0x01, 0x2a,
	0x12, 0x7e, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78,
	0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61,
	0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x30, 0x01,
	0x42, 0x0b, 0x5a, 0x09, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_group_apply_service_proto_rawDescOnce sync.Once
	file_friend_group_apply_service_proto_rawDescData = file_friend_group_apply_service_proto_rawDesc
)

func file_friend_group_apply_service_proto_rawDescGZIP() []byte {
	file_friend_group_apply_service_proto_rawDescOnce.Do(func() {
		file_friend_group_apply_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_group_apply_service_proto_rawDescData)
	})
	return file_friend_group_apply_service_proto_rawDescData
}

var file_friend_group_apply_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_friend_group_apply_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_friend_group_apply_service_proto_goTypes = []interface{}{
	(ApplyType)(0),                         // 0: qianxia.IMChat.ApplyType
	(Status)(0),                            // 1: qianxia.IMChat.Status
	(*CreateFriendGroupApplyRequest)(nil),  // 2: qianxia.IMChat.CreateFriendGroupApplyRequest
	(*CreateFriendGroupApplyResponse)(nil), // 3: qianxia.IMChat.CreateFriendGroupApplyResponse
	(*ReplyFriendGroupApplyRequest)(nil),   // 4: qianxia.IMChat.ReplyFriendGroupApplyRequest
	(*ReplyFriendGroupApplyResponse)(nil),  // 5: qianxia.IMChat.ReplyFriendGroupApplyResponse
	(*ListFriendGroupApplyResponse)(nil),   // 6: qianxia.IMChat.ListFriendGroupApplyResponse
	(*timestamppb.Timestamp)(nil),          // 7: google.protobuf.Timestamp
	(*EmptyRequest)(nil),                   // 8: qianxia.IMChat.EmptyRequest
}
var file_friend_group_apply_service_proto_depIdxs = []int32{
	0, // 0: qianxia.IMChat.CreateFriendGroupApplyRequest.apply_type:type_name -> qianxia.IMChat.ApplyType
	1, // 1: qianxia.IMChat.ReplyFriendGroupApplyRequest.status:type_name -> qianxia.IMChat.Status
	0, // 2: qianxia.IMChat.ReplyFriendGroupApplyRequest.apply_type:type_name -> qianxia.IMChat.ApplyType
	7, // 3: qianxia.IMChat.ListFriendGroupApplyResponse.apply_time:type_name -> google.protobuf.Timestamp
	2, // 4: qianxia.IMChat.FriendGroupApplyService.CreateFriendGroupApply:input_type -> qianxia.IMChat.CreateFriendGroupApplyRequest
	4, // 5: qianxia.IMChat.FriendGroupApplyService.ReplyFriendGroupApply:input_type -> qianxia.IMChat.ReplyFriendGroupApplyRequest
	8, // 6: qianxia.IMChat.FriendGroupApplyService.ListFriendGroupApply:input_type -> qianxia.IMChat.EmptyRequest
	3, // 7: qianxia.IMChat.FriendGroupApplyService.CreateFriendGroupApply:output_type -> qianxia.IMChat.CreateFriendGroupApplyResponse
	5, // 8: qianxia.IMChat.FriendGroupApplyService.ReplyFriendGroupApply:output_type -> qianxia.IMChat.ReplyFriendGroupApplyResponse
	6, // 9: qianxia.IMChat.FriendGroupApplyService.ListFriendGroupApply:output_type -> qianxia.IMChat.ListFriendGroupApplyResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_friend_group_apply_service_proto_init() }
func file_friend_group_apply_service_proto_init() {
	if File_friend_group_apply_service_proto != nil {
		return
	}
	file_command_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_friend_group_apply_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFriendGroupApplyRequest); i {
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
		file_friend_group_apply_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFriendGroupApplyResponse); i {
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
		file_friend_group_apply_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyFriendGroupApplyRequest); i {
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
		file_friend_group_apply_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyFriendGroupApplyResponse); i {
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
		file_friend_group_apply_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendGroupApplyResponse); i {
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
			RawDescriptor: file_friend_group_apply_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_group_apply_service_proto_goTypes,
		DependencyIndexes: file_friend_group_apply_service_proto_depIdxs,
		EnumInfos:         file_friend_group_apply_service_proto_enumTypes,
		MessageInfos:      file_friend_group_apply_service_proto_msgTypes,
	}.Build()
	File_friend_group_apply_service_proto = out.File
	file_friend_group_apply_service_proto_rawDesc = nil
	file_friend_group_apply_service_proto_goTypes = nil
	file_friend_group_apply_service_proto_depIdxs = nil
}