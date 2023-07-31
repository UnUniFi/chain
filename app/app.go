package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	"github.com/cosmos/cosmos-sdk/client/flags"
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
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
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
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
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
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
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
	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v7/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	transfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	ibctestingtypes "github.com/cosmos/ibc-go/v7/testing/types"
	"github.com/spf13/cast"

	"github.com/UnUniFi/chain/wasmbinding"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	epochs "github.com/UnUniFi/chain/x/epochs"
	epochskeeper "github.com/UnUniFi/chain/x/epochs/keeper"
	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	"github.com/UnUniFi/chain/x/yieldaggregator"
	icacallbacks "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/icacallbacks"
	icacallbackskeeper "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/icacallbacks/keeper"
	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/icacallbacks/types"
	interchainquery "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery"
	interchainquerykeeper "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/keeper"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/types"
	records "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/records"
	recordskeeper "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/records/keeper"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/records/types"
	stakeibc "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc"
	stakeibckeeper "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/keeper"
	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/types"
	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"

	nftbackedloan "github.com/UnUniFi/chain/x/nftbackedloan"
	nftbackedloankeeper "github.com/UnUniFi/chain/x/nftbackedloan/keeper"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"

	"github.com/UnUniFi/chain/x/derivatives"
	derivativeskeeper "github.com/UnUniFi/chain/x/derivatives/keeper"
	derivativestypes "github.com/UnUniFi/chain/x/derivatives/types"
	nftfactory "github.com/UnUniFi/chain/x/nftfactory"
	nftfactorykeeper "github.com/UnUniFi/chain/x/nftfactory/keeper"
	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
	"github.com/UnUniFi/chain/x/pricefeed"
	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	ecosystemincentive "github.com/UnUniFi/chain/x/ecosystemincentive"
	ecosystemincentivekeeper "github.com/UnUniFi/chain/x/ecosystemincentive/keeper"
	ecosystemincentivetypes "github.com/UnUniFi/chain/x/ecosystemincentive/types"

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

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
				upgradeclient.LegacyProposalHandler,
				upgradeclient.LegacyCancelProposalHandler,
				ibcclientclient.UpdateClientProposalHandler,
				ibcclientclient.UpgradeProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		consensus.AppModuleBasic{},
		// non sdk modules
		wasm.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		transfer.AppModuleBasic{},
		ica.AppModuleBasic{},
		ibcfee.AppModuleBasic{},

		// original modules
		pricefeed.AppModuleBasic{},
		derivatives.AppModuleBasic{},

		nftbackedloan.AppModuleBasic{},
		nftfactory.AppModuleBasic{},
		ecosystemincentive.AppModuleBasic{},

		yieldaggregator.AppModuleBasic{},
		stakeibc.AppModuleBasic{},
		epochs.AppModuleBasic{},
		interchainquery.AppModuleBasic{},
		records.AppModuleBasic{},
		icacallbacks.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		nft.ModuleName:                 nil,
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		ibcfeetypes.ModuleName:         nil,
		icatypes.ModuleName:            nil,
		wasm.ModuleName:                {authtypes.Burner},

		// original modules
		nftbackedloantypes.ModuleName: nil,
		// nftbackedloantypes.NftTradingFee: nil,
		nftfactorytypes.ModuleName: nil,

		yieldaggregatortypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		stakeibctypes.ModuleName:        {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		interchainquerytypes.ModuleName: nil,

		derivativestypes.ModuleName:             {authtypes.Minter, authtypes.Burner},
		derivativestypes.DerivativeFeeCollector: nil,
		derivativestypes.MarginManager:          nil,

		ecosystemincentivetypes.ModuleName: nil,
	}

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
	legacyAmino *codec.LegacyAmino
	appCodec    codec.Codec
	// txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.BaseKeeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             govkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	GroupKeeper           groupkeeper.Keeper
	NFTKeeper             nftkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper

	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper
	WasmKeeper          wasm.Keeper

	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedIBCFeeKeeper        capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper

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

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator

	keepers.AppKeepers

	// msgSvcRouter      *authmiddleware.MsgServiceRouter
	// legacyRouter      sdk.Router
}

