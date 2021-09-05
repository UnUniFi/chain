package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/lcnem/jpyx/app"
	auctiontypes "github.com/lcnem/jpyx/x/auction/types"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"
)

func TestSurplusAuctionBasic(t *testing.T) {
	// Setup
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	buyer := addrs[0]
	sellerModName := cdptypes.LiquidatorMacc
	sellerAddr := authtypes.NewModuleAddress(sellerModName)

	tApp := app.NewTestApp()
	bk := tApp.GetBankKeeper()

	sellerAcc := authtypes.NewEmptyModuleAccount(sellerModName, authtypes.Burner) // forward auctions burn proceeds
	tApp.InitializeFromGenesisStates(
		// NewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	authtypes.NewBaseAccount(buyer, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	sellerAcc,
		// }),
		app.NewAuthGenState(tApp, []sdk.AccAddress{buyer}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{sellerAcc}),
	)
	ctx := tApp.NewContext(false, tmproto.Header{})
	require.NoError(t, bk.SetBalances(ctx, sellerAddr, cs(c("token1", 100), c("token2", 100))))

	keeper := tApp.GetAuctionKeeper()

	// Create an auction (lot: 20 token1, initialBid: 0 token2)
	auctionID, err := keeper.StartSurplusAuction(ctx, sellerModName, c("token1", 20), "token2") // lot, bid denom
	require.NoError(t, err)
	// Check seller's coins have decreased
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 100)))

	// PlaceBid (bid: 10 token, lot: same as starting)
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, buyer, c("token2", 10)))
	// Check buyer's coins have decreased
	tApp.CheckBalance(t, ctx, buyer, cs(c("token1", 100), c("token2", 90)))
	// Check seller's coins have not increased (because proceeds are burned)
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 100)))

	// increment bid same bidder
	err = keeper.PlaceBid(ctx, auctionID, buyer, c("token2", 20))
	require.NoError(t, err)

	// Close auction at just at auction expiry time
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(auctiontypes.DefaultBidDuration))
	require.NoError(t, keeper.CloseAuction(ctx, auctionID))
	// Check buyer's coins increased
	tApp.CheckBalance(t, ctx, buyer, cs(c("token1", 120), c("token2", 80)))
}

func TestDebtAuctionBasic(t *testing.T) {
	// Setup
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	seller := addrs[0]
	buyerModName := cdptypes.LiquidatorMacc
	buyerAddr := authtypes.NewModuleAddress(buyerModName)

	tApp := app.NewTestApp()
	bk := tApp.GetBankKeeper()

	buyerAcc := authtypes.NewEmptyModuleAccount(buyerModName, authtypes.Minter) // reverse auctions mint payout
	tApp.InitializeFromGenesisStates(
		// zNewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	authtypes.NewBaseAccount(seller, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	buyerAcc,
		// }),
		app.NewAuthGenState(tApp, []sdk.AccAddress{seller}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{buyerAcc}),
	)
	ctx := tApp.NewContext(false, tmproto.Header{})
	require.NoError(t, bk.SetBalances(ctx, buyerAddr, cs(c("debt", 100))))

	keeper := tApp.GetAuctionKeeper()

	// Start auction
	auctionID, err := keeper.StartDebtAuction(ctx, buyerModName, c("token1", 20), c("token2", 99999), c("debt", 20))
	require.NoError(t, err)
	// Check buyer's coins have not decreased (except for debt), as lot is minted at the end
	tApp.CheckBalance(t, ctx, buyerAddr, cs(c("debt", 80)))

	// Place a bid
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, seller, c("token2", 10)))
	// Check seller's coins have decreased
	tApp.CheckBalance(t, ctx, seller, cs(c("token1", 80), c("token2", 100)))
	// Check buyer's coins have increased
	tApp.CheckBalance(t, ctx, buyerAddr, cs(c("token1", 20), c("debt", 100)))

	// Close auction at just after auction expiry
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(auctiontypes.DefaultBidDuration))
	require.NoError(t, keeper.CloseAuction(ctx, auctionID))
	// Check seller's coins increased
	tApp.CheckBalance(t, ctx, seller, cs(c("token1", 80), c("token2", 110)))
}

