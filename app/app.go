package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	ante "github.com/UnUniFi/chain/app/ante"
	appparams "github.com/UnUniFi/chain/app/params"

	// Upgrades from earlier versions of Ununifi
	v1_beta3 "github.com/UnUniFi/chain/app/upgrades/v1-beta.3"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/UnUniFi/chain/x/auction"
	auctionkeeper "github.com/UnUniFi/chain/x/auction/keeper"
	auctiontypes "github.com/UnUniFi/chain/x/auction/types"
	"github.com/UnUniFi/chain/x/cdp"
	cdpkeeper "github.com/UnUniFi/chain/x/cdp/keeper"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
	epochsmodule "github.com/UnUniFi/chain/x/epochs"
	epochsmodulekeeper "github.com/UnUniFi/chain/x/epochs/keeper"
	epochsmoduletypes "github.com/UnUniFi/chain/x/epochs/types"
	icacallbacksmodule "github.com/UnUniFi/chain/x/icacallbacks"
	icacallbacksmodulekeeper "github.com/UnUniFi/chain/x/icacallbacks/keeper"
	icacallbacksmoduletypes "github.com/UnUniFi/chain/x/icacallbacks/types"
	"github.com/UnUniFi/chain/x/incentive"
	incentivekeeper "github.com/UnUniFi/chain/x/incentive/keeper"
	incentivetypes "github.com/UnUniFi/chain/x/incentive/types"
	"github.com/UnUniFi/chain/x/interchainquery"
	interchainquerykeeper "github.com/UnUniFi/chain/x/interchainquery/keeper"
	interchainquerytypes "github.com/UnUniFi/chain/x/interchainquery/types"
	"github.com/UnUniFi/chain/x/pricefeed"
	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
	recordsmodule "github.com/UnUniFi/chain/x/records"
	recordsmodulekeeper "github.com/UnUniFi/chain/x/records/keeper"
	recordsmoduletypes "github.com/UnUniFi/chain/x/records/types"
	stakeibcmodule "github.com/UnUniFi/chain/x/stakeibc"
	stakeibcmodulekeeper "github.com/UnUniFi/chain/x/stakeibc/keeper"
	stakeibcmoduletypes "github.com/UnUniFi/chain/x/stakeibc/types"
	"github.com/UnUniFi/chain/x/ununifidist"
	ununifidistkeeper "github.com/UnUniFi/chain/x/ununifidist/keeper"
	ununifidisttypes "github.com/UnUniFi/chain/x/ununifidist/types"
	"github.com/UnUniFi/chain/x/yieldaggregator"
	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"
	"github.com/UnUniFi/chain/x/yieldfarm"
	yieldfarmkeeper "github.com/UnUniFi/chain/x/yieldfarm/keeper"
	yieldfarmtypes "github.com/UnUniFi/chain/x/yieldfarm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"

	// "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icacontroller "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	icagenesistypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v3/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v3/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v3/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	tmjson "github.com/tendermint/tendermint/libs/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/prometheus/client_golang/prometheus"

	ica "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts"
	icahost "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
)

const Name = "ununifi"

// We pull these out so we can set them with LDFLAGS in the Makefile
var (
	// If EnabledSpecificProposals is "", and this is "true", then enable all x/wasm proposals.
	// If EnabledSpecificProposals is "", and this is not "true", then disable all x/wasm proposals.
	ProposalsEnabled = "false"
	// If set to non-empty string it must be comma-separated list of values that are all a subset
	// of "EnableAllProposals" (takes precedence over ProposalsEnabled)
	// https://github.com/CosmWasm/wasmd/blob/02a54d33ff2c064f3539ae12d75d027d9c665f05/x/wasm/internal/types/proposal.go#L28-L34
	EnableSpecificProposals = ""
)

