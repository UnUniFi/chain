package keeper_test

import (
	"fmt"
	"reflect"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

func (suite *KeeperTestSuite) Test_msgServer_NftLocked() {
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
		msg types.MsgNftLocked
	}
	tests := []struct {
		name       string
		args       args
		want       *types.MsgNftLockedResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			"success",
			args{
				types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				},
			},
			&types.MsgNftLockedResponse{},
			false,
			"",
		},
		{
			"error",
			args{
				types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[2]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				},
			},
			&types.MsgNftLockedResponse{},
			true,
			"sender does not mach oracle address",
		},
	}

	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			got, err := suite.msgServer.NftLocked(suite.ctx, &tt.args.msg)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("msgServer.NftLocked() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("msgServer.NftLocked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_msgServer_NftUnlocked() {
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
		msg *types.MsgNftUnlocked
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		want       *types.MsgNftUnlockedResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			"success",
			func() error {
				suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &types.MsgNftLocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
					Uri:       testUri,
					UriHash:   testUriHash,
				})
				return nil
			},
			args{
				&types.MsgNftUnlocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
				},
			},
			&types.MsgNftUnlockedResponse{},
			false,
			"",
		},
		{
			"error not the owner of nft",
			func() error {
				return nil
			},
			args{
				&types.MsgNftUnlocked{
					Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
					ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:     nftid,
				},
			},
			&types.MsgNftUnlockedResponse{},
			true,
			"not the owner of nft",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)
			got, err := suite.msgServer.NftUnlocked(suite.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("msgServer.NftUnlocked() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("msgServer.NftUnlocked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_msgServer_NftTransferRequest() {
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
	msg := types.MsgNftLocked{
		Sender:    ununifitypes.StringAccAddress(suite.addrs[0]),
		ToAddress: ununifitypes.StringAccAddress(suite.addrs[1]),
		NftId:     nftid,
		Uri:       testUri,
		UriHash:   testUriHash,
	}
	suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, mp)
	err := suite.app.DecentralizedvaultKeeper.NftLocked(suite.ctx, &msg)
	suite.Require().NoError(err)

	type args struct {
		msg *types.MsgNftTransferRequest
	}
	tests := []struct {
		name       string
		preRun     func() error
		args       args
		want       *types.MsgNftTransferRequestResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			"error not owner",
			func() error {
				return nil
			},
			args{
				&types.MsgNftTransferRequest{
					Sender:     ununifitypes.StringAccAddress(suite.addrs[0]),
					NftId:      nftid,
					EthAddress: "xxxxxx",
				},
			},
			&types.MsgNftTransferRequestResponse{},
			true,
			"not the owner of nft",
		},
		{
			"success",
			func() error {
				return nil
			},
			args{
				&types.MsgNftTransferRequest{
					Sender:     ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:      nftid,
					EthAddress: "xxxxxx",
				},
			},
			&types.MsgNftTransferRequestResponse{},
			false,
			"",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			err := tt.preRun()
			suite.Require().NoError(err)

			got, err := suite.msgServer.NftTransferRequest(suite.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("msgServer.NftTransferRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("msgServer.NftTransferRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_msgServer_NftRejectTransfer() {
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
	suite.Require().NoError(err)

	type args struct {
		msg *types.MsgNftRejectTransfer
	}
	tests := []struct {
		name       string
		args       args
		want       *types.MsgNftRejectTransferResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			"error not trusted sender",
			args{
				&types.MsgNftRejectTransfer{
					Sender: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:  nftid,
				},
			},
			&types.MsgNftRejectTransferResponse{},
			true,
			"sender does not mach oracle address",
		},
		{
			"success",
			args{
				&types.MsgNftRejectTransfer{
					Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
					NftId:  nftid,
				},
			},
			&types.MsgNftRejectTransferResponse{},
			false,
			"",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			got, err := suite.msgServer.NftRejectTransfer(suite.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("msgServer.NftRejectTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("msgServer.NftRejectTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_msgServer_NftTransferred() {
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
	suite.Require().NoError(err)

	type args struct {
		msg *types.MsgNftTransferred
	}
	tests := []struct {
		name       string
		args       args
		want       *types.MsgNftTransferredResponse
		wantErr    bool
		wantErrMsg string
	}{
		{
			"error not trusted sender",
			args{
				&types.MsgNftTransferred{
					Sender: ununifitypes.StringAccAddress(suite.addrs[1]),
					NftId:  nftid,
				},
			},
			nil,
			true,
			"sender does not mach oracle address",
		},
		{
			"success",
			args{
				&types.MsgNftTransferred{
					Sender: ununifitypes.StringAccAddress(suite.addrs[0]),
					NftId:  nftid,
				},
			},
			&types.MsgNftTransferredResponse{},
			false,
			"",
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			got, err := suite.msgServer.NftTransferred(suite.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("msgServer.NftTransferred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				// error case
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tt.wantErrMsg)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("msgServer.NftTransferred() = %v, want %v", got, tt.want)
			}
		})
	}
}
