<!--
order: 6
-->

# Client

## CLI

A user can query and interact with the `nftmarket` module using the CLI.

### Query

The `query` commands allow users to query `nftmarket` state.

```sh
ununifid query nftmarket --help
```
<!-- todo: write section -->
#### classes

The `classes` endpoint allows users to query all listing nft series.

```sh
ununifid query nftmarket classes [flags]
```

Example:

```sh
ununifid query nftmarket classes 
```

Example Output:

```yml
classes:
- id: a10
  name: crypotpunk
  description: crypotpunk is awsome
  symbol: cryp
  uri: http...
  uriHash: xxxxx
  nft:
  - id: ax10
    uri: http...
    uriHash: xxxxx
  nftCount: 20
- id: b10
  name: ape
  description: ape is awsome
  symbol: ape
  uri: http...
  uriHash: xxxxx
  nft:
  - id: bx10
    uri: http...
    uriHash: xxxxx
  nftCount: 5
pagination:
  total: '2'
```

### Transactions

The `tx` commands allow users to interact with the `nftmarket` module.

```sh
ununifid tx nftmarket --help
```

#### listing

The `listing` command listing NFT.

```sh
ununifid tx nftmarket listing [class-id] [nft-id] [flags]
```

Example:

```sh
ununifid tx nftmarket listing a10 a10 --from myKeyName --chain-id ununifi-x
```

<!-- todo: write section -->
## gRPC

A user can query the `nftmarket` module using gRPC endpoints.

### Classes

The `Classes` endpoint allows users to query all listing nft series.

```sh
ununifif.nftmarket.v1beta1.Query/Classes
```

Example:

```sh
grpcurl -plaintext \
    -d '{"nftLimit":"1"}' \
    localhost:9090 \
    ununifif.nftmarket.v1beta1.Query/Classes
```

Example Output:

```json
{
  "classes": [
    {
      "id":"a10",
      "name":"crypotpunk",
      "description":"crypotpunk is awsome",
      "symbol":"cryp",
      "uri":"http...",
      "uriHash":"xxxxx",
      "nft":[
        {
          "id":"ax10",
          "uri":"http...",
          "uriHash":"xxxxx"
        }
      ],
      "nftCount":20
    },
    {
      "id":"b10",
      "name":"ape",
      "description":"ape is awsome",
      "symbol":"ape",
      "uri":"http...",
      "uriHash":"xxxxx",
      "nft":[
        {
          "id":"bx10",
          "uri":"http...",
          "uriHash":"xxxxx"
        }
      ],
      "nftCount":5
    }
  ],
  "pagination": {
    "total": "2"
  }
}
```

### Class

The `Class` endpoint allows users to query listing nft series.

```sh
ununifif.nftmarket.v1beta1.Query/Class
```

Example:

```sh
grpcurl -plaintext \
    -d '{"classId":"a10", "nftLimit":"1"}' \
    localhost:9090 \
    ununifif.nftmarket.v1beta1.Query/Class
```

Example Output:

```json
{
  "class": {
      "id":"a10",
      "name":"crypotpunk",
      "description":"crypotpunk is awsome",
      "symbol":"cryp",
      "uri":"http...",
      "uriHash":"xxxxx",
      "nft":[
        {
          "id":"a10",
          "name":"crypotpunk",
          "description":"crypotpunk is awsome",
          "symbol":"cryp",
          "uri":"http...",
          "uriHash":"xxxxx",
        }
      ],
      "nftCount":20,
  }
}
```

### NFT

The `NFT` endpoint allows users to query nft.

```sh
ununifif.nftmarket.v1beta1.Query/Nft
```

Example:

```sh
grpcurl -plaintext \
    -d '{"classId":"a10","nftId":"a10"}' \
    localhost:9090 \
    ununifif.nftmarket.v1beta1.Query/Nft
```

Example Output:

```json
{
  "nft":{
    "id":"a10",
    "name":"crypotpunk",
    "description":"crypotpunk is awsome",
    "symbol":"cryp",
    "uri":"http...",
    "uriHash":"xxxxx",
    "uriHash":"xxxxx",
    "listingType":"DIRECT_ASSET_BORROW",
    "bidToken":"uguu",
    "state":"BIDDING",
    "minBid":"1",
    "bidActiverank":"2",
    "bids":[
      {
        "bidder":"ununifi1...",
        "amount":"100uguu",
      },
      {
        "bidder":"ununifi1...",
        "amount":"99uguu",
      },
      {
        "bidder":"ununifi1...",
        "amount":"98uguu",
      }
    ],
  }
}
```
