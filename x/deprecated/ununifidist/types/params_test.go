package types_test

import (
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifidisttypes "github.com/UnUniFi/chain/x/ununifidist/types"
)

type paramTest struct {
	params     ununifidisttypes.Params
	expectPass bool
}

type ParamTestSuite struct {
	suite.Suite

	tests []paramTest
}

func (suite *ParamTestSuite) SetupTest() {
	p1 := ununifidisttypes.Params{
		Active: true,
		Periods: ununifidisttypes.Periods{
			ununifidisttypes.Period{
				Start:     time.Date(2020, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
			},
			ununifidisttypes.Period{
				Start:     time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2022, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
			},
		},
	}
	p2 := ununifidisttypes.Params{
		Active: true,
		Periods: ununifidisttypes.Periods{
			ununifidisttypes.Period{
				Start:     time.Date(2022, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
			},
			ununifidisttypes.Period{
				Start:     time.Date(2023, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2024, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
			},
		},
	}
	p3 := ununifidisttypes.Params{
		Active: true,
		Periods: ununifidisttypes.Periods{
			ununifidisttypes.Period{
				Start:     time.Date(2020, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
			},
			ununifidisttypes.Period{
				Start:     time.Date(2020, time.March, 1, 1, 0, 0, 0, time.UTC),
				End:       time.Date(2022, time.March, 1, 1, 0, 0, 0, time.UTC),
				Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
			},
		},
	}

	suite.tests = []paramTest{
		{
			params:     p1,
			expectPass: true,
		},
		{
			params:     p2,
			expectPass: false,
		},
		{
			params:     p3,
			expectPass: false,
		},
	}
}

func (suite *ParamTestSuite) TestParamValidation() {
	for _, t := range suite.tests {
		err := t.params.Validate()
		if t.expectPass {
			suite.NoError(err)
		} else {
			suite.Error(err)
		}
	}
}

// func TestGenesisTestSuite(t *testing.T) {
// 	suite.Run(t, new(ParamTestSuite))
// }
