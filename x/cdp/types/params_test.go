package types_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
)

type ParamsTestSuite struct {
	suite.Suite
}

func (suite *ParamsTestSuite) SetupTest() {
}

func (suite *ParamsTestSuite) TestParamValidation() {
	type args struct {
		collateralParams cdptypes.CollateralParams
		debtParams       cdptypes.DebtParams
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}

	testCases := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			name: "default",
			args: args{
				collateralParams: cdptypes.DefaultCollateralParams,
				debtParams:       cdptypes.DefaultDebtParams,
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "valid single-collateral",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 4000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "invalid single-collateral mismatched debt denoms",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "sjpy",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 4000000000000),
						DebtDenom:               "debtsjpy",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "does not match global debt denom",
			},
		},
		{
			name: "invalid single-collateral mismatched debt denoms",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 4000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "euu",
						ReferenceAsset:          "eur",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 4000000000000),
						DebtDenom:               "debteuu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "does not match global debt denom",
			},
		},
		{
			name: "invalid single-collateral over debt limit",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 1000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "exceeds global debt limit",
			},
		},
		{
			name: "invalid single-collateral over debt limit ver euu",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("euu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:eur",
						LiquidationMarketId:              "bnb:eur",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 21000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "euu",
						ReferenceAsset:          "eur",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("euu", 1000000000000),
						DebtDenom:               "debteuu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "exceeds global debt limit",
			},
		},
		{
			name: "valid multi-collateral",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "xrp:jpy",
						LiquidationMarketId:              "xrp:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "bnb",
						Type:                             "bnb-b",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("euu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x22,
						SpotMarketId:                     "bnb:eur",
						LiquidationMarketId:              "bnb:eur",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-b",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("euu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x23,
						SpotMarketId:                     "xrp:eur",
						LiquidationMarketId:              "xrp:eur",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 4000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "xxx",
						ReferenceAsset:          "yyy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("xxx", 2000000000000),
						DebtDenom:               "debtxxx",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "euu",
						ReferenceAsset:          "eur",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("euu", 4000000000000),
						DebtDenom:               "debteuu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "invalid multi-collateral over first debt limit",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "xrp:jpy",
						LiquidationMarketId:              "xrp:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "sum of collateral debt limits",
			},
		},
		{
			name: "invalid multi-collateral over last debt limit",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "xrp:jpy",
						LiquidationMarketId:              "xrp:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "bnb",
						Type:                             "bnb-b",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("euu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x22,
						SpotMarketId:                     "bnb:eur",
						LiquidationMarketId:              "bnb:eur",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-b",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("euu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x23,
						SpotMarketId:                     "xrp:eur",
						LiquidationMarketId:              "xrp:eur",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 4000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "xxx",
						ReferenceAsset:          "yyy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("xxx", 2000000000000),
						DebtDenom:               "debtxxx",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "euu",
						ReferenceAsset:          "eur",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("euu", 2000000000000),
						DebtDenom:               "debteuu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "sum of collateral debt limits",
			},
		},
		{
			name: "invalid multi-collateral over all debt limit",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "xrp:jpy",
						LiquidationMarketId:              "xrp:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "bnb",
						Type:                             "bnb-b",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("euu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x22,
						SpotMarketId:                     "bnb:eur",
						LiquidationMarketId:              "bnb:eur",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-b",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("euu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x23,
						SpotMarketId:                     "xrp:eur",
						LiquidationMarketId:              "xrp:eur",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "xxx",
						ReferenceAsset:          "yyy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("xxx", 2000000000000),
						DebtDenom:               "debtxxx",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "euu",
						ReferenceAsset:          "eur",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("euu", 2000000000000),
						DebtDenom:               "debteuu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "sum of collateral debt limits",
			},
		},
		{
			name: "invalid multi-collateral multiple debt denoms",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("sjpy", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "xrp:jpy",
						LiquidationMarketId:              "xrp:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 4000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "does not match global debt limit denom",
			},
		},
		{
			name: "invalid multi-collateral duplicate type",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("sjpy", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "xrp:jpy",
						LiquidationMarketId:              "xrp:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(6),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DefaultDebtParams,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "duplicate cdp collateral type",
			},
		},
		{
			name: "invalid collateral params empty denom",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "collateral denom invalid",
			},
		},
		{
			name: "invalid collateral params empty market id",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "",
						LiquidationMarketId:              "",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "market id cannot be blank",
			},
		},
		{
			name: "invalid collateral params duplicate denom + type",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "duplicate cdp collateral type",
			},
		},
		{
			name: "valid collateral params duplicate denom + different type",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "bnb",
						Type:                             "bnb-b",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x21,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "invalid collateral params duplicate prefix",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
					{
						Denom:                            "xrp",
						Type:                             "xrp-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "xrp:jpy",
						LiquidationMarketId:              "xrp:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "duplicate prefix for collateral denom",
			},
		},
		{
			name: "invalid collateral params nil debt limit",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.Coin{},
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debt limit for all collaterals should be positive",
			},
		},
		{
			name: "invalid collateral params liquidation ratio out of range",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("1.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "liquidation penalty should be between 0 and 1",
			},
		},
		{
			name: "invalid collateral params auction size zero",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.ZeroInt(),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "auction size should be positive",
			},
		},
		{
			name: "invalid collateral params stability fee out of range",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.1"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "stability fee must be ≥ 1.0",
			},
		},
		{
			name: "invalid debt param empty denom",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "",
						ReferenceAsset:          "eur",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("euu", 2000000000000),
						DebtDenom:               "debteuu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debtParam Denom invalid",
			},
		},
		{
			name: "invalid debt param empty dept denom",
			args: args{
				collateralParams: cdptypes.CollateralParams{
					{
						Denom:                            "bnb",
						Type:                             "bnb-a",
						LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
						DebtLimit:                        sdk.NewInt64Coin("jpu", 2000000000000),
						StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"),
						LiquidationPenalty:               sdk.MustNewDecFromStr("0.05"),
						AuctionSize:                      sdk.NewInt(50000000000),
						Prefix:                           0x20,
						SpotMarketId:                     "bnb:jpy",
						LiquidationMarketId:              "bnb:jpy",
						KeeperRewardPercentage:           sdk.MustNewDecFromStr("0.01"),
						ConversionFactor:                 sdk.NewInt(8),
						CheckCollateralizationIndexCount: sdk.NewInt(10),
					},
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
					{
						Denom:                   "euu",
						ReferenceAsset:          "eur",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("euu", 2000000000000),
						DebtDenom:               "",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debtParam DebtDenom invalid",
			},
		},
		{
			name: "nil debt limit",
			args: args{
				collateralParams: cdptypes.DefaultCollateralParams,
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.Coin{},
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "global debt limit <nil>: invalid coins",
			},
		},
		{
			name: "zero surplus auction threshold",
			args: args{
				collateralParams: cdptypes.DefaultCollateralParams,
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.ZeroInt(),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "surplus auction threshold should be positive",
			},
		},
		{
			name: "zero debt auction threshold",
			args: args{
				collateralParams: cdptypes.DefaultCollateralParams,
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.ZeroInt(),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debt auction threshold should be positive",
			},
		},
		{
			name: "zero surplus auction lot",
			args: args{
				collateralParams: cdptypes.DefaultCollateralParams,
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.ZeroInt(),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.NewInt(10000000000),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "surplus auction lot should be positive",
			},
		},
		{
			name: "zero debt auction lot",
			args: args{
				collateralParams: cdptypes.DefaultCollateralParams,
				debtParams: cdptypes.DebtParams{
					{
						Denom:                   "jpu",
						ReferenceAsset:          "jpy",
						ConversionFactor:        sdk.NewInt(6),
						DebtFloor:               sdk.NewInt(10000000),
						GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
						DebtDenom:               "debtjpu",
						SurplusAuctionThreshold: sdk.NewInt(500000000000),
						SurplusAuctionLot:       sdk.NewInt(10000000000),
						DebtAuctionThreshold:    sdk.NewInt(100000000000),
						DebtAuctionLot:          sdk.ZeroInt(),
						CircuitBreaker:          false,
					},
				},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debt auction lot should be positive",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			params := cdptypes.NewParams(
				tc.args.collateralParams,
				tc.args.debtParams,
			)
			err := params.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func TestParamsTestSuite(t *testing.T) {
	suite.Run(t, new(ParamsTestSuite))
}

func findTestSetup() (cdptypes.DebtParams, cdptypes.DebtParam, cdptypes.DebtParam) {
	type dps = cdptypes.DebtParams
	type dp = cdptypes.DebtParam
	jpu_debt := dp{
		Denom:            "jpu",
		ReferenceAsset:   "jpy",
		ConversionFactor: sdk.NewInt(6),
		DebtFloor:        sdk.NewInt(1),
		GlobalDebtLimit:  sdk.NewCoin("jpux", sdk.NewInt(100)),
		DebtDenom:        "debtjpu",
	}
	euu_debt := dp{
		Denom:            "euu",
		ReferenceAsset:   "eur",
		ConversionFactor: sdk.NewInt(6),
		DebtFloor:        sdk.NewInt(1),
		GlobalDebtLimit:  sdk.NewCoin("euux", sdk.NewInt(500)),
		DebtDenom:        "debteuu",
	}
	dummy_dept := dp{
		Denom:            "dum",
		ReferenceAsset:   "du",
		ConversionFactor: sdk.NewInt(6),
		DebtFloor:        sdk.NewInt(1),
		GlobalDebtLimit:  sdk.NewCoin("dum", sdk.NewInt(500)),
		DebtDenom:        "debtdum",
	}
	t1 := dps{jpu_debt, euu_debt, dummy_dept}
	return t1, jpu_debt, euu_debt
}
func TestFindDenom(t *testing.T) {
	type dp = cdptypes.DebtParam
	t1, jpu_debt, euu_debt := findTestSetup()

	result, exits := t1.FindDenom("jpu")
	except := jpu_debt
	assert.Equalf(t, true, exits, "not exists")
	assert.NotEqualf(t, euu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.Equalf(t, except, result, "not equal, except: %v\n result: %v", except, result)

	result, exits = t1.FindDenom("euu")
	except = euu_debt
	assert.Equalf(t, true, exits, "not exists")
	assert.NotEqualf(t, jpu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.Equalf(t, except, result, "not equal, except: %v\n result: %v", except, result)

	// not match test
	result, exits = t1.FindDenom("xxx")
	except = dp{}
	assert.Equalf(t, false, exits, "not exists")
	assert.NotEqualf(t, euu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.NotEqualf(t, jpu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.Equalf(t, except, result, "not equal, except: %v\n result: %v", except, result)
}

func TestFindGlobalDebtLimitDenom(t *testing.T) {
	type dp = cdptypes.DebtParam
	t1, jpu_debt, euu_debt := findTestSetup()

	result, exits := t1.FindGlobalDebtLimitDenom("jpux")
	except := jpu_debt
	assert.Equalf(t, true, exits, "not exists")
	assert.NotEqualf(t, euu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.Equalf(t, except, result, "not equal, except: %v\n result: %v", except, result)

	result, exits = t1.FindGlobalDebtLimitDenom("euux")
	except = euu_debt
	assert.Equalf(t, true, exits, "not exists")
	assert.NotEqualf(t, jpu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.Equalf(t, except, result, "not equal, except: %v\n result: %v", except, result)

	// not match test
	result, exits = t1.FindDenom("xxx")
	except = dp{}
	assert.Equalf(t, false, exits, "not exists")
	assert.NotEqualf(t, euu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.NotEqualf(t, jpu_debt, result, "except not equal, except: %v\n result: %v", except, result)
	assert.Equalf(t, except, result, "not equal, except: %v\n result: %v", except, result)
}
