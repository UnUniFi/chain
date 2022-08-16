// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: yieldaggregator/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the yieldaggregator module's genesis state.
type GenesisState struct {
	Params                  Params                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	AssetManagementAccounts []AssetManagementAccount `protobuf:"bytes,2,rep,name=asset_management_accounts,json=assetManagementAccounts,proto3" json:"asset_management_accounts"`
	AssetManagementTargets  []AssetManagementTarget  `protobuf:"bytes,3,rep,name=asset_management_targets,json=assetManagementTargets,proto3" json:"asset_management_targets"`
	FarmingOrders           []FarmingOrder           `protobuf:"bytes,4,rep,name=farming_orders,json=farmingOrders,proto3" json:"farming_orders"`
	FarmingUnits            []FarmingUnit            `protobuf:"bytes,5,rep,name=farming_units,json=farmingUnits,proto3" json:"farming_units"`
	UserInfos               []UserInfo               `protobuf:"bytes,6,rep,name=user_infos,json=userInfos,proto3" json:"user_infos"`
	LastFarmingUnitId       uint64                   `protobuf:"varint,7,opt,name=last_farming_unit_id,json=lastFarmingUnitId,proto3" json:"last_farming_unit_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4a7772696467f75, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAssetManagementAccounts() []AssetManagementAccount {
	if m != nil {
		return m.AssetManagementAccounts
	}
	return nil
}

func (m *GenesisState) GetAssetManagementTargets() []AssetManagementTarget {
	if m != nil {
		return m.AssetManagementTargets
	}
	return nil
}

func (m *GenesisState) GetFarmingOrders() []FarmingOrder {
	if m != nil {
		return m.FarmingOrders
	}
	return nil
}

func (m *GenesisState) GetFarmingUnits() []FarmingUnit {
	if m != nil {
		return m.FarmingUnits
	}
	return nil
}

func (m *GenesisState) GetUserInfos() []UserInfo {
	if m != nil {
		return m.UserInfos
	}
	return nil
}

func (m *GenesisState) GetLastFarmingUnitId() uint64 {
	if m != nil {
		return m.LastFarmingUnitId
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ununifi.chain.yieldaggregator.GenesisState")
}

func init() { proto.RegisterFile("yieldaggregator/genesis.proto", fileDescriptor_e4a7772696467f75) }

var fileDescriptor_e4a7772696467f75 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x13, 0x6f, 0xac, 0x38, 0xf7, 0x2a, 0x38, 0x5c, 0x74, 0xac, 0x36, 0x16, 0x41, 0x2c,
	0x0a, 0x89, 0x54, 0x7d, 0x80, 0x56, 0xa8, 0x54, 0x14, 0xa5, 0x1a, 0x10, 0x37, 0x61, 0x9a, 0x4c,
	0xa6, 0x03, 0xcd, 0x4c, 0x99, 0x33, 0x41, 0xfb, 0x16, 0x3e, 0x56, 0x97, 0x5d, 0xba, 0x12, 0x69,
	0xdf, 0xc0, 0x27, 0x90, 0x4c, 0x52, 0x2d, 0xb1, 0x58, 0xef, 0x2e, 0x39, 0xe7, 0x7c, 0xff, 0xf7,
	0x2f, 0x06, 0x75, 0x96, 0x82, 0xcd, 0x53, 0xca, 0xb9, 0x66, 0x9c, 0x1a, 0xa5, 0x43, 0xce, 0x24,
	0x03, 0x01, 0xc1, 0x42, 0x2b, 0xa3, 0x70, 0xa7, 0x90, 0x85, 0x14, 0x99, 0x08, 0x92, 0x19, 0x15,
	0x32, 0x68, 0x1c, 0xb7, 0xcf, 0xb9, 0xe2, 0xca, 0x5e, 0x86, 0xe5, 0x57, 0x05, 0xb5, 0xef, 0x36,
	0x33, 0x17, 0x54, 0xd3, 0xbc, 0x8e, 0x6c, 0xdf, 0x69, 0x6e, 0x29, 0x00, 0x33, 0xd5, 0xf2, 0xfe,
	0x4f, 0x0f, 0x9d, 0xbd, 0xac, 0x1a, 0xbc, 0x37, 0xd4, 0x30, 0xfc, 0x02, 0xb5, 0x2a, 0x9a, 0xb8,
	0x5d, 0xb7, 0x77, 0xda, 0x7f, 0x10, 0xfc, 0xb3, 0x51, 0xf0, 0xce, 0x1e, 0x0f, 0xbd, 0xd5, 0xf7,
	0x7b, 0xce, 0xa4, 0x46, 0xf1, 0x67, 0x74, 0xdb, 0x4a, 0xe2, 0x9c, 0x4a, 0xca, 0x59, 0xce, 0xa4,
	0x89, 0x69, 0x92, 0xa8, 0x42, 0x1a, 0x20, 0x97, 0xba, 0x27, 0xbd, 0xd3, 0xfe, 0xf3, 0x23, 0xb9,
	0x83, 0x92, 0x7f, 0xf3, 0x1b, 0x1f, 0x54, 0x74, 0xed, 0xb9, 0x45, 0x0f, 0x6e, 0x01, 0x1b, 0x44,
	0xfe, 0x12, 0x1b, 0xaa, 0x39, 0x33, 0x40, 0x4e, 0xac, 0xf7, 0xd9, 0xc5, 0xbc, 0x1f, 0x2c, 0x5c,
	0x6b, 0x6f, 0xd2, 0x43, 0x4b, 0xc0, 0x1f, 0xd1, 0xf5, 0x8c, 0xea, 0x5c, 0x48, 0x1e, 0x2b, 0x9d,
	0x32, 0x0d, 0xc4, 0xb3, 0xae, 0xc7, 0x47, 0x5c, 0xa3, 0x0a, 0x7a, 0x5b, 0x32, 0xb5, 0xe2, 0x5a,
	0xb6, 0x37, 0x03, 0x1c, 0xa1, 0xdd, 0x20, 0x2e, 0xa4, 0x30, 0x40, 0x2e, 0xdb, 0xe0, 0x47, 0xff,
	0x17, 0x1c, 0x49, 0xb1, 0xab, 0x7e, 0x96, 0xfd, 0x19, 0x01, 0x7e, 0x8d, 0x50, 0x01, 0x4c, 0xc7,
	0x42, 0x66, 0x0a, 0x48, 0xcb, 0x66, 0x3e, 0x3c, 0x92, 0x19, 0x01, 0xd3, 0x63, 0x99, 0xa9, 0x3a,
	0xf0, 0x6a, 0x51, 0xff, 0x03, 0x0e, 0xd1, 0xf9, 0x9c, 0x82, 0x89, 0xf7, 0x9b, 0xc6, 0x22, 0x25,
	0x57, 0xba, 0x6e, 0xcf, 0x9b, 0xdc, 0x28, 0x77, 0x7b, 0x85, 0xc6, 0xe9, 0xf0, 0xd5, 0x6a, 0xe3,
	0xbb, 0xeb, 0x8d, 0xef, 0xfe, 0xd8, 0xf8, 0xee, 0xd7, 0xad, 0xef, 0xac, 0xb7, 0xbe, 0xf3, 0x6d,
	0xeb, 0x3b, 0x9f, 0x9e, 0x70, 0x61, 0x66, 0xc5, 0x34, 0x48, 0x54, 0x1e, 0x46, 0x32, 0x92, 0x62,
	0x24, 0x42, 0x5b, 0x27, 0xfc, 0x12, 0x36, 0x9f, 0xb1, 0x59, 0x2e, 0x18, 0x4c, 0x5b, 0xf6, 0x1d,
	0x3f, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x05, 0x12, 0x5f, 0x97, 0x58, 0x03, 0x00, 0x00,
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
	if m.LastFarmingUnitId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastFarmingUnitId))
		i--
		dAtA[i] = 0x38
	}
	if len(m.UserInfos) > 0 {
		for iNdEx := len(m.UserInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.FarmingUnits) > 0 {
		for iNdEx := len(m.FarmingUnits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FarmingUnits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.FarmingOrders) > 0 {
		for iNdEx := len(m.FarmingOrders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FarmingOrders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AssetManagementTargets) > 0 {
		for iNdEx := len(m.AssetManagementTargets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetManagementTargets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AssetManagementAccounts) > 0 {
		for iNdEx := len(m.AssetManagementAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetManagementAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.AssetManagementAccounts) > 0 {
		for _, e := range m.AssetManagementAccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AssetManagementTargets) > 0 {
		for _, e := range m.AssetManagementTargets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FarmingOrders) > 0 {
		for _, e := range m.FarmingOrders {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FarmingUnits) > 0 {
		for _, e := range m.FarmingUnits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserInfos) > 0 {
		for _, e := range m.UserInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastFarmingUnitId != 0 {
		n += 1 + sovGenesis(uint64(m.LastFarmingUnitId))
	}
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetManagementAccounts", wireType)
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
			m.AssetManagementAccounts = append(m.AssetManagementAccounts, AssetManagementAccount{})
			if err := m.AssetManagementAccounts[len(m.AssetManagementAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetManagementTargets", wireType)
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
			m.AssetManagementTargets = append(m.AssetManagementTargets, AssetManagementTarget{})
			if err := m.AssetManagementTargets[len(m.AssetManagementTargets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmingOrders", wireType)
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
			m.FarmingOrders = append(m.FarmingOrders, FarmingOrder{})
			if err := m.FarmingOrders[len(m.FarmingOrders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmingUnits", wireType)
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
			m.FarmingUnits = append(m.FarmingUnits, FarmingUnit{})
			if err := m.FarmingUnits[len(m.FarmingUnits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfos", wireType)
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
			m.UserInfos = append(m.UserInfos, UserInfo{})
			if err := m.UserInfos[len(m.UserInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastFarmingUnitId", wireType)
			}
			m.LastFarmingUnitId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastFarmingUnitId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
