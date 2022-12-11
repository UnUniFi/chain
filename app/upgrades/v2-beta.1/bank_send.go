package v2_beta1

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

	fromAddr, err := sdk.AccAddressFromBech32(FromAddressValidator)
	if err != nil {
		panic(err)
	}
	// Validator
	for index, value := range bank_send_list.Validator {
		ctx.Logger().Info(fmt.Sprintf("bank send validator :%s[%s]", strconv.Itoa(index), value.ToAddress))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin.Add(coin)
		normalToken := sdk.NewCoin(Denom, sdk.NewInt(FundAmountValidator))
		toAddr, _ := sdk.AccAddressFromBech32(value.ToAddress)
		if err := bankkeeper.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(normalToken)); err != nil {
			panic(err)
		}
		total_allocate_coin.Add(normalToken)
	}

	// Lend validatos
	for index, value := range bank_send_list.LendValidator {
		ctx.Logger().Info(fmt.Sprintf("bank send validator :%s[%s]", strconv.Itoa(index), value.ToAddress))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountValidator)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the validator does not match.: Actual[%d] Assumed[%d]",
			total_allocate_coin.Amount,
			assumed_coin.Amount))
	}

	fromAddr, err = sdk.AccAddressFromBech32(FromAddressAirdrop)
	if err != nil {
		panic(err)
	}
	// Airdrop, Community reward and Moderator
	for index, value := range bank_send_list.AirdropCommunityRewardModerator {
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountExceptValidator)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the Airdrop, Community reward and Moderator does not match.: Actual[%d] Assumed[%d]",
			total_allocate_coin.Amount,
			assumed_coin.Amount))
	}

	// airdrop forfeit
	toAddr, err := sdk.AccAddressFromBech32(ToAddressAirdropForfeit)
	if err != nil {
		panic(err)
	}
	for index, value := range bank_send_list.AirdropForfeit {
		err = forfeitToken(ctx, authkeeper, bankkeeper, index, value, toAddr)
		if err != nil {
			panic(err)
		}
	}

	// others
	for index, value := range bank_send_list.Others {
		fromAddr, _ := sdk.AccAddressFromBech32(value.FromAddress)
		bankTarget := value.BankSendTarget
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))
		// subtract token from sender
		token := sdk.NewCoin(bankTarget.Denom, sdk.NewInt(bankTarget.Amount))
		changeVestingAmount(ctx, authkeeper, token, fromAddr, false)
		_ = tokenAllocation(ctx, authkeeper, bankkeeper, index, bankTarget, fromAddr)
	}

	// after get total supply
	after_total_supply := bankkeeper.GetSupply(ctx, Denom)
	ctx.Logger().Info(fmt.Sprintf("bank send : total supply[%d]", after_total_supply.Amount))

	return nil
}

func forfeitToken(
	ctx sdk.Context,
	authkeeper authkeeper.AccountKeeper,
	bankkeeper bankkeeper.Keeper,
	index int,
	fromAddr string,
	toAddr sdk.AccAddress,
) error {
	addr, err := sdk.AccAddressFromBech32(fromAddr)
	if err != nil {
		return err
	}
	accI := authkeeper.GetAccount(ctx, addr)
	if accI == nil {
		panic(fmt.Sprintf("error address not exist: [%s][%s]", strconv.Itoa(index), fromAddr))
	}
	cont_acc, ok := accI.(*authvesting.ContinuousVestingAccount)
	zeroCoins := sdk.NewCoins(sdk.NewCoin(Denom, sdk.ZeroInt()))
	if ok {
		// add coin amount to send forfeited amount of token to ToAirdropAddress
		add_coins := sdk.NewCoins(sdk.NewCoin(Denom, cont_acc.OriginalVesting.AmountOf(Denom)))
		cont_acc.OriginalVesting = zeroCoins

		if err := cont_acc.Validate(); err != nil {
			panic(fmt.Errorf("failed to validate ContinuousVestingAccount: %w", err))
		}

		authkeeper.SetAccount(ctx, cont_acc)

		if err := bankkeeper.SendCoins(ctx, addr, toAddr, add_coins); err != nil {
			return err
		}
	}
	return nil
}

func changeVestingAmount(
	ctx sdk.Context,
	authkeeper authkeeper.AccountKeeper,
	amount sdk.Coin,
	addr sdk.AccAddress,
	add bool,
) {
	// subtract token from sender
	accI := authkeeper.GetAccount(ctx, addr)

	cont_acc, _ := accI.(*authvesting.ContinuousVestingAccount)
	if add {
		modifiedAmount := cont_acc.OriginalVesting.Add(amount)
		cont_acc.OriginalVesting = modifiedAmount
	} else {
		modifiedAmount := cont_acc.OriginalVesting.Sub(sdk.NewCoins(amount))
		cont_acc.OriginalVesting = modifiedAmount
	}

	if err := cont_acc.Validate(); err != nil {
		panic(fmt.Errorf("failed to validate ContinuousVestingAccount: %w", err))
	}
	authkeeper.SetAccount(ctx, cont_acc)
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

	// if the account is not existant, this method creates account internally
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
