package keepers

import (
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	builderkeeper "github.com/skip-mev/pob/x/builder/keeper"

	derivativeskeeper "github.com/UnUniFi/chain/x/derivatives/keeper"
	epochskeeper "github.com/UnUniFi/chain/x/epochs/keeper"
	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"
	yieldaggregatorKeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
)

type AppKeepers struct {
	// keepers, by order of initialization
	// "Special" keepers
	ParamsKeeper          *paramskeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ConsensusParamsKeeper *consensusparamkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper  capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper     capabilitykeeper.ScopedKeeper
	ScopedICQKeeper      capabilitykeeper.ScopedKeeper

	// "Normal" keepers
	AccountKeeper *authkeeper.AccountKeeper
	BankKeeper    *bankkeeper.BaseKeeper
	AuthzKeeper   *authzkeeper.Keeper
	StakingKeeper *stakingkeeper.Keeper
	DistrKeeper   *distrkeeper.Keeper

	WasmKeeper  *wasmkeeper.Keeper
	BuildKeeper builderkeeper.Keeper

	// ununifi original keepers
	EpochsKeeper          *epochskeeper.Keeper
	YieldaggregatorKeeper *yieldaggregatorKeeper.Keeper
	PricefeedKeeper       *pricefeedkeeper.Keeper
	DerivativesKeeper     *derivativeskeeper.Keeper
}
