package wasmbinding

import (
	"encoding/json"
	"strconv"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/UnUniFi/chain/wasmbinding/bindings"
	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	icqkeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/keeper"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/types"
	recordskeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/keeper"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(bankKeeper *bankkeeper.BaseKeeper, icqKeeper *icqkeeper.Keeper, recordsKeeper *recordskeeper.Keeper, yieldaggregatorKeeper *yieldaggregatorkeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:               old,
			bankKeeper:            bankKeeper,
			icqKeeper:             icqKeeper,
			recordsKeeper:         recordsKeeper,
			yieldaggregatorKeeper: yieldaggregatorKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped               wasmkeeper.Messenger
	bankKeeper            *bankkeeper.BaseKeeper
	icqKeeper             *icqkeeper.Keeper
	recordsKeeper         *recordskeeper.Keeper
	yieldaggregatorKeeper *yieldaggregatorkeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		var contractMsg bindings.UnunifiMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "ununifi msg")
		}
		if contractMsg.SubmitICQRequest != nil {
			return m.submitICQRequest(ctx, contractAddr, contractMsg.SubmitICQRequest)
		}
		if contractMsg.RequestKvIcq != nil {
			return m.submitICQRequest(ctx, contractAddr, contractMsg.RequestKvIcq)
		}
		if contractMsg.IBCTransfer != nil {
			return m.ibcTransfer(ctx, contractAddr, contractMsg.IBCTransfer)
		}
		if contractMsg.DeputyDepositToVault != nil {
			return m.deputyDepositToVault(ctx, contractAddr, contractMsg.DeputyDepositToVault)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) submitICQRequest(ctx sdk.Context, contractAddr sdk.AccAddress, submitICQRequest *bindings.SubmitICQRequest) ([]sdk.Event, [][]byte, error) {
	err := PerformSubmitICQRequest(m.icqKeeper, m.bankKeeper, ctx, contractAddr, submitICQRequest)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform icq request submission")
	}
	return nil, nil, nil
}

func PerformSubmitICQRequest(f *icqkeeper.Keeper, b *bankkeeper.BaseKeeper, ctx sdk.Context, contractAddr sdk.AccAddress, submitICQRequest *bindings.SubmitICQRequest) error {
	if submitICQRequest == nil {
		return wasmvmtypes.InvalidRequest{Err: "icq request empty"}
	}

	ttl := ctx.BlockTime().Add(time.Hour*504).Unix() * time.Second.Nanoseconds() // 3 weeks
	err := f.MakeRequest(
		ctx,
		submitICQRequest.ConnectionId,
		submitICQRequest.ChainId,
		submitICQRequest.QueryPrefix,
		submitICQRequest.QueryKey,
		sdk.NewInt(-1),
		interchainquerytypes.ModuleName,
		contractAddr.String(), // set contract address on callback id
		uint64(ttl),           // ttl
		0,                     // height always 0 (which means current height)
	)
	if err != nil {
		return sdkerrors.Wrap(err, "creating icq request")
	}
	return nil
}

func (m *CustomMessenger) ibcTransfer(ctx sdk.Context, contractAddr sdk.AccAddress, ibcTransfer *wasmvmtypes.TransferMsg) ([]sdk.Event, [][]byte, error) {
	err := PerformIBCTransfer(m.recordsKeeper, ctx, contractAddr, ibcTransfer)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform ibc transfer")
	}
	return nil, nil, nil
}

func PerformIBCTransfer(f *recordskeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, ibcTransfer *wasmvmtypes.TransferMsg) error {
	if ibcTransfer == nil {
		return wasmvmtypes.InvalidRequest{Err: "icq request empty"}
	}

	amount, err := wasmkeeper.ConvertWasmCoinToSdkCoin(ibcTransfer.Amount)
	if err != nil {
		return err
	}

	err = f.ContractTransfer(
		ctx,
		&ibctransfertypes.MsgTransfer{
			SourcePort:       "transfer",
			SourceChannel:    ibcTransfer.ChannelID,
			Token:            amount,
			Sender:           contractAddr.String(),
			Receiver:         ibcTransfer.ToAddress,
			TimeoutHeight:    wasmkeeper.ConvertWasmIBCTimeoutHeightToCosmosHeight(ibcTransfer.Timeout.Block),
			TimeoutTimestamp: ibcTransfer.Timeout.Timestamp,
		})
	if err != nil {
		return sdkerrors.Wrap(err, "sending ibc transfer")
	}
	return nil
}

func (m *CustomMessenger) deputyDepositToVault(ctx sdk.Context, contractAddr sdk.AccAddress, deputyDepositToVault *bindings.DeputyDepositToVault) ([]sdk.Event, [][]byte, error) {
	err := PerformDeputyDepositToVault(m.bankKeeper, m.yieldaggregatorKeeper, ctx, contractAddr, deputyDepositToVault)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform ibc transfer")
	}
	return nil, nil, nil
}

func PerformDeputyDepositToVault(bk *bankkeeper.BaseKeeper, iyaKeeper *yieldaggregatorkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, deputyDepositToVault *bindings.DeputyDepositToVault) error {
	if deputyDepositToVault == nil {
		return wasmvmtypes.InvalidRequest{Err: "icq request empty"}
	}

	amount, err := wasmkeeper.ConvertWasmCoinToSdkCoin(deputyDepositToVault.Amount)
	if err != nil {
		return err
	}

	depositor, err := sdk.AccAddressFromBech32(deputyDepositToVault.Depositor)
	if err != nil {
		return err
	}

	err = bk.SendCoins(ctx, contractAddr, depositor, sdk.Coins{amount})
	if err != nil {
		return sdkerrors.Wrap(err, "sending coins from contract to depositor account")
	}

	vaultId, err := strconv.ParseUint(deputyDepositToVault.VaultId, 10, 64)
	if err != nil {
		return err
	}

	err = iyaKeeper.DepositAndMintLPToken(
		ctx,
		depositor,
		vaultId,
		amount,
	)
	if err != nil {
		return sdkerrors.Wrap(err, "depositing to vault")
	}
	return nil
}
