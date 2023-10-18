package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	"github.com/golang/protobuf/proto"

	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	icacallbackskeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	icqtypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
)

type (
	Keeper struct {
		// *cosmosibckeeper.Keeper
		Cdc                codec.BinaryCodec
		storeKey           storetypes.StoreKey
		memKey             storetypes.StoreKey
		paramstore         paramtypes.Subspace
		scopedKeeper       capabilitykeeper.ScopedKeeper
		AccountKeeper      types.AccountKeeper
		TransferKeeper     ibctransferkeeper.Keeper
		IBCKeeper          ibckeeper.Keeper
		ICACallbacksKeeper icacallbackskeeper.Keeper
		wasmKeeper         icqtypes.WasmKeeper
		wasmReader         wasmkeeper.Keeper
	}
)

func NewKeeper(
	Cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	AccountKeeper types.AccountKeeper,
	TransferKeeper ibctransferkeeper.Keeper,
	ibcKeeper ibckeeper.Keeper,
	ICACallbacksKeeper icacallbackskeeper.Keeper,
	WasmKeeper icqtypes.WasmKeeper,
	wasmReader wasmkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Cdc:                Cdc,
		storeKey:           storeKey,
		memKey:             memKey,
		paramstore:         ps,
		scopedKeeper:       scopedKeeper,
		AccountKeeper:      AccountKeeper,
		TransferKeeper:     TransferKeeper,
		IBCKeeper:          ibcKeeper,
		ICACallbacksKeeper: ICACallbacksKeeper,
		wasmKeeper:         WasmKeeper,
		wasmReader:         wasmReader,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k Keeper) Transfer(ctx sdk.Context, msg *ibctypes.MsgTransfer, depositRecordId uint64) error {
	goCtx := sdk.WrapSDKContext(ctx)

	// because TransferKeeper.Transfer doesn't return a sequence number, we need to fetch it manually
	// the sequence number isn't actually incremented here, that happens in `SendPacket`, which is triggered
	// by calling `Transfer`
	// see: https://github.com/cosmos/ibc-go/blob/48a6ae512b4ea42c29fdf6c6f5363f50645591a2/modules/core/04-channel/keeper/packet.go#L125
	sequence, found := k.IBCKeeper.ChannelKeeper.GetNextSequenceSend(ctx, msg.SourcePort, msg.SourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", msg.SourcePort, msg.SourceChannel,
		)
	}

	// trigger transfer
	_, err := k.TransferKeeper.Transfer(goCtx, msg)
	if err != nil {
		return err
	}

	// add callback data
	transferCallback := types.TransferCallback{
		DepositRecordId: depositRecordId,
	}
	k.Logger(ctx).Info(fmt.Sprintf("Marshalling TransferCallback args: %v", transferCallback))
	marshalledCallbackArgs, err := k.MarshalTransferCallbackArgs(ctx, transferCallback)
	if err != nil {
		return err
	}
	// Store the callback data
	callback := icacallbackstypes.CallbackData{
		CallbackKey:  icacallbackstypes.PacketID(msg.SourcePort, msg.SourceChannel, sequence),
		PortId:       msg.SourcePort,
		ChannelId:    msg.SourceChannel,
		Sequence:     sequence,
		CallbackId:   TRANSFER,
		CallbackArgs: marshalledCallbackArgs,
	}
	k.Logger(ctx).Info(fmt.Sprintf("Storing callback data: %v", callback))
	k.ICACallbacksKeeper.SetCallbackData(ctx, callback)
	return nil
}

func (k Keeper) ContractTransfer(ctx sdk.Context, msg *ibctypes.MsgTransfer) error {
	goCtx := sdk.WrapSDKContext(ctx)
	sequence, found := k.IBCKeeper.ChannelKeeper.GetNextSequenceSend(ctx, msg.SourcePort, msg.SourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", msg.SourcePort, msg.SourceChannel,
		)
	}

	// trigger transfer
	_, err := k.TransferKeeper.Transfer(goCtx, msg)
	if err != nil {
		return err
	}

	// Store the callback data
	callback := icacallbackstypes.CallbackData{
		CallbackKey:  icacallbackstypes.PacketID(msg.SourcePort, msg.SourceChannel, sequence),
		PortId:       msg.SourcePort,
		ChannelId:    msg.SourceChannel,
		Sequence:     sequence,
		CallbackId:   CONTRACT_TRANSFER,
		CallbackArgs: []byte{},
	}
	k.Logger(ctx).Info(fmt.Sprintf("Storing callback data: %v", callback))
	k.ICACallbacksKeeper.SetCallbackData(ctx, callback)
	return nil
}

func (k Keeper) VaultTransfer(ctx sdk.Context, vaultId uint64, contractAddr sdk.AccAddress, msg *ibctypes.MsgTransfer) error {
	goCtx := sdk.WrapSDKContext(ctx)
	sequence, found := k.IBCKeeper.ChannelKeeper.GetNextSequenceSend(ctx, msg.SourcePort, msg.SourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", msg.SourcePort, msg.SourceChannel,
		)
	}

	transferCallback := types.VaultTransferCallback{
		VaultId:          vaultId,
		StrategyContract: contractAddr.String(),
	}
	k.Logger(ctx).Info(fmt.Sprintf("Marshalling TransferCallback args: %v", transferCallback))
	marshalledCallbackArgs, err := proto.Marshal(&transferCallback)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("MarshalTransferCallbackArgs %v", err.Error()))
		return err
	}

	// Store the callback data
	callback := icacallbackstypes.CallbackData{
		CallbackKey:  icacallbackstypes.PacketID(msg.SourcePort, msg.SourceChannel, sequence),
		PortId:       msg.SourcePort,
		ChannelId:    msg.SourceChannel,
		Sequence:     sequence,
		CallbackId:   CONTRACT_TRANSFER,
		CallbackArgs: marshalledCallbackArgs,
	}
	k.Logger(ctx).Info(fmt.Sprintf("Storing callback data: %v", callback))
	k.ICACallbacksKeeper.SetCallbackData(ctx, callback)

	// trigger transfer
	_, err = k.TransferKeeper.Transfer(goCtx, msg)
	if err != nil {
		return err
	}
	return nil
}
