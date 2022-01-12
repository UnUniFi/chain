package auction_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/auction"
	auctiontypes "github.com/UnUniFi/chain/x/auction/types"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
	"github.com/cosmos/cosmos-sdk/simapp"
)

func TestKeeper_BeginBlocker(t *testing.T) {
	// Setup
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	buyer := addrs[0]
	returnAddrs := addrs[1:]
	returnWeights := []sdk.Int{sdk.NewInt(1)}
	sellerModName := cdptypes.LiquidatorMacc
	modBaseAcc := authtypes.NewBaseAccount(authtypes.NewModuleAddress(sellerModName), nil, 0, 0)
	modAcc := authtypes.NewModuleAccount(modBaseAcc, sellerModName, []string{authtypes.Minter, authtypes.Burner}...)

	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{})

	require.NoError(t, simapp.FundModuleAccount(tApp.BankKeeper ,ctx, modAcc.Name, cs(c("token1", 100), c("token2", 100), c("debt", 100))))
	tApp.InitializeFromGenesisStates(
		// NewAuthGenStateFromAccs(authtypes.GenesisAccounts{
		// 	auth.NewBaseAccount(buyer, cs(c("token1", 100), c("token2", 100)), nil, 0, 0),
		// 	sellerAcc,
		// }),
		app.NewAuthGenState(tApp, []sdk.AccAddress{buyer}, []sdk.Coins{cs(c("token1", 100), c("token2", 100))}),
		app.NewAuthGenStateModAcc(tApp, []*authtypes.ModuleAccount{modAcc}),
	)

	keeper := tApp.GetAuctionKeeper()

	// Start an auction and place a bid
	auctionID, err := keeper.StartCollateralAuction(ctx, sellerModName, c("token1", 20), c("token2", 50), returnAddrs, returnWeights, c("debt", 40))
	require.NoError(t, err)
	require.NoError(t, keeper.PlaceBid(ctx, auctionID, buyer, c("token2", 30)))

	// Run the beginblocker, simulating a block time 1ns before auction expiry
	preExpiryTime := ctx.BlockTime().Add(auctiontypes.DefaultBidDuration - 1)
	auction.BeginBlocker(ctx.WithBlockTime(preExpiryTime), keeper)

	// Check auction has not been closed yet
	_, found := keeper.GetAuction(ctx, auctionID)
	require.True(t, found)

	// Run the endblocker, simulating a block time equal to auction expiry
	expiryTime := ctx.BlockTime().Add(auctiontypes.DefaultBidDuration)
	auction.BeginBlocker(ctx.WithBlockTime(expiryTime), keeper)

	// Check auction has been closed
	_, found = keeper.GetAuction(ctx, auctionID)
	require.False(t, found)
}

func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

func NewAuthGenStateFromAccs(accounts authtypes.GenesisAccounts) app.GenesisState {
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), accounts)
	return app.GenesisState{authtypes.ModuleName: authtypes.ModuleCdc.MustMarshalJSON(authGenesis)}
}
