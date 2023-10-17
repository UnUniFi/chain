package keeper

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/iancoleman/orderedmap"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) GetStrategyVersion(ctx sdk.Context, strategy types.Strategy) uint8 {
	wasmQuery := fmt.Sprintf(`{"version":{}}`)
	contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
	result, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
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

type DenomInfo struct {
	Denom            string `json:"denom"`
	TargetChainId    string `json:"target_chain_id"`
	TargetChainDenom string `json:"target_chain_denom"`
	TargetChainAddr  string `json:"target_chain_addr"`
}

func (k Keeper) GetStrategyDepositInfo(ctx sdk.Context, strategy types.Strategy) (info DenomInfo) {
	wasmQuery := fmt.Sprintf(`{"deposit_denom":{}}`)
	contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
	result, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &info)
	if err != nil {
		return DenomInfo{}
	}

	return
}

// GetStrategyCount get the total number of Strategy
func (k Keeper) GetStrategyCount(ctx sdk.Context, vaultDenom string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetStrategyCount set the total number of Strategy
func (k Keeper) SetStrategyCount(ctx sdk.Context, vaultDenom string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendStrategy appends a Strategy in the store with a new id and update the count
func (k Keeper) AppendStrategy(
	ctx sdk.Context,
	vaultDenom string,
	strategy types.Strategy,
) uint64 {
	// Create the strategy
	count := k.GetStrategyCount(ctx, vaultDenom)

	// Set the ID of the appended value
	strategy.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	bz := k.cdc.MustMarshal(&strategy)
	store.Set(GetStrategyIDBytes(strategy.Id), bz)

	// Update strategy count
	k.SetStrategyCount(ctx, vaultDenom, count+1)

	return count
}

// SetStrategy set a specific Strategy in the store
func (k Keeper) SetStrategy(ctx sdk.Context, strategy types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(strategy.Denom))
	b := k.cdc.MustMarshal(&strategy)
	store.Set(GetStrategyIDBytes(strategy.Id), b)
}

// GetStrategy returns a Strategy from its id
func (k Keeper) GetStrategy(ctx sdk.Context, vaultDenom string, id uint64) (val types.Strategy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	b := store.Get(GetStrategyIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStrategy removes a Strategy from the store
func (k Keeper) RemoveStrategy(ctx sdk.Context, vaultDenom string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	store.Delete(GetStrategyIDBytes(id))
}

func (k Keeper) MigrateAllLegacyStrategies(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(""))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	legacyStrategies := []types.LegacyStrategy{}
	for ; iterator.Valid(); iterator.Next() {
		var val types.LegacyStrategy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		legacyStrategies = append(legacyStrategies, val)
	}

	for _, legacyStrategy := range legacyStrategies {
		strategy := types.Strategy{
			Denom:           legacyStrategy.Denom,
			Id:              legacyStrategy.Id,
			ContractAddress: legacyStrategy.ContractAddress,
			Name:            legacyStrategy.Name,
			Description:     "",
			GitUrl:          legacyStrategy.GitUrl,
		}
		k.SetStrategy(ctx, strategy)
	}
}

// GetAllStrategy returns all Strategy
func (k Keeper) GetAllStrategy(ctx sdk.Context, vaultDenom string) (list []types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Strategy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStrategyIDBytes returns the byte representation of the ID
func GetStrategyIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetStrategyIDFromBytes returns ID in uint64 format from a byte array
func GetStrategyIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func CalculateTransferRoute(currChannels, tarChannels []types.TransferChannel) []types.TransferChannel {
	diffStartIndex := int(0)
	for index, currChan := range currChannels {
		if len(tarChannels) <= index {
			diffStartIndex = index
		}
		tarChan := tarChannels[index]
		if currChan.ChainId != tarChan.ChainId {
			diffStartIndex = index
			break
		}
	}

	route := []types.TransferChannel{}
	for index := len(currChannels) - 1; index >= diffStartIndex; index-- {
		route = append(route, currChannels[index])
	}
	for index := diffStartIndex; index < len(tarChannels); index++ {
		route = append(route, tarChannels[index])
	}
	return route
}

func (k Keeper) ComposePacketForwardMetadata(ctx sdk.Context, channels []types.TransferChannel, finalReceiver string) (string, *PacketMetadata) {
	if len(channels) == 0 {
		return "", nil
	}

	if len(channels) == 1 {
		return finalReceiver, nil
	}

	receiver, nextForward := k.ComposePacketForwardMetadata(ctx, channels[1:], finalReceiver)
	nextForwardBz, err := json.Marshal(nextForward)
	if err != nil {
		return "", nil
	}
	retries := uint8(2)
	params, err := k.GetParams(ctx)
	if err != nil {
		return "", nil
	}
	ibcTransferTimeoutNanos := params.IbcTransferTimeoutNanos
	return k.GetIntermediaryReceiver(ctx, channels[0].ChainId), &PacketMetadata{
		Forward: &ForwardMetadata{
			Receiver: receiver,
			Port:     ibctransfertypes.PortID,
			Channel:  channels[0].ChannelId,
			Timeout:  Duration(ibcTransferTimeoutNanos),
			Retries:  &retries,
			Next:     NewJSONObject(false, nextForwardBz, orderedmap.OrderedMap{}),
		},
	}
}

// stake into strategy
func (k Keeper) StakeToStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy, amount sdk.Int) error {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	balances := k.bankKeeper.GetAllBalances(ctx, vaultModAddr)

	switch strategy.ContractAddress {
	case "x/ibc-staking":
		stakeCoin := sdk.NewCoin(strategy.Denom, amount)
		return k.stakeibcKeeper.LiquidStake(
			ctx,
			vaultModAddr,
			stakeCoin,
		)
	default:
		version := k.GetStrategyVersion(ctx, strategy)
		_ = version
		contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
		strategyDenomAmount := amount
		if balances.AmountOf(strategy.Denom).LT(amount) {
			strategyDenomAmount = balances.AmountOf(strategy.Denom)
		}

		if strategyDenomAmount.IsPositive() {
			wasmMsg := `{"stake":{}}`
			_, err := k.wasmKeeper.Execute(ctx, contractAddr, vaultModAddr, []byte(wasmMsg), sdk.Coins{sdk.NewCoin(strategy.Denom, strategyDenomAmount)})
			if err != nil {
				return err
			}
		}

		remaining := amount.Sub(strategyDenomAmount)
		for _, balance := range balances {
			if remaining.IsZero() {
				return nil
			}
			denomInfo := k.GetDenomInfo(ctx, balance.Denom)
			if balance.Denom != strategy.Denom && denomInfo.Symbol == vault.Symbol {
				stakeAmount := remaining
				if balance.Amount.LT(remaining) {
					stakeAmount = balance.Amount
				}
				err := k.ExecuteVaultTransfer(ctx, vault, strategy, sdk.NewCoin(balance.Denom, stakeAmount))
				if err != nil {
					return err
				}
				k.recordsKeeper.IncreaseVaultPendingDeposit(ctx, vault.Id, stakeAmount)
				remaining = remaining.Sub(stakeAmount)
			}
		}

		return nil
	}
}

func (k Keeper) ExecuteVaultTransfer(ctx sdk.Context, vault types.Vault, strategy types.Strategy, stakeCoin sdk.Coin) error {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
	info := k.GetStrategyDepositInfo(ctx, strategy)
	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}
	ibcTransferTimeoutNanos := params.IbcTransferTimeoutNanos
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + ibcTransferTimeoutNanos
	denomInfo := k.GetDenomInfo(ctx, info.Denom)
	symbolInfo := k.GetSymbolInfo(ctx, vault.Symbol)
	tarChannel := types.TransferChannel{}
	for _, channel := range symbolInfo.Channels {
		if channel.ChainId == info.TargetChainId {
			tarChannel = channel
		}
	}
	transferRoute := CalculateTransferRoute(denomInfo.Channels, []types.TransferChannel{tarChannel})
	initialReceiver, metadata := k.ComposePacketForwardMetadata(ctx, transferRoute, info.TargetChainAddr)
	memo, err := json.Marshal(metadata)

	msg := ibctypes.NewMsgTransfer(
		ibctransfertypes.PortID,
		transferRoute[0].ChannelId,
		stakeCoin,
		vaultModAddr.String(),
		initialReceiver,
		clienttypes.Height{},
		timeoutTimestamp,
		string(memo),
	)
	err = k.recordsKeeper.VaultTransfer(ctx, vault.Id, contractAddr, msg)
	return err
}

// unstake worth of withdrawal amount from the strategy
func (k Keeper) UnstakeFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy, amount sdk.Int, recipient string) error {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	if recipient == "" {
		recipient = vaultModAddr.String()
	}
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		{
			err := k.stakeibcKeeper.RedeemStake(
				ctx,
				vaultModAddr,
				sdk.NewCoin(strategy.Denom, amount),
				recipient,
			)
			if err != nil {
				return err
			}

			return nil
		}
	default:
		version := k.GetStrategyVersion(ctx, strategy)
		wasmMsg := ""
		switch version {
		case 0:
			wasmMsg = fmt.Sprintf(`{"unstake":{"amount":"%s"}}`, amount.String())
		default: // case 1+
			wasmMsg = fmt.Sprintf(`{"unstake":{"share_amount":"%s", "recipient": "%s"}}`, amount.String(), recipient)
		}
		contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, vaultModAddr, []byte(wasmMsg), sdk.Coins{})
		return err
	}
}

