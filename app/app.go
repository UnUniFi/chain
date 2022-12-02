package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
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
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"
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
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"

	ante "github.com/UnUniFi/chain/app/ante"
	appparams "github.com/UnUniFi/chain/app/params"

	// "github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	// ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	// ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	// ibc "github.com/cosmos/ibc-go/v3/modules/core"
	// ibcclient "github.com/cosmos/ibc-go/v3/modules/core/02-client"
	// ibcclientclient "github.com/cosmos/ibc-go/v3/modules/core/02-client/client"
	// ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	// porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	// ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	// ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	// "github.com/gravity-devs/liquidity/x/liquidity"
	// liquiditykeeper "github.com/gravity-devs/liquidity/x/liquidity/keeper"
	// liquiditytypes "github.com/gravity-devs/liquidity/x/liquidity/types"

	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	// this line is used by starport scaffolding # stargate/app/moduleImport
	"github.com/UnUniFi/chain/x/auction"
	auctionkeeper "github.com/UnUniFi/chain/x/auction/keeper"
	auctiontypes "github.com/UnUniFi/chain/x/auction/types"
	"github.com/UnUniFi/chain/x/cdp"
	cdpkeeper "github.com/UnUniFi/chain/x/cdp/keeper"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
	ecosystemincentive "github.com/UnUniFi/chain/x/ecosystem-incentive"
	ecosystemincentivekeeper "github.com/UnUniFi/chain/x/ecosystem-incentive/keeper"
	ecosystemincentivetypes "github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	"github.com/UnUniFi/chain/x/incentive"
	incentivekeeper "github.com/UnUniFi/chain/x/incentive/keeper"
	incentivetypes "github.com/UnUniFi/chain/x/incentive/types"
	"github.com/UnUniFi/chain/x/nftmarket"
	nftmarketkeeper "github.com/UnUniFi/chain/x/nftmarket/keeper"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
	"github.com/UnUniFi/chain/x/nftmint"
	nftmintkeeper "github.com/UnUniFi/chain/x/nftmint/keeper"
	nftminttypes "github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/UnUniFi/chain/x/pricefeed"
	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
	"github.com/UnUniFi/chain/x/ununifidist"
	ununifidistkeeper "github.com/UnUniFi/chain/x/ununifidist/keeper"
	ununifidisttypes "github.com/UnUniFi/chain/x/ununifidist/types"
	// "github.com/CosmWasm/wasmd/x/wasm"
	// wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
)

const Name = "ununifi"
const upgradeName = "Alpha"

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
// func GetEnabledProposals() []wasm.ProposalType {
// 	if EnableSpecificProposals == "" {
// 		if ProposalsEnabled == "true" {
// 			return wasm.EnableAllProposals
// 		}
// 		return wasm.DisableAllProposals
// 	}
// 	chunks := strings.Split(EnableSpecificProposals, ",")
// 	proposals, err := wasm.ConvertToProposals(chunks)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return proposals
// }

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		// ibcclientclient.UpdateClientProposalHandler,
		// ibcclientclient.UpgradeProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)
	// govProposalHandlers = append(govProposalHandlers, wasmclient.ProposalHandlers...)

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
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		// ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		// transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		// liquidity.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
		auction.AppModuleBasic{},
		cdp.AppModuleBasic{},
		ecosystemincentive.AppModuleBasic{},
		pricefeed.AppModuleBasic{},
		ununifidist.AppModuleBasic{},
		incentive.AppModuleBasic{},
		nftmint.AppModuleBasic{},
		nftmarket.AppModuleBasic{},
		// wasm.AppModuleBasic{},
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
		// ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		auctiontypes.ModuleName:            nil,
		cdptypes.ModuleName:                {authtypes.Minter, authtypes.Burner},
		cdptypes.LiquidatorMacc:            {authtypes.Minter, authtypes.Burner},
		ecosystemincentivetypes.ModuleName: nil,
		ununifidisttypes.ModuleName:        {authtypes.Minter},
		// wasm.ModuleName:             {authtypes.Burner},
		nft.ModuleName:            nil,
		nftminttypes.ModuleName:   nil,
		nftmarkettypes.ModuleName: nil,
		// nftmarkettypes.NftTradingFee: nil,
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName:   true,
		cdptypes.LiquidatorMacc: true,
	}
)

