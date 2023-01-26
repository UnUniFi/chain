// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ecosystem-incentive/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// At first, we go with this one type.
// NFTMARKET_FRONTEND type reward will be disributed for the creators of frontend of UnUniFi's services.
type RewardType int32

const (
	RewardType_NFTMARKET_FRONTEND RewardType = 0
)

var RewardType_name = map[int32]string{
	0: "NFTMARKET_FRONTEND",
}

var RewardType_value = map[string]int32{
	"NFTMARKET_FRONTEND": 0,
}

func (x RewardType) String() string {
	return proto.EnumName(RewardType_name, int32(x))
}

func (RewardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec4d1df7c0aa1f33, []int{0}
}

// Params defines the parameters for the module.
type Params struct {
	RewardParams            []*RewardParams `protobuf:"bytes,1,rep,name=reward_params,json=rewardParams,proto3" json:"reward_params,omitempty" yaml:"reward_params"`
	MaxIncentiveUnitIdLen   uint64          `protobuf:"varint,2,opt,name=max_incentive_unit_id_len,json=maxIncentiveUnitIdLen,proto3" json:"max_incentive_unit_id_len,omitempty" yaml:"max_incentive_unit_id_len"`
	MaxSubjectInfoNumInUnit uint64          `protobuf:"varint,3,opt,name=max_subject_info_num_in_unit,json=maxSubjectInfoNumInUnit,proto3" json:"max_subject_info_num_in_unit,omitempty" yaml:"max_subject_info_num_in_unit"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec4d1df7c0aa1f33, []int{0}
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

func (m *Params) GetRewardParams() []*RewardParams {
	if m != nil {
		return m.RewardParams
	}
	return nil
}

func (m *Params) GetMaxIncentiveUnitIdLen() uint64 {
	if m != nil {
		return m.MaxIncentiveUnitIdLen
	}
	return 0
}

func (m *Params) GetMaxSubjectInfoNumInUnit() uint64 {
	if m != nil {
		return m.MaxSubjectInfoNumInUnit
	}
	return 0
}

// RewardParams defines which module have which rewards type and rate
// to maintain the correctness of the fee rate in a module
// e.g. if nftmarket module have "Frontend" and "Collection" incentive,
// the combined those rates for the incentive cannot be exceed 1
type RewardParams struct {
	ModuleName string       `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty" yaml:"module_name"`
	RewardRate []RewardRate `protobuf:"bytes,2,rep,name=reward_rate,json=rewardRate,proto3" json:"reward_rate" yaml:"reward_rate"`
}

func (m *RewardParams) Reset()         { *m = RewardParams{} }
func (m *RewardParams) String() string { return proto.CompactTextString(m) }
func (*RewardParams) ProtoMessage()    {}
func (*RewardParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec4d1df7c0aa1f33, []int{1}
}
func (m *RewardParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardParams.Merge(m, src)
}
func (m *RewardParams) XXX_Size() int {
	return m.Size()
}
func (m *RewardParams) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardParams.DiscardUnknown(m)
}

var xxx_messageInfo_RewardParams proto.InternalMessageInfo

func (m *RewardParams) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *RewardParams) GetRewardRate() []RewardRate {
	if m != nil {
		return m.RewardRate
	}
	return nil
}

// RewardRate defines the ratio to take reward for a specific reward_type.
// The total sum of reward_rate in a module cannot be exceed 1
type RewardRate struct {
	RewardType RewardType                             `protobuf:"varint,1,opt,name=reward_type,json=rewardType,proto3,enum=ununifi.ecosystemincentive.RewardType" json:"reward_type,omitempty" yaml:"reward_type"`
	Rate       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate" yaml:"rate"`
}

func (m *RewardRate) Reset()         { *m = RewardRate{} }
func (m *RewardRate) String() string { return proto.CompactTextString(m) }
func (*RewardRate) ProtoMessage()    {}
func (*RewardRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec4d1df7c0aa1f33, []int{2}
}
func (m *RewardRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardRate.Merge(m, src)
}
func (m *RewardRate) XXX_Size() int {
	return m.Size()
}
func (m *RewardRate) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardRate.DiscardUnknown(m)
}

var xxx_messageInfo_RewardRate proto.InternalMessageInfo

func (m *RewardRate) GetRewardType() RewardType {
	if m != nil {
		return m.RewardType
	}
	return RewardType_NFTMARKET_FRONTEND
}

