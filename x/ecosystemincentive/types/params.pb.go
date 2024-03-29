// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ununifi/ecosystemincentive/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// STAKERS type reward will be distributed for the stakers of GUU token.
// FRONTEND_DEVELOPERS type reward will be disributed for the creators of frontend of UnUniFi's services.
// COMMUNITY_POOL type reward will be distributed for the community pool.
type RewardType int32

const (
	RewardType_UNKNOWN             RewardType = 0
	RewardType_STAKERS             RewardType = 1
	RewardType_FRONTEND_DEVELOPERS RewardType = 2
	RewardType_COMMUNITY_POOL      RewardType = 3
)

var RewardType_name = map[int32]string{
	0: "UNKNOWN",
	1: "STAKERS",
	2: "FRONTEND_DEVELOPERS",
	3: "COMMUNITY_POOL",
}

var RewardType_value = map[string]int32{
	"UNKNOWN":             0,
	"STAKERS":             1,
	"FRONTEND_DEVELOPERS": 2,
	"COMMUNITY_POOL":      3,
}

func (x RewardType) String() string {
	return proto.EnumName(RewardType_name, int32(x))
}

func (RewardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4a7b5cb0e4986851, []int{0}
}

// Params defines the parameters for the module.
type Params struct {
	RewardParams []*RewardParams `protobuf:"bytes,1,rep,name=reward_params,json=rewardParams,proto3" json:"reward_params,omitempty" yaml:"reward_params"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7b5cb0e4986851, []int{0}
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

// RewardParams defines which module have which rewards type and rate
// to maintain the correctness of the fee rate in a module
// e.g. if nftbackedloan module have "Frontend" and "Collection" incentive,
// the combined those rates for the incentive cannot be exceed 1
type RewardParams struct {
	ModuleName string       `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty" yaml:"module_name"`
	RewardRate []RewardRate `protobuf:"bytes,2,rep,name=reward_rate,json=rewardRate,proto3" json:"reward_rate" yaml:"reward_rate"`
}

func (m *RewardParams) Reset()         { *m = RewardParams{} }
func (m *RewardParams) String() string { return proto.CompactTextString(m) }
func (*RewardParams) ProtoMessage()    {}
func (*RewardParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7b5cb0e4986851, []int{1}
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
	return fileDescriptor_4a7b5cb0e4986851, []int{2}
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
	return RewardType_UNKNOWN
}

func init() {
	proto.RegisterEnum("ununifi.ecosystemincentive.RewardType", RewardType_name, RewardType_value)
	proto.RegisterType((*Params)(nil), "ununifi.ecosystemincentive.Params")
	proto.RegisterType((*RewardParams)(nil), "ununifi.ecosystemincentive.RewardParams")
	proto.RegisterType((*RewardRate)(nil), "ununifi.ecosystemincentive.RewardRate")
}

func init() {
	proto.RegisterFile("ununifi/ecosystemincentive/params.proto", fileDescriptor_4a7b5cb0e4986851)
}