// GetEnabledProposals parses the ProposalsEnabled / EnableSpecificProposals values to
// produce a list of enabled proposals to pass into wasmd app.
func GetEnabledProposals() []wasm.ProposalType {
	if EnableSpecificProposals == "" {
		if ProposalsEnabled == "true" {
			return wasm.EnableAllProposals
		}
		return wasm.DisableAllProposals
	}
	chunks := strings.Split(EnableSpecificProposals, ",")
	proposals, err := wasm.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}
	return proposals
}

func GetWasmOpts(appOpts servertypes.AppOptions) []wasm.Option {
	var wasmOpts []wasm.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return wasmOpts
}

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers
	govProposalHandlers = wasmclient.ProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.ProposalHandler,
		upgradeclient.CancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()...),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		// liquidity.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
		auction.AppModuleBasic{},
		cdp.AppModuleBasic{},
		pricefeed.AppModuleBasic{},
		ununifidist.AppModuleBasic{},
		incentive.AppModuleBasic{},
		wasm.AppModuleBasic{},
		yieldfarm.AppModuleBasic{},
		yieldaggregator.AppModuleBasic{},
		stakeibcmodule.AppModuleBasic{},
		epochsmodule.AppModuleBasic{},
		interchainquery.AppModuleBasic{},
		ica.AppModuleBasic{},
		recordsmodule.AppModuleBasic{},
		icacallbacksmodule.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		// liquiditytypes.ModuleName:      {authtypes.Minter, authtypes.Burner},
		ibctransfertypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
		icatypes.ModuleName:             nil,
		stakeibcmoduletypes.ModuleName:  {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		interchainquerytypes.ModuleName: nil,
		auctiontypes.ModuleName:         nil,
		cdptypes.ModuleName:             {authtypes.Minter, authtypes.Burner},
		cdptypes.LiquidatorMacc:         {authtypes.Minter, authtypes.Burner},
		ununifidisttypes.ModuleName:     {authtypes.Minter},
		wasm.ModuleName:                 {authtypes.Burner},
		yieldfarmtypes.ModuleName:       {authtypes.Minter},
		yieldaggregatortypes.ModuleName: nil,
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName:          true,
		cdptypes.LiquidatorMacc:        true,
		stakeibcmoduletypes.ModuleName: true,
	}
)

