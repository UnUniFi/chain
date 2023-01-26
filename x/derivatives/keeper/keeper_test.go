package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	simapp "github.com/UnUniFi/chain/app"
	appparams "github.com/UnUniFi/chain/app/params"
	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *simapp.App
	addrs       []sdk.AccAddress
	queryClient types.QueryClient
	keeper      keeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	app := simapp.Setup(suite.T(), isCheckTx)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DerivativesKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	encodingConfig := appparams.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		app.GetKey(banktypes.StoreKey),
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)

	pricefeedKeeper := pricefeedkeeper.NewKeeper(
		appCodec,
		app.GetKey(pricefeedtypes.StoreKey),
		app.GetMemKey(pricefeedtypes.MemStoreKey),
		app.GetSubspace(pricefeedtypes.ModuleName),
	)
	keeper := keeper.NewKeeper(appCodec, app.GetKey(types.StoreKey), app.GetKey(types.MemStoreKey), suite.app.GetSubspace(types.ModuleName), bankKeeper, pricefeedKeeper)
	suite.keeper = keeper
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
