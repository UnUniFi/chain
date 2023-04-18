# yield-aggregator

The `yield-aggregator` module provides the function for yield aggregation.
Users deposit their funds to the **vault**, and then this module uses the funds to earn yield automatically.

This module is the first yield aggregator that supports "interchain" yield aggregation. So this is also called as Interchain Yield Aggregator (IYA).

## Contents

1. **[Concepts](#concepts)**
2. **[Parameters](#network-parameters)**
3. **[Messages](#messages)**
4. **[Transactions](#transactions)**

## Concepts

NFT backed loan in UnUniFi does not have an automatic earning function.
Users borrowed assets in NFT backed loan seek the way to manage their assets.
This yield aggregator module will serve such an opportunity to such users.

### Vault

- One token many Vaults
  There can be multiple vaults for a single token. You can choose the Vault that best suits your preferences and manage your assets.

- Users can create Vaults
  Users can create Vaults without governance, but it needs a fee and deposit to prevent spam.
  Vault creator can configure the commission rate. It makes the vault creation competitive and creates an incentive for creation.

- One Vault has a combination of many strategies
  The Vault can be created by combining the strategies described below.
  You can create a Vault by selecting the strategies to be used and their weights.
  The strategy weights cannot be changed. If you want to change the weights, abolish the vault and let them go to another vault of the same token.

### Strategy

The Strategy is a method of how the tokens will be managed to earn a yield.
Developers can add available strategies through governance with proposals.

## Network-parameters

## Messages

### MsgDepositToVault

[MsgDepositToVault]()

### MsgWithdrawFromVault

[MsgWithdrawFromVault]()

### MsgCreateVault

[MsgCreateVault]()

### MsgDeleteVault

[MsgDeleteVault]()

### MsgTransferVaultOwnership

[MsgTransferVaultOwnership]()

| Method Name              | Request Type                                                                     | Response Type                                                                                    | Description | HTTP Verb | Endpoint |
| ------------------------ | -------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------ | ----------- | --------- | -------- |
| `DepositToVault`         | [MsgDepositToVault](#ununifi.yield-aggregator.MsgDepositToVault)                 | [MsgDepositToVaultResponse](#ununifi.yield-aggregator.MsgDepositToVaultResponse)                 |             |           |
| `WithdrawFromVault`      | [MsgWithdrawFromVault](#ununifi.yield-aggregator.MsgWithdrawFromVault)           | [MsgWithdrawFromVaultResponse](#ununifi.yield-aggregator.MsgWithdrawFromVaultResponse)           |             |           |
| `CreateVault`            | [MsgCreateVault](#ununifi.yield-aggregator.MsgCreateVault)                       | [MsgCreateVaultResponse](#ununifi.yield-aggregator.MsgCreateVaultResponse)                       |             |           |
| `DeleteVault`            | [MsgDeleteVault](#ununifi.yield-aggregator.MsgDeleteVault)                       | [MsgDeleteVaultResponse](#ununifi.yield-aggregator.MsgDeleteVaultResponse)                       |             |           |
| `TransferVaultOwnership` | [MsgTransferVaultOwnership](#ununifi.yield-aggregator.MsgTransferVaultOwnership) | [MsgTransferVaultOwnershipResponse](#ununifi.yield-aggregator.MsgTransferVaultOwnershipResponse) |             |           |

## Transactions
