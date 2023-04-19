# Yield-aggregator

The `yield-aggregator` module provides the function for yield aggregation.
Users deposit their funds to the **Vault**, and then this module uses the funds to earn yield automatically.

This module is the first yield aggregator that supports "interchain" yield aggregation. So this is also called as Interchain Yield Aggregator (IYA).

## Contents

1. **[Concepts](#concepts)**
2. **[Parameters](#network-parameters)**
3. **[Messages](#messages)**
4. **[Transactions](#transactions)**
5. **[Queries](#queries)**

## Concepts

This yield aggregator module provides an automatic earning function.
This allows users to manage their own assets according to their preferences.

### Vault

- One token many Vaults
  There can be multiple vaults for a single token. You can choose the Vault that best suits your preferences and manage your assets.

- Users can create Vaults
  Users can create Vaults without governance.
  However, it needs a fee and deposit. The fee is to prevent spam and the deposit is to provide an incentive to remove unnecessary vaults.

  The vault creator can configure the commission rate. It makes the vault creation competitive and creates an incentive for creation.

- One Vault has a combination of many strategies
  The Vault can be created by combining the strategies described below.
  You can create a Vault by selecting the strategies to be used and their weights.
  The strategy weights cannot be changed. If you want to change the weights, abolish the vault and let them go to another vault of the same token.

### Strategy

Strategies are methods of how the tokens will be managed to earn a yield. Users can add available strategies through governance with proposals.
Strategies can be developed using the **CosmWasm** smart contract.

## Network-parameters

| Field                    | Type                                                  | Label | Description                   |
| ------------------------ | ----------------------------------------------------- | ----- | ----------------------------- |
| `commission_rate`        | [cosmos.Dec](#cosmos.Dec)                             |       | Default commission rate       |
| `vault_creation_fee`     | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       | The fee to create a vault     |
| `vault_creation_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |       | The deposit to create a vault |

## Messages

### MsgDepositToVault

[MsgDepositToVault](https://github.com/UnUniFi/chain/blob/feat/iya_v2_sdk_v047/proto/yield-aggregator/tx.proto#L25-L34)

### MsgWithdrawFromVault

[MsgWithdrawFromVault](https://github.com/UnUniFi/chain/blob/feat/iya_v2_sdk_v047/proto/yield-aggregator/tx.proto#L38-L51)

### MsgCreateVault

[MsgCreateVault](https://github.com/UnUniFi/chain/blob/feat/iya_v2_sdk_v047/proto/yield-aggregator/tx.proto#L55-L78)

### MsgDeleteVault

[MsgDeleteVault](https://github.com/UnUniFi/chain/blob/feat/iya_v2_sdk_v047/proto/yield-aggregator/tx.proto#L82-L90)

### MsgTransferVaultOwnership

[MsgTransferVaultOwnership](https://github.com/UnUniFi/chain/blob/feat/iya_v2_sdk_v047/proto/yield-aggregator/tx.proto#L94-L103)

| Method Name              | Request Type                                                                     | Response Type                                                                                    | Description | HTTP Verb | Endpoint |
| ------------------------ | -------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------ | ----------- | --------- | -------- |
| `DepositToVault`         | [MsgDepositToVault](#ununifi.yield-aggregator.MsgDepositToVault)                 | [MsgDepositToVaultResponse](#ununifi.yield-aggregator.MsgDepositToVaultResponse)                 |             |           |
| `WithdrawFromVault`      | [MsgWithdrawFromVault](#ununifi.yield-aggregator.MsgWithdrawFromVault)           | [MsgWithdrawFromVaultResponse](#ununifi.yield-aggregator.MsgWithdrawFromVaultResponse)           |             |           |
| `CreateVault`            | [MsgCreateVault](#ununifi.yield-aggregator.MsgCreateVault)                       | [MsgCreateVaultResponse](#ununifi.yield-aggregator.MsgCreateVaultResponse)                       |             |           |
| `DeleteVault`            | [MsgDeleteVault](#ununifi.yield-aggregator.MsgDeleteVault)                       | [MsgDeleteVaultResponse](#ununifi.yield-aggregator.MsgDeleteVaultResponse)                       |             |           |
| `TransferVaultOwnership` | [MsgTransferVaultOwnership](#ununifi.yield-aggregator.MsgTransferVaultOwnership) | [MsgTransferVaultOwnershipResponse](#ununifi.yield-aggregator.MsgTransferVaultOwnershipResponse) |             |           |

## Transactions

### Deposit to Vault

deposit tokens to a vault.

```sh
ununifid tx yieldaggregator deposit-to-vault [id] [principal-amount] --from --chain-id
```

::: details Example

Deposit `50uguu` to the vault `#1`.

```sh
ununifid tx yieldaggregator deposit-to-vault 1 50uguu --from user --chain-id test
```

### Withdraw from Vault

withdraw tokens from a vault.

```sh
ununifid tx yieldaggregator withdraw-from-vault [id] [principal-amount] --from --chain-id
```

::: details Example

Withdraw `50uguu` from the vault `#1`.

```sh
ununifid tx yieldaggregator withdraw-from-vault 1 50uguu --from user --chain-id test
```

### Create Vault

Create a new vault.

```sh
ununifid tx yieldaggregator create-vault [denom] [commission-rate] [withdraw-reserve-rate] [fee] [deposit] [strategyWeights] --from --chain-id
```

::: details Example

Create a `GUU` vault.

- Its commission rate is `1%`.
- Its reserve rate for withdrawing is `30%`.
- Its fee is `10000uguu` & its deposit is `20000uguu`.
- It contains strategies #1:`10%` & #2:`90%`.

```sh
ununifid tx yieldaggregator create-vault uguu 0.01 0.3 10000uguu 20000uguu 1:0.1,2:0.9 --from user --chain-id test
```

### Delete Vault

Delete own vault.

```sh
ununifid tx yieldaggregator delete-vault [id] --from --chain-id
```

::: details Example

Delete the vault `#1`.

```sh
ununifid tx yieldaggregator delete-vault 1 --from user --chain-id test
```

### Transfer Vault Ownership

Transfer the own vault ownership to another address.

```sh
ununifid tx yieldaggregator transfer-vault-ownership [id] [recipient] --from --chain-id
```

::: details Example

Transfer the ownership of the vault `#1` to the address `ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w`.

```sh
ununifid tx yieldaggregator transfer-vault-ownership 1 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --from user --chain-id test
```

## Queries

### list_vault

Show all vaults

```sh
ununifid query yieldaggregator list-vault
```

```rest
ununifi/yield-aggregator/vaults
```

Response

```json
{
  "vaults": [
    {
      "id": "string",
      "denom": "string",
      "owner": "string",
      "owner_deposit": {
        "denom": "string",
        "amount": "string"
      },
      "withdraw_commission_rate": "string",
      "withdraw_reserve_rate": "string",
      "strategy_weights": [
        {
          "strategy_id": "string",
          "weight": "string"
        }
      ]
    }
  ],
  "pagination": {
    "next_key": "string",
    "total": "string"
  }
}
```

### show_vault

Show a vault

```sh
ununifid query yieldaggregator show-vault [id]
```

```rest
ununifi/yield-aggregator/vaults/{id}
```

Response

```json
{
  "vault": {
    "id": "string",
    "denom": "string",
    "owner": "string",
    "owner_deposit": {
      "denom": "string",
      "amount": "string"
    },
    "withdraw_commission_rate": "string",
    "withdraw_reserve_rate": "string",
    "strategy_weights": [
      {
        "strategy_id": "string",
        "weight": "string"
      }
    ]
  },
  "strategies": [
    {
      "denom": "string",
      "id": "string",
      "contract_address": "string",
      "name": "string"
    }
  ]
}
```

### list_strategy

Show all strategies

```sh
ununifid query yieldaggregator list-strategy [vault-denom]
```

```rest
ununifi/yield-aggregator/strategies/{denom}
```

Response

```json
{
  "strategies": [
    {
      "denom": "string",
      "id": "string",
      "contract_address": "string",
      "name": "string"
    }
  ],
  "pagination": {
    "next_key": "string",
    "total": "string"
  }
}
```

### show_strategy

Show a strategy

```sh
ununifid query yieldaggregator show-strategy [vault-denom] [id]
```

```rest
ununifi/yield-aggregator/strategies/{denom}/{id}
```

Response

```json
{
  "strategy": {
    "denom": "string",
    "id": "string",
    "contract_address": "string",
    "name": "string"
  }
}
```

### params

shows yield-aggregator params

```sh
ununifid query yieldaggregator params
```

```rest
ununifi/yield-aggregator/params
```

Response

```json
{
  "params": {
    "commission_rate": "string",
    "vault_creation_fee": {
      "denom": "string",
      "amount": "string"
    },
    "vault_creation_deposit": {
      "denom": "string",
      "amount": "string"
    }
  }
}
```