type AmountsResp struct {
	TotalDeposited string `json:"total_deposited"`
	BondingStandby string `json:"bonding_standby"`
	Bonded         string `json:"bonded"`
	Unbonding      string `json:"unbonding"`
}

func (k Keeper) GetAmountFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy) (sdk.Coin, error) {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		updatedAmount := k.stakeibcKeeper.GetUpdatedBalance(ctx, vaultModAddr, strategy.Denom)
		return sdk.NewCoin(strategy.Denom, updatedAmount), nil
	default:
		version := k.GetStrategyVersion(ctx, strategy)
		switch version {
		case 0:
			wasmQuery := fmt.Sprintf(`{"bonded":{"addr": "%s"}}`, vaultModAddr.String())
			contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
			resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
			if err != nil {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
			}
			amountStr := strings.ReplaceAll(string(resp), "\"", "")
			amount, ok := sdk.NewIntFromString(amountStr)
			if !ok {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), nil
			}
			return sdk.NewCoin(strategy.Denom, amount), err
		default: // case 1+
			wasmQuery := fmt.Sprintf(`{"amounts":{"addr": "%s"}}`, vaultModAddr.String())
			contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
			resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
			if err != nil {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
			}

			parsedAmounts := AmountsResp{}
			err = json.Unmarshal(resp, &parsedAmounts)
			if err != nil {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
			}

			amount, ok := sdk.NewIntFromString(parsedAmounts.Bonded)
			if !ok {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), nil
			}
			return sdk.NewCoin(strategy.Denom, amount), err
		}
	}
}

