package keeper

import (
	// "fmt"
	// "math/big"
	// "time"

	// cdcTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifiTypes "github.com/UnUniFi/chain/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) OpenPerpetualOptionsPosition(ctx sdk.Context, positionId string, sender ununifiTypes.StringAccAddress, margin sdk.Coin, market types.Market, positionInstance types.PerpetualOptionsPositionInstance) (*types.Position, error) {
	// todo implement
	return nil, nil
}

func (k Keeper) ClosePerpetualOptionsPosition(ctx sdk.Context, position types.Position, positionInstance types.PerpetualOptionsPositionInstance) error {
	// todo implement
	return nil
}
func (k Keeper) ReportLiquidationNeededPerpetualOptionsPosition(ctx sdk.Context, rewardRecipient ununifiTypes.StringAccAddress, position types.Position, positionInstance types.PerpetualOptionsPositionInstance) error {
	// todo implement
	return nil
}
