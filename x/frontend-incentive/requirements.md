## Basic requirement

### Incentive system

---

- Frontend-incentive reward comes from some specific transaction fee (e.g. MsgPayAuctionFee's optional fee, not tx fee)
- Those rewards have some denoms which are used in nftmarket
- Those rewards are accumulated at EndBlock
- Those rewards are determined by the `reward_rate` of the global option of this module
- Those rewards' calculation is the fee * `reward_rate`, the fee indicates some specific transaction fee (e.g. MsgPayAuctionFee's optional fee, not tx fee)
- Subjects can decide the wights and each addresses of the distribution amount of the reward in trasaction memo field (e.g. one: 0.5, two: 0.5)
- 
## Register

1. Subjects first register the addresses, weights of the proportion of the rewards and `registering_name` to be used to identify the subject to take rewards (e.g. ununifi1a~, 0.5, ununifi1b~, 0.5 registering_name)
1. Weights mush be tottaly added to 1.0
1. The `registering_name` must be unique for each

## Accumulate reward

1. If the `registering_name` is put into the specified field (currently considering memo field) for the target messages like MsgPayAuctionFee, the `reward_rate` of the consumed trading fee which is made in that transaction (not tx fee) is accumulated to the subjects
1. The reward accumulation is exucuted at EndBlock
1. The consumed trading fee is calculated from the target message's argument
1. The reward is stored as just data in the `frontend-incentive` module

### Withdrawal

1. Subject account of frontend-incentive can withdraw those accumulated rewards
1. Subject account can withdraw all accumulated rewards for all denom at once
1. Subject account can withdraw accumulated reward of the specific denom

### Distribution

1. The reward is distributed when eligible subject sends withdrawal message
1. The reward comes from the module account that accumulates the consmed trading fee (currently it'll be valut module or distribution module)


### Query

1. The accumulated rewards of any account for all denom can be queried
1. The accumulated reward of any account for the specific denom can be queried
1. The weghts for each subjects in the `registering_name` can be queried
1. 
