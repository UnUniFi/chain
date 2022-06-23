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

<!-- todo: write section -->
### Balance

The `Balance` endpoint allows users to query account balance by address for a given denomination.

```sh
cosmos.bank.v1beta1.Query/Balance
```

Example:

```sh
grpcurl -plaintext \
    -d '{"address":"cosmos1..","denom":"stake"}' \
    localhost:9090 \
    cosmos.bank.v1beta1.Query/Balance
```

Example Output:

```json
{
  "balance": {
    "denom": "stake",
    "amount": "1000000000"
  }
}
```
