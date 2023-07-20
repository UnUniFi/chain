package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestSellingDecision() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftbackedloanKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase      string
		classId       string
		nftId         string
		nftOwner      sdk.AccAddress
		executor      sdk.AccAddress
		numBids       int
		enoughAutoPay bool
		autoPayment   bool
		listBefore    bool
		expectPass    bool
	}{
		{
			testCase:      "not existing listing",
			classId:       "class1",
			nftId:         "nft1",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    false,
			expectPass:    false,
		},
		{
			testCase:      "not owned nft listing",
			classId:       "class2",
			nftId:         "nft2",
			nftOwner:      acc1,
			executor:      acc2,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			expectPass:    false,
		},
		{
			testCase:      "successful nft selling decision with automatic payment",
			classId:       "class4",
			nftId:         "nft4",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   true,
			listBefore:    true,
			expectPass:    true,
		},
		{
			testCase:      "successful nft selling decision with automatic payment enabled with not enough balance",
			classId:       "class5",
			nftId:         "nft5",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: false,
			autoPayment:   true,
			listBefore:    true,
			expectPass:    true,
		},
		{
			testCase:      "successful nft selling decision without automatic payment",
			classId:       "class6",
			nftId:         "nft6",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			expectPass:    true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{coin})
		suite.NoError(err)

		_ = suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
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

		lastBidder := sdk.AccAddress{}
		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
			lastBidder = bidder

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			halfCoin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)/2))
			mintCoin := coin
			if !tc.enoughAutoPay {
				mintCoin = sdk.NewInt64Coin("uguu", int64(1000000*(i+1)/2))
			}
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{mintCoin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{mintCoin})
			suite.NoError(err)

			err := suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            nftIdentifier,
				Price:            coin,
				Deposit:          halfCoin,
				AutomaticPayment: tc.autoPayment,
				InterestRate:     sdk.MustNewDecFromStr("0.1"),
				Expiry:           time.Now().Add(time.Hour * 24),
			})
			suite.Require().NoError(err)
		}

		err = suite.app.NftbackedloanKeeper.SellingDecision(suite.ctx, &types.MsgSellingDecision{
			Sender: tc.executor.String(),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)
			if tc.autoPayment {
				bid, err := suite.app.NftbackedloanKeeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), lastBidder)
				suite.Require().NoError(err)
				if tc.enoughAutoPay {
					// check automatic payment execution when user has enough balance
					suite.Require().Equal(bid.PaidAmount.Amount.Add(bid.Deposit.Amount), bid.Price.Amount, tc.testCase)
				} else {
					// check automatic payment when the user does not have enough balance
					suite.Require().NotEqual(bid.PaidAmount, bid.Price.Amount)
				}
			}

			// check full payment end time update
			listing, err := suite.app.NftbackedloanKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().NoError(err)
			suite.Require().Equal(listing.State, types.ListingState_SELLING_DECISION)
			suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingFullPaymentPeriod)), listing.FullPaymentEndAt)
		} else {
			suite.Require().Error(err)
		}
	}
}
