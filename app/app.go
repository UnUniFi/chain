package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"

	"github.com/spf13/cast"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	pobabci "github.com/skip-mev/pob/abci"
	"github.com/skip-mev/pob/mempool"

	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"

	"github.com/UnUniFi/chain/app/keepers"
	"github.com/UnUniFi/chain/app/upgrades"
	v1_beta3 "github.com/UnUniFi/chain/app/upgrades/v1-beta.3"
	v2_1 "github.com/UnUniFi/chain/app/upgrades/v2.1"
	v2_2 "github.com/UnUniFi/chain/app/upgrades/v2.2"
	v3 "github.com/UnUniFi/chain/app/upgrades/v3"
	v3_1 "github.com/UnUniFi/chain/app/upgrades/v3.1"
)

const Name = "ununifi"

// We pull these out so we can set them with LDFLAGS in the Makefile
var (
	// If EnabledSpecificProposals is "", and this is "true", then enable all x/wasm proposals.
	// If EnabledSpecificProposals is "", and this is not "true", then disable all x/wasm proposals.
	ProposalsEnabled = "true"
	// If set to non-empty string it must be comma-separated list of values that are all a subset
	// of "EnableAllProposals" (takes precedence over ProposalsEnabled)
	// https://github.com/CosmWasm/wasmd/blob/02a54d33ff2c064f3539ae12d75d027d9c665f05/x/wasm/internal/types/proposal.go#L28-L34
	EnableSpecificProposals = "MigrateContract,UpdateAdmin,ClearAdmin"
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

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName:    true,
		stakeibctypes.ModuleName: true,
	}

	Upgrades = []upgrades.Upgrade{v1_beta3.Upgrade, v2_1.Upgrade, v2_2.Upgrade, v3.Upgrade, v3_1.Upgrade}
)

var (
	_ runtime.AppI            = (*App)(nil)
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
	keepers.AppKeepers

	legacyAmino *codec.LegacyAmino
	appCodec    codec.Codec
	// txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator

	// Custom checkTx handler
	checkTxHandler pobabci.CheckTx
}

// NewApp returns a reference to an initialized Gaia.
// NewSimApp returns a reference to an initialized SimApp.
func NewApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option,
	enabledProposals []wasm.ProposalType,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	encodingConfig := MakeEncodingConfig()

	appCodec, legacyAmino := encodingConfig.Codec, encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(Name, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	// for original AuctionFactory
	// maxTx can be changed (0 is unlimited)
	maxTx := 0
	factory := mempool.NewDefaultAuctionFactory(txConfig.TxDecoder())
	mempool := mempool.NewAuctionMempool(txConfig.TxDecoder(), txConfig.TxEncoder(), maxTx, factory)
	bApp.SetMempool(mempool)

	app := &App{
		BaseApp:           bApp,
		AppKeepers:        keepers.AppKeepers{},
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
	}

	moduleAccountAddresses := app.ModuleAccountAddrs()

	// Setup keepers
	app.AppKeepers = *keepers.NewAppKeeper(
		appCodec,
		bApp,
		legacyAmino,
		maccPerms,
		moduleAccountAddresses,
		app.BlockedModuleAccountAddrs(moduleAccountAddresses),
		appOpts,
		AccountAddressPrefix,
	)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		appModules(app, encodingConfig, skipGenesisInvariants)...,
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(orderBeginBlockers()...)

	app.mm.SetOrderEndBlockers(orderEndBlockers()...)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	// NOTE: wasm module should be at the end as it can call other module functionality direct or via message dispatching during
	// genesis phase. For example bank transfer, auth account check, staking, ...
	app.mm.SetOrderInitGenesis(orderInitGenesis()...)

	app.mm.RegisterInvariants(&app.AppKeepers.CrisisKeeper)
	app.configurator = module.NewConfigurator(appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))

	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AppKeepers.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	keys := app.AppKeepers.GetKVStoreKey()
	app.MountKVStores(keys)
	app.MountTransientStores(app.AppKeepers.GetTransientStoreKey())
	app.MountMemoryStores(app.AppKeepers.GetMemoryStoreKey())

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, logger, keys); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}
	anteHandler := app.CreateAnteHandler(encodingConfig.TxConfig, wasmConfig, keys[wasm.StoreKey], mempool)
	app.SetAnteHandler(anteHandler)
	app.SetupProposalHandler(encodingConfig.TxConfig, mempool, anteHandler)
	app.SetupChackTx(encodingConfig.TxConfig, mempool, anteHandler)

	app.SetupUpgradeHandlers()
	app.SetupUpgradeStoreLoaders()

	// must be before Loading version
	// requires the snapshot store to be created and registered as a BaseAppOption
	// see cmd/wasmd/root.go: 206 - 214 approx
	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.AppKeepers.WasmKeeper),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %s", err))
		}
	}

	// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
	// antehandlers, but are run _after_ the `runMsgs` execution. They are also
	// defined as a chain, and have the same signature as antehandlers.
	//
	// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
	// meaning that both `runMsgs` and `postHandler` state will be committed if
	// both are successful, and both will be reverted if any of the two fails.
	//
	// The SDK exposes a default postHandlers chain, which comprises of only
	// one decorator: the Transaction Tips decorator. However, some chains do
	// not need it by default, so feel free to comment the next line if you do
	// not need tips.
	// To read more about tips:
	// https://docs.cosmos.network/main/core/tips.html
	//
	// Please note that changing any of the anteHandler or postHandler chain is
	// likely to be a state-machine breaking change, which needs a coordinated
	// upgrade.
	app.setPostHandler()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			logger.Error("error on loading last version", "err", err)
			os.Exit(1)
		}
		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		// Initialize pinned codes in wasmvm as they are not persisted there
		if err := app.AppKeepers.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
	}

	return app
}

