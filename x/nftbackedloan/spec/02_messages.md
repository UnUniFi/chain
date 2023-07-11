<!--
order: 2
-->

# Messages

In this section we describe the processing of the nftbackedloan messages.

## MsgListNft

```protobuf
message MsgListNft {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
  ListingType listing_type = 3;
  string bid_token = 4;
  string min_bid = 5 [
    (gogoproto.moretags) = "yaml:\"min_bid\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable) = false
  ];
  uint64 bid_active_rank = 6;
}
```

## MsgCancelNftListing

```protobuf
message MsgCancelNftListing {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
}
```

## MsgNftBuyBack

```protobuf
message MsgNftBuyBack {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
}
```

## MsgExpandListingPeriod

```protobuf
message MsgExpandListingPeriod {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
}
```

## MsgPlaceBid

```protobuf
message MsgPlaceBid {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
  repeated cosmos.base.v1beta1.Coin amount = 3 [ (gogoproto.nullable) = false ];
}
```

## MsgCancelBid

```protobuf
message MsgCancelBid {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
}
```

## MsgEndNftListing

```protobuf
message MsgEndNftListing {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
}

```

## MsgPayFullBid

```protobuf
message MsgPayFullBid {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
}
```

## MsgBorrow

```protobuf
message MsgBorrow {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
  repeated cosmos.base.v1beta1.Coin amount = 3 [ (gogoproto.nullable) = false ];
}
```

## MsgRepay

```protobuf
message MsgRepay {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
  repeated cosmos.base.v1beta1.Coin amount = 3 [ (gogoproto.nullable) = false ];
}
```

## MsgMintStableCoin

```protobuf
message MsgMintStableCoin {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

## MsgBurnStableCoin

```protobuf
message MsgBurnStableCoin {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

## MsgLiquidate

```protobuf
message MsgLiquidate {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  NftIdentifier nft_id = 2 [ (gogoproto.nullable) = false ];
}
```
