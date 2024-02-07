package v4_0_1

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

func upgradeBankSend(
	ctx sdk.Context,
	authkeeper authkeeper.AccountKeeper,
	bankkeeper bankkeeper.Keeper,
	bank_send_list ResultList) error {
	ctx.Logger().Info(fmt.Sprintf("upgrade :%s", UpgradeName))

	total_allocate_coin := sdk.NewCoin(Denom, sdk.NewInt(0))
	assumed_coin := sdk.NewCoin(Denom, sdk.NewInt(0))

	// before get total supply
	before_total_supply := bankkeeper.GetSupply(ctx, Denom)
	ctx.Logger().Info(fmt.Sprintf("bank send : total supply[%d]", before_total_supply.Amount))

	fromAddr, err := sdk.AccAddressFromBech32(FromAddress)
	if err != nil {
		panic(err)
	}
	// Campaign
	for index, value := range bank_send_list.Campaign {
		ctx.Logger().Info(fmt.Sprintf("bank send validator :%s[%s]", strconv.Itoa(index), value.ToAddress))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountCampaign)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the campaign rewards does not match.: Actual[%d] Assumed[%d]",
			total_allocate_coin.Amount,
			assumed_coin.Amount))
	}

	// after get total supply
	after_total_supply := bankkeeper.GetSupply(ctx, Denom)
	ctx.Logger().Info(fmt.Sprintf("bank send : total supply[%d]", after_total_supply.Amount))

	return nil
}

func tokenAllocation(
	ctx sdk.Context,
	authkeeper authkeeper.AccountKeeper,
	bankkeeper bankkeeper.Keeper,
	index int,
	value BankSendTarget,
	fromAddr sdk.AccAddress) sdk.Coin {
	not_exist_vesting_account := true
	add_coin := sdk.NewCoin(value.Denom, sdk.NewInt(value.Amount))

	// check exits VestingAccount
	toAddr, err := sdk.AccAddressFromBech32(value.ToAddress)
	if err != nil {
		panic(err)
	}

	// if the account is not existent, this method creates account internally
	if err := bankkeeper.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(add_coin)); err != nil {
		panic(err)
	}
	accI := authkeeper.GetAccount(ctx, toAddr)

	cont_acc, ok := accI.(*authvesting.ContinuousVestingAccount)

	if ok {
		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : ContinuousVestingAccount is exits [%s]", strconv.Itoa(index), cont_acc.String()))
		not_exist_vesting_account = false

		// 	add coins
		newAmount := cont_acc.OriginalVesting.Add(add_coin)
		cont_acc.OriginalVesting = newAmount

		// start time sets a more past date.
		if cont_acc.GetStartTime() > value.VestingStarts {
			cont_acc.StartTime = value.VestingStarts
		}
		// end time sets a more future date.
		if cont_acc.GetEndTime() < value.VestingEnds {
			cont_acc.EndTime = value.VestingEnds
		}

		if err := cont_acc.Validate(); err != nil {
			panic(fmt.Errorf("failed to validate ContinuousVestingAccount: %w", err))
		}

		authkeeper.SetAccount(ctx, cont_acc)

		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : ContinuousVestingAccount [%s]", strconv.Itoa(index), cont_acc.String()))
	}

	delayed_acc, ok := accI.(*authvesting.DelayedVestingAccount)
	if ok {
		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : DelayedVestingAccount is exits [%s]", strconv.Itoa(index), delayed_acc.String()))
		not_exist_vesting_account = false

		// 	add coins
		newAmount := delayed_acc.DelegatedVesting.Add(add_coin)
		delayed_acc.DelegatedVesting = newAmount

		// end time sets a more future date.
		if delayed_acc.GetEndTime() < value.VestingEnds {
			delayed_acc.EndTime = value.VestingEnds
		}
		if err := delayed_acc.Validate(); err != nil {
			panic(fmt.Errorf("failed to validate DelayedVestingAccount: %w", err))
		}

		authkeeper.SetAccount(ctx, delayed_acc)

		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : DelayedVestingAccount [%s]", strconv.Itoa(index), delayed_acc.String()))
	}

	if not_exist_vesting_account {
		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : ContinuousVestingAccount / DelayedVestingAccount not exits", strconv.Itoa(index)))
		// not exist
		//  create vesting account
		cont_vesting_acc := authvesting.NewContinuousVestingAccount(
			accI.(*authtypes.BaseAccount),
			sdk.NewCoins(add_coin),
			value.VestingStarts,
			value.VestingEnds)

		if err := cont_vesting_acc.Validate(); err != nil {
			panic(fmt.Errorf("failed to validate new ContinuousVestingAccount: %w", err))
		}

		authkeeper.SetAccount(ctx, cont_vesting_acc)

		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : NewContinuousVestingAccount [%s]", strconv.Itoa(index), cont_vesting_acc.String()))
	}

	return add_coin
}
