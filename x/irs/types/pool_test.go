package types

import (
	"fmt"
	"testing"
	"time"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var (
	defaultSpreadFactor = sdk.MustNewDecFromStr("0.025")
	defaultSwapDee      = defaultSpreadFactor
	defaultExitFee      = sdk.ZeroDec()
	defaultPoolId       = uint64(1)
	defaultContractAddr = sdk.AccAddress([]byte("contract")).String()
	defaultStartTime    = uint64(1701388800) // 2023-12-01 00:00:00 +0000 UTC
	defaultMaturity     = uint64(31536000)   // 1 year

	defaultTwoAssetScalingFactors   = []uint64{1, 1}
	defaultThreeAssetScalingFactors = []uint64{1, 1, 1}
	defaultFiveAssetScalingFactors  = []uint64{1, 1, 1, 1, 1}
	defaultFutureGovernor           = ""

	twoEvenStablePoolAssets = sdk.NewCoins(
		sdk.NewInt64Coin("bar", 1000000000),
		sdk.NewInt64Coin("foo", 1000000000),
	)
	twoUnevenStablePoolAssets = sdk.NewCoins(
		sdk.NewInt64Coin("bar", 1000000000),
		sdk.NewInt64Coin("foo", 2000000000),
	)
	threeEvenStablePoolAssets = sdk.NewCoins(
		sdk.NewInt64Coin("asset/a", 1000000),
		sdk.NewInt64Coin("asset/b", 1000000),
		sdk.NewInt64Coin("asset/c", 1000000),
	)
	threeUnevenStablePoolAssets = sdk.NewCoins(
		sdk.NewInt64Coin("asset/a", 1000000),
		sdk.NewInt64Coin("asset/b", 2000000),
		sdk.NewInt64Coin("asset/c", 3000000),
	)
	fiveEvenStablePoolAssets = sdk.NewCoins(
		sdk.NewInt64Coin("asset/a", 1000000000),
		sdk.NewInt64Coin("asset/b", 1000000000),
		sdk.NewInt64Coin("asset/c", 1000000000),
		sdk.NewInt64Coin("asset/d", 1000000000),
		sdk.NewInt64Coin("asset/e", 1000000000),
	)
	fiveUnevenStablePoolAssets = sdk.NewCoins(
		sdk.NewInt64Coin("asset/a", 1000000000),
		sdk.NewInt64Coin("asset/b", 2000000000),
		sdk.NewInt64Coin("asset/c", 3000000000),
		sdk.NewInt64Coin("asset/d", 4000000000),
		sdk.NewInt64Coin("asset/e", 5000000000),
	)
)

func TestEnsureDenomInPool(t *testing.T) {
	tests := []struct {
		name       string
		poolAssets []sdk.Coin
		TokensIn   []sdk.Coin
		err        bool
	}{
		{
			name: "exist in pool",
			poolAssets: []sdk.Coin{
				sdk.NewCoin("denom", sdk.NewInt(1000)),
			},
			TokensIn: []sdk.Coin{
				sdk.NewCoin("denom", sdk.NewInt(1000)),
			},
			err: false,
		},
		{
			name: "not exist in pool",
			poolAssets: []sdk.Coin{
				sdk.NewCoin("denom", sdk.NewInt(1000)),
			},
			TokensIn: []sdk.Coin{
				sdk.NewCoin("denom2", sdk.NewInt(1000)),
			},
			err: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ensureDenomInPool(tt.poolAssets, tt.TokensIn)
			if tt.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMaximalExactRatioJoin(t *testing.T) {
	emptyContext := sdk.Context{}

	tranchePoolAssets := sdk.NewCoins(sdk.NewInt64Coin("foo", 100), sdk.NewInt64Coin("bar", 100))
	pool := poolStructFromAssets(tranchePoolAssets)

	tests := []struct {
		name        string
		pool        TranchePool
		tokensIn    sdk.Coins
		expNumShare sdk.Int
		expRemCoin  sdk.Coins
	}{
		{
			name:        "two asset pool, same tokenIn ratio",
			tokensIn:    sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(10)), sdk.NewCoin("bar", sdk.NewInt(10))),
			expNumShare: sdk.NewIntFromUint64(10000000000000000000),
			expRemCoin:  sdk.Coins{},
		},
		{
			name:        "two asset pool, different tokenIn ratio with pool",
			tokensIn:    sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(10)), sdk.NewCoin("bar", sdk.NewInt(11))),
			expNumShare: sdk.NewIntFromUint64(10000000000000000000),
			expRemCoin:  sdk.NewCoins(sdk.NewCoin("bar", sdk.NewIntFromUint64(1))),
		},
	}

	for _, test := range tests {
		numShare, remCoins, err := MaximalExactRatioJoin(&pool, emptyContext, test.tokensIn)

		require.NoError(t, err)
		require.Equal(t, test.expNumShare, numShare)
		require.Equal(t, test.expRemCoin, remCoins)
	}
}

