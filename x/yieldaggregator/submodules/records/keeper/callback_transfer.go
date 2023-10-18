package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/golang/protobuf/proto" //nolint:staticcheck

	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
)

func (k Keeper) MarshalTransferCallbackArgs(ctx sdk.Context, delegateCallback types.TransferCallback) ([]byte, error) {
	out, err := proto.Marshal(&delegateCallback)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("MarshalTransferCallbackArgs %v", err.Error()))
		return nil, err
	}
	return out, nil
}

func (k Keeper) UnmarshalTransferCallbackArgs(ctx sdk.Context, delegateCallback []byte) (*types.TransferCallback, error) {
	unmarshalledTransferCallback := types.TransferCallback{}
	if err := proto.Unmarshal(delegateCallback, &unmarshalledTransferCallback); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("UnmarshalTransferCallbackArgs %v", err.Error()))
		return nil, err
	}
	return &unmarshalledTransferCallback, nil
}

func TransferCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ack *channeltypes.Acknowledgement, args []byte) error {
	k.Logger(ctx).Info("TransferCallback executing", "packet", packet)
	if ack.GetError() != "" {
		k.Logger(ctx).Error(fmt.Sprintf("TransferCallback does not handle errors %s", ack.GetError()))
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "TransferCallback does not handle errors: %s", ack.GetError())
	}
	if ack == nil {
		// timeout
		k.Logger(ctx).Error(fmt.Sprintf("TransferCallback timeout, ack is nil, packet %v", packet))
		return nil
	}

	var data ibctransfertypes.FungibleTokenPacketData
	if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error unmarshalling packet  %v", err.Error()))
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}
	k.Logger(ctx).Info(fmt.Sprintf("TransferCallback unmarshalled FungibleTokenPacketData %v", data))

	// deserialize the args
	transferCallbackData, err := k.UnmarshalTransferCallbackArgs(ctx, args)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrUnmarshalFailure, "cannot unmarshal transfer callback args: %s", err.Error())
	}
	k.Logger(ctx).Info(fmt.Sprintf("TransferCallback %v", transferCallbackData))
	depositRecord, found := k.GetDepositRecord(ctx, transferCallbackData.DepositRecordId)
	if !found {
		k.Logger(ctx).Error(fmt.Sprintf("TransferCallback deposit record not found, packet %v", packet))
		return sdkerrors.Wrapf(types.ErrUnknownDepositRecord, "deposit record not found %d", transferCallbackData.DepositRecordId)
	}
	depositRecord.Status = types.DepositRecord_DELEGATION_QUEUE
	k.SetDepositRecord(ctx, depositRecord)
	k.Logger(ctx).Info(fmt.Sprintf("\t [IBC-TRANSFER] Deposit record updated: {%v}", depositRecord.Id))
	k.Logger(ctx).Info(fmt.Sprintf("[IBC-TRANSFER] success to %s", depositRecord.HostZoneId))
	return nil
}

func (k Keeper) GetStrategyVersion(ctx sdk.Context, strategyAddr sdk.AccAddress) uint8 {
	wasmQuery := fmt.Sprintf(`{"version":{}}`)
	result, err := k.wasmReader.QuerySmart(ctx, strategyAddr, []byte(wasmQuery))
	if err != nil {
		return 0
	}

	jsonMap := make(map[string]uint8)
	err = json.Unmarshal(result, &jsonMap)
	if err != nil {
		return 0
	}

	return jsonMap["version"]
}

func ContractTransferCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ack *channeltypes.Acknowledgement, args []byte) error {
	k.Logger(ctx).Info("TransferCallback executing", "packet", packet)
	if ack.GetError() != "" {
		k.Logger(ctx).Error(fmt.Sprintf("TransferCallback does not handle errors %s", ack.GetError()))
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "TransferCallback does not handle errors: %s", ack.GetError())
	}
	if ack == nil {
		// timeout
		k.Logger(ctx).Error(fmt.Sprintf("TransferCallback timeout, ack is nil, packet %v", packet))
		return nil
	}

	var data ibctransfertypes.FungibleTokenPacketData
	if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error unmarshalling packet  %v", err.Error()))
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}
	k.Logger(ctx).Info(fmt.Sprintf("TransferCallback unmarshalled FungibleTokenPacketData %v", data))

	contractAddress, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrUnmarshalFailure, "cannot retrieve contract address: %s", err.Error())
	}

	version := k.GetStrategyVersion(ctx, contractAddress)
	_ = version

	x := types.MessageTransferCallback{}
	x.TransferCallback.Denom = data.Denom
	x.TransferCallback.Amount = data.Amount
	x.TransferCallback.Sender = data.Sender
	x.TransferCallback.Receiver = data.Receiver
	x.TransferCallback.Memo = data.Memo
	x.TransferCallback.Success = true

	callbackBytes, err := json.Marshal(x)
	if err != nil {
		return fmt.Errorf("failed to marshal MessageTransferCallback: %v", err)
	}

	_, err = k.wasmKeeper.Sudo(ctx, contractAddress, callbackBytes)
	if err != nil {
		k.Logger(ctx).Info("SudoTxQueryResult: failed to Sudo", string(callbackBytes), "error", err, "contract_address", contractAddress)
		return fmt.Errorf("failed to Sudo: %v", err)
	}
	return nil
}

func VaultTransferCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ack *channeltypes.Acknowledgement, args []byte) error {
	k.Logger(ctx).Info("VaultTransferCallback executing", "packet", packet)
	if ack.GetError() != "" {
		k.Logger(ctx).Error(fmt.Sprintf("VaultTransferCallback does not handle errors %s", ack.GetError()))
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "VaultTransferCallback does not handle errors: %s", ack.GetError())
	}
	if ack == nil {
		// timeout
		k.Logger(ctx).Error(fmt.Sprintf("VaultTransferCallback timeout, ack is nil, packet %v", packet))
		return nil
	}

	var data ibctransfertypes.FungibleTokenPacketData
	if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error unmarshalling packet  %v", err.Error()))
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}
	k.Logger(ctx).Info(fmt.Sprintf("VaultTransferCallback unmarshalled FungibleTokenPacketData %v", data))

	unmarshalledTransferCallback := types.VaultTransferCallback{}
	if err := proto.Unmarshal(args, &unmarshalledTransferCallback); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("UnmarshalVaultTransferCallbackArgs %v", err.Error()))
		return err
	}
	k.Logger(ctx).Info(fmt.Sprintf("VaultTransferCallback %v", unmarshalledTransferCallback))

	contractAddress, err := sdk.AccAddressFromBech32(unmarshalledTransferCallback.StrategyContract)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrUnmarshalFailure, "cannot retrieve contract address: %s", err.Error())
	}

	version := k.GetStrategyVersion(ctx, contractAddress)
	_ = version

	x := types.MessageDepositCallback{}
	x.DepositCallback.Denom = data.Denom
	x.DepositCallback.Amount = data.Amount
	x.DepositCallback.Sender = data.Sender
	x.DepositCallback.Receiver = data.Receiver
	x.DepositCallback.Success = true

	callbackBytes, err := json.Marshal(x)
	if err != nil {
		return fmt.Errorf("failed to marshal MessageDepositCallback: %v", err)
	}

	amount, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		return fmt.Errorf("failed to parse transfer amount: %s", data.Amount)
	}
	k.DecreaseVaultPendingDeposit(ctx, unmarshalledTransferCallback.VaultId, amount)

	_, err = k.wasmKeeper.Sudo(ctx, contractAddress, callbackBytes)
	if err != nil {
		k.Logger(ctx).Info("SudoTxQueryResult: failed to Sudo", string(callbackBytes), "error", err, "contract_address", contractAddress)
		return fmt.Errorf("failed to Sudo: %v", err)
	}

	return nil
}
