// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ununifi/eventhook/eventhook.proto

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

type Hook struct {
	Id              uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ContractAddress string          `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	GitUrl          string          `protobuf:"bytes,4,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
	EventType       string          `protobuf:"bytes,5,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	EventAttributes []*KeyValuePair `protobuf:"bytes,6,rep,name=event_attributes,json=eventAttributes,proto3" json:"event_attributes,omitempty"`
}

func (m *Hook) Reset()         { *m = Hook{} }
func (m *Hook) String() string { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()    {}
func (*Hook) Descriptor() ([]byte, []int) {
	return fileDescriptor_8298502638ea5b43, []int{0}
}
func (m *Hook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Hook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Hook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Hook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hook.Merge(m, src)
}
func (m *Hook) XXX_Size() int {
	return m.Size()
}
func (m *Hook) XXX_DiscardUnknown() {
	xxx_messageInfo_Hook.DiscardUnknown(m)
}

var xxx_messageInfo_Hook proto.InternalMessageInfo

func (m *Hook) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Hook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hook) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Hook) GetGitUrl() string {
	if m != nil {
		return m.GitUrl
	}
	return ""
}

func (m *Hook) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Hook) GetEventAttributes() []*KeyValuePair {
	if m != nil {
		return m.EventAttributes
	}
	return nil
}

type KeyValuePair struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_8298502638ea5b43, []int{1}
}
func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return m.Size()
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Hook)(nil), "nftvault.eventhook.Hook")
	proto.RegisterType((*KeyValuePair)(nil), "nftvault.eventhook.KeyValuePair")
}

func init() { proto.RegisterFile("ununifi/eventhook/eventhook.proto", fileDescriptor_8298502638ea5b43) }

var fileDescriptor_8298502638ea5b43 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x97, 0xad, 0x9b, 0x2c, 0x8a, 0x1b, 0x61, 0x60, 0x11, 0x0c, 0x75, 0xa7, 0xe9, 0xa1,
	0x05, 0x05, 0xef, 0x13, 0x11, 0x61, 0x17, 0x29, 0xce, 0x83, 0x97, 0x91, 0xb5, 0x59, 0x17, 0xd6,
	0x25, 0x23, 0x7d, 0x19, 0xf6, 0xbf, 0xf0, 0xcf, 0xf2, 0xb8, 0xa3, 0x78, 0x92, 0xf5, 0x1f, 0x91,
	0xa6, 0x3a, 0x07, 0xde, 0xbe, 0x7c, 0xbf, 0x2f, 0xf0, 0xbe, 0xf7, 0xf0, 0xb9, 0x91, 0x46, 0x8a,
	0x99, 0x08, 0xf8, 0x9a, 0x4b, 0x98, 0x2b, 0xb5, 0xf8, 0x53, 0xfe, 0x4a, 0x2b, 0x50, 0x84, 0xc8,
	0x19, 0xac, 0x99, 0x49, 0xc1, 0xdf, 0x91, 0xd3, 0x5e, 0xa2, 0x12, 0x65, 0x71, 0x50, 0xaa, 0x2a,
	0xd9, 0xff, 0x44, 0xd8, 0x79, 0x50, 0x6a, 0x41, 0x8e, 0x71, 0x5d, 0xc4, 0x2e, 0xf2, 0xd0, 0xc0,
	0x09, 0xeb, 0x22, 0x26, 0x04, 0x3b, 0x92, 0x2d, 0xb9, 0x5b, 0xf7, 0xd0, 0xa0, 0x1d, 0x5a, 0x4d,
	0x2e, 0x70, 0x37, 0x52, 0x12, 0x34, 0x8b, 0x60, 0xc2, 0xe2, 0x58, 0xf3, 0x2c, 0x73, 0x1b, 0x96,
	0x77, 0x7e, 0xfd, 0x61, 0x65, 0x93, 0x13, 0x7c, 0x90, 0x08, 0x98, 0x18, 0x9d, 0xba, 0x8e, 0x4d,
	0xb4, 0x12, 0x01, 0x63, 0x9d, 0x92, 0x33, 0x8c, 0xed, 0x4c, 0x13, 0xc8, 0x57, 0xdc, 0x6d, 0x5a,
	0xd6, 0xb6, 0xce, 0x53, 0xbe, 0xe2, 0x64, 0x84, 0xbb, 0x15, 0x66, 0x00, 0x5a, 0x4c, 0x0d, 0xf0,
	0xcc, 0x6d, 0x79, 0x8d, 0xc1, 0xe1, 0x95, 0xe7, 0xff, 0x2f, 0xe5, 0x8f, 0x78, 0xfe, 0xcc, 0x52,
	0xc3, 0x1f, 0x99, 0xd0, 0x61, 0xc7, 0xfa, 0xc3, 0xdd, 0xc7, 0xfe, 0x0d, 0x3e, 0xda, 0x0f, 0x90,
	0x2e, 0x6e, 0x2c, 0x78, 0x6e, 0x4b, 0xb6, 0xc3, 0x52, 0x92, 0x1e, 0x6e, 0xae, 0x4b, 0xfc, 0x53,
	0xb3, 0x7a, 0xdc, 0xde, 0xbd, 0x6f, 0x29, 0xda, 0x6c, 0x29, 0xfa, 0xda, 0x52, 0xf4, 0x56, 0xd0,
	0xda, 0xa6, 0xa0, 0xb5, 0x8f, 0x82, 0xd6, 0x5e, 0x2e, 0x13, 0x01, 0x73, 0x33, 0xf5, 0x23, 0xb5,
	0x0c, 0xc6, 0x72, 0x2c, 0xc5, 0xbd, 0x08, 0xa2, 0x39, 0x13, 0x32, 0x78, 0xdd, 0x3b, 0x47, 0x59,
	0x2d, 0x9b, 0xb6, 0xec, 0x86, 0xaf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x45, 0xf4, 0x77,
	0xb0, 0x01, 0x00, 0x00,
}

func (m *Hook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Hook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EventAttributes) > 0 {
		for iNdEx := len(m.EventAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EventAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEventhook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.EventType) > 0 {
		i -= len(m.EventType)
		copy(dAtA[i:], m.EventType)
		i = encodeVarintEventhook(dAtA, i, uint64(len(m.EventType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GitUrl) > 0 {
		i -= len(m.GitUrl)
		copy(dAtA[i:], m.GitUrl)
		i = encodeVarintEventhook(dAtA, i, uint64(len(m.GitUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEventhook(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEventhook(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEventhook(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *KeyValuePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyValuePair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyValuePair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintEventhook(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintEventhook(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEventhook(dAtA []byte, offset int, v uint64) int {
	offset -= sovEventhook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Hook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEventhook(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEventhook(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEventhook(uint64(l))
	}
	l = len(m.GitUrl)
	if l > 0 {
		n += 1 + l + sovEventhook(uint64(l))
	}
	l = len(m.EventType)
	if l > 0 {
		n += 1 + l + sovEventhook(uint64(l))
	}
	if len(m.EventAttributes) > 0 {
		for _, e := range m.EventAttributes {
			l = e.Size()
			n += 1 + l + sovEventhook(uint64(l))
		}
	}
	return n
}

func (m *KeyValuePair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEventhook(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovEventhook(uint64(l))
	}
	return n
}

func sovEventhook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEventhook(x uint64) (n int) {
	return sovEventhook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Hook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventhook
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
			return fmt.Errorf("proto: Hook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
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
				return ErrInvalidLengthEventhook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventhook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
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
				return ErrInvalidLengthEventhook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventhook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
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
				return ErrInvalidLengthEventhook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventhook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
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
				return ErrInvalidLengthEventhook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventhook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
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
				return ErrInvalidLengthEventhook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEventhook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventAttributes = append(m.EventAttributes, &KeyValuePair{})
			if err := m.EventAttributes[len(m.EventAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventhook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEventhook
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
func (m *KeyValuePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventhook
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
			return fmt.Errorf("proto: KeyValuePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValuePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
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
				return ErrInvalidLengthEventhook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventhook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventhook
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
				return ErrInvalidLengthEventhook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventhook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventhook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEventhook
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
func skipEventhook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventhook
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
					return 0, ErrIntOverflowEventhook
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
					return 0, ErrIntOverflowEventhook
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
				return 0, ErrInvalidLengthEventhook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEventhook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEventhook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEventhook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEventhook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEventhook = fmt.Errorf("proto: unexpected end of group")
)
