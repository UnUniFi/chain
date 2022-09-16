// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ecosystem-incentive/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_UnUniFi_chain_types "github.com/UnUniFi/chain/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgIncentiveRegister struct {
	IncentiveId string                                            `protobuf:"bytes,1,opt,name=incentive_id,json=incentiveId,proto3" json:"incentive_id,omitempty"`
	Subjects    []github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,2,rep,name=subjects,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"subjects" yaml:"sender"`
	Weights     []string                                          `protobuf:"bytes,3,rep,name=weights,proto3" json:"weights,omitempty"`
}

func (m *MsgIncentiveRegister) Reset()         { *m = MsgIncentiveRegister{} }
func (m *MsgIncentiveRegister) String() string { return proto.CompactTextString(m) }
func (*MsgIncentiveRegister) ProtoMessage()    {}
func (*MsgIncentiveRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{0}
}
func (m *MsgIncentiveRegister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncentiveRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncentiveRegister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncentiveRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncentiveRegister.Merge(m, src)
}
func (m *MsgIncentiveRegister) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncentiveRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncentiveRegister.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncentiveRegister proto.InternalMessageInfo

func (m *MsgIncentiveRegister) GetIncentiveId() string {
	if m != nil {
		return m.IncentiveId
	}
	return ""
}

func (m *MsgIncentiveRegister) GetWeights() []string {
	if m != nil {
		return m.Weights
	}
	return nil
}

type MsgIncentiveRegisterResponse struct {
}

func (m *MsgIncentiveRegisterResponse) Reset()         { *m = MsgIncentiveRegisterResponse{} }
func (m *MsgIncentiveRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIncentiveRegisterResponse) ProtoMessage()    {}
func (*MsgIncentiveRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{1}
}
func (m *MsgIncentiveRegisterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncentiveRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncentiveRegisterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncentiveRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncentiveRegisterResponse.Merge(m, src)
}
func (m *MsgIncentiveRegisterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncentiveRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncentiveRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncentiveRegisterResponse proto.InternalMessageInfo

type MsgWithdrawAllRewards struct {
	Sender github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,1,opt,name=sender,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"sender" yaml:"sender"`
}

func (m *MsgWithdrawAllRewards) Reset()         { *m = MsgWithdrawAllRewards{} }
func (m *MsgWithdrawAllRewards) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawAllRewards) ProtoMessage()    {}
func (*MsgWithdrawAllRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{2}
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

type MsgWithdrawAllRewardsResponse struct {
}

func (m *MsgWithdrawAllRewardsResponse) Reset()         { *m = MsgWithdrawAllRewardsResponse{} }
func (m *MsgWithdrawAllRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawAllRewardsResponse) ProtoMessage()    {}
func (*MsgWithdrawAllRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{3}
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
	Sender github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,1,opt,name=sender,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"sender" yaml:"sender"`
	Denom  string                                          `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *MsgWithdrawReward) Reset()         { *m = MsgWithdrawReward{} }
func (m *MsgWithdrawReward) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawReward) ProtoMessage()    {}
func (*MsgWithdrawReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{4}
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

func (m *MsgWithdrawReward) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type MsgWithdrawRewardsResponse struct {
}

func (m *MsgWithdrawRewardsResponse) Reset()         { *m = MsgWithdrawRewardsResponse{} }
func (m *MsgWithdrawRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawRewardsResponse) ProtoMessage()    {}
func (*MsgWithdrawRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f4f2600114c5509, []int{5}
}
func (m *MsgWithdrawRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawRewardsResponse.Merge(m, src)
}
func (m *MsgWithdrawRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawRewardsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgIncentiveRegister)(nil), "ununifi.ecosystemincentive.MsgIncentiveRegister")
	proto.RegisterType((*MsgIncentiveRegisterResponse)(nil), "ununifi.ecosystemincentive.MsgIncentiveRegisterResponse")
	proto.RegisterType((*MsgWithdrawAllRewards)(nil), "ununifi.ecosystemincentive.MsgWithdrawAllRewards")
	proto.RegisterType((*MsgWithdrawAllRewardsResponse)(nil), "ununifi.ecosystemincentive.MsgWithdrawAllRewardsResponse")
	proto.RegisterType((*MsgWithdrawReward)(nil), "ununifi.ecosystemincentive.MsgWithdrawReward")
	proto.RegisterType((*MsgWithdrawRewardsResponse)(nil), "ununifi.ecosystemincentive.MsgWithdrawRewardsResponse")
}

func init() { proto.RegisterFile("ecosystem-incentive/tx.proto", fileDescriptor_5f4f2600114c5509) }

