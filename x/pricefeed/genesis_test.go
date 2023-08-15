package pricefeed_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/app"
	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"

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
	var strAddr []string
	for _, addr := range addrs {
		strAddr = append(strAddr, addr.String())
	}

	tApp = app.NewTestApp()
	suite.NotPanics(func() {
		tApp.InitializeFromGenesisStates(
			NewPricefeedGenStateWithOracles(tApp, strAddr),
		)
	})
}

// func TestGenesisTestSuite(t *testing.T) {
// 	suite.Run(t, new(GenesisTestSuite))
// }