func TestCalcJoinPoolNoSwapShares(t *testing.T) {
	tenPercentOfTwoPool := int64(1000000000 / 10)
	tenPercentOfThreePool := int64(1000000 / 10)
	tests := map[string]struct {
		tokensIn        sdk.Coins
		poolAssets      sdk.Coins
		expNumShare     sdk.Int
		expTokensJoined sdk.Coins
		expPoolAssets   sdk.Coins
		expectPass      bool
	}{
		"even two asset pool, same tokenIn ratio": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool))),
			poolAssets:      twoEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool))),
			expPoolAssets:   twoEvenStablePoolAssets,
			expectPass:      true,
		},
		"even two asset pool, different tokenIn ratio with pool": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool+1))),
			poolAssets:      twoEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool))),
			expPoolAssets:   twoEvenStablePoolAssets,
			expectPass:      true,
		},
		"uneven two asset pool, same tokenIn ratio": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(2*tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool))),
			poolAssets:      twoUnevenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(2*tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool))),
			expPoolAssets:   twoUnevenStablePoolAssets,
			expectPass:      true,
		},
		"uneven two asset pool, different tokenIn ratio with pool": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(2*tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool+1))),
			poolAssets:      twoUnevenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(2*tenPercentOfTwoPool)), sdk.NewCoin("bar", sdk.NewInt(tenPercentOfTwoPool))),
			expPoolAssets:   twoUnevenStablePoolAssets,
			expectPass:      true,
		},
		"even three asset pool, same tokenIn ratio": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(tenPercentOfThreePool))),
			poolAssets:      threeEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(tenPercentOfThreePool))),
			expPoolAssets:   threeEvenStablePoolAssets,
			expectPass:      true,
		},
		"even three asset pool, different tokenIn ratio with pool": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(tenPercentOfThreePool+1))),
			poolAssets:      threeEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(tenPercentOfThreePool))),
			expPoolAssets:   threeEvenStablePoolAssets,
			expectPass:      true,
		},
		"uneven three asset pool, same tokenIn ratio": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(2*tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(3*tenPercentOfThreePool))),
			poolAssets:      threeUnevenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(2*tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(3*tenPercentOfThreePool))),
			expPoolAssets:   threeUnevenStablePoolAssets,
			expectPass:      true,
		},
		"uneven three asset pool, different tokenIn ratio with pool": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(2*tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(3*tenPercentOfThreePool+1))),
			poolAssets:      threeUnevenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(2*tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(3*tenPercentOfThreePool))),
			expPoolAssets:   threeUnevenStablePoolAssets,
			expectPass:      true,
		},
		"uneven three asset pool, uneven scaling factors": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(2*tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(3*tenPercentOfThreePool))),
			poolAssets:      threeUnevenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(10000000000000000000),
			expTokensJoined: sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(2*tenPercentOfThreePool)), sdk.NewCoin("asset/c", sdk.NewInt(3*tenPercentOfThreePool))),
			expPoolAssets:   threeUnevenStablePoolAssets,
			expectPass:      true,
		},

		// error catching
		"even two asset pool, no-swap join attempt with one asset": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tenPercentOfTwoPool))),
			poolAssets:      twoEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(0),
			expTokensJoined: sdk.Coins{},
			expPoolAssets:   twoEvenStablePoolAssets,
			expectPass:      false,
		},
		"even two asset pool, no-swap join attempt with one valid and one invalid asset": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tenPercentOfTwoPool)), sdk.NewCoin("baz", sdk.NewInt(tenPercentOfTwoPool))),
			poolAssets:      twoEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(0),
			expTokensJoined: sdk.Coins{},
			expPoolAssets:   twoEvenStablePoolAssets,
			expectPass:      false,
		},
		"even two asset pool, no-swap join attempt with two invalid assets": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("baz", sdk.NewInt(tenPercentOfTwoPool)), sdk.NewCoin("qux", sdk.NewInt(tenPercentOfTwoPool))),
			poolAssets:      twoEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(0),
			expTokensJoined: sdk.Coins{},
			expPoolAssets:   twoEvenStablePoolAssets,
			expectPass:      false,
		},
		"even three asset pool, no-swap join attempt with an invalid asset": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("asset/a", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("asset/b", sdk.NewInt(tenPercentOfThreePool)), sdk.NewCoin("qux", sdk.NewInt(tenPercentOfThreePool))),
			poolAssets:      threeEvenStablePoolAssets,
			expNumShare:     sdk.NewIntFromUint64(0),
			expTokensJoined: sdk.Coins{},
			expPoolAssets:   threeEvenStablePoolAssets,
			expectPass:      false,
		},
		"single asset pool, no-swap join attempt with one asset": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(sdk.MaxSortableDec.TruncateInt64()))),
			poolAssets:      sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(1))),
			expNumShare:     sdk.NewIntFromUint64(0),
			expTokensJoined: sdk.Coins{},
			expPoolAssets:   sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(1))),
			expectPass:      false,
		},
		"attempt joining pool with no assets in it": {
			tokensIn:        sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(1))),
			poolAssets:      sdk.Coins{},
			expNumShare:     sdk.NewIntFromUint64(0),
			expTokensJoined: sdk.Coins{},
			expPoolAssets:   sdk.Coins{},
			expectPass:      false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := sdk.Context{}
			pool := poolStructFromAssets(test.poolAssets)
			numShare, tokensJoined, err := pool.CalcJoinPoolNoSwapShares(ctx, test.tokensIn, pool.SwapFee)

			if test.expectPass {
				require.NoError(t, err)
				require.Equal(t, test.expPoolAssets, sdk.Coins(pool.PoolAssets))
				require.Equal(t, test.expNumShare, numShare)
				require.Equal(t, test.expTokensJoined, tokensJoined)
			} else {
				require.Error(t, err)
				require.Equal(t, test.expPoolAssets, sdk.Coins(pool.PoolAssets))
				require.Equal(t, test.expNumShare, numShare)
				require.Equal(t, test.expTokensJoined, tokensJoined)
			}
		})
	}
}

