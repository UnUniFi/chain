package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/decentralized-vault/types"

	ununifitypes "github.com/UnUniFi/chain/types"
)

func (suite *KeeperTestSuite) TestNftLocked() {
	nftid := "a10"
	testUri := "testUri"
	testUriHash := "testUriHash"
	mp := types.Params{
		Networks: []types.Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
				Active:    true,
			},
		},
	}
	suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, mp)

	testCases := []struct {
		msg      string
		expError string
		postMsg  types.MsgNftLocked
		owner    sdk.AccAddress
	}{
		{
			"success",
			"",
			types.MsgNftLocked{
				Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
				ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:     nftid,
				Uri:       testUri,
				UriHash:   testUriHash,
			},
			suite.addrs[1],
		},
		{
			"not trusted sender",
			"sender does not mach oracle address",
			types.MsgNftLocked{
				Sender:    ununifitypes.StringAccAddress(suite.addrs[2]),
				ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:     nftid,
				Uri:       testUri,
				UriHash:   testUriHash,
			},
			suite.addrs[0],
		},
	}

	for index, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := suite.Require()
			err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &tc.postMsg)
			if tc.expError == "" {
				require.NoError(err)
				val := suite.app.NFTKeeper.GetOwner(suite.ctx, types.WrappedClassId, nftid)
				require.Equal(val, tc.owner, "the error occurred on:%d", index)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestNftUnlocked() {
	nftid := "a10"
	testUri := "testUri"
	testUriHash := "testUriHash"
	mp := types.Params{
		Networks: []types.Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
				Active:    true,
			},
		},
	}
	suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, mp)

	testCases := []struct {
		msg      string
		expError string
		postMsg  types.MsgNftUnlocked
		owner    sdk.AccAddress
	}{
		{
			"not trusted sender",
			"sender does not mach oracle address",
			types.MsgNftUnlocked{
				Sender:    ununifitypes.StringAccAddress(suite.addrs[2]),
				ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:     nftid,
			},
			suite.addrs[1],
		},
		{
			"success",
			"",
			types.MsgNftUnlocked{
				Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
				ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:     nftid,
			},
			nil,
		},
	}

	suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &types.MsgNftLocked{
		Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
		ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
		NftId:     nftid,
		Uri:       testUri,
		UriHash:   testUriHash,
	})
	for index, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := suite.Require()
			err := suite.app.DecentralizedvaultKeeper.NftUnlocked(suite.ctx, &tc.postMsg)
			if tc.expError == "" {
				require.NoError(err)
				val := suite.app.NFTKeeper.GetOwner(suite.ctx, types.WrappedClassId, nftid)
				require.Equal(val, tc.owner, "the error occurred on:%d", index)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
				val := suite.app.NFTKeeper.GetOwner(suite.ctx, types.WrappedClassId, nftid)
				require.Equal(val, tc.owner, "the error occurred on:%d", index)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestNftTransferRequest() {
	nftid := "a10"
	testUri := "testUri"
	testUriHash := "testUriHash"
	mp := types.Params{
		Networks: []types.Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
				Active:    true,
			},
		},
	}
	suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, mp)
	modAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)

	testCases := []struct {
		msg      string
		expError string
		postMsg  types.MsgNftTransferRequest
		owner    sdk.AccAddress
	}{
		{
			"success",
			"",
			types.MsgNftTransferRequest{
				Sender:     ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:      nftid,
				EthAddress: "xxxxxx",
			},
			modAddr,
		},
		{
			"not owner",
			"not the owner of nft",
			types.MsgNftTransferRequest{
				Sender:     ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:      nftid,
				EthAddress: "xxxxxx",
			},
			modAddr,
		},
		{
			"not exits nft",
			"nft does not exist",
			types.MsgNftTransferRequest{
				Sender:     ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:      "xxxxxx",
				EthAddress: "xxxxxx",
			},
			suite.addrs[0],
		},
		// todo: implement nft listed case
		// {
		// 	"nft listed",
		// 	"nft does not exist",
		// 	types.MsgNftTransferRequest{
		// 		Sender:     ununifitypes.StringAccAddress(suite.addrs[0]),
		// 		NftId:      "xxxxxx",
		// 		EthAddress: "xxxxxx",
		// 	},
		// 	suite.addrs[0],
		// },
	}

	msg := types.MsgNftLocked{
		Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
		ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
		NftId:     nftid,
		Uri:       testUri,
		UriHash:   testUriHash,
	}

	err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
	suite.Require().NoError(err)
	for index, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := suite.Require()
			err := suite.app.DecentralizedvaultKeeper.NftTransferRequest(suite.ctx, &tc.postMsg)
			if tc.expError == "" {
				require.NoError(err)
				val := suite.app.NFTKeeper.GetOwner(suite.ctx, types.WrappedClassId, nftid)
				require.Equal(val, tc.owner, "the error occurred on:%d", index)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestNftRejectTransfer() {

	type testPostMsgType types.MsgNftRejectTransfer
	nftid := "a10"
	nftid2 := "a11"
	testUri := "testUri"
	testUriHash := "testUriHash"
	mp := types.Params{
		Networks: []types.Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
				Active:    true,
			},
		},
	}
	suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, mp)

	testCases := []struct {
		msg      string
		expError string
		postMsg  testPostMsgType
		owner    sdk.AccAddress
	}{
		{
			"not trusted sender",
			"sender does not mach oracle address",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:  nftid,
			},
			sdk.AccAddress(nil),
		},
		{
			"not exits nft",
			"nft does not exist",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:  "xxxxxx",
			},
			sdk.AccAddress(nil),
		},
		{
			"not exits nft",
			"nft does not exist",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:  "xxxxxx",
			},
			sdk.AccAddress(nil),
		},
		{
			"not deposit",
			"nft is not deposited",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:  nftid2,
			},
			suite.addrs[0],
		},
		{
			"success",
			"",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:  nftid,
			},
			suite.addrs[1],
		},
		// todo: implement nft listed case
		// {
		// 	"nft listed",
		// 	"nft does not exist",
		// 	types.MsgNftTransferRequest{
		// 		Sender:     ununifitypes.StringAccAddress(suite.addrs[0]),
		// 		NftId:      "xxxxxx",
		// 		EthAddress: "xxxxxx",
		// 	},
		// 	suite.addrs[0],
		// },
	}

	msg := types.MsgNftLocked{
		Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
		ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
		NftId:     nftid,
		Uri:       testUri,
		UriHash:   testUriHash,
	}

	err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
	suite.Require().NoError(err)

	msg = types.MsgNftLocked{
		Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
		ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
		NftId:     nftid2,
		Uri:       testUri,
		UriHash:   testUriHash,
	}

	err = suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
	suite.Require().NoError(err)
	err = suite.app.DecentralizedvaultKeeper.NftTransferRequest(suite.ctx, &types.MsgNftTransferRequest{
		Sender:     ununifitypes.StringAccAddress(suite.addrs[1]),
		EthAddress: "xxxxx",
		NftId:      nftid,
	})
	suite.Require().NoError(err)

	for index, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := suite.Require()
			postMsg := types.MsgNftRejectTransfer(tc.postMsg)
			err := suite.app.DecentralizedvaultKeeper.NftRejectTransfer(suite.ctx, &postMsg)
			if tc.expError == "" {
				require.NoError(err)
				val := suite.app.NFTKeeper.GetOwner(suite.ctx, types.WrappedClassId, nftid)
				require.Equal(val, tc.owner, "the error occurred on:%d", index)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestNftTransferred() {

	type testPostMsgType types.MsgNftTransferred
	nftid := "a10"
	nftid2 := "a11"
	testUri := "testUri"
	testUriHash := "testUriHash"
	mp := types.Params{
		Networks: []types.Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
				Active:    true,
			},
		},
	}
	suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, mp)
	modAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)

	testCases := []struct {
		msg      string
		expError string
		postMsg  testPostMsgType
		owner    sdk.AccAddress
	}{
		{
			"not trusted sender",
			"sender does not mach oracle address",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:  nftid,
			},
			modAddr,
		},
		{
			"not exits nft",
			"nft is not deposited",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:  "xxxxxx",
			},
			sdk.AccAddress(nil),
		},
		{
			"not deposit",
			"nft is not deposited",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:  nftid2,
			},
			suite.addrs[1],
		},
		{
			"success",
			"",
			testPostMsgType{
				Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
				NftId:  nftid,
			},
			sdk.AccAddress(nil),
		},
		// todo: implement nft listed case
		// {
		// 	"nft listed",
		// 	"nft does not exist",
		// 	types.MsgNftTransferRequest{
		// 		Sender:     ununifitypes.StringAccAddress(suite.addrs[0]),
		// 		NftId:      "xxxxxx",
		// 		EthAddress: "xxxxxx",
		// 	},
		// 	suite.addrs[0],
		// },
	}

	msg := types.MsgNftLocked{
		Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
		ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
		NftId:     nftid,
		Uri:       testUri,
		UriHash:   testUriHash,
	}

	err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
	suite.Require().NoError(err)

	msg = types.MsgNftLocked{
		Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
		ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
		NftId:     nftid2,
		Uri:       testUri,
		UriHash:   testUriHash,
	}

	err = suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
	suite.Require().NoError(err)
	err = suite.app.DecentralizedvaultKeeper.NftTransferRequest(suite.ctx, &types.MsgNftTransferRequest{
		Sender:     ununifitypes.StringAccAddress(suite.addrs[1]),
		EthAddress: "xxxxx",
		NftId:      nftid,
	})
	suite.Require().NoError(err)

	for index, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := suite.Require()
			postMsg := types.MsgNftTransferred(tc.postMsg)
			err := suite.app.DecentralizedvaultKeeper.NftTransferred(suite.ctx, &postMsg)
			if tc.expError == "" {
				require.NoError(err)
				val := suite.app.NFTKeeper.GetOwner(suite.ctx, types.WrappedClassId, postMsg.NftId)
				require.Equal(val, tc.owner, "the error occurred on:%d", index)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
				if !tc.owner.Equals(sdk.AccAddress(nil)) {
					val := suite.app.NFTKeeper.GetOwner(suite.ctx, types.WrappedClassId, tc.postMsg.NftId)
					require.Equal(val, tc.owner, "the error occurred on:%d", index)
				}
			}
		})
	}
}
