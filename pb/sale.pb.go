// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: sale.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Sale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Alamat    string               `protobuf:"bytes,3,opt,name=alamat,proto3" json:"alamat,omitempty"`
	Email     string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Telepon   string               `protobuf:"bytes,5,opt,name=telepon,proto3" json:"telepon,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Sale) Reset() {
	*x = Sale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sale) ProtoMessage() {}

func (x *Sale) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sale.ProtoReflect.Descriptor instead.
func (*Sale) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{0}
}

func (x *Sale) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sale) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sale) GetAlamat() string {
	if x != nil {
		return x.Alamat
	}
	return ""
}

func (x *Sale) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Sale) GetTelepon() string {
	if x != nil {
		return x.Telepon
	}
	return ""
}

func (x *Sale) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Sale) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SalesRequest) Reset() {
	*x = SalesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalesRequest) ProtoMessage() {}

func (x *SalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalesRequest.ProtoReflect.Descriptor instead.
func (*SalesRequest) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{1}
}

type SaleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sale *Sale `protobuf:"bytes,1,opt,name=sale,proto3" json:"sale,omitempty"`
}

func (x *SaleResponse) Reset() {
	*x = SaleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleResponse) ProtoMessage() {}

func (x *SaleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleResponse.ProtoReflect.Descriptor instead.
func (*SaleResponse) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{2}
}

func (x *SaleResponse) GetSale() *Sale {
	if x != nil {
		return x.Sale
	}
	return nil
}

type SalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sales []*Sale `protobuf:"bytes,1,rep,name=sales,proto3" json:"sales,omitempty"`
}

func (x *SalesResponse) Reset() {
	*x = SalesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalesResponse) ProtoMessage() {}

func (x *SalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalesResponse.ProtoReflect.Descriptor instead.
func (*SalesResponse) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{3}
}

func (x *SalesResponse) GetSales() []*Sale {
	if x != nil {
		return x.Sales
	}
	return nil
}

type CreateSaleInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alamat  string `protobuf:"bytes,2,opt,name=alamat,proto3" json:"alamat,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Telepon string `protobuf:"bytes,4,opt,name=telepon,proto3" json:"telepon,omitempty"`
}

func (x *CreateSaleInput) Reset() {
	*x = CreateSaleInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSaleInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSaleInput) ProtoMessage() {}

func (x *CreateSaleInput) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSaleInput.ProtoReflect.Descriptor instead.
func (*CreateSaleInput) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSaleInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSaleInput) GetAlamat() string {
	if x != nil {
		return x.Alamat
	}
	return ""
}

func (x *CreateSaleInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateSaleInput) GetTelepon() string {
	if x != nil {
		return x.Telepon
	}
	return ""
}

type UpdateSaleInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Alamat  string `protobuf:"bytes,3,opt,name=alamat,proto3" json:"alamat,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Telepon string `protobuf:"bytes,5,opt,name=telepon,proto3" json:"telepon,omitempty"`
}

func (x *UpdateSaleInput) Reset() {
	*x = UpdateSaleInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSaleInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSaleInput) ProtoMessage() {}

func (x *UpdateSaleInput) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSaleInput.ProtoReflect.Descriptor instead.
func (*UpdateSaleInput) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSaleInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSaleInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSaleInput) GetAlamat() string {
	if x != nil {
		return x.Alamat
	}
	return ""
}

func (x *UpdateSaleInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateSaleInput) GetTelepon() string {
	if x != nil {
		return x.Telepon
	}
	return ""
}

type SaleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaleRequest) Reset() {
	*x = SaleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleRequest) ProtoMessage() {}

func (x *SaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleRequest.ProtoReflect.Descriptor instead.
func (*SaleRequest) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{6}
}

func (x *SaleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSaleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteSaleResponse) Reset() {
	*x = DeleteSaleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSaleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSaleResponse) ProtoMessage() {}

func (x *DeleteSaleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSaleResponse.ProtoReflect.Descriptor instead.
func (*DeleteSaleResponse) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSaleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_sale_proto protoreflect.FileDescriptor

var file_sale_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x04, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x0e, 0x0a, 0x0c,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0c,
	0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04,
	0x73, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x61, 0x6c, 0x65, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x65, 0x22, 0x2f, 0x0a, 0x0d, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x73,
	0x61, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x61, 0x6c, 0x65, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x6e, 0x22, 0x7d, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x0b, 0x53, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x97, 0x02, 0x0a, 0x0b, 0x53, 0x61, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x65, 0x63, 0x68, 0x6f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sale_proto_rawDescOnce sync.Once
	file_sale_proto_rawDescData = file_sale_proto_rawDesc
)

func file_sale_proto_rawDescGZIP() []byte {
	file_sale_proto_rawDescOnce.Do(func() {
		file_sale_proto_rawDescData = protoimpl.X.CompressGZIP(file_sale_proto_rawDescData)
	})
	return file_sale_proto_rawDescData
}

var file_sale_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sale_proto_goTypes = []interface{}{
	(*Sale)(nil),                // 0: pb.Sale
	(*SalesRequest)(nil),        // 1: pb.SalesRequest
	(*SaleResponse)(nil),        // 2: pb.SaleResponse
	(*SalesResponse)(nil),       // 3: pb.SalesResponse
	(*CreateSaleInput)(nil),     // 4: pb.CreateSaleInput
	(*UpdateSaleInput)(nil),     // 5: pb.UpdateSaleInput
	(*SaleRequest)(nil),         // 6: pb.SaleRequest
	(*DeleteSaleResponse)(nil),  // 7: pb.DeleteSaleResponse
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_sale_proto_depIdxs = []int32{
	8, // 0: pb.Sale.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: pb.Sale.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: pb.SaleResponse.sale:type_name -> pb.Sale
	0, // 3: pb.SalesResponse.sales:type_name -> pb.Sale
	4, // 4: pb.SaleService.CreateSale:input_type -> pb.CreateSaleInput
	1, // 5: pb.SaleService.GetSales:input_type -> pb.SalesRequest
	6, // 6: pb.SaleService.GetSale:input_type -> pb.SaleRequest
	5, // 7: pb.SaleService.UpdateSale:input_type -> pb.UpdateSaleInput
	6, // 8: pb.SaleService.DeleteSale:input_type -> pb.SaleRequest
	2, // 9: pb.SaleService.CreateSale:output_type -> pb.SaleResponse
	3, // 10: pb.SaleService.GetSales:output_type -> pb.SalesResponse
	2, // 11: pb.SaleService.GetSale:output_type -> pb.SaleResponse
	2, // 12: pb.SaleService.UpdateSale:output_type -> pb.SaleResponse
	7, // 13: pb.SaleService.DeleteSale:output_type -> pb.DeleteSaleResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sale_proto_init() }
func file_sale_proto_init() {
	if File_sale_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sale_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sale); i {
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
		file_sale_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalesRequest); i {
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
		file_sale_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleResponse); i {
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
		file_sale_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalesResponse); i {
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
		file_sale_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSaleInput); i {
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
		file_sale_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSaleInput); i {
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
		file_sale_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleRequest); i {
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
		file_sale_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSaleResponse); i {
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
			RawDescriptor: file_sale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sale_proto_goTypes,
		DependencyIndexes: file_sale_proto_depIdxs,
		MessageInfos:      file_sale_proto_msgTypes,
	}.Build()
	File_sale_proto = out.File
	file_sale_proto_rawDesc = nil
	file_sale_proto_goTypes = nil
	file_sale_proto_depIdxs = nil
}