func TestDebtAuctionDebtRemaining(t *testing.T) {
	// Setup
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	seller := addrs[0]
	buyerModName := cdptypes.LiquidatorMacc
	buyerAddr := authtypes.NewModuleAddress(buyerModName)

	tApp := app.NewTestApp()
	bk := tApp.GetBankKeeper()

	buyerAcc := authtypes.NewEmptyModuleAccount(buyerModName, authtypes.Minter) // reverse auctions mint payout
	tApp.InitializeFromGenesisStates(
		// NewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	authtypes.NewBaseAccount(seller, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	buyerAcc,
		// }),
		// app.NewAuthGenState(tApp, []sdk.AccAddress{seller, buyerAddr}, []sdk.Coins{cs(c("token1", 100), c("token2", 100)), {}}),
		app.NewAuthGenState(tApp, []sdk.AccAddress{seller}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{buyerAcc}),
	)
	ctx := tApp.NewContext(false, tmproto.Header{})
	require.NoError(t, bk.SetBalances(ctx, buyerAddr, cs(c("debt", 100))))

	keeper := tApp.GetAuctionKeeper()

	// Start auction
	auctionID, err := keeper.StartDebtAuction(ctx, buyerModName, c("token1", 10), c("token2", 99999), c("debt", 20))
	require.NoError(t, err)
	// Check buyer's coins have not decreased (except for debt), as lot is minted at the end
	tApp.CheckBalance(t, ctx, buyerAddr, cs(c("debt", 80)))

	// Place a bid
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, seller, c("token2", 10)))
	// Check seller's coins have decreased
	tApp.CheckBalance(t, ctx, seller, cs(c("token1", 90), c("token2", 100)))
	// Check buyer's coins have increased
	tApp.CheckBalance(t, ctx, buyerAddr, cs(c("token1", 10), c("debt", 90)))

	// Close auction at just after auction expiry
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(auctiontypes.DefaultBidDuration))
	require.NoError(t, keeper.CloseAuction(ctx, auctionID))
	// Check seller's coins increased
	tApp.CheckBalance(t, ctx, seller, cs(c("token1", 90), c("token2", 110)))
	// check that debt has increased due to corresponding debt being greater than bid
	tApp.CheckBalance(t, ctx, buyerAddr, cs(c("token1", 10), c("debt", 100)))
}

