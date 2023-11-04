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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

	TargetId    int32  `protobuf:"varint,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ApplyType   int32  `protobuf:"varint,3,opt,name=apply_type,json=applyType,proto3" json:"apply_type,omitempty"`
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

func (x *CreateFriendGroupApplyRequest) GetTargetId() int32 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *CreateFriendGroupApplyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateFriendGroupApplyRequest) GetApplyType() int32 {
	if x != nil {
		return x.ApplyType
	}
	return 0
}

type ReplyFriendGroupApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId  int32  `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Status    int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	ApplyType int32  `protobuf:"varint,3,opt,name=apply_type,json=applyType,proto3" json:"apply_type,omitempty"`
	Comment   string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *ReplyFriendGroupApplyRequest) Reset() {
	*x = ReplyFriendGroupApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_group_apply_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyFriendGroupApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyFriendGroupApplyRequest) ProtoMessage() {}

func (x *ReplyFriendGroupApplyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ReplyFriendGroupApplyRequest.ProtoReflect.Descriptor instead.
func (*ReplyFriendGroupApplyRequest) Descriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyFriendGroupApplyRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *ReplyFriendGroupApplyRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ReplyFriendGroupApplyRequest) GetApplyType() int32 {
	if x != nil {
		return x.ApplyType
	}
	return 0
}

func (x *ReplyFriendGroupApplyRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type ListFriendGroupApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyId     int32                  `protobuf:"varint,1,opt,name=apply_id,json=applyId,proto3" json:"apply_id,omitempty"`
	Nickname    string                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ApplyType   int32                  `protobuf:"varint,4,opt,name=apply_type,json=applyType,proto3" json:"apply_type,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ListFriendGroupApplyResponse) Reset() {
	*x = ListFriendGroupApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_group_apply_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendGroupApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendGroupApplyResponse) ProtoMessage() {}

func (x *ListFriendGroupApplyResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListFriendGroupApplyResponse.ProtoReflect.Descriptor instead.
func (*ListFriendGroupApplyResponse) Descriptor() ([]byte, []int) {
	return file_friend_group_apply_service_proto_rawDescGZIP(), []int{2}
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

func (x *ListFriendGroupApplyResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListFriendGroupApplyResponse) GetApplyType() int32 {
	if x != nil {
		return x.ApplyType
	}
	return 0
}

func (x *ListFriendGroupApplyResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_friend_group_apply_service_proto protoreflect.FileDescriptor

var file_friend_group_apply_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68,
	0x61, 0x74, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8c, 0x01,
	0x0a, 0x1c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xd1, 0x01, 0x0a,
	0x1c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x2a, 0x24, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x55,
	0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x91, 0x03, 0x0a, 0x17, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12,
	0x2d, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x7c, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2c,
	0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x71,
	0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x2c, 0x2e, 0x71, 0x69, 0x61, 0x6e, 0x78, 0x69, 0x61, 0x2e, 0x49, 0x4d,
	0x43, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x0b, 0x5a,
	0x09, 0x49, 0x4d, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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
var file_friend_group_apply_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_friend_group_apply_service_proto_goTypes = []interface{}{
	(ApplyType)(0),                        // 0: qianxia.IMChat.ApplyType
	(Status)(0),                           // 1: qianxia.IMChat.Status
	(*CreateFriendGroupApplyRequest)(nil), // 2: qianxia.IMChat.CreateFriendGroupApplyRequest
	(*ReplyFriendGroupApplyRequest)(nil),  // 3: qianxia.IMChat.ReplyFriendGroupApplyRequest
	(*ListFriendGroupApplyResponse)(nil),  // 4: qianxia.IMChat.ListFriendGroupApplyResponse
	(*timestamppb.Timestamp)(nil),         // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                 // 6: google.protobuf.Empty
	(*Response)(nil),                      // 7: qianxia.IMChat.Response
}
var file_friend_group_apply_service_proto_depIdxs = []int32{
	5, // 0: qianxia.IMChat.ListFriendGroupApplyResponse.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: qianxia.IMChat.FriendGroupApplyService.CreateFriendGroupApply:input_type -> qianxia.IMChat.CreateFriendGroupApplyRequest
	3, // 2: qianxia.IMChat.FriendGroupApplyService.ReplyFriendGroupApply:input_type -> qianxia.IMChat.ReplyFriendGroupApplyRequest
	6, // 3: qianxia.IMChat.FriendGroupApplyService.ListFriendGroupApply:input_type -> google.protobuf.Empty
	7, // 4: qianxia.IMChat.FriendGroupApplyService.CreateFriendGroupApply:output_type -> qianxia.IMChat.Response
	7, // 5: qianxia.IMChat.FriendGroupApplyService.ReplyFriendGroupApply:output_type -> qianxia.IMChat.Response
	4, // 6: qianxia.IMChat.FriendGroupApplyService.ListFriendGroupApply:output_type -> qianxia.IMChat.ListFriendGroupApplyResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friend_group_apply_service_proto_init() }
func file_friend_group_apply_service_proto_init() {
	if File_friend_group_apply_service_proto != nil {
		return
	}
	file_response_proto_init()
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
		file_friend_group_apply_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   3,
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