package decentralizedvault_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/UnUniFi/chain/app"
	ununifitypes "github.com/UnUniFi/chain/types"
	decentralizedvault "github.com/UnUniFi/chain/x/decentralized-vault"
	"github.com/UnUniFi/chain/x/decentralized-vault/keeper"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

func TestInitGenesis(t *testing.T) {
	isCheckTx := false
	app := simapp.Setup(t, isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})

	type args struct {
		ctx      sdk.Context
		k        keeper.Keeper
		genState types.GenesisState
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"success",
			args{
				ctx,
				app.DecentralizedvaultKeeper,
				types.GenesisState{
					Params: types.DefaultParams(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanics(t, func() {
				decentralizedvault.InitGenesis(tt.args.ctx, tt.args.k, tt.args.genState)
			})
		})
	}
}

func TestExportGenesis(t *testing.T) {
	isCheckTx := false
	app := simapp.Setup(t, isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(30000000))
	mp := types.Params{
		Networks: []types.Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{ununifitypes.StringAccAddress(addrs[0])},
				Active:    true,
			},
		},
	}
	app.DecentralizedvaultKeeper.SetParamSet(ctx, mp)

	type args struct {
		ctx sdk.Context
		k   keeper.Keeper
	}
	tests := []struct {
		name string
		args args
		want *types.GenesisState
	}{
		{
			"success",
			args{
				ctx,
				app.DecentralizedvaultKeeper,
			},
			&types.GenesisState{
				Params: mp,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decentralizedvault.ExportGenesis(tt.args.ctx, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExportGenesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
