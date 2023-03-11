// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: copy-trading/exemplary_trader.proto

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

type ExemplaryTrader struct {
	Address              string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name                 string                                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ProfitCommissionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=profit_commission_rate,json=profitCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"profit_commission_rate"`
}

func (m *ExemplaryTrader) Reset()         { *m = ExemplaryTrader{} }
func (m *ExemplaryTrader) String() string { return proto.CompactTextString(m) }
func (*ExemplaryTrader) ProtoMessage()    {}
func (*ExemplaryTrader) Descriptor() ([]byte, []int) {
	return fileDescriptor_973e30b53d12a7a1, []int{0}
}
func (m *ExemplaryTrader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExemplaryTrader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExemplaryTrader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExemplaryTrader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExemplaryTrader.Merge(m, src)
}
func (m *ExemplaryTrader) XXX_Size() int {
	return m.Size()
}
func (m *ExemplaryTrader) XXX_DiscardUnknown() {
	xxx_messageInfo_ExemplaryTrader.DiscardUnknown(m)
}

var xxx_messageInfo_ExemplaryTrader proto.InternalMessageInfo

func (m *ExemplaryTrader) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ExemplaryTrader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExemplaryTrader) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ExemplaryTrader)(nil), "ununifi.chain.copytrading.ExemplaryTrader")
}

func init() {
	proto.RegisterFile("copy-trading/exemplary_trader.proto", fileDescriptor_973e30b53d12a7a1)
}

var fileDescriptor_973e30b53d12a7a1 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0x42, 0x31,
	0x18, 0xc5, 0xef, 0x55, 0xa2, 0xf1, 0x3a, 0x98, 0x34, 0xc4, 0x5c, 0x1d, 0x0a, 0xd1, 0xc4, 0xb8,
	0xd0, 0x0e, 0xbe, 0x01, 0xfe, 0xdb, 0x89, 0x2c, 0x2e, 0xa4, 0xb4, 0xe5, 0xf2, 0x45, 0xdb, 0xaf,
	0x69, 0x4b, 0x02, 0x6f, 0xe1, 0x5b, 0xc9, 0xc8, 0x68, 0x1c, 0x88, 0x81, 0x17, 0x31, 0xb7, 0x80,
	0xc1, 0xa9, 0xa7, 0xa7, 0xe7, 0xfc, 0x9a, 0x9c, 0xe2, 0x5a, 0xa2, 0x9b, 0x75, 0xa2, 0x17, 0x0a,
	0x6c, 0xc5, 0xf5, 0x54, 0x1b, 0xf7, 0x2e, 0xfc, 0x6c, 0x50, 0x3b, 0xda, 0x33, 0xe7, 0x31, 0x22,
	0xb9, 0x98, 0xd8, 0x89, 0x85, 0x11, 0x30, 0x39, 0x16, 0x60, 0x59, 0x5d, 0xd9, 0x36, 0x2e, 0x9b,
	0x15, 0x56, 0x98, 0x52, 0xbc, 0x56, 0x9b, 0xc2, 0xd5, 0x67, 0x5e, 0x9c, 0x3d, 0xee, 0x58, 0x2f,
	0x09, 0x45, 0xca, 0xe2, 0x58, 0x28, 0xe5, 0x75, 0x08, 0x65, 0xde, 0xce, 0x6f, 0x4f, 0x7a, 0xbb,
	0x2b, 0x21, 0x45, 0xc3, 0x0a, 0xa3, 0xcb, 0x83, 0x64, 0x27, 0x4d, 0xda, 0xc5, 0xa9, 0xd2, 0x41,
	0x7a, 0x70, 0x11, 0xd0, 0x96, 0x87, 0xe9, 0x69, 0xdf, 0x22, 0xaa, 0x38, 0x77, 0x1e, 0x47, 0x10,
	0x07, 0x12, 0x8d, 0x81, 0x10, 0x00, 0xed, 0xc0, 0x8b, 0xa8, 0xcb, 0x46, 0x1d, 0xee, 0xb2, 0xf9,
	0xb2, 0x95, 0x7d, 0x2f, 0x5b, 0x37, 0x15, 0xc4, 0xf1, 0x64, 0xc8, 0x24, 0x1a, 0x2e, 0x31, 0x18,
	0x0c, 0xdb, 0xa3, 0x13, 0xd4, 0x1b, 0x8f, 0x33, 0xa7, 0x03, 0x7b, 0xd0, 0xb2, 0xd7, 0xdc, 0xd0,
	0xee, 0xff, 0x60, 0x3d, 0x11, 0x75, 0xf7, 0x79, 0xbe, 0xa2, 0xf9, 0x62, 0x45, 0xf3, 0x9f, 0x15,
	0xcd, 0x3f, 0xd6, 0x34, 0x5b, 0xac, 0x69, 0xf6, 0xb5, 0xa6, 0xd9, 0x6b, 0x67, 0x8f, 0xdb, 0xb7,
	0x7d, 0x0b, 0x4f, 0xc0, 0xd3, 0x3e, 0x7c, 0xca, 0xff, 0x8d, 0x9a, 0xbe, 0x18, 0x1e, 0xa5, 0x65,
	0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x7e, 0x7f, 0xcd, 0x71, 0x01, 0x00, 0x00,
}

func (m *ExemplaryTrader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExemplaryTrader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExemplaryTrader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ProfitCommissionRate.Size()
		i -= size
		if _, err := m.ProfitCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExemplaryTrader(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintExemplaryTrader(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintExemplaryTrader(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintExemplaryTrader(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExemplaryTrader(dAtA []byte, offset int, v uint64) int {
	offset -= sovExemplaryTrader(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExemplaryTrader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovExemplaryTrader(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExemplaryTrader(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovExemplaryTrader(uint64(l))
	}
	l = m.ProfitCommissionRate.Size()
	n += 1 + l + sovExemplaryTrader(uint64(l))
	return n
}

func sovExemplaryTrader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExemplaryTrader(x uint64) (n int) {
	return sovExemplaryTrader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExemplaryTrader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExemplaryTrader
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
			return fmt.Errorf("proto: ExemplaryTrader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExemplaryTrader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExemplaryTrader
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
				return ErrInvalidLengthExemplaryTrader
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExemplaryTrader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExemplaryTrader
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
				return ErrInvalidLengthExemplaryTrader
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExemplaryTrader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExemplaryTrader
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
				return ErrInvalidLengthExemplaryTrader
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExemplaryTrader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfitCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExemplaryTrader
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
				return ErrInvalidLengthExemplaryTrader
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExemplaryTrader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProfitCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExemplaryTrader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExemplaryTrader
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
func skipExemplaryTrader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExemplaryTrader
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
					return 0, ErrIntOverflowExemplaryTrader
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
					return 0, ErrIntOverflowExemplaryTrader
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
				return 0, ErrInvalidLengthExemplaryTrader
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExemplaryTrader
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExemplaryTrader
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExemplaryTrader        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExemplaryTrader          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExemplaryTrader = fmt.Errorf("proto: unexpected end of group")
)
