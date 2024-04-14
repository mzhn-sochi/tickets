// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: tickets.proto

package ts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	share "tickets/api/share"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Statuses int32

const (
	Statuses_WAITING_OCR        Statuses = 0
	Statuses_WAITING_VALIDATION Statuses = 1
	Statuses_WAITING_APPROVAL   Statuses = 2
	Statuses_CLOSED             Statuses = 3
	Statuses_REJECTED           Statuses = 4
)

// Enum value maps for Statuses.
var (
	Statuses_name = map[int32]string{
		0: "WAITING_OCR",
		1: "WAITING_VALIDATION",
		2: "WAITING_APPROVAL",
		3: "CLOSED",
		4: "REJECTED",
	}
	Statuses_value = map[string]int32{
		"WAITING_OCR":        0,
		"WAITING_VALIDATION": 1,
		"WAITING_APPROVAL":   2,
		"CLOSED":             3,
		"REJECTED":           4,
	}
)

func (x Statuses) Enum() *Statuses {
	p := new(Statuses)
	*p = x
	return p
}

func (x Statuses) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Statuses) Descriptor() protoreflect.EnumDescriptor {
	return file_tickets_proto_enumTypes[0].Descriptor()
}

func (Statuses) Type() protoreflect.EnumType {
	return &file_tickets_proto_enumTypes[0]
}

func (x Statuses) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Statuses.Descriptor instead.
func (Statuses) EnumDescriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{0}
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"` // uuid v7
	ImageUrl    string `protobuf:"bytes,3,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	ShopAddress string `protobuf:"bytes,4,opt,name=shopAddress,proto3" json:"shopAddress,omitempty"`
	CreatedAt   int64  `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   *int64 `protobuf:"varint,6,opt,name=updatedAt,proto3,oneof" json:"updatedAt,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ticket) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Ticket) GetShopAddress() string {
	if x != nil {
		return x.ShopAddress
	}
	return ""
}

func (x *Ticket) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Ticket) GetUpdatedAt() int64 {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"` // uuid v7
	ImageUrl string `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	ShopAddr string `protobuf:"bytes,3,opt,name=shopAddr,proto3" json:"shopAddr,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreateRequest) GetShopAddr() string {
	if x != nil {
		return x.ShopAddr
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string `protobuf:"bytes,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

type CloseTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string `protobuf:"bytes,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
}

func (x *CloseTicketRequest) Reset() {
	*x = CloseTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseTicketRequest) ProtoMessage() {}

func (x *CloseTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseTicketRequest.ProtoReflect.Descriptor instead.
func (*CloseTicketRequest) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{3}
}

func (x *CloseTicketRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    *Statuses        `protobuf:"varint,1,opt,name=status,proto3,enum=ticket.Statuses,oneof" json:"status,omitempty"`
	TimeRange *share.TimeRange `protobuf:"bytes,2,opt,name=timeRange,proto3,oneof" json:"timeRange,omitempty"`
	UserId    *string          `protobuf:"bytes,3,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{4}
}

func (x *Filter) GetStatus() Statuses {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Statuses_WAITING_OCR
}

func (x *Filter) GetTimeRange() *share.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *Filter) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bounds *share.Bounds `protobuf:"bytes,1,opt,name=bounds,proto3" json:"bounds,omitempty"`
	Filter *Filter       `protobuf:"bytes,2,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{5}
}

func (x *ListRequest) GetBounds() *share.Bounds {
	if x != nil {
		return x.Bounds
	}
	return nil
}

func (x *ListRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickets []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
	Count   int64     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

func (x *ListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type FindByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string `protobuf:"bytes,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
}

func (x *FindByIdRequest) Reset() {
	*x = FindByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tickets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdRequest) ProtoMessage() {}

func (x *FindByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdRequest.ProtoReflect.Descriptor instead.
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return file_tickets_proto_rawDescGZIP(), []int{7}
}

func (x *FindByIdRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

var File_tickets_proto protoreflect.FileDescriptor

var file_tickets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x5f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x22, 0x2c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x33, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x48, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x52, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x2a, 0x63, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x0f,
	0x0a, 0x0b, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x43, 0x52, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x41, 0x49, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x32, 0xe9, 0x01, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickets_proto_rawDescOnce sync.Once
	file_tickets_proto_rawDescData = file_tickets_proto_rawDesc
)

func file_tickets_proto_rawDescGZIP() []byte {
	file_tickets_proto_rawDescOnce.Do(func() {
		file_tickets_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickets_proto_rawDescData)
	})
	return file_tickets_proto_rawDescData
}

var file_tickets_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tickets_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tickets_proto_goTypes = []interface{}{
	(Statuses)(0),              // 0: ticket.Statuses
	(*Ticket)(nil),             // 1: ticket.Ticket
	(*CreateRequest)(nil),      // 2: ticket.CreateRequest
	(*CreateResponse)(nil),     // 3: ticket.CreateResponse
	(*CloseTicketRequest)(nil), // 4: ticket.CloseTicketRequest
	(*Filter)(nil),             // 5: ticket.Filter
	(*ListRequest)(nil),        // 6: ticket.ListRequest
	(*ListResponse)(nil),       // 7: ticket.ListResponse
	(*FindByIdRequest)(nil),    // 8: ticket.FindByIdRequest
	(*share.TimeRange)(nil),    // 9: share.TimeRange
	(*share.Bounds)(nil),       // 10: share.Bounds
	(*share.Empty)(nil),        // 11: share.Empty
}
var file_tickets_proto_depIdxs = []int32{
	0,  // 0: ticket.Filter.status:type_name -> ticket.Statuses
	9,  // 1: ticket.Filter.timeRange:type_name -> share.TimeRange
	10, // 2: ticket.ListRequest.bounds:type_name -> share.Bounds
	5,  // 3: ticket.ListRequest.filter:type_name -> ticket.Filter
	1,  // 4: ticket.ListResponse.tickets:type_name -> ticket.Ticket
	2,  // 5: ticket.TicketService.Create:input_type -> ticket.CreateRequest
	6,  // 6: ticket.TicketService.List:input_type -> ticket.ListRequest
	8,  // 7: ticket.TicketService.FindById:input_type -> ticket.FindByIdRequest
	4,  // 8: ticket.TicketService.CloseTicket:input_type -> ticket.CloseTicketRequest
	3,  // 9: ticket.TicketService.Create:output_type -> ticket.CreateResponse
	7,  // 10: ticket.TicketService.List:output_type -> ticket.ListResponse
	1,  // 11: ticket.TicketService.FindById:output_type -> ticket.Ticket
	11, // 12: ticket.TicketService.CloseTicket:output_type -> share.Empty
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_tickets_proto_init() }
func file_tickets_proto_init() {
	if File_tickets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tickets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_tickets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_tickets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_tickets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseTicketRequest); i {
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
		file_tickets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_tickets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_tickets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_tickets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIdRequest); i {
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
	file_tickets_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_tickets_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_tickets_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tickets_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tickets_proto_goTypes,
		DependencyIndexes: file_tickets_proto_depIdxs,
		EnumInfos:         file_tickets_proto_enumTypes,
		MessageInfos:      file_tickets_proto_msgTypes,
	}.Build()
	File_tickets_proto = out.File
	file_tickets_proto_rawDesc = nil
	file_tickets_proto_goTypes = nil
	file_tickets_proto_depIdxs = nil
}