var (
	_ CosmosApp               = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
	_ ibctesting.TestingApp   = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	AuthzKeeper      authzkeeper.Keeper
	// LiquidityKeeper  liquiditykeeper.Keeper
	WasmKeeper wasm.Keeper

	ICAHostKeeper       icahostkeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration
	auctionKeeper         auctionkeeper.Keeper
	cdpKeeper             cdpkeeper.Keeper
	incentiveKeeper       incentivekeeper.Keeper
	ununifidistKeeper     ununifidistkeeper.Keeper
	pricefeedKeeper       pricefeedkeeper.Keeper
	YieldfarmKeeper       yieldfarmkeeper.Keeper
	YieldaggregatorKeeper yieldaggregatorkeeper.Keeper

	ScopedStakeibcKeeper capabilitykeeper.ScopedKeeper
	StakeibcKeeper       stakeibcmodulekeeper.Keeper

	EpochsKeeper             epochsmodulekeeper.Keeper
	InterchainqueryKeeper    interchainquerykeeper.Keeper
	ScopedRecordsKeeper      capabilitykeeper.ScopedKeeper
	RecordsKeeper            recordsmodulekeeper.Keeper
	ScopedIcacallbacksKeeper capabilitykeeper.ScopedKeeper
	IcacallbacksKeeper       icacallbacksmodulekeeper.Keeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

// NewApp returns a reference to an initialized Gaia.
// NewSimApp returns a reference to an initialized SimApp.
func NewApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig appparams.EncodingConfig,
	enabledProposals []wasm.ProposalType,
	appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {

	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey,
		icahosttypes.StoreKey,
		evidencetypes.StoreKey,
		// liquiditytypes.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey, feegrant.StoreKey, authzkeeper.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
		auctiontypes.StoreKey, cdptypes.StoreKey, incentivetypes.StoreKey,
		ununifidisttypes.StoreKey, pricefeedtypes.StoreKey,
		wasm.StoreKey,
		yieldfarmtypes.StoreKey,
		yieldaggregatortypes.StoreKey,

		stakeibcmoduletypes.StoreKey,
		epochsmoduletypes.StoreKey,
		interchainquerytypes.StoreKey,
		icacontrollertypes.StoreKey,
		recordsmoduletypes.StoreKey,
		icacallbacksmoduletypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(
		app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()),
	)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)

	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)
	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.BaseApp.MsgServiceRouter(),
	)
	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		&stakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&stakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
	)
	// app.LiquidityKeeper = liquiditykeeper.NewKeeper(
	// 	appCodec,
	// 	keys[liquiditytypes.StoreKey],
	// 	app.GetSubspace(liquiditytypes.ModuleName),
	// 	app.BankKeeper,
	// 	app.AccountKeeper,
	// 	app.DistrKeeper,
	// )

	// upgrade handlers
	cfg := module.NewConfigurator(appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks(), app.incentiveKeeper.Hooks()),
	)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.TransferKeeper)

	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec, keys[icacontrollertypes.StoreKey], app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper, app.MsgServiceRouter(),
	)

	scopedIcacallbacksKeeper := app.CapabilityKeeper.ScopeToModule(icacallbacksmoduletypes.ModuleName)
	app.ScopedIcacallbacksKeeper = scopedIcacallbacksKeeper
	app.IcacallbacksKeeper = *icacallbacksmodulekeeper.NewKeeper(
		appCodec,
		keys[icacallbacksmoduletypes.StoreKey],
		keys[icacallbacksmoduletypes.MemStoreKey],
		app.GetSubspace(icacallbacksmoduletypes.ModuleName),
		scopedIcacallbacksKeeper,
		*app.IBCKeeper,
		app.ICAControllerKeeper,
	)

	app.InterchainqueryKeeper = interchainquerykeeper.NewKeeper(appCodec, keys[interchainquerytypes.StoreKey], app.IBCKeeper)
	interchainQueryModule := interchainquery.NewAppModule(appCodec, app.InterchainqueryKeeper)

	scopedRecordsKeeper := app.CapabilityKeeper.ScopeToModule(recordsmoduletypes.ModuleName)
	app.ScopedRecordsKeeper = scopedRecordsKeeper
	app.RecordsKeeper = *recordsmodulekeeper.NewKeeper(
		appCodec,
		keys[recordsmoduletypes.StoreKey],
		keys[recordsmoduletypes.MemStoreKey],
		app.GetSubspace(recordsmoduletypes.ModuleName),
		scopedRecordsKeeper,
		app.AccountKeeper,
		app.TransferKeeper,
		*app.IBCKeeper,
		app.IcacallbacksKeeper,
	)
	recordsModule := recordsmodule.NewAppModule(appCodec, app.RecordsKeeper, app.AccountKeeper, app.BankKeeper)

	scopedStakeibcKeeper := app.CapabilityKeeper.ScopeToModule(stakeibcmoduletypes.ModuleName)
	app.ScopedStakeibcKeeper = scopedStakeibcKeeper
	app.StakeibcKeeper = stakeibcmodulekeeper.NewKeeper(
		appCodec,
		keys[stakeibcmoduletypes.StoreKey],
		keys[stakeibcmoduletypes.MemStoreKey],
		app.GetSubspace(stakeibcmoduletypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		// &app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.ICAControllerKeeper,
		*app.IBCKeeper,
		scopedStakeibcKeeper,
		app.InterchainqueryKeeper,
		app.RecordsKeeper,
		app.StakingKeeper,
		app.IcacallbacksKeeper,
	)

	stakeibcModule := stakeibcmodule.NewAppModule(appCodec, app.StakeibcKeeper, app.AccountKeeper, app.BankKeeper)
	stakeibcIBCModule := stakeibcmodule.NewIBCModule(app.StakeibcKeeper)

	// Register ICQ callbacks
	err := app.InterchainqueryKeeper.SetCallbackHandler(stakeibcmoduletypes.ModuleName, app.StakeibcKeeper.CallbackHandler())
	if err != nil {
		return nil
	}

	icacallbacksModule := icacallbacksmodule.NewAppModule(appCodec, app.IcacallbacksKeeper, app.AccountKeeper, app.BankKeeper)
	// Register ICA calllbacks
	// stakeibc
	err = app.IcacallbacksKeeper.SetICACallbackHandler(stakeibcmoduletypes.ModuleName, app.StakeibcKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}
	// records
	err = app.IcacallbacksKeeper.SetICACallbackHandler(recordsmoduletypes.ModuleName, app.RecordsKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}

	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
	)
	icaModule := ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper)
	icaHostIBCModule := icahost.NewIBCModule(app.ICAHostKeeper)

	// Stack one contains
	// - IBC
	// - ICA
	// - icacallbacks
	// - stakeibc
	// - base app
	var icamiddlewareStack porttypes.IBCModule
	icamiddlewareStack = icacallbacksmodule.NewIBCModule(app.IcacallbacksKeeper, stakeibcIBCModule)
	icamiddlewareStack = icacontroller.NewIBCModule(app.ICAControllerKeeper, icamiddlewareStack)

	// Stack two contains
	// - IBC
	// - records
	// - transfer
	// - base app
	recordsStack := recordsmodule.NewIBCModule(app.RecordsKeeper, transferIBCModule)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)

	app.EvidenceKeeper = *evidenceKeeper

	// this line is used by starport scaffolding # stargate/app/keeperDefinition
	app.auctionKeeper = auctionkeeper.NewKeeper(
		appCodec,
		keys[auctiontypes.StoreKey],
		keys[auctiontypes.MemStoreKey],
		app.GetSubspace(auctiontypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
	)
	app.pricefeedKeeper = pricefeedkeeper.NewKeeper(
		appCodec,
		keys[pricefeedtypes.StoreKey],
		keys[pricefeedtypes.MemStoreKey],
		app.GetSubspace(pricefeedtypes.ModuleName),
	)
	cdpKeeper := cdpkeeper.NewKeeper(
		appCodec,
		keys[cdptypes.StoreKey],
		keys[cdptypes.MemStoreKey],
		app.GetSubspace(cdptypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.auctionKeeper,
		app.pricefeedKeeper,
		maccPerms,
	)
	app.incentiveKeeper = incentivekeeper.NewKeeper(
		appCodec,
		keys[incentivetypes.StoreKey],
		keys[incentivetypes.MemStoreKey],
		app.GetSubspace(incentivetypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&cdpKeeper,
	)
	app.ununifidistKeeper = ununifidistkeeper.NewKeeper(
		appCodec,
		keys[ununifidisttypes.StoreKey],
		keys[ununifidisttypes.MemStoreKey],
		app.GetSubspace(ununifidisttypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
	)

	app.cdpKeeper = *cdpKeeper.SetHooks(cdptypes.NewMultiCdpHooks(app.incentiveKeeper.Hooks()))

	app.YieldfarmKeeper = *yieldfarmkeeper.NewKeeper(
		appCodec,
		keys[yieldfarmtypes.StoreKey],
		app.GetSubspace(yieldfarmtypes.ModuleName),
		app.BankKeeper,
	)

	app.YieldaggregatorKeeper = yieldaggregatorkeeper.NewKeeper(
		appCodec,
		keys[yieldaggregatortypes.StoreKey],
		app.GetSubspace(yieldaggregatortypes.ModuleName),
		app.BankKeeper,
		app.YieldfarmKeeper,
		wasmkeeper.NewDefaultPermissionKeeper(app.WasmKeeper),
		app.StakeibcKeeper,
	)

	epochsKeeper := epochsmodulekeeper.NewKeeper(appCodec, keys[epochsmoduletypes.StoreKey])
	app.EpochsKeeper = *epochsKeeper.SetHooks(
		epochsmoduletypes.NewMultiEpochHooks(
			app.StakeibcKeeper.Hooks(),
			app.YieldaggregatorKeeper.Hooks(),
		),
	)
	epochsModule := epochsmodule.NewAppModule(appCodec, app.EpochsKeeper)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(yieldaggregatortypes.RouterKey, yieldaggregator.NewProposalHandler(app.YieldaggregatorKeeper)).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	supportedFeatures := "iterator,staking,stargate"
	app.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		wasmOpts...,
	)
	// The gov proposal types can be individually enabled
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals))
	}
	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.
		AddRoute(ibctransfertypes.ModuleName, recordsStack).
		AddRoute(icacontrollertypes.SubModuleName, icamiddlewareStack).
		AddRoute(icahosttypes.SubModuleName, icaHostIBCModule).
		// Note, authentication module packets are routed to the top level of the middleware stack
		AddRoute(stakeibcmoduletypes.ModuleName, icamiddlewareStack).
		AddRoute(icacallbacksmoduletypes.ModuleName, icamiddlewareStack)
	app.IBCKeeper.SetRouter(ibcRouter)

	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		govRouter,
	)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		// liquidity.NewAppModule(appCodec, app.LiquidityKeeper, app.AccountKeeper, app.BankKeeper, app.DistrKeeper),
		transferModule,
		stakeibcModule,
		epochsModule,
		interchainQueryModule,
		icaModule,
		recordsModule,
		icacallbacksModule,

		// this line is used by starport scaffolding # stargate/app/appModule
		auction.NewAppModule(appCodec, app.auctionKeeper, app.AccountKeeper, app.BankKeeper),
		cdp.NewAppModule(appCodec, app.cdpKeeper, app.AccountKeeper, app.BankKeeper, app.pricefeedKeeper),
		incentive.NewAppModule(appCodec, app.incentiveKeeper, app.AccountKeeper, app.BankKeeper, app.cdpKeeper),
		ununifidist.NewAppModule(appCodec, app.ununifidistKeeper, app.AccountKeeper, app.BankKeeper),
		pricefeed.NewAppModule(appCodec, app.pricefeedKeeper, app.AccountKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper), yieldfarm.NewAppModule(appCodec, app.YieldfarmKeeper, app.AccountKeeper, app.BankKeeper),
		yieldfarm.NewAppModule(appCodec, app.YieldfarmKeeper, app.AccountKeeper, app.BankKeeper),
		yieldaggregator.NewAppModule(appCodec, app.YieldaggregatorKeeper, app.AccountKeeper, app.BankKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		// additional non simd modules
		// liquiditytypes.ModuleName,
		ununifidisttypes.ModuleName,
		auctiontypes.ModuleName,
		cdptypes.ModuleName,
		incentivetypes.ModuleName,
		pricefeedtypes.ModuleName,

		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		stakeibcmoduletypes.ModuleName,
		epochsmoduletypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordsmoduletypes.ModuleName,
		icacallbacksmoduletypes.ModuleName,

		wasm.ModuleName,
		yieldfarmtypes.ModuleName,
		yieldaggregatortypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		// liquiditytypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		// additional non simd modules
		ununifidisttypes.ModuleName,
		auctiontypes.ModuleName,
		cdptypes.ModuleName,
		incentivetypes.ModuleName,
		pricefeedtypes.ModuleName,

		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		stakeibcmoduletypes.ModuleName,
		epochsmoduletypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordsmoduletypes.ModuleName,
		icacallbacksmoduletypes.ModuleName,

		wasm.ModuleName,
		yieldfarmtypes.ModuleName,
		yieldaggregatortypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	// NOTE: wasm module should be at the end as it can call other module functionality direct or via message dispatching during
	// genesis phase. For example bank transfer, auth account check, staking, ...
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		// liquiditytypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		authz.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
		auctiontypes.ModuleName,
		pricefeedtypes.ModuleName,
		cdptypes.ModuleName,
		incentivetypes.ModuleName,
		ununifidisttypes.ModuleName,

		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		stakeibcmoduletypes.ModuleName,
		epochsmoduletypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordsmoduletypes.ModuleName,
		icacallbacksmoduletypes.ModuleName,

		wasm.ModuleName,
		yieldfarmtypes.ModuleName,
		yieldaggregatortypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(cfg)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		// liquidity.NewAppModule(appCodec, app.LiquidityKeeper, app.AccountKeeper, app.BankKeeper, app.DistrKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
		// TODO
		// auction.NewAppModule(appCodec, app.auctionKeeper, app.AccountKeeper, app.BankKeeper),
		// cdp.NewAppModule(appCodec, app.cdpKeeper, app.AccountKeeper, app.BankKeeper, app.pricefeedKeeper),
		// incentive.NewAppModule(appCodec, app.incentiveKeeper, app.AccountKeeper, app.BankKeeper, app.cdpKeeper),
		// ununifidist.NewAppModule(appCodec, app.ununifidistKeeper, app.AccountKeeper, app.BankKeeper),
		// pricefeed.NewAppModule(appCodec, app.pricefeedKeeper, app.AccountKeeper),
		yieldfarm.NewAppModule(appCodec, app.YieldfarmKeeper, app.AccountKeeper, app.BankKeeper),
		yieldaggregator.NewAppModule(appCodec, app.YieldaggregatorKeeper, app.AccountKeeper, app.BankKeeper),
	)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)
	anteHandler, err := ante.NewAnteHandler(
		appCodec,
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			FeegrantKeeper:  app.FeeGrantKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)

	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.setupUpgradeHandlers()

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == v1_beta3.UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(fmt.Sprintf("failed to load latest version: %s", err))
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedWasmKeeper = scopedWasmKeeper
	app.ScopedICAHostKeeper = scopedICAHostKeeper
	app.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.ScopedStakeibcKeeper = scopedStakeibcKeeper

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app *App) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// GetStakingKeeper implements the TestingApp interface.
func (app *App) GetStakingKeeper() stakingkeeper.Keeper {
	return app.StakingKeeper
}

