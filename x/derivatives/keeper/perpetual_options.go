package keeper

import (
	"fmt"
	"math/big"
	"time"

	cdcTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) OpenPerpetualOptionsPosition(ctx sdk.Context, positionId string, sender sdk.AccAddress, positionInstance *types.PerpetualOptionsPosition) error {
	return nil
}

func (k Keeper) ClosePerpetualOptionsPosition(ctx sdk.Context, openedPosition types.OpenedPosition, positionInstance *types.PerpetualOptionsOpenedPosition) error {
	return nil
}
