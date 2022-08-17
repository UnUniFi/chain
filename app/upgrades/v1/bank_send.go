package v1

import (
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	authvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

func upgradeBankSend(
	ctx sdk.Context,
	authkeeper authkeeper.AccountKeeper,
	bankkeeper bankkeeper.Keeper,
	bank_send_list ResultList) error {
	ctx.Logger().Info(fmt.Sprintf("upgrade :%s", UpgradeName))

	total_allocate_coin := sdk.NewCoin("uguu", sdk.NewInt(0))
	assumed_coin := sdk.NewCoin("uguu", sdk.NewInt(0))

	// before get total supply
	before_total_supply := bankkeeper.GetSupply(ctx, "uguu")
	ctx.Logger().Info(fmt.Sprintf("bank send : total supply[%d]", before_total_supply.Amount))

	// Validator
	for index, value := range bank_send_list.Validator {
		ctx.Logger().Info(fmt.Sprintf("bank send validator :%s[%s]", strconv.Itoa(index), value.ToAddress))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value)
		total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin.Add(sdk.NewCoin("uguu", sdk.NewInt(TOTAL_AMOUNT_VALIDATOR)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the validator does not match.: Actual[%d] Assumed[%d]",
			total_allocate_coin.Amount,
			assumed_coin.Amount))
	}

	// Airdrop, Community reward and Moderator
	for index, value := range bank_send_list.AirdropCommunityRewardModerator {
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value)
		total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin.Add(sdk.NewCoin("uguu", sdk.NewInt(TOTAL_AMOUNT_EXCEPT_VALIDATOR)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the Airdrop, Community reward and Moderator does not match.: Actual[%d] Assumed[%d]",
			total_allocate_coin.Amount,
			assumed_coin.Amount))
	}

	// after get total supply
	after_total_supply := bankkeeper.GetSupply(ctx, "uguu")
	ctx.Logger().Info(fmt.Sprintf("bank send : total supply[%d]", after_total_supply.Amount))

	return nil
}

func tokenAllocation(
	ctx sdk.Context,
	authkeeper authkeeper.AccountKeeper,
	bankkeeper bankkeeper.Keeper,
	index int,
	value BankSendTarget) sdk.Coin {
	not_exist_vesting_account := true

	// check exits VestingAccount
	fromAddr, err := sdk.AccAddressFromBech32(value.FromAddress)
	if err != nil {
		panic(err)
	}
	toAddr, err := sdk.AccAddressFromBech32(value.ToAddress)
	if err != nil {
		panic(err)
	}
	accI := authkeeper.GetAccount(ctx, toAddr)
	if accI == nil {
		panic(fmt.Sprintf("error address not exist: [%s][%s]", strconv.Itoa(index), value.ToAddress))
	}

	cont_acc, ok := accI.(*authvesting.ContinuousVestingAccount)
	add_coin := sdk.NewCoins(sdk.NewCoin(value.Amount, sdk.NewInt(value.Denom)))
	if ok {
		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : ContinuousVestingAccount is exits [%s]", strconv.Itoa(index), cont_acc.String()))
		not_exist_vesting_account = false

		// 	add coins
		cont_acc.TrackDelegation(
			time.Now(),
			cont_acc.GetOriginalVesting(),
			add_coin)
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
	}

	delayed_acc, ok := accI.(*authvesting.DelayedVestingAccount)
	if ok {
		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : DelayedVestingAccount is exits [%s]", strconv.Itoa(index), delayed_acc.String()))
		not_exist_vesting_account = false

		// 	add coins
		delayed_acc.TrackDelegation(
			time.Unix(delayed_acc.EndTime, 0),
			delayed_acc.GetOriginalVesting(),
			add_coin)

		// end time sets a more future date.
		if delayed_acc.GetEndTime() < value.VestingEnds {
			delayed_acc.EndTime = value.VestingEnds
		}
		if err := delayed_acc.Validate(); err != nil {
			panic(fmt.Errorf("failed to validate DelayedVestingAccount: %w", err))
		}

		authkeeper.SetAccount(ctx, delayed_acc)
	}

	if not_exist_vesting_account {
		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : ContinuousVestingAccount / DelayedVestingAccount not exits", strconv.Itoa(index)))
		// not exist
		// 	create vesting account
		cont_vesting_acc := authvesting.NewContinuousVestingAccount(
			accI.(*types.BaseAccount),
			add_coin,
			value.VestingStarts,
			value.VestingEnds)

		if err := cont_vesting_acc.Validate(); err != nil {
			panic(fmt.Errorf("failed to validate new ContinuousVestingAccount: %w", err))
		}

		authkeeper.SetAccount(ctx, cont_vesting_acc)

		if err := bankkeeper.SendCoins(ctx, fromAddr, toAddr, add_coin); err != nil {
			panic(err)
		}
		ctx.Logger().Info(fmt.Sprintf("bank send[%s] : NewContinuousVestingAccount [%s]", strconv.Itoa(index), cont_vesting_acc.String()))
	}

	return add_coin[0]
}
