// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: nft/nft.proto

package nft

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Nft struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Collection    string                 `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	Metadata      *structpb.Struct       `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Nft) Reset() {
	*x = Nft{}
	mi := &file_nft_nft_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Nft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nft) ProtoMessage() {}

func (x *Nft) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nft.ProtoReflect.Descriptor instead.
func (*Nft) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{0}
}

func (x *Nft) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Nft) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Nft) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *Nft) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type NftAvailable struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Collection    string                 `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftAvailable) Reset() {
	*x = NftAvailable{}
	mi := &file_nft_nft_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftAvailable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftAvailable) ProtoMessage() {}

func (x *NftAvailable) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftAvailable.ProtoReflect.Descriptor instead.
func (*NftAvailable) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{1}
}

func (x *NftAvailable) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NftAvailable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NftAvailable) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *NftAvailable) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type NftPopular struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Collection    string                 `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	TradeVolume   float32                `protobuf:"fixed32,4,opt,name=trade_volume,json=tradeVolume,proto3" json:"trade_volume,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftPopular) Reset() {
	*x = NftPopular{}
	mi := &file_nft_nft_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftPopular) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftPopular) ProtoMessage() {}

func (x *NftPopular) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftPopular.ProtoReflect.Descriptor instead.
func (*NftPopular) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{2}
}

func (x *NftPopular) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NftPopular) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NftPopular) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *NftPopular) GetTradeVolume() float32 {
	if x != nil {
		return x.TradeVolume
	}
	return 0
}

type NftsAtAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     string                 `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftsAtAccountRequest) Reset() {
	*x = NftsAtAccountRequest{}
	mi := &file_nft_nft_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftsAtAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftsAtAccountRequest) ProtoMessage() {}

func (x *NftsAtAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftsAtAccountRequest.ProtoReflect.Descriptor instead.
func (*NftsAtAccountRequest) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{3}
}

func (x *NftsAtAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type NftsAtAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nfts          []*Nft                 `protobuf:"bytes,1,rep,name=nfts,proto3" json:"nfts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftsAtAccountResponse) Reset() {
	*x = NftsAtAccountResponse{}
	mi := &file_nft_nft_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftsAtAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftsAtAccountResponse) ProtoMessage() {}

func (x *NftsAtAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftsAtAccountResponse.ProtoReflect.Descriptor instead.
func (*NftsAtAccountResponse) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{4}
}

func (x *NftsAtAccountResponse) GetNfts() []*Nft {
	if x != nil {
		return x.Nfts
	}
	return nil
}

type NftsAvailableRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CollectionId  string                 `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftsAvailableRequest) Reset() {
	*x = NftsAvailableRequest{}
	mi := &file_nft_nft_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftsAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftsAvailableRequest) ProtoMessage() {}

func (x *NftsAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftsAvailableRequest.ProtoReflect.Descriptor instead.
func (*NftsAvailableRequest) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{5}
}

func (x *NftsAvailableRequest) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

type NftsAvailableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nfts          []*NftAvailable        `protobuf:"bytes,1,rep,name=nfts,proto3" json:"nfts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftsAvailableResponse) Reset() {
	*x = NftsAvailableResponse{}
	mi := &file_nft_nft_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftsAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftsAvailableResponse) ProtoMessage() {}

func (x *NftsAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftsAvailableResponse.ProtoReflect.Descriptor instead.
func (*NftsAvailableResponse) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{6}
}

func (x *NftsAvailableResponse) GetNfts() []*NftAvailable {
	if x != nil {
		return x.Nfts
	}
	return nil
}

type NftsPopularRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftsPopularRequest) Reset() {
	*x = NftsPopularRequest{}
	mi := &file_nft_nft_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftsPopularRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftsPopularRequest) ProtoMessage() {}

func (x *NftsPopularRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftsPopularRequest.ProtoReflect.Descriptor instead.
func (*NftsPopularRequest) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{7}
}

type NftsPopularResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nfts          []*NftPopular          `protobuf:"bytes,1,rep,name=nfts,proto3" json:"nfts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NftsPopularResponse) Reset() {
	*x = NftsPopularResponse{}
	mi := &file_nft_nft_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NftsPopularResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftsPopularResponse) ProtoMessage() {}

