// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmint/nftmint.proto

package types

import (
	fmt "fmt"
	github_com_UnUniFi_chain_types "github.com/UnUniFi/chain/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MintingPermission int32

const (
	MintingPermission_OnlyOwner MintingPermission = 0
	MintingPermission_Anyone    MintingPermission = 1
)

var MintingPermission_name = map[int32]string{
	0: "OnlyOwner",
	1: "Anyone",
}

var MintingPermission_value = map[string]int32{
	"OnlyOwner": 0,
	"Anyone":    1,
}

func (x MintingPermission) String() string {
	return proto.EnumName(MintingPermission_name, int32(x))
}

func (MintingPermission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_828f2ddeb82d86e5, []int{0}
}

type ClassAttributes struct {
	ClassId           string                                          `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Owner             github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,2,opt,name=owner,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"owner" yaml:"owner"`
	BaseTokenUri      string                                          `protobuf:"bytes,3,opt,name=base_token_uri,json=baseTokenUri,proto3" json:"base_token_uri,omitempty"`
	MintingPermission MintingPermission                               `protobuf:"varint,4,opt,name=minting_permission,json=mintingPermission,proto3,enum=ununifi.nftmint.MintingPermission" json:"minting_permission,omitempty"`
	TokenSupplyCap    uint64                                          `protobuf:"varint,5,opt,name=token_supply_cap,json=tokenSupplyCap,proto3" json:"token_supply_cap,omitempty"`
}

func (m *ClassAttributes) Reset()         { *m = ClassAttributes{} }
func (m *ClassAttributes) String() string { return proto.CompactTextString(m) }
func (*ClassAttributes) ProtoMessage()    {}
func (*ClassAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_828f2ddeb82d86e5, []int{0}
}
func (m *ClassAttributes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassAttributes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassAttributes.Merge(m, src)
}
func (m *ClassAttributes) XXX_Size() int {
	return m.Size()
}
func (m *ClassAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ClassAttributes proto.InternalMessageInfo

func (m *ClassAttributes) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *ClassAttributes) GetBaseTokenUri() string {
	if m != nil {
		return m.BaseTokenUri
	}
	return ""
}

func (m *ClassAttributes) GetMintingPermission() MintingPermission {
	if m != nil {
		return m.MintingPermission
	}
	return MintingPermission_OnlyOwner
}

func (m *ClassAttributes) GetTokenSupplyCap() uint64 {
	if m != nil {
		return m.TokenSupplyCap
	}
	return 0
}

type OwningClassIdList struct {
	Owner   github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,1,opt,name=owner,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"owner" yaml:"owner"`
	ClassId []string                                        `protobuf:"bytes,2,rep,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (m *OwningClassIdList) Reset()         { *m = OwningClassIdList{} }
func (m *OwningClassIdList) String() string { return proto.CompactTextString(m) }
func (*OwningClassIdList) ProtoMessage()    {}
func (*OwningClassIdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_828f2ddeb82d86e5, []int{1}
}
func (m *OwningClassIdList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwningClassIdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwningClassIdList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwningClassIdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwningClassIdList.Merge(m, src)
}
func (m *OwningClassIdList) XXX_Size() int {
	return m.Size()
}
func (m *OwningClassIdList) XXX_DiscardUnknown() {
	xxx_messageInfo_OwningClassIdList.DiscardUnknown(m)
}

var xxx_messageInfo_OwningClassIdList proto.InternalMessageInfo

func (m *OwningClassIdList) GetClassId() []string {
	if m != nil {
		return m.ClassId
	}
	return nil
}

type ClassNameIdList struct {
	ClassName string   `protobuf:"bytes,1,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ClassId   []string `protobuf:"bytes,2,rep,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (m *ClassNameIdList) Reset()         { *m = ClassNameIdList{} }
func (m *ClassNameIdList) String() string { return proto.CompactTextString(m) }
func (*ClassNameIdList) ProtoMessage()    {}
func (*ClassNameIdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_828f2ddeb82d86e5, []int{2}
}
func (m *ClassNameIdList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassNameIdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassNameIdList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassNameIdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassNameIdList.Merge(m, src)
}
func (m *ClassNameIdList) XXX_Size() int {
	return m.Size()
}
func (m *ClassNameIdList) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassNameIdList.DiscardUnknown(m)
}

var xxx_messageInfo_ClassNameIdList proto.InternalMessageInfo

func (m *ClassNameIdList) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *ClassNameIdList) GetClassId() []string {
	if m != nil {
		return m.ClassId
	}
	return nil
}

type Params struct {
	MaxNFTSupplyCap   uint64 `protobuf:"varint,1,opt,name=MaxNFTSupplyCap,proto3" json:"MaxNFTSupplyCap,omitempty"`
	MinClassNameLen   uint64 `protobuf:"varint,2,opt,name=MinClassNameLen,proto3" json:"MinClassNameLen,omitempty"`
	MaxClassNameLen   uint64 `protobuf:"varint,3,opt,name=MaxClassNameLen,proto3" json:"MaxClassNameLen,omitempty"`
	MinUriLen         uint64 `protobuf:"varint,4,opt,name=MinUriLen,proto3" json:"MinUriLen,omitempty"`
	MaxUriLen         uint64 `protobuf:"varint,5,opt,name=MaxUriLen,proto3" json:"MaxUriLen,omitempty"`
	MaxSymbolLen      uint64 `protobuf:"varint,6,opt,name=MaxSymbolLen,proto3" json:"MaxSymbolLen,omitempty"`
	MaxDescriptionLen uint64 `protobuf:"varint,7,opt,name=MaxDescriptionLen,proto3" json:"MaxDescriptionLen,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_828f2ddeb82d86e5, []int{3}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxNFTSupplyCap() uint64 {
	if m != nil {
		return m.MaxNFTSupplyCap
	}
	return 0
}

func (m *Params) GetMinClassNameLen() uint64 {
	if m != nil {
		return m.MinClassNameLen
	}
	return 0
}

func (m *Params) GetMaxClassNameLen() uint64 {
	if m != nil {
		return m.MaxClassNameLen
	}
	return 0
}

func (m *Params) GetMinUriLen() uint64 {
	if m != nil {
		return m.MinUriLen
	}
	return 0
}

func (m *Params) GetMaxUriLen() uint64 {
	if m != nil {
		return m.MaxUriLen
	}
	return 0
}

func (m *Params) GetMaxSymbolLen() uint64 {
	if m != nil {
		return m.MaxSymbolLen
	}
	return 0
}

func (m *Params) GetMaxDescriptionLen() uint64 {
	if m != nil {
		return m.MaxDescriptionLen
	}
	return 0
}

func init() {
	proto.RegisterEnum("ununifi.nftmint.MintingPermission", MintingPermission_name, MintingPermission_value)
	proto.RegisterType((*ClassAttributes)(nil), "ununifi.nftmint.ClassAttributes")
	proto.RegisterType((*OwningClassIdList)(nil), "ununifi.nftmint.OwningClassIdList")
	proto.RegisterType((*ClassNameIdList)(nil), "ununifi.nftmint.ClassNameIdList")
	proto.RegisterType((*Params)(nil), "ununifi.nftmint.Params")
}

func init() { proto.RegisterFile("nftmint/nftmint.proto", fileDescriptor_828f2ddeb82d86e5) }

var fileDescriptor_828f2ddeb82d86e5 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x6f, 0xd3, 0x3c,
	0x18, 0xc0, 0x9b, 0xae, 0xeb, 0xde, 0x5a, 0x7d, 0xfb, 0xc7, 0x02, 0xa9, 0x4c, 0x90, 0x56, 0x11,
	0x87, 0x08, 0xa1, 0x44, 0x82, 0x1b, 0xb7, 0xb6, 0x68, 0x08, 0xb1, 0xae, 0xa3, 0x5d, 0x2f, 0x48,
	0xa8, 0x72, 0x52, 0x37, 0xb3, 0x88, 0xed, 0x28, 0x76, 0xb4, 0xe6, 0x43, 0x20, 0x71, 0xe0, 0x73,
	0xf0, 0x39, 0x76, 0xdc, 0x11, 0x71, 0xa8, 0x50, 0xfb, 0x0d, 0xf8, 0x04, 0xc8, 0x4e, 0xda, 0xad,
	0x9d, 0xe0, 0xc6, 0x29, 0xf1, 0xef, 0xf9, 0xe9, 0xb1, 0xfd, 0x3c, 0x8f, 0xc1, 0x43, 0x36, 0x97,
	0x94, 0x30, 0xe9, 0xe6, 0x5f, 0x27, 0x8a, 0xb9, 0xe4, 0xb0, 0x9e, 0xb0, 0x84, 0x91, 0x39, 0x71,
	0x72, 0x7c, 0xfc, 0x20, 0xe0, 0x01, 0xd7, 0x31, 0x57, 0xfd, 0x65, 0xda, 0x71, 0x3b, 0xe0, 0x3c,
	0x08, 0xb1, 0xab, 0x57, 0x5e, 0x32, 0x77, 0x25, 0xa1, 0x58, 0x48, 0x44, 0xa3, 0x4c, 0xb0, 0xbe,
	0x15, 0x41, 0xbd, 0x1f, 0x22, 0x21, 0xba, 0x52, 0xc6, 0xc4, 0x4b, 0x24, 0x16, 0xf0, 0x11, 0xf8,
	0xcf, 0x57, 0x68, 0x4a, 0x66, 0x2d, 0xa3, 0x63, 0xd8, 0x95, 0xd1, 0x91, 0x5e, 0xbf, 0x9d, 0xc1,
	0x8f, 0xe0, 0x90, 0x5f, 0x31, 0x1c, 0xb7, 0x8a, 0x8a, 0xf7, 0xde, 0x5c, 0x2f, 0xdb, 0x85, 0x1f,
	0xcb, 0xb6, 0x1b, 0x10, 0x79, 0x99, 0x78, 0x8e, 0xcf, 0xa9, 0x3b, 0x61, 0x13, 0x46, 0x4e, 0x88,
	0xeb, 0x5f, 0x22, 0xc2, 0x5c, 0x99, 0x46, 0x58, 0x38, 0x63, 0x19, 0x13, 0x16, 0x74, 0x7d, 0xbf,
	0x3b, 0x9b, 0xc5, 0x58, 0x88, 0x5f, 0xcb, 0x76, 0x35, 0x45, 0x34, 0x7c, 0x65, 0xe9, 0x6c, 0xd6,
	0x28, 0xcb, 0x0a, 0x9f, 0x82, 0x9a, 0x87, 0x04, 0x9e, 0x4a, 0xfe, 0x09, 0xb3, 0x69, 0x12, 0x93,
	0xd6, 0x81, 0xde, 0xbf, 0xaa, 0xe8, 0x85, 0x82, 0x93, 0x98, 0xc0, 0xf7, 0x00, 0xaa, 0x2b, 0x13,
	0x16, 0x4c, 0x23, 0x1c, 0x53, 0x22, 0x04, 0xe1, 0xac, 0x55, 0xea, 0x18, 0x76, 0xed, 0x85, 0xe5,
	0xec, 0x15, 0xc6, 0x19, 0x64, 0xea, 0xf9, 0xd6, 0x1c, 0x35, 0xe9, 0x3e, 0x82, 0x36, 0x68, 0x64,
	0x7b, 0x8a, 0x24, 0x8a, 0xc2, 0x74, 0xea, 0xa3, 0xa8, 0x75, 0xd8, 0x31, 0xec, 0xd2, 0xa8, 0xa6,
	0xf9, 0x58, 0xe3, 0x3e, 0x8a, 0xac, 0xcf, 0x06, 0x68, 0x0e, 0xaf, 0x18, 0x61, 0x41, 0x3f, 0xab,
	0xc9, 0x29, 0x11, 0xf2, 0xb6, 0x2e, 0xc6, 0x3f, 0xa9, 0xcb, 0xdd, 0x8e, 0x14, 0x3b, 0x07, 0x77,
	0x3a, 0x62, 0xbd, 0xcb, 0xfb, 0x77, 0x86, 0x28, 0xce, 0x0f, 0xf3, 0x04, 0x80, 0xcc, 0x66, 0x88,
	0xe2, 0xbc, 0x83, 0x15, 0x7f, 0x23, 0xfd, 0x2d, 0xd9, 0xd7, 0x22, 0x28, 0x9f, 0xa3, 0x18, 0x51,
	0x01, 0x6d, 0x50, 0x1f, 0xa0, 0xc5, 0xd9, 0xc9, 0xc5, 0xf6, 0xea, 0x3a, 0x53, 0x69, 0xb4, 0x8f,
	0xb5, 0x49, 0xd8, 0xf6, 0x10, 0xa7, 0x98, 0xe9, 0xe9, 0x50, 0xe6, 0x2e, 0xce, 0x73, 0xee, 0x98,
	0x07, 0xdb, 0x9c, 0x3b, 0xe6, 0x63, 0x50, 0x19, 0x10, 0xd5, 0x6c, 0xe5, 0x94, 0xb4, 0x73, 0x0b,
	0x74, 0x14, 0x2d, 0xf2, 0xe8, 0x61, 0x1e, 0xdd, 0x00, 0x68, 0x81, 0xea, 0x00, 0x2d, 0xc6, 0x29,
	0xf5, 0x78, 0xa8, 0x84, 0xb2, 0x16, 0x76, 0x18, 0x7c, 0x0e, 0x9a, 0x03, 0xb4, 0x78, 0x8d, 0x85,
	0x1f, 0x93, 0x48, 0x12, 0xce, 0x94, 0x78, 0xa4, 0xc5, 0xfb, 0x81, 0x67, 0x0e, 0x68, 0xde, 0x9b,
	0x22, 0xf8, 0x3f, 0xa8, 0x0c, 0x59, 0x98, 0x0e, 0x55, 0x83, 0x1a, 0x05, 0x08, 0x40, 0xb9, 0xcb,
	0x52, 0xce, 0x70, 0xc3, 0xe8, 0xf5, 0xae, 0x57, 0xa6, 0x71, 0xb3, 0x32, 0x8d, 0x9f, 0x2b, 0xd3,
	0xf8, 0xb2, 0x36, 0x0b, 0x37, 0x6b, 0xb3, 0xf0, 0x7d, 0x6d, 0x16, 0x3e, 0xd8, 0x7f, 0x1c, 0x88,
	0xc5, 0xe6, 0x81, 0x67, 0xa3, 0xe1, 0x95, 0xf5, 0xfb, 0x7c, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xa0, 0x50, 0xc3, 0x8d, 0x00, 0x04, 0x00, 0x00,
}

func (m *ClassAttributes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassAttributes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassAttributes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TokenSupplyCap != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.TokenSupplyCap))
		i--
		dAtA[i] = 0x28
	}
	if m.MintingPermission != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MintingPermission))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BaseTokenUri) > 0 {
		i -= len(m.BaseTokenUri)
		copy(dAtA[i:], m.BaseTokenUri)
		i = encodeVarintNftmint(dAtA, i, uint64(len(m.BaseTokenUri)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Owner.Size()
		i -= size
		if _, err := m.Owner.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNftmint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintNftmint(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OwningClassIdList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwningClassIdList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwningClassIdList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassId) > 0 {
		for iNdEx := len(m.ClassId) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClassId[iNdEx])
			copy(dAtA[i:], m.ClassId[iNdEx])
			i = encodeVarintNftmint(dAtA, i, uint64(len(m.ClassId[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.Owner.Size()
		i -= size
		if _, err := m.Owner.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNftmint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClassNameIdList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassNameIdList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassNameIdList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassId) > 0 {
		for iNdEx := len(m.ClassId) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClassId[iNdEx])
			copy(dAtA[i:], m.ClassId[iNdEx])
			i = encodeVarintNftmint(dAtA, i, uint64(len(m.ClassId[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClassName) > 0 {
		i -= len(m.ClassName)
		copy(dAtA[i:], m.ClassName)
		i = encodeVarintNftmint(dAtA, i, uint64(len(m.ClassName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxDescriptionLen != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MaxDescriptionLen))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxSymbolLen != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MaxSymbolLen))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxUriLen != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MaxUriLen))
		i--
		dAtA[i] = 0x28
	}
	if m.MinUriLen != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MinUriLen))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxClassNameLen != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MaxClassNameLen))
		i--
		dAtA[i] = 0x18
	}
	if m.MinClassNameLen != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MinClassNameLen))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxNFTSupplyCap != 0 {
		i = encodeVarintNftmint(dAtA, i, uint64(m.MaxNFTSupplyCap))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftmint(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftmint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassAttributes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovNftmint(uint64(l))
	}
	l = m.Owner.Size()
	n += 1 + l + sovNftmint(uint64(l))
	l = len(m.BaseTokenUri)
	if l > 0 {
		n += 1 + l + sovNftmint(uint64(l))
	}
	if m.MintingPermission != 0 {
		n += 1 + sovNftmint(uint64(m.MintingPermission))
	}
	if m.TokenSupplyCap != 0 {
		n += 1 + sovNftmint(uint64(m.TokenSupplyCap))
	}
	return n
}

func (m *OwningClassIdList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Owner.Size()
	n += 1 + l + sovNftmint(uint64(l))
	if len(m.ClassId) > 0 {
		for _, s := range m.ClassId {
			l = len(s)
			n += 1 + l + sovNftmint(uint64(l))
		}
	}
	return n
}

func (m *ClassNameIdList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassName)
	if l > 0 {
		n += 1 + l + sovNftmint(uint64(l))
	}
	if len(m.ClassId) > 0 {
		for _, s := range m.ClassId {
			l = len(s)
			n += 1 + l + sovNftmint(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxNFTSupplyCap != 0 {
		n += 1 + sovNftmint(uint64(m.MaxNFTSupplyCap))
	}
	if m.MinClassNameLen != 0 {
		n += 1 + sovNftmint(uint64(m.MinClassNameLen))
	}
	if m.MaxClassNameLen != 0 {
		n += 1 + sovNftmint(uint64(m.MaxClassNameLen))
	}
	if m.MinUriLen != 0 {
		n += 1 + sovNftmint(uint64(m.MinUriLen))
	}
	if m.MaxUriLen != 0 {
		n += 1 + sovNftmint(uint64(m.MaxUriLen))
	}
	if m.MaxSymbolLen != 0 {
		n += 1 + sovNftmint(uint64(m.MaxSymbolLen))
	}
	if m.MaxDescriptionLen != 0 {
		n += 1 + sovNftmint(uint64(m.MaxDescriptionLen))
	}
	return n
}

func sovNftmint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftmint(x uint64) (n int) {
	return sovNftmint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassAttributes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftmint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClassAttributes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassAttributes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftmint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftmint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftmint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftmint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseTokenUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftmint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftmint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseTokenUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingPermission", wireType)
			}
			m.MintingPermission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintingPermission |= MintingPermission(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenSupplyCap", wireType)
			}
			m.TokenSupplyCap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenSupplyCap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNftmint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftmint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OwningClassIdList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftmint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OwningClassIdList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwningClassIdList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftmint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftmint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftmint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftmint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = append(m.ClassId, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftmint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftmint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClassNameIdList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftmint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClassNameIdList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassNameIdList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftmint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftmint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNftmint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftmint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = append(m.ClassId, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftmint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftmint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftmint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNFTSupplyCap", wireType)
			}
			m.MaxNFTSupplyCap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNFTSupplyCap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinClassNameLen", wireType)
			}
			m.MinClassNameLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinClassNameLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxClassNameLen", wireType)
			}
			m.MaxClassNameLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxClassNameLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUriLen", wireType)
			}
			m.MinUriLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUriLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxUriLen", wireType)
			}
			m.MaxUriLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxUriLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSymbolLen", wireType)
			}
			m.MaxSymbolLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSymbolLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDescriptionLen", wireType)
			}
			m.MaxDescriptionLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDescriptionLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNftmint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftmint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNftmint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftmint
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNftmint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNftmint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftmint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftmint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftmint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftmint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftmint = fmt.Errorf("proto: unexpected end of group")
)
