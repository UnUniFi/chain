package keeper_test

import (
	"context"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

func (suite *KeeperTestSuite) TestKeeper_Params() {
	// params := types.DefaultParams()
	params := types.Params{
		Networks: []types.Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(suite.addrs[0])},
				Active:    true,
			},
		},
	}
	suite.app.DecentralizedvaultKeeper.SetParamSet(suite.ctx, params)

	type args struct {
		c   context.Context
		req *types.QueryParamsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QueryParamsResponse
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				sdk.WrapSDKContext(suite.ctx),
				&types.QueryParamsRequest{},
			},
			want: &types.QueryParamsResponse{Params: params},
		},
		{
			name: "Error",
			args: args{
				sdk.WrapSDKContext(suite.ctx),
				nil,
			},
			want:    &types.QueryParamsResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(fmt.Sprintf("Case %s", tt.name), func() {
			got, err := suite.app.DecentralizedvaultKeeper.Params(tt.args.c, tt.args.req)
			if (err != nil) != tt.wantErr {
				suite.T().Errorf("Keeper.Params() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && reflect.DeepEqual(err, status.Error(codes.InvalidArgument, "invalid request")) {
				// Error case
				return
			} else if !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("Keeper.Params() = %v, want %v", got, tt.want)
			}
		})
	}
}