func TestCalcExitPool(t *testing.T) {
	emptyContext := sdk.Context{}

	twoStablePoolAssets := sdk.NewCoins(
		sdk.NewInt64Coin("foo", 1000000000),
		sdk.NewInt64Coin("bar", 1000000000),
	)
	threePoolAssets := sdk.NewCoins(sdk.NewInt64Coin("foo", 2000000000), sdk.NewInt64Coin("bar", 3000000000), sdk.NewInt64Coin("baz", 4000000000))

	// create these pools used for testing
	twoAssetPool := poolStructFromAssets(twoStablePoolAssets)
	threeAssetPool := poolStructFromAssets(threePoolAssets)

	twoAssetPoolWithExitFee := poolStructFromAssets(twoStablePoolAssets)
	twoAssetPoolWithExitFee.ExitFee = sdk.MustNewDecFromStr("0.0001")

	threeAssetPoolWithExitFee := poolStructFromAssets(threePoolAssets)
	threeAssetPoolWithExitFee.ExitFee = sdk.MustNewDecFromStr("0.0002")

	tests := []struct {
		name          string
		pool          TranchePool
		exitingShares sdk.Int
		expError      bool
	}{
		{
			name:          "two-asset pool, exiting shares grater than total shares",
			pool:          twoAssetPool,
			exitingShares: twoAssetPool.GetTotalShares().Amount.AddRaw(1),
			expError:      true,
		},
		{
			name:          "three-asset pool, exiting shares grater than total shares",
			pool:          threeAssetPool,
			exitingShares: threeAssetPool.GetTotalShares().Amount.AddRaw(1),
			expError:      true,
		},
		{
			name:          "two-asset pool, valid exiting shares",
			pool:          twoAssetPool,
			exitingShares: twoAssetPool.GetTotalShares().Amount.QuoRaw(2),
			expError:      false,
		},
		{
			name:          "three-asset pool, valid exiting shares",
			pool:          threeAssetPool,
			exitingShares: sdk.NewIntFromUint64(3000000000000),
			expError:      false,
		},
		{
			name:          "two-asset pool with exit fee, valid exiting shares",
			pool:          twoAssetPoolWithExitFee,
			exitingShares: twoAssetPoolWithExitFee.GetTotalShares().Amount.QuoRaw(2),
			expError:      false,
		},
		{
			name:          "three-asset pool with exit fee, valid exiting shares",
			pool:          threeAssetPoolWithExitFee,
			exitingShares: sdk.NewIntFromUint64(7000000000000),
			expError:      false,
		},
	}

	for _, test := range tests {
		// using empty context since, currently, the context is not used anyway. This might be changed in the future
		exitFee := test.pool.ExitFee
		exitCoins, err := CalcExitPool(emptyContext, &test.pool, test.exitingShares, exitFee)
		if test.expError {
			require.Error(t, err, "test: %v", test.name)
		} else {
			require.NoError(t, err, "test: %v", test.name)

			// exitCoins = ( (1 - exitFee) * exitingShares / poolTotalShares ) * poolTotalLiquidity
			expExitCoins := mulCoins(test.pool.GetPoolAssets(), (sdk.OneDec().Sub(exitFee)).MulInt(test.exitingShares).QuoInt(test.pool.GetTotalShares().Amount))
			require.Equal(t, expExitCoins.Sort().String(), exitCoins.Sort().String(), "test: %v", test.name)
		}
	}
}

