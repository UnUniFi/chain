// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ununifi/kyc/verification.proto

package types

import (
	fmt "fmt"
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

type Verification struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ProviderId uint64 `protobuf:"varint,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
}

func (m *Verification) Reset()         { *m = Verification{} }
func (m *Verification) String() string { return proto.CompactTextString(m) }
func (*Verification) ProtoMessage()    {}
func (*Verification) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd365ef280640105, []int{0}
}
func (m *Verification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Verification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Verification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Verification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verification.Merge(m, src)
}
func (m *Verification) XXX_Size() int {
	return m.Size()
}
func (m *Verification) XXX_DiscardUnknown() {
	xxx_messageInfo_Verification.DiscardUnknown(m)
}

var xxx_messageInfo_Verification proto.InternalMessageInfo

func (m *Verification) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Verification) GetProviderId() uint64 {
	if m != nil {
		return m.ProviderId
	}
	return 0
}

func init() {
	proto.RegisterType((*Verification)(nil), "ununifi.kyc.Verification")
}

func init() { proto.RegisterFile("ununifi/kyc/verification.proto", fileDescriptor_fd365ef280640105) }

var fileDescriptor_fd365ef280640105 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xcd, 0x2b, 0xcd,
	0xcb, 0x4c, 0xcb, 0xd4, 0xcf, 0xae, 0x4c, 0xd6, 0x2f, 0x4b, 0x2d, 0xca, 0x4c, 0xcb, 0x4c, 0x4e,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xca, 0xeb, 0x65,
	0x57, 0x26, 0x2b, 0x79, 0x72, 0xf1, 0x84, 0x21, 0x29, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49,
	0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xe4, 0xb9,
	0xb8, 0x0b, 0x8a, 0xf2, 0xcb, 0x32, 0x53, 0x52, 0x8b, 0xe2, 0x33, 0x53, 0x24, 0x98, 0x14, 0x18,
	0x35, 0x58, 0x82, 0xb8, 0x60, 0x42, 0x9e, 0x29, 0x4e, 0x76, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc,
	0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x1f, 0x9a, 0x17, 0x9a, 0x97, 0xe9, 0x96, 0xa9, 0x9f, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x5f, 0x01,
	0x76, 0x64, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x79, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x96, 0x17, 0x22, 0x26, 0xc0, 0x00, 0x00, 0x00,
}

func (m *Verification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Verification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Verification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProviderId != 0 {
		i = encodeVarintVerification(dAtA, i, uint64(m.ProviderId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintVerification(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVerification(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerification(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Verification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovVerification(uint64(l))
	}
	if m.ProviderId != 0 {
		n += 1 + sovVerification(uint64(m.ProviderId))
	}
	return n
}

func sovVerification(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerification(x uint64) (n int) {
	return sovVerification(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Verification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerification
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
			return fmt.Errorf("proto: Verification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Verification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerification
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
				return ErrInvalidLengthVerification
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderId", wireType)
			}
			m.ProviderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerification
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProviderId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVerification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerification
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
func skipVerification(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerification
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
					return 0, ErrIntOverflowVerification
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
					return 0, ErrIntOverflowVerification
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
				return 0, ErrInvalidLengthVerification
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerification
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerification
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerification        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerification          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerification = fmt.Errorf("proto: unexpected end of group")
)
