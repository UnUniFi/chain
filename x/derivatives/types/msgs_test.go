package types

import (
	"testing"

	codecType "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/testutil/sample"
)

func TestMsgDepositToPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDepositToPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDepositToPool{
				Sender: "invalid_address",
				Amount: sdk.NewCoin("uusdc", sdk.NewInt(1)),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "insufficient amount",
			msg: MsgDepositToPool{
				Sender: sample.AccAddress(),
				Amount: sdk.NewCoin("uusdc", sdk.NewInt(0)),
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "valid msg",
			msg: MsgDepositToPool{
				Sender: sample.AccAddress(),
				Amount: sdk.NewCoin("uusdc", sdk.NewInt(1)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgWithdrawFromPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgWithdrawFromPool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgWithdrawFromPool{
				Sender:      "invalid_address",
				LptAmount:   sdk.NewInt(1),
				RedeemDenom: "uusdc",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "insufficient amount",
			msg: MsgWithdrawFromPool{
				Sender:      sample.AccAddress(),
				LptAmount:   sdk.NewInt(0),
				RedeemDenom: "uusdc",
			},
			err: ErrInsufficientAmount,
		},
		{
			name: "valid msg",
			msg: MsgWithdrawFromPool{
				Sender:      sample.AccAddress(),
				LptAmount:   sdk.NewInt(1),
				RedeemDenom: "uusdc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgOpenPosition_ValidateBasic(t *testing.T) {
	positionInstVal := PerpetualFuturesPositionInstance{
		PositionType: PositionType_LONG,
		Size_:        sdk.MustNewDecFromStr("0.001"),
		Leverage:     uint32(1),
	}

	piAny, _ := codecType.NewAnyWithValue(&positionInstVal)
	tests := []struct {
		name string
		msg  MsgOpenPosition
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgOpenPosition{
				Sender: "invalid_address",
				Margin: sdk.NewCoin("uusdc", sdk.NewInt(1)),
				Market: Market{
					BaseDenom:  "uusdc",
					QuoteDenom: "atom",
				},
				PositionInstance: *piAny,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid msg",
			msg: MsgOpenPosition{
				Sender: sample.AccAddress(),
				Margin: sdk.NewCoin("uusdc", sdk.NewInt(1)),
				Market: Market{
					BaseDenom:  "uusdc",
					QuoteDenom: "atom",
				},
				PositionInstance: *piAny,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgClosePosition_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgClosePosition
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgClosePosition{
				Sender:     "invalid_address",
				PositionId: "1",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid msg",
			msg: MsgClosePosition{
				Sender:     sample.AccAddress(),
				PositionId: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgReportLiquidation_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgReportLiquidation
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgReportLiquidation{
				Sender:          "invalid_address",
				PositionId:      "1",
				RewardRecipient: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid reward recipient",
			msg: MsgReportLiquidation{
				Sender:          sample.AccAddress(),
				PositionId:      "1",
				RewardRecipient: "invalid address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid msg",
			msg: MsgReportLiquidation{
				Sender:          sample.AccAddress(),
				PositionId:      "1",
				RewardRecipient: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgReportLevyPeriod_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgReportLevyPeriod
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgReportLevyPeriod{
				Sender:          "invalid_address",
				PositionId:      "1",
				RewardRecipient: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid reward recipient",
			msg: MsgReportLevyPeriod{
				Sender:          sample.AccAddress(),
				PositionId:      "1",
				RewardRecipient: "invalid address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid msg",
			msg: MsgReportLevyPeriod{
				Sender:          sample.AccAddress(),
				PositionId:      "1",
				RewardRecipient: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
