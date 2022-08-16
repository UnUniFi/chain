package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (f FarmingUnit) GetAddress() sdk.AccAddress {
	return authtypes.NewModuleAddress(fmt.Sprintf("farming_unit_%d", f.Id))
}
