package keeper_test

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
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
		{
			"exist nft",
			"nft already exist",
			types.MsgNftLocked{
				Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
				ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
				NftId:     nftid,
				Uri:       testUri,
				UriHash:   testUriHash,
			},
			suite.addrs[1],
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

func (suite *KeeperTestSuite) TestKeeper_IsTrustedSender() {
	type args struct {
		senderAddress sdk.AccAddress
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		want       bool
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "error not exist oracle",
			preRun: func() error {
				mp := types.Params{
					Networks: []types.Network{
						{
							NetworkId: "Ununifi",
							Asset:     "GUU",
							Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
							Active:    true,
						},
					},
				}
				suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, mp)
				return nil
			},
			args: args{
				suite.addrs[0],
			},
			want:       true,
			wantErr:    true,
			wantErrMsg: "oracle does not exist",
		},
		{
			name: "error",
			preRun: func() error {
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
				return nil
			},
			args: args{
				sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
			},
			want:       true,
			wantErr:    true,
			wantErrMsg: "sender does not mach oracle address",
		},
		{
			name: "success",
			preRun: func() error {
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
				return nil
			},
			args: args{
				suite.addrs[0],
			},
			want:       true,
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			got, err := suite.app.DecentralizedvaultKeeper.IsTrustedSender(suite.ctx, tt.args.senderAddress)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.IsTrustedSender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if got != tt.want {
				suite.T().Errorf("Keeper.IsTrustedSender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_MintWrappedNft() {
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
	sender := suite.addrs[0]
	receiver := suite.addrs[1]
	type args struct {
		nftId    string
		uri      string
		uriHash  string
		receiver sdk.AccAddress
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "success",
			preRun: func() error {
				_, err := suite.app.DecentralizedvaultKeeper.IsTrustedSender(suite.ctx, sender)
				return err
			},
			args: args{
				nftid,
				testUri,
				testUriHash,
				receiver,
			},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "error nft already exist",
			preRun: func() error {
				_, err := suite.app.DecentralizedvaultKeeper.IsTrustedSender(suite.ctx, sender)
				return err
			},
			args: args{
				nftid,
				testUri,
				testUriHash,
				receiver,
			},
			wantErr:    true,
			wantErrMsg: "nft already exist",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			err = suite.app.DecentralizedvaultKeeper.MintWrappedNft(
				suite.ctx, tt.args.nftId, tt.args.uri, tt.args.uriHash, tt.args.receiver)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.MintWrappedNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_BurnWrappedNft() {
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

	type args struct {
		nftId string
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "error not exist nft",
			preRun: func() error {
				return nil
			},
			args: args{
				"xxxx",
			},
			wantErr:    true,
			wantErrMsg: "nft does not exist",
		},
		{
			name: "success",
			preRun: func() error {
				msg := types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				}
				err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
				return err
			},
			args: args{
				nftid,
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			err = suite.app.DecentralizedvaultKeeper.BurnWrappedNft(suite.ctx, tt.args.nftId)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.BurnWrappedNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DepositWrappedNft() {
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

	type args struct {
		depositor  sdk.AccAddress
		nftId      string
		ethAddress string
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "error",
			preRun: func() error {
				return nil
			},
			args: args{
				suite.addrs[1],
				nftid,
				"xxxxxx",
			},
			wantErr:    true,
			wantErrMsg: "nft class does not exist",
		},
		{
			name: "success",
			preRun: func() error {
				msg := types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				}
				err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
				return err
			},
			args: args{
				suite.addrs[1],
				nftid,
				"xxxxxx",
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			err = suite.app.DecentralizedvaultKeeper.DepositWrappedNft(
				suite.ctx, tt.args.depositor, tt.args.nftId, tt.args.ethAddress)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.DepositWrappedNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_WithdrawWrappedNft() {
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

	type args struct {
		msg *types.MsgNftRejectTransfer
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "error",
			preRun: func() error {
				return nil
			},
			args: args{
				&types.MsgNftRejectTransfer{
					Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
					NftId:  nftid,
				},
			},
			wantErr:    true,
			wantErrMsg: "transfer request does not exist",
		},
		{
			name: "success",
			preRun: func() error {
				msg := types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				}
				err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)

				suite.Require().NoError(err)

				err = suite.app.DecentralizedvaultKeeper.NftTransferRequest(suite.ctx, &types.MsgNftTransferRequest{
					Sender:     ununifitypes.StringAccAddress(suite.addrs[1]),
					EthAddress: "xxxxx",
					NftId:      nftid,
				})
				return err
			},
			args: args{
				&types.MsgNftRejectTransfer{
					Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
					NftId:  nftid,
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			err = suite.app.DecentralizedvaultKeeper.WithdrawWrappedNft(suite.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.WithdrawWrappedNft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetTransferRequestByIdBytes() {
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

	type args struct {
		nftIdBytes []byte
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		want       types.TransferRequest
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "error",
			preRun: func() error {
				return nil
			},
			args: args{
				[]byte("xxxx"),
			},
			want: types.TransferRequest{
				Owner:      suite.addrs[1].String(),
				EthAddress: "xxxxx",
				NftId:      nftid,
			},
			wantErr:    true,
			wantErrMsg: "transfer request does not exist",
		},
		{
			name: "success",
			preRun: func() error {
				msg := types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				}
				err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)

				suite.Require().NoError(err)

				err = suite.app.DecentralizedvaultKeeper.NftTransferRequest(suite.ctx, &types.MsgNftTransferRequest{
					Sender:     ununifitypes.StringAccAddress(suite.addrs[1]),
					EthAddress: "xxxxx",
					NftId:      nftid,
				})
				return err
			},
			args: args{
				[]byte(nftid),
			},
			want: types.TransferRequest{
				Owner:      suite.addrs[1].String(),
				EthAddress: "xxxxx",
				NftId:      nftid,
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			got, err := suite.app.DecentralizedvaultKeeper.GetTransferRequestByIdBytes(suite.ctx, tt.args.nftIdBytes)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.GetTransferRequestByIdBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("Keeper.GetTransferRequestByIdBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_IsListedNft() {
	nftid := "a10"
	// testUri := "testUri"
	// testUriHash := "testUriHash"

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

	type args struct {
		nftId string
	}
	tests := []struct {
		name   string
		preRun func() error
		args   args
		want   bool
	}{
		{
			name: "success false",
			preRun: func() error {
				return nil
			},
			args: args{
				nftid,
			},
			want: false,
		},
		// {
		// 	name: "success true",
		// 	preRun: func() error {
		// 		now := time.Now().UTC()
		// 		owner := suite.addrs[0]
		// 		listings := []nftmarkettypes.NftListing{
		// 			{
		// 				NftId: nftmarkettypes.NftIdentifier{
		// 					ClassId: types.WrappedClassId,
		// 					NftId:   nftid,
		// 				},
		// 				Owner:              owner.String(),
		// 				ListingType:        nftmarkettypes.ListingType_DIRECT_ASSET_BORROW,
		// 				State:              nftmarkettypes.ListingState_LISTING,
		// 				BidToken:           "uguu",
		// 				MinBid:             sdk.OneInt(),
		// 				BidActiveRank:      1,
		// 				StartedAt:          now,
		// 				EndAt:              now,
		// 				FullPaymentEndAt:   time.Time{},
		// 				SuccessfulBidEndAt: time.Time{},
		// 				AutoRelistedCount:  0,
		// 			},
		// 		}

		// 		for _, listing := range listings {
		// 			suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)
		// 		}
		// 		for _, listing := range listings {
		// 			gotListing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, listing.IdBytes())
		// 			suite.Require().NoError(err)
		// 			suite.Require().Equal(listing, gotListing)
		// 		}
		// 		return nil
		// 	},
		// 	args: args{
		// 		nftid,
		// 	},
		// 	want: true,
		// },
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			if got := suite.app.DecentralizedvaultKeeper.IsListedNft(suite.ctx, tt.args.nftId); got != tt.want {
				suite.T().Errorf("Keeper.IsListedNft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetTransferRequests() {
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

	type args struct {
		limit int
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		want       []types.TransferRequest
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "success",
			preRun: func() error {
				msg := types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				}
				err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)

				suite.Require().NoError(err)

				err = suite.app.DecentralizedvaultKeeper.NftTransferRequest(suite.ctx, &types.MsgNftTransferRequest{
					Sender:     ununifitypes.StringAccAddress(suite.addrs[1]),
					EthAddress: "xxxxx",
					NftId:      nftid,
				})
				return err
			},
			args: args{
				10,
			},
			want: []types.TransferRequest{
				{
					Owner:      suite.addrs[1].String(),
					EthAddress: "xxxxx",
					NftId:      nftid,
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "success",
			preRun: func() error {
				msg := types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[2]),
					NftId:     nftid2,
					Uri:       testUri,
					UriHash:   testUriHash,
				}
				err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)

				suite.Require().NoError(err)

				err = suite.app.DecentralizedvaultKeeper.NftTransferRequest(suite.ctx, &types.MsgNftTransferRequest{
					Sender:     ununifitypes.StringAccAddress(suite.addrs[2]),
					EthAddress: "xxxxx",
					NftId:      nftid2,
				})
				return err
			},
			args: args{
				1,
			},
			want: []types.TransferRequest{
				{
					Owner:      suite.addrs[1].String(),
					EthAddress: "xxxxx",
					NftId:      nftid,
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "success",
			preRun: func() error {
				return nil
			},
			args: args{
				2,
			},
			want: []types.TransferRequest{
				{
					Owner:      suite.addrs[1].String(),
					EthAddress: "xxxxx",
					NftId:      nftid,
				},
				{
					Owner:      suite.addrs[2].String(),
					EthAddress: "xxxxx",
					NftId:      nftid2,
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			got, err := suite.app.DecentralizedvaultKeeper.GetTransferRequests(suite.ctx, tt.args.limit)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.GetTransferRequests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("Keeper.GetTransferRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
