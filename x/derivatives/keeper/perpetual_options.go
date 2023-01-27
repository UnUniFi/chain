package keeper

import (
	// "fmt"
	// "math/big"
	// "time"

	// cdcTypes "github.com/cosmos/cosmos-sdk/codec/types"
	ununifiTypes "github.com/UnUniFi/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) OpenPerpetualOptionsPosition(ctx sdk.Context, positionId string, sender ununifiTypes.StringAccAddress, margin sdk.Coin, market types.Market, positionInstance types.PerpetualOptionsPositionInstance) (*types.Position, error) {
	return nil, nil
}

func (k Keeper) ClosePerpetualOptionsPosition(ctx sdk.Context, position types.Position, positionInstance types.PerpetualOptionsPositionInstance) error {
	return nil
}
func (k Keeper) ReportLiquidationNeededPerpetualOptionsPosition(ctx sdk.Context, rewardRecipient ununifiTypes.StringAccAddress, remainingMargin sdk.Coin, position types.Position, positionInstance types.PerpetualOptionsPositionInstance) error {
	return nil
}