// GetTransferKeeper implements the TestingApp interface.
func (app *App) GetTransferKeeper() *ibctransferkeeper.Keeper {
	return &app.TransferKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *App) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *App) GetTxConfig() client.TxConfig {
	cfg := MakeEncodingConfig()
	return cfg.TxConfig
}

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	icaRawGenesisState := genesisState[icatypes.ModuleName]

	var icaGenesisState icagenesistypes.GenesisState
	if err := app.cdc.UnmarshalJSON(icaRawGenesisState, &icaGenesisState); err != nil {
		panic(err)
	}

	icaGenesisState.HostGenesisState.Params.AllowMessages = []string{"*"} // allow all msgs
	genesisJson, err := app.cdc.MarshalJSON(icaGenesisState)
	if err != nil {
		panic(err)
	}

	genesisState[icatypes.ModuleName] = genesisJson

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive external tokens.
func (app *App) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blockedAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	// paramsKeeper.Subspace(liquiditytypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(stakeibcmoduletypes.ModuleName)
	paramsKeeper.Subspace(epochsmoduletypes.ModuleName)
	paramsKeeper.Subspace(interchainquerytypes.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(recordsmoduletypes.ModuleName)
	paramsKeeper.Subspace(icacallbacksmoduletypes.ModuleName)

	// this line is used by starport scaffolding # stargate/app/paramSubspace
	paramsKeeper.Subspace(auctiontypes.ModuleName)
	paramsKeeper.Subspace(cdptypes.ModuleName)
	paramsKeeper.Subspace(incentivetypes.ModuleName)
	paramsKeeper.Subspace(ununifidisttypes.ModuleName)
	paramsKeeper.Subspace(pricefeedtypes.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(yieldfarmtypes.ModuleName)
	paramsKeeper.Subspace(yieldaggregatortypes.ModuleName)

	return paramsKeeper
}

func (app *App) setupUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(
		v1_beta3.UpgradeName,
		v1_beta3.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.AccountKeeper,
			app.BankKeeper))
}
