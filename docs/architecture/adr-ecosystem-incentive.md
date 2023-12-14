# ADR: ecosystem-incentive Module

## Status

PROPOSED

# Term list

**NOTE: You can maybe remove this file once state page is finalized.**

The definition of the specific terms which are related to this module.

`incentive_id`  
incentive_id is the unique identifier in the `incentive_store`. Hence, it can't be duplicated.

`weight`  
The ratio of the reward distribution in a `incentive_store` unit.  
`incentive_store` can contain several `subject`s and ratio for each.

`reward_setting`  
The setting about `reward_type` and `reward_rate`.

`incentive_type`  
The type of incentive for the subject.  
We first have `frontend`.

`reward_rate`  
The rate to determine the percentage for the `ecosystem-incentive` reward out of the total trading fee.  
This value is connected to the `incentive_type`.  
e.g.  
`frontend` - `0.8`

`denom`  
The token ticker like GUU, BTC etc.  
In UnUniFi NFT market, some tokens other than GUU (native UnUniFi token) can be used to purchase NFTs. So the rewards are accumuleted in some denoms.

`reward_payer`  
The account to pay the reward for `ecosystem-incentive` protocol.  
It's the module account that collect the protocol earned fees from x/nftmarket module.

# Basic specs

## Incentive system

- ecosystem-incentive reward comes from the fee that is made in `nftmarket` module
- Those rewards have some denoms which are used in nftmarket
- Those rewards are accumulated at the timing hooks are called
- Those rewards are determined by the `reward_rate` of the global option of this module in `reward_setting` and the protocol earned NFT trading fee amount
- Those rewards' calculation is `the trading fee * reward_rate`
- Subjects register `incentive_id` and each addresses and its weights (`subject_weight_map`) to receive the reward by sending a message at first
- Subjects have to put registerd `incentive_id` in a target message's memo field to accumulate the nftmarket_reward
- Subjects can send a withdrawal message to actually receive the nftmarket_incentive reward
- Subjects can see how much reward is accumulated for their address
- The reward is distributed from the module account that accumulates the NFT trading fee

## Service Flow

1. Subject send a `MsgRegister` message to register `incentive_id` and addresses.
1. At the `AfterNftListed` hooks function, connect that `NftIdentifier` with the `incentive_id` by passing memo filed data through an argument.
1. At the `AfterNftPaid` hooks function, record the reward which the `incentive_id` earned by passing exact `fee_amount` and `fee_denom` in the argument of that hook function.
1. At the `MsgWithdrawResward`, execute bank send method from the module account which accumulate fees internally.

## Register

1. Subjects first register the addresses, weights of the proportion of the rewards and `incentive_id` to be used to identify the subject to take rewards in `incentive_store` (e.g. ununifi1a~, 0.5, ununifi1b~, 0.5 registering_name)
1. Weights mush be totaly added to 1.0 (undetermined)
1. The `incentive_id` must be unique for each
1. The object related to one `incentive_id` is static
1. Any address pairs or an address can resister `incentive_store`

## Accumulate reward

1. If the `incentive_id` is put into the specified field (currently considering memo field) for the target messages like MsgPayAuctionFee, the `reward_rate` of the consumed trading fee which is made in that transaction (not tx fee) is accumulated to the subjects
1. The reward accumulation is exucuted at the timing hooks are called
1. The reward is stored as just data in the `ecosystem-incentive` module with the subject address as key (not for each `incentive_id`)

## Withdrawal

1. Subject account of ecosystem-incentive can withdraw those accumulated rewards
1. Subject account can withdraw all accumulated rewards for all denom at once
1. Subject account can withdraw accumulated reward of the specific denom
1. If the sender module account doesn't have enough funds for the withdrawal, the witthdrawal tx fails with emitting the specific event with the reason.

## Distribution

1. The reward is distributed when eligible subject sends withdrawal message
1. The reward comes from the module account that accumulates the consumed trading fee

### The way to achieve distribution

1. Actually sending corresponding coin for reward in a process using SendCoinFromModuleAccount
1. (Or possibly in a process, mint corresponding coin for reward for the subject address and just subtract corresponding coin from the subject module account)

## Query

1. The accumulated rewards of any account for all denom can be queried by `address`
1. The accumulated reward of any account for the specific denom can be queried by `address` and `denom`
1. The weghts for each subjects in the `incentive_id` can be queried by `incentive_id`

## Params

`reward_setting`  
This contains `reward_type` and `reward_rate` in array for the incentive configuration.  
e.g.

```protobuf
message RewardParams {
  repeated RewardRate reward_rate = 1;
}
```

`reward_rate`  
The factor to multipy the trading fee for the reward of this module.  
e.g. If `reward_rate` is 80% and the trading fee that is made in a target message is 100GUU, the actual reward for target `incentive_id` subjects is `100GUU * 0.80 = 80GUU`.

```protobuf
message RewardRate {
  string reward_type = 1;
  unsure rate = 2;
}
```

## Hooks

**This logic is not finalized. Needed to be researched.**

1. Update the reward amount when the according hooks function is called in `nftmarket` module
1. The reward amount is in the hooks function
1. At this moment, what it's needed to do is just update the stored data regarding reward amount for the denom of the subjects address by number
1. Hooks are also called for resistration of the incentive with `incentive_id` and `NftIdentifier`.
1. To pass the `incentive_id` from the memo data of `MsgListNft` requires a method to get memo data in the process of `MsgListNft` in `x/nftmarket` module.

The interfaces:

```go
type NftmarketHooks interface {
   AfterNftListed(ctx sdk.Context, nft_id types.NftId, incentive_id string)
   AfterNftPaid(ctx sdk.Context, nft_id types.NftId, fee_amount mathInt, fee_denom string)
}
```

# Check List

**NOTE: You can remove this file once it's done.**

## Needed to be inspected

The check list to achieve all requirements of this module.

- [x] At EndBlock, is it possible to distinguish transactions from message type?
- [x] Is it possible to extract arguments of that message?
- [x] Is it possible to extract memo field data of that message?
- [x] Is possible to distinguish transaction succeeds
- [x] Best way to get subject address and its weight via CLI (json file or map?)
- [x] The way to contain reward information for each addresses and denoms
- [x] Can you use the handler in the transaction tips article
