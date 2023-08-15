// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ununifi/ecosystemincentive/ecosystemincentive.proto

package types

import (
	fmt "fmt"
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

type RewardRecord struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Rewards github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"rewards" yaml:"rewards"`
}

func (m *RewardRecord) Reset()         { *m = RewardRecord{} }
func (m *RewardRecord) String() string { return proto.CompactTextString(m) }
func (*RewardRecord) ProtoMessage()    {}
func (*RewardRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a12ba5b7ba43f547, []int{0}
}
func (m *RewardRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardRecord.Merge(m, src)
}
func (m *RewardRecord) XXX_Size() int {
	return m.Size()
}
func (m *RewardRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardRecord.DiscardUnknown(m)
}

var xxx_messageInfo_RewardRecord proto.InternalMessageInfo

func (m *RewardRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RewardRecord) GetRewards() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func init() {
	proto.RegisterType((*RewardRecord)(nil), "ununifi.ecosystemincentive.RewardRecord")
}

func init() {
	proto.RegisterFile("ununifi/ecosystemincentive/ecosystemincentive.proto", fileDescriptor_a12ba5b7ba43f547)
}

var fileDescriptor_a12ba5b7ba43f547 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x90, 0xa8, 0x08, 0x88, 0x21, 0x62, 0x28, 0x1d, 0xdc, 0x2a, 0x53, 0x07, 0xb0,
	0x55, 0xb2, 0x31, 0x06, 0x89, 0x8d, 0x25, 0x52, 0x17, 0x36, 0xc7, 0x31, 0xa9, 0x05, 0xf1, 0xab,
	0xf2, 0x9c, 0x96, 0xdc, 0x82, 0x73, 0xb0, 0x73, 0x87, 0x8e, 0x1d, 0x99, 0x0a, 0x4a, 0x6e, 0xc0,
	0x09, 0x50, 0x93, 0x54, 0x20, 0xd1, 0xc9, 0x96, 0xdf, 0xff, 0x7f, 0x9f, 0xf5, 0xdc, 0xa0, 0x30,
	0x85, 0xd1, 0x8f, 0x9a, 0x2b, 0x09, 0x58, 0xa2, 0x55, 0x99, 0x36, 0x52, 0x19, 0xab, 0x17, 0x6a,
	0xcf, 0x13, 0x9b, 0xe7, 0x60, 0xc1, 0x1b, 0x74, 0x25, 0xf6, 0x3f, 0x31, 0x38, 0x4f, 0x21, 0x85,
	0x26, 0xc6, 0xb7, 0xb7, 0xb6, 0x31, 0xa0, 0x12, 0x30, 0x03, 0xe4, 0xb1, 0x40, 0xc5, 0x17, 0x93,
	0x58, 0x59, 0x31, 0xe1, 0x12, 0xb4, 0x69, 0xe7, 0xfe, 0x3b, 0x71, 0x4f, 0x23, 0xb5, 0x14, 0x79,
	0x12, 0x29, 0x09, 0x79, 0xe2, 0x5d, 0xba, 0x3d, 0x91, 0x24, 0xb9, 0x42, 0xec, 0x93, 0x11, 0x19,
	0x1f, 0x87, 0xde, 0xf7, 0x66, 0x78, 0x56, 0x8a, 0xec, 0xf9, 0xc6, 0xef, 0x06, 0x7e, 0xb4, 0x8b,
	0x78, 0x4b, 0xb7, 0x97, 0x37, 0x6d, 0xec, 0x1f, 0x8c, 0x0e, 0xc7, 0x27, 0xd7, 0x17, 0xac, 0x15,
	0xb2, 0xad, 0x90, 0x75, 0x42, 0x76, 0x0b, 0xda, 0x84, 0xe1, 0x6a, 0x33, 0x74, 0x7e, 0x61, 0x5d,
	0xcf, 0x7f, 0xfb, 0x1c, 0x8e, 0x53, 0x6d, 0x67, 0x45, 0xcc, 0x24, 0x64, 0xbc, 0xfb, 0x6f, 0x7b,
	0x5c, 0x61, 0xf2, 0xc4, 0x6d, 0x39, 0x57, 0xd8, 0x20, 0x30, 0xda, 0xd9, 0xc2, 0xfb, 0x55, 0x45,
	0xc9, 0xba, 0xa2, 0xe4, 0xab, 0xa2, 0xe4, 0xb5, 0xa6, 0xce, 0xba, 0xa6, 0xce, 0x47, 0x4d, 0x9d,
	0x87, 0xe0, 0x0f, 0x6c, 0x6a, 0xa6, 0x46, 0xdf, 0x69, 0x2e, 0x67, 0x42, 0x1b, 0xfe, 0xb2, 0x6f,
	0xd7, 0x0d, 0x3d, 0x3e, 0x6a, 0xb6, 0x11, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x3e, 0x3c,
	0xee, 0x96, 0x01, 0x00, 0x00,
}

func (m *RewardRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEcosystemincentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEcosystemincentive(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEcosystemincentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovEcosystemincentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEcosystemincentive(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovEcosystemincentive(uint64(l))
		}
	}
	return n
}

func sovEcosystemincentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEcosystemincentive(x uint64) (n int) {
	return sovEcosystemincentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RewardRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEcosystemincentive
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
			return fmt.Errorf("proto: RewardRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemincentive
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
				return ErrInvalidLengthEcosystemincentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemincentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemincentive
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
				return ErrInvalidLengthEcosystemincentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemincentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEcosystemincentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEcosystemincentive
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
func skipEcosystemincentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEcosystemincentive
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
					return 0, ErrIntOverflowEcosystemincentive
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
					return 0, ErrIntOverflowEcosystemincentive
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
				return 0, ErrInvalidLengthEcosystemincentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEcosystemincentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEcosystemincentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEcosystemincentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEcosystemincentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEcosystemincentive = fmt.Errorf("proto: unexpected end of group")
)
