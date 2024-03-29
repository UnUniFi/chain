// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ununifi/ecosystemincentive/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgWithdrawAllRewards struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}

func (m *MsgWithdrawAllRewards) Reset()         { *m = MsgWithdrawAllRewards{} }
func (m *MsgWithdrawAllRewards) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawAllRewards) ProtoMessage()    {}
func (*MsgWithdrawAllRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_df57aae78a8732aa, []int{0}
}
func (m *MsgWithdrawAllRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawAllRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawAllRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawAllRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawAllRewards.Merge(m, src)
}
func (m *MsgWithdrawAllRewards) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawAllRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawAllRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawAllRewards proto.InternalMessageInfo

func (m *MsgWithdrawAllRewards) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgWithdrawAllRewardsResponse struct {
}

func (m *MsgWithdrawAllRewardsResponse) Reset()         { *m = MsgWithdrawAllRewardsResponse{} }
func (m *MsgWithdrawAllRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawAllRewardsResponse) ProtoMessage()    {}
func (*MsgWithdrawAllRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df57aae78a8732aa, []int{1}
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawAllRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawAllRewardsResponse.Merge(m, src)
}
func (m *MsgWithdrawAllRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawAllRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawAllRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawAllRewardsResponse proto.InternalMessageInfo

type MsgWithdrawReward struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	Denom  string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *MsgWithdrawReward) Reset()         { *m = MsgWithdrawReward{} }
func (m *MsgWithdrawReward) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawReward) ProtoMessage()    {}
func (*MsgWithdrawReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_df57aae78a8732aa, []int{2}
}
func (m *MsgWithdrawReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawReward.Merge(m, src)
}
func (m *MsgWithdrawReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawReward proto.InternalMessageInfo

func (m *MsgWithdrawReward) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgWithdrawReward) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type MsgWithdrawRewardResponse struct {
}

func (m *MsgWithdrawRewardResponse) Reset()         { *m = MsgWithdrawRewardResponse{} }
func (m *MsgWithdrawRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawRewardResponse) ProtoMessage()    {}
func (*MsgWithdrawRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df57aae78a8732aa, []int{3}
}
func (m *MsgWithdrawRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawRewardResponse.Merge(m, src)
}
func (m *MsgWithdrawRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawRewardResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgWithdrawAllRewards)(nil), "ununifi.ecosystemincentive.MsgWithdrawAllRewards")
	proto.RegisterType((*MsgWithdrawAllRewardsResponse)(nil), "ununifi.ecosystemincentive.MsgWithdrawAllRewardsResponse")
	proto.RegisterType((*MsgWithdrawReward)(nil), "ununifi.ecosystemincentive.MsgWithdrawReward")
	proto.RegisterType((*MsgWithdrawRewardResponse)(nil), "ununifi.ecosystemincentive.MsgWithdrawRewardResponse")
}

func init() {
	proto.RegisterFile("ununifi/ecosystemincentive/tx.proto", fileDescriptor_df57aae78a8732aa)
}

