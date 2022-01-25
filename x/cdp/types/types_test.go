package types_test

import (
	"testing"

	types "github.com/UnUniFi/chain/x/cdp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func makeMockData() types.DebtParams {
	return types.DebtParams{

		types.DebtParam{
			Denom:            "jpu",
			ReferenceAsset:   "jpy",
			ConversionFactor: sdk.NewInt(6),
			DebtFloor:        sdk.NewInt(1),
			GlobalDebtLimit:  sdk.NewCoin("jpu", sdk.ZeroInt()),
			DebtDenom:        "jpudebt",
		},
		types.DebtParam{
			Denom:            "euu",
			ReferenceAsset:   "eur",
			ConversionFactor: sdk.NewInt(6),
			DebtFloor:        sdk.NewInt(1),
			GlobalDebtLimit:  sdk.NewCoin("euu", sdk.ZeroInt()),
			DebtDenom:        "euudebt",
		},
	}
}

func TestNewDebtDenomMap(t *testing.T) {
	debt_params := makeMockData()
	debt_map := types.NewDebtDenomMap(debt_params)
	except := make(types.DebtDenomMap)
	except[debt_params[0].Denom] = debt_params[0].DebtDenom
	except[debt_params[1].Denom] = debt_params[1].DebtDenom
	assert.Equalf(t, except, debt_map, "not match")
	// todo test panic line
}

func TestByte(t *testing.T) {
	debt_params := makeMockData()
	debt_map := types.NewDebtDenomMap(debt_params)
	except_string := "{\"euu\":\"euudebt\",\"jpu\":\"jpudebt\"}"
	assert.Equalf(t, []byte(except_string), debt_map.Byte(), "not match")
}

// func ()
