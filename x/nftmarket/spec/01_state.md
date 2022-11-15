<!--
order: 1
-->

# State

## NftListing

`NftListing` is created when a nft is listed for sell by owner.

```protobuf
enum ListingType {
  DIRECT_ASSET_BORROW = 0;
  SYNTHETIC_ASSET_CREATION = 1;
}
enum ListingState {
  SELLING = 0;
  BIDDING = 1;
  LIQUIDATION = 2;
  END_LISTING = 3;
  SUCCESSFUL_BID = 4;
}

message NftIdentifier {
  string class_id = 1;
  string nft_id = 2;
}

message NftListing {
  NftIdentifier nft_id = 1 [ (gogoproto.nullable) = false ];
  ListingType listing_type = 2;
  ListingState state = 3;
  string bid_token = 4;
  uint64 minimum_deposit_rate = 5;
  bool automatic_overdraft = 6;
}
```

- NftListing: `"nft_listing" | format(nft) -> NftListing`
- NftListing by address: `"address_nft_listing" | format(address) | format(nft) -> format(nft)`

## Bid

`Bid` is created when a bidder bid on a nft listing.

```protobuf
message NftBid {
  NftIdentifier nft_id = 1 [ (gogoproto.nullable) = false ];
  string bidder = 2;
  cosmos.base.v1beta1.Coin bid_amount = 3 [ (gogoproto.nullable) = false ];
  cosmos.base.v1beta1.Coin deposit_amount = 4 [ (gogoproto.nullable) = false ];
  google.protobuf.Timestamp bidding_period = 5 [ (gogoproto.nullable) = false ];
  uint64 deposit_lending_rate = 6 [ (gogoproto.nullable) = false ];
  bool automatic_payment = 7;
}
```

- Bid: `"nft_bid" | format(nft) | format(bidder) -> Bid`
- Bid by address: `"address_bid" | format(bidder) | format(nft) -> Bid`

## Loan

`Loan` is created when a nft lister make a loan from the protocol for using a listed nft as collateral.

```protobuf
message Loan {
  NftIdentifier nft_id = 1 [ (gogoproto.nullable) = false ];
  repeated cosmos.base.v1beta1.Coin loan = 2 [ (gogoproto.nullable) = false ];
}
```

- Loan: `"nft_loan" | format(nft) -> Loan`
- Loan by address: `"nft_loan" | format(address) | format(nft) -> Loan`

## Rewards

`Rewards` stores how many tokens are eligible for an address to claim.

- Rewards: `"rewards" | format(address) -> Coins`

## Params

`Params` describes global parameters that are maintained by governance.

```protobuf
message Params {
  string min_staking_for_listing = 1 [
    (gogoproto.moretags) = "yaml:\"min_staking_for_listing\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable) = false
  ];
  repeated string bid_tokens = 2;
  uint64 auto_relisting_count_if_no_bid = 3;
  uint64 nft_listing_delay_seconds = 4;
  uint64 nft_listing_period_initial = 5;
  uint64 nft_listing_cancel_required_seconds = 6;
  uint64 nft_listing_buy_back_extra_percentage = 7;
  uint64 nft_listing_gap_time = 8;
  uint64 bid_cancel_required_seconds = 9;
  uint64 bid_token_disburse_seconds_after_cancel = 10;
  uint64 nft_listing_full_payment_period = 11;
  uint64 nft_listing_nft_delivery_period = 12;
  uint64 nft_creator_share_percentage = 13;
  string market_administrator = 14;
  cosmos.base.v1beta1.Coin nft_listing_commission_fee = 15 [ (gogoproto.nullable) = false ];
  uint64 nft_listing_extend_seconds = 16;
  cosmos.base.v1beta1.Coin nft_listing_period_extend_fee_per_hour = 17 [ (gogoproto.nullable) = false ];
}
```
