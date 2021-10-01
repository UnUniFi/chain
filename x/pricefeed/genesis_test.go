package pricefeed_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/app"
	pricefeedkeeper "github.com/lcnem/jpyx/x/pricefeed/keeper"

	"github.com/lcnem/jpyx/types"

	"github.com/stretchr/testify/suite"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper pricefeedkeeper.Keeper
}

func (suite *GenesisTestSuite) TestValidGenState() {
	tApp := app.NewTestApp()

	suite.NotPanics(func() {
		tApp.InitializeFromGenesisStates(
			NewPricefeedGenStateMulti(tApp),
		)
	})
	_, addrs := app.GeneratePrivKeyAddressPairs(10)

	tApp = app.NewTestApp()
	suite.NotPanics(func() {
		tApp.InitializeFromGenesisStates(
			NewPricefeedGenStateWithOracles(tApp, types.StringAccAddresses(addrs)),
		)
	})
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
