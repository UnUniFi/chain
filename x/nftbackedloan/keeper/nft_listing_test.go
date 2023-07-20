package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestListNft() {
	acc1 := suite.addrs[0]
	acc2 := suite.addrs[1]
	keeper := suite.keeper
	nftKeeper := suite.nftKeeper

	tests := []struct {
		testCase         string
		classId          string
		nftId            string
		nftOwner         sdk.AccAddress
		lister           sdk.AccAddress
		BidDenom         string
		mintBefore       bool
		listBefore       bool
		expectPass       bool
		statusListedHook bool
	}{
		{
			testCase:         "not existing nft",
			classId:          "class1",
			nftId:            "nft1",
			nftOwner:         acc1,
			lister:           acc1,
			BidDenom:         "uguu",
			mintBefore:       false,
			listBefore:       false,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "already listed",
			classId:          "class2",
			nftId:            "nft2",
			nftOwner:         acc1,
			lister:           acc1,
			BidDenom:         "uguu",
			mintBefore:       true,
			listBefore:       true,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "not owned nft",
			classId:          "class3",
			nftId:            "nft3",
			nftOwner:         acc1,
			lister:           acc2,
			BidDenom:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "unsupported bid token",
			classId:          "class4",
			nftId:            "nft4",
			nftOwner:         acc1,
			lister:           acc1,
			BidDenom:         "xxxx",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "successful listing with default active rank",
			classId:          "class5",
			nftId:            "nft5",
			nftOwner:         acc1,
			lister:           acc1,
			BidDenom:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       true,
			statusListedHook: false,
		},
		{
			testCase:         "successful listing with non-default active rank",
			classId:          "class6",
			nftId:            "nft6",
			nftOwner:         acc1,
			lister:           acc1,
			BidDenom:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       true,
			statusListedHook: false,
		},
		{
			testCase:         "successful anther owner",
			classId:          "class7",
			nftId:            "nft7",
			nftOwner:         acc2,
			lister:           acc2,
			BidDenom:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       true,
			statusListedHook: false,
		},
	}

	for _, tc := range tests {
		statusAfterNftListed = false
		if tc.mintBefore {
			_ = nftKeeper.SaveClass(suite.ctx, nfttypes.Class{
				Id:          tc.classId,
				Name:        tc.classId,
				Symbol:      tc.classId,
				Description: tc.classId,
				Uri:         tc.classId,
			})
			err := nftKeeper.Mint(suite.ctx, nfttypes.NFT{
				ClassId: tc.classId,
				Id:      tc.nftId,
				Uri:     tc.nftId,
				UriHash: tc.nftId,
			}, tc.nftOwner)
			suite.Require().NoError(err)
		}
		if tc.listBefore {
			err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:             tc.lister.String(),
				NftId:              types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
				BidDenom:           tc.BidDenom,
				MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			})
			suite.Require().NoError(err)
		}
		err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             tc.lister.String(),
			NftId:              types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
			BidDenom:           tc.BidDenom,
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// get listing
			listing, err := keeper.GetNftListingByIdBytes(suite.ctx, (types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}).IdBytes())
			suite.Require().NoError(err)

			// check ownership is transferred
			// moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
			nftOwner := nftKeeper.GetOwner(suite.ctx, tc.classId, tc.nftId)
			suite.Require().Equal(nftOwner.String(), tc.nftOwner.String())

			// check startedAt is set as current time
			suite.Require().Equal(suite.ctx.BlockTime(), listing.StartedAt)
		} else {
			suite.Require().Error(err)
		}

		suite.Require().Equal(tc.statusListedHook, statusAfterNftListed, tc.testCase)
	}
}

