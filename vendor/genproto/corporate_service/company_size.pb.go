// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: company_size.proto

package corporate_service

import (
	common "genproto/common"
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

type CreateCompanySizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request     *common.AdminRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Name        map[string]string    `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description map[string]string    `protobuf:"bytes,3,rep,name=description,proto3" json:"description,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	From        int32                `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	To          int32                `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *CreateCompanySizeRequest) Reset() {
	*x = CreateCompanySizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_size_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanySizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanySizeRequest) ProtoMessage() {}

func (x *CreateCompanySizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_size_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanySizeRequest.ProtoReflect.Descriptor instead.
func (*CreateCompanySizeRequest) Descriptor() ([]byte, []int) {
	return file_company_size_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCompanySizeRequest) GetRequest() *common.AdminRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *CreateCompanySizeRequest) GetName() map[string]string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CreateCompanySizeRequest) GetDescription() map[string]string {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *CreateCompanySizeRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *CreateCompanySizeRequest) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

type ShortCompanySize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        map[string]string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description map[string]string `protobuf:"bytes,3,rep,name=description,proto3" json:"description,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	From        int32             `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	To          int32             `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *ShortCompanySize) Reset() {
	*x = ShortCompanySize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_size_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortCompanySize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortCompanySize) ProtoMessage() {}

func (x *ShortCompanySize) ProtoReflect() protoreflect.Message {
	mi := &file_company_size_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortCompanySize.ProtoReflect.Descriptor instead.
func (*ShortCompanySize) Descriptor() ([]byte, []int) {
	return file_company_size_proto_rawDescGZIP(), []int{1}
}

func (x *ShortCompanySize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShortCompanySize) GetName() map[string]string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ShortCompanySize) GetDescription() map[string]string {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *ShortCompanySize) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *ShortCompanySize) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

type GetAllCompanySizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *common.Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Limit   int32           `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page    int32           `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Search  string          `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllCompanySizeRequest) Reset() {
	*x = GetAllCompanySizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_size_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCompanySizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCompanySizeRequest) ProtoMessage() {}

func (x *GetAllCompanySizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_size_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCompanySizeRequest.ProtoReflect.Descriptor instead.
func (*GetAllCompanySizeRequest) Descriptor() ([]byte, []int) {
	return file_company_size_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllCompanySizeRequest) GetRequest() *common.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *GetAllCompanySizeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllCompanySizeRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllCompanySizeRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetAllCompanySizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*ShortCompanySize `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetAllCompanySizeResponse) Reset() {
	*x = GetAllCompanySizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_size_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCompanySizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCompanySizeResponse) ProtoMessage() {}

func (x *GetAllCompanySizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_company_size_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCompanySizeResponse.ProtoReflect.Descriptor instead.
func (*GetAllCompanySizeResponse) Descriptor() ([]byte, []int) {
	return file_company_size_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCompanySizeResponse) GetData() []*ShortCompanySize {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetAllCompanySizeResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_company_size_proto protoreflect.FileDescriptor

var file_company_size_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x6f, 0x1a, 0x37, 0x0a, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xb6, 0x02, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69,
	0x7a, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x74, 0x6f, 0x1a, 0x37, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a,
	0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x22, 0x58, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_size_proto_rawDescOnce sync.Once
	file_company_size_proto_rawDescData = file_company_size_proto_rawDesc
)

func file_company_size_proto_rawDescGZIP() []byte {
	file_company_size_proto_rawDescOnce.Do(func() {
		file_company_size_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_size_proto_rawDescData)
	})
	return file_company_size_proto_rawDescData
}

var file_company_size_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_company_size_proto_goTypes = []interface{}{
	(*CreateCompanySizeRequest)(nil),  // 0: CreateCompanySizeRequest
	(*ShortCompanySize)(nil),          // 1: ShortCompanySize
	(*GetAllCompanySizeRequest)(nil),  // 2: GetAllCompanySizeRequest
	(*GetAllCompanySizeResponse)(nil), // 3: GetAllCompanySizeResponse
	nil,                               // 4: CreateCompanySizeRequest.NameEntry
	nil,                               // 5: CreateCompanySizeRequest.DescriptionEntry
	nil,                               // 6: ShortCompanySize.NameEntry
	nil,                               // 7: ShortCompanySize.DescriptionEntry
	(*common.AdminRequest)(nil),       // 8: AdminRequest
	(*common.Request)(nil),            // 9: Request
}
var file_company_size_proto_depIdxs = []int32{
	8, // 0: CreateCompanySizeRequest.request:type_name -> AdminRequest
	4, // 1: CreateCompanySizeRequest.name:type_name -> CreateCompanySizeRequest.NameEntry
	5, // 2: CreateCompanySizeRequest.description:type_name -> CreateCompanySizeRequest.DescriptionEntry
	6, // 3: ShortCompanySize.name:type_name -> ShortCompanySize.NameEntry
	7, // 4: ShortCompanySize.description:type_name -> ShortCompanySize.DescriptionEntry
	9, // 5: GetAllCompanySizeRequest.request:type_name -> Request
	1, // 6: GetAllCompanySizeResponse.data:type_name -> ShortCompanySize
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_company_size_proto_init() }
func file_company_size_proto_init() {
	if File_company_size_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_company_size_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanySizeRequest); i {
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
		file_company_size_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortCompanySize); i {
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
		file_company_size_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCompanySizeRequest); i {
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
		file_company_size_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCompanySizeResponse); i {
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
			RawDescriptor: file_company_size_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_company_size_proto_goTypes,
		DependencyIndexes: file_company_size_proto_depIdxs,
		MessageInfos:      file_company_size_proto_msgTypes,
	}.Build()
	File_company_size_proto = out.File
	file_company_size_proto_rawDesc = nil
	file_company_size_proto_goTypes = nil
	file_company_size_proto_depIdxs = nil
}
