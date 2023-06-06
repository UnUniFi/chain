package v1_beta4

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

	// others
	for index, value := range bank_send_list.Others {
		fromAddr, _ := sdk.AccAddressFromBech32(value.FromAddress)
		bankTarget := value.BankSendTarget
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))
		normalToken := sdk.NewCoin(bankTarget.Denom, sdk.NewInt(bankTarget.Amount))
		toAddr, _ := sdk.AccAddressFromBech32(bankTarget.ToAddress)
		if err := bankkeeper.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(normalToken)); err != nil {
			panic(err)
		}
		total_allocate_coin = total_allocate_coin.Add(normalToken)
	}

	// Check the amount of tokens sent
	assumed_coin = assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountTransferredValidator)))
	assumed_coin = assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalDelegationAmountValidator)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the others validator does not match.: Actual[%v] Assumed[%v]",
			total_allocate_coin,
			assumed_coin))
	}

	fromAddr, err := sdk.AccAddressFromBech32(FromAddressValidator)
	if err != nil {
		panic(err)
	}
	// Validator
	for index, value := range bank_send_list.Validator {
		ctx.Logger().Info(fmt.Sprintf("bank send validator :%s[%s]", strconv.Itoa(index), value.ToAddress))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin = total_allocate_coin.Add(coin)
		normalToken := sdk.NewCoin(Denom, sdk.NewInt(FundAmountValidator))
		toAddr, _ := sdk.AccAddressFromBech32(value.ToAddress)
		if err := bankkeeper.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(normalToken)); err != nil {
			panic(err)
		}
		total_allocate_coin = total_allocate_coin.Add(normalToken)
	}

	// Lend validatos
	for index, value := range bank_send_list.LendValidator {
		ctx.Logger().Info(fmt.Sprintf("bank send validator :%s[%s]", strconv.Itoa(index), value.ToAddress))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin = total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin = assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountValidator)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the validator does not match.: Actual[%v] Assumed[%v]",
			total_allocate_coin,
			assumed_coin))
	}

	fromAddr, err = sdk.AccAddressFromBech32(FromAddressEcocsytemDevelopment)
	if err != nil {
		panic(err)
	}
	// Ecocsytem Development(Community Program, Competition, Moderator)
	for index, value := range bank_send_list.EcocsytemDevelopment {
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin = total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin = assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountEcocsytemDevelopment)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the Ecocsytem Development(Community Program, Competition, Moderator) does not match.: Actual[%v] Assumed[%v]",
			total_allocate_coin,
			assumed_coin))
	}

	fromAddr, err = sdk.AccAddressFromBech32(FromAddressMarketing)
	if err != nil {
		panic(err)
	}
	// Marketing(Existing VC of Japanese company)
	for index, value := range bank_send_list.Marketing {
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin = total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin = assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountMarketing)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the Marketing(Existing VC of Japanese company) does not match.: Actual[%v] Assumed[%v]",
			total_allocate_coin,
			assumed_coin))
	}

	fromAddr, err = sdk.AccAddressFromBech32(FromAddressAdvisors)
	if err != nil {
		panic(err)
	}
	// advisors(Advisor)
	for index, value := range bank_send_list.Advisors {
		ctx.Logger().Info(fmt.Sprintf("bank send :%s", strconv.Itoa(index)))
		coin := tokenAllocation(ctx, authkeeper, bankkeeper, index, value, fromAddr)
		total_allocate_coin = total_allocate_coin.Add(coin)
	}

	// Check the amount of tokens sent
	assumed_coin = assumed_coin.Add(sdk.NewCoin(Denom, sdk.NewInt(TotalAmountAdvisors)))
	if !total_allocate_coin.IsEqual(assumed_coin) {
		panic(fmt.Sprintf("error: assumed amount sent to the advisors(Advisor) does not match.: Actual[%v] Assumed[%v]",
			total_allocate_coin,
			assumed_coin))
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
