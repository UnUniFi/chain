package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v7/modules/apps/29-fee"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	transfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	"github.com/skip-mev/pob/x/builder"
	buildertypes "github.com/skip-mev/pob/x/builder/types"

	ununifinftmodule "github.com/UnUniFi/chain/x/nft/module"

	epochs "github.com/UnUniFi/chain/x/epochs"
	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	"github.com/UnUniFi/chain/x/yieldaggregator"
	icacallbacks "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks"
	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/types"
	interchainquery "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/types"
	records "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
	stakeibc "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc"
	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"

	nftbackedloan "github.com/UnUniFi/chain/x/nftbackedloan"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"

	"github.com/UnUniFi/chain/x/derivatives"
	derivativestypes "github.com/UnUniFi/chain/x/derivatives/types"
	nftfactory "github.com/UnUniFi/chain/x/nftfactory"
	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
	"github.com/UnUniFi/chain/x/pricefeed"

	ecosystemincentive "github.com/UnUniFi/chain/x/ecosystemincentive"
	ecosystemincentivetypes "github.com/UnUniFi/chain/x/ecosystemincentive/types"

	ibctestingtypes "github.com/cosmos/ibc-go/v7/testing/types"

	appparams "github.com/UnUniFi/chain/app/params"
)

var maccPerms = map[string][]string{
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
	buildertypes.ModuleName:        nil,

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
	derivativestypes.PendingPaymentManager:  nil,

	ecosystemincentivetypes.ModuleName: nil,
}

