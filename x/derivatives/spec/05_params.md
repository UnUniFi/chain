# Params

`Params` is included in `GenesisState`. It has below three properties which will be explaned in each section.

```go
type Params struct {
	Pool             Pool                   `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool" yaml:"pool"`
	PerpetualFutures PerpetualFuturesParams `protobuf:"bytes,2,opt,name=perpetual_futures,json=perpetualFutures,proto3" json:"perpetual_futures" yaml:"perpetual_futures"`
	PerpetualOptions PerpetualOptionsParams `protobuf:"bytes,3,opt,name=perpetual_options,json=perpetualOptions,proto3" json:"perpetual_options" yaml:"perpetual_options"`
}
```

## Pool

```go
type Pool struct {
	QuoteTicker                       string                                 `protobuf:"bytes,1,opt,name=quote_ticker,json=quoteTicker,proto3" json:"quote_ticker,omitempty" yaml:"quote_ticker"`
	BaseLptMintFee                    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=base_lpt_mint_fee,json=baseLptMintFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_lpt_mint_fee" yaml:"base_lpt_mint_fee"`
	BaseLptRedeemFee                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=base_lpt_redeem_fee,json=baseLptRedeemFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_lpt_redeem_fee" yaml:"base_lpt_redeem_fee"`
	BorrowingFeeRatePerHour           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=borrowing_fee_rate_per_hour,json=borrowingFeeRatePerHour,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"borrowing_fee_rate_per_hour" yaml:"borrowing_fee_rate_per_hour"`
	LiquidationNeededReportRewardRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=liquidation_needed_report_reward_rate,json=liquidationNeededReportRewardRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_needed_report_reward_rate" yaml:"liquidation_needed_report_reward_rate"`
	AcceptedAssets                    []*Pool_Asset                          `protobuf:"bytes,6,rep,name=accepted_assets,json=acceptedAssets,proto3" json:"accepted_assets,omitempty" yaml:"accepted_assets"`
}
```

The tokens in `AcceptedAssets` have to have `DenomMetadata` in bank module.


## PerpetualFutures

```go
type PerpetualFuturesParams struct {
	CommissionRate                              github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=commission_rate,json=commissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission_rate" yaml:"commission_rate"`
	MarginMaintenanceRate                       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=margin_maintenance_rate,json=marginMaintenanceRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"margin_maintenance_rate" yaml:"margin_maintenance_rate"`
	ImaginaryFundingRateProportionalCoefficient github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=imaginary_funding_rate_proportional_coefficient,json=imaginaryFundingRateProportionalCoefficient,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"imaginary_funding_rate_proportional_coefficient" yaml:"imaginary_funding_rate_proportonal_coefficient"`
	Markets                                     []*Market                              `protobuf:"bytes,4,rep,name=markets,proto3" json:"markets,omitempty" yaml:"markets"`
}
```

## PerpetualOptioins

nothing is defined yet.
