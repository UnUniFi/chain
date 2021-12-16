package types_test

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmtime "github.com/tendermint/tendermint/types/time"

	ununifitypes "github.com/UnUniFi/chain/types"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
)

type CdpValidationSuite struct {
	suite.Suite

	addrs []sdk.AccAddress
}

func (suite *CdpValidationSuite) SetupTest() {
	r := rand.New(rand.NewSource(12345))
	privkeySeed := make([]byte, 15)
	r.Read(privkeySeed)
	addr := sdk.AccAddress(secp256k1.GenPrivKeySecp256k1(privkeySeed).PubKey().Address())
	suite.addrs = []sdk.AccAddress{addr}
}

func (suite *CdpValidationSuite) TestCdpValidation() {
	type errArgs struct {
		expectPass bool
		contains   string
	}
	testCases := []struct {
		name    string
		cdp     cdptypes.Cdp
		errArgs errArgs
	}{
		{
			name: "valid cdp",
			cdp:  cdptypes.NewCdp(1, suite.addrs[0], sdk.NewInt64Coin("bnb", 100000), "bnb-a", sdk.NewInt64Coin("jpu", 100000), tmtime.Now(), sdk.Dec{}),
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "invalid cdp id",
			cdp:  cdptypes.NewCdp(0, suite.addrs[0], sdk.NewInt64Coin("bnb", 100000), "bnb-a", sdk.NewInt64Coin("jpu", 100000), tmtime.Now(), sdk.Dec{}),
			errArgs: errArgs{
				expectPass: false,
				contains:   "cdp id cannot be 0",
			},
		},
		{
			name: "invalid collateral",
			cdp:  cdptypes.Cdp{1, ununifitypes.StringAccAddress(suite.addrs[0]), "bnb-a", sdk.Coin{Denom: "", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(0)}, tmtime.Now(), sdk.Dec{}},
			errArgs: errArgs{
				expectPass: false,
				contains:   "collateral 100: invalid coins",
			},
		},
		{
			name: "invalid prinicpal",
			cdp:  cdptypes.Cdp{1, ununifitypes.StringAccAddress(suite.addrs[0]), "xrp-a", sdk.Coin{Denom: "xrp", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(0)}, tmtime.Now(), sdk.Dec{}},
			errArgs: errArgs{
				expectPass: false,
				contains:   "principal 100: invalid coins",
			},
		},
		{
			name: "invalid fees",
			cdp:  cdptypes.Cdp{1, ununifitypes.StringAccAddress(suite.addrs[0]), "xrp-a", sdk.Coin{Denom: "xrp", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "", Amount: sdk.NewInt(0)}, tmtime.Now(), sdk.Dec{}},
			errArgs: errArgs{
				expectPass: false,
				contains:   "accumulated fees 0: invalid coins",
			},
		},
		{
			name: "invalid fees updated",
			cdp:  cdptypes.Cdp{1, ununifitypes.StringAccAddress(suite.addrs[0]), "xrp-a", sdk.Coin{Denom: "xrp", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(0)}, time.Time{}, sdk.Dec{}},
			errArgs: errArgs{
				expectPass: false,
				contains:   "cdp updated fee time cannot be zero",
			},
		},
		{
			name: "invalid type",
			cdp:  cdptypes.Cdp{1, ununifitypes.StringAccAddress(suite.addrs[0]), "", sdk.Coin{Denom: "xrp", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(100)}, sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(0)}, tmtime.Now(), sdk.Dec{}},
			errArgs: errArgs{
				expectPass: false,
				contains:   "cdp type cannot be empty",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.cdp.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err, tc.name)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *CdpValidationSuite) TestDepositValidation() {
	type errArgs struct {
		expectPass bool
		contains   string
	}
	testCases := []struct {
		name    string
		deposit cdptypes.Deposit
		errArgs errArgs
	}{
		{
			name:    "valid deposit",
			deposit: cdptypes.NewDeposit(1, suite.addrs[0], sdk.NewInt64Coin("bnb", 1000000)),
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name:    "invalid cdp id",
			deposit: cdptypes.NewDeposit(0, suite.addrs[0], sdk.NewInt64Coin("bnb", 1000000)),
			errArgs: errArgs{
				expectPass: false,
				contains:   "deposit's cdp id cannot be 0",
			},
		},
		{
			name:    "empty depositor",
			deposit: cdptypes.NewDeposit(1, sdk.AccAddress{}, sdk.NewInt64Coin("bnb", 1000000)),
			errArgs: errArgs{
				expectPass: false,
				contains:   "depositor cannot be empty",
			},
		},
		{
			name:    "invalid deposit coins",
			deposit: cdptypes.NewDeposit(1, suite.addrs[0], sdk.Coin{Denom: "Invalid Denom", Amount: sdk.NewInt(1000000)}),
			errArgs: errArgs{
				expectPass: false,
				contains:   "deposit 1000000Invalid Denom: invalid coins",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.deposit.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err, tc.name)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *CdpValidationSuite) TestCdpGetTotalPrinciple() {
	principal := sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(100500)}
	accumulatedFees := sdk.Coin{Denom: "jpu", Amount: sdk.NewInt(25000)}

	cdp := cdptypes.Cdp{Principal: principal, AccumulatedFees: accumulatedFees}

	suite.Require().Equal(cdp.GetTotalPrincipal(), principal.Add(accumulatedFees))
}

func TestCdpValidationSuite(t *testing.T) {
	suite.Run(t, new(CdpValidationSuite))
}
