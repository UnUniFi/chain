// Code generated by MockGen. DO NOT EDIT.
// Source: x/nftmint/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetSequence mocks base method.
func (m *MockAccountKeeper) GetSequence(ctx types.Context, addr types.AccAddress) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSequence", ctx, addr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSequence indicates an expected call of GetSequence.
func (mr *MockAccountKeeperMockRecorder) GetSequence(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSequence", reflect.TypeOf((*MockAccountKeeper)(nil).GetSequence), ctx, addr)
}

// MockNftKeeper is a mock of NftKeeper interface.
type MockNftKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockNftKeeperMockRecorder
}

// MockNftKeeperMockRecorder is the mock recorder for MockNftKeeper.
type MockNftKeeperMockRecorder struct {
	mock *MockNftKeeper
}

// NewMockNftKeeper creates a new mock instance.
func NewMockNftKeeper(ctrl *gomock.Controller) *MockNftKeeper {
	mock := &MockNftKeeper{ctrl: ctrl}
	mock.recorder = &MockNftKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNftKeeper) EXPECT() *MockNftKeeperMockRecorder {
	return m.recorder
}

// Burn mocks base method.
func (m *MockNftKeeper) Burn(ctx types.Context, classID, nftID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Burn", ctx, classID, nftID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Burn indicates an expected call of Burn.
func (mr *MockNftKeeperMockRecorder) Burn(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Burn", reflect.TypeOf((*MockNftKeeper)(nil).Burn), ctx, classID, nftID)
}

// GetBalance mocks base method.
func (m *MockNftKeeper) GetBalance(ctx types.Context, classID string, owner types.AccAddress) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, classID, owner)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockNftKeeperMockRecorder) GetBalance(ctx, classID, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockNftKeeper)(nil).GetBalance), ctx, classID, owner)
}

// GetClass mocks base method.
func (m *MockNftKeeper) GetClass(ctx types.Context, classID string) (nft.Class, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClass", ctx, classID)
	ret0, _ := ret[0].(nft.Class)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetClass indicates an expected call of GetClass.
func (mr *MockNftKeeperMockRecorder) GetClass(ctx, classID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClass", reflect.TypeOf((*MockNftKeeper)(nil).GetClass), ctx, classID)
}

// GetClasses mocks base method.
func (m *MockNftKeeper) GetClasses(ctx types.Context) []*nft.Class {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClasses", ctx)
	ret0, _ := ret[0].([]*nft.Class)
	return ret0
}

// GetClasses indicates an expected call of GetClasses.
func (mr *MockNftKeeperMockRecorder) GetClasses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClasses", reflect.TypeOf((*MockNftKeeper)(nil).GetClasses), ctx)
}

// GetNFT mocks base method.
func (m *MockNftKeeper) GetNFT(ctx types.Context, classID, nftID string) (nft.NFT, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFT", ctx, classID, nftID)
	ret0, _ := ret[0].(nft.NFT)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNFT indicates an expected call of GetNFT.
func (mr *MockNftKeeperMockRecorder) GetNFT(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFT", reflect.TypeOf((*MockNftKeeper)(nil).GetNFT), ctx, classID, nftID)
}

// GetNFTsOfClass mocks base method.
func (m *MockNftKeeper) GetNFTsOfClass(ctx types.Context, classID string) []nft.NFT {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTsOfClass", ctx, classID)
	ret0, _ := ret[0].([]nft.NFT)
	return ret0
}

// GetNFTsOfClass indicates an expected call of GetNFTsOfClass.
func (mr *MockNftKeeperMockRecorder) GetNFTsOfClass(ctx, classID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTsOfClass", reflect.TypeOf((*MockNftKeeper)(nil).GetNFTsOfClass), ctx, classID)
}

// GetNFTsOfClassByOwner mocks base method.
func (m *MockNftKeeper) GetNFTsOfClassByOwner(ctx types.Context, classID string, owner types.AccAddress) []nft.NFT {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTsOfClassByOwner", ctx, classID, owner)
	ret0, _ := ret[0].([]nft.NFT)
	return ret0
}

// GetNFTsOfClassByOwner indicates an expected call of GetNFTsOfClassByOwner.
func (mr *MockNftKeeperMockRecorder) GetNFTsOfClassByOwner(ctx, classID, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTsOfClassByOwner", reflect.TypeOf((*MockNftKeeper)(nil).GetNFTsOfClassByOwner), ctx, classID, owner)
}

// GetOwner mocks base method.
func (m *MockNftKeeper) GetOwner(ctx types.Context, classID, nftID string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwner", ctx, classID, nftID)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetOwner indicates an expected call of GetOwner.
func (mr *MockNftKeeperMockRecorder) GetOwner(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwner", reflect.TypeOf((*MockNftKeeper)(nil).GetOwner), ctx, classID, nftID)
}

// GetTotalSupply mocks base method.
func (m *MockNftKeeper) GetTotalSupply(ctx types.Context, classID string) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalSupply", ctx, classID)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetTotalSupply indicates an expected call of GetTotalSupply.
func (mr *MockNftKeeperMockRecorder) GetTotalSupply(ctx, classID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalSupply", reflect.TypeOf((*MockNftKeeper)(nil).GetTotalSupply), ctx, classID)
}

// HasClass mocks base method.
func (m *MockNftKeeper) HasClass(ctx types.Context, classId string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasClass", ctx, classId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasClass indicates an expected call of HasClass.
func (mr *MockNftKeeperMockRecorder) HasClass(ctx, classId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasClass", reflect.TypeOf((*MockNftKeeper)(nil).HasClass), ctx, classId)
}

// HasNFT mocks base method.
func (m *MockNftKeeper) HasNFT(ctx types.Context, classID, nftID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNFT", ctx, classID, nftID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNFT indicates an expected call of HasNFT.
func (mr *MockNftKeeperMockRecorder) HasNFT(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNFT", reflect.TypeOf((*MockNftKeeper)(nil).HasNFT), ctx, classID, nftID)
}

// Mint mocks base method.
func (m *MockNftKeeper) Mint(ctx types.Context, token nft.NFT, receiver types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mint", ctx, token, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mint indicates an expected call of Mint.
func (mr *MockNftKeeperMockRecorder) Mint(ctx, token, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mint", reflect.TypeOf((*MockNftKeeper)(nil).Mint), ctx, token, receiver)
}

// SaveClass mocks base method.
func (m *MockNftKeeper) SaveClass(ctx types.Context, class nft.Class) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveClass", ctx, class)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveClass indicates an expected call of SaveClass.
func (mr *MockNftKeeperMockRecorder) SaveClass(ctx, class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveClass", reflect.TypeOf((*MockNftKeeper)(nil).SaveClass), ctx, class)
}

// Transfer mocks base method.
func (m *MockNftKeeper) Transfer(ctx types.Context, classID, nftID string, receiver types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, classID, nftID, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transfer indicates an expected call of Transfer.
func (mr *MockNftKeeperMockRecorder) Transfer(ctx, classID, nftID, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockNftKeeper)(nil).Transfer), ctx, classID, nftID, receiver)
}

// Update mocks base method.
func (m *MockNftKeeper) Update(ctx types.Context, token nft.NFT) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNftKeeperMockRecorder) Update(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNftKeeper)(nil).Update), ctx, token)
}