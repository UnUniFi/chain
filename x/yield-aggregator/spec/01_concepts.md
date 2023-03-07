# Concepts

- One token many vaults
- One token many strategies in db
- One vault has a combination of many strategy ids (weight can't be changed via voting)
- If user want to change the weight, abolish the vault and let them go to other vault of same token (edited)
  - this make the rebalancing process very easy
- Users can create vault without governance, but it needs fee and deposit to prevent spams.
- Vault creator can configure the commission rate. It makes the vault creation competitive and creates an incentive for creation.

## Introduction  

NFT backed loan in UnUniFi does not have an automatic earning function.
Users borrowed assets in NFT backed loan seek the way to manage their assets.
This yield aggregator module will serve such an opportunity to such users.

## about Yield Aggregator (YA)

`yield-aggregator` module provides the function for yield aggregation.
Users deposit their funds to the **vault**, and then this module uses the funds to earn yield automatically.

This module is the first yield aggregator that support "interchain" yield aggregation. So this is also called as Interchain Yield Aggregator (IYA).

## Yield Farming Contract

The funds from users that is deposited to the **vault** will be allocated to the yield farming contract, in proportion to the weight that is determined in governance.
