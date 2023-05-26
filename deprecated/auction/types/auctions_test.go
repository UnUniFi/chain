package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	ununifitypes "github.com/UnUniFi/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TestInitiatorModuleName = "liquidator"
	TestLotDenom            = "jpu"
	TestLotAmount           = 100
	TestBidDenom            = "guu"
	TestBidAmount           = 20
	TestDebtDenom           = "debt"
	TestDebtAmount1         = 20
	TestDebtAmount2         = 15
	TestExtraEndTime        = 10000
	TestAuctionID           = 9999123
	testAccAddress1         = "ununifi1hsh64ca3q68tnvqyt8mn5knfcm7jyesr5prcns"
	testAccAddress2         = "ununifi1qq6xpcc4xhzxgkw4m9n8lew8sfmt3sn2r6usnl"
)

func init() {
	sdk.GetConfig().SetBech32PrefixForAccount("ununifi", "ununifi"+sdk.PrefixPublic)
}

func d(amount string) sdk.Dec               { return sdk.MustNewDecFromStr(amount) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func i(n int64) sdk.Int                     { return sdk.NewInt(n) }
func is(ns ...int64) (is []sdk.Int) {
	for _, n := range ns {
		is = append(is, sdk.NewInt(n))
	}
	return
}

func TestNewWeightedAddresses(t *testing.T) {
	addr1, err := sdk.AccAddressFromBech32(testAccAddress1)
	require.NoError(t, err)

	addr2, err := sdk.AccAddressFromBech32(testAccAddress2)
	require.NoError(t, err)

	tests := []struct {
		name      string
		addresses []sdk.AccAddress
		weights   []sdk.Int
		expPass   bool
	}{
		{
			"normal",
			[]sdk.AccAddress{addr1, addr2},
			[]sdk.Int{sdk.NewInt(6), sdk.NewInt(8)},
			true,
		},
		{
			"empty address",
			[]sdk.AccAddress{nil, nil},
			[]sdk.Int{sdk.NewInt(6), sdk.NewInt(8)},
			false,
		},
		{
			"mismatched",
			[]sdk.AccAddress{addr1, addr2},
			[]sdk.Int{sdk.NewInt(6)},
			false,
		},
		{
			"negative weight",
			[]sdk.AccAddress{addr1, addr2},
			is(6, -8),
			false,
		},
		{
			"zero weight",
			[]sdk.AccAddress{addr1, addr2},
			is(0, 0),
			false,
		},
	}

	// Run NewWeightedAdresses tests
	for _, tc := range tests {
		// Attempt to instantiate new WeightedAddresses
		weightedAddresses, err := NewWeightedAddresses(tc.addresses, tc.weights)

		if tc.expPass {
			require.NoError(t, err)
			require.Equal(t, tc.addresses, weightedAddresses.Addresses())
			require.Equal(t, tc.weights, weightedAddresses.Weights())
		} else {
			require.Error(t, err)
		}
	}
}

func TestBaseAuctionValidate(t *testing.T) {
	addr1, err := sdk.AccAddressFromBech32(testAccAddress1)
	require.NoError(t, err)

	now := time.Now()

	tests := []struct {
		msg     string
		auction BaseAuction
		expPass bool
	}{
		{
			"valid auction",
			BaseAuction{
				Id:              1,
				Initiator:       testAccAddress1,
				Lot:             c("guu", 1),
				Bidder:          ununifitypes.StringAccAddress(addr1),
				Bid:             c("guu", 1),
				EndTime:         now,
				MaxEndTime:      now,
				HasReceivedBids: true,
			},
			true,
		},
		{
			"blank initiator",
			BaseAuction{
				Id:        1,
				Initiator: "",
			},
			false,
		},
		{
			"invalid lot",
			BaseAuction{
				Id:        1,
				Initiator: testAccAddress1,
				Lot:       sdk.Coin{Denom: "%DENOM", Amount: sdk.NewInt(1)},
			},
			false,
		},
		{
			"empty bidder",
			BaseAuction{
				Id:        1,
				Initiator: testAccAddress1,
				Lot:       c("guu", 1),
				Bidder:    ununifitypes.StringAccAddress{},
			},
			false,
		},
		{
			"invalid bidder",
			BaseAuction{
				Id:        1,
				Initiator: testAccAddress1,
				Lot:       c("guu", 1),
				Bidder:    ununifitypes.StringAccAddress(addr1[:10]),
			},
			false,
		},
		{
			"invalid bid",
			BaseAuction{
				Id:        1,
				Initiator: testAccAddress1,
				Lot:       c("guu", 1),
				Bidder:    ununifitypes.StringAccAddress(addr1),
				Bid:       sdk.Coin{Denom: "%DENOM", Amount: sdk.NewInt(1)},
			},
			false,
		},
		{
			"invalid end time",
			BaseAuction{
				Id:        1,
				Initiator: testAccAddress1,
				Lot:       c("guu", 1),
				Bidder:    ununifitypes.StringAccAddress(addr1),
				Bid:       c("guu", 1),
				EndTime:   time.Unix(0, 0),
			},
			false,
		},
		{
			"max end time > endtime",
			BaseAuction{
				Id:         1,
				Initiator:  testAccAddress1,
				Lot:        c("guu", 1),
				Bidder:     ununifitypes.StringAccAddress(addr1),
				Bid:        c("guu", 1),
				EndTime:    now.Add(time.Minute),
				MaxEndTime: now,
			},
			false,
		},
	}

	for _, tc := range tests {

		err := tc.auction.Validate()

		if tc.expPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func TestDebtAuctionValidate(t *testing.T) {
	addr1, err := sdk.AccAddressFromBech32(testAccAddress1)
	require.NoError(t, err)

	now := time.Now()

	tests := []struct {
		msg     string
		auction DebtAuction
		expPass bool
	}{
		{
			"valid auction",
			DebtAuction{
				BaseAuction: BaseAuction{
					Id:              1,
					Initiator:       testAccAddress1,
					Lot:             c("guu", 1),
					Bidder:          ununifitypes.StringAccAddress(addr1),
					Bid:             c("guu", 1),
					EndTime:         now,
					MaxEndTime:      now,
					HasReceivedBids: true,
				},
				CorrespondingDebt: c("guu", 1),
			},
			true,
		},
		{
			"invalid corresponding debt",
			DebtAuction{
				BaseAuction: BaseAuction{
					Id:              1,
					Initiator:       testAccAddress1,
					Lot:             c("guu", 1),
					Bidder:          ununifitypes.StringAccAddress(addr1),
					Bid:             c("guu", 1),
					EndTime:         now,
					MaxEndTime:      now,
					HasReceivedBids: true,
				},
				CorrespondingDebt: sdk.Coin{Denom: "%DENOM", Amount: sdk.NewInt(1)},
			},
			false,
		},
	}

	for _, tc := range tests {

		err := tc.auction.Validate()

		if tc.expPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func TestCollateralAuctionValidate(t *testing.T) {
	addr1, err := sdk.AccAddressFromBech32(testAccAddress1)
	require.NoError(t, err)

	now := time.Now()

	tests := []struct {
		msg     string
		auction CollateralAuction
		expPass bool
	}{
		{
			"valid auction",
			CollateralAuction{
				BaseAuction: BaseAuction{
					Id:              1,
					Initiator:       testAccAddress1,
					Lot:             c("guu", 1),
					Bidder:          ununifitypes.StringAccAddress(addr1),
					Bid:             c("guu", 1),
					EndTime:         now,
					MaxEndTime:      now,
					HasReceivedBids: true,
				},
				CorrespondingDebt: c("guu", 1),
				MaxBid:            c("guu", 1),
				LotReturns: WeightedAddresses{
					{
						Address: ununifitypes.StringAccAddress(addr1),
						Weight:  sdk.NewInt(1),
					},
				},
			},
			true,
		},
		{
			"invalid corresponding debt",
			CollateralAuction{
				BaseAuction: BaseAuction{
					Id:              1,
					Initiator:       testAccAddress1,
					Lot:             c("guu", 1),
					Bidder:          ununifitypes.StringAccAddress(addr1),
					Bid:             c("guu", 1),
					EndTime:         now,
					MaxEndTime:      now,
					HasReceivedBids: true,
				},
				CorrespondingDebt: sdk.Coin{Denom: "%DENOM", Amount: sdk.NewInt(1)},
			},
			false,
		},
		{
			"invalid max bid",
			CollateralAuction{
				BaseAuction: BaseAuction{
					Id:              1,
					Initiator:       testAccAddress1,
					Lot:             c("guu", 1),
					Bidder:          ununifitypes.StringAccAddress(addr1),
					Bid:             c("guu", 1),
					EndTime:         now,
					MaxEndTime:      now,
					HasReceivedBids: true,
				},
				CorrespondingDebt: c("guu", 1),
				MaxBid:            sdk.Coin{Denom: "%DENOM", Amount: sdk.NewInt(1)},
			},
			false,
		},
		{
			"invalid lot returns",
			CollateralAuction{
				BaseAuction: BaseAuction{
					Id:              1,
					Initiator:       testAccAddress1,
					Lot:             c("guu", 1),
					Bidder:          ununifitypes.StringAccAddress(addr1),
					Bid:             c("guu", 1),
					EndTime:         now,
					MaxEndTime:      now,
					HasReceivedBids: true,
				},
				CorrespondingDebt: c("guu", 1),
				MaxBid:            c("guu", 1),
				LotReturns: WeightedAddresses{
					{
						Address: nil,
						Weight:  sdk.NewInt(1),
					},
				},
			},
			false,
		},
	}

	for _, tc := range tests {

		err := tc.auction.Validate()

		if tc.expPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func TestBaseAuctionGetters(t *testing.T) {
	endTime := time.Now().Add(TestExtraEndTime)

	// Create a new BaseAuction (via SurplusAuction)
	auction := NewSurplusAuction(
		TestInitiatorModuleName,
		c(TestLotDenom, TestLotAmount),
		TestBidDenom, endTime,
	)

	auctionID := auction.GetID()
	auctionBid := auction.GetBid()
	auctionLot := auction.GetLot()
	auctionEndTime := auction.GetEndTime()
	auctionString := auction.String()

	require.Equal(t, auction.Id, auctionID)
	require.Equal(t, auction.Bid, auctionBid)
	require.Equal(t, auction.Lot, auctionLot)
	require.Equal(t, auction.EndTime, auctionEndTime)
	require.NotNil(t, auctionString)
}

func TestNewSurplusAuction(t *testing.T) {
	endTime := time.Now().Add(TestExtraEndTime)

	// Create a new SurplusAuction
	surplusAuction := NewSurplusAuction(
		TestInitiatorModuleName,
		c(TestLotDenom, TestLotAmount),
		TestBidDenom, endTime,
	)

	require.Equal(t, surplusAuction.Initiator, TestInitiatorModuleName)
	require.Equal(t, surplusAuction.Lot, c(TestLotDenom, TestLotAmount))
	require.Equal(t, surplusAuction.Bid, c(TestBidDenom, 0))
	require.Equal(t, surplusAuction.EndTime, endTime)
	require.Equal(t, surplusAuction.MaxEndTime, endTime)
}

func TestNewDebtAuction(t *testing.T) {
	endTime := time.Now().Add(TestExtraEndTime)

	// Create a new DebtAuction
	debtAuction := NewDebtAuction(
		TestInitiatorModuleName,
		c(TestBidDenom, TestBidAmount),
		c(TestLotDenom, TestLotAmount),
		endTime,
		c(TestDebtDenom, TestDebtAmount1),
	)

	require.Equal(t, debtAuction.Initiator, TestInitiatorModuleName)
	require.Equal(t, debtAuction.Lot, c(TestLotDenom, TestLotAmount))
	require.Equal(t, debtAuction.Bid, c(TestBidDenom, TestBidAmount))
	require.Equal(t, debtAuction.EndTime, endTime)
	require.Equal(t, debtAuction.MaxEndTime, endTime)
	require.Equal(t, debtAuction.CorrespondingDebt, c(TestDebtDenom, TestDebtAmount1))
}

func TestNewCollateralAuction(t *testing.T) {
	// Set up WeightedAddresses
	addresses := []sdk.AccAddress{
		sdk.AccAddress([]byte(testAccAddress1)),
		sdk.AccAddress([]byte(testAccAddress2)),
	}

	weights := []sdk.Int{
		sdk.NewInt(6),
		sdk.NewInt(8),
	}

	weightedAddresses, _ := NewWeightedAddresses(addresses, weights)

	endTime := time.Now().Add(TestExtraEndTime)

	collateralAuction := NewCollateralAuction(
		TestInitiatorModuleName,
		c(TestLotDenom, TestLotAmount),
		endTime,
		c(TestBidDenom, TestBidAmount),
		weightedAddresses,
		c(TestDebtDenom, TestDebtAmount2),
	)

	require.Equal(t, collateralAuction.BaseAuction.Initiator, TestInitiatorModuleName)
	require.Equal(t, collateralAuction.BaseAuction.Lot, c(TestLotDenom, TestLotAmount))
	require.Equal(t, collateralAuction.BaseAuction.Bid, c(TestBidDenom, 0))
	require.Equal(t, collateralAuction.BaseAuction.EndTime, endTime)
	require.Equal(t, collateralAuction.BaseAuction.MaxEndTime, endTime)
	require.Equal(t, collateralAuction.MaxBid, c(TestBidDenom, TestBidAmount))
	require.Equal(t, collateralAuction.LotReturns, weightedAddresses)
	require.Equal(t, collateralAuction.CorrespondingDebt, c(TestDebtDenom, TestDebtAmount2))
}
