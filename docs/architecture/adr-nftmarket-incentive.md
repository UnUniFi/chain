# ADR: nftmarket-incentive Module

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
The rate to determine the percentage for the `nftmarket-incentive` reward out of the total trading fee.    
This value is connected to the `incentive_type`.    
e.g.  
`frontend` - `0.8`   

`denom`   
The token ticker like GUU, BTC etc.   
In UnUniFi NFT market, some tokens other than GUU (native UnUniFi token) can be used to purchase NFTs. So the rewards are accumuleted in some denoms.

`reward_payer`   
The account to pay the reward for `nftmarket-incentive` protocol.   
It's the module account that collect the protocol earned fees from x/nftmarket module.

# Basic specs

## Incentive system

- nftmarket-incentive reward comes from the fee that is made in `nftmarket` module
- Those rewards have some denoms which are used in nftmarket
- Those rewards are accumulated at the timing hooks are called
- Those rewards are determined by the `reward_rate` of the global option of this module in `reward_setting` and the protocol earned NFT trading fee amount
- Those rewards' calculation is `the trading fee * reward_rate`
- Subjects register `incentive_id` and each addresses and its weights (`subject_weight_map`) to receive the reward by sending a message at first
- Subjects have to put registerd `incentive_id` in a target message's memo field to accumulate the frontend_reward
- Subjects can send a withdrawal message to actually receive the frontend_incentive reward
- Subjects can see how much reward is accumulated for their address
- The reward is distributed from the module account that accumulates the NFT trading fee

## Register

1. Subjects first register the addresses, weights of the proportion of the rewards and `incentive_id` to be used to identify the subject to take rewards in `incentive_store` (e.g. ununifi1a~, 0.5, ununifi1b~, 0.5 registering_name)
1. Weights mush be tottaly added to 1.0
1. The `incentive_id` must be unique for each
1. The object related to one `incentive_id` is static
1. Any address pairs or an address can resister `incentive_store`

## Accumulate reward

1. If the `incentive_id` is put into the specified field (currently considering memo field) for the target messages like MsgPayAuctionFee, the `reward_rate` of the consumed trading fee which is made in that transaction (not tx fee) is accumulated to the subjects
1. The reward accumulation is exucuted at EndBlock
1. The consumed trading fee is calculated from the target message's argument
1. The reward is stored as just data in the `nftmarket-incentive` module with the subject address as key

## Withdrawal

1. Subject account of nftmarket-incentive can withdraw those accumulated rewards
1. Subject account can withdraw all accumulated rewards for all denom at once
1. Subject account can withdraw accumulated reward of the specific denom
1. If the sender module account doesn't have enough funds for the withdrawal, the witthdrawal tx fails with emitting the specific event with the reason.

## Distribution

1. The reward is distributed when eligible subject sends withdrawal message
1. The reward comes from the module account that accumulates the consumed trading fee
1. The `reward_rate` determines the maximum rate of amount for the `nftmarket-incentive` in the consumed trading fee

### The way to achieve distribution

1. Actually sending corresponding coin for reward in a process using SendCoinFromModuleAccount 
1. (Or possibly in a process, mint corresponding coin for reward for the subject address and just subtract corresponding coin from the subject module account)

## Query

1. The accumulated rewards of any account for all denom can be queried by `address`
1. The accumulated reward of any account for the specific denom can be queried by `address` and `denom`
1. The weghts for each subjects in the `incentive_id` can be queried by `incentive_id`

## Params

`reward_rate`   
The factor to multipy the trading fee for the reward of this module.   
e.g. If `reward_rate` is 80% and the trading fee that is made in a target message is 100GUU, the actual reward for target `incentive_id` subjects is `100GUU * 0.80 = 80GUU`.   

`message_type`   
The specific message types which are subject to the `nftmarket-incentive`

## EndBlock

**This logic is not finalized. Needed to be researched.**

1. Update the reward amount if there are target transactions in a block and that has `incentive_id` in memo field
1. The reward is calculated `trading fee * reward_rate`, trading fee indicates the protocol earned fee by NFT trading
1. At this moment, what it's needed to do is just update the stored data regarding reward amount for the denom of the subjects address by number
1. Need to distinguish between success and failure of the txs before.

## Target message type

1. There's specific message type which is subject to `nftmarket-incentive`.
1. The criteria to choose the message type to be suject is the cash flow to the lister
1. Current idea is MsgPayFullBid type is what that is.
1. Possibly the message type to be subject for `nftmarket-incentive` will increase.

# Check List

**NOTE: You can remove this file once it's done.**

## Needed to be inspected

The check list to achieve all requirements of this module.

- [ ] At EndBlock, is it possible to distinguish transactions from message type?
- [ ] Is it possible to extract arguments of that message?
- [ ] Is it possible to extract memo field data of that message?
- [ ] Is possible to distinguish transaction succeeds
- [ ] Best way to get subject address and its weight via CLI (json file or map?)
- [ ] The way to contain reward information for each addresses and denoms
- [ ] Can you use the handler in the transaction tips article
