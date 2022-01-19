package types_test

import (
	"strings"
	"testing"

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
		surplusThreshold sdk.Int
		surplusLot       sdk.Int
		debtThreshold    sdk.Int
		debtLot          sdk.Int
		breaker          bool
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
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 4000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "sjpy",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 4000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 1000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
				},
				debtParams: cdptypes.DebtParams{
					{
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 4000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "invalid multi-collateral over debt limit",
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 4000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "does not match global debt limit denom",
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
						Denom:            "",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.NewInt64Coin("jpu", 2000000000000),
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debt denom invalid",
			},
		},
		{
			name: "nil debt limit",
			args: args{
				collateralParams: cdptypes.DefaultCollateralParams,
				debtParams: cdptypes.DebtParams{
					{
						Denom:            "jpu",
						ReferenceAsset:   "jpy",
						ConversionFactor: sdk.NewInt(6),
						DebtFloor:        sdk.NewInt(10000000),
						GlobalDebtLimit:  sdk.Coin{},
					},
				},
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
				debtParams:       cdptypes.DefaultDebtParams,
				surplusThreshold: sdk.ZeroInt(),
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
				debtParams:       cdptypes.DefaultDebtParams,
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    sdk.ZeroInt(),
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
				debtParams:       cdptypes.DefaultDebtParams,
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       sdk.ZeroInt(),
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          cdptypes.DefaultDebtLot,
				breaker:          cdptypes.DefaultCircuitBreaker,
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
				debtParams:       cdptypes.DefaultDebtParams,
				surplusThreshold: cdptypes.DefaultSurplusThreshold,
				surplusLot:       cdptypes.DefaultSurplusLot,
				debtThreshold:    cdptypes.DefaultDebtThreshold,
				debtLot:          sdk.ZeroInt(),
				breaker:          cdptypes.DefaultCircuitBreaker,
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
				tc.args.surplusThreshold,
				tc.args.surplusLot,
				tc.args.debtThreshold,
				tc.args.debtLot, tc.args.breaker,
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