var (
	_ CosmosApp               = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
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
	// msgSvcRouter      *authmiddleware.MsgServiceRouter
	// legacyRouter      sdk.Router

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

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
	// IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper evidencekeeper.Keeper
	// TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper feegrantkeeper.Keeper
	AuthzKeeper    authzkeeper.Keeper
	NFTKeeper      nftkeeper.Keeper
	// LiquidityKeeper  liquiditykeeper.Keeper
	// WasmKeeper wasm.Keeper

	// make scoped keepers public for test purposes
	// ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	// ScopedWasmKeeper     capabilitykeeper.ScopedKeeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration
	auctionKeeper            auctionkeeper.Keeper
	cdpKeeper                cdpkeeper.Keeper
	EcosystemincentiveKeeper ecosystemincentivekeeper.Keeper
	incentiveKeeper          incentivekeeper.Keeper
	ununifidistKeeper        ununifidistkeeper.Keeper
	pricefeedKeeper          pricefeedkeeper.Keeper
	NftmintKeeper            nftmintkeeper.Keeper
	NftmarketKeeper          nftmarketkeeper.Keeper

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
	// enabledProposals []wasm.ProposalType,
	// this line is used by starport scaffolding # stargate/app/newArgument
	appOpts servertypes.AppOptions,
	// wasmOpts []wasm.Option,
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
		govtypes.StoreKey, paramstypes.StoreKey,
		// ibchost.StoreKey,
		upgradetypes.StoreKey,
		evidencetypes.StoreKey,
		// liquiditytypes.StoreKey,
		// ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey, feegrant.StoreKey, authzkeeper.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
		auctiontypes.StoreKey, cdptypes.StoreKey,
		ecosystemincentivetypes.StoreKey, incentivetypes.StoreKey,
		ununifidisttypes.StoreKey, pricefeedtypes.StoreKey,
		// wasm.StoreKey,
		nftkeeper.StoreKey,
		nftminttypes.StoreKey,
		nftmarkettypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		// legacyRouter:      authmiddleware.NewLegacyRouter(),
		// msgSvcRouter:      authmiddleware.NewMsgServiceRouter(interfaceRegistry),
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tkeys:          tkeys,
		memKeys:        memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(
		app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()),
	)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	// scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	// scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	// scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
	app.CapabilityKeeper.Seal()

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms, sdk.Bech32MainPrefix,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)
	app.AuthzKeeper = authzkeeper.NewKeeper(keys[authzkeeper.StoreKey], appCodec, app.BaseApp.MsgServiceRouter(), app.AccountKeeper)
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
		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName,
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
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())
	app.NFTKeeper = nftkeeper.NewKeeper(keys[nftkeeper.StoreKey], appCodec, app.AccountKeeper, app.BankKeeper)

	// app.LiquidityKeeper = liquiditykeeper.NewKeeper(
	// 	appCodec,
	// 	keys[liquiditytypes.StoreKey],
	// 	app.GetSubspace(liquiditytypes.ModuleName),
	// 	app.BankKeeper,
	// 	app.AccountKeeper,
	// 	app.DistrKeeper,
	// )

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks(), app.incentiveKeeper.Hooks()),
	)

	// Create IBC Keeper
	// app.IBCKeeper = ibckeeper.NewKeeper(
	// 	appCodec,
	// 	keys[ibchost.StoreKey],
	// 	app.GetSubspace(ibchost.ModuleName),
	// 	app.StakingKeeper,
	// 	app.UpgradeKeeper,
	// 	scopedIBCKeeper,
	// )

	// register the proposal types
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)) //.
		// AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	// Create Transfer Keepers
	// app.TransferKeeper = ibctransferkeeper.NewKeeper(
	// 	appCodec,
	// 	keys[ibctransfertypes.StoreKey],
	// 	app.GetSubspace(ibctransfertypes.ModuleName),
	// 	app.IBCKeeper.ChannelKeeper,
	// 	&app.IBCKeeper.PortKeeper,
	// 	app.AccountKeeper,
	// 	app.BankKeeper,
	// 	scopedTransferKeeper,
	// )
	// transferModule := transfer.NewAppModule(app.TransferKeeper)

	// create static IBC router, add transfer route, then set and seal it
	// ibcRouter := porttypes.NewRouter()
	// ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)

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
	app.EcosystemincentiveKeeper = ecosystemincentivekeeper.NewKeeper(
		appCodec,
		keys[ecosystemincentivetypes.StoreKey],
		app.GetSubspace(ecosystemincentivetypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
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

	app.NftmintKeeper = nftmintkeeper.NewKeeper(
		appCodec,
		keys[nftminttypes.StoreKey],
		keys[nftminttypes.MemStoreKey],
		app.GetSubspace(nftminttypes.ModuleName),
		app.AccountKeeper,
		app.NFTKeeper,
	)

	nftmarketKeeper := nftmarketkeeper.NewKeeper(
		appCodec,
		encodingConfig.TxConfig,
		keys[nftmarkettypes.StoreKey],
		keys[nftmarkettypes.MemStoreKey],
		app.GetSubspace(nftmarkettypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.NFTKeeper,
	)

	// create Keeper objects which have Hooks
	app.cdpKeeper = *cdpKeeper.SetHooks(cdptypes.NewMultiCdpHooks(app.incentiveKeeper.Hooks()))
	app.NftmarketKeeper = *nftmarketKeeper.SetHooks(nftmarkettypes.NewMultiNftmarketHooks(app.EcosystemincentiveKeeper.Hooks()))

	// wasmDir := filepath.Join(homePath, "wasm")
	// wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	// if err != nil {
	// 	panic(fmt.Sprintf("error while reading wasm config: %s", err))
	// }
	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	// supportedFeatures := "iterator,staking,stargate"
	// app.WasmKeeper = wasm.NewKeeper(
	// 	appCodec,
	// 	keys[wasm.StoreKey],
	// 	app.GetSubspace(wasm.ModuleName),
	// 	app.AccountKeeper,
	// 	app.BankKeeper,
	// 	app.StakingKeeper,
	// 	app.DistrKeeper,
	// 	app.IBCKeeper.ChannelKeeper,
	// 	&app.IBCKeeper.PortKeeper,
	// 	scopedWasmKeeper,
	// 	app.TransferKeeper,
	// 	app.MsgServiceRouter(),
	// 	app.GRPCQueryRouter(),
	// 	wasmDir,
	// 	wasmConfig,
	// 	supportedFeatures,
	// 	wasmOpts...,
	// )
	// // The gov proposal types can be individually enabled
	// if len(enabledProposals) != 0 {
	// 	govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals))
	// }
	// ibcRouter.AddRoute(wasm.ModuleName, wasm.NewIBCHandler(app.WasmKeeper, app.IBCKeeper.ChannelKeeper))
	// app.IBCKeeper.SetRouter(ibcRouter)

	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter, app.BaseApp.MsgServiceRouter(), govConfig,
	)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
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
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		nftmodule.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		// ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		// liquidity.NewAppModule(appCodec, app.LiquidityKeeper, app.AccountKeeper, app.BankKeeper, app.DistrKeeper),
		// transferModule,
		// this line is used by starport scaffolding # stargate/app/appModule
		auction.NewAppModule(appCodec, app.auctionKeeper, app.AccountKeeper, app.BankKeeper),
		cdp.NewAppModule(appCodec, app.cdpKeeper, app.AccountKeeper, app.BankKeeper, app.pricefeedKeeper),
		ecosystemincentive.NewAppModule(appCodec, app.EcosystemincentiveKeeper, app.AccountKeeper, app.BankKeeper),
		incentive.NewAppModule(appCodec, app.incentiveKeeper, app.AccountKeeper, app.BankKeeper, app.cdpKeeper),
		ununifidist.NewAppModule(appCodec, app.ununifidistKeeper, app.AccountKeeper, app.BankKeeper),
		pricefeed.NewAppModule(appCodec, app.pricefeedKeeper, app.AccountKeeper),
		nftmint.NewAppModule(appCodec, app.NftmintKeeper, app.AccountKeeper),
		nftmarket.NewAppModule(appCodec, app.NftmarketKeeper, app.AccountKeeper, app.BankKeeper),
		// wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper),
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
		nft.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		// additional non simd modules
		// liquiditytypes.ModuleName,
		ununifidisttypes.ModuleName,
		auctiontypes.ModuleName,
		cdptypes.ModuleName,
		ecosystemincentivetypes.ModuleName,
		incentivetypes.ModuleName,
		pricefeedtypes.ModuleName,
		nftminttypes.ModuleName,
		nftmarkettypes.ModuleName,

		// ibchost.ModuleName,
		// ibctransfertypes.ModuleName,
		// wasm.ModuleName,
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
		nft.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		// additional non simd modules
		ununifidisttypes.ModuleName,
		auctiontypes.ModuleName,
		cdptypes.ModuleName,
		incentivetypes.ModuleName,
		pricefeedtypes.ModuleName,
		nftminttypes.ModuleName,
		nftmarkettypes.ModuleName,
		ecosystemincentivetypes.ModuleName,
		// ibchost.ModuleName,
		// ibctransfertypes.ModuleName,
		// wasm.ModuleName,
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
		// ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		authz.ModuleName,
		nft.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
		auctiontypes.ModuleName,
		pricefeedtypes.ModuleName,
		cdptypes.ModuleName,
		incentivetypes.ModuleName,
		ununifidisttypes.ModuleName,
		nftminttypes.ModuleName,
		nftmarkettypes.ModuleName,
		ecosystemincentivetypes.ModuleName,
		// ibchost.ModuleName,
		// ibctransfertypes.ModuleName,
		// wasm after ibc transfer
		// wasm.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

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
		nftmodule.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		// liquidity.NewAppModule(appCodec, app.LiquidityKeeper, app.AccountKeeper, app.BankKeeper, app.DistrKeeper),
		// wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper),
		// ibc.NewAppModule(app.IBCKeeper),
		// transferModule,
		// TODO
		// auction.NewAppModule(appCodec, app.auctionKeeper, app.AccountKeeper, app.BankKeeper),
		// cdp.NewAppModule(appCodec, app.cdpKeeper, app.AccountKeeper, app.BankKeeper, app.pricefeedKeeper),
		// incentive.NewAppModule(appCodec, app.incentiveKeeper, app.AccountKeeper, app.BankKeeper, app.cdpKeeper),
		// ununifidist.NewAppModule(appCodec, app.ununifidistKeeper, app.AccountKeeper, app.BankKeeper),
		// pricefeed.NewAppModule(appCodec, app.pricefeedKeeper, app.AccountKeeper),
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
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}

	app.SetAnteHandler(anteHandler)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx sdk.Context, _ upgradetypes.Plan, _ module.VersionMap) (module.VersionMap, error) {
			// app.IBCKeeper.ConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())

			fromVM := make(map[string]uint64)
			for moduleName := range app.mm.Modules {
				fromVM[moduleName] = 1
			}
			// override versions for _new_ modules as to not skip InitGenesis
			fromVM[authz.ModuleName] = 0
			fromVM[feegrant.ModuleName] = 0

			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Added: []string{authz.ModuleName, feegrant.ModuleName},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(fmt.Sprintf("failed to load latest version: %s", err))
		}
	}

	// app.ScopedIBCKeeper = scopedIBCKeeper
	// app.ScopedTransferKeeper = scopedTransferKeeper
	// app.ScopedWasmKeeper = scopedWasmKeeper
	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

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
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
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
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
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

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(ctx client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
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
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	// paramsKeeper.Subspace(liquiditytypes.ModuleName)
	// paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	// paramsKeeper.Subspace(ibchost.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace
	paramsKeeper.Subspace(auctiontypes.ModuleName)
	paramsKeeper.Subspace(cdptypes.ModuleName)
	paramsKeeper.Subspace(incentivetypes.ModuleName)
	paramsKeeper.Subspace(ununifidisttypes.ModuleName)
	paramsKeeper.Subspace(pricefeedtypes.ModuleName)
	paramsKeeper.Subspace(nftmarkettypes.ModuleName)
	// paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(nftminttypes.ModuleName)
	paramsKeeper.Subspace(ecosystemincentivetypes.ModuleName)
	return paramsKeeper
}
