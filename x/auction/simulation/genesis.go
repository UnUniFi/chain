package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	banksim "github.com/cosmos/cosmos-sdk/x/bank/simulation"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/UnUniFi/chain/x/auction/types"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
)

const (
	// Block time params are un-exported constants in cosmos-sdk/x/simulation.
	// Copy them here in lieu of importing them.
	minTimePerBlock time.Duration = (10000 / 2) * time.Second
	maxTimePerBlock time.Duration = 10000 * time.Second

	// Calculate the average block time
	AverageBlockTime time.Duration = (maxTimePerBlock - minTimePerBlock) / 2
	// MaxBidDuration is a crude way of ensuring that BidDuration ≤ MaxAuctionDuration for all generated params
	MaxBidDuration time.Duration = AverageBlockTime * 50
)

func GenBidDuration(r *rand.Rand) time.Duration {
	d, err := RandomPositiveDuration(r, 0, MaxBidDuration)
	if err != nil {
		panic(err)
	}
	return d
}
func GenMaxAuctionDuration(r *rand.Rand) time.Duration {
	d, err := RandomPositiveDuration(r, MaxBidDuration, AverageBlockTime*200)
	if err != nil {
		panic(err)
	}
	return d
}

func GenIncrementCollateral(r *rand.Rand) sdk.Dec {
	return simulation.RandomDecAmount(r, sdk.MustNewDecFromStr("1"))
}

var GenIncrementDebt = GenIncrementCollateral
var GenIncrementSurplus = GenIncrementCollateral

// RandomizedGenState generates a random GenesisState for auction
func RandomizedGenState(simState *module.SimulationState) {

	p := types.NewParams(
		GenMaxAuctionDuration(simState.Rand),
		GenBidDuration(simState.Rand),
		GenIncrementSurplus(simState.Rand),
		GenIncrementDebt(simState.Rand),
		GenIncrementCollateral(simState.Rand),
	)
	if err := p.Validate(); err != nil {
		panic(err)
	}
	auctionGenesis := types.NewGenesisState(
		types.DefaultNextAuctionID,
		p,
		nil,
	)

	// Add auctions
	debtAuction := types.NewDebtAuction(
		cdptypes.LiquidatorMacc, // using cdp account rather than generic test one to avoid having to set permissions on the supply keeper
		sdk.NewInt64Coin("usdx", 100),
		sdk.NewInt64Coin("uguu", 1000000000000),
		simState.GenTimestamp.Add(time.Hour*5),
		sdk.NewInt64Coin("debt", 100), // same as usdx
	)
	auctions := types.GenesisAuctions{&debtAuction}

	var startingID = auctionGenesis.NextAuctionId
	var ok bool
	var totalAuctionCoins sdk.Coins
	for i, a := range auctions {
		auctions[i], ok = a.WithID(uint64(i) + startingID).(types.GenesisAuction)
		if !ok {
			panic("can't convert Auction to GenesisAuction")
		}
		totalAuctionCoins = totalAuctionCoins.Add(a.GetModuleAccountCoins()...)
	}
	auctionGenesis.NextAuctionId = startingID + uint64(len(auctions))
	packAuction, _ := types.PackGenesisAuctions(auctions)
	auctionGenesis.Auctions = append(auctionGenesis.Auctions, packAuction...)

	bidderCoins := sdk.NewCoins(sdk.NewInt64Coin("usdx", 10000000000))

	// Update the bank genesis state to reflect the new coins
	// TODO find some way for this to happen automatically / move it elsewhere

	var sendEnabledParams banktypes.SendEnabledParams
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(banktypes.KeySendEnabled), &sendEnabledParams, simState.Rand,
		func(r *rand.Rand) { sendEnabledParams = banksim.RandomGenesisSendParams(r) },
	)

	var defaultSendEnabledParam bool
	simState.AppParams.GetOrGenerate(
		simState.Cdc, string(banktypes.KeyDefaultSendEnabled), &defaultSendEnabledParam, simState.Rand,
		func(r *rand.Rand) { defaultSendEnabledParam = banksim.RandomGenesisDefaultSendParam(r) },
	)

	supply := sdk.NewCoins(totalAuctionCoins...)
	supply.Add(bidderCoins...)

	bankGenesis := banktypes.GenesisState{
		Params: banktypes.Params{
			SendEnabled:        sendEnabledParams,
			DefaultSendEnabled: defaultSendEnabledParam,
		},
		Balances: banksim.RandomGenesisBalances(simState),
		Supply:   supply,
	}

	simState.GenState[banktypes.ModuleName] = simState.Cdc.MustMarshalJSON(&bankGenesis)
	// TODO liquidator mod account doesn't need to be initialized for this example
	// - it just mints uguu, doesn't need a starting balance
	// - and supply.GetModuleAccount creates one if it doesn't exist

	// Note: this line prints out the auction genesis state, not just the auction parameters. Some sdk modules print out just the parameters.
	bz, err := json.MarshalIndent(&auctionGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&auctionGenesis)
}

func RandomPositiveDuration(r *rand.Rand, inclusiveMin, exclusiveMax time.Duration) (time.Duration, error) {
	min := int64(inclusiveMin)
	max := int64(exclusiveMax)
	if min < 0 || max < 0 {
		return 0, fmt.Errorf("min and max must be positive")
	}
	if min >= max {
		return 0, fmt.Errorf("max must be < min")
	}
	randPositiveInt64 := r.Int63n(max-min) + min
	return time.Duration(randPositiveInt64), nil
}
