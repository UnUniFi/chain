package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func GetVaultModuleAddress(strategyContract string, maturity uint64) sdk.AccAddress {
	return authtypes.NewModuleAddress(fmt.Sprintf("irs/vault/%s/%d", strategyContract, maturity))
}

func GetLiquidityPoolModuleAddress(strategyContract string, maturity uint64) sdk.AccAddress {
	return authtypes.NewModuleAddress(fmt.Sprintf("irs/amm/%s/%d", strategyContract, maturity))
}
