// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incentive/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the incentive module's genesis state.
type GenesisState struct {
	Params               Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	CdpAccumulationTimes []GenesisAccumulationTime `protobuf:"bytes,2,rep,name=cdp_accumulation_times,json=cdpAccumulationTimes,proto3" json:"cdp_accumulation_times" yaml:"cdp_accumulation_times"`
	CdpMintingClaims     []CdpMintingClaim         `protobuf:"bytes,3,rep,name=cdp_minting_claims,json=cdpMintingClaims,proto3" json:"cdp_minting_claims" yaml:"cdp_minting_claims"`
	Denoms               *GenesisDenoms            `protobuf:"bytes,4,opt,name=denoms,proto3" json:"denoms,omitempty" yaml:"denoms"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ea08f29a85d2f6, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetCdpAccumulationTimes() []GenesisAccumulationTime {
	if m != nil {
		return m.CdpAccumulationTimes
	}
	return nil
}

func (m *GenesisState) GetCdpMintingClaims() []CdpMintingClaim {
	if m != nil {
		return m.CdpMintingClaims
	}
	return nil
}

func (m *GenesisState) GetDenoms() *GenesisDenoms {
	if m != nil {
		return m.Denoms
	}
	return nil
}

type GenesisAccumulationTime struct {
	CollateralType           string    `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty" yaml:"collateral_type"`
	PreviousAccumulationTime time.Time `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time" yaml:"previous_accumulation_time"`
}

func (m *GenesisAccumulationTime) Reset()         { *m = GenesisAccumulationTime{} }
func (m *GenesisAccumulationTime) String() string { return proto.CompactTextString(m) }
func (*GenesisAccumulationTime) ProtoMessage()    {}
func (*GenesisAccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ea08f29a85d2f6, []int{1}
}
func (m *GenesisAccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccumulationTime.Merge(m, src)
}
func (m *GenesisAccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccumulationTime proto.InternalMessageInfo

func (m *GenesisAccumulationTime) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisAccumulationTime) GetPreviousAccumulationTime() time.Time {
	if m != nil {
		return m.PreviousAccumulationTime
	}
	return time.Time{}
}

type GenesisDenoms struct {
	PrincipalDenom        string `protobuf:"bytes,1,opt,name=principal_denom,json=principalDenom,proto3" json:"principal_denom,omitempty" yaml:"principal_denom"`
	CdpMintingRewardDenom string `protobuf:"bytes,2,opt,name=cdp_minting_reward_denom,json=cdpMintingRewardDenom,proto3" json:"cdp_minting_reward_denom,omitempty" yaml:"principal_denom"`
}

func (m *GenesisDenoms) Reset()         { *m = GenesisDenoms{} }
func (m *GenesisDenoms) String() string { return proto.CompactTextString(m) }
func (*GenesisDenoms) ProtoMessage()    {}
func (*GenesisDenoms) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ea08f29a85d2f6, []int{2}
}
func (m *GenesisDenoms) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisDenoms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisDenoms.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisDenoms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDenoms.Merge(m, src)
}
func (m *GenesisDenoms) XXX_Size() int {
	return m.Size()
}
func (m *GenesisDenoms) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDenoms.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDenoms proto.InternalMessageInfo

func (m *GenesisDenoms) GetPrincipalDenom() string {
	if m != nil {
		return m.PrincipalDenom
	}
	return ""
}

