package types_test

import (
	"testing"
	time "time"

	simapp "github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/nftmarket/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx   sdk.Context
	app   *simapp.App
	addrs []sdk.AccAddress
}

func (suite *KeeperTestSuite) SetupTest(hooks types.NftmarketHooks) {
	isCheckTx := false

	app := simapp.Setup(suite.T(), isCheckTx)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.addrs = simapp.AddTestAddrsIncremental(app, suite.ctx, 1, sdk.NewInt(30000000))
	suite.app = app

	if hooks != nil {
		suite.app.NftmarketKeeper.SetHooks(hooks)
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

var _ types.NftmarketHooks = &dummyNftmarketHook{}

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
		hookRefs := []types.NftmarketHooks{}

		for _, hook := range tc.hooks {
			hookRefs = append(hookRefs, hook.Clone())
		}

		// insert dummy hook struct as part of NftmarketHooks
		hooks := types.NewMultiNftmarketHooks(hookRefs...)
		// suite.app.NftmarketKeeper.SetHooks(hooks)

		if tc.lenEvents == 0 {
			suite.Panics(func() {
				hooks.AfterNftListed(suite.ctx, nftId, "test")

				hooks.AfterNftPaymentWithCommission(suite.ctx, nftId, sdk.Coin{Denom: "uguu", Amount: sdk.OneInt()})

				hooks.AfterNftUnlistedWithoutPayment(suite.ctx, nftId)
			})
		} else {
			suite.NotPanics(func() {
				// hooks.AfterNftListed(suite.ctx, nftId, "test")
				suite.MockListNft(suite.addrs[0], nftId)

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

func (suite *KeeperTestSuite) TestAfterNftListedInListNft() {
	hooks := dummyNftmarketHook{}
	suite.SetupTest(&hooks)
	//suite.app.NftmarketKeeper.SetHooks(&hooks)
	nftId := types.NftIdentifier{
		ClassId: "dummyhook",
		NftId:   "dummyhook",
	}
	suite.MockListNft(suite.addrs[0], nftId)

	suite.Require().Equal(1, hooks.successCounter)
}

func (suite *KeeperTestSuite) MockListNft(sender sdk.AccAddress, nftId types.NftIdentifier) error {
	suite.app.NFTKeeper.SaveClass(suite.ctx, nft.Class{
		Id: nftId.ClassId,
	})
	suite.app.NFTKeeper.Mint(suite.ctx, nft.NFT{
		ClassId: nftId.ClassId,
		Id:      nftId.NftId,
	}, suite.addrs[0])
	msg := types.MsgListNft{
		Sender:        sender.Bytes(),
		NftId:         nftId,
		ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:      "uguu",
		MinBid:        sdk.OneInt(),
		BidActiveRank: 3,
	}

	// check listing already exists
	_, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftId.IdBytes())
	if err == nil {
		return types.ErrNftListingAlreadyExists
	}

	// Check nft exists
	_, found := suite.app.NFTKeeper.GetNFT(suite.ctx, nftId.ClassId, nftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of nft
	owner := suite.app.NFTKeeper.GetOwner(suite.ctx, nftId.ClassId, nftId.NftId)
	if owner.String() != sender.String() {
		return types.ErrNotNftOwner
	}
	// create listing
	bidActiveRank := msg.BidActiveRank

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
	listing := types.NftListing{
		NftId:         msg.NftId,
		Owner:         owner.String(),
		ListingType:   msg.ListingType,
		State:         types.ListingState_LISTING,
		BidToken:      msg.BidToken,
		MinBid:        msg.MinBid,
		BidActiveRank: bidActiveRank,
		StartedAt:     suite.ctx.BlockTime(),
		EndAt:         suite.ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingPeriodInitial)),
	}
	suite.app.NftmarketKeeper.SaveNftListing(suite.ctx, listing)

	// this method is mock for list nft to serve in test to avoid tx memo retrieving.
	txMemo := "test"
	suite.app.NftmarketKeeper.AfterNftListed(suite.ctx, msg.NftId, txMemo)
	return nil
}