func (x *NftsPopularResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nft_nft_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftsPopularResponse.ProtoReflect.Descriptor instead.
func (*NftsPopularResponse) Descriptor() ([]byte, []int) {
	return file_nft_nft_proto_rawDescGZIP(), []int{8}
}

func (x *NftsPopularResponse) GetNfts() []*NftPopular {
	if x != nil {
		return x.Nfts
	}
	return nil
}

var File_nft_nft_proto protoreflect.FileDescriptor

var file_nft_nft_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x6e, 0x66, 0x74, 0x2f, 0x6e, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6e, 0x66, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x03, 0x4e, 0x66, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x68, 0x0a, 0x0c, 0x4e, 0x66, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x73, 0x0a, 0x0a,
	0x4e, 0x66, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x22, 0x35, 0x0a, 0x14, 0x4e, 0x66, 0x74, 0x73, 0x41, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x15, 0x4e, 0x66, 0x74, 0x73,
	0x41, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x66, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x4e, 0x66, 0x74, 0x52, 0x04, 0x6e, 0x66, 0x74, 0x73, 0x22,
	0x3b, 0x0a, 0x14, 0x4e, 0x66, 0x74, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x15,
	0x4e, 0x66, 0x74, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x66, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x4e, 0x66, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x04, 0x6e, 0x66, 0x74, 0x73, 0x22, 0x14, 0x0a, 0x12,
	0x4e, 0x66, 0x74, 0x73, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3a, 0x0a, 0x13, 0x4e, 0x66, 0x74, 0x73, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x66, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x4e, 0x66,
	0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x52, 0x04, 0x6e, 0x66, 0x74, 0x73, 0x42, 0x62,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x66, 0x74, 0x42, 0x08, 0x4e, 0x66, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x2d, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x6e, 0x66, 0x74, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02,
	0x03, 0x4e, 0x66, 0x74, 0xca, 0x02, 0x03, 0x4e, 0x66, 0x74, 0xe2, 0x02, 0x0f, 0x4e, 0x66, 0x74,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03, 0x4e,
	0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nft_nft_proto_rawDescOnce sync.Once
	file_nft_nft_proto_rawDescData []byte
)

func file_nft_nft_proto_rawDescGZIP() []byte {
	file_nft_nft_proto_rawDescOnce.Do(func() {
		file_nft_nft_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nft_nft_proto_rawDesc), len(file_nft_nft_proto_rawDesc)))
	})
	return file_nft_nft_proto_rawDescData
}

var file_nft_nft_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_nft_nft_proto_goTypes = []any{
	(*Nft)(nil),                   // 0: nft.Nft
	(*NftAvailable)(nil),          // 1: nft.NftAvailable
	(*NftPopular)(nil),            // 2: nft.NftPopular
	(*NftsAtAccountRequest)(nil),  // 3: nft.NftsAtAccountRequest
	(*NftsAtAccountResponse)(nil), // 4: nft.NftsAtAccountResponse
	(*NftsAvailableRequest)(nil),  // 5: nft.NftsAvailableRequest
	(*NftsAvailableResponse)(nil), // 6: nft.NftsAvailableResponse
	(*NftsPopularRequest)(nil),    // 7: nft.NftsPopularRequest
	(*NftsPopularResponse)(nil),   // 8: nft.NftsPopularResponse
	(*structpb.Struct)(nil),       // 9: google.protobuf.Struct
}
var file_nft_nft_proto_depIdxs = []int32{
	9, // 0: nft.Nft.metadata:type_name -> google.protobuf.Struct
	0, // 1: nft.NftsAtAccountResponse.nfts:type_name -> nft.Nft
	1, // 2: nft.NftsAvailableResponse.nfts:type_name -> nft.NftAvailable
	2, // 3: nft.NftsPopularResponse.nfts:type_name -> nft.NftPopular
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nft_nft_proto_init() }
func file_nft_nft_proto_init() {
	if File_nft_nft_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nft_nft_proto_rawDesc), len(file_nft_nft_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nft_nft_proto_goTypes,
		DependencyIndexes: file_nft_nft_proto_depIdxs,
		MessageInfos:      file_nft_nft_proto_msgTypes,
	}.Build()
	File_nft_nft_proto = out.File
	file_nft_nft_proto_goTypes = nil
	file_nft_nft_proto_depIdxs = nil
}