// NewApp returns a reference to an initialized Gaia.
// NewSimApp returns a reference to an initialized SimApp.
func NewApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	enabledProposals []wasm.ProposalType,
	appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	encodingConfig := MakeEncodingConfig()

	appCodec, legacyAmino := encodingConfig.Marshaler, encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(Name, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey, crisistypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, consensusparamtypes.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
		evidencetypes.StoreKey, capabilitytypes.StoreKey,
		authzkeeper.StoreKey, nftkeeper.StoreKey, group.StoreKey,
		// non sdk store keys
		ibcexported.StoreKey, ibctransfertypes.StoreKey, ibcfeetypes.StoreKey,
		wasm.StoreKey, icahosttypes.StoreKey,
		icacontrollertypes.StoreKey,

		// original modules
		derivativestypes.StoreKey,
		pricefeedtypes.StoreKey,

		nftbackedloantypes.StoreKey,
		nftfactorytypes.StoreKey,
		ecosystemincentivetypes.StoreKey,

		yieldaggregatortypes.StoreKey,
		stakeibctypes.StoreKey,
		epochstypes.StoreKey,
		interchainquerytypes.StoreKey,
		recordstypes.StoreKey,
		icacallbackstypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, logger, keys); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	app := &App{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
		AppKeepers:        keepers.AppKeepers{},
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, keys[consensusparamtypes.StoreKey], authtypes.NewModuleAddress(govtypes.ModuleName).String())
	bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		AccountAddressPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.BlockedAddrs(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.StakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		legacyAmino,
		keys[slashingtypes.StoreKey],
		app.StakingKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	invCheckPeriod := cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks(),
		),
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(keys[authzkeeper.StoreKey], appCodec, app.MsgServiceRouter(), app.AccountKeeper)

	groupConfig := group.DefaultConfig()
	app.GroupKeeper = groupkeeper.NewKeeper(keys[group.StoreKey], appCodec, app.MsgServiceRouter(), app.AccountKeeper, groupConfig)

	// get skipUpgradeHeights from the app options
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	// set the governance module account as the authority for conducting upgrades
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibcexported.StoreKey],
		app.GetSubspace(ibcexported.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.MsgServiceRouter(),
		govConfig,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	app.NFTKeeper = nftkeeper.NewKeeper(
		keys[nftkeeper.StoreKey],
		appCodec,
		app.AccountKeeper,
		app.BankKeeper,
	)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		app.StakingKeeper,
		app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	// IBC Fee Module keeper
	app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		appCodec, keys[ibcfeetypes.StoreKey],
		app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper, app.AccountKeeper, app.BankKeeper,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCFeeKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)

	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
	)
	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec,
		keys[icacontrollertypes.StoreKey],
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCFeeKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper,
		app.MsgServiceRouter(),
	)

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	availableCapabilities := "iterator,staking,stargate,cosmwasm_1_1,cosmwasm_1_2"

	wasmOpts = append(wasmbinding.RegisterCustomPlugins(&app.BankKeeper, &app.InterchainqueryKeeper, &app.RecordsKeeper), wasmOpts...)

	app.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		distrkeeper.NewQuerier(app.DistrKeeper),
		app.IBCFeeKeeper, // ISC4 Wrapper: fee IBC middleware
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		availableCapabilities,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		wasmOpts...,
	)

	app.PricefeedKeeper = pricefeedkeeper.NewKeeper(
		appCodec,
		keys[pricefeedtypes.StoreKey],
		keys[pricefeedtypes.MemStoreKey],
		app.GetSubspace(pricefeedtypes.ModuleName),
		app.BankKeeper,
	)

	app.DerivativesKeeper = derivativeskeeper.NewKeeper(
		appCodec,
		keys[derivativestypes.StoreKey],
		keys[derivativestypes.MemStoreKey],
		app.GetSubspace(derivativestypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.PricefeedKeeper,
		app.NftfactoryKeeper,
	)

	app.NftfactoryKeeper = nftfactorykeeper.NewKeeper(
		appCodec,
		keys[nftfactorytypes.StoreKey],
		keys[nftfactorytypes.MemStoreKey],
		app.GetSubspace(nftfactorytypes.ModuleName),
		app.AccountKeeper,
		app.NFTKeeper,
	)

	nftbackedloanKeeper := nftbackedloankeeper.NewKeeper(
		appCodec,
		encodingConfig.TxConfig,
		keys[nftbackedloantypes.StoreKey],
		keys[nftbackedloantypes.MemStoreKey],
		app.GetSubspace(nftbackedloantypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.NFTKeeper,
	)

	app.EcosystemincentiveKeeper = ecosystemincentivekeeper.NewKeeper(
		appCodec,
		keys[ecosystemincentivetypes.StoreKey],
		app.GetSubspace(ecosystemincentivetypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.DistrKeeper,
		// same as the feeCollectorName in the distribution module
		authtypes.FeeCollectorName,
	)

	// create Keeper objects which have Hooks
	app.NftbackedloanKeeper = *nftbackedloanKeeper.SetHooks(nftbackedloantypes.NewMultiNftbackedloanHooks(app.EcosystemincentiveKeeper.Hooks()))

	scopedIcacallbacksKeeper := app.CapabilityKeeper.ScopeToModule(icacallbackstypes.ModuleName)
	app.ScopedIcacallbacksKeeper = scopedIcacallbacksKeeper
	app.IcacallbacksKeeper = *icacallbackskeeper.NewKeeper(
		appCodec,
		keys[icacallbackstypes.StoreKey],
		keys[icacallbackstypes.MemStoreKey],
		app.GetSubspace(icacallbackstypes.ModuleName),
		scopedIcacallbacksKeeper,
		*app.IBCKeeper,
		app.ICAControllerKeeper,
	)

	app.InterchainqueryKeeper = interchainquerykeeper.NewKeeper(
		appCodec,
		keys[interchainquerytypes.StoreKey],
		app.IBCKeeper,
		&app.WasmKeeper,
	)

	scopedRecordsKeeper := app.CapabilityKeeper.ScopeToModule(recordstypes.ModuleName)
	app.ScopedRecordsKeeper = scopedRecordsKeeper
	app.RecordsKeeper = *recordskeeper.NewKeeper(
		appCodec,
		keys[recordstypes.StoreKey],
		keys[recordstypes.MemStoreKey],
		app.GetSubspace(recordstypes.ModuleName),
		scopedRecordsKeeper,
		app.AccountKeeper,
		app.TransferKeeper,
		*app.IBCKeeper,
		app.IcacallbacksKeeper,
		&app.WasmKeeper,
	)

	scopedStakeibcKeeper := app.CapabilityKeeper.ScopeToModule(stakeibctypes.ModuleName)
	app.ScopedStakeibcKeeper = scopedStakeibcKeeper
	app.StakeibcKeeper = stakeibckeeper.NewKeeper(
		appCodec,
		keys[stakeibctypes.StoreKey],
		keys[stakeibctypes.MemStoreKey],
		app.GetSubspace(stakeibctypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		// &app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.ICAControllerKeeper,
		*app.IBCKeeper,
		scopedStakeibcKeeper,
		scopedIBCKeeper,
		app.InterchainqueryKeeper,
		app.RecordsKeeper,
		*app.StakingKeeper,
		app.IcacallbacksKeeper,
	)

	// Register ICQ callbacks
	err = app.InterchainqueryKeeper.SetCallbackHandler(stakeibctypes.ModuleName, app.StakeibcKeeper.CallbackHandler())
	if err != nil {
		return nil
	}

	// Register ICA calllbacks
	// stakeibc
	err = app.IcacallbacksKeeper.SetICACallbackHandler(icacontrollertypes.SubModuleName, app.StakeibcKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}
	// records
	err = app.IcacallbacksKeeper.SetICACallbackHandler(recordstypes.ModuleName, app.RecordsKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}

	app.YieldaggregatorKeeper = yieldaggregatorkeeper.NewKeeper(
		appCodec,
		keys[yieldaggregatortypes.StoreKey],
		app.GetSubspace(yieldaggregatortypes.ModuleName),
		app.BankKeeper,
		wasmkeeper.NewDefaultPermissionKeeper(app.WasmKeeper),
		app.WasmKeeper,
		app.StakeibcKeeper,
		app.RecordsKeeper,
	)

	epochsKeeper := epochskeeper.NewKeeper(appCodec, keys[epochstypes.StoreKey])
	app.EpochsKeeper = *epochsKeeper.SetHooks(
		epochstypes.NewMultiEpochHooks(
			app.YieldaggregatorKeeper.Hooks(),
			app.StakeibcKeeper.Hooks(),
		),
	)

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://docs.cosmos.network/main/modules/gov#proposal-messages
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(yieldaggregatortypes.RouterKey, yieldaggregator.NewYieldAggregatorProposalHandler(app.YieldaggregatorKeeper))

	// The gov proposal types can be individually enabled
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals))
	}

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://docs.cosmos.network/main/modules/gov#proposal-messages
	app.GovKeeper.SetLegacyRouter(govRouter)

	// Create Transfer Stack
	var transferStack porttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.TransferKeeper)
	transferStack = records.NewIBCModule(app.RecordsKeeper, transferStack)
	transferStack = ibcfee.NewIBCMiddleware(transferStack, app.IBCFeeKeeper)

	// RecvPacket, message that originates from core IBC and goes down to app, the flow is:
	// channel.RecvPacket -> fee.OnRecvPacket -> icaHost.OnRecvPacket
	var icaHostStack porttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(app.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, app.IBCFeeKeeper)

	// Create fee enabled wasm ibc Stack
	var wasmStack porttypes.IBCModule
	wasmStack = wasm.NewIBCHandler(app.WasmKeeper, app.IBCKeeper.ChannelKeeper, app.IBCFeeKeeper)
	wasmStack = ibcfee.NewIBCMiddleware(wasmStack, app.IBCFeeKeeper)

	// Stack two (Stakeibc Stack) contains
	// - IBC
	// - ICA
	// - stakeibc
	// - base app
	var stakeibcStack porttypes.IBCModule = stakeibc.NewIBCModule(app.StakeibcKeeper)
	stakeibcStack = icacontroller.NewIBCMiddleware(stakeibcStack, app.ICAControllerKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.
		AddRoute(ibctransfertypes.ModuleName, transferStack).
		AddRoute(wasm.ModuleName, wasmStack).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		// Stakeibc Stack
		AddRoute(icacontrollertypes.SubModuleName, stakeibcStack).
		AddRoute(stakeibctypes.ModuleName, stakeibcStack)

	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	app.setupAppkeeper()

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	// tbbbt
	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, &app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName)),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		groupmodule.NewAppModule(appCodec, app.GroupKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		nftmodule.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),

		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.MsgServiceRouter(), app.GetSubspace(wasmtypes.ModuleName)),
		ibc.NewAppModule(app.IBCKeeper),
		transfer.NewAppModule(app.TransferKeeper),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),

		// original modules
		pricefeed.NewAppModule(appCodec, app.PricefeedKeeper, app.AccountKeeper),
		derivatives.NewAppModule(appCodec, app.DerivativesKeeper, app.BankKeeper),

		nftfactory.NewAppModule(appCodec, app.NftfactoryKeeper, app.NFTKeeper),
		nftbackedloan.NewAppModule(appCodec, app.NftbackedloanKeeper, app.AccountKeeper, app.BankKeeper),
		ecosystemincentive.NewAppModule(appCodec, app.EcosystemincentiveKeeper, app.BankKeeper),

		yieldaggregator.NewAppModule(appCodec, app.YieldaggregatorKeeper, app.AccountKeeper, app.BankKeeper),
		stakeibc.NewAppModule(appCodec, app.StakeibcKeeper, app.AccountKeeper, app.BankKeeper),
		epochs.NewAppModule(appCodec, app.EpochsKeeper),
		interchainquery.NewAppModule(appCodec, app.InterchainqueryKeeper),
		records.NewAppModule(appCodec, app.RecordsKeeper, app.AccountKeeper, app.BankKeeper),
		icacallbacks.NewAppModule(appCodec, app.IcacallbacksKeeper, app.AccountKeeper, app.BankKeeper),
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
		group.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		consensusparamtypes.ModuleName,
		// original modules
		pricefeedtypes.ModuleName,
		derivativestypes.ModuleName,

		nftfactorytypes.ModuleName,
		nftbackedloantypes.ModuleName,
		ecosystemincentivetypes.ModuleName,

		stakeibctypes.ModuleName,
		epochstypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordstypes.ModuleName,
		icacallbackstypes.ModuleName,

		yieldaggregatortypes.ModuleName,

		// additional non simd modules
		ibctransfertypes.ModuleName,
		ibcexported.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		wasm.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
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
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		consensusparamtypes.ModuleName,
		// original modules
		pricefeedtypes.ModuleName,
		derivativestypes.ModuleName,

		nftfactorytypes.ModuleName,
		nftbackedloantypes.ModuleName,
		ecosystemincentivetypes.ModuleName,

		stakeibctypes.ModuleName,
		epochstypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordstypes.ModuleName,
		icacallbackstypes.ModuleName,

		yieldaggregatortypes.ModuleName,

		// additional non simd modules
		ibctransfertypes.ModuleName,
		ibcexported.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		wasm.ModuleName,
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
		authz.ModuleName,
		feegrant.ModuleName,
		nft.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		consensusparamtypes.ModuleName,

		// original modules
		pricefeedtypes.ModuleName,
		derivativestypes.ModuleName,

		nftfactorytypes.ModuleName,
		nftbackedloantypes.ModuleName,
		ecosystemincentivetypes.ModuleName,

		stakeibctypes.ModuleName,
		epochstypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordstypes.ModuleName,
		icacallbackstypes.ModuleName,

		yieldaggregatortypes.ModuleName,

		// additional non simd modules
		ibctransfertypes.ModuleName,
		ibcexported.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// wasm after ibc transfer
		wasm.ModuleName,
	)

	app.mm.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))

	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.setAnteHandler(encodingConfig.TxConfig, wasmConfig, keys[wasm.StoreKey])

	app.setupUpgradeHandlers()
	app.setupUpgradeStoreLoaders()

	// must be before Loading version
	// requires the snapshot store to be created and registered as a BaseAppOption
	// see cmd/wasmd/root.go: 206 - 214 approx
	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.WasmKeeper),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %s", err))
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedWasmKeeper = scopedWasmKeeper
	app.ScopedICAHostKeeper = scopedICAHostKeeper
	app.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.ScopedStakeibcKeeper = scopedStakeibcKeeper

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
		if err := app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
	}

	return app
}

