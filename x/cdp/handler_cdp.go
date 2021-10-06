package cdp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/cdp/keeper"
	"github.com/lcnem/jpyx/x/cdp/types"
)

func handleMsgCreateCdp(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateCdp) (*sdk.Result, error) {
	fmt.Printf("yakitori handleMsgCreateCdp msg.Sender.AccAddress() %s\n", msg.Sender.AccAddress())
	fmt.Printf("yakitori handleMsgCreateCdp msg.Collateral %s\n", msg.Collateral)
	fmt.Printf("yakitori handleMsgCreateCdp msg.Principal %s\n", msg.Principal)
	fmt.Printf("yakitori handleMsgCreateCdp msg.CollateralType %s\n", msg.CollateralType)
	err := k.AddCdp(ctx, msg.Sender.AccAddress(), msg.Collateral, msg.Principal, msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.AccAddress().String()),
		),
	)
	fmt.Printf("yakitori handleMsgCreateCdp msg.Sender.AccAddress() %s\n", msg.Sender.AccAddress())
	fmt.Printf("yakitori handleMsgCreateCdp msg.CollateralType %s\n", msg.CollateralType)
	id, _ := k.GetCdpID(ctx, msg.Sender.AccAddress(), msg.CollateralType)
	fmt.Printf("yakitori handleMsgCreateCdp id %d\n", id)

	fmt.Printf("yakitori handleMsgCreateCdp types.GetCdpIDBytes(id) %s\n", types.GetCdpIDBytes(id))
	return &sdk.Result{
		Data:   types.GetCdpIDBytes(id),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgDeposit(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeposit) (*sdk.Result, error) {
	err := k.DepositCollateral(ctx, msg.Owner.AccAddress(), msg.Depositor.AccAddress(), msg.Collateral, msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor.AccAddress().String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgWithdraw(ctx sdk.Context, k keeper.Keeper, msg *types.MsgWithdraw) (*sdk.Result, error) {
	err := k.WithdrawCollateral(ctx, msg.Owner.AccAddress(), msg.Depositor.AccAddress(), msg.Collateral, msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor.AccAddress().String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDrawDebt(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDrawDebt) (*sdk.Result, error) {
	err := k.AddPrincipal(ctx, msg.Sender.AccAddress(), msg.CollateralType, msg.Principal)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.AccAddress().String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgRepayDebt(ctx sdk.Context, k keeper.Keeper, msg *types.MsgRepayDebt) (*sdk.Result, error) {
	err := k.RepayPrincipal(ctx, msg.Sender.AccAddress(), msg.CollateralType, msg.Payment)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.AccAddress().String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgLiquidate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgLiquidate) (*sdk.Result, error) {
	err := k.AttemptKeeperLiquidation(ctx, msg.Keeper.AccAddress(), msg.Borrower.AccAddress(), msg.CollateralType)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Keeper.AccAddress().String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
