// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jsmndist/jsmndist.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	Active  bool      `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty" yaml:"active"`
	Periods []*Period `protobuf:"bytes,2,rep,name=periods,proto3" json:"periods,omitempty" yaml:"periods"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c639fed8d674e80e, []int{0}
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

func (m *Params) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Params) GetPeriods() []*Period {
	if m != nil {
		return m.Periods
	}
	return nil
}

type Period struct {
	Start     time.Time                              `protobuf:"bytes,1,opt,name=start,proto3,stdtime" json:"start" yaml:"start"`
	End       time.Time                              `protobuf:"bytes,2,opt,name=end,proto3,stdtime" json:"end" yaml:"end"`
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation" yaml:"inflation"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_c639fed8d674e80e, []int{1}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *Period) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Params)(nil), "jpyx.jsmndist.Params")
	proto.RegisterType((*Period)(nil), "jpyx.jsmndist.Period")
}

func init() { proto.RegisterFile("jsmndist/jsmndist.proto", fileDescriptor_c639fed8d674e80e) }

var fileDescriptor_c639fed8d674e80e = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0xee, 0xd3, 0x30,
	0x10, 0xc7, 0xe3, 0x5f, 0x45, 0xa0, 0x2e, 0x45, 0x10, 0xf1, 0x27, 0xea, 0x10, 0x57, 0x19, 0x50,
	0x19, 0xb0, 0xd5, 0xb2, 0x31, 0xa1, 0xd0, 0x89, 0xa9, 0x8a, 0x98, 0x98, 0x70, 0x12, 0x37, 0xb8,
	0xc4, 0x71, 0x14, 0xbb, 0x55, 0xfb, 0x16, 0x7d, 0xac, 0x8e, 0x1d, 0x11, 0x43, 0x40, 0xed, 0x1b,
	0x64, 0x66, 0x40, 0x89, 0x93, 0x02, 0x13, 0x93, 0xcf, 0x77, 0xdf, 0xfb, 0x9c, 0xbf, 0x27, 0xc3,
	0x17, 0x1b, 0x25, 0xf2, 0x84, 0x2b, 0x4d, 0xfa, 0x00, 0x17, 0xa5, 0xd4, 0xd2, 0x19, 0x6f, 0x8a,
	0xc3, 0x1e, 0xf7, 0xc9, 0xc9, 0xd3, 0x54, 0xa6, 0xb2, 0xad, 0x90, 0x26, 0x32, 0xa2, 0x09, 0x4a,
	0xa5, 0x4c, 0x33, 0x46, 0xda, 0x5b, 0xb4, 0x5d, 0x13, 0xcd, 0x05, 0x53, 0x9a, 0x8a, 0xa2, 0x13,
	0x78, 0xb1, 0x54, 0x42, 0x2a, 0x12, 0x51, 0xc5, 0xc8, 0x6e, 0x1e, 0x31, 0x4d, 0xe7, 0x24, 0x96,
	0x3c, 0x37, 0x75, 0x7f, 0x0f, 0xed, 0x15, 0x2d, 0xa9, 0x50, 0xce, 0x2b, 0x68, 0xd3, 0x58, 0xf3,
	0x1d, 0x73, 0xc1, 0x14, 0xcc, 0x1e, 0x04, 0x4f, 0xea, 0x0a, 0x8d, 0x0f, 0x54, 0x64, 0x6f, 0x7d,
	0x93, 0xf7, 0xc3, 0x4e, 0xe0, 0xbc, 0x87, 0xf7, 0x0b, 0x56, 0x72, 0x99, 0x28, 0xf7, 0x6e, 0x3a,
	0x98, 0x8d, 0x16, 0xcf, 0xf0, 0x3f, 0x8f, 0xc5, 0xab, 0xb6, 0x1a, 0x38, 0x75, 0x85, 0x1e, 0x19,
	0x44, 0xa7, 0xf7, 0xc3, 0xbe, 0xd3, 0xff, 0x05, 0xa0, 0x6d, 0x74, 0xce, 0x07, 0x78, 0x4f, 0x69,
	0x5a, 0xea, 0x76, 0xf2, 0x68, 0x31, 0xc1, 0xc6, 0x15, 0xee, 0x5d, 0xe1, 0x8f, 0xbd, 0xab, 0xc0,
	0x3d, 0x55, 0xc8, 0xaa, 0x2b, 0xf4, 0xd0, 0x60, 0xdb, 0x36, 0xff, 0xf8, 0x03, 0x81, 0xd0, 0x20,
	0x9c, 0x25, 0x1c, 0xb0, 0x3c, 0x71, 0xef, 0xfe, 0x4b, 0x7a, 0xde, 0x91, 0xa0, 0x21, 0xb1, 0x3c,
	0x31, 0x9c, 0xa6, 0xdd, 0xf9, 0x0c, 0x87, 0x3c, 0x5f, 0x67, 0x54, 0x73, 0x99, 0xbb, 0xa3, 0x29,
	0x98, 0x0d, 0x83, 0xa0, 0xd1, 0x7f, 0xaf, 0xd0, 0xcb, 0x94, 0xeb, 0x2f, 0xdb, 0x08, 0xc7, 0x52,
	0x90, 0x6e, 0xb9, 0xe6, 0x78, 0xad, 0x92, 0xaf, 0x44, 0x1f, 0x0a, 0xa6, 0xf0, 0x92, 0xc5, 0x75,
	0x85, 0x1e, 0x1b, 0xf2, 0x0d, 0xe4, 0x87, 0x7f, 0xa0, 0xc1, 0xbb, 0xd3, 0xc5, 0x03, 0xe7, 0x8b,
	0x07, 0x7e, 0x5e, 0x3c, 0x70, 0xbc, 0x7a, 0xd6, 0xf9, 0xea, 0x59, 0xdf, 0xae, 0x9e, 0xf5, 0xe9,
	0xef, 0x01, 0x59, 0x9c, 0x33, 0x41, 0x9a, 0xe5, 0x92, 0xfd, 0xed, 0x83, 0x98, 0x21, 0x91, 0xdd,
	0x9a, 0x7a, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x89, 0xdf, 0x30, 0x2a, 0x42, 0x02, 0x00, 0x00,
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
	if len(m.Periods) > 0 {
		for iNdEx := len(m.Periods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Periods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintJsmndist(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintJsmndist(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintJsmndist(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintJsmndist(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintJsmndist(dAtA []byte, offset int, v uint64) int {
	offset -= sovJsmndist(v)
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
	if m.Active {
		n += 2
	}
	if len(m.Periods) > 0 {
		for _, e := range m.Periods {
			l = e.Size()
			n += 1 + l + sovJsmndist(uint64(l))
		}
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovJsmndist(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovJsmndist(uint64(l))
	l = m.Inflation.Size()
	n += 1 + l + sovJsmndist(uint64(l))
	return n
}

func sovJsmndist(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJsmndist(x uint64) (n int) {
	return sovJsmndist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJsmndist
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
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJsmndist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Periods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJsmndist
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
				return ErrInvalidLengthJsmndist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJsmndist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Periods = append(m.Periods, &Period{})
			if err := m.Periods[len(m.Periods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJsmndist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJsmndist
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJsmndist
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJsmndist
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
				return ErrInvalidLengthJsmndist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJsmndist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJsmndist
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
				return ErrInvalidLengthJsmndist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJsmndist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJsmndist
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
				return ErrInvalidLengthJsmndist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJsmndist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJsmndist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJsmndist
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
func skipJsmndist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJsmndist
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
					return 0, ErrIntOverflowJsmndist
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
					return 0, ErrIntOverflowJsmndist
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
				return 0, ErrInvalidLengthJsmndist
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJsmndist
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJsmndist
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJsmndist        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJsmndist          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJsmndist = fmt.Errorf("proto: unexpected end of group")
)