var fileDescriptor_df57aae78a8732aa = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xcd, 0x2b, 0xcd,
	0xcb, 0x4c, 0xcb, 0xd4, 0x4f, 0x4d, 0xce, 0x2f, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0xcd, 0xcc, 0x4b,
	0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x82, 0x2a, 0xd2, 0xc3, 0x54, 0x24, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa6, 0x0f,
	0x62, 0x41, 0x74, 0x48, 0x49, 0x26, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xc7, 0x43, 0x24, 0x20, 0x1c,
	0x88, 0x94, 0x92, 0x13, 0x97, 0xa8, 0x6f, 0x71, 0x7a, 0x78, 0x66, 0x49, 0x46, 0x4a, 0x51, 0x62,
	0xb9, 0x63, 0x4e, 0x4e, 0x50, 0x6a, 0x79, 0x62, 0x51, 0x4a, 0xb1, 0x90, 0x26, 0x17, 0x5b, 0x71,
	0x6a, 0x5e, 0x4a, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xe0, 0xa7, 0x7b, 0xf2,
	0xbc, 0x95, 0x89, 0xb9, 0x39, 0x56, 0x4a, 0x10, 0x71, 0xa5, 0x20, 0xa8, 0x02, 0x25, 0x79, 0x2e,
	0x59, 0xac, 0x66, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0xa5, 0x71, 0x09, 0x22,
	0x29, 0x80, 0xc8, 0x92, 0x60, 0x81, 0x90, 0x1a, 0x17, 0x6b, 0x4a, 0x6a, 0x5e, 0x7e, 0xae, 0x04,
	0x13, 0x58, 0xa5, 0xc0, 0xa7, 0x7b, 0xf2, 0x3c, 0x10, 0x95, 0x60, 0x61, 0xa5, 0x20, 0x88, 0xb4,
	0x92, 0x34, 0x97, 0x24, 0x86, 0x3d, 0x30, 0x47, 0x18, 0x35, 0x31, 0x71, 0x31, 0xfb, 0x16, 0xa7,
	0x0b, 0x35, 0x31, 0x72, 0x09, 0x61, 0xf1, 0xaf, 0xa1, 0x1e, 0xee, 0x60, 0xd5, 0xc3, 0xea, 0x3d,
	0x29, 0x4b, 0x92, 0xb5, 0xc0, 0x1c, 0x23, 0x54, 0xc6, 0xc5, 0x87, 0x16, 0x1c, 0xba, 0x44, 0x1a,
	0x06, 0x51, 0x2e, 0x65, 0x4a, 0x92, 0x72, 0x98, 0xbd, 0x4e, 0xbe, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9c, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0xab, 0x1f, 0x9a, 0x17, 0x9a, 0x97, 0xe9, 0x96, 0xa9, 0x9f, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x5f,
	0x81, 0x35, 0x35, 0x56, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x13, 0x91, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xc7, 0x23, 0xbf, 0xb8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	WithdrawAllRewards(ctx context.Context, in *MsgWithdrawAllRewards, opts ...grpc.CallOption) (*MsgWithdrawAllRewardsResponse, error)
	WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) WithdrawAllRewards(ctx context.Context, in *MsgWithdrawAllRewards, opts ...grpc.CallOption) (*MsgWithdrawAllRewardsResponse, error) {
	out := new(MsgWithdrawAllRewardsResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/WithdrawAllRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardResponse, error) {
	out := new(MsgWithdrawRewardResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/WithdrawReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	WithdrawAllRewards(context.Context, *MsgWithdrawAllRewards) (*MsgWithdrawAllRewardsResponse, error)
	WithdrawReward(context.Context, *MsgWithdrawReward) (*MsgWithdrawRewardResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) WithdrawAllRewards(ctx context.Context, req *MsgWithdrawAllRewards) (*MsgWithdrawAllRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawAllRewards not implemented")
}
func (*UnimplementedMsgServer) WithdrawReward(ctx context.Context, req *MsgWithdrawReward) (*MsgWithdrawRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawReward not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_WithdrawAllRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawAllRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawAllRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ununifi.ecosystemincentive.Msg/WithdrawAllRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawAllRewards(ctx, req.(*MsgWithdrawAllRewards))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ununifi.ecosystemincentive.Msg/WithdrawReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawReward(ctx, req.(*MsgWithdrawReward))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ununifi.ecosystemincentive.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WithdrawAllRewards",
			Handler:    _Msg_WithdrawAllRewards_Handler,
		},
		{
			MethodName: "WithdrawReward",
			Handler:    _Msg_WithdrawReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ununifi/ecosystemincentive/tx.proto",
}

func (m *MsgWithdrawAllRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawAllRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawAllRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawAllRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawAllRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawAllRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgWithdrawAllRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawAllRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgWithdrawAllRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawAllRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawAllRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawAllRewardsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawAllRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawAllRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawRewardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
