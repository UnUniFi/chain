# Concepts

`yield-aggregator` module provides the function for yield aggregation.
It find the optimal target of funds for asset management (e.g. the highest interest rate pool).

`nft-marketmaker` module and CosmWasm contracts will call the keeper of this module.

## DailyPercent

`nft-marketmaker` module and CosmWasm contracts must report annual percents and must pay reported daily percent rate (DPR) back to depositors.

Annual percent rate (APR) and annual percent yield (APY) can be calculated with recent `n` (you can choose) days' `DailyPercent` data.