// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: types.proto

package apis

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

type AccountNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNo int64 `protobuf:"varint,1,opt,name=accountNo,proto3" json:"accountNo,omitempty"`
}

func (x *AccountNumber) Reset() {
	*x = AccountNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountNumber) ProtoMessage() {}

func (x *AccountNumber) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountNumber.ProtoReflect.Descriptor instead.
func (*AccountNumber) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *AccountNumber) GetAccountNo() int64 {
	if x != nil {
		return x.AccountNo
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type PageNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PageNumber) Reset() {
	*x = PageNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageNumber) ProtoMessage() {}

func (x *PageNumber) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageNumber.ProtoReflect.Descriptor instead.
func (*PageNumber) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *PageNumber) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type OnlineUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo []*UserInfo `protobuf:"bytes,1,rep,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *OnlineUsers) Reset() {
	*x = OnlineUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineUsers) ProtoMessage() {}

func (x *OnlineUsers) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineUsers.ProtoReflect.Descriptor instead.
func (*OnlineUsers) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *OnlineUsers) GetUserInfo() []*UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"primaryKey;column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: gorm:"column:account_no"
	AccountNo int64 `protobuf:"varint,2,opt,name=accountNo,proto3" json:"accountNo,omitempty"`
	// @gotags: gorm:"column:password"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// @gotags: gorm:"column:nickname"
	NickName string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetAccountNo() int64 {
	if x != nil {
		return x.AccountNo
	}
	return 0
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type ChatRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record []*ChatRecord `protobuf:"bytes,1,rep,name=record,proto3" json:"record,omitempty"`
}

func (x *ChatRecords) Reset() {
	*x = ChatRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRecords) ProtoMessage() {}

func (x *ChatRecords) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRecords.ProtoReflect.Descriptor instead.
func (*ChatRecords) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{5}
}

func (x *ChatRecords) GetRecord() []*ChatRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type ChatRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"primaryKey;column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: gorm:"column:sender_id"
	SenderID int64 `protobuf:"varint,2,opt,name=senderID,proto3" json:"senderID,omitempty"`
	// @gotags: gorm:"column:recipient_id"
	RecipientID int64 `protobuf:"varint,3,opt,name=recipientID,proto3" json:"recipientID,omitempty"`
	// @gotags: gorm:"column:chat_time"
	ChatTime int64 `protobuf:"varint,4,opt,name=chatTime,proto3" json:"chatTime,omitempty"`
	// @gotags: gorm:"column:content"
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatRecord) Reset() {
	*x = ChatRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRecord) ProtoMessage() {}

func (x *ChatRecord) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRecord.ProtoReflect.Descriptor instead.
func (*ChatRecord) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{6}
}

func (x *ChatRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatRecord) GetSenderID() int64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *ChatRecord) GetRecipientID() int64 {
	if x != nil {
		return x.RecipientID
	}
	return 0
}

func (x *ChatRecord) GetChatTime() int64 {
	if x != nil {
		return x.ChatTime
	}
	return 0
}

func (x *ChatRecord) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61,
	0x70, 0x69, 0x73, 0x22, 0x2d, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x0b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x70, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x90, 0x01, 0x0a,
	0x0a, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x68, 0x61,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32,
	0xec, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x07, 0x45, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_types_proto_goTypes = []interface{}{
	(*AccountNumber)(nil), // 0: apis.AccountNumber
	(*Status)(nil),        // 1: apis.Status
	(*PageNumber)(nil),    // 2: apis.PageNumber
	(*OnlineUsers)(nil),   // 3: apis.OnlineUsers
	(*UserInfo)(nil),      // 4: apis.UserInfo
	(*ChatRecords)(nil),   // 5: apis.ChatRecords
	(*ChatRecord)(nil),    // 6: apis.ChatRecord
}
var file_types_proto_depIdxs = []int32{
	4, // 0: apis.OnlineUsers.userInfo:type_name -> apis.UserInfo
	6, // 1: apis.ChatRecords.record:type_name -> apis.ChatRecord
	4, // 2: apis.ChatService.Register:input_type -> apis.UserInfo
	4, // 3: apis.ChatService.Login:input_type -> apis.UserInfo
	2, // 4: apis.ChatService.GetOnlineUsers:input_type -> apis.PageNumber
	0, // 5: apis.ChatService.Chat:input_type -> apis.AccountNumber
	0, // 6: apis.ChatService.ChatHistory:input_type -> apis.AccountNumber
	0, // 7: apis.ChatService.Subscribe:input_type -> apis.AccountNumber
	0, // 8: apis.ChatService.EndChat:input_type -> apis.AccountNumber
	4, // 9: apis.ChatService.Register:output_type -> apis.UserInfo
	4, // 10: apis.ChatService.Login:output_type -> apis.UserInfo
	3, // 11: apis.ChatService.GetOnlineUsers:output_type -> apis.OnlineUsers
	5, // 12: apis.ChatService.Chat:output_type -> apis.ChatRecords
	5, // 13: apis.ChatService.ChatHistory:output_type -> apis.ChatRecords
	1, // 14: apis.ChatService.Subscribe:output_type -> apis.Status
	1, // 15: apis.ChatService.EndChat:output_type -> apis.Status
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountNumber); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageNumber); i {
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
		file_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineUsers); i {
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
		file_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRecords); i {
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
		file_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRecord); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
