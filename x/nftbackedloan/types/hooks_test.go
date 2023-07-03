package types_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"

	simapp "github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/nftbackedloan/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx   sdk.Context
	app   *simapp.App
	addrs []sdk.AccAddress
}

func (suite *KeeperTestSuite) SetupTest(hooks types.NftbackedloanHooks) {
	isCheckTx := false

	app := simapp.Setup(suite.T(), ([]wasm.Option{})...)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.addrs = simapp.AddTestAddrsIncremental(app, suite.ctx, 1, sdk.NewInt(30000000))
	suite.app = app

	if hooks != nil {
		suite.app.NftbackedloanKeeper.SetHooks(hooks)
	}
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func dummyAfterNftListedEvent(nftId types.NftIdentifier) sdk.Event {
	return sdk.NewEvent(
		"afterNftListed",
		sdk.NewAttribute("nftId", nftId.String()),
	)
}

func dummyAfterNftPaymentWithCommissionEvent(nftId types.NftIdentifier) sdk.Event {
	return sdk.NewEvent(
		"afterNftPaymentWithCommission",
		sdk.NewAttribute("nftId", nftId.String()),
	)
}

func dummyAfterNftUnlistedWithoutPaymentEvent(nftId types.NftIdentifier) sdk.Event {
	return sdk.NewEvent(
		"afterNftUnlistedWithoutPayment",
		sdk.NewAttribute("nftId", nftId.String()),
	)
}

// dummyNftmarketHook is a struct satisfying the nftmarket hook interface,
// that maintains a counter for how many times its been succesfully called,
// and a boolean for whether it should panic during its execution.
type dummyNftmarketHook struct {
	successCounter int
	shouldPanic    bool
}

func (hook *dummyNftmarketHook) AfterNftListed(ctx sdk.Context, nftId types.NftIdentifier, txMemo string) {
	if hook.shouldPanic {
		panic("dummyNftmarketHook AfterNftListed is panicking")
	}

	hook.successCounter += 1
	ctx.EventManager().EmitEvent(dummyAfterNftListedEvent(nftId))
}

func (hook *dummyNftmarketHook) AfterNftPaymentWithCommission(ctx sdk.Context, nftId types.NftIdentifier, fee sdk.Coin) {
	if hook.shouldPanic {
		panic("dummyNftmarketHook AfterNftPaymentWithCommission is panicking")
	}

	hook.successCounter += 1
	ctx.EventManager().EmitEvent(dummyAfterNftPaymentWithCommissionEvent(nftId))
}

func (hook *dummyNftmarketHook) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftId types.NftIdentifier) {
	if hook.shouldPanic {
		panic("dummyNftmarketHook AfterNftUnlistedWithoutPayment is panicking")
	}

	hook.successCounter += 1
	ctx.EventManager().EmitEvent(dummyAfterNftUnlistedWithoutPaymentEvent(nftId))
}

func (hook *dummyNftmarketHook) Clone() *dummyNftmarketHook {
	newHook := dummyNftmarketHook{shouldPanic: hook.shouldPanic, successCounter: hook.successCounter}
	return &newHook
}

var _ types.NftbackedloanHooks = &dummyNftmarketHook{}

func (suite *KeeperTestSuite) TestHooksPanicRecovery() {
	panicHook := dummyNftmarketHook{shouldPanic: true}
	noPanicHook := dummyNftmarketHook{shouldPanic: false}
	nftId := types.NftIdentifier{
		ClassId: "dummyhook",
		NftId:   "dummyhook",
	}

	tests := []struct {
		hooks                 []dummyNftmarketHook
		expectedCounterValues []int
		lenEvents             int
	}{
		{[]dummyNftmarketHook{noPanicHook}, []int{3}, 3},
		{[]dummyNftmarketHook{panicHook}, []int{0}, 0},
	}

	for tcIndex, tc := range tests {
		suite.SetupTest(nil)
		hookRefs := []types.NftbackedloanHooks{}

		for _, hook := range tc.hooks {
			hookRefs = append(hookRefs, hook.Clone())
		}

		// insert dummy hook struct as part of NftbackedloanHooks
		hooks := types.NewMultiNftbackedloanHooks(hookRefs...)
		// suite.app.NftbackedloanKeeper.SetHooks(hooks)

		if tc.lenEvents == 0 {
			suite.Panics(func() {
				hooks.AfterNftListed(suite.ctx, nftId, "test")

				hooks.AfterNftPaymentWithCommission(suite.ctx, nftId, sdk.Coin{Denom: "uguu", Amount: sdk.OneInt()})

				hooks.AfterNftUnlistedWithoutPayment(suite.ctx, nftId)
			})
		} else {
			suite.NotPanics(func() {
				hooks.AfterNftListed(suite.ctx, nftId, "test")

				hooks.AfterNftPaymentWithCommission(suite.ctx, nftId, sdk.Coin{Denom: "uguu", Amount: sdk.OneInt()})

				hooks.AfterNftUnlistedWithoutPayment(suite.ctx, nftId)
			})
		}

		for i := 0; i < len(hooks); i++ {
			nftmarketHook := hookRefs[i].(*dummyNftmarketHook)
			suite.Require().Equal(tc.expectedCounterValues[i], nftmarketHook.successCounter, "test case index %d", tcIndex)
		}
	}
}
