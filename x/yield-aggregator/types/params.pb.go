// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: yield-aggregator/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	CommissionRate       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=commission_rate,json=commissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission_rate"`
	VaultCreationFee     types.Coin                             `protobuf:"bytes,2,opt,name=vault_creation_fee,json=vaultCreationFee,proto3" json:"vault_creation_fee"`
	VaultCreationDeposit types.Coin                             `protobuf:"bytes,3,opt,name=vault_creation_deposit,json=vaultCreationDeposit,proto3" json:"vault_creation_deposit"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d4d0a492ea9705, []int{0}
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

func (m *Params) GetVaultCreationFee() types.Coin {
	if m != nil {
		return m.VaultCreationFee
	}
	return types.Coin{}
}

func (m *Params) GetVaultCreationDeposit() types.Coin {
	if m != nil {
		return m.VaultCreationDeposit
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "ununifi.yieldaggregator.Params")
}

func init() { proto.RegisterFile("yield-aggregator/params.proto", fileDescriptor_b5d4d0a492ea9705) }

var fileDescriptor_b5d4d0a492ea9705 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4a, 0x02, 0x41,
	0x1c, 0xdf, 0x35, 0x11, 0xda, 0xa0, 0x62, 0x91, 0x52, 0xa1, 0x51, 0x3a, 0x84, 0x17, 0x67, 0xb0,
	0x6e, 0xd1, 0x49, 0xc5, 0x4b, 0x04, 0x21, 0x78, 0xe9, 0x22, 0xb3, 0xe3, 0xdf, 0x75, 0xc8, 0x9d,
	0x59, 0x76, 0x66, 0x25, 0xdf, 0xa2, 0x63, 0xd0, 0xa5, 0x87, 0xe8, 0x21, 0x3c, 0x4a, 0xa7, 0xe8,
	0x20, 0xe1, 0xbe, 0x48, 0xec, 0xee, 0x80, 0x7d, 0x5c, 0x3a, 0xed, 0xfe, 0xf9, 0x7d, 0xf2, 0x1b,
	0xe7, 0x64, 0xc1, 0x61, 0x36, 0x6e, 0x51, 0xdf, 0x8f, 0xc0, 0xa7, 0x5a, 0x46, 0x24, 0xa4, 0x11,
	0x0d, 0x14, 0x0e, 0x23, 0xa9, 0xa5, 0x7b, 0x1c, 0x8b, 0x58, 0xf0, 0x09, 0xc7, 0x19, 0x6d, 0xcb,
	0xaa, 0x95, 0x7d, 0xe9, 0xcb, 0x8c, 0x43, 0xd2, 0xbf, 0x9c, 0x5e, 0xab, 0x32, 0xa9, 0x02, 0xa9,
	0x46, 0x39, 0x90, 0x1f, 0x06, 0x42, 0xf9, 0x45, 0x3c, 0xaa, 0x80, 0xcc, 0xdb, 0x1e, 0x68, 0xda,
	0x26, 0x4c, 0x72, 0x91, 0xe3, 0xa7, 0xcf, 0x05, 0xa7, 0x74, 0x9b, 0x45, 0xbb, 0xe0, 0x1c, 0x30,
	0x19, 0x04, 0x5c, 0x29, 0x2e, 0xc5, 0x28, 0xa2, 0x1a, 0x2a, 0x76, 0xc3, 0x6e, 0xee, 0x76, 0xae,
	0x96, 0xeb, 0xba, 0xf5, 0xb1, 0xae, 0x9f, 0xf9, 0x5c, 0x4f, 0x63, 0x0f, 0x33, 0x19, 0x98, 0x10,
	0xf3, 0x69, 0xa9, 0xf1, 0x3d, 0xd1, 0x8b, 0x10, 0x14, 0xee, 0x01, 0x7b, 0x7b, 0x6d, 0x39, 0xa6,
	0x43, 0x0f, 0xd8, 0x60, 0x7f, 0x6b, 0x3a, 0xa0, 0x1a, 0xdc, 0x1b, 0xc7, 0x9d, 0xd3, 0x78, 0xa6,
	0x47, 0x2c, 0x02, 0xaa, 0xd3, 0xa8, 0x09, 0x40, 0xa5, 0xd0, 0xb0, 0x9b, 0x7b, 0xe7, 0x55, 0x6c,
	0x84, 0x69, 0x5d, 0x6c, 0xea, 0xe2, 0xae, 0xe4, 0xa2, 0x53, 0x4c, 0x4b, 0x0c, 0x0e, 0x33, 0x69,
	0xd7, 0x28, 0xfb, 0x00, 0xee, 0xd0, 0x39, 0xfa, 0x65, 0x37, 0x86, 0x50, 0x2a, 0xae, 0x2b, 0x3b,
	0xff, 0xb3, 0x2c, 0xff, 0xb0, 0xec, 0xe5, 0xe2, 0xcb, 0xe2, 0xd3, 0x4b, 0xdd, 0xea, 0x5c, 0x2f,
	0x37, 0xc8, 0x5e, 0x6d, 0x90, 0xfd, 0xb9, 0x41, 0xf6, 0x63, 0x82, 0xac, 0x55, 0x82, 0xac, 0xf7,
	0x04, 0x59, 0x77, 0xed, 0x6f, 0x5b, 0x0c, 0xc5, 0x50, 0xf0, 0x3e, 0x27, 0x6c, 0x4a, 0xb9, 0x20,
	0x0f, 0xe4, 0xcf, 0xdb, 0x66, 0xd3, 0x78, 0xa5, 0x6c, 0xf1, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4d, 0x4d, 0x5b, 0xe3, 0xfc, 0x01, 0x00, 0x00,
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
	{
		size, err := m.VaultCreationDeposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.VaultCreationFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.CommissionRate.Size()
		i -= size
		if _, err := m.CommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.CommissionRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.VaultCreationFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.VaultCreationDeposit.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
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
			if err := m.CommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultCreationFee", wireType)
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
			if err := m.VaultCreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultCreationDeposit", wireType)
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
			if err := m.VaultCreationDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
