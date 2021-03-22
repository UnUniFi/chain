<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [auction/auction.proto](#auction/auction.proto)
    - [BaseAuction](#jpyx.auction.BaseAuction)
    - [CollateralAuction](#jpyx.auction.CollateralAuction)
    - [DebtAuction](#jpyx.auction.DebtAuction)
    - [MsgPlaceBid](#jpyx.auction.MsgPlaceBid)
    - [Params](#jpyx.auction.Params)
    - [SurplusAuction](#jpyx.auction.SurplusAuction)
    - [WeightedAddresses](#jpyx.auction.WeightedAddresses)
  
- [auction/genesis.proto](#auction/genesis.proto)
    - [GenesisState](#jpyx.auction.GenesisState)
  
- [auction/query.proto](#auction/query.proto)
    - [QueryAllAuctionRequest](#jpyx.auction.QueryAllAuctionRequest)
    - [QueryAllAuctionResponse](#jpyx.auction.QueryAllAuctionResponse)
    - [QueryGetAuctionRequest](#jpyx.auction.QueryGetAuctionRequest)
    - [QueryGetAuctionResponse](#jpyx.auction.QueryGetAuctionResponse)
    - [QueryParamsRequest](#jpyx.auction.QueryParamsRequest)
    - [QueryParamsResponse](#jpyx.auction.QueryParamsResponse)
  
    - [Query](#jpyx.auction.Query)
  
- [cdp/cdp.proto](#cdp/cdp.proto)
    - [AugmentedCdp](#jpyx.cdp.AugmentedCdp)
    - [Cdp](#jpyx.cdp.Cdp)
    - [CollateralParam](#jpyx.cdp.CollateralParam)
    - [DebtParam](#jpyx.cdp.DebtParam)
    - [Deposit](#jpyx.cdp.Deposit)
    - [MsgCreateCdp](#jpyx.cdp.MsgCreateCdp)
    - [MsgDeposit](#jpyx.cdp.MsgDeposit)
    - [MsgDrawDebt](#jpyx.cdp.MsgDrawDebt)
    - [MsgLiquidate](#jpyx.cdp.MsgLiquidate)
    - [MsgRepayDebt](#jpyx.cdp.MsgRepayDebt)
    - [MsgWithdraw](#jpyx.cdp.MsgWithdraw)
    - [Params](#jpyx.cdp.Params)
  
- [cdp/genesis.proto](#cdp/genesis.proto)
    - [GenesisAccumulationTime](#jpyx.cdp.GenesisAccumulationTime)
    - [GenesisState](#jpyx.cdp.GenesisState)
    - [GenesisTotalPrincipal](#jpyx.cdp.GenesisTotalPrincipal)
  
- [cdp/query.proto](#cdp/query.proto)
    - [QueryAllCdpRequest](#jpyx.cdp.QueryAllCdpRequest)
    - [QueryAllCdpResponse](#jpyx.cdp.QueryAllCdpResponse)
    - [QueryGetCdpRequest](#jpyx.cdp.QueryGetCdpRequest)
    - [QueryGetCdpResponse](#jpyx.cdp.QueryGetCdpResponse)
    - [QueryParamsRequest](#jpyx.cdp.QueryParamsRequest)
    - [QueryParamsResponse](#jpyx.cdp.QueryParamsResponse)
  
    - [Query](#jpyx.cdp.Query)
  
- [incentive/incentive.proto](#incentive/incentive.proto)
    - [BaseClaim](#jpyx.incentive.BaseClaim)
    - [BaseMultiClaim](#jpyx.incentive.BaseMultiClaim)
    - [JpyxMintingClaim](#jpyx.incentive.JpyxMintingClaim)
    - [MsgClaimJpyxMintingReward](#jpyx.incentive.MsgClaimJpyxMintingReward)
    - [Multiplier](#jpyx.incentive.Multiplier)
    - [Params](#jpyx.incentive.Params)
    - [RewardIndex](#jpyx.incentive.RewardIndex)
    - [RewardPeriod](#jpyx.incentive.RewardPeriod)
  
- [incentive/genesis.proto](#incentive/genesis.proto)
    - [GenesisAccumulationTime](#jpyx.incentive.GenesisAccumulationTime)
    - [GenesisState](#jpyx.incentive.GenesisState)
  
- [incentive/query.proto](#incentive/query.proto)
    - [QueryParamsRequest](#jpyx.incentive.QueryParamsRequest)
    - [QueryParamsResponse](#jpyx.incentive.QueryParamsResponse)
  
    - [Query](#jpyx.incentive.Query)
  
- [jsmndist/jsmndist.proto](#jsmndist/jsmndist.proto)
    - [Params](#jpyx.jsmndist.Params)
    - [Period](#jpyx.jsmndist.Period)
  
- [jsmndist/genesis.proto](#jsmndist/genesis.proto)
    - [GenesisState](#jpyx.jsmndist.GenesisState)
  
- [jsmndist/query.proto](#jsmndist/query.proto)
    - [QueryGetBalancesRequest](#jpyx.jsmndist.QueryGetBalancesRequest)
    - [QueryGetBalancesResponse](#jpyx.jsmndist.QueryGetBalancesResponse)
    - [QueryParamsRequest](#jpyx.jsmndist.QueryParamsRequest)
    - [QueryParamsResponse](#jpyx.jsmndist.QueryParamsResponse)
  
    - [Query](#jpyx.jsmndist.Query)
  
- [pricefeed/pricefeed.proto](#pricefeed/pricefeed.proto)
    - [CurrentPrice](#jpyx.pricefeed.CurrentPrice)
    - [Market](#jpyx.pricefeed.Market)
    - [MsgPostPrice](#jpyx.pricefeed.MsgPostPrice)
    - [Params](#jpyx.pricefeed.Params)
    - [PostedPrice](#jpyx.pricefeed.PostedPrice)
  
- [pricefeed/genesis.proto](#pricefeed/genesis.proto)
    - [GenesisState](#jpyx.pricefeed.GenesisState)
  
- [pricefeed/query.proto](#pricefeed/query.proto)
    - [QueryAllMarketRequest](#jpyx.pricefeed.QueryAllMarketRequest)
    - [QueryAllMarketResponse](#jpyx.pricefeed.QueryAllMarketResponse)
    - [QueryAllOracleRequest](#jpyx.pricefeed.QueryAllOracleRequest)
    - [QueryAllOracleResponse](#jpyx.pricefeed.QueryAllOracleResponse)
    - [QueryAllPriceRequest](#jpyx.pricefeed.QueryAllPriceRequest)
    - [QueryAllPriceResponse](#jpyx.pricefeed.QueryAllPriceResponse)
    - [QueryAllRawPriceRequest](#jpyx.pricefeed.QueryAllRawPriceRequest)
    - [QueryAllRawPriceResponse](#jpyx.pricefeed.QueryAllRawPriceResponse)
    - [QueryGetPriceRequest](#jpyx.pricefeed.QueryGetPriceRequest)
    - [QueryGetPriceResponse](#jpyx.pricefeed.QueryGetPriceResponse)
    - [QueryParamsRequest](#jpyx.pricefeed.QueryParamsRequest)
    - [QueryParamsResponse](#jpyx.pricefeed.QueryParamsResponse)
  
    - [Query](#jpyx.pricefeed.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auction/auction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/auction.proto



<a name="jpyx.auction.BaseAuction"></a>

### BaseAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `initiator` | [string](#string) |  |  |
| `lot` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidder` | [string](#string) |  |  |
| `bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `has_received_bids` | [bool](#bool) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `max_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="jpyx.auction.CollateralAuction"></a>

### CollateralAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#jpyx.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lot_returns` | [WeightedAddresses](#jpyx.auction.WeightedAddresses) |  |  |






<a name="jpyx.auction.DebtAuction"></a>

### DebtAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#jpyx.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jpyx.auction.MsgPlaceBid"></a>

### MsgPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jpyx.auction.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `increment_surplus` | [string](#string) |  |  |
| `increment_debt` | [string](#string) |  |  |
| `increment_collateral` | [string](#string) |  |  |






<a name="jpyx.auction.SurplusAuction"></a>

### SurplusAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#jpyx.auction.BaseAuction) |  |  |






<a name="jpyx.auction.WeightedAddresses"></a>

### WeightedAddresses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [string](#string) | repeated |  |
| `weights` | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/genesis.proto



<a name="jpyx.auction.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_auction_id` | [uint64](#uint64) |  |  |
| `params` | [Params](#jpyx.auction.Params) |  |  |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/query.proto



<a name="jpyx.auction.QueryAllAuctionRequest"></a>

### QueryAllAuctionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jpyx.auction.QueryAllAuctionResponse"></a>

### QueryAllAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jpyx.auction.QueryGetAuctionRequest"></a>

### QueryGetAuctionRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="jpyx.auction.QueryGetAuctionResponse"></a>

### QueryGetAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="jpyx.auction.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="jpyx.auction.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.auction.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jpyx.auction.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jpyx.auction.QueryParamsRequest) | [QueryParamsResponse](#jpyx.auction.QueryParamsResponse) |  | GET|/jpyx/auction/params|
| `Auction` | [QueryGetAuctionRequest](#jpyx.auction.QueryGetAuctionRequest) | [QueryGetAuctionResponse](#jpyx.auction.QueryGetAuctionResponse) | this line is used by starport scaffolding # 2 | GET|/jpyx/auction/auctions/{id}|
| `AuctionAll` | [QueryAllAuctionRequest](#jpyx.auction.QueryAllAuctionRequest) | [QueryAllAuctionResponse](#jpyx.auction.QueryAllAuctionResponse) |  | GET|/jpyx/auction/auctions|

 <!-- end services -->



<a name="cdp/cdp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/cdp.proto



<a name="jpyx.cdp.AugmentedCdp"></a>

### AugmentedCdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [Cdp](#jpyx.cdp.Cdp) |  |  |
| `collateral_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateralization_ratio` | [string](#string) |  |  |






<a name="jpyx.cdp.Cdp"></a>

### Cdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="jpyx.cdp.CollateralParam"></a>

### CollateralParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `liquidation_ratio` | [string](#string) |  |  |
| `debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `stability_fee` | [string](#string) |  |  |
| `auction_size` | [string](#string) |  |  |
| `liquidation_penalty` | [string](#string) |  |  |
| `prefix` | [uint32](#uint32) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `liquidation_market_id` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |
| `check_collateralization_index_count` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |






<a name="jpyx.cdp.DebtParam"></a>

### DebtParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `reference_asset` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `debt_floor` | [string](#string) |  |  |






<a name="jpyx.cdp.Deposit"></a>

### Deposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jpyx.cdp.MsgCreateCdp"></a>

### MsgCreateCdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="jpyx.cdp.MsgDeposit"></a>

### MsgDeposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="jpyx.cdp.MsgDrawDebt"></a>

### MsgDrawDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jpyx.cdp.MsgLiquidate"></a>

### MsgLiquidate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="jpyx.cdp.MsgRepayDebt"></a>

### MsgRepayDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jpyx.cdp.MsgWithdraw"></a>

### MsgWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="jpyx.cdp.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_params` | [CollateralParam](#jpyx.cdp.CollateralParam) | repeated |  |
| `debt_param` | [DebtParam](#jpyx.cdp.DebtParam) |  |  |
| `global_debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `surplus_auction_threshold` | [string](#string) |  |  |
| `surplus_auction_lot` | [string](#string) |  |  |
| `debt_auction_threshold` | [string](#string) |  |  |
| `debt_auction_lot` | [string](#string) |  |  |
| `circuit_breaker` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cdp/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/genesis.proto



<a name="jpyx.cdp.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="jpyx.cdp.GenesisState"></a>

### GenesisState
GenesisState defines the cdp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.cdp.Params) |  |  |
| `cdps` | [Cdp](#jpyx.cdp.Cdp) | repeated |  |
| `deposits` | [Deposit](#jpyx.cdp.Deposit) | repeated |  |
| `starting_cdp_id` | [uint64](#uint64) |  |  |
| `debt_denom` | [string](#string) |  |  |
| `gov_denom` | [string](#string) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#jpyx.cdp.GenesisAccumulationTime) | repeated |  |
| `total_principals` | [GenesisTotalPrincipal](#jpyx.cdp.GenesisTotalPrincipal) | repeated | this line is used by starport scaffolding # genesis/proto/state |






<a name="jpyx.cdp.GenesisTotalPrincipal"></a>

### GenesisTotalPrincipal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `total_principal` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cdp/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/query.proto



<a name="jpyx.cdp.QueryAllCdpRequest"></a>

### QueryAllCdpRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jpyx.cdp.QueryAllCdpResponse"></a>

### QueryAllCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [Cdp](#jpyx.cdp.Cdp) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jpyx.cdp.QueryGetCdpRequest"></a>

### QueryGetCdpRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="jpyx.cdp.QueryGetCdpResponse"></a>

### QueryGetCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [Cdp](#jpyx.cdp.Cdp) |  |  |






<a name="jpyx.cdp.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="jpyx.cdp.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.cdp.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jpyx.cdp.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jpyx.cdp.QueryParamsRequest) | [QueryParamsResponse](#jpyx.cdp.QueryParamsResponse) |  | GET|/jpyx/cdp/params|
| `Cdp` | [QueryGetCdpRequest](#jpyx.cdp.QueryGetCdpRequest) | [QueryGetCdpResponse](#jpyx.cdp.QueryGetCdpResponse) | this line is used by starport scaffolding # 2 | GET|/jpyx/cdp/cdps/{id}|
| `CdpAll` | [QueryAllCdpRequest](#jpyx.cdp.QueryAllCdpRequest) | [QueryAllCdpResponse](#jpyx.cdp.QueryAllCdpResponse) |  | GET|/jpyx/cdp/cdps|

 <!-- end services -->



<a name="incentive/incentive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/incentive.proto



<a name="jpyx.incentive.BaseClaim"></a>

### BaseClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jpyx.incentive.BaseMultiClaim"></a>

### BaseMultiClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="jpyx.incentive.JpyxMintingClaim"></a>

### JpyxMintingClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseClaim](#jpyx.incentive.BaseClaim) |  |  |
| `reward_indexes` | [RewardIndex](#jpyx.incentive.RewardIndex) | repeated |  |






<a name="jpyx.incentive.MsgClaimJpyxMintingReward"></a>

### MsgClaimJpyxMintingReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |






<a name="jpyx.incentive.Multiplier"></a>

### Multiplier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `months_lockup` | [int64](#int64) |  |  |
| `factor` | [string](#string) |  |  |






<a name="jpyx.incentive.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `jpyx_minting_reward_periods` | [RewardPeriod](#jpyx.incentive.RewardPeriod) | repeated |  |
| `claim_multipliers` | [Multiplier](#jpyx.incentive.Multiplier) | repeated |  |
| `claim_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="jpyx.incentive.RewardIndex"></a>

### RewardIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_factor` | [string](#string) |  |  |






<a name="jpyx.incentive.RewardPeriod"></a>

### RewardPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/genesis.proto



<a name="jpyx.incentive.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="jpyx.incentive.GenesisState"></a>

### GenesisState
GenesisState defines the incentive module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.incentive.Params) |  |  |
| `jpyx_accumulation_times` | [GenesisAccumulationTime](#jpyx.incentive.GenesisAccumulationTime) | repeated |  |
| `jpyx_minting_claims` | [JpyxMintingClaim](#jpyx.incentive.JpyxMintingClaim) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/query.proto



<a name="jpyx.incentive.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="jpyx.incentive.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.incentive.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jpyx.incentive.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jpyx.incentive.QueryParamsRequest) | [QueryParamsResponse](#jpyx.incentive.QueryParamsResponse) | this line is used by starport scaffolding # 2 | GET|/jpyx/incentive/params|

 <!-- end services -->



<a name="jsmndist/jsmndist.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## jsmndist/jsmndist.proto



<a name="jpyx.jsmndist.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `periods` | [Period](#jpyx.jsmndist.Period) | repeated |  |






<a name="jpyx.jsmndist.Period"></a>

### Period



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `inflation` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="jsmndist/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## jsmndist/genesis.proto



<a name="jpyx.jsmndist.GenesisState"></a>

### GenesisState
GenesisState defines the jsmndist module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.jsmndist.Params) |  |  |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="jsmndist/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## jsmndist/query.proto



<a name="jpyx.jsmndist.QueryGetBalancesRequest"></a>

### QueryGetBalancesRequest







<a name="jpyx.jsmndist.QueryGetBalancesResponse"></a>

### QueryGetBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="jpyx.jsmndist.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="jpyx.jsmndist.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.jsmndist.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jpyx.jsmndist.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jpyx.jsmndist.QueryParamsRequest) | [QueryParamsResponse](#jpyx.jsmndist.QueryParamsResponse) |  | GET|/jpyx/jsmndist/params|
| `Balances` | [QueryGetBalancesRequest](#jpyx.jsmndist.QueryGetBalancesRequest) | [QueryGetBalancesResponse](#jpyx.jsmndist.QueryGetBalancesResponse) | this line is used by starport scaffolding # 2 | GET|/jpyx/jsmndist/balances|

 <!-- end services -->



<a name="pricefeed/pricefeed.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/pricefeed.proto



<a name="jpyx.pricefeed.CurrentPrice"></a>

### CurrentPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="jpyx.pricefeed.Market"></a>

### Market



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [string](#string) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="jpyx.pricefeed.MsgPostPrice"></a>

### MsgPostPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="jpyx.pricefeed.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#jpyx.pricefeed.Market) | repeated |  |






<a name="jpyx.pricefeed.PostedPrice"></a>

### PostedPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="pricefeed/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/genesis.proto



<a name="jpyx.pricefeed.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.pricefeed.Params) |  |  |
| `posted_prices` | [PostedPrice](#jpyx.pricefeed.PostedPrice) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="pricefeed/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/query.proto



<a name="jpyx.pricefeed.QueryAllMarketRequest"></a>

### QueryAllMarketRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jpyx.pricefeed.QueryAllMarketResponse"></a>

### QueryAllMarketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#jpyx.pricefeed.Market) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jpyx.pricefeed.QueryAllOracleRequest"></a>

### QueryAllOracleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jpyx.pricefeed.QueryAllOracleResponse"></a>

### QueryAllOracleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [string](#string) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jpyx.pricefeed.QueryAllPriceRequest"></a>

### QueryAllPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jpyx.pricefeed.QueryAllPriceResponse"></a>

### QueryAllPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPrice](#jpyx.pricefeed.CurrentPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jpyx.pricefeed.QueryAllRawPriceRequest"></a>

### QueryAllRawPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jpyx.pricefeed.QueryAllRawPriceResponse"></a>

### QueryAllRawPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [PostedPrice](#jpyx.pricefeed.PostedPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jpyx.pricefeed.QueryGetPriceRequest"></a>

### QueryGetPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="jpyx.pricefeed.QueryGetPriceResponse"></a>

### QueryGetPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPrice](#jpyx.pricefeed.CurrentPrice) |  |  |






<a name="jpyx.pricefeed.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="jpyx.pricefeed.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jpyx.pricefeed.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jpyx.pricefeed.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jpyx.pricefeed.QueryParamsRequest) | [QueryParamsResponse](#jpyx.pricefeed.QueryParamsResponse) |  | GET|/jpyx/pricefeed/params|
| `MarketAll` | [QueryAllMarketRequest](#jpyx.pricefeed.QueryAllMarketRequest) | [QueryAllMarketResponse](#jpyx.pricefeed.QueryAllMarketResponse) | this line is used by starport scaffolding # 2 | GET|/jpyx/pricefeed/markets|
| `OracleAll` | [QueryAllOracleRequest](#jpyx.pricefeed.QueryAllOracleRequest) | [QueryAllOracleResponse](#jpyx.pricefeed.QueryAllOracleResponse) |  | GET|/jpyx/pricefeed/markets/{market_id}/oracles|
| `Price` | [QueryGetPriceRequest](#jpyx.pricefeed.QueryGetPriceRequest) | [QueryGetPriceResponse](#jpyx.pricefeed.QueryGetPriceResponse) |  | GET|/jpyx/pricefeed/markets/{market_id}/price|
| `PriceAll` | [QueryAllPriceRequest](#jpyx.pricefeed.QueryAllPriceRequest) | [QueryAllPriceResponse](#jpyx.pricefeed.QueryAllPriceResponse) |  | GET|/jpyx/pricefeed/prices|
| `RawPriceAll` | [QueryAllRawPriceRequest](#jpyx.pricefeed.QueryAllRawPriceRequest) | [QueryAllRawPriceResponse](#jpyx.pricefeed.QueryAllRawPriceResponse) |  | GET|/jpyx/pricefeed/markets/{market_id}/raw_prices|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
