package keepers

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/CosmWasm/wasmd/x/wasm"

	buildertypes "github.com/skip-mev/pob/x/builder/types"

	// ecosystemincentivetypes "github.com/UnUniFi/chain/x/ecosystemincentive/types"
	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/types"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/types"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"

	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
)

func (appKeepers *AppKeepers) GenerateKeys() {
	// Define what keys will be used in the cosmos-sdk key/value store.
	// Cosmos-SDK modules each have a "key" that allows the application to reference what they've stored on the chain.
	appKeepers.keys = sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey, crisistypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, consensusparamtypes.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
		evidencetypes.StoreKey, capabilitytypes.StoreKey,
		authzkeeper.StoreKey, nftkeeper.StoreKey, group.StoreKey,
		// non sdk store keys
		ibcexported.StoreKey, ibctransfertypes.StoreKey, ibcfeetypes.StoreKey,
		wasm.StoreKey, icahosttypes.StoreKey,
		icacontrollertypes.StoreKey,
		buildertypes.StoreKey,

		// original modules
		nftbackedloantypes.StoreKey,
		nftfactorytypes.StoreKey,
		// ecosystemincentivetypes.StoreKey,

		// derivativestypes.StoreKey,
		// pricefeedtypes.StoreKey,

		yieldaggregatortypes.StoreKey,
		stakeibctypes.StoreKey,
		epochstypes.StoreKey,
		interchainquerytypes.StoreKey,
		recordstypes.StoreKey,
		icacallbackstypes.StoreKey,
	)

	// Define transient store keys
	appKeepers.tkeys = sdk.NewTransientStoreKeys(paramstypes.TStoreKey)

	// MemKeys are for information that is stored only in RAM.
	appKeepers.memKeys = sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
}

func (appKeepers *AppKeepers) GetKVStoreKey() map[string]*storetypes.KVStoreKey {
	return appKeepers.keys
}

func (appKeepers *AppKeepers) GetTransientStoreKey() map[string]*storetypes.TransientStoreKey {
	return appKeepers.tkeys
}

func (appKeepers *AppKeepers) GetMemoryStoreKey() map[string]*storetypes.MemoryStoreKey {
	return appKeepers.memKeys
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetKey(storeKey string) *storetypes.KVStoreKey {
	return appKeepers.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return appKeepers.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (appKeepers *AppKeepers) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return appKeepers.memKeys[storeKey]
}
