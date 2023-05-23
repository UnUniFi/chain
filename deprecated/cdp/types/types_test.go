package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	types "github.com/UnUniFi/chain/deprecated/cdp/types"
)

type TypeTestSuite struct {
	suite.Suite
}

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

func (suite *TypeTestSuite) TestNewDebtDenomMap() {
	debt_params := makeMockData()
	debt_map := types.NewDebtDenomMap(debt_params)
	except := make(types.DebtDenomMap)
	except[debt_params[0].Denom] = debt_params[0].DebtDenom
	except[debt_params[1].Denom] = debt_params[1].DebtDenom

	suite.Equal(except, debt_map)
	suite.Panics(func() {
		types.NewDebtDenomMap(types.DebtParams{})
	})
	suite.Panics(func() {
		types.NewDebtDenomMap(types.DebtParams{
			types.DebtParam{
				Denom:            "jpu",
				ReferenceAsset:   "jpy",
				ConversionFactor: sdk.NewInt(6),
				DebtFloor:        sdk.NewInt(1),
				GlobalDebtLimit:  sdk.NewCoin("jpu", sdk.ZeroInt()),
			},
			types.DebtParam{
				Denom:            "euu",
				ReferenceAsset:   "eur",
				ConversionFactor: sdk.NewInt(6),
				DebtFloor:        sdk.NewInt(1),
				GlobalDebtLimit:  sdk.NewCoin("euu", sdk.ZeroInt()),
				DebtDenom:        "euudebt",
			},
		})
	})
	suite.Panics(func() {
		types.NewDebtDenomMap(types.DebtParams{
			types.DebtParam{
				Denom:            "jpu",
				ReferenceAsset:   "jpy",
				ConversionFactor: sdk.NewInt(6),
				DebtFloor:        sdk.NewInt(1),
				GlobalDebtLimit:  sdk.NewCoin("jpu", sdk.ZeroInt()),
				DebtDenom:        "debtjpu",
			},
			types.DebtParam{
				Denom:            "euu",
				ReferenceAsset:   "eur",
				ConversionFactor: sdk.NewInt(6),
				DebtFloor:        sdk.NewInt(1),
				GlobalDebtLimit:  sdk.NewCoin("euu", sdk.ZeroInt()),
			},
		})
	})
}

func (suite *TypeTestSuite) TestNewDenomMapFromByte() {
	mockStr := "{\"euu\":\"euudebt\",\"jpu\":\"jpudebt\"}"
	debt_params := makeMockData()
	debt_map := types.NewDebtDenomMapFromByte([]byte(mockStr))
	except := make(types.DebtDenomMap)
	except[debt_params[0].Denom] = debt_params[0].DebtDenom
	except[debt_params[1].Denom] = debt_params[1].DebtDenom

	suite.Equal(except, debt_map)
	suite.Panics(func() {
		types.NewDebtDenomMapFromByte([]byte{})
	})
	suite.Panics(func() {
		types.NewDebtDenomMapFromByte([]byte("not json"))
	})
}

func TestByte(t *testing.T) {
	debt_params := makeMockData()
	debt_map := types.NewDebtDenomMap(debt_params)
	except_string := "{\"euu\":\"euudebt\",\"jpu\":\"jpudebt\"}"
	assert.Equalf(t, []byte(except_string), debt_map.Byte(), "not match")
}

func TestDrawTestSuite(t *testing.T) {
	suite.Run(t, new(TypeTestSuite))
}
