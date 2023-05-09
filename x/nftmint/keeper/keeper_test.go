package keeper_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"

	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/cosmos/cosmos-sdk/x/nft/module"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	// nfttestutil "github.com/cosmos/cosmos-sdk/x/nft/testutil"
	"github.com/UnUniFi/chain/x/nftmint/types/testutil"

	simapp "github.com/UnUniFi/chain/app"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx   sdk.Context
	addrs []sdk.AccAddress
	// queryClient nft.QueryClient
	// nftKeeper     nftkeeper.Keeper
	nftKeeper     *testutil.MockNftKeeper
	nftmintKeeper keeper.Keeper
	// accountKeeper *nfttestutil.MockAccountKeeper
	accountKeeper *testutil.MockAccountKeeper

	encCfg moduletestutil.TestEncodingConfig
}

func (s *KeeperTestSuite) SetupTest() {
	// s.addrs = simtestutil.CreateIncrementalAccounts(3)
	// s.encCfg = moduletestutil.MakeTestEncodingConfig(module.AppModuleBasic{})

	// key := sdk.NewKVStoreKey(nft.StoreKey)
	// testCtx := sdktestutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	// ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	// // gomock initializations
	// ctrl := gomock.NewController(s.T())
	// accountKeeper := testutil.NewMockAccountKeeper(ctrl)
	// bankKeeper := testutil.NewMockBankKeeper(ctrl)
	// xbankKeeper := testutil.NewMockNftKeeper(ctrl)
	// accountKeeper.EXPECT().GetModuleAddress("nft").Return(s.addrs[0]).AnyTimes()

	// nftKeeper := nftkeeper.NewKeeper(key, s.encCfg.Codec, accountKeeper, bankKeeper)
	// queryHelper := baseapp.NewQueryServerTestHelper(ctx, s.encCfg.InterfaceRegistry)
	// nft.RegisterQueryServer(queryHelper, nftKeeper)

	// storeKey := sdk.NewKVStoreKey(types.StoreKey)
	// memKey := sdk.NewKVStoreKey(types.MemStoreKey)
	// app := simapp.Setup(s.T(), ([]wasm.Option{})...)
	// // encodingConfig := appparams.MakeEncodingConfig()
	// // appCodec := s.encCfg.Marshaler
	// nftmintKeeper := keeper.NewKeeper(
	// 	s.encCfg.Codec,
	// 	storeKey,
	// 	memKey,
	// 	app.GetSubspace(types.ModuleName),
	// 	accountKeeper,
	// 	nftKeeper,
	// )

	// s.nftKeeper = nftKeeper
	// s.nftmintKeeper = nftmintKeeper
	// s.queryClient = nft.NewQueryClient(queryHelper)
	// s.ctx = ctx

	s.addrs = simtestutil.CreateIncrementalAccounts(3)
	s.encCfg = moduletestutil.MakeTestEncodingConfig(module.AppModuleBasic{})

	key := sdk.NewKVStoreKey(nft.StoreKey)
	testCtx := sdktestutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	accountKeeper := testutil.NewMockAccountKeeper(ctrl)
	nftKeeper := testutil.NewMockNftKeeper(ctrl)
	// uint64c := make(chan uint64)
	// errorc := make(chan error)
	// fmt.Println("s.addrs[0]")
	// fmt.Println(s.addrs[0])
	// accountKeeper.EXPECT().GetSequence(gomock.Any(), uint(1)).Return(uint64(1), nil).AnyTimes()

	// loop addrs length
	numIncrements := 5
	// for i := 0; i < len(s.addrs); i++ {
	// 	for j := 0; j <= numIncrements; j++ {
	// 		// mockCounter.EXPECT().Increment().Return(i)
	// 		// accountKeeper.EXPECT().GetSequence(ctx, s.addrs[0]).Return(uint64(1), nil).AnyTimes()
	// 		accountKeeper.EXPECT().GetSequence(ctx, s.addrs[i]).Return(uint64(j), nil)
	// 	}
	// }
	for j := 0; j <= numIncrements; j++ {
		accountKeeper.EXPECT().GetSequence(ctx, s.addrs[0]).Return(uint64(j), nil).Times(j).AnyTimes()
	}

	// accountKeeper.EXPECT().GetSequence(ctx, s.addrs[0]).Return(uint64(1), nil).AnyTimes()
	// accountKeeper.EXPECT().GetSequence(ctx,s.addrs[0]).Return(s.addrs[0]).AnyTimes()

	// nftKeeper := nftkeeper.NewKeeper(key, s.encCfg.Codec, accountKeeper, bankKeeper)
	// queryHelper := baseapp.NewQueryServerTestHelper(ctx, s.encCfg.InterfaceRegistry)
	// nft.RegisterQueryServer(queryHelper, nftKeeper)

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memKey := sdk.NewKVStoreKey(types.MemStoreKey)
	app := simapp.Setup(s.T(), ([]wasm.Option{})...)
	// encodingConfig := appparams.MakeEncodingConfig()
	// appCodec := s.encCfg.Marshaler
	nftmintKeeper := keeper.NewKeeper(
		s.encCfg.Codec,
		storeKey,
		memKey,
		app.GetSubspace(types.ModuleName),
		accountKeeper,
		nftKeeper,
	)

	s.nftKeeper = nftKeeper
	s.nftmintKeeper = nftmintKeeper
	s.accountKeeper = accountKeeper
	// s.queryClient = nft.NewQueryClient(queryHelper)
	s.ctx = ctx
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
