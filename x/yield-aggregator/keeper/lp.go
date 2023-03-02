package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetLPTokenDenom(principalDenom string) string {
	return fmt.Sprintf("yield-aggregator/lp/%s", principalDenom)
}

func (k Keeper) MintLPToken(recipient string, principalDenom string, amount sdk.Int) {
	denom := k.GetLPTokenDenom(principalDenom)
	panic(denom)
	panic("implement me")
}

func (k Keeper) BurnLPToken(sender string, principalDenom string, amount sdk.Int) {
	denom := k.GetLPTokenDenom(principalDenom)
	panic(denom)
	panic("implement me")
}