// we create a pool struct directly to bypass checks in NewStableswapPool()
func poolStructFromAssets(assets sdk.Coins) TranchePool {

	p := TranchePool{
		Id:               defaultPoolId,
		StrategyContract: defaultContractAddr,
		StartTime:        defaultStartTime,
		Maturity:         defaultMaturity,
		SwapFee:          defaultSwapDee,
		ExitFee:          defaultExitFee,
		TotalShares:      sdk.Coin{},
		PoolAssets:       assets,
	}
	InitPoolSharesSupply := OneShare.MulRaw(100)
	p.TotalShares = sdk.NewCoin(LsDenom(p), InitPoolSharesSupply)
	return p
}

func mulCoins(coins sdk.Coins, multiplier sdk.Dec) sdk.Coins {
	outCoins := sdk.Coins{}
	for _, coin := range coins {
		outCoin := sdk.NewCoin(coin.Denom, multiplier.MulInt(coin.Amount).TruncateInt())
		if !outCoin.Amount.IsZero() {
			outCoins = append(outCoins, outCoin)
		}
	}
	return outCoins
}

// SwapOutAmtGivenIn is a mutative method for CalcOutAmtGivenIn, which includes the actual swap.
func TestSwapOutAmtGivenIn(t *testing.T) {
	tests := map[string]struct {
		poolAssets            sdk.Coins
		tokenIn               sdk.Coin
		expectedTokenOut      sdk.Coin
		expectedPoolLiquidity sdk.Coins
		spreadFactor          sdk.Dec
		expError              bool
	}{
		"even pool basic trade": {
			poolAssets:            twoEvenStablePoolAssets,
			tokenIn:               sdk.NewInt64Coin("foo", 100),
			expectedTokenOut:      sdk.NewInt64Coin("bar", 99),
			expectedPoolLiquidity: twoEvenStablePoolAssets.Add(sdk.NewInt64Coin("foo", 100)).Sub(sdk.NewCoins(sdk.NewInt64Coin("bar", 99))...),
			spreadFactor:          sdk.ZeroDec(),
			expError:              false,
		},
		"100:1 scaling factor ratio, even swap": {
			poolAssets: sdk.NewCoins(
				sdk.NewInt64Coin("bar", 1000000000),
				sdk.NewInt64Coin("foo", 10000000),
			),
			tokenIn:          sdk.NewInt64Coin("foo", 100),
			expectedTokenOut: sdk.NewInt64Coin("bar", 6762), // todo calculate this
			expectedPoolLiquidity: sdk.NewCoins(
				sdk.NewInt64Coin("bar", 1000000000).SubAmount(sdk.NewIntFromUint64(9999)),
				sdk.NewInt64Coin("foo", 10000000).AddAmount(sdk.NewIntFromUint64(100)),
			),
			spreadFactor: sdk.ZeroDec(),
			expError:     false,
		},
		// TODO: Add test cases here, where they're off 1-1 ratio
		// * (we just need to verify that the further off they are, further slippage is)
		// * Add test cases with non-zero spread factor.
		// looks like its really an error due to slippage at limit
		"trade hits max pool capacity for asset": {
			poolAssets: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 9_999_999_998),
				sdk.NewInt64Coin("bar", 9_999_999_998),
			),
			tokenIn:          sdk.NewInt64Coin("foo", 1),
			expectedTokenOut: sdk.Coin{},
			expectedPoolLiquidity: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 9_999_999_999),
				sdk.NewInt64Coin("bar", 9_999_999_997),
			),
			spreadFactor: sdk.ZeroDec(),
			expError:     true,
		},
		"trade exceeds max pool capacity for asset": {
			poolAssets: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 10_000_000_000),
				sdk.NewInt64Coin("bar", 10_000_000_000),
			),
			tokenIn:          sdk.NewInt64Coin("foo", 1),
			expectedTokenOut: sdk.Coin{},
			expectedPoolLiquidity: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 10_000_000_000),
				sdk.NewInt64Coin("bar", 10_000_000_000),
			),
			spreadFactor: sdk.ZeroDec(),
			expError:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db := tmdb.NewMemDB()
			stateStore := store.NewCommitMultiStore(db)
			ctx := sdk.NewContext(stateStore, tmproto.Header{Time: time.Unix(1704067200, 0)}, false, log.NewNopLogger())
			p := poolStructFromAssets(tc.poolAssets)

			tokenOut, err := p.SwapOutAmtGivenIn(ctx, tc.tokenIn, tc.expectedTokenOut.Denom, tc.spreadFactor)
			if tc.expError {
				require.Error(t, err)
				return
			}
			fmt.Println("expErr", tc.expError, "err", err)
			require.NoError(t, err)
			poolLiquidity := sdk.Coins(p.PoolAssets)
			if !tc.expError {
				require.Equal(t, tc.expectedTokenOut.Amount, tokenOut.Amount)
				require.True(t, poolLiquidity.IsAllGTE(tc.expectedPoolLiquidity),
					"p.PoolLiquidity.IsAllGTE(tc.expectedPoolLiquidity) failed. Pool liq %v, expected %v",
					poolLiquidity, tc.expectedPoolLiquidity)
			}
		})
	}
}

