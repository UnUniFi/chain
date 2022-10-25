package keeper_test

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

func (suite *KeeperTestSuite) TestKeeper_GetSetParamSet() {
	type args struct {
		ctx sdk.Context
		mp  types.Params
	}
	tests := []struct {
		name string
		args args
		want types.Params
	}{
		{
			"Match case 1",
			args{
				suite.ctx,
				types.Params{
					Networks: []types.Network{
						{
							NetworkId: "Ethereum",
							Asset:     "ETH",
							Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
							Active:    true,
						},
					},
				},
			},
			types.Params{
				Networks: []types.Network{
					{
						NetworkId: "Ethereum",
						Asset:     "ETH",
						Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
						Active:    true,
					},
				},
			},
		},
		{
			"Match case 2",
			args{
				suite.ctx,
				types.Params{
					Networks: []types.Network{
						{
							NetworkId: "hoge",
							Asset:     "hoge",
							Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[2])},
							Active:    false,
						},
					},
				},
			},
			types.Params{
				Networks: []types.Network{
					{
						NetworkId: "hoge",
						Asset:     "hoge",
						Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[2])},
						Active:    false,
					},
				},
			},
		},
		{
			"Match case non set params",
			args{
				suite.ctx,
				types.Params{},
			},
			types.Params{},
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, tt.args.mp)
			if got := suite.app.DecentralizedvaultKeeper.GetParamSet(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				if got.String() == tt.want.String() {
					return
				}
				suite.T().Errorf("Keeper.GetParamSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetOracles() {
	type args struct {
		ctx       sdk.Context
		networkId string
		mp        types.Params
	}

	tests := []struct {
		name string
		args args
		want []sdk.AccAddress
	}{
		{
			"Match NetworkId case 1",
			args{
				suite.ctx,
				"Ethereum",
				types.Params{
					Networks: []types.Network{
						{
							NetworkId: "Ethereum",
							Asset:     "ETH",
							Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
							Active:    true,
						},
					},
				},
			},
			[]sdk.AccAddress{suite.addrs[0]},
		},
		{
			"Match NetworkId case 2",
			args{
				suite.ctx,
				"Ethereum",
				types.Params{
					Networks: []types.Network{
						{
							NetworkId: "Ethereum",
							Asset:     "ETH",
							Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[2])},
							Active:    true,
						},
					},
				},
			},
			[]sdk.AccAddress{suite.addrs[2]},
		},
		{
			"No Match NetworkId",
			args{
				suite.ctx,
				"NoMatchId",
				types.Params{
					Networks: []types.Network{
						{
							NetworkId: "Ethereum",
							Asset:     "ETH",
							Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
							Active:    true,
						},
					},
				},
			},
			[]sdk.AccAddress{},
		},
		// {
		// 	"GetNetworks = 0",
		// 	args{
		// 		suite.ctx,
		// 		"NoMatchId",
		// 		types.Params{
		// 			Networks: []types.Network{
		// 				{
		// 					NetworkId: "Ethereum",
		// 					Asset:     "ETH",
		// 					Oracles:   []ununifitypes.StringAccAddress{},
		// 					Active:    true,
		// 				},
		// 			},
		// 		},
		// 	},
		// 	[]sdk.AccAddress{},
		// },
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, tt.args.mp)
			if got := suite.app.DecentralizedvaultKeeper.GetOracles(tt.args.ctx, tt.args.networkId); !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("Keeper.GetOracles() = %v, want %v", got, tt.want)
			}
		})
	}
}
