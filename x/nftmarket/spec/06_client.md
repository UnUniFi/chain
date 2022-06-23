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
#### balances

The `balances` command allows users to query account balances by address.

```sh
ununifid query bank balances [address] [flags]
```

Example:

```sh
ununifid query bank balances cosmos1..
```

Example Output:

```yml
balances:
- amount: "1000000000"
  denom: stake
pagination:
  next_key: null
  total: "0"
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
  ],
  "pagination": {
    "total": "1"
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
    -d '{"nftLimit":"1"}' \
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