func (k Keeper) GetUnbondingAmountFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy) (sdk.Coin, error) {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		zone, err := k.stakeibcKeeper.GetHostZoneFromIBCDenom(ctx, strategy.Denom)
		if err != nil {
			return sdk.Coin{}, err
		}
		unbondingAmount := k.recordsKeeper.GetUserRedemptionRecordBySenderAndHostZone(ctx, vaultModAddr, zone.ChainId)
		return sdk.NewCoin(strategy.Denom, unbondingAmount), nil
	default:
		version := k.GetStrategyVersion(ctx, strategy)
		switch version {
		case 0:
			wasmQuery := fmt.Sprintf(`{"unbonding":{"addr": "%s"}}`, vaultModAddr.String())
			contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
			resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
			if err != nil {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
			}
			amountStr := strings.ReplaceAll(string(resp), "\"", "")
			amount, ok := sdk.NewIntFromString(amountStr)
			if !ok {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), nil
			}
			return sdk.NewCoin(strategy.Denom, amount), err
		default: // case 1+
			wasmQuery := fmt.Sprintf(`{"amounts":{"addr": "%s"}}`, vaultModAddr.String())
			contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
			resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
			if err != nil {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
			}

			parsedAmounts := AmountsResp{}
			err = json.Unmarshal(resp, &parsedAmounts)
			if err != nil {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
			}

			amount, ok := sdk.NewIntFromString(parsedAmounts.Unbonding)
			if !ok {
				return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), nil
			}
			return sdk.NewCoin(strategy.Denom, amount), err
		}
	}
}