// ModuleBasics defines the module BasicManager is in charge of setting up basic,
// non-dependant module elements, such as codec registration
// and genesis verification.
var ModuleBasics = module.NewBasicManager(
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
	ununifinftmodule.AppModuleBasic{},
	consensus.AppModuleBasic{},
	// non sdk modules
	wasm.AppModuleBasic{},
	ibc.AppModuleBasic{},
	ibctm.AppModuleBasic{},
	transfer.AppModuleBasic{},
	ica.AppModuleBasic{},
	ibcfee.AppModuleBasic{},
	builder.AppModuleBasic{},

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

func appModules(
	app *App,
	encodingConfig appparams.EncodingConfig,
	skipGenesisInvariants bool,
) []module.AppModule {
	appCodec := encodingConfig.Codec

	return []module.AppModule{
		builder.NewAppModule(appCodec, app.AppKeepers.BuilderKeeper),
		genutil.NewAppModule(
			app.AppKeepers.AccountKeeper,
			app.AppKeepers.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AppKeepers.AccountKeeper, authsims.RandomGenesisAccounts, app.AppKeepers.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
		bank.NewAppModule(appCodec, app.AppKeepers.BankKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.AppKeepers.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, &app.AppKeepers.GovKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.AppKeepers.MintKeeper, app.AppKeepers.AccountKeeper, nil, app.AppKeepers.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.AppKeepers.SlashingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.StakingKeeper, app.AppKeepers.GetSubspace(slashingtypes.ModuleName)),
		distr.NewAppModule(appCodec, app.AppKeepers.DistrKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.StakingKeeper, app.AppKeepers.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, &app.AppKeepers.StakingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.AppKeepers.GetSubspace(stakingtypes.ModuleName)),
		upgrade.NewAppModule(&app.AppKeepers.UpgradeKeeper),
		evidence.NewAppModule(app.AppKeepers.EvidenceKeeper),
		params.NewAppModule(app.AppKeepers.ParamsKeeper),
		authzmodule.NewAppModule(appCodec, app.AppKeepers.AuthzKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.interfaceRegistry),
		groupmodule.NewAppModule(appCodec, app.AppKeepers.GroupKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.interfaceRegistry),
		ununifinftmodule.NewAppModule(nftmodule.NewAppModule(appCodec, app.AppKeepers.NFTKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.interfaceRegistry), app.AppKeepers.UnUniFiNFTKeeper),
		consensus.NewAppModule(appCodec, app.AppKeepers.ConsensusParamsKeeper),

		wasm.NewAppModule(appCodec, &app.AppKeepers.WasmKeeper, app.AppKeepers.StakingKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper, app.MsgServiceRouter(), app.AppKeepers.GetSubspace(wasmtypes.ModuleName)),
		ibc.NewAppModule(app.AppKeepers.IBCKeeper),
		transfer.NewAppModule(app.AppKeepers.TransferKeeper),
		ibcfee.NewAppModule(app.AppKeepers.IBCFeeKeeper),
		ica.NewAppModule(&app.AppKeepers.ICAControllerKeeper, &app.AppKeepers.ICAHostKeeper),
		crisis.NewAppModule(&app.AppKeepers.CrisisKeeper, skipGenesisInvariants, app.AppKeepers.GetSubspace(crisistypes.ModuleName)),

		// original modules
		nftfactory.NewAppModule(appCodec, app.AppKeepers.NftfactoryKeeper, app.AppKeepers.UnUniFiNFTKeeper),
		// nftbackedloan.NewAppModule(appCodec, app.AppKeepers.NftbackedloanKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
		// ecosystemincentive.NewAppModule(appCodec, app.AppKeepers.EcosystemincentiveKeeper, app.AppKeepers.BankKeeper),

		// pricefeed.NewAppModule(appCodec, app.AppKeepers.PricefeedKeeper, app.AppKeepers.AccountKeeper),
		// derivatives.NewAppModule(appCodec, app.AppKeepers.DerivativesKeeper, app.AppKeepers.BankKeeper),

		yieldaggregator.NewAppModule(appCodec, app.AppKeepers.YieldaggregatorKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
		stakeibc.NewAppModule(appCodec, app.AppKeepers.StakeibcKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
		epochs.NewAppModule(appCodec, app.AppKeepers.EpochsKeeper),
		interchainquery.NewAppModule(appCodec, app.AppKeepers.InterchainqueryKeeper),
		records.NewAppModule(appCodec, app.AppKeepers.RecordsKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
		icacallbacks.NewAppModule(appCodec, app.AppKeepers.IcacallbacksKeeper, app.AppKeepers.AccountKeeper, app.AppKeepers.BankKeeper),
	}
}

// simulationModules returns modules for simulation manager
// define the order of the modules for deterministic simulations
func simulationModules(
	app *App,
	encodingConfig appparams.EncodingConfig,
	_ bool,
) []module.AppModuleSimulation {
	// appCodec := encodingConfig.Codec

	return []module.AppModuleSimulation{}
}

/*
orderBeginBlockers tells the app's module manager how to set the order of
BeginBlockers, which are run at the beginning of every block.

Interchain Security Requirements:
During begin block slashing happens after distr.BeginBlocker so that
there is nothing left over in the validator fee pool, so as to keep the
CanWithdrawInvariant invariant.
NOTE: staking module is required if HistoricalEntries param > 0
NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
*/

func orderBeginBlockers() []string {
	return []string{
		buildertypes.ModuleName,
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
		nftfactorytypes.ModuleName,
		// nftbackedloantypes.ModuleName,
		// ecosystemincentivetypes.ModuleName,

		// pricefeedtypes.ModuleName,
		// derivativestypes.ModuleName,

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
	}
}

/*
Interchain Security Requirements:
- provider.EndBlock gets validator updates from the staking module;
thus, staking.EndBlock must be executed before provider.EndBlock;
- creating a new consumer chain requires the following order,
CreateChildClient(), staking.EndBlock, provider.EndBlock;
thus, gov.EndBlock must be executed before staking.EndBlock
*/
func orderEndBlockers() []string {
	return []string{
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
		nftfactorytypes.ModuleName,
		// nftbackedloantypes.ModuleName,
		// ecosystemincentivetypes.ModuleName,

		// pricefeedtypes.ModuleName,
		// derivativestypes.ModuleName,

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
		buildertypes.ModuleName,
	}
}

/*
NOTE: The genutils module must occur after staking so that pools are
properly initialized with tokens from genesis accounts.
NOTE: The genutils module must also occur after auth so that it can access the params from auth.
NOTE: Capability module must occur first so that it can initialize any capabilities
so that other modules that want to create or claim capabilities afterwards in InitChain
can do so safely.
*/
func orderInitGenesis() []string {
	return []string{
		buildertypes.ModuleName,
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
		nftfactorytypes.ModuleName,
		// nftbackedloantypes.ModuleName,
		// ecosystemincentivetypes.ModuleName,

		// pricefeedtypes.ModuleName,
		// derivativestypes.ModuleName,

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
	}
}

// GetStakingKeeper implements the TestingApp interface.
func (app *App) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.AppKeepers.StakingKeeper
}

// GetTransferKeeper implements the TestingApp interface.
func (app *App) GetTransferKeeper() *ibctransferkeeper.Keeper {
	return &app.AppKeepers.TransferKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.AppKeepers.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *App) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.AppKeepers.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *App) GetTxConfig() client.TxConfig {
	cfg := MakeEncodingConfig()
	return cfg.TxConfig
}