func (suite *KeeperTestSuite) TestCancelNftListing() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	keeper := suite.keeper
	nftKeeper := suite.nftKeeper

	params := suite.app.NftbackedloanKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase           string
		classId            string
		nftId              string
		nftOwner           sdk.AccAddress
		canceller          sdk.AccAddress
		cancelAfter        time.Duration
		numBids            int
		listBefore         bool
		expectPass         bool
		statusUnlistedHook bool
	}{
		{
			testCase:           "not existing listing",
			classId:            "class1",
			nftId:              "nft1",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         false,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "not owned nft listing",
			classId:            "class2",
			nftId:              "nft2",
			nftOwner:           acc1,
			canceller:          acc2,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         true,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "cancel time not pass",
			classId:            "class3",
			nftId:              "nft3",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        0,
			numBids:            0,
			listBefore:         true,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "successful cancel without cancel fee",
			classId:            "class5",
			nftId:              "nft5",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         true,
			expectPass:         true,
			statusUnlistedHook: false,
		},
		{
			testCase:           "successful cancel with cancel fee",
			classId:            "class6",
			nftId:              "nft6",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         true,
			expectPass:         true,
			statusUnlistedHook: false,
		},
	}

	for _, tc := range tests {
		statusAfterNftUnlistedWithoutPayment = false

		_ = nftKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err := nftKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		if tc.listBefore {
			err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:             tc.nftOwner.String(),
				NftId:              nftIdentifier,
				BidDenom:           "uguu",
				MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			})
			suite.Require().NoError(err)
		}

		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

			// init tokens to addr
			price := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			deposit := sdk.NewInt64Coin("uguu", int64(100000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{price})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{price})
			suite.NoError(err)

			err := suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            nftIdentifier,
				Price:            price,
				Expiry:           time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.MustNewDecFromStr("0.05"),
				AutomaticPayment: false,
				Deposit:          deposit,
			})
			suite.Require().NoError(err)
		}

		oldCancellerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.canceller, "uguu")
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(tc.cancelAfter))
		err = keeper.CancelNftListing(suite.ctx, &types.MsgCancelNftListing{
			Sender: tc.canceller.String(),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// check all bids are closed and returned
			nftBids := keeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Len(nftBids, 0)

			// check cancel fee is reduced from listing owner
			if tc.numBids > 0 {
				cancellerNewBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.canceller, "uguu")
				suite.Require().True(cancellerNewBalance.Amount.LT(oldCancellerBalance.Amount))
			}

			// check nft ownership is returned back to owner
			owner := nftKeeper.GetOwner(suite.ctx, tc.classId, tc.nftId)
			suite.Require().Equal(owner, tc.nftOwner)

			// check nft listing is deleted
			_, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Error(err)
		} else {
			suite.Require().Error(err)
		}

		suite.Require().Equal(tc.statusUnlistedHook, statusAfterNftUnlistedWithoutPayment, tc.testCase)
	}
}

func (suite *KeeperTestSuite) TestDeliverSuccessfulBids() {
	suite.SetupTest()
	keeper := suite.keeper

	classId := "class1"
	nftId := "nf1"
	nftOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()

	suite.ctx = suite.ctx.WithBlockTime(now)
	price := sdk.NewInt64Coin("uguu", int64(1000000000))
	deposit := sdk.NewInt64Coin("uguu", int64(100000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{price})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{price})
	suite.NoError(err)

	_ = suite.app.NewNFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
		Id:          classId,
		Name:        classId,
		Symbol:      classId,
		Description: classId,
		Uri:         classId,
	})
	err = suite.app.NewNFTKeeper.Mint(suite.ctx, nfttypes.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     nftId,
		UriHash: nftId,
	}, nftOwner)
	suite.Require().NoError(err)

	nftIdentifier := types.NftIdentifier{ClassId: classId, NftId: nftId}
	err = suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
		Sender:             nftOwner.String(),
		NftId:              nftIdentifier,
		BidDenom:           "uguu",
		MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
	})
	suite.Require().NoError(err)

	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// init tokens to addr
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{price})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{price})
	suite.NoError(err)

	err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
		Sender:           bidder.String(),
		NftId:            nftIdentifier,
		Price:            price,
		Expiry:           time.Now().Add(time.Hour * 24),
		InterestRate:     sdk.MustNewDecFromStr("0.05"),
		AutomaticPayment: true,
		Deposit:          deposit,
	})
	suite.Require().NoError(err)
	err = keeper.PayRemainder(suite.ctx, &types.MsgPayRemainder{
		Sender: bidder.String(),
		NftId:  nftIdentifier,
	})
	suite.Require().NoError(err)

	listing, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().NoError(err)
	listing.SuccessfulBidEndAt = now
	listing.LiquidatedAt = now
	listing.State = types.ListingState_SUCCESSFUL_BID
	keeper.SetNftListing(suite.ctx, listing)

	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second))
	oldNftOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")

	keeper.DeliverSuccessfulBids(suite.ctx)

	// check nft transfer
	newNftOwner := suite.app.NewNFTKeeper.GetOwner(suite.ctx, classId, nftId)
	suite.Require().NoError(err)
	suite.Require().Equal(newNftOwner.String(), bidder.String())

	// check fee paid
	newOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")
	suite.Require().True(newOwnerBalance.Amount.GT(oldNftOwnerBalance.Amount))

	// check bid deleted
	bids := keeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Len(bids, 0)

	// check nft listing deleted
	_, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Error(err)

	// check if AfterNftPaymentWithCommission is called
	suite.Require().True(statusAfterNftPaymentWithCommission)
}