func (app *App) setAnteHandler(txConfig client.TxConfig, wasmConfig wasmtypes.WasmConfig, txCounterStoreKey storetypes.StoreKey) {
	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				SignModeHandler: txConfig.SignModeHandler(),
				FeegrantKeeper:  app.FeeGrantKeeper,
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCKeeper:         app.IBCKeeper,
			WasmConfig:        &wasmConfig,
			TXCounterStoreKey: txCounterStoreKey,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}
	app.SetAnteHandler(anteHandler)
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

// GetStakingKeeper implements the TestingApp interface.
func (app *App) GetStakingKeeper() ibctestingtypes.StakingKeeper {
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

func (app *App) Configurator() module.Configurator {
	return app.configurator
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
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

	// original modules
	paramsKeeper.Subspace(pricefeedtypes.ModuleName)
	paramsKeeper.Subspace(derivativestypes.ModuleName)

	paramsKeeper.Subspace(nftfactorytypes.ModuleName)
	paramsKeeper.Subspace(nftbackedloantypes.ModuleName)
	paramsKeeper.Subspace(ecosystemincentivetypes.ModuleName)

	paramsKeeper.Subspace(stakeibctypes.ModuleName)
	paramsKeeper.Subspace(epochstypes.ModuleName)
	paramsKeeper.Subspace(interchainquerytypes.ModuleName)
	paramsKeeper.Subspace(recordstypes.ModuleName)
	paramsKeeper.Subspace(icacallbackstypes.ModuleName)
	paramsKeeper.Subspace(yieldaggregatortypes.ModuleName)

	return paramsKeeper
}

func (app *App) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades))
		}
	}

}

func (app *App) setupUpgradeHandlers() {
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
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

func (app *App) setupAppkeeper() {
	app.AppKeepers.AccountKeeper = &app.AccountKeeper
	app.AppKeepers.BankKeeper = &app.BankKeeper
	app.AppKeepers.ConsensusParamsKeeper = &app.ConsensusParamsKeeper
	app.AppKeepers.ParamsKeeper = &app.ParamsKeeper
	app.AppKeepers.WasmKeeper = &app.WasmKeeper
	app.AppKeepers.EpochsKeeper = &app.EpochsKeeper
	app.AppKeepers.YieldaggregatorKeeper = &app.YieldaggregatorKeeper
	app.AppKeepers.PricefeedKeeper = &app.PricefeedKeeper
	app.AppKeepers.DerivativesKeeper = &app.DerivativesKeeper
}