func (m *GenesisDenoms) GetCdpMintingRewardDenom() string {
	if m != nil {
		return m.CdpMintingRewardDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ununifi.incentive.GenesisState")
	proto.RegisterType((*GenesisAccumulationTime)(nil), "ununifi.incentive.GenesisAccumulationTime")
	proto.RegisterType((*GenesisDenoms)(nil), "ununifi.incentive.GenesisDenoms")
}

func init() { proto.RegisterFile("incentive/genesis.proto", fileDescriptor_b5ea08f29a85d2f6) }

var fileDescriptor_b5ea08f29a85d2f6 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0xb5, 0xec, 0x62, 0xa8, 0xd2, 0xa4, 0x8d, 0xc8, 0x43, 0x36, 0x54, 0x76, 0x06, 0x0a, 0x21,
	0x50, 0x09, 0xd2, 0x5d, 0x77, 0x95, 0x43, 0x5b, 0x28, 0x85, 0xa2, 0x38, 0x9b, 0x6e, 0xcc, 0x78,
	0x34, 0x51, 0x06, 0x34, 0x0f, 0x34, 0xa3, 0xb4, 0xfe, 0x82, 0x6e, 0xf3, 0x11, 0xdd, 0xf4, 0x4f,
	0xb2, 0xcc, 0xb2, 0x2b, 0xb7, 0xd8, 0x9b, 0xae, 0xf3, 0x05, 0x45, 0xa3, 0xb1, 0x15, 0xbf, 0xc8,
	0x6e, 0x34, 0x73, 0xce, 0xb9, 0xe7, 0x9e, 0x7b, 0x65, 0x1f, 0x12, 0x86, 0x30, 0x53, 0xe4, 0x1a,
	0x07, 0x09, 0x66, 0x58, 0x12, 0xe9, 0x8b, 0x8c, 0x2b, 0xee, 0xec, 0xe6, 0x2c, 0x67, 0xe4, 0x92,
	0xf8, 0x73, 0x40, 0x7b, 0x2f, 0xe1, 0x09, 0xd7, 0xaf, 0x41, 0x71, 0x2a, 0x81, 0xed, 0x4e, 0xc2,
	0x79, 0x92, 0xe2, 0x40, 0x7f, 0x0d, 0xf3, 0xcb, 0x40, 0x11, 0x8a, 0xa5, 0x82, 0x54, 0x18, 0x40,
	0xab, 0x2a, 0x31, 0x3f, 0x95, 0x4f, 0xe0, 0x67, 0xc3, 0x7e, 0xf6, 0xa1, 0x2c, 0x7b, 0xae, 0xa0,
	0xc2, 0xce, 0x47, 0xbb, 0x29, 0x60, 0x06, 0xa9, 0x74, 0xad, 0xae, 0x75, 0xbc, 0x75, 0xda, 0xf2,
	0x57, 0x6c, 0xf8, 0x5f, 0x34, 0x20, 0xdc, 0xbf, 0x1d, 0x77, 0x6a, 0xf7, 0xe3, 0xce, 0xf6, 0x08,
	0xd2, 0xf4, 0x2d, 0x28, 0x69, 0x20, 0x32, 0x7c, 0xe7, 0x87, 0x65, 0x1f, 0xa0, 0x58, 0x0c, 0x20,
	0x42, 0x39, 0xcd, 0x53, 0xa8, 0x08, 0x67, 0x03, 0x6d, 0xcd, 0xad, 0x77, 0x1b, 0xc7, 0x5b, 0xa7,
	0x27, 0x6b, 0xa4, 0x8d, 0x97, 0x77, 0x0f, 0x38, 0x7d, 0x42, 0x71, 0xf8, 0xca, 0xd4, 0x7a, 0x59,
	0xd6, 0x5a, 0xaf, 0x0b, 0xa2, 0x3d, 0x14, 0x8b, 0x65, 0xae, 0x74, 0xa4, 0xed, 0x14, 0x04, 0x4a,
	0x98, 0x22, 0x2c, 0x19, 0xa0, 0x14, 0x12, 0x2a, 0xdd, 0x86, 0x36, 0x01, 0xd6, 0x98, 0xe8, 0xc5,
	0xe2, 0x73, 0x89, 0xed, 0x15, 0xd0, 0xf0, 0xc8, 0x14, 0x6f, 0x55, 0xc5, 0x17, 0xb5, 0x40, 0xf4,
	0x02, 0x2d, 0x72, 0xa4, 0xf3, 0xc9, 0x6e, 0xc6, 0x98, 0x71, 0x2a, 0xdd, 0x27, 0x3a, 0xc8, 0xee,
	0xe6, 0x6e, 0xcf, 0x34, 0x2e, 0xdc, 0xad, 0xb2, 0x2c, 0x99, 0x20, 0x32, 0x12, 0xe0, 0x9f, 0x65,
	0x1f, 0x6e, 0x88, 0xc6, 0xe9, 0xd9, 0xcf, 0x11, 0x4f, 0x53, 0xa8, 0x70, 0x06, 0xd3, 0x81, 0x1a,
	0x09, 0xac, 0x47, 0xf7, 0x34, 0x6c, 0xdf, 0x8f, 0x3b, 0x07, 0xc6, 0xf2, 0x22, 0x00, 0x44, 0x3b,
	0xd5, 0x4d, 0x7f, 0x24, 0x70, 0x31, 0xac, 0xb6, 0xc8, 0xf0, 0x35, 0xe1, 0xb9, 0x5c, 0x4d, 0xd6,
	0xad, 0xeb, 0x16, 0xda, 0x7e, 0xb9, 0x69, 0xfe, 0x6c, 0xd3, 0xfc, 0xfe, 0x6c, 0xd3, 0xc2, 0xd7,
	0x26, 0xa3, 0x23, 0xb3, 0x0c, 0x1b, 0xb5, 0xc0, 0xcd, 0x9f, 0x8e, 0x15, 0xb9, 0x33, 0xc0, 0x72,
	0x3b, 0xe0, 0x97, 0x65, 0x6f, 0x2f, 0xe4, 0x52, 0x34, 0x28, 0x32, 0xc2, 0x10, 0x11, 0x30, 0x1d,
	0xe8, 0x40, 0x56, 0x1b, 0x5c, 0x02, 0x80, 0x68, 0x67, 0x7e, 0xa3, 0x55, 0x9c, 0x73, 0xdb, 0x7d,
	0x38, 0xb7, 0x0c, 0x7f, 0x83, 0x59, 0x6c, 0xd4, 0xea, 0x8f, 0xaa, 0xed, 0x57, 0xe3, 0x8d, 0x34,
	0x53, 0x8b, 0x86, 0x67, 0xb7, 0x13, 0xcf, 0xba, 0x9b, 0x78, 0xd6, 0xdf, 0x89, 0x67, 0xdd, 0x4c,
	0xbd, 0xda, 0xdd, 0xd4, 0xab, 0xfd, 0x9e, 0x7a, 0xb5, 0xaf, 0x27, 0x09, 0x51, 0x57, 0xf9, 0xd0,
	0x47, 0x9c, 0x06, 0x17, 0xec, 0x82, 0x91, 0xf7, 0x24, 0x40, 0x57, 0x90, 0xb0, 0xe0, 0x7b, 0xf5,
	0x0f, 0x06, 0xc5, 0x2c, 0xe4, 0xb0, 0xa9, 0xe3, 0x7c, 0xf3, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x19,
	0x90, 0xfd, 0x63, 0x0a, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Denoms != nil {
		{
			size, err := m.Denoms.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.CdpMintingClaims) > 0 {
		for iNdEx := len(m.CdpMintingClaims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CdpMintingClaims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.CdpAccumulationTimes) > 0 {
		for iNdEx := len(m.CdpAccumulationTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CdpAccumulationTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisAccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisDenoms) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisDenoms) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisDenoms) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CdpMintingRewardDenom) > 0 {
		i -= len(m.CdpMintingRewardDenom)
		copy(dAtA[i:], m.CdpMintingRewardDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CdpMintingRewardDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PrincipalDenom) > 0 {
		i -= len(m.PrincipalDenom)
		copy(dAtA[i:], m.PrincipalDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PrincipalDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.CdpAccumulationTimes) > 0 {
		for _, e := range m.CdpAccumulationTimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CdpMintingClaims) > 0 {
		for _, e := range m.CdpMintingClaims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Denoms != nil {
		l = m.Denoms.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisAccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisDenoms) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PrincipalDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.CdpMintingRewardDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdpAccumulationTimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CdpAccumulationTimes = append(m.CdpAccumulationTimes, GenesisAccumulationTime{})
			if err := m.CdpAccumulationTimes[len(m.CdpAccumulationTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdpMintingClaims", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CdpMintingClaims = append(m.CdpMintingClaims, CdpMintingClaim{})
			if err := m.CdpMintingClaims[len(m.CdpMintingClaims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Denoms == nil {
				m.Denoms = &GenesisDenoms{}
			}
			if err := m.Denoms.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisAccumulationTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisAccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisDenoms) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisDenoms: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisDenoms: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrincipalDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrincipalDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdpMintingRewardDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CdpMintingRewardDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)