var fileDescriptor_5f4f2600114c5509 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x3b, 0x5b, 0x5c, 0xdd, 0xd7, 0x3f, 0xb0, 0x43, 0x85, 0x10, 0x6a, 0xba, 0xe6, 0xb4,
	0x97, 0x4d, 0xfc, 0x87, 0xa8, 0xb7, 0xee, 0x41, 0xdc, 0x43, 0x3d, 0x44, 0x16, 0xc1, 0x8b, 0xb6,
	0x99, 0xd7, 0xe9, 0x48, 0x3b, 0x53, 0xf2, 0x4e, 0xcc, 0xe6, 0x24, 0x88, 0x37, 0x2f, 0x7e, 0x1d,
	0xbf, 0xc1, 0x1e, 0xf7, 0x28, 0x1e, 0x8a, 0xb4, 0xdf, 0xc0, 0x4f, 0x20, 0x6d, 0xec, 0x20, 0x36,
	0x8b, 0x16, 0xf4, 0x96, 0x49, 0x9e, 0xe7, 0x7d, 0x7f, 0xe1, 0x79, 0x18, 0x68, 0x63, 0x6a, 0xa8,
	0x24, 0x8b, 0xe3, 0x03, 0xa5, 0x53, 0xd4, 0x56, 0xbd, 0xc5, 0xd8, 0x9e, 0x44, 0x93, 0xcc, 0x58,
	0xc3, 0xfd, 0x5c, 0xe7, 0x5a, 0xbd, 0x56, 0x91, 0x53, 0x39, 0x91, 0xdf, 0x92, 0x46, 0x9a, 0xa5,
	0x2c, 0x5e, 0x3c, 0x55, 0x8e, 0xf0, 0x33, 0x83, 0x56, 0x8f, 0xe4, 0xd1, 0x4a, 0x96, 0xa0, 0x54,
	0x64, 0x31, 0xe3, 0x37, 0xe1, 0x8a, 0xf3, 0xbe, 0x54, 0xc2, 0x63, 0x7b, 0x6c, 0x7f, 0x27, 0xb9,
	0xec, 0xde, 0x1d, 0x09, 0x2e, 0xe0, 0x12, 0xe5, 0x83, 0x37, 0x98, 0x5a, 0xf2, 0xb6, 0xf6, 0x9a,
	0xfb, 0x3b, 0x87, 0x4f, 0x4e, 0xa7, 0x9d, 0xc6, 0xd7, 0x69, 0x27, 0x96, 0xca, 0x0e, 0xf3, 0x41,
	0x94, 0x9a, 0x71, 0x7c, 0xac, 0x8f, 0xb5, 0x7a, 0xac, 0xe2, 0x74, 0xd8, 0x57, 0x3a, 0xb6, 0xe5,
	0x04, 0x29, 0x7a, 0x66, 0x33, 0xa5, 0x65, 0x37, 0x4d, 0xbb, 0x42, 0x64, 0x48, 0xf4, 0x7d, 0xda,
	0xb9, 0x5a, 0xf6, 0xc7, 0xa3, 0x47, 0x21, 0xa1, 0x16, 0x98, 0x85, 0x89, 0x9b, 0xcc, 0x3d, 0xb8,
	0x58, 0xa0, 0x92, 0x43, 0x4b, 0x5e, 0x73, 0xb1, 0x24, 0x59, 0x1d, 0xc3, 0x00, 0xda, 0x75, 0xe8,
	0x09, 0xd2, 0xc4, 0x68, 0xc2, 0xb0, 0x84, 0xeb, 0x3d, 0x92, 0xcf, 0x95, 0x1d, 0x8a, 0xac, 0x5f,
	0x74, 0x47, 0xa3, 0x04, 0x8b, 0x7e, 0x26, 0x88, 0xbf, 0x82, 0xed, 0x6a, 0x4f, 0xf5, 0x57, 0xff,
	0x10, 0xfb, 0xe7, 0xdc, 0xb0, 0x03, 0x37, 0x6a, 0x57, 0x3b, 0xb6, 0x8f, 0x0c, 0x76, 0x7f, 0x51,
	0x54, 0x9f, 0xff, 0x3f, 0x18, 0x6f, 0xc1, 0x05, 0x81, 0xda, 0x8c, 0xbd, 0xad, 0x65, 0x9e, 0xd5,
	0x21, 0x6c, 0x83, 0xbf, 0x06, 0xe3, 0x58, 0xef, 0x7c, 0x68, 0x42, 0xb3, 0x47, 0x92, 0xbf, 0x83,
	0xdd, 0xf5, 0x9e, 0xdc, 0x8a, 0xce, 0xef, 0x5c, 0x54, 0x17, 0x8f, 0xff, 0x60, 0x53, 0xc7, 0x0a,
	0x84, 0xbf, 0x67, 0xc0, 0x6b, 0xe2, 0xbc, 0xfd, 0x87, 0x81, 0xeb, 0x16, 0xff, 0xe1, 0xc6, 0x16,
	0x07, 0x51, 0xc0, 0xb5, 0xdf, 0x52, 0x3b, 0xf8, 0xcb, 0x61, 0x95, 0xdc, 0xbf, 0xbf, 0x91, 0xdc,
	0x2d, 0x3e, 0x7c, 0x7a, 0x3a, 0x0b, 0xd8, 0xd9, 0x2c, 0x60, 0xdf, 0x66, 0x01, 0xfb, 0x34, 0x0f,
	0x1a, 0x67, 0xf3, 0xa0, 0xf1, 0x65, 0x1e, 0x34, 0x5e, 0xdc, 0x3b, 0xb7, 0x1e, 0x27, 0x71, 0xed,
	0x7d, 0xb1, 0x28, 0xcd, 0x60, 0x7b, 0x79, 0x03, 0xdc, 0xfd, 0x11, 0x00, 0x00, 0xff, 0xff, 0x77,
	0x9d, 0x52, 0x49, 0x53, 0x04, 0x00, 0x00,
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
	IncentiveRegister(ctx context.Context, in *MsgIncentiveRegister, opts ...grpc.CallOption) (*MsgIncentiveRegisterResponse, error)
	WithdrawAllRewards(ctx context.Context, in *MsgWithdrawAllRewards, opts ...grpc.CallOption) (*MsgWithdrawAllRewardsResponse, error)
	WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IncentiveRegister(ctx context.Context, in *MsgIncentiveRegister, opts ...grpc.CallOption) (*MsgIncentiveRegisterResponse, error) {
	out := new(MsgIncentiveRegisterResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/IncentiveRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawAllRewards(ctx context.Context, in *MsgWithdrawAllRewards, opts ...grpc.CallOption) (*MsgWithdrawAllRewardsResponse, error) {
	out := new(MsgWithdrawAllRewardsResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/WithdrawAllRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawReward(ctx context.Context, in *MsgWithdrawReward, opts ...grpc.CallOption) (*MsgWithdrawRewardsResponse, error) {
	out := new(MsgWithdrawRewardsResponse)
	err := c.cc.Invoke(ctx, "/ununifi.ecosystemincentive.Msg/WithdrawReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	IncentiveRegister(context.Context, *MsgIncentiveRegister) (*MsgIncentiveRegisterResponse, error)
	WithdrawAllRewards(context.Context, *MsgWithdrawAllRewards) (*MsgWithdrawAllRewardsResponse, error)
	WithdrawReward(context.Context, *MsgWithdrawReward) (*MsgWithdrawRewardsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) IncentiveRegister(ctx context.Context, req *MsgIncentiveRegister) (*MsgIncentiveRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncentiveRegister not implemented")
}
func (*UnimplementedMsgServer) WithdrawAllRewards(ctx context.Context, req *MsgWithdrawAllRewards) (*MsgWithdrawAllRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawAllRewards not implemented")
}
func (*UnimplementedMsgServer) WithdrawReward(ctx context.Context, req *MsgWithdrawReward) (*MsgWithdrawRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawReward not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_IncentiveRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIncentiveRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IncentiveRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ununifi.ecosystemincentive.Msg/IncentiveRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IncentiveRegister(ctx, req.(*MsgIncentiveRegister))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "IncentiveRegister",
			Handler:    _Msg_IncentiveRegister_Handler,
		},
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
	Metadata: "ecosystem-incentive/tx.proto",
}

func (m *MsgIncentiveRegister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncentiveRegister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncentiveRegister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Weights) > 0 {
		for iNdEx := len(m.Weights) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Weights[iNdEx])
			copy(dAtA[i:], m.Weights[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Weights[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Subjects) > 0 {
		for iNdEx := len(m.Subjects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Subjects[iNdEx].Size()
				i -= size
				if _, err := m.Subjects[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IncentiveId) > 0 {
		i -= len(m.IncentiveId)
		copy(dAtA[i:], m.IncentiveId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IncentiveId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgIncentiveRegisterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncentiveRegisterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncentiveRegisterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
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
	{
		size := m.Sender.Size()
		i -= size
		if _, err := m.Sender.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	{
		size := m.Sender.Size()
		i -= size
		if _, err := m.Sender.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgIncentiveRegister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IncentiveId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Subjects) > 0 {
		for _, e := range m.Subjects {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Weights) > 0 {
		for _, s := range m.Weights {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgIncentiveRegisterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawAllRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Sender.Size()
	n += 1 + l + sovTx(uint64(l))
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
	l = m.Sender.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawRewardsResponse) Size() (n int) {
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
func (m *MsgIncentiveRegister) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIncentiveRegister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncentiveRegister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentiveId", wireType)
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
			m.IncentiveId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subjects", wireType)
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
			var v github_com_UnUniFi_chain_types.StringAccAddress
			m.Subjects = append(m.Subjects, v)
			if err := m.Subjects[len(m.Subjects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weights", wireType)
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
			m.Weights = append(m.Weights, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgIncentiveRegisterResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIncentiveRegisterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncentiveRegisterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgWithdrawRewardsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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