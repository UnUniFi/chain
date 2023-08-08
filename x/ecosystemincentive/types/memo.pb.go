// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ununifi/ecosystemincentive/memo.proto

package types

import (
	fmt "fmt"
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

type FrontendMetadata struct {
	Version   uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty" yaml:"version"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
}

func (m *FrontendMetadata) Reset()         { *m = FrontendMetadata{} }
func (m *FrontendMetadata) String() string { return proto.CompactTextString(m) }
func (*FrontendMetadata) ProtoMessage()    {}
func (*FrontendMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a43f0dcca886bed, []int{0}
}
func (m *FrontendMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FrontendMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FrontendMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FrontendMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendMetadata.Merge(m, src)
}
func (m *FrontendMetadata) XXX_Size() int {
	return m.Size()
}
func (m *FrontendMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendMetadata proto.InternalMessageInfo

func (m *FrontendMetadata) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FrontendMetadata) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func init() {
	proto.RegisterType((*FrontendMetadata)(nil), "ununifi.ecosystemincentive.FrontendMetadata")
}

func init() {
	proto.RegisterFile("ununifi/ecosystemincentive/memo.proto", fileDescriptor_6a43f0dcca886bed)
}

var fileDescriptor_6a43f0dcca886bed = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xcd, 0x2b, 0xcd,
	0xcb, 0x4c, 0xcb, 0xd4, 0x4f, 0x4d, 0xce, 0x2f, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0xcd, 0xcc, 0x4b,
	0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0xcf, 0x4d, 0xcd, 0xcd, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x82, 0x2a, 0xd3, 0xc3, 0x54, 0x26, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56,
	0xa6, 0x0f, 0x62, 0x41, 0x74, 0x28, 0x95, 0x70, 0x09, 0xb8, 0x15, 0xe5, 0xe7, 0x95, 0xa4, 0xe6,
	0xa5, 0xf8, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0x0a, 0xe9, 0x70, 0xb1, 0x97, 0xa5, 0x16,
	0x15, 0x67, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x3a, 0x09, 0x7d, 0xba, 0x27, 0xcf,
	0x57, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x04, 0x95, 0x50, 0x0a, 0x82, 0x29, 0x11, 0x32, 0xe2, 0xe2,
	0x2c, 0x4a, 0x4d, 0xce, 0x2c, 0xc8, 0x4c, 0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x74,
	0x12, 0xf9, 0x74, 0x4f, 0x5e, 0x00, 0xa2, 0x1e, 0x2e, 0xa5, 0x14, 0x84, 0x50, 0xe6, 0xe4, 0x7b,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7,
	0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xc6, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xa1, 0x79, 0xa1, 0x79, 0x99, 0x6e, 0x99, 0xfa, 0xc9, 0x19,
	0x89, 0x99, 0x79, 0xfa, 0x15, 0xd8, 0xfc, 0x5e, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6,
	0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x88, 0xdf, 0xdc, 0x19, 0x26, 0x01, 0x00, 0x00,
}

func (m *FrontendMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FrontendMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FrontendMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintMemo(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintMemo(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMemo(dAtA []byte, offset int, v uint64) int {
	offset -= sovMemo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FrontendMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovMemo(uint64(m.Version))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovMemo(uint64(l))
	}
	return n
}

func sovMemo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMemo(x uint64) (n int) {
	return sovMemo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FrontendMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMemo
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
			return fmt.Errorf("proto: FrontendMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FrontendMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemo
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
				return ErrInvalidLengthMemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMemo
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
func skipMemo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMemo
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
					return 0, ErrIntOverflowMemo
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
					return 0, ErrIntOverflowMemo
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
				return 0, ErrInvalidLengthMemo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMemo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMemo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMemo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMemo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMemo = fmt.Errorf("proto: unexpected end of group")
)
