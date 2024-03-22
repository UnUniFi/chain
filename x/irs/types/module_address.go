package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func GetVaultModuleAddress(pool TranchePool) sdk.AccAddress {
	return authtypes.NewModuleAddress(fmt.Sprintf("irs/vault/%d", pool.Id))
}

func GetLiquidityPoolModuleAddress(pool TranchePool) sdk.AccAddress {
	return authtypes.NewModuleAddress(fmt.Sprintf("irs/amm/%d", pool.Id))
}
