package v1

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

func upgradeBankSend(
	ctx sdk.Context,
	bankkeeper bankkeeper.Keeper,
	bank_send_list ResultList) error {
	ctx.Logger().Info(fmt.Sprintf("upgrade bank send:%s", UpgradeName))

	for index, value := range bank_send_list.Response {
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))

		fromAddr, err := sdk.AccAddressFromBech32(value.FromAddress)
		if err != nil {
			panic(err)
		}

		toAddr, err := sdk.AccAddressFromBech32(value.ToAddress)
		if err != nil {
			panic(err)
		}

		err = bankkeeper.SendCoins(
			ctx,
			fromAddr,
			toAddr,
			sdk.NewCoins(sdk.NewCoin(value.Amount, sdk.NewInt(value.Denom))))
		if err != nil {
			panic(err)
		}
	}

	return nil
}