func (app *App) CreateAnteHandler(txConfig client.TxConfig, wasmConfig wasmtypes.WasmConfig, txCounterStoreKey storetypes.StoreKey, mempool *mempool.AuctionMempool) sdk.AnteHandler {
	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AppKeepers.AccountKeeper,
				BankKeeper:      app.AppKeepers.BankKeeper,
				SignModeHandler: txConfig.SignModeHandler(),
				FeegrantKeeper:  app.AppKeepers.FeeGrantKeeper,
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCKeeper:         app.AppKeepers.IBCKeeper,
			WasmConfig:        &wasmConfig,
			TXCounterStoreKey: txCounterStoreKey,

			TxEncoder:     txConfig.TxEncoder(),
			BuilderKeeper: app.AppKeepers.BuilderKeeper,
			Mempool:       mempool,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}
	return anteHandler
}

func (app *App) SetupProposalHandler(txConfig client.TxConfig, mempool *mempool.AuctionMempool, anteHandler sdk.AnteHandler) {
	// Set the custom gov proposal handler on BaseApp.
	// These methods perform custom logic when validator nodes are instructed to create proposals
	proposalHandlers := pobabci.NewProposalHandler(
		mempool,
		app.Logger(),
		anteHandler,
		txConfig.TxEncoder(),
		txConfig.TxDecoder(),
	)
	app.SetPrepareProposal(proposalHandlers.PrepareProposalHandler())
	app.SetProcessProposal(proposalHandlers.ProcessProposalHandler())
}

func (app *App) SetupChackTx(txConfig client.TxConfig, mempool *mempool.AuctionMempool, anteHandler sdk.AnteHandler) {
	// Set the custom CheckTx handler on BaseApp.
	checkTxHandler := pobabci.NewCheckTxHandler(
		app.BaseApp,
		txConfig.TxDecoder(),
		mempool,
		anteHandler,
		app.ChainID(),
	)
	app.SetCheckTx(checkTxHandler.CheckTx())
}

// ChainID gets chainID from private fields of BaseApp
// Should be removed once SDK 0.50.x will be adopted
func (app *App) ChainID() string {
	field := reflect.ValueOf(app.BaseApp).Elem().FieldByName("chainID")
	return field.String()
}

// CheckTx will check the transaction with the provided checkTxHandler. We override the default
// handler so that we can verify bid transactions before they are inserted into the mempool.
// With the POB CheckTx, we can verify the bid transaction and all of the bundled transactions
// before inserting the bid transaction into the mempool.
func (app *App) CheckTx(req abci.RequestCheckTx) abci.ResponseCheckTx {
	return app.checkTxHandler(req)
}

// SetCheckTx sets the checkTxHandler for the app.
func (app *App) SetCheckTx(handler pobabci.CheckTx) {
	app.checkTxHandler = handler
}

func (app *App) setPostHandler() {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		panic(err)
	}
	app.SetPostHandler(postHandler)
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app *App) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *App) Configurator() module.Configurator {
	return app.configurator
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.AppKeepers.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
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
	return app.legacyAmino
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

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.AppKeepers.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

// RegisterNodeService registers the node gRPC service on the app gRPC router.
func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

func (app *App) SetupUpgradeStoreLoaders() {
	upgradeInfo, err := app.AppKeepers.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.AppKeepers.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades))
		}
	}

}

func (app *App) SetupUpgradeHandlers() {
	for _, upgrade := range Upgrades {
		app.AppKeepers.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.mm,
				app.configurator,
				app.BaseApp,
				&app.AppKeepers,
			),
		)
	}
}
