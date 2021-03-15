// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jsmndist/genesis.proto

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

// GenesisState defines the jsmndist module's genesis state.
type GenesisState struct {
	Params            *Params   `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty" yaml:"params"`
	PreviousBlockTime time.Time `protobuf:"bytes,2,opt,name=previous_block_time,json=previousBlockTime,proto3,stdtime" json:"previous_block_time" yaml:"previous_block_time"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e0325b935fb7b9e, []int{0}
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

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetPreviousBlockTime() time.Time {
	if m != nil {
		return m.PreviousBlockTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "jpyx.jsmndist.GenesisState")
}

func init() { proto.RegisterFile("jsmndist/genesis.proto", fileDescriptor_7e0325b935fb7b9e) }

var fileDescriptor_7e0325b935fb7b9e = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x86, 0x63, 0x86, 0x0e, 0x81, 0x0e, 0x0d, 0x7f, 0x55, 0x06, 0x07, 0x65, 0xa8, 0x98, 0x6c,
	0x09, 0x36, 0xa6, 0x2a, 0x0b, 0x2b, 0x2a, 0x4c, 0x2c, 0x55, 0x12, 0x8c, 0x49, 0x89, 0x63, 0x2b,
	0x76, 0x50, 0x73, 0x8b, 0x1e, 0x8a, 0xa1, 0x63, 0x47, 0xa6, 0x80, 0x92, 0x1b, 0xf4, 0x04, 0x28,
	0x76, 0x5c, 0x09, 0xb1, 0x7d, 0xd6, 0xf3, 0x7e, 0xaf, 0x1f, 0xdb, 0xbd, 0x58, 0x49, 0x56, 0xbc,
	0x64, 0x52, 0x61, 0x4a, 0x0a, 0x22, 0x33, 0x89, 0x44, 0xc9, 0x15, 0xf7, 0xc6, 0x2b, 0x51, 0xaf,
	0x91, 0x85, 0xfe, 0x19, 0xe5, 0x94, 0x6b, 0x82, 0xfb, 0xc9, 0x84, 0xfc, 0x80, 0x72, 0x4e, 0x73,
	0x82, 0xf5, 0x29, 0xa9, 0x5e, 0xb1, 0xca, 0x18, 0x91, 0x2a, 0x66, 0x62, 0x08, 0x5c, 0x1e, 0xda,
	0xed, 0x60, 0x40, 0xf8, 0x09, 0xdc, 0x93, 0x7b, 0x73, 0xe1, 0xa3, 0x8a, 0x15, 0xf1, 0xe6, 0xee,
	0x48, 0xc4, 0x65, 0xcc, 0xe4, 0x14, 0x5c, 0x81, 0xeb, 0xe3, 0x9b, 0x73, 0xf4, 0x47, 0x00, 0x3d,
	0x68, 0x18, 0x4d, 0xf6, 0x4d, 0x30, 0xae, 0x63, 0x96, 0xdf, 0x85, 0x26, 0x1e, 0x2e, 0x86, 0x3d,
	0xaf, 0x74, 0x4f, 0x45, 0x49, 0x3e, 0x32, 0x5e, 0xc9, 0x65, 0x92, 0xf3, 0xf4, 0x7d, 0xd9, 0xdb,
	0x4c, 0x8f, 0x74, 0x9d, 0x8f, 0x8c, 0x2a, 0xb2, 0xaa, 0xe8, 0xc9, 0xaa, 0x46, 0xb3, 0x6d, 0x13,
	0x38, 0xfb, 0x26, 0xf0, 0x87, 0xde, 0xff, 0x25, 0xe1, 0xe6, 0x3b, 0x00, 0x8b, 0x89, 0x25, 0x51,
	0x0f, 0xfa, 0xfd, 0x68, 0xbe, 0x6d, 0x21, 0xd8, 0xb5, 0x10, 0xfc, 0xb4, 0x10, 0x6c, 0x3a, 0xe8,
	0xec, 0x3a, 0xe8, 0x7c, 0x75, 0xd0, 0x79, 0x9e, 0xd1, 0x4c, 0xbd, 0x55, 0x09, 0x4a, 0x39, 0xc3,
	0x79, 0x5a, 0x10, 0x86, 0xfb, 0xf7, 0xe0, 0xf5, 0xe1, 0x23, 0xb0, 0xaa, 0x05, 0x91, 0xc9, 0x48,
	0x0b, 0xdd, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x76, 0xe5, 0x52, 0x3f, 0x88, 0x01, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousBlockTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Params != nil {
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousBlockTime)
	n += 1 + l + sovGenesis(uint64(l))
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
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousBlockTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousBlockTime, dAtA[iNdEx:postIndex]); err != nil {
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