func TestCollateralAuctionBasic(t *testing.T) {
	// Setup
	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	buyer := addrs[0]
	returnAddrs := addrs[1:]
	returnWeights := is(30, 20, 10)
	sellerModName := cdptypes.LiquidatorMacc
	sellerAddr := authtypes.NewModuleAddress(sellerModName)

	tApp := app.NewTestApp()
	bk := tApp.GetBankKeeper()

	sellerAcc := authtypes.NewEmptyModuleAccount(sellerModName)
	tApp.InitializeFromGenesisStates(
		// NewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	authtypes.NewBaseAccount(buyer, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	authtypes.NewBaseAccount(returnAddrs[0], cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	authtypes.NewBaseAccount(returnAddrs[1], cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	authtypes.NewBaseAccount(returnAddrs[2], cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	sellerAcc,
		// }),
		app.NewAuthGenState(tApp, []sdk.AccAddress{buyer}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{sellerAcc}),
		app.NewAuthGenState(tApp, []sdk.AccAddress{returnAddrs[0]}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenState(tApp, []sdk.AccAddress{returnAddrs[1]}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenState(tApp, []sdk.AccAddress{returnAddrs[2]}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
	)
	ctx := tApp.NewContext(false, tmproto.Header{})
	require.NoError(t, bk.SetBalances(ctx, sellerAddr, cs(c("token1", 100), c("token2", 100), c("debt", 100))))

	keeper := tApp.GetAuctionKeeper()

	// Start auction
	auctionID, err := keeper.StartCollateralAuction(ctx, sellerModName, c("token1", 20), c("token2", 50), returnAddrs, returnWeights, c("debt", 40))
	require.NoError(t, err)
	// Check seller's coins have decreased
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 100), c("debt", 60)))

	// Place a forward bid
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, buyer, c("token2", 10)))
	// Check bidder's coins have decreased
	tApp.CheckBalance(t, ctx, buyer, cs(c("token1", 100), c("token2", 90)))
	// Check seller's coins have increased
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 110), c("debt", 70)))
	// Check return addresses have not received coins
	for _, ra := range returnAddrs {
		tApp.CheckBalance(t, ctx, ra, cs(c("token1", 100), c("token2", 100)))
	}

	// Place a reverse bid
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, buyer, c("token2", 50))) // first bid up to max bid to switch phases
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, buyer, c("token1", 15)))
	// Check bidder's coins have decreased
	tApp.CheckBalance(t, ctx, buyer, cs(c("token1", 100), c("token2", 50)))
	// Check seller's coins have increased
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 150), c("debt", 100)))
	// Check return addresses have received coins
	tApp.CheckBalance(t, ctx, returnAddrs[0], cs(c("token1", 102), c("token2", 100)))
	tApp.CheckBalance(t, ctx, returnAddrs[1], cs(c("token1", 102), c("token2", 100)))
	tApp.CheckBalance(t, ctx, returnAddrs[2], cs(c("token1", 101), c("token2", 100)))

	// Close auction at just after auction expiry
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(auctiontypes.DefaultBidDuration))
	require.NoError(t, keeper.CloseAuction(ctx, auctionID))
	// Check buyer's coins increased
	tApp.CheckBalance(t, ctx, buyer, cs(c("token1", 115), c("token2", 50)))
}

func TestCollateralAuctionDebtRemaining(t *testing.T) {
	// Setup
	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	buyer := addrs[0]
	returnAddrs := addrs[1:]
	returnWeights := is(30, 20, 10)
	sellerModName := cdptypes.LiquidatorMacc
	sellerAddr := authtypes.NewModuleAddress(sellerModName)

	tApp := app.NewTestApp()
	bk := tApp.GetBankKeeper()
	sellerAcc := authtypes.NewEmptyModuleAccount(sellerModName)
	// require.NoError(t, sellerAcc.SetCoins(cs(c("token1", 100), c("token2", 100), c("debt", 100))))
	tApp.InitializeFromGenesisStates(
		// NewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	authtypes.NewBaseAccount(buyer, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	authtypes.NewBaseAccount(returnAddrs[0], cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	authtypes.NewBaseAccount(returnAddrs[1], cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	authtypes.NewBaseAccount(returnAddrs[2], cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	sellerAcc,
		// }),
		app.NewAuthGenState(tApp, []sdk.AccAddress{buyer}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{sellerAcc}),
		app.NewAuthGenState(tApp, []sdk.AccAddress{returnAddrs[0]}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenState(tApp, []sdk.AccAddress{returnAddrs[1]}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenState(tApp, []sdk.AccAddress{returnAddrs[2]}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
	)
	ctx := tApp.NewContext(false, tmproto.Header{})
	require.NoError(t, bk.SetBalances(ctx, sellerAddr, cs(c("token1", 100), c("token2", 100), c("debt", 100))))

	keeper := tApp.GetAuctionKeeper()

	// Start auction
	auctionID, err := keeper.StartCollateralAuction(ctx, sellerModName, c("token1", 20), c("token2", 50), returnAddrs, returnWeights, c("debt", 40))
	require.NoError(t, err)
	// Check seller's coins have decreased
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 100), c("debt", 60)))

	// Place a forward bid
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, buyer, c("token2", 10)))
	// Check bidder's coins have decreased
	tApp.CheckBalance(t, ctx, buyer, cs(c("token1", 100), c("token2", 90)))
	// Check seller's coins have increased
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 110), c("debt", 70)))
	// Check return addresses have not received coins
	for _, ra := range returnAddrs {
		tApp.CheckBalance(t, ctx, ra, cs(c("token1", 100), c("token2", 100)))
	}
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(auctiontypes.DefaultBidDuration))
	require.NoError(t, keeper.CloseAuction(ctx, auctionID))

	// check that buyers coins have increased
	tApp.CheckBalance(t, ctx, buyer, cs(c("token1", 120), c("token2", 90)))
	// Check return addresses have not received coins
	for _, ra := range returnAddrs {
		tApp.CheckBalance(t, ctx, ra, cs(c("token1", 100), c("token2", 100)))
	}
	// check that token2 has increased by 10, debt by 40, for a net debt increase of 30 debt
	tApp.CheckBalance(t, ctx, sellerAddr, cs(c("token1", 80), c("token2", 110), c("debt", 100)))
}

