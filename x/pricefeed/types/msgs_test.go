package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmtime "github.com/cometbft/cometbft/types/time"
)

func TestMsgPlaceBid_ValidateBasic(t *testing.T) {
	addr := "someName"
	price, _ := sdk.NewDecFromStr("0.3005")
	expiry := tmtime.Now()
	negativePrice, _ := sdk.NewDecFromStr("-3.05")

	tests := []struct {
		name       string
		msg        MsgPostPrice
		expectPass bool
	}{
		{"normal", MsgPostPrice{addr, "xrp", price, expiry, sdk.NewCoin("uguu", sdk.NewInt(1000))}, true},
		{"emptyAddr", MsgPostPrice{"", "xrp", price, expiry, sdk.NewCoin("uguu", sdk.NewInt(1000))}, false},
		{"emptyAsset", MsgPostPrice{addr, "", price, expiry, sdk.NewCoin("uguu", sdk.NewInt(1000))}, false},
		{"negativePrice", MsgPostPrice{addr, "xrp", negativePrice, expiry, sdk.NewCoin("uguu", sdk.NewInt(1000))}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectPass {
				require.Nil(t, tc.msg.ValidateBasic())
			} else {
				require.NotNil(t, tc.msg.ValidateBasic())
			}
		})
	}
}
