package keepers

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibcfee "github.com/cosmos/ibc-go/v7/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	transfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"

	"github.com/UnUniFi/chain/wasmbinding"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7"
	ibchookskeeper "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/keeper"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/types"

	builderkeeper "github.com/skip-mev/pob/x/builder/keeper"
	buildertypes "github.com/skip-mev/pob/x/builder/types"

	ununifinftkeeper "github.com/UnUniFi/chain/x/nft/keeper"

	epochskeeper "github.com/UnUniFi/chain/x/epochs/keeper"
	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	icacallbackskeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/keeper"
	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/types"
	interchainquerykeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/keeper"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records"
	recordskeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/keeper"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
	stakeibc "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc"
	stakeibckeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/keeper"
	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"

	nftbackedloankeeper "github.com/UnUniFi/chain/x/nftbackedloan/keeper"
	// nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"

	derivativeskeeper "github.com/UnUniFi/chain/x/derivatives/keeper"
	nftfactorykeeper "github.com/UnUniFi/chain/x/nftfactory/keeper"
	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"

	ecosystemincentivekeeper "github.com/UnUniFi/chain/x/ecosystemincentive/keeper"
	// ecosystemincentivetypes "github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

type AppKeepers struct {
	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.BaseKeeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	StakingKeeper         stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             govkeeper.Keeper
	CrisisKeeper          crisiskeeper.Keeper
	UpgradeKeeper         upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	GroupKeeper           groupkeeper.Keeper
	NFTKeeper             nftkeeper.Keeper
	UnUniFiNFTKeeper      ununifinftkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper

	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper
	WasmKeeper          wasm.Keeper
	// IBC hooks
	IBCHooksKeeper         ibchookskeeper.Keeper
	Ics20WasmHooks         ibchooks.WasmHooks
	ContractKeeper         *wasmkeeper.PermissionedKeeper
	HooksTransferIBCModule ibchooks.IBCMiddleware
	HooksICS4Wrapper       ibchooks.ICS4Middleware

	NftbackedloanKeeper nftbackedloankeeper.Keeper
	NftfactoryKeeper    nftfactorykeeper.Keeper

	YieldaggregatorKeeper    yieldaggregatorkeeper.Keeper
	StakeibcKeeper           stakeibckeeper.Keeper
	ScopedStakeibcKeeper     capabilitykeeper.ScopedKeeper
	EpochsKeeper             epochskeeper.Keeper
	InterchainqueryKeeper    interchainquerykeeper.Keeper
	ScopedRecordsKeeper      capabilitykeeper.ScopedKeeper
	RecordsKeeper            recordskeeper.Keeper
	ScopedIcacallbacksKeeper capabilitykeeper.ScopedKeeper
	IcacallbacksKeeper       icacallbackskeeper.Keeper

	DerivativesKeeper derivativeskeeper.Keeper
	PricefeedKeeper   pricefeedkeeper.Keeper

	EcosystemincentiveKeeper ecosystemincentivekeeper.Keeper

	// BuilderKeeper is the keeper that handles processing auction transactions
	BuilderKeeper builderkeeper.Keeper

	// Modules
	ICAModule      ica.AppModule
	TransferModule transfer.AppModule

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedIBCFeeKeeper        capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper
}

