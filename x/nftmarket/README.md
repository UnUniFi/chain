# NFTMARKET

The ``NFTMARKET`` module provides the logic to create and interact with auction on the UnUniFi for NFT liquidity

## Contents

1. **[Concepts](#concepts)**
1. **[Liquidation](#liquidation)**
1. **[Parameters](#network-parameters)**
1. **[Messages](#messages)**
1. **[Transactions](#transactions)**
1. **[Queries and Transactions](#queries)**

## Concepts

The `x/nftmarket` module implements an NFT Auction

Here we will explain basic NFT Auction concepts.

### Auction

#### List

You can list your NFTs for auction.

#### Cancel List

You can cancel list.  
The requirement is that the lister does not borrow any tokens on this listing.

#### Bid

You can bid on the auction.  
To be eligible to bid, you must meet the following requirements

New bids must meet the following criteria

##### Definition

- $i \in I$: index of bids
- $n = |I|$: number of bids
- $\{p_i\}_{i \in I}$: the price of $i$ th bid
- $\{d_i\}_{i \in I}$: the deposit amount of $i$ th bid
- $\{r_i\}_{i \in I}$: the interest rate of $i$ th bid
- $\{t_i\}_{i \in I}$: the expiration date of $i$ th bid
- $q = \frac{1}{n} \sum_{i \in I} p_i$
- $s = \sum_{i \in I} d_i$: means the amount which lister can borrow with NFT as collateral
- $\{a_i\}_{i \in I}$: means the amount borrowed from $i$ th bid deposit
- $b = \sum_{i \in I} a_i$
- $i_p(j)$: means the index of the $j$ th highest price bid
- $i_d(j)$: means the index of the $j$ th highest deposit amount bid
- $i_r(j)$: means the index of the $j$ th lowest interest rate bid
- $i_t(j)$: means the index of the $j$ th farthest expiration date bid
- $c$: minimum deposit rate

##### New bid formulation

When $(p_{\text{new}}, d_{\text{new}}, r_{\text{new}}, t_{\text{new}})$ will be added to the set of bids, the new bids sequence will be

- $I' = I \cup \{n+1\}$
- $n' = n + 1$
- $\{p_i'\}_{i \in I'} = \{p_i\}_{i \in I} \cup \{p_{\text{new}}\}$
- $\{d_i'\}_{i \in I'} = \{d_i\}_{i \in I} \cup \{d_{\text{new}}\}$
- $\{r_i'\}_{i \in I'} = \{r_i\}_{i \in I} \cup \{r_{\text{new}}\}$
- $\{t_i'\}_{i \in I'} = \{t_i\}_{i \in I} \cup \{t_{\text{new}}\}$
- $q' = \frac{1}{n'} \sum_{i \in I'} p_i'$
- $s' = \sum_{i \in I'} d_i'$

where the prime means the next state.

The constraint of $d_{n+1}'$ is

$$
  c p_{n+1}' \le d_{n+1}' \le q' - s
$$

In easy expression, it means

- $c p_{n+1}' \le d_{n+1}'$
- $s' = s + d_{n+1}' \le q'$

where $c$ means minimum deposit rate.

#### Cancel Bid

You can cancel the bid.  
Provided that the deposit of the bid is not rented out.

#### Sell

lister can sell the NFT to the highest bidder.  
When the bidder accepts the offer and pays the price, the sale is completed.  
If the bidder does not pay the `bidder price - deposit price` while the sell offer is being made, the bidder's deposit will be collected by the PROTOCOL.  

#### Borrow

The lister can borrow a token when there is a bid for an auction that he/she has listed.  
Failure to return borrowed tokens before the bid expires will result in liquidation.  

### Liquidation

If the borrowed tokens are not returned by the bid's expiration date, a liquidation of the auction occurs.  

When liquidation occurs, the following will happen:

- The bidder must pay the liquidation amount (bid price - deposit price) by the payment deadline.
- If the bidder does not pay the liquidation amount, the deposit amount will be collected.
- The winning bidder is determined by verifying the payment in descending order of deposit amount (not bid price), and the winning bidder is established once the payment is confirmed.
- The guaranteed winning amount for the lister is only the total deposit price.
- Unsuccessful bidders will have their deposits returned.
- If interest accrues for the bidder, interest will be paid under the following conditions:
  - If deposit interest < (winning bid amount - winning bidder's deposit amount): the bidder receives the full interest amount
  - If deposit interest > (winning bid amount - winning bidder's deposit amount): the bidder receives a portion of the interest based on the proportion of all interest at the time of liquidation
- If no bidder pays the liquidation amount, the NFT and the total deposit amount become the property of the lister.

#### Auto Re-refinancing

Auto Re-refinancing is a feature that automatically refinances the borrowed deposit of a listed bid when its expiration date is reached, preventing liquidation from occurring. This function requires another bid to be available, and the amount that can be borrowed from that bid must exceed the amount to be refinanced. The criteria for selecting the deposit to refinance is that the auto refinancing is performed in order of the lowest interest rate.
The lister sets this when listing.

#### Auto Payment

Auto Payment is a feature that automatically pays the liquidation amount when the auction's liquidation occurs for a bid. If there are insufficient funds at the time of liquidation, manual payment will be required.
The bidder sets this when bidding.

## network-parameters

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_staking_for_listing` | [string](#string) |  |  |
| `bid_tokens` | [string](#string) | repeated |  |
| `auto_relisting_count_if_no_bid` | [uint64](#uint64) |  |  |
| `nft_listing_delay_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_period_initial` | [uint64](#uint64) |  |  |
| `nft_listing_cancel_required_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_cancel_fee_percentage` | [uint64](#uint64) |  |  |
| `nft_listing_gap_time` | [uint64](#uint64) |  |  |
| `bid_cancel_required_seconds` | [uint64](#uint64) |  |  |
| `bid_token_disburse_seconds_after_cancel` | [uint64](#uint64) |  |  |
| `nft_listing_full_payment_period` | [uint64](#uint64) |  |  |
| `nft_listing_nft_delivery_period` | [uint64](#uint64) |  |  |
| `nft_creator_share_percentage` | [uint64](#uint64) |  |  |
| `market_administrator` | [string](#string) |  |  |
| `nft_listing_commission_fee` | [uint64](#uint64) |  |  |
| `nft_listing_extend_seconds` | [uint64](#uint64) |  |  |
| `nft_listing_period_extend_fee_per_hour` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |

1. min_staking_for_listing - not use
1. bid_tokens - Types of tokens available for auction
1. auto_relisting_count_if_no_bid - not use
1. nft_listing_delay_seconds - not use
1. nft_listing_period_initial - not use
1. nft_listing_cancel_required_seconds - Waiting time to be able to cancel list
1. nft_listing_cancel_fee_percentage - not use
1. nft_listing_gap_time - not use
1. bid_cancel_required_seconds - Waiting time to be able to cancel a bid
1. bid_token_disburse_seconds_after_cancel - tokens will be reimbursed X days after the bid cancellation is approved
1. nft_listing_full_payment_period - Period during which the bidder pays the liquidation amount
1. nft_listing_nft_delivery_period - Waiting time between successful bid and delivery of NFT
1. nft_creator_share_percentage - not use
1. market_administrator - not use
1. nft_listing_commission_fee - Auction Fee
1. nft_listing_extend_seconds - not use
1. nft_listing_period_extend_fee_per_hourCoin - not use

## messages

### MsgListNft

[MsgListNft](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L39-L59)

### MsgCancelNftListing

[MsgCancelNftListing](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L62-L69)

### MsgBorrow

[MsgBorrow](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L131-L139)

### MsgCancelBid

[MsgCancelBid](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L91-L98)

### MsgMintNft

[MsgMintNft](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L26-L36)

### MsgPayFullBid

[MsgPayFullBid](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L121-L128)

### MsgPlaceBid

[MsgPlaceBid](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L72-L88)

### MsgRepay

[MsgRepay](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L142-L150)

### MsgSellingDecision

[MsgSellingDecision](https://github.com/UnUniFi/chain/blob/newDevelop/proto/nftmarket/tx.proto#L111-L118)

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ListNft` | [MsgListNft](#ununifi.nftmarket.MsgListNft) | [MsgListNftResponse](#ununifi.nftmarket.MsgListNftResponse) |  | |
| `CancelNftListing` | [MsgCancelNftListing](#ununifi.nftmarket.MsgCancelNftListing) | [MsgCancelNftListingResponse](#ununifi.nftmarket.MsgCancelNftListingResponse) |  | |
| `PlaceBid` | [MsgPlaceBid](#ununifi.nftmarket.MsgPlaceBid) | [MsgPlaceBidResponse](#ununifi.nftmarket.MsgPlaceBidResponse) |  | |
| `CancelBid` | [MsgCancelBid](#ununifi.nftmarket.MsgCancelBid) | [MsgCancelBidResponse](#ununifi.nftmarket.MsgCancelBidResponse) |  | |
| `SellingDecision` | [MsgSellingDecision](#ununifi.nftmarket.MsgSellingDecision) | [MsgSellingDecisionResponse](#ununifi.nftmarket.MsgSellingDecisionResponse) |  | |
| `PayFullBid` | [MsgPayFullBid](#ununifi.nftmarket.MsgPayFullBid) | [MsgPayFullBidResponse](#ununifi.nftmarket.MsgPayFullBidResponse) |  | |
| `Borrow` | [MsgBorrow](#ununifi.nftmarket.MsgBorrow) | [MsgBorrowResponse](#ununifi.nftmarket.MsgBorrowResponse) |  | |
| `Repay` | [MsgRepay](#ununifi.nftmarket.MsgRepay) | [MsgRepayResponse](#ununifi.nftmarket.MsgRepayResponse) |  | |

## transactions

### List Nft

Put the NFT up for auction.

```sh
ununifid tx nftmarket listing [nft-class-id] [nft-id] --min-minimum-deposit-rate [dec] --bid-token [token] --automatic-refinancing --min-bidding-period-hours [num]  --from --chain-id
```

::: details Example

The NFT of `a10/a10` will be listed.
The token used for bidding is `uguu` and the minimum bid deposit is `0.1%`.
Enable the automatic-refinancing function

```sh
ununifid tx nftmarket listing a10 a10 --min-minimum-deposit-rate 0.01 --bid-token uguu --automatic-refinancing --from user --chain-id test
```

### Cancel List

cancel list

```sh
ununifid tx nftmarket cancel_listing [nft-class-id] [nft-id] --from --chain-id  
```

::: details Example

The List of `a10/a10` will be cancel.

```sh
ununifid tx nftmarket cancel_listing a10 a10 --chain-id test --from user --gas=300000 -y 
```

### Bid

Bid on the auction.

```sh
ununifid tx nftmarket placebid [nft-class-id] [nft-id] [bid-amount] [deposit-amount] [deposit-interest-rate] [bidding_hour_time] --from --chain-id
```

::: details Example

Bid on `a10/a10` auction.  
The bid price is `100uguu` and `50uguu` deposit.  
The lending interest rate on the deposit is `0.1%`.  
The bid is valid for `48` hours.

```sh
ununifid tx nftmarket placebid a10 a10 100uguu 50uguu 0.1 48 --from user2 --chain-id test
```

### Cancel Bid

Cancel Bid.

```sh
ununifid tx nftmarket cancelbid [nft-class-id] [nft-id] --from --chain-id
```

::: details Example

Cancel Bid on `a10/a10` auction.  

```sh
ununifid tx nftmarket cancelbid a10 a10 --from user2 --chain-id test
```

### Borrow

Borrow tokens from the auction you have listed.

```sh
ununifid tx nftmarket borrow [nft-class-id] [nft-id] [amount] --from --chain-id test
```

::: details Example

Borrow 50uguu from a10/a10 auction

```sh
ununifid tx nftmarket borrow a10 a10 50uguu --from user --chain-id test
```

### Repay

Returns borrowed tokens.

```sh
ununifid tx nftmarket repay [nft-class-id] [nft-id] [amount] --from --chain-id test
```

::: details Example

Repay 50uguu from a10/a10 auction

```sh
ununifid tx nftmarket repay a10 a10 50uguu --from user --chain-id test
```

### Sell

Offer to sell NFT to the highest bidder.

```sh
ununifid tx nftmarket selling_decision [nft-class-id] [nft-id] --from --chain-id test
```

::: details Example

Make an offer to the highest bidder in the a10/a10 auction

```sh
ununifid tx nftmarket selling_decision a10 a10 --from user --chain-id test
```

## queries

todo: write spec

### bidder_bids

Show User's bids

```sh
ununifid query nftmarket bidder_bids [bidder_address]
```

### listed_class

Show NFT class on listed

```bash

ununifid query nftmarket listed_class
```

### listed_nfts

Show NFT on listed

```sh
ununifid query nftmarket listed_nfts
```

### loan

show loans in the auction

```sh
ununifid query nftmarket loan [class-id] [nft-id]
```

### loans

show loans in all auctions

```sh
ununifid query nftmarket loans
```

### nft_bids

show bid in auction

```sh
ununifid query nftmarket nft_bids [class_id] [nft_id]
```

### nft_listing

show list info

```sh
ununifid query nftmarket nft_listing [class_id] [nft_id]
```

### params

shows nftmarket params

```sh
ununifid q nftmarket params
```