func TestApplySwap(t *testing.T) {
	tests := []struct {
		name              string
		currentPoolAssets sdk.Coins
		tokenIn           sdk.Coin
		tokenOut          sdk.Coin
		swapFeeIn         sdk.Dec
		swapFeeOut        sdk.Dec
		newPoolAssets     sdk.Coins
		expectPass        bool
	}{
		{
			name:              "zero swap fee in",
			currentPoolAssets: twoEvenStablePoolAssets,
			tokenIn:           sdk.NewInt64Coin("bar", 100),
			tokenOut:          sdk.NewInt64Coin("foo", 100),
			swapFeeIn:         sdk.ZeroDec(),
			swapFeeOut:        sdk.ZeroDec(),
			newPoolAssets: sdk.NewCoins(
				sdk.NewInt64Coin("bar", 1000000100),
				sdk.NewInt64Coin("foo", 999999900),
			),
			expectPass: true,
		},
		{
			name:              "positive swap fee in",
			currentPoolAssets: twoEvenStablePoolAssets,
			tokenIn:           sdk.NewInt64Coin("bar", 1000),
			tokenOut:          sdk.NewInt64Coin("foo", 1000),
			swapFeeIn:         sdk.MustNewDecFromStr("0.025"),
			swapFeeOut:        sdk.ZeroDec(),
			newPoolAssets: sdk.NewCoins(
				sdk.NewInt64Coin("bar", 1000000975),
				sdk.NewInt64Coin("foo", 999999000),
			),
			expectPass: true,
		},
		{
			name:              "positive swap fee out",
			currentPoolAssets: twoEvenStablePoolAssets,
			tokenIn:           sdk.NewInt64Coin("bar", 1000),
			tokenOut:          sdk.NewInt64Coin("foo", 1000),
			swapFeeIn:         sdk.MustNewDecFromStr("0.025"),
			swapFeeOut:        sdk.MustNewDecFromStr("0.05"),
			newPoolAssets: sdk.NewCoins(
				sdk.NewInt64Coin("bar", 1000000975),
				sdk.NewInt64Coin("foo", 999999050),
			),
			expectPass: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pool := poolStructFromAssets(test.currentPoolAssets)
			err := pool.applySwap(test.tokenIn, test.tokenOut, test.swapFeeIn, test.swapFeeOut)
			if test.expectPass {
				require.NoError(t, err)
				require.Equal(t, test.newPoolAssets, sdk.Coins(pool.PoolAssets))
			} else {
				require.Error(t, err)
			}
		})
	}
}
