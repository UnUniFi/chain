package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/UnUniFi/chain/x/auction/simulation"
	"github.com/UnUniFi/chain/x/auction/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func TestDecodeDistributionStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)
	c := sdk.NewInt64Coin

	surplusAuction := types.NewSurplusAuction("me", sdk.NewCoin("coin", sdk.OneInt()), "coin", time.Now().UTC()).WithID(1)
	surplusAuctionAny, _ := codectypes.NewAnyWithValue(surplusAuction)
	surplusBz, _ := surplusAuctionAny.Marshal()

	debtAuction := types.NewDebtAuction("buyerMod", c("denom", 12345678), c("anotherdenom", 12345678), time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC), c("debt", 12345678)).WithID(2)
	debtAuctionAny, _ := codectypes.NewAnyWithValue(debtAuction)
	debtBz, _ := debtAuctionAny.Marshal()

	collateralAuction := types.NewCollateralAuction("sellerMod", c("denom", 12345678), time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC), c("anotherdenom", 12345678), nil, c("debt", 12345678)).WithID(3)
	collateralAuctionAny, _ := codectypes.NewAnyWithValue(collateralAuction)
	collateralBz, _ := collateralAuctionAny.Marshal()

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   []byte(types.AuctionKey),
				Value: surplusBz,
			},
			{
				Key:   []byte(types.AuctionKey),
				Value: debtBz,
			},
			{
				Key:   []byte(types.AuctionKey),
				Value: collateralBz,
			},
			{
				Key:   []byte(types.AuctionByTimeKey),
				Value: sdk.Uint64ToBigEndian(2),
			},
			{
				Key:   []byte(types.NextAuctionIDKey),
				Value: sdk.Uint64ToBigEndian(10),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Auction", fmt.Sprintf("%v\n%v", surplusAuction, surplusAuction)},
		{"Auction", fmt.Sprintf("%v\n%v", debtAuction, debtAuction)},
		{"Auction", fmt.Sprintf("%v\n%v", collateralAuction, collateralAuction)},
		{"AuctionByTime", "2\n2"},
		{"NextAuctionI", "10\n10"},
		{"other", ""},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