func TestStartSurplusAuction(t *testing.T) {
	someTime := time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC)
	type args struct {
		seller   string
		lot      sdk.Coin
		bidDenom string
	}
	testCases := []struct {
		name       string
		blockTime  time.Time
		args       args
		expectPass bool
		expPanic   bool
	}{
		{
			"normal",
			someTime,
			args{cdptypes.LiquidatorMacc, c("stable", 10), "gov"},
			true, false,
		},
		{
			"no module account",
			someTime,
			args{"nonExistentModule", c("stable", 10), "gov"},
			false, true,
		},
		{
			"not enough coins",
			someTime,
			args{cdptypes.LiquidatorMacc, c("stable", 101), "gov"},
			false, false,
		},
		{
			"incorrect denom",
			someTime,
			args{cdptypes.LiquidatorMacc, c("notacoin", 10), "gov"},
			false, false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			initialLiquidatorCoins := cs(c("stable", 100))
			tApp := app.NewTestApp()
			sk := tApp.GetBankKeeper()

			liqAddr := authtypes.NewModuleAddress(cdptypes.LiquidatorMacc)
			liqAcc := authtypes.NewEmptyModuleAccount(cdptypes.LiquidatorMacc, authtypes.Burner)
			tApp.InitializeFromGenesisStates(
				NewAuthGenStateFromAccs(tApp, authtypes.GenesisAccounts{liqAcc}),
			)
			ctx := tApp.NewContext(false, tmproto.Header{}).WithBlockTime(tc.blockTime)
			require.NoError(t, sk.SetBalances(ctx, liqAddr, initialLiquidatorCoins))

			keeper := tApp.GetAuctionKeeper()

			// run function under test
			var (
				id  uint64
				err error
			)
			if tc.expPanic {
				require.Panics(t, func() { _, _ = keeper.StartSurplusAuction(ctx, tc.args.seller, tc.args.lot, tc.args.bidDenom) }, tc.name)
			} else {
				id, err = keeper.StartSurplusAuction(ctx, tc.args.seller, tc.args.lot, tc.args.bidDenom)
			}

			// check
			// sk := tApp.GetauthtypesKeeper()
			// liquidatorCoins := sk.GetModuleAccount(ctx, cdptypes.LiquidatorMacc).GetCoins()
			liquidatorCoins := sk.GetAllBalances(ctx, liqAddr)
			actualAuc, found := keeper.GetAuction(ctx, id)

			if tc.expectPass {
				require.NoError(t, err, tc.name)
				// check coins moved
				require.Equal(t, initialLiquidatorCoins.Sub(cs(tc.args.lot)), liquidatorCoins, tc.name)
				// check auction in store and is correct
				require.True(t, found, tc.name)
				expectedAuction := auctiontypes.SurplusAuction{BaseAuction: auctiontypes.BaseAuction{
					Id:              id,
					Initiator:       tc.args.seller,
					Lot:             tc.args.lot,
					Bidder:          nil,
					Bid:             c(tc.args.bidDenom, 0),
					HasReceivedBids: false,
					EndTime:         auctiontypes.DistantFuture,
					MaxEndTime:      auctiontypes.DistantFuture,
				}}
				require.Equal(t, &expectedAuction, actualAuc, tc.name)
			} else if !tc.expPanic && !tc.expectPass {
				require.Error(t, err, tc.name)
				// check coins not moved
				require.Equal(t, initialLiquidatorCoins, liquidatorCoins, tc.name)
				// check auction not in store
				require.False(t, found, tc.name)
			}
		})
	}
}

