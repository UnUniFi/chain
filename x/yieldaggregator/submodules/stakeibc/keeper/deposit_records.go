package keeper

import (
	"fmt"

	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/spf13/cast"

	"github.com/UnUniFi/chain/utils"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"
)

func (k Keeper) CreateDepositRecordsForEpoch(ctx sdk.Context, epochNumber uint64) {
	// Create one new deposit record / host zone for the next epoch
	createDepositRecords := func(ctx sdk.Context, index int64, zoneInfo types.HostZone) error {
		k.Logger(ctx).Info(fmt.Sprintf("createDepositRecords, index: %d, zoneInfo: %s", index, zoneInfo.ConnectionId))
		depositRecord := recordstypes.DepositRecord{
			Id:                 0,
			Amount:             sdk.ZeroInt(),
			Denom:              zoneInfo.HostDenom,
			HostZoneId:         zoneInfo.ChainId,
			Status:             recordstypes.DepositRecord_TRANSFER_QUEUE,
			DepositEpochNumber: epochNumber,
		}
		k.RecordsKeeper.AppendDepositRecord(ctx, depositRecord)
		return nil
	}
	k.IterateHostZones(ctx, createDepositRecords)
}

func (k Keeper) TransferExistingDepositsToHostZones(ctx sdk.Context, epochNumber uint64, depositRecords []recordstypes.DepositRecord) {
	transferDepositRecords := utils.FilterDepositRecords(depositRecords, func(record recordstypes.DepositRecord) (condition bool) {
		isTransferRecord := record.Status == recordstypes.DepositRecord_TRANSFER_QUEUE
		isBeforeCurrentEpoch := record.DepositEpochNumber < epochNumber
		return isTransferRecord && isBeforeCurrentEpoch
	})

	ibcTransferTimeoutNanos := k.GetParams(ctx).IbcTransferTimeoutNanos

	for _, depositRecord := range transferDepositRecords {
		pstr := fmt.Sprintf("\t[TransferExistingDepositsToHostZones] Processing deposits {%d} {%s} {%d}", depositRecord.Id, depositRecord.Denom, depositRecord.Amount)
		k.Logger(ctx).Info(pstr)

		// if a TRANSFER record has 0 balance and was created in the previous epoch, it's safe to remove since it will never be updated or used"
		if depositRecord.Amount.LTE(sdk.ZeroInt()) {
			k.Logger(ctx).Info("[TransferExistingDepositsToHostZones] Empty deposit record (ID: %s)! Removing.", depositRecord.Id)
			k.RecordsKeeper.RemoveDepositRecord(ctx, depositRecord.Id)
			continue
		}

		hostZone, hostZoneFound := k.GetHostZone(ctx, depositRecord.HostZoneId)
		if !hostZoneFound {
			k.Logger(ctx).Error(fmt.Sprintf("[TransferExistingDepositsToHostZones] Host zone not found for deposit record id %d", depositRecord.Id))
			continue
		}

		hostZoneModuleAddress := hostZone.GetAddress()
		delegateAccount := hostZone.GetDelegationAccount()
		if delegateAccount == nil || delegateAccount.GetAddress() == "" {
			k.Logger(ctx).Error(fmt.Sprintf("[TransferExistingDepositsToHostZones] Zone %s is missing a delegation address!", hostZone.ChainId))
			continue
		}
		delegateAddress := delegateAccount.GetAddress()

		transferCoin := sdk.NewCoin(hostZone.GetIBCDenom(), depositRecord.Amount)
		// timeout 30 min in the future
		// NOTE: this assumes no clock drift between chains, which tendermint guarantees
		// if we onboard non-tendermint chains, we need to use the time on the host chain to
		// calculate the timeout
		// https://github.com/tendermint/tendermint/blob/v0.34.x/spec/consensus/bft-time.md
		timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + ibcTransferTimeoutNanos
		msg := ibctypes.NewMsgTransfer(ibctransfertypes.PortID, hostZone.TransferChannelId, transferCoin, hostZoneModuleAddress, delegateAddress, clienttypes.Height{}, timeoutTimestamp, "")
		k.Logger(ctx).Info(fmt.Sprintf("TransferExistingDepositsToHostZones msg %v", msg))

		err := k.RecordsKeeper.Transfer(ctx, msg, depositRecord.Id)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("\t[TransferExistingDepositsToHostZones] Failed to initiate IBC transfer to host zone, HostZone: %v, Channel: %v, Amount: %v, ModuleAddress: %v, DelegateAddress: %v, Timeout: %v",
				hostZone.ChainId, hostZone.TransferChannelId, transferCoin, hostZoneModuleAddress, delegateAddress, timeoutTimestamp))
			k.Logger(ctx).Error(fmt.Sprintf("\t[TransferExistingDepositsToHostZones] err {%s}", err.Error()))
			continue
		}
	}
}

func (k Keeper) StakeExistingDepositsOnHostZones(ctx sdk.Context, epochNumber uint64, depositRecords []recordstypes.DepositRecord) {
	stakeDepositRecords := utils.FilterDepositRecords(depositRecords, func(record recordstypes.DepositRecord) (condition bool) {
		isStakeRecord := record.Status == recordstypes.DepositRecord_DELEGATION_QUEUE
		isBeforeCurrentEpoch := record.DepositEpochNumber < epochNumber
		return isStakeRecord && isBeforeCurrentEpoch
	})

	// limit the number of staking deposits to process per epoch
	maxDepositRecordsToStake := utils.Min(len(stakeDepositRecords), cast.ToInt(k.GetParams(ctx).MaxStakeIcaCallsPerEpoch))
	k.Logger(ctx).Info(fmt.Sprintf("Staking %d out of %d deposit records", maxDepositRecordsToStake, len(stakeDepositRecords)))

	for _, depositRecord := range stakeDepositRecords[:maxDepositRecordsToStake] {
		k.Logger(ctx).Info(fmt.Sprintf("\t[StakeExistingDepositsOnHostZones] Processing deposit ID:{%d} DENOM:{%s} AMT:{%d}",
			depositRecord.Id, depositRecord.Denom, depositRecord.Amount))

		hostZone, hostZoneFound := k.GetHostZone(ctx, depositRecord.HostZoneId)
		if !hostZoneFound {
			k.Logger(ctx).Error(fmt.Sprintf("[StakeExistingDepositsOnHostZones] Host zone not found for deposit record {%d}", depositRecord.Id))
			continue
		}

		delegateAccount := hostZone.GetDelegationAccount()
		if delegateAccount == nil || delegateAccount.GetAddress() == "" {
			k.Logger(ctx).Error(fmt.Sprintf("[StakeExistingDepositsOnHostZones] Zone %s is missing a delegation address!", hostZone.ChainId))
			continue
		}

		k.Logger(ctx).Info(fmt.Sprintf("\t[StakeExistingDepositsOnHostZones] Staking %d on %s", depositRecord.Amount, hostZone.HostDenom))
		stakeAmount := sdk.NewCoin(hostZone.HostDenom, depositRecord.Amount)

		err := k.DelegateOnHost(ctx, hostZone, stakeAmount, depositRecord.Id)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Did not stake %s on %s | err: %s", stakeAmount.String(), hostZone.ChainId, err.Error()))
			continue
		} else {
			k.Logger(ctx).Info(fmt.Sprintf("Successfully submitted stake for %s on %s", stakeAmount.String(), hostZone.ChainId))
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute("hostZone", hostZone.ChainId),
				sdk.NewAttribute("newAmountStaked", depositRecord.Amount.String()),
			),
		)
	}
}