func NewAppKeeper(
	appCodec codec.Codec,
	bApp *baseapp.BaseApp,
	legacyAmino *codec.LegacyAmino,
	maccPerms map[string][]string,
	modAccAddrs map[string]bool,
	blockedAddress map[string]bool,
	appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option,
	enabledProposals []wasm.ProposalType,
	accountAddressPrefix string,
) *AppKeepers {
	appKeepers := AppKeepers{}

	invCheckPeriod := cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))
	// get skipUpgradeHeights from the app options
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))

	// Set keys KVStoreKey, TransientStoreKey, MemoryStoreKey
	appKeepers.GenerateKeys()

	appKeepers.ParamsKeeper = initParamsKeeper(
		appCodec,
		legacyAmino,
		appKeepers.keys[paramstypes.StoreKey],
		appKeepers.tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	appKeepers.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, appKeepers.keys[consensusparamtypes.StoreKey], authtypes.NewModuleAddress(govtypes.ModuleName).String())
	bApp.SetParamStore(&appKeepers.ConsensusParamsKeeper)

	// add capability keeper and ScopeToModule for ibc module
	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		appKeepers.keys[capabilitytypes.StoreKey],
		appKeepers.memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := appKeepers.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedICAHostKeeper := appKeepers.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedICAControllerKeeper := appKeepers.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedTransferKeeper := appKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedWasmKeeper := appKeepers.CapabilityKeeper.ScopeToModule(wasm.ModuleName)

	// add keepers
	appKeepers.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		appKeepers.keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		accountAddressPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		appKeepers.keys[banktypes.StoreKey],
		appKeepers.AccountKeeper,
		blockedAddress,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.StakingKeeper = *stakingkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[stakingtypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[minttypes.StoreKey],
		appKeepers.StakingKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[distrtypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		legacyAmino,
		appKeepers.keys[slashingtypes.StoreKey],
		appKeepers.StakingKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.CrisisKeeper = *crisiskeeper.NewKeeper(
		appCodec,
		appKeepers.keys[crisistypes.StoreKey],
		invCheckPeriod,
		appKeepers.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[feegrant.StoreKey],
		appKeepers.AccountKeeper,
	)

	// register the staking hooks
	appKeepers.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			appKeepers.DistrKeeper.Hooks(),
			appKeepers.SlashingKeeper.Hooks(),
		),
	)

	appKeepers.AuthzKeeper = authzkeeper.NewKeeper(appKeepers.keys[authzkeeper.StoreKey], appCodec, bApp.MsgServiceRouter(), appKeepers.AccountKeeper)

	groupConfig := group.DefaultConfig()
	appKeepers.GroupKeeper = groupkeeper.NewKeeper(appKeepers.keys[group.StoreKey], appCodec, bApp.MsgServiceRouter(), appKeepers.AccountKeeper, groupConfig)

	// set the governance module account as the authority for conducting upgrades
	appKeepers.UpgradeKeeper = *upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		appKeepers.keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		bApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibcexported.StoreKey],
		appKeepers.GetSubspace(ibcexported.ModuleName),
		appKeepers.StakingKeeper,
		appKeepers.UpgradeKeeper,
		scopedIBCKeeper,
	)

	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[govtypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		bApp.MsgServiceRouter(),
		govConfig,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	appKeepers.NFTKeeper = nftkeeper.NewKeeper(appKeepers.keys[nftkeeper.StoreKey],
		appCodec,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
	)

	appKeepers.UnUniFiNFTKeeper = ununifinftkeeper.NewKeeper(
		appKeepers.NFTKeeper,
		appCodec,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[evidencetypes.StoreKey],
		appKeepers.StakingKeeper,
		appKeepers.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	appKeepers.EvidenceKeeper = *evidenceKeeper

	// IBC Fee Module keeper
	appKeepers.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		appCodec, appKeepers.keys[ibcfeetypes.StoreKey],
		appKeepers.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper, appKeepers.AccountKeeper, appKeepers.BankKeeper,
	)

	appKeepers.keys[ibchookstypes.StoreKey] = storetypes.NewKVStoreKey(ibchookstypes.StoreKey)
	appKeepers.IBCHooksKeeper = ibchookskeeper.NewKeeper(
		appKeepers.keys[ibchookstypes.StoreKey],
	)
	appKeepers.Ics20WasmHooks = ibchooks.NewWasmHooks(&appKeepers.IBCHooksKeeper, nil, accountAddressPrefix) // The contract keeper needs to be set later

	// Create Transfer Keepers
	appKeepers.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibctransfertypes.StoreKey],
		appKeepers.GetSubspace(ibctransfertypes.ModuleName),
		appKeepers.HooksICS4Wrapper, // essentially still app.IBCKeeper.ChannelKeeper under the hood because no hook overrides
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		scopedTransferKeeper,
	)

	appKeepers.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[icahosttypes.StoreKey],
		appKeepers.GetSubspace(icahosttypes.SubModuleName),
		appKeepers.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		scopedICAHostKeeper,
		bApp.MsgServiceRouter(),
	)
	appKeepers.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[icacontrollertypes.StoreKey],
		appKeepers.GetSubspace(icacontrollertypes.SubModuleName),
		appKeepers.IBCFeeKeeper, // may be replaced with middleware such as ics29 fee
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper,
		bApp.MsgServiceRouter(),
	)

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	availableCapabilities := "iterator,staking,stargate,cosmwasm_1_1,cosmwasm_1_2"

	wasmOpts = append(wasmbinding.RegisterCustomPlugins(&appKeepers.BankKeeper, &appKeepers.InterchainqueryKeeper, &appKeepers.RecordsKeeper), wasmOpts...)

	appKeepers.WasmKeeper = wasm.NewKeeper(
		appCodec,
		appKeepers.keys[wasm.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		distrkeeper.NewQuerier(appKeepers.DistrKeeper),
		appKeepers.IBCFeeKeeper, // ISC4 Wrapper: fee IBC middleware
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		appKeepers.TransferKeeper,
		bApp.MsgServiceRouter(),
		bApp.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		availableCapabilities,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		wasmOpts...,
	)

	// Pass the contract keeper to all the structs (generally ICS4Wrappers for ibc middlewares) that need it
	appKeepers.ContractKeeper = wasmkeeper.NewDefaultPermissionKeeper(appKeepers.WasmKeeper)
	appKeepers.Ics20WasmHooks.ContractKeeper = &appKeepers.WasmKeeper
	appKeepers.HooksICS4Wrapper = ibchooks.NewICS4Middleware(
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.Ics20WasmHooks,
	)
	// Hooks Middleware
	transferIBCModule := transfer.NewIBCModule(appKeepers.TransferKeeper)
	appKeepers.HooksTransferIBCModule = ibchooks.NewIBCMiddleware(&transferIBCModule, &appKeepers.HooksICS4Wrapper)

	// Instantiate the builder keeper, store keys, and module manager
	appKeepers.BuilderKeeper = builderkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[buildertypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.DistrKeeper,
		appKeepers.StakingKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.NftfactoryKeeper = nftfactorykeeper.NewKeeper(
		appCodec,
		appKeepers.keys[nftfactorytypes.StoreKey],
		appKeepers.keys[nftfactorytypes.MemStoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.UnUniFiNFTKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// nftbackedloanKeeper := nftbackedloankeeper.NewKeeper(
	// 	appCodec,
	// 	appKeepers.keys[nftbackedloantypes.StoreKey],
	// 	appKeepers.keys[nftbackedloantypes.MemStoreKey],
	// 	appKeepers.GetSubspace(nftbackedloantypes.ModuleName),
	// 	appKeepers.AccountKeeper,
	// 	appKeepers.BankKeeper,
	// 	appKeepers.UnUniFiNFTKeeper,
	// )

	// appKeepers.EcosystemincentiveKeeper = ecosystemincentivekeeper.NewKeeper(
	// 	appCodec,
	// 	appKeepers.keys[ecosystemincentivetypes.StoreKey],
	// 	appKeepers.GetSubspace(ecosystemincentivetypes.ModuleName),
	// 	appKeepers.AccountKeeper,
	// 	appKeepers.BankKeeper,
	// 	appKeepers.DistrKeeper,
	// 	// same as the feeCollectorName in the distribution module
	// 	authtypes.FeeCollectorName,
	// )

	// create Keeper objects which have Hooks
	// appKeepers.NftbackedloanKeeper = nftbackedloanKeeper
	// appKeepers.NftbackedloanKeeper = *nftbackedloanKeeper.SetHooks(nftbackedloantypes.NewMultiNftbackedloanHooks(appKeepers.EcosystemincentiveKeeper.Hooks()))

	// appKeepers.PricefeedKeeper = pricefeedkeeper.NewKeeper(
	// 	appCodec,
	// 	appKeepers.keys[pricefeedtypes.StoreKey],
	// 	appKeepers.keys[pricefeedtypes.MemStoreKey],
	// 	appKeepers.GetSubspace(pricefeedtypes.ModuleName),
	// 	appKeepers.BankKeeper,
	// )

	// appKeepers.DerivativesKeeper = derivativeskeeper.NewKeeper(
	// 	appCodec,
	// 	appKeepers.keys[derivativestypes.StoreKey],
	// 	appKeepers.keys[derivativestypes.MemStoreKey],
	// 	appKeepers.GetSubspace(derivativestypes.ModuleName),
	// 	appKeepers.AccountKeeper,
	// 	appKeepers.BankKeeper,
	// 	appKeepers.PricefeedKeeper,
	// 	appKeepers.UnUniFiNFTKeeper,
	// )

	scopedIcacallbacksKeeper := appKeepers.CapabilityKeeper.ScopeToModule(icacallbackstypes.ModuleName)
	appKeepers.ScopedIcacallbacksKeeper = scopedIcacallbacksKeeper
	appKeepers.IcacallbacksKeeper = *icacallbackskeeper.NewKeeper(
		appCodec,
		appKeepers.keys[icacallbackstypes.StoreKey],
		appKeepers.keys[icacallbackstypes.MemStoreKey],
		appKeepers.GetSubspace(icacallbackstypes.ModuleName),
		scopedIcacallbacksKeeper,
		*appKeepers.IBCKeeper,
		appKeepers.ICAControllerKeeper,
	)

	appKeepers.InterchainqueryKeeper = interchainquerykeeper.NewKeeper(
		appCodec,
		appKeepers.keys[interchainquerytypes.StoreKey],
		appKeepers.IBCKeeper,
		&appKeepers.WasmKeeper,
	)

	scopedRecordsKeeper := appKeepers.CapabilityKeeper.ScopeToModule(recordstypes.ModuleName)
	appKeepers.ScopedRecordsKeeper = scopedRecordsKeeper
	appKeepers.RecordsKeeper = *recordskeeper.NewKeeper(
		appCodec,
		appKeepers.keys[recordstypes.StoreKey],
		appKeepers.keys[recordstypes.MemStoreKey],
		appKeepers.GetSubspace(recordstypes.ModuleName),
		scopedRecordsKeeper,
		appKeepers.AccountKeeper,
		appKeepers.TransferKeeper,
		*appKeepers.IBCKeeper,
		appKeepers.IcacallbacksKeeper,
		&appKeepers.WasmKeeper,
	)

	scopedStakeibcKeeper := appKeepers.CapabilityKeeper.ScopeToModule(stakeibctypes.ModuleName)
	appKeepers.ScopedStakeibcKeeper = scopedStakeibcKeeper
	appKeepers.StakeibcKeeper = stakeibckeeper.NewKeeper(
		appCodec,
		appKeepers.keys[stakeibctypes.StoreKey],
		appKeepers.keys[stakeibctypes.MemStoreKey],
		appKeepers.GetSubspace(stakeibctypes.ModuleName),
		appKeepers.IBCKeeper.ChannelKeeper,
		// &appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.ICAControllerKeeper,
		*appKeepers.IBCKeeper,
		scopedStakeibcKeeper,
		scopedIBCKeeper,
		appKeepers.InterchainqueryKeeper,
		appKeepers.RecordsKeeper,
		appKeepers.StakingKeeper,
		appKeepers.IcacallbacksKeeper,
	)

	// Register ICQ callbacks
	err = appKeepers.InterchainqueryKeeper.SetCallbackHandler(stakeibctypes.ModuleName, appKeepers.StakeibcKeeper.CallbackHandler())
	if err != nil {
		return nil
	}

	// Register ICA calllbacks
	// stakeibc
	err = appKeepers.IcacallbacksKeeper.SetICACallbackHandler(icacontrollertypes.SubModuleName, appKeepers.StakeibcKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}
	// records
	err = appKeepers.IcacallbacksKeeper.SetICACallbackHandler(recordstypes.ModuleName, appKeepers.RecordsKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}

	appKeepers.YieldaggregatorKeeper = yieldaggregatorkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[yieldaggregatortypes.StoreKey],
		appKeepers.BankKeeper,
		wasmkeeper.NewDefaultPermissionKeeper(appKeepers.WasmKeeper),
		appKeepers.WasmKeeper,
		appKeepers.StakeibcKeeper,
		appKeepers.RecordsKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	epochsKeeper := epochskeeper.NewKeeper(appCodec, appKeepers.keys[epochstypes.StoreKey])
	appKeepers.EpochsKeeper = *epochsKeeper.SetHooks(
		epochstypes.NewMultiEpochHooks(
			appKeepers.YieldaggregatorKeeper.Hooks(),
			appKeepers.StakeibcKeeper.Hooks(),
		),
	)

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://docs.cosmos.network/main/modules/gov#proposal-messages
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(appKeepers.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(&appKeepers.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(appKeepers.IBCKeeper.ClientKeeper))

	// The gov proposal types can be individually enabled
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(appKeepers.WasmKeeper, enabledProposals))
	}

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://docs.cosmos.network/main/modules/gov#proposal-messages
	appKeepers.GovKeeper.SetLegacyRouter(govRouter)

	// Create Transfer Stack
	var transferStack porttypes.IBCModule
	transferStack = transfer.NewIBCModule(appKeepers.TransferKeeper)
	transferStack = records.NewIBCModule(appKeepers.RecordsKeeper, transferStack)
	transferStack = ibcfee.NewIBCMiddleware(transferStack, appKeepers.IBCFeeKeeper)

	// RecvPacket, message that originates from core IBC and goes down to app, the flow is:
	// channel.RecvPacket -> fee.OnRecvPacket -> icaHost.OnRecvPacket
	var icaHostStack porttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(appKeepers.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, appKeepers.IBCFeeKeeper)

	// Create fee enabled wasm ibc Stack
	var wasmStack porttypes.IBCModule
	wasmStack = wasm.NewIBCHandler(appKeepers.WasmKeeper, appKeepers.IBCKeeper.ChannelKeeper, appKeepers.IBCFeeKeeper)
	wasmStack = ibcfee.NewIBCMiddleware(wasmStack, appKeepers.IBCFeeKeeper)

	// Stack two (Stakeibc Stack) contains
	// - IBC
	// - ICA
	// - stakeibc
	// - base app
	var stakeibcStack porttypes.IBCModule = stakeibc.NewIBCModule(appKeepers.StakeibcKeeper)
	stakeibcStack = icacontroller.NewIBCMiddleware(stakeibcStack, appKeepers.ICAControllerKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.
		AddRoute(ibctransfertypes.ModuleName, transferStack).
		AddRoute(wasm.ModuleName, wasmStack).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		// Stakeibc Stack
		AddRoute(icacontrollertypes.SubModuleName, stakeibcStack).
		AddRoute(stakeibctypes.ModuleName, stakeibcStack)

	appKeepers.IBCKeeper.SetRouter(ibcRouter)

	appKeepers.ScopedIBCKeeper = scopedIBCKeeper
	appKeepers.ScopedTransferKeeper = scopedTransferKeeper
	appKeepers.ScopedWasmKeeper = scopedWasmKeeper
	appKeepers.ScopedICAHostKeeper = scopedICAHostKeeper
	appKeepers.ScopedICAControllerKeeper = scopedICAControllerKeeper
	appKeepers.ScopedStakeibcKeeper = scopedStakeibcKeeper

	return &appKeepers
}

// GetSubspace returns a param subspace for a given module name.
func (appKeepers *AppKeepers) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := appKeepers.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibcexported.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(buildertypes.ModuleName)

	// original modules
	paramsKeeper.Subspace(nftfactorytypes.ModuleName)
	// paramsKeeper.Subspace(nftbackedloantypes.ModuleName)
	// paramsKeeper.Subspace(ecosystemincentivetypes.ModuleName)

	// paramsKeeper.Subspace(pricefeedtypes.ModuleName)
	// paramsKeeper.Subspace(derivativestypes.ModuleName)

	paramsKeeper.Subspace(stakeibctypes.ModuleName)
	paramsKeeper.Subspace(epochstypes.ModuleName)
	paramsKeeper.Subspace(interchainquerytypes.ModuleName)
	paramsKeeper.Subspace(recordstypes.ModuleName)
	paramsKeeper.Subspace(icacallbackstypes.ModuleName)

	// Deprecated: Just for migration
	paramsKeeper.Subspace(yieldaggregatortypes.ModuleName)

	return paramsKeeper
}