func init() {
	proto.RegisterEnum("ununifi.ecosystemincentive.RewardType", RewardType_name, RewardType_value)
	proto.RegisterType((*Params)(nil), "ununifi.ecosystemincentive.Params")
	proto.RegisterType((*RewardParams)(nil), "ununifi.ecosystemincentive.RewardParams")
	proto.RegisterType((*RewardRate)(nil), "ununifi.ecosystemincentive.RewardRate")
}

func init() { proto.RegisterFile("ecosystem-incentive/params.proto", fileDescriptor_ec4d1df7c0aa1f33) }

var fileDescriptor_ec4d1df7c0aa1f33 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x9b, 0xee, 0xb2, 0xe0, 0x74, 0x15, 0x19, 0xd6, 0x35, 0x16, 0x49, 0x42, 0x5c, 0xd6,
	0x22, 0x6c, 0x02, 0x2a, 0x08, 0x82, 0x07, 0xc3, 0x6e, 0xa1, 0xa8, 0x51, 0x63, 0xf7, 0xe2, 0xc1,
	0x61, 0x9a, 0x4c, 0xbb, 0xa3, 0x9d, 0x99, 0x92, 0x4c, 0xb4, 0x7d, 0x0b, 0x1f, 0xc4, 0x17, 0xf0,
	0xe6, 0x71, 0x8f, 0x7b, 0x14, 0x0f, 0x41, 0xda, 0x37, 0xc8, 0x13, 0x48, 0x26, 0xd9, 0x34, 0x82,
	0xc5, 0x3d, 0x25, 0x61, 0xfe, 0xdf, 0xef, 0xf7, 0x7d, 0x5f, 0x18, 0x60, 0x91, 0x50, 0x24, 0x8b,
	0x44, 0x12, 0x76, 0x44, 0x79, 0x48, 0xb8, 0xa4, 0x9f, 0x89, 0x3b, 0xc3, 0x31, 0x66, 0x89, 0x33,
	0x8b, 0x85, 0x14, 0xb0, 0x9b, 0xf2, 0x94, 0xd3, 0x31, 0x75, 0xea, 0x64, 0x1d, 0xec, 0xee, 0x4d,
	0xc4, 0x44, 0xa8, 0x98, 0x5b, 0xbc, 0x95, 0x15, 0xf6, 0xf7, 0x36, 0xd8, 0x79, 0xa3, 0x10, 0x70,
	0x02, 0xae, 0xc7, 0xe4, 0x0b, 0x8e, 0x23, 0x54, 0x32, 0x75, 0xcd, 0xda, 0xea, 0x75, 0x1e, 0xf6,
	0x9c, 0xcd, 0x50, 0x27, 0x50, 0x05, 0x25, 0xc0, 0xd3, 0xf3, 0xcc, 0xdc, 0x5b, 0x60, 0x36, 0x7d,
	0x6a, 0xff, 0x05, 0xb2, 0x83, 0xdd, 0xb8, 0x91, 0x83, 0x1f, 0xc0, 0x1d, 0x86, 0xe7, 0xa8, 0xa6,
	0xa0, 0x94, 0x53, 0x89, 0x68, 0x84, 0xa6, 0x84, 0xeb, 0x6d, 0x4b, 0xeb, 0x6d, 0x7b, 0x07, 0x79,
	0x66, 0x5a, 0x25, 0x6a, 0x63, 0xd4, 0x0e, 0x6e, 0x31, 0x3c, 0x1f, 0x5c, 0x1e, 0x9d, 0x72, 0x2a,
	0x07, 0xd1, 0x4b, 0xc2, 0xe1, 0x18, 0xdc, 0x2d, 0x8a, 0x92, 0x74, 0xf4, 0x91, 0x84, 0x12, 0x51,
	0x3e, 0x16, 0x88, 0xa7, 0x0c, 0x51, 0xae, 0xca, 0xf5, 0x2d, 0xa5, 0xb8, 0x9f, 0x67, 0xe6, 0xbd,
	0xb5, 0x62, 0x53, 0xda, 0x0e, 0x6e, 0x33, 0x3c, 0x7f, 0x57, 0x9e, 0x0e, 0xf8, 0x58, 0xf8, 0x29,
	0x1b, 0xf0, 0x42, 0x66, 0x7f, 0xd3, 0xc0, 0x6e, 0x73, 0x01, 0xf0, 0x09, 0xe8, 0x30, 0x11, 0xa5,
	0x53, 0x82, 0x38, 0x66, 0x44, 0xd7, 0x2c, 0xad, 0x77, 0xcd, 0xdb, 0xcf, 0x33, 0x13, 0x56, 0x9e,
	0xf5, 0xa1, 0x1d, 0x80, 0xf2, 0xcb, 0xc7, 0x8c, 0xc0, 0x10, 0x74, 0xaa, 0x8d, 0xc5, 0x58, 0x12,
	0xbd, 0xad, 0x16, 0x7f, 0xf8, 0xff, 0xc5, 0x07, 0x58, 0x12, 0xaf, 0x7b, 0x9e, 0x99, 0xad, 0xb5,
	0xa4, 0x01, 0xb2, 0x03, 0x10, 0xd7, 0x39, 0xfb, 0x87, 0x06, 0xc0, 0xba, 0x0c, 0xa2, 0xda, 0x29,
	0x17, 0xb3, 0xb2, 0xd9, 0x1b, 0x57, 0x71, 0x0e, 0x17, 0x33, 0xd2, 0x1c, 0xaa, 0x01, 0xa9, 0x7d,
	0x45, 0x06, 0xbe, 0x05, 0xdb, 0xd5, 0x34, 0xc5, 0x1a, 0x9e, 0x15, 0x5d, 0xfe, 0xca, 0xcc, 0xc3,
	0x09, 0x95, 0x67, 0xe9, 0xc8, 0x09, 0x05, 0x73, 0x43, 0x91, 0x30, 0x91, 0x54, 0x8f, 0xa3, 0x24,
	0xfa, 0xe4, 0x16, 0x94, 0xc4, 0x39, 0x26, 0x61, 0x9e, 0x99, 0x9d, 0x8a, 0xaf, 0x06, 0x51, 0xa8,
	0x07, 0x07, 0x97, 0x13, 0x28, 0xc1, 0x3e, 0x80, 0x7e, 0x7f, 0xf8, 0xea, 0x79, 0xf0, 0xe2, 0x64,
	0x88, 0xfa, 0xc1, 0x6b, 0x7f, 0x78, 0xe2, 0x1f, 0xdf, 0x6c, 0x79, 0xfe, 0xf9, 0xd2, 0xd0, 0x2e,
	0x96, 0x86, 0xf6, 0x7b, 0x69, 0x68, 0x5f, 0x57, 0x46, 0xeb, 0x62, 0x65, 0xb4, 0x7e, 0xae, 0x8c,
	0xd6, 0xfb, 0xc7, 0x0d, 0xf9, 0x69, 0xf1, 0x13, 0xfb, 0xd4, 0x0d, 0xcf, 0x30, 0xe5, 0xee, 0xdc,
	0xfd, 0xd7, 0xe5, 0x52, 0xed, 0x8c, 0x76, 0xd4, 0x55, 0x79, 0xf4, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0x63, 0x4d, 0x74, 0x80, 0x03, 0x00, 0x00,
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
	if m.MaxSubjectInfoNumInUnit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSubjectInfoNumInUnit))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxIncentiveUnitIdLen != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxIncentiveUnitIdLen))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RewardParams) > 0 {
		for iNdEx := len(m.RewardParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RewardParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardRate) > 0 {
		for iNdEx := len(m.RewardRate) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardRate[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RewardRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.RewardType != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RewardType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RewardParams) > 0 {
		for _, e := range m.RewardParams {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MaxIncentiveUnitIdLen != 0 {
		n += 1 + sovParams(uint64(m.MaxIncentiveUnitIdLen))
	}
	if m.MaxSubjectInfoNumInUnit != 0 {
		n += 1 + sovParams(uint64(m.MaxSubjectInfoNumInUnit))
	}
	return n
}

func (m *RewardParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.RewardRate) > 0 {
		for _, e := range m.RewardRate {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *RewardRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RewardType != 0 {
		n += 1 + sovParams(uint64(m.RewardType))
	}
	l = m.Rate.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardParams = append(m.RewardParams, &RewardParams{})
			if err := m.RewardParams[len(m.RewardParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxIncentiveUnitIdLen", wireType)
			}
			m.MaxIncentiveUnitIdLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxIncentiveUnitIdLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSubjectInfoNumInUnit", wireType)
			}
			m.MaxSubjectInfoNumInUnit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSubjectInfoNumInUnit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *RewardParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: RewardParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardRate = append(m.RewardRate, RewardRate{})
			if err := m.RewardRate[len(m.RewardRate)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *RewardRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: RewardRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardType", wireType)
			}
			m.RewardType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardType |= RewardType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)