// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Validator_ValidatorStatus int32

const (
	Validator_Active   Validator_ValidatorStatus = 0
	Validator_Inactive Validator_ValidatorStatus = 1
)

var Validator_ValidatorStatus_name = map[int32]string{
	0: "Active",
	1: "Inactive",
}

var Validator_ValidatorStatus_value = map[string]int32{
	"Active":   0,
	"Inactive": 1,
}

func (x Validator_ValidatorStatus) String() string {
	return proto.EnumName(Validator_ValidatorStatus_name, int32(x))
}

func (Validator_ValidatorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_135ed83653830bac, []int{1, 0}
}

type ValidatorExchangeRate struct {
	InternalTokensToSharesRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=internalTokensToSharesRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"internalTokensToSharesRate"`
	EpochNumber                uint64                                 `protobuf:"varint,2,opt,name=epochNumber,proto3" json:"epochNumber,omitempty"`
}

func (m *ValidatorExchangeRate) Reset()         { *m = ValidatorExchangeRate{} }
func (m *ValidatorExchangeRate) String() string { return proto.CompactTextString(m) }
func (*ValidatorExchangeRate) ProtoMessage()    {}
func (*ValidatorExchangeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_135ed83653830bac, []int{0}
}
func (m *ValidatorExchangeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorExchangeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorExchangeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorExchangeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorExchangeRate.Merge(m, src)
}
func (m *ValidatorExchangeRate) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorExchangeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorExchangeRate.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorExchangeRate proto.InternalMessageInfo

func (m *ValidatorExchangeRate) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

type Validator struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string                    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Status               Validator_ValidatorStatus `protobuf:"varint,3,opt,name=status,proto3,enum=Stridelabs.stride.stakeibc.Validator_ValidatorStatus" json:"status,omitempty"`
	CommissionRate       uint64                    `protobuf:"varint,4,opt,name=commissionRate,proto3" json:"commissionRate,omitempty"`
	DelegationAmt        uint64                    `protobuf:"varint,5,opt,name=delegationAmt,proto3" json:"delegationAmt,omitempty"`
	Weight               uint64                    `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	InternalExchangeRate *ValidatorExchangeRate    `protobuf:"bytes,7,opt,name=internalExchangeRate,proto3" json:"internalExchangeRate,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_135ed83653830bac, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetStatus() Validator_ValidatorStatus {
	if m != nil {
		return m.Status
	}
	return Validator_Active
}

func (m *Validator) GetCommissionRate() uint64 {
	if m != nil {
		return m.CommissionRate
	}
	return 0
}

func (m *Validator) GetDelegationAmt() uint64 {
	if m != nil {
		return m.DelegationAmt
	}
	return 0
}

func (m *Validator) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Validator) GetInternalExchangeRate() *ValidatorExchangeRate {
	if m != nil {
		return m.InternalExchangeRate
	}
	return nil
}

func init() {
	proto.RegisterEnum("Stridelabs.stride.stakeibc.Validator_ValidatorStatus", Validator_ValidatorStatus_name, Validator_ValidatorStatus_value)
	proto.RegisterType((*ValidatorExchangeRate)(nil), "Stridelabs.stride.stakeibc.ValidatorExchangeRate")
	proto.RegisterType((*Validator)(nil), "Stridelabs.stride.stakeibc.Validator")
}

func init() {
	// proto.RegisterFile("stakeibc/validator.proto", fileDescriptor_135ed83653830bac)
}

var fileDescriptor_135ed83653830bac = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xdf, 0xb1, 0x71, 0x6b, 0x5e, 0xb5, 0xca, 0x10, 0x65, 0xcd, 0x61, 0xbb, 0x04, 0x29, 0x01,
	0xc9, 0x2e, 0x46, 0xbc, 0x79, 0x49, 0xa8, 0xa0, 0xa0, 0x1e, 0x36, 0xc5, 0x83, 0x17, 0x99, 0x9d,
	0x7d, 0xec, 0x0e, 0xc9, 0xce, 0x84, 0x9d, 0x49, 0xad, 0xe0, 0x87, 0xf0, 0x03, 0xf8, 0x31, 0x0a,
	0x7e, 0x85, 0x1e, 0x4b, 0x4f, 0xe2, 0xa1, 0x48, 0xf2, 0x45, 0x24, 0xb3, 0x9b, 0x18, 0x83, 0x4a,
	0x4f, 0xf3, 0xde, 0xef, 0xfd, 0xfb, 0xbd, 0x37, 0x3f, 0xf0, 0xb4, 0x61, 0x63, 0x14, 0x09, 0x8f,
	0x4e, 0xd8, 0x44, 0xa4, 0xcc, 0xa8, 0x32, 0x9c, 0x96, 0xca, 0x28, 0xda, 0x1e, 0x99, 0x52, 0xa4,
	0x38, 0x61, 0x89, 0x0e, 0xb5, 0x35, 0xc3, 0x55, 0x6e, 0xfb, 0x21, 0x57, 0xba, 0x50, 0xfa, 0x83,
	0xcd, 0x8c, 0x2a, 0xa7, 0x2a, 0x6b, 0xb7, 0x32, 0x95, 0xa9, 0x0a, 0x5f, 0x5a, 0x15, 0xda, 0xf9,
	0x46, 0xe0, 0xfe, 0xbb, 0xd5, 0x80, 0x17, 0xa7, 0x3c, 0x67, 0x32, 0xc3, 0x98, 0x19, 0xa4, 0x9f,
	0xa1, 0x2d, 0xa4, 0xc1, 0x52, 0xb2, 0xc9, 0xb1, 0x1a, 0xa3, 0xd4, 0xc7, 0x6a, 0x94, 0xb3, 0x12,
	0xf5, 0x32, 0xea, 0x91, 0x80, 0x74, 0x9b, 0xc3, 0xe7, 0xe7, 0x57, 0x07, 0xce, 0x8f, 0xab, 0x83,
	0xc3, 0x4c, 0x98, 0x7c, 0x96, 0x84, 0x5c, 0x15, 0xf5, 0xd0, 0xfa, 0xe9, 0xe9, 0x74, 0x1c, 0x99,
	0x4f, 0x53, 0xd4, 0xe1, 0x11, 0xf2, 0xcb, 0xb3, 0x1e, 0xd4, 0x9c, 0x8e, 0x90, 0xc7, 0xff, 0xe9,
	0x4f, 0x03, 0xd8, 0xc3, 0xa9, 0xe2, 0xf9, 0xdb, 0x59, 0x91, 0x60, 0xe9, 0xdd, 0x08, 0x48, 0xb7,
	0x11, 0x6f, 0x42, 0x9d, 0xaf, 0x3b, 0xd0, 0x5c, 0x33, 0xa7, 0x14, 0x1a, 0x92, 0x15, 0x35, 0xaf,
	0xd8, 0xda, 0xb4, 0x0f, 0xbb, 0x2c, 0x4d, 0x4b, 0xd4, 0xda, 0xd6, 0x37, 0x87, 0xde, 0xe5, 0x59,
	0xaf, 0x55, 0x13, 0x18, 0x54, 0x91, 0xe5, 0x2d, 0x65, 0x16, 0xaf, 0x12, 0xe9, 0x1b, 0x70, 0xb5,
	0x61, 0x66, 0xa6, 0xbd, 0x9d, 0x80, 0x74, 0xf7, 0xfb, 0xcf, 0xc2, 0x7f, 0x5f, 0x3b, 0x5c, 0x8f,
	0xff, 0x6d, 0x8d, 0x6c, 0x71, 0x5c, 0x37, 0xa1, 0x87, 0xb0, 0xcf, 0x55, 0x51, 0x08, 0xad, 0x85,
	0x92, 0xf6, 0x70, 0x0d, 0xbb, 0xc9, 0x16, 0x4a, 0x1f, 0xc1, 0x9d, 0x14, 0x27, 0x98, 0x31, 0x23,
	0x94, 0x1c, 0x14, 0xc6, 0xbb, 0x69, 0xd3, 0xfe, 0x04, 0xe9, 0x03, 0x70, 0x3f, 0xa2, 0xc8, 0x72,
	0xe3, 0xb9, 0x36, 0x5c, 0x7b, 0x14, 0xa1, 0xb5, 0x3a, 0xe5, 0xe6, 0x17, 0x7a, 0xbb, 0x01, 0xe9,
	0xee, 0xf5, 0x9f, 0x5c, 0x6b, 0x85, 0xcd, 0xc2, 0xf8, 0xaf, 0xed, 0x3a, 0x8f, 0xe1, 0xee, 0xd6,
	0x9e, 0x14, 0xc0, 0x1d, 0x70, 0x23, 0x4e, 0xf0, 0x9e, 0x43, 0x6f, 0xc3, 0xad, 0x57, 0x92, 0x55,
	0x1e, 0x19, 0xbe, 0x3c, 0x9f, 0xfb, 0xe4, 0x62, 0xee, 0x93, 0x9f, 0x73, 0x9f, 0x7c, 0x59, 0xf8,
	0xce, 0xc5, 0xc2, 0x77, 0xbe, 0x2f, 0x7c, 0xe7, 0x7d, 0xb8, 0x21, 0x96, 0x8a, 0x59, 0xef, 0x35,
	0x4b, 0x74, 0x54, 0x51, 0x8b, 0x4e, 0xa3, 0xb5, 0xf2, 0xad, 0x70, 0x12, 0xd7, 0x2a, 0xf5, 0xe9,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x3e, 0x54, 0xdb, 0x12, 0x03, 0x00, 0x00,
}

func (m *ValidatorExchangeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorExchangeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorExchangeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.InternalTokensToSharesRate.Size()
		i -= size
		if _, err := m.InternalTokensToSharesRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InternalExchangeRate != nil {
		{
			size, err := m.InternalExchangeRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Weight != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x30
	}
	if m.DelegationAmt != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.DelegationAmt))
		i--
		dAtA[i] = 0x28
	}
	if m.CommissionRate != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.CommissionRate))
		i--
		dAtA[i] = 0x20
	}
	if m.Status != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorExchangeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InternalTokensToSharesRate.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.EpochNumber != 0 {
		n += 1 + sovValidator(uint64(m.EpochNumber))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovValidator(uint64(m.Status))
	}
	if m.CommissionRate != 0 {
		n += 1 + sovValidator(uint64(m.CommissionRate))
	}
	if m.DelegationAmt != 0 {
		n += 1 + sovValidator(uint64(m.DelegationAmt))
	}
	if m.Weight != 0 {
		n += 1 + sovValidator(uint64(m.Weight))
	}
	if m.InternalExchangeRate != nil {
		l = m.InternalExchangeRate.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorExchangeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorExchangeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorExchangeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalTokensToSharesRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InternalTokensToSharesRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Validator_ValidatorStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			m.CommissionRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissionRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationAmt", wireType)
			}
			m.DelegationAmt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegationAmt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalExchangeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InternalExchangeRate == nil {
				m.InternalExchangeRate = &ValidatorExchangeRate{}
			}
			if err := m.InternalExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