func TestCloseAuction(t *testing.T) {
	// Set up
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	buyer := addrs[0]
	sellerModName := cdptypes.LiquidatorMacc
	sellerAddr := authtypes.NewModuleAddress(sellerModName)

	tApp := app.NewTestApp()
	bk := tApp.GetBankKeeper()

	sellerAcc := authtypes.NewEmptyModuleAccount(sellerModName, authtypes.Burner) // forward auctions burn proceeds
	// require.NoError(t, sellerAcc.SetCoins(cs(c("token1", 100), c("token2", 100))))
	tApp.InitializeFromGenesisStates(
		// NewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	authtypes.NewBaseAccount(buyer, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	sellerAcc,
		// }),
		app.NewAuthGenState(tApp, []sdk.AccAddress{buyer}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{sellerAcc}),
	)
	ctx := tApp.NewContext(false, tmproto.Header{})
	require.NoError(t, bk.SetBalances(ctx, sellerAddr, cs(c("token1", 100), c("token2", 100))))

	keeper := tApp.GetAuctionKeeper()

	// Create an auction (lot: 20 token1, initialBid: 0 token2)
	id, err := keeper.StartSurplusAuction(ctx, sellerModName, c("token1", 20), "token2") // lot, bid denom
	require.NoError(t, err)

	// Attempt to close the auction before EndTime
	require.Error(t, keeper.CloseAuction(ctx, id))

	// Attempt to close auction that does not exist
	require.Error(t, keeper.CloseAuction(ctx, 999))
}

func TestCloseExpiredAuctions(t *testing.T) {
	// Set up
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	buyer := addrs[0]
	sellerModName := "liquidator"
	sellerAddr := authtypes.NewModuleAddress(sellerModName)

	tApp := app.NewTestApp()
	bk := tApp.GetBankKeeper()

	sellerAcc := authtypes.NewEmptyModuleAccount(sellerModName, authtypes.Burner) // forward auctions burn proceeds
	// require.NoError(t, sellerAcc.SetCoins(cs(c("token1", 100), c("token2", 100))))
	tApp.InitializeFromGenesisStates(
		// NewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	authtypes.NewBaseAccount(buyer, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	sellerAcc,
		// }),
		app.NewAuthGenState(tApp, []sdk.AccAddress{buyer}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{sellerAcc}),
	)
	ctx := tApp.NewContext(false, tmproto.Header{})
	require.NoError(t, bk.SetBalances(ctx, sellerAddr, cs(c("token1", 100), c("token2", 100))))

	keeper := tApp.GetAuctionKeeper()

	// Start auction 1
	_, err := keeper.StartSurplusAuction(ctx, sellerModName, c("token1", 20), "token2") // lot, bid denom
	require.NoError(t, err)

	// Start auction 2
	_, err = keeper.StartSurplusAuction(ctx, sellerModName, c("token1", 20), "token2") // lot, bid denom
	require.NoError(t, err)

	// Fast forward the block time
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(auctiontypes.DefaultMaxAuctionDuration).Add(1))

	// Close expired auctions
	err = keeper.CloseExpiredAuctions(ctx)
	require.NoError(t, err)
}