var fileDescriptor_4a7b5cb0e4986851 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x33, 0xad, 0x54, 0x9c, 0xd4, 0xb2, 0x8c, 0x45, 0x97, 0x3d, 0x24, 0x25, 0x87, 0xba,
	0x08, 0x26, 0x60, 0x0f, 0x82, 0xe0, 0xc1, 0x75, 0x53, 0x90, 0x76, 0x93, 0x35, 0x9b, 0x28, 0x7a,
	0x09, 0xd3, 0xec, 0x98, 0x0e, 0x76, 0x66, 0x62, 0xfe, 0xa8, 0xf9, 0x16, 0x7e, 0x10, 0x3f, 0x84,
	0xc7, 0x1e, 0x7b, 0x14, 0x0f, 0x41, 0x76, 0xbf, 0xc1, 0x7e, 0x02, 0xc9, 0x24, 0xee, 0x46, 0x54,
	0xf4, 0x94, 0x0c, 0xf3, 0xbc, 0xbf, 0xe7, 0x79, 0x86, 0x17, 0xde, 0x2d, 0x78, 0xc1, 0xe9, 0x1b,
	0x6a, 0x91, 0x48, 0x64, 0x65, 0x96, 0x13, 0x46, 0x79, 0x44, 0x78, 0x4e, 0xdf, 0x13, 0x2b, 0xc1,
	0x29, 0x66, 0x99, 0x99, 0xa4, 0x22, 0x17, 0x68, 0xd0, 0x0a, 0xcd, 0xdf, 0x85, 0x83, 0xfd, 0x58,
	0xc4, 0x42, 0xca, 0xac, 0xfa, 0xaf, 0x99, 0x30, 0xde, 0xc1, 0x9d, 0xa9, 0x24, 0xa0, 0x18, 0xde,
	0x4c, 0xc9, 0x07, 0x9c, 0xce, 0xc3, 0x06, 0xd9, 0x07, 0x07, 0xdb, 0x43, 0xf5, 0xc1, 0xd0, 0xfc,
	0x3b, 0xd3, 0xf4, 0xe4, 0x40, 0x03, 0x18, 0xf5, 0x57, 0x95, 0xbe, 0x5f, 0x62, 0x76, 0xf1, 0xc8,
	0xf8, 0x05, 0x64, 0x78, 0xbb, 0x69, 0x47, 0x67, 0x7c, 0x06, 0x70, 0xb7, 0x3b, 0x88, 0x1e, 0x42,
	0x95, 0x89, 0x79, 0x71, 0x41, 0x42, 0x8e, 0x19, 0xe9, 0x83, 0x03, 0x30, 0xbc, 0x31, 0xba, 0xbd,
	0xaa, 0x74, 0xd4, 0xd0, 0x3a, 0x97, 0x86, 0x07, 0x9b, 0x93, 0x83, 0x19, 0x41, 0x11, 0x54, 0x5b,
	0xa7, 0x14, 0xe7, 0xa4, 0xbf, 0x25, 0x03, 0x1f, 0xfe, 0x3b, 0xb0, 0x87, 0x73, 0x32, 0x1a, 0x5c,
	0x56, 0xba, 0xb2, 0x31, 0xe9, 0x80, 0x0c, 0x0f, 0xa6, 0x6b, 0x9d, 0xf1, 0x05, 0x40, 0xb8, 0x19,
	0x43, 0xe1, 0xda, 0x33, 0x2f, 0x93, 0x26, 0xec, 0xde, 0xff, 0x78, 0xfa, 0x65, 0x42, 0xba, 0xa5,
	0x3a, 0x90, 0xb5, 0x5f, 0xad, 0x41, 0xcf, 0xe1, 0xb5, 0xb6, 0x4d, 0xfd, 0x0c, 0x8f, 0xeb, 0x94,
	0xdf, 0x2a, 0xfd, 0x30, 0xa6, 0xf9, 0x79, 0x71, 0x66, 0x46, 0x82, 0x59, 0x91, 0xc8, 0x98, 0xc8,
	0xda, 0xcf, 0xfd, 0x6c, 0xfe, 0xd6, 0xaa, 0x29, 0x99, 0x39, 0x26, 0xd1, 0xaa, 0xd2, 0xd5, 0x96,
	0x2f, 0x8b, 0x48, 0xd4, 0xbd, 0xd9, 0xcf, 0x06, 0xd2, 0x40, 0x85, 0xd7, 0x03, 0xe7, 0xc4, 0x71,
	0x5f, 0x3a, 0x3d, 0xa5, 0x3e, 0xcc, 0xfc, 0x27, 0x27, 0xb6, 0x37, 0xeb, 0x01, 0x74, 0x07, 0xde,
	0x3a, 0xf6, 0x5c, 0xc7, 0xb7, 0x9d, 0x71, 0x38, 0xb6, 0x5f, 0xd8, 0xa7, 0xee, 0xb4, 0xbe, 0xd8,
	0x42, 0x08, 0xee, 0x3d, 0x75, 0x27, 0x93, 0xc0, 0x79, 0xe6, 0xbf, 0x0a, 0xa7, 0xae, 0x7b, 0xda,
	0xdb, 0x1e, 0x4d, 0x2e, 0x17, 0x1a, 0xb8, 0x5a, 0x68, 0xe0, 0xfb, 0x42, 0x03, 0x9f, 0x96, 0x9a,
	0x72, 0xb5, 0xd4, 0x94, 0xaf, 0x4b, 0x4d, 0x79, 0x7d, 0xd4, 0xc9, 0x1a, 0xf0, 0x80, 0xd3, 0x63,
	0x6a, 0x45, 0xe7, 0x98, 0x72, 0xeb, 0xe3, 0x9f, 0x36, 0x58, 0x86, 0x3f, 0xdb, 0x91, 0xfb, 0x78,
	0xf4, 0x23, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xe1, 0xd3, 0x6d, 0xec, 0x02, 0x00, 0x00,
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
