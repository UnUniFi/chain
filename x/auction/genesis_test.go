package auction_test

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/x/auction"
	auctiontypes "github.com/lcnem/jpyx/x/auction/types"
)

var _, testAddrs = app.GeneratePrivKeyAddressPairs(2)
var testTime = time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
var weightedAddresses, _ = auctiontypes.NewWeightedAddresses(testAddrs, []sdk.Int{sdk.OneInt(), sdk.OneInt()})
var testAuction = auctiontypes.NewCollateralAuction(
	"seller",
	c("lotdenom", 10),
	testTime,
	c("biddenom", 1000),
	weightedAddresses,
	c("debt", 1000),
).WithID(3).(*auctiontypes.CollateralAuction)
var testAuctions, _ = auctiontypes.PackGenesisAuctions(auctiontypes.GenesisAuctions{testAuction})

func TestInitGenesis(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		// setup keepers
		tApp := app.NewTestApp()
		bk := tApp.GetBankKeeper()
		keeper := tApp.GetAuctionKeeper()
		ctx := tApp.NewContext(true, tmproto.Header{})
		// setup module account
		// supplyKeeper := tApp.GetSupplyKeeper()
		accountKeeper := tApp.GetAccountKeeper()
		moduleAddr := authtypes.NewModuleAddress(auctiontypes.ModuleName)
		moduleAcc := accountKeeper.GetModuleAccount(ctx, auctiontypes.ModuleName)
		// require.NoError(t, moduleAcc.SetCoins(testAuction.GetModuleAccountCoins()))
		require.NoError(t, bk.SetBalances(ctx, moduleAddr, testAuction.GetModuleAccountCoins()))
		// supplyKeeper.SetModuleAccount(ctx, moduleAcc)
		accountKeeper.SetModuleAccount(ctx, moduleAcc)

		// create genesis
		gs := auctiontypes.NewGenesisState(
			10,
			auctiontypes.DefaultParams(),
			// auctiontypes.GenesisAuctions{testAuction},
			testAuctions,
		)
		require.Equal(t, gs.Auctions[0], testAuctions[0])

		// run init
		require.NotPanics(t, func() {
			// auction.InitGenesis(ctx, keeper, supplyKeeper, gs)
			auction.InitGenesis(ctx, keeper, accountKeeper, bk, gs)
		})

		// check state is as expected
		actualID, err := keeper.GetNextAuctionID(ctx)
		require.NoError(t, err)
		require.Equal(t, gs.NextAuctionId, actualID)

		require.Equal(t, gs.Params, keeper.GetParams(ctx))

		// TODO is there a nicer way of comparing state?
		sort.Slice(gs.Auctions, func(i, j int) bool {
			// Any型からAuction型へUnpackする
			unpackAuctions, _ := auctiontypes.UnpackGenesisAuctions(gs.Auctions)
			return unpackAuctions[i].GetID() > unpackAuctions[j].GetID()
		})
		i := 0
		keeper.IterateAuctions(ctx, func(a auctiontypes.Auction) bool {
			unpacked, _ := auctiontypes.UnpackAuction(gs.Auctions[i])
			require.Equal(t, unpacked, a)
			i++
			return false
		})
	})
	t.Run("invalid (invalid nextAuctionID)", func(t *testing.T) {
		// setup keepers
		tApp := app.NewTestApp()
		sk := tApp.GetBankKeeper()
		keeper := tApp.GetAuctionKeeper()
		ctx := tApp.NewContext(true, tmproto.Header{})
		accountKeeper := tApp.GetAccountKeeper()

		// create invalid genesis
		gs := auctiontypes.NewGenesisState(
			0, // next id < testAuction ID
			auctiontypes.DefaultParams(),
			// auctiontypes.GenesisAuctions{testAuction},
			testAuctions,
		)

		// check init fails
		require.Panics(t, func() {
			auction.InitGenesis(ctx, keeper, accountKeeper, sk, gs)
		})
	})
	t.Run("invalid (missing mod account coins)", func(t *testing.T) {
		// setup keepers
		tApp := app.NewTestApp()
		sk := tApp.GetBankKeeper()
		keeper := tApp.GetAuctionKeeper()
		ctx := tApp.NewContext(true, tmproto.Header{})
		accountKeeper := tApp.GetAccountKeeper()

		// create invalid genesis
		gs := auctiontypes.NewGenesisState(
			10,
			auctiontypes.DefaultParams(),
			// auctiontypes.GenesisAuctions{testAuction},
			testAuctions,
		)
		// invalid as there is no module account setup

		// check init fails
		require.Panics(t, func() {
			auction.InitGenesis(ctx, keeper, accountKeeper, sk, gs)
		})
	})
}

func TestExportGenesis(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// setup state
		tApp := app.NewTestApp()
		ctx := tApp.NewContext(true, tmproto.Header{})
		tApp.InitializeFromGenesisStates()

		// export
		gs := auction.ExportGenesis(ctx, tApp.GetAuctionKeeper())

		// check state matches
		// require.Equal(t, auction.DefaultGenesisState(), gs)
		require.Equal(t, auctiontypes.DefaultGenesis(), gs)
	})
	t.Run("one auction", func(t *testing.T) {
		// setup state
		tApp := app.NewTestApp()
		ctx := tApp.NewContext(true, tmproto.Header{})
		tApp.InitializeFromGenesisStates()
		tApp.GetAuctionKeeper().SetAuction(ctx, testAuction)

		// export
		gs := auction.ExportGenesis(ctx, tApp.GetAuctionKeeper())

		// check state matches
		// expectedGenesisState := auction.DefaultGenesisState()
		expectedGenesisState := auctiontypes.DefaultGenesis()
		expectedGenesisState.Auctions = append(expectedGenesisState.Auctions, testAuctions...)
		require.Equal(t, expectedGenesisState, gs)
	})
}
