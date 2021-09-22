<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [auction/auction.proto](#auction/auction.proto)
    - [BaseAuction](#botany.auction.BaseAuction)
    - [CollateralAuction](#botany.auction.CollateralAuction)
    - [DebtAuction](#botany.auction.DebtAuction)
    - [Params](#botany.auction.Params)
    - [SurplusAuction](#botany.auction.SurplusAuction)
    - [WeightedAddress](#botany.auction.WeightedAddress)
  
- [auction/genesis.proto](#auction/genesis.proto)
    - [GenesisState](#botany.auction.GenesisState)
  
- [auction/query.proto](#auction/query.proto)
    - [QueryAllAuctionRequest](#botany.auction.QueryAllAuctionRequest)
    - [QueryAllAuctionResponse](#botany.auction.QueryAllAuctionResponse)
    - [QueryGetAuctionRequest](#botany.auction.QueryGetAuctionRequest)
    - [QueryGetAuctionResponse](#botany.auction.QueryGetAuctionResponse)
    - [QueryParamsRequest](#botany.auction.QueryParamsRequest)
    - [QueryParamsResponse](#botany.auction.QueryParamsResponse)
  
    - [Query](#botany.auction.Query)
  
- [auction/tx.proto](#auction/tx.proto)
    - [MsgPlaceBid](#botany.auction.MsgPlaceBid)
    - [MsgPlaceBidResponse](#botany.auction.MsgPlaceBidResponse)
  
    - [Msg](#botany.auction.Msg)
  
- [botanydist/botanydist.proto](#botanydist/botanydist.proto)
    - [Params](#botany.botanydist.Params)
    - [Period](#botany.botanydist.Period)
  
- [botanydist/genesis.proto](#botanydist/genesis.proto)
    - [GenesisState](#botany.botanydist.GenesisState)
  
- [botanydist/query.proto](#botanydist/query.proto)
    - [QueryGetBalancesRequest](#botany.botanydist.QueryGetBalancesRequest)
    - [QueryGetBalancesResponse](#botany.botanydist.QueryGetBalancesResponse)
    - [QueryParamsRequest](#botany.botanydist.QueryParamsRequest)
    - [QueryParamsResponse](#botany.botanydist.QueryParamsResponse)
  
    - [Query](#botany.botanydist.Query)
  
- [cdp/cdp.proto](#cdp/cdp.proto)
    - [AugmentedCdp](#botany.cdp.AugmentedCdp)
    - [Cdp](#botany.cdp.Cdp)
    - [CollateralParam](#botany.cdp.CollateralParam)
    - [DebtParam](#botany.cdp.DebtParam)
    - [Deposit](#botany.cdp.Deposit)
    - [Params](#botany.cdp.Params)
  
- [cdp/genesis.proto](#cdp/genesis.proto)
    - [GenesisAccumulationTime](#botany.cdp.GenesisAccumulationTime)
    - [GenesisState](#botany.cdp.GenesisState)
    - [GenesisTotalPrincipal](#botany.cdp.GenesisTotalPrincipal)
  
- [cdp/query.proto](#cdp/query.proto)
    - [QueryAllAccountRequest](#botany.cdp.QueryAllAccountRequest)
    - [QueryAllAccountResponse](#botany.cdp.QueryAllAccountResponse)
    - [QueryAllCdpRequest](#botany.cdp.QueryAllCdpRequest)
    - [QueryAllCdpResponse](#botany.cdp.QueryAllCdpResponse)
    - [QueryAllDepositRequest](#botany.cdp.QueryAllDepositRequest)
    - [QueryAllDepositResponse](#botany.cdp.QueryAllDepositResponse)
    - [QueryGetCdpRequest](#botany.cdp.QueryGetCdpRequest)
    - [QueryGetCdpResponse](#botany.cdp.QueryGetCdpResponse)
    - [QueryParamsRequest](#botany.cdp.QueryParamsRequest)
    - [QueryParamsResponse](#botany.cdp.QueryParamsResponse)
  
    - [Query](#botany.cdp.Query)
  
- [cdp/tx.proto](#cdp/tx.proto)
    - [MsgCreateCdp](#botany.cdp.MsgCreateCdp)
    - [MsgCreateCdpResponse](#botany.cdp.MsgCreateCdpResponse)
    - [MsgDeposit](#botany.cdp.MsgDeposit)
    - [MsgDepositResponse](#botany.cdp.MsgDepositResponse)
    - [MsgDrawDebt](#botany.cdp.MsgDrawDebt)
    - [MsgDrawDebtResponse](#botany.cdp.MsgDrawDebtResponse)
    - [MsgLiquidate](#botany.cdp.MsgLiquidate)
    - [MsgLiquidateResponse](#botany.cdp.MsgLiquidateResponse)
    - [MsgRepayDebt](#botany.cdp.MsgRepayDebt)
    - [MsgRepayDebtResponse](#botany.cdp.MsgRepayDebtResponse)
    - [MsgWithdraw](#botany.cdp.MsgWithdraw)
    - [MsgWithdrawResponse](#botany.cdp.MsgWithdrawResponse)
  
    - [Msg](#botany.cdp.Msg)
  
- [incentive/incentive.proto](#incentive/incentive.proto)
    - [BaseClaim](#botany.incentive.BaseClaim)
    - [BaseMultiClaim](#botany.incentive.BaseMultiClaim)
    - [CdpMintingClaim](#botany.incentive.CdpMintingClaim)
    - [Multiplier](#botany.incentive.Multiplier)
    - [Params](#botany.incentive.Params)
    - [RewardIndex](#botany.incentive.RewardIndex)
    - [RewardPeriod](#botany.incentive.RewardPeriod)
  
- [incentive/genesis.proto](#incentive/genesis.proto)
    - [GenesisAccumulationTime](#botany.incentive.GenesisAccumulationTime)
    - [GenesisDenoms](#botany.incentive.GenesisDenoms)
    - [GenesisState](#botany.incentive.GenesisState)
  
- [incentive/query.proto](#incentive/query.proto)
    - [QueryParamsRequest](#botany.incentive.QueryParamsRequest)
    - [QueryParamsResponse](#botany.incentive.QueryParamsResponse)
  
    - [Query](#botany.incentive.Query)
  
- [incentive/tx.proto](#incentive/tx.proto)
    - [MsgClaimCdpMintingReward](#botany.incentive.MsgClaimCdpMintingReward)
    - [MsgClaimCdpMintingRewardResponse](#botany.incentive.MsgClaimCdpMintingRewardResponse)
  
    - [Msg](#botany.incentive.Msg)
  
- [pricefeed/pricefeed.proto](#pricefeed/pricefeed.proto)
    - [CurrentPrice](#botany.pricefeed.CurrentPrice)
    - [Market](#botany.pricefeed.Market)
    - [Params](#botany.pricefeed.Params)
    - [PostedPrice](#botany.pricefeed.PostedPrice)
  
- [pricefeed/genesis.proto](#pricefeed/genesis.proto)
    - [GenesisState](#botany.pricefeed.GenesisState)
  
- [pricefeed/query.proto](#pricefeed/query.proto)
    - [QueryAllMarketRequest](#botany.pricefeed.QueryAllMarketRequest)
    - [QueryAllMarketResponse](#botany.pricefeed.QueryAllMarketResponse)
    - [QueryAllOracleRequest](#botany.pricefeed.QueryAllOracleRequest)
    - [QueryAllOracleResponse](#botany.pricefeed.QueryAllOracleResponse)
    - [QueryAllPriceRequest](#botany.pricefeed.QueryAllPriceRequest)
    - [QueryAllPriceResponse](#botany.pricefeed.QueryAllPriceResponse)
    - [QueryAllRawPriceRequest](#botany.pricefeed.QueryAllRawPriceRequest)
    - [QueryAllRawPriceResponse](#botany.pricefeed.QueryAllRawPriceResponse)
    - [QueryGetPriceRequest](#botany.pricefeed.QueryGetPriceRequest)
    - [QueryGetPriceResponse](#botany.pricefeed.QueryGetPriceResponse)
    - [QueryParamsRequest](#botany.pricefeed.QueryParamsRequest)
    - [QueryParamsResponse](#botany.pricefeed.QueryParamsResponse)
  
    - [Query](#botany.pricefeed.Query)
  
- [pricefeed/tx.proto](#pricefeed/tx.proto)
    - [MsgPostPrice](#botany.pricefeed.MsgPostPrice)
    - [MsgPostPriceResponse](#botany.pricefeed.MsgPostPriceResponse)
  
    - [Msg](#botany.pricefeed.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auction/auction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/auction.proto



<a name="botany.auction.BaseAuction"></a>

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






<a name="botany.auction.CollateralAuction"></a>

### CollateralAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#botany.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lot_returns` | [WeightedAddress](#botany.auction.WeightedAddress) | repeated |  |






<a name="botany.auction.DebtAuction"></a>

### DebtAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#botany.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="botany.auction.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `increment_surplus` | [string](#string) |  |  |
| `increment_debt` | [string](#string) |  |  |
| `increment_collateral` | [string](#string) |  |  |






<a name="botany.auction.SurplusAuction"></a>

### SurplusAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#botany.auction.BaseAuction) |  |  |






<a name="botany.auction.WeightedAddress"></a>

### WeightedAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/genesis.proto



<a name="botany.auction.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_auction_id` | [uint64](#uint64) |  |  |
| `params` | [Params](#botany.auction.Params) |  |  |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/query.proto



<a name="botany.auction.QueryAllAuctionRequest"></a>

### QueryAllAuctionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="botany.auction.QueryAllAuctionResponse"></a>

### QueryAllAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="botany.auction.QueryGetAuctionRequest"></a>

### QueryGetAuctionRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="botany.auction.QueryGetAuctionResponse"></a>

### QueryGetAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="botany.auction.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="botany.auction.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.auction.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.auction.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#botany.auction.QueryParamsRequest) | [QueryParamsResponse](#botany.auction.QueryParamsResponse) |  | GET|/botany/auction/params|
| `Auction` | [QueryGetAuctionRequest](#botany.auction.QueryGetAuctionRequest) | [QueryGetAuctionResponse](#botany.auction.QueryGetAuctionResponse) | this line is used by starport scaffolding # 2 | GET|/botany/auction/auctions/{id}|
| `AuctionAll` | [QueryAllAuctionRequest](#botany.auction.QueryAllAuctionRequest) | [QueryAllAuctionResponse](#botany.auction.QueryAllAuctionResponse) |  | GET|/botany/auction/auctions|

 <!-- end services -->



<a name="auction/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/tx.proto



<a name="botany.auction.MsgPlaceBid"></a>

### MsgPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="botany.auction.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.auction.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PlaceBid` | [MsgPlaceBid](#botany.auction.MsgPlaceBid) | [MsgPlaceBidResponse](#botany.auction.MsgPlaceBidResponse) |  | |

 <!-- end services -->



<a name="botanydist/botanydist.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## botanydist/botanydist.proto



<a name="botany.botanydist.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `periods` | [Period](#botany.botanydist.Period) | repeated |  |






<a name="botany.botanydist.Period"></a>

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



<a name="botanydist/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## botanydist/genesis.proto



<a name="botany.botanydist.GenesisState"></a>

### GenesisState
GenesisState defines the botanydist module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.botanydist.Params) |  |  |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `gov_denom` | [string](#string) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="botanydist/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## botanydist/query.proto



<a name="botany.botanydist.QueryGetBalancesRequest"></a>

### QueryGetBalancesRequest







<a name="botany.botanydist.QueryGetBalancesResponse"></a>

### QueryGetBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="botany.botanydist.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="botany.botanydist.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.botanydist.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.botanydist.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#botany.botanydist.QueryParamsRequest) | [QueryParamsResponse](#botany.botanydist.QueryParamsResponse) |  | GET|/botany/botanydist/params|
| `Balances` | [QueryGetBalancesRequest](#botany.botanydist.QueryGetBalancesRequest) | [QueryGetBalancesResponse](#botany.botanydist.QueryGetBalancesResponse) | this line is used by starport scaffolding # 2 | GET|/botany/botanydist/balances|

 <!-- end services -->



<a name="cdp/cdp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/cdp.proto



<a name="botany.cdp.AugmentedCdp"></a>

### AugmentedCdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [Cdp](#botany.cdp.Cdp) |  |  |
| `collateral_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateralization_ratio` | [string](#string) |  |  |






<a name="botany.cdp.Cdp"></a>

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






<a name="botany.cdp.CollateralParam"></a>

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






<a name="botany.cdp.DebtParam"></a>

### DebtParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `reference_asset` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `debt_floor` | [string](#string) |  |  |






<a name="botany.cdp.Deposit"></a>

### Deposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="botany.cdp.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_params` | [CollateralParam](#botany.cdp.CollateralParam) | repeated |  |
| `debt_param` | [DebtParam](#botany.cdp.DebtParam) |  |  |
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



<a name="botany.cdp.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="botany.cdp.GenesisState"></a>

### GenesisState
GenesisState defines the cdp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.cdp.Params) |  |  |
| `cdps` | [Cdp](#botany.cdp.Cdp) | repeated |  |
| `deposits` | [Deposit](#botany.cdp.Deposit) | repeated |  |
| `starting_cdp_id` | [uint64](#uint64) |  |  |
| `debt_denom` | [string](#string) |  |  |
| `gov_denom` | [string](#string) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#botany.cdp.GenesisAccumulationTime) | repeated |  |
| `total_principals` | [GenesisTotalPrincipal](#botany.cdp.GenesisTotalPrincipal) | repeated | this line is used by starport scaffolding # genesis/proto/state |






<a name="botany.cdp.GenesisTotalPrincipal"></a>

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



<a name="botany.cdp.QueryAllAccountRequest"></a>

### QueryAllAccountRequest







<a name="botany.cdp.QueryAllAccountResponse"></a>

### QueryAllAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="botany.cdp.QueryAllCdpRequest"></a>

### QueryAllCdpRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="botany.cdp.QueryAllCdpResponse"></a>

### QueryAllCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [AugmentedCdp](#botany.cdp.AugmentedCdp) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="botany.cdp.QueryAllDepositRequest"></a>

### QueryAllDepositRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="botany.cdp.QueryAllDepositResponse"></a>

### QueryAllDepositResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#botany.cdp.Deposit) | repeated |  |






<a name="botany.cdp.QueryGetCdpRequest"></a>

### QueryGetCdpRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="botany.cdp.QueryGetCdpResponse"></a>

### QueryGetCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [AugmentedCdp](#botany.cdp.AugmentedCdp) |  |  |






<a name="botany.cdp.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="botany.cdp.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.cdp.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.cdp.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#botany.cdp.QueryParamsRequest) | [QueryParamsResponse](#botany.cdp.QueryParamsResponse) |  | GET|/botany/cdp/params|
| `Cdp` | [QueryGetCdpRequest](#botany.cdp.QueryGetCdpRequest) | [QueryGetCdpResponse](#botany.cdp.QueryGetCdpResponse) | this line is used by starport scaffolding # 2 | GET|/botany/cdp/cdps/owners/{owner}/collateral-types/{collateral_type}/cdp|
| `CdpAll` | [QueryAllCdpRequest](#botany.cdp.QueryAllCdpRequest) | [QueryAllCdpResponse](#botany.cdp.QueryAllCdpResponse) |  | GET|/botany/cdp/cdps|
| `AccountAll` | [QueryAllAccountRequest](#botany.cdp.QueryAllAccountRequest) | [QueryAllAccountResponse](#botany.cdp.QueryAllAccountResponse) |  | GET|/botany/cdp/accounts|
| `DepositAll` | [QueryAllDepositRequest](#botany.cdp.QueryAllDepositRequest) | [QueryAllDepositResponse](#botany.cdp.QueryAllDepositResponse) |  | GET|/botany/cdp/deposits/owners/{owner}/collateral-types/{collateral_type}|

 <!-- end services -->



<a name="cdp/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/tx.proto



<a name="botany.cdp.MsgCreateCdp"></a>

### MsgCreateCdp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="botany.cdp.MsgCreateCdpResponse"></a>

### MsgCreateCdpResponse







<a name="botany.cdp.MsgDeposit"></a>

### MsgDeposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="botany.cdp.MsgDepositResponse"></a>

### MsgDepositResponse







<a name="botany.cdp.MsgDrawDebt"></a>

### MsgDrawDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="botany.cdp.MsgDrawDebtResponse"></a>

### MsgDrawDebtResponse







<a name="botany.cdp.MsgLiquidate"></a>

### MsgLiquidate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="botany.cdp.MsgLiquidateResponse"></a>

### MsgLiquidateResponse







<a name="botany.cdp.MsgRepayDebt"></a>

### MsgRepayDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="botany.cdp.MsgRepayDebtResponse"></a>

### MsgRepayDebtResponse







<a name="botany.cdp.MsgWithdraw"></a>

### MsgWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="botany.cdp.MsgWithdrawResponse"></a>

### MsgWithdrawResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.cdp.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateCdp` | [MsgCreateCdp](#botany.cdp.MsgCreateCdp) | [MsgCreateCdpResponse](#botany.cdp.MsgCreateCdpResponse) |  | |
| `Deposit` | [MsgDeposit](#botany.cdp.MsgDeposit) | [MsgDepositResponse](#botany.cdp.MsgDepositResponse) |  | |
| `Withdraw` | [MsgWithdraw](#botany.cdp.MsgWithdraw) | [MsgWithdrawResponse](#botany.cdp.MsgWithdrawResponse) |  | |
| `DrawDebt` | [MsgDrawDebt](#botany.cdp.MsgDrawDebt) | [MsgDrawDebtResponse](#botany.cdp.MsgDrawDebtResponse) |  | |
| `RepayDebt` | [MsgRepayDebt](#botany.cdp.MsgRepayDebt) | [MsgRepayDebtResponse](#botany.cdp.MsgRepayDebtResponse) |  | |
| `Liquidate` | [MsgLiquidate](#botany.cdp.MsgLiquidate) | [MsgLiquidateResponse](#botany.cdp.MsgLiquidateResponse) |  | |

 <!-- end services -->



<a name="incentive/incentive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/incentive.proto



<a name="botany.incentive.BaseClaim"></a>

### BaseClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="botany.incentive.BaseMultiClaim"></a>

### BaseMultiClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="botany.incentive.CdpMintingClaim"></a>

### CdpMintingClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseClaim](#botany.incentive.BaseClaim) |  |  |
| `reward_indexes` | [RewardIndex](#botany.incentive.RewardIndex) | repeated |  |






<a name="botany.incentive.Multiplier"></a>

### Multiplier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `months_lockup` | [int64](#int64) |  |  |
| `factor` | [string](#string) |  |  |






<a name="botany.incentive.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_minting_reward_periods` | [RewardPeriod](#botany.incentive.RewardPeriod) | repeated |  |
| `claim_multipliers` | [Multiplier](#botany.incentive.Multiplier) | repeated |  |
| `claim_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="botany.incentive.RewardIndex"></a>

### RewardIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_factor` | [string](#string) |  |  |






<a name="botany.incentive.RewardPeriod"></a>

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



<a name="botany.incentive.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="botany.incentive.GenesisDenoms"></a>

### GenesisDenoms



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `principal_denom` | [string](#string) |  |  |
| `cdp_minting_reward_denom` | [string](#string) |  |  |






<a name="botany.incentive.GenesisState"></a>

### GenesisState
GenesisState defines the incentive module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.incentive.Params) |  |  |
| `cdp_accumulation_times` | [GenesisAccumulationTime](#botany.incentive.GenesisAccumulationTime) | repeated |  |
| `cdp_minting_claims` | [CdpMintingClaim](#botany.incentive.CdpMintingClaim) | repeated |  |
| `denoms` | [GenesisDenoms](#botany.incentive.GenesisDenoms) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/query.proto



<a name="botany.incentive.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="botany.incentive.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.incentive.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.incentive.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#botany.incentive.QueryParamsRequest) | [QueryParamsResponse](#botany.incentive.QueryParamsResponse) | this line is used by starport scaffolding # 2 | GET|/botany/incentive/params|

 <!-- end services -->



<a name="incentive/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/tx.proto



<a name="botany.incentive.MsgClaimCdpMintingReward"></a>

### MsgClaimCdpMintingReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |






<a name="botany.incentive.MsgClaimCdpMintingRewardResponse"></a>

### MsgClaimCdpMintingRewardResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.incentive.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ClaimCdpMintingReward` | [MsgClaimCdpMintingReward](#botany.incentive.MsgClaimCdpMintingReward) | [MsgClaimCdpMintingRewardResponse](#botany.incentive.MsgClaimCdpMintingRewardResponse) |  | |

 <!-- end services -->



<a name="pricefeed/pricefeed.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/pricefeed.proto



<a name="botany.pricefeed.CurrentPrice"></a>

### CurrentPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="botany.pricefeed.Market"></a>

### Market



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [string](#string) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="botany.pricefeed.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#botany.pricefeed.Market) | repeated |  |






<a name="botany.pricefeed.PostedPrice"></a>

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



<a name="botany.pricefeed.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.pricefeed.Params) |  |  |
| `posted_prices` | [PostedPrice](#botany.pricefeed.PostedPrice) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="pricefeed/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/query.proto



<a name="botany.pricefeed.QueryAllMarketRequest"></a>

### QueryAllMarketRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="botany.pricefeed.QueryAllMarketResponse"></a>

### QueryAllMarketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#botany.pricefeed.Market) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="botany.pricefeed.QueryAllOracleRequest"></a>

### QueryAllOracleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="botany.pricefeed.QueryAllOracleResponse"></a>

### QueryAllOracleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [string](#string) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="botany.pricefeed.QueryAllPriceRequest"></a>

### QueryAllPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="botany.pricefeed.QueryAllPriceResponse"></a>

### QueryAllPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPrice](#botany.pricefeed.CurrentPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="botany.pricefeed.QueryAllRawPriceRequest"></a>

### QueryAllRawPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="botany.pricefeed.QueryAllRawPriceResponse"></a>

### QueryAllRawPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [PostedPrice](#botany.pricefeed.PostedPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="botany.pricefeed.QueryGetPriceRequest"></a>

### QueryGetPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="botany.pricefeed.QueryGetPriceResponse"></a>

### QueryGetPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPrice](#botany.pricefeed.CurrentPrice) |  |  |






<a name="botany.pricefeed.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="botany.pricefeed.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#botany.pricefeed.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.pricefeed.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#botany.pricefeed.QueryParamsRequest) | [QueryParamsResponse](#botany.pricefeed.QueryParamsResponse) |  | GET|/botany/pricefeed/params|
| `MarketAll` | [QueryAllMarketRequest](#botany.pricefeed.QueryAllMarketRequest) | [QueryAllMarketResponse](#botany.pricefeed.QueryAllMarketResponse) | this line is used by starport scaffolding # 2 | GET|/botany/pricefeed/markets|
| `OracleAll` | [QueryAllOracleRequest](#botany.pricefeed.QueryAllOracleRequest) | [QueryAllOracleResponse](#botany.pricefeed.QueryAllOracleResponse) |  | GET|/botany/pricefeed/markets/{market_id}/oracles|
| `Price` | [QueryGetPriceRequest](#botany.pricefeed.QueryGetPriceRequest) | [QueryGetPriceResponse](#botany.pricefeed.QueryGetPriceResponse) |  | GET|/botany/pricefeed/markets/{market_id}/price|
| `PriceAll` | [QueryAllPriceRequest](#botany.pricefeed.QueryAllPriceRequest) | [QueryAllPriceResponse](#botany.pricefeed.QueryAllPriceResponse) |  | GET|/botany/pricefeed/prices|
| `RawPriceAll` | [QueryAllRawPriceRequest](#botany.pricefeed.QueryAllRawPriceRequest) | [QueryAllRawPriceResponse](#botany.pricefeed.QueryAllRawPriceResponse) |  | GET|/botany/pricefeed/markets/{market_id}/raw_prices|

 <!-- end services -->



<a name="pricefeed/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/tx.proto



<a name="botany.pricefeed.MsgPostPrice"></a>

### MsgPostPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="botany.pricefeed.MsgPostPriceResponse"></a>

### MsgPostPriceResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="botany.pricefeed.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PostPriceResponse` | [MsgPostPriceResponse](#botany.pricefeed.MsgPostPriceResponse) | [MsgPostPriceResponse](#botany.pricefeed.MsgPostPriceResponse) |  | |

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
