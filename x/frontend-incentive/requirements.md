## Basic requirement

### Incentive system

---

- Frontend-incentive reward comes from some specific transaction fee (e.g. MsgPayAuctionFee's optional fee, not tx fee)
- Those rewards have some denoms which are used in nftmarket
- Those rewards are accumulated at EndBlock
- Those rewards are determined by the `reward-rate` of the global option of this module
- Those rewards' calculation is the fee * `reward-rate`, the fee indicates some specific transaction fee (e.g. MsgPayAuctionFee's optional fee, not tx fee)
- Subjects can decide the wights and each addresses of the distribution amount of the reward in trasaction memo field (e.g. one: 0.5, two: 0.5)


### Withdrawal

1. Subject account of frontend-incentive can withdraw those accumulated rewards
1. Subject account can withdraw all accumulated rewards for all denom at once
1. Subject account can withdraw accumulated reward of the specific denom

### Query

1. Any account can make a query if there are accumulated rewards for all denom
1. Any account can make query with the specific denom if there is reward for that denom
