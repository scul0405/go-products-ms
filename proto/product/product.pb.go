// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/product/product.proto

package productService

import (
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CategoryID  string                 `protobuf:"bytes,2,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64                `protobuf:"fixed64,5,opt,name=Price,proto3" json:"Price,omitempty"`
	ImageURL    string                 `protobuf:"bytes,6,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Photos      []string               `protobuf:"bytes,7,rep,name=Photos,proto3" json:"Photos,omitempty"`
	Quantity    int64                  `protobuf:"varint,8,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Rating      int32                  `protobuf:"varint,9,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Created_At  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=Created_At,json=CreatedAt,proto3" json:"Created_At,omitempty"`
	Updated_At  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=Updated_At,json=UpdatedAt,proto3" json:"Updated_At,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Product) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *Product) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *Product) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Product) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Product) GetCreated_At() *timestamppb.Timestamp {
	if x != nil {
		return x.Created_At
	}
	return nil
}

func (x *Product) GetUpdated_At() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated_At
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{1}
}

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID  string   `protobuf:"bytes,1,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64  `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
	ImageURL    string   `protobuf:"bytes,5,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Photos      []string `protobuf:"bytes,6,rep,name=Photos,proto3" json:"Photos,omitempty"`
	Quantity    int64    `protobuf:"varint,7,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Rating      int32    `protobuf:"varint,8,opt,name=Rating,proto3" json:"Rating,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReq) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

func (x *CreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateReq) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateReq) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *CreateReq) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *CreateReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateReq) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=Product,proto3" json:"Product,omitempty"`
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRes) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   string   `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	CategoryID  string   `protobuf:"bytes,2,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64  `protobuf:"fixed64,5,opt,name=Price,proto3" json:"Price,omitempty"`
	ImageURL    string   `protobuf:"bytes,6,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Photos      []string `protobuf:"bytes,7,rep,name=Photos,proto3" json:"Photos,omitempty"`
	Quantity    int64    `protobuf:"varint,8,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Rating      int32    `protobuf:"varint,9,opt,name=Rating,proto3" json:"Rating,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *UpdateReq) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

func (x *UpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateReq) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateReq) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *UpdateReq) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *UpdateReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *UpdateReq) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type UpdateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=Product,proto3" json:"Product,omitempty"`
}

func (x *UpdateRes) Reset() {
	*x = UpdateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRes) ProtoMessage() {}

func (x *UpdateRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRes.ProtoReflect.Descriptor instead.
func (*UpdateRes) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRes) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type GetByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
}

func (x *GetByIDReq) Reset() {
	*x = GetByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDReq) ProtoMessage() {}

func (x *GetByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDReq.ProtoReflect.Descriptor instead.
func (*GetByIDReq) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{6}
}

func (x *GetByIDReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

type GetByIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=Product,proto3" json:"Product,omitempty"`
}

func (x *GetByIDRes) Reset() {
	*x = GetByIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRes) ProtoMessage() {}

func (x *GetByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRes.ProtoReflect.Descriptor instead.
func (*GetByIDRes) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{7}
}

func (x *GetByIDRes) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=Search,proto3" json:"Search,omitempty"`
	Page   int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size   int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{8}
}

func (x *SearchReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *SearchReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SearchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64      `protobuf:"varint,1,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	TotalPages int64      `protobuf:"varint,2,opt,name=TotalPages,proto3" json:"TotalPages,omitempty"`
	Page       int64      `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Size       int64      `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	HasMore    bool       `protobuf:"varint,5,opt,name=HasMore,proto3" json:"HasMore,omitempty"`
	Products   []*Product `protobuf:"bytes,6,rep,name=Products,proto3" json:"Products,omitempty"`
}

func (x *SearchRes) Reset() {
	*x = SearchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_product_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRes) ProtoMessage() {}

func (x *SearchRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRes.ProtoReflect.Descriptor instead.
func (*SearchRes) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{9}
}

func (x *SearchRes) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *SearchRes) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *SearchRes) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchRes) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchRes) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *SearchRes) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_proto_product_product_proto protoreflect.FileDescriptor

var file_proto_product_product_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3,
	0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x39,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xdf, 0x01,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0x3e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0xfd, 0x01, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0x3e, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x2a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x4b, 0x0a, 0x09,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x09, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x9c,
	0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x12, 0x5a,
	0x10, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_product_proto_rawDescOnce sync.Once
	file_proto_product_product_proto_rawDescData = file_proto_product_product_proto_rawDesc
)

func file_proto_product_product_proto_rawDescGZIP() []byte {
	file_proto_product_product_proto_rawDescOnce.Do(func() {
		file_proto_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_product_proto_rawDescData)
	})
	return file_proto_product_product_proto_rawDescData
}

var file_proto_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_product_product_proto_goTypes = []interface{}{
	(*Product)(nil),               // 0: productService.Product
	(*Empty)(nil),                 // 1: productService.Empty
	(*CreateReq)(nil),             // 2: productService.CreateReq
	(*CreateRes)(nil),             // 3: productService.CreateRes
	(*UpdateReq)(nil),             // 4: productService.UpdateReq
	(*UpdateRes)(nil),             // 5: productService.UpdateRes
	(*GetByIDReq)(nil),            // 6: productService.GetByIDReq
	(*GetByIDRes)(nil),            // 7: productService.GetByIDRes
	(*SearchReq)(nil),             // 8: productService.SearchReq
	(*SearchRes)(nil),             // 9: productService.SearchRes
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_proto_product_product_proto_depIdxs = []int32{
	10, // 0: productService.Product.Created_At:type_name -> google.protobuf.Timestamp
	10, // 1: productService.Product.Updated_At:type_name -> google.protobuf.Timestamp
	0,  // 2: productService.CreateRes.Product:type_name -> productService.Product
	0,  // 3: productService.UpdateRes.Product:type_name -> productService.Product
	0,  // 4: productService.GetByIDRes.Product:type_name -> productService.Product
	0,  // 5: productService.SearchRes.Products:type_name -> productService.Product
	4,  // 6: productService.ProductsService.Create:input_type -> productService.UpdateReq
	4,  // 7: productService.ProductsService.Update:input_type -> productService.UpdateReq
	6,  // 8: productService.ProductsService.GetByID:input_type -> productService.GetByIDReq
	8,  // 9: productService.ProductsService.Search:input_type -> productService.SearchReq
	5,  // 10: productService.ProductsService.Create:output_type -> productService.UpdateRes
	5,  // 11: productService.ProductsService.Update:output_type -> productService.UpdateRes
	7,  // 12: productService.ProductsService.GetByID:output_type -> productService.GetByIDRes
	9,  // 13: productService.ProductsService.Search:output_type -> productService.SearchRes
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_product_product_proto_init() }
func file_proto_product_product_proto_init() {
	if File_proto_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_product_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_proto_product_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRes); i {
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
		file_proto_product_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_proto_product_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRes); i {
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
		file_proto_product_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDReq); i {
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
		file_proto_product_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRes); i {
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
		file_proto_product_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_proto_product_product_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRes); i {
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
			RawDescriptor: file_proto_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_product_proto_goTypes,
		DependencyIndexes: file_proto_product_product_proto_depIdxs,
		MessageInfos:      file_proto_product_product_proto_msgTypes,
	}.Build()
	File_proto_product_product_proto = out.File
	file_proto_product_product_proto_rawDesc = nil
	file_proto_product_product_proto_goTypes = nil
	file_proto_product_product_proto_depIdxs = nil
}
