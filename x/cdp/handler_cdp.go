package cdp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/cdp/keeper"
	"github.com/lcnem/jpyx/x/cdp/types"
)

func handleMsgCreateCdp(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateCDP) (*sdk.Result, error) {
	err := k.AddCdp(ctx, msg.Sender, msg.Collateral, msg.Principal, msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	)
	id, _ := k.GetCdpID(ctx, msg.Sender, msg.CollateralType)

	return &sdk.Result{
		Data:   types.GetCdpIDBytes(id),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgDeposit(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeposit) (*sdk.Result, error) {
	err := k.DepositCollateral(ctx, msg.Owner, msg.Depositor, msg.Collateral, msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgWithdraw(ctx sdk.Context, k keeper.Keeper, msg *types.MsgWithdraw) (*sdk.Result, error) {
	err := k.WithdrawCollateral(ctx, msg.Owner, msg.Depositor, msg.Collateral, msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDrawDebt(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDrawDebt) (*sdk.Result, error) {
	err := k.AddPrincipal(ctx, msg.Sender, msg.CollateralType, msg.Principal)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgRepayDebt(ctx sdk.Context, k keeper.Keeper, msg *types.MsgRepayDebt) (*sdk.Result, error) {
	err := k.RepayPrincipal(ctx, msg.Sender, msg.CollateralType, msg.Payment)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgLiquidate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgLiquidate) (*sdk.Result, error) {
	err := k.AttemptKeeperLiquidation(ctx, msg.Keeper, msg.Borrower, msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Keeper.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
