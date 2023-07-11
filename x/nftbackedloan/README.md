# NFT BACKED LOAN

The `NFTBACKEDLOAN` module provides NFT-Fi services that allow borrowers to borrow tokens by putting NFTs as collateral.

## Contents

1. **[Concepts](#concepts)**
1. **[Liquidation](#liquidation)**
1. **[Parameters](#network-parameters)**
1. **[Messages](#messages)**
1. **[Transactions](#transactions)**
1. **[Queries and Transactions](#queries)**

## Concepts

The `x/nftbackedloan` module implements NFT-Fi services.

The borrowers list the NFT on the marketplace for collateral. The Lenders bid on the NFTs in the marketplace, and the deposit amounts at this time become the borrower's lending pool.

If the loan is not repaid by the due date or the borrower decides to sell, the NFT belongs to the lender.

### NFT Market

#### List

You can list your NFTs for auction.

#### Cancel List

You can cancel list.  
This operation is valid only when there are no bids for the NFT.

#### Bid

You can bid on NFTs to lend tokens. A deposit of at least the minimum deposit rate must be paid at the time of bidding.
To be eligible to bid, you must meet the following requirements

New bids must meet the following criteria

##### New Bid Definition

$i \in I$: index of bids
$n = |I|$: number of bids
$\{p_i\}_{i \in I}$: the price of $i$ th bid
$\{d_i\}_{i \in I}$: the deposit amount of $i$ th bid
$\{r_i\}_{i \in I}$: the interest rate of $i$ th bid
$\{t_i\}_{i \in I}$: the expiration date of $i$ th bid
$q = \frac{1}{n} \sum_{i \in I} p_i$
$s = \sum_{i \in I} d_i$: means the amount which lister can borrow with NFT as collateral
$\{a_i\}_{i \in I}$: means the amount borrowed from $i$ th bid deposit
$b = \sum_{i \in I} a_i$
$i_p(j)$: means the index of the $j$ th highest price bid
$i_d(j)$: means the index of the $j$ th highest deposit amount bid
$i_r(j)$: means the index of the $j$ th lowest interest rate bid
$i_t(j)$: means the index of the $j$ th farthest expiration date bid
$c$: minimum deposit rate

##### New bid formulation

When $(p_{\text{new}}, d_{\text{new}}, r_{\text{new}}, t_{\text{new}})$ will be added to the set of bids, the new bids sequence will be

$I' = I \cup \{n+1\}$
$n' = n + 1$
$\{p_i'\}_{i \in I'} = \{p_i\}_{i \in I} \cup \{p_{\text{new}}\}$
$\{d_i'\}_{i \in I'} = \{d_i\}_{i \in I} \cup \{d_{\text{new}}\}$
$\{r_i'\}_{i \in I'} = \{r_i\}_{i \in I} \cup \{r_{\text{new}}\}$
$\{t_i'\}_{i \in I'} = \{t_i\}_{i \in I} \cup \{t_{\text{new}}\}$
$q' = \frac{1}{n'} \sum_{i \in I'} p_i'$
$s' = \sum_{i \in I'} d_i'$

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
However, it is limited to the extent that NFT can liquidate and may not be canceled. Rebid is similarly limited.

#### Sell

lister can sell the NFT to the highest bidder.  
When the bidder accepts the offer and pays the price, the sale is completed.  
If the bidder does not pay the `bidder price - deposit price` while the sell offer is being made, the bidder's deposit will be collected by the PROTOCOL.

#### Borrow

The lister can borrow a token when there is a bid for an auction that he/she has listed.  
Failure to return borrowed tokens before the bid expires will result in liquidation.

Borrows must meet the following criteria

##### Borrow Definition

$p_i$: Bidder $i$'s bid price.

$d_i$: Bidder $i$'s deposit.

$r_i$: Bidder $i$'s interest rate (annual rate).

$r'(r, x)$: Interest rate converted to a period of $x$.

$i_p(k)$: Bidder index for the $k$th bid price.

$i_d(k)$: Bidder index for the $k$th deposit.

$i_r(k)$: Bidder index for the $k$th interest rate.

##### Borrow formulation

Borrowing capacity $a$ is defined as follows. $x$ is the borrowing period.

$$
\begin{aligned}a(x) &= \sum_{k=1}^K(x) d_{i_r(k)} \\ K(x) &= \max\left\{K \middle| \sum_{k=1}^K (1 + r'(r_{i_r(k)}, x)) d_{i_r(k)} \le \min_L \left\{ p_{i_p(l)} + \sum_{l=1}^{L-1} d_{i_p(l)} \right\} \right\} \end{aligned}
$$

In the scenario where "after several consecutive deposit forfeitures, the bidder is eventually successful in an auction, and the amount becomes the smallest," it is possible to borrow only up to the amount that can be covered, including interest, at that amount.

### Liquidation

If the borrowed tokens are not returned by the bid's expiration date, a liquidation of the auction occurs.

When liquidation occurs, the following will happen:

- All bidder must pay the liquidation amount (bid price - deposit price) by the payment deadline. If the winning bidder did not pay the liquidation amount, the deposit amount will be collected.
- The winning bidder is determined by verifying the payment in descending order of deposit amount (not bid price), and the winning bidder is established once the payment is confirmed.
- Unsuccessful bidders will have their deposits returned.
- The bidder's interest will be paid from the liquidation amount of the winning bid price.
- If all bidders did not pay the liquidation amount, the NFT and the total deposit amount become the property of the lister.

#### Auto Payment

Auto Payment is a option of bidding that automatically pays the liquidation amount when the auction's liquidation occurs for a bid. If there are insufficient funds at the time of liquidation, manual payment will be required.
The bidder can set when bidding.

## network-parameters

| Field                                     | Type                                                  | Label    | Description |
| ----------------------------------------- | ----------------------------------------------------- | -------- | ----------- |
| `min_staking_for_listing`                 | [string](#string)                                     |          |             |
| `bid_tokens`                              | [string](#string)                                     | repeated |             |
| `auto_relisting_count_if_no_bid`          | [uint64](#uint64)                                     |          |             |
| `nft_listing_delay_seconds`               | [uint64](#uint64)                                     |          |             |
| `nft_listing_period_initial`              | [uint64](#uint64)                                     |          |             |
| `nft_listing_cancel_required_seconds`     | [uint64](#uint64)                                     |          |             |
| `nft_listing_cancel_fee_percentage`       | [uint64](#uint64)                                     |          |             |
| `nft_listing_gap_time`                    | [uint64](#uint64)                                     |          |             |
| `bid_cancel_required_seconds`             | [uint64](#uint64)                                     |          |             |
| `bid_token_disburse_seconds_after_cancel` | [uint64](#uint64)                                     |          |             |
| `nft_listing_full_payment_period`         | [uint64](#uint64)                                     |          |             |
| `nft_listing_nft_delivery_period`         | [uint64](#uint64)                                     |          |             |
| `nft_creator_share_percentage`            | [uint64](#uint64)                                     |          |             |
| `market_administrator`                    | [string](#string)                                     |          |             |
| `nft_listing_commission_fee`              | [uint64](#uint64)                                     |          |             |
| `nft_listing_extend_seconds`              | [uint64](#uint64)                                     |          |             |
| `nft_listing_period_extend_fee_per_hour`  | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |          |             |

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

[MsgListNft](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L34-L49)

### MsgCancelNftListing

[MsgCancelNftListing](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L52-L55)

### MsgPlaceBid

[MsgPlaceBid](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L58-L70)

### MsgCancelBid

[MsgCancelBid](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L73-L76)

### MsgBorrow

[MsgBorrow](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L102-L106)

### MsgRepay

[MsgRepay](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L109-L113)

### MsgSellingDecision

[MsgSellingDecision](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L85-L88)

### MsgPayRemainder

[MsgPayRemainder](https://github.com/UnUniFi/chain/blob/newDevelop/proto/ununifi/nftbackedloan/tx.proto#L91-L94)

| Method Name        | Request Type                                | Response Type                                               | Description | HTTP Verb | Endpoint |
| ------------------ | ------------------------------------------- | ----------------------------------------------------------- | ----------- | --------- | -------- |
| `ListNft`          | [MsgListNft](#MsgListNft)                   | [MsgListNftResponse](#MsgListNftResponse)                   |             |           |
| `CancelNftListing` | [MsgCancelNftListing](#MsgCancelNftListing) | [MsgCancelNftListingResponse](#MsgCancelNftListingResponse) |             |           |
| `PlaceBid`         | [MsgPlaceBid](#MsgPlaceBid)                 | [MsgPlaceBidResponse](#MsgPlaceBidResponse)                 |             |           |
| `CancelBid`        | [MsgCancelBid](#MsgCancelBid)               | [MsgCancelBidResponse](#MsgCancelBidResponse)               |             |           |
| `Borrow`           | [MsgBorrow](#MsgBorrow)                     | [MsgBorrowResponse](#MsgBorrowResponse)                     |             |           |
| `Repay`            | [MsgRepay](#MsgRepay)                       | [MsgRepayResponse](#MsgRepayResponse)                       |             |           |
| `SellingDecision`  | [MsgSellingDecision](#MsgSellingDecision)   | [MsgSellingDecisionResponse](#MsgSellingDecisionResponse)   |             |           |
| `PayRemainder`     | [MsgPayRemainder](#MsgPayRemainder)         | [MsgPayRemainderResponse](#MsgPayRemainderResponse)         |             |           |

## transactions

### List Nft

Put the NFT up for auction.

```sh
ununifid tx nftbackedloan list [nft-class-id] [nft-id] --min-deposit-rate [dec] --bid-token [token] --automatic-refinancing --min-bidding-period-hours [dec]  --from --chain-id
```

::: details Example

The NFT of `a10/a10` will be listed.
The token used for bidding is `uguu` and the minimum bid deposit is `0.1%`.
Enable the automatic-refinancing function

```sh
ununifid tx nftbackedloan list a10 a10 --min-deposit-rate 0.01 --bid-token uguu --automatic-refinancing --from user --chain-id test
```

### Cancel Listing

Cancel NFT listing.

```sh
ununifid tx nftbackedloan cancel-listing [nft-class-id] [nft-id] --from --chain-id
```

::: details Example

The List of `a10/a10` will be cancel.

```sh
ununifid tx nftbackedloan cancel-listing a10 a10 --chain-id test --from user --gas=300000 -y
```

### Place bid

Bid on the NFT.

```sh
ununifid tx nftbackedloan place-bid [nft-class-id] [nft-id] [bid-amount] [deposit-amount] [deposit-interest-rate] [bidding_hour_time] --from --chain-id
```

::: details Example

Bid on `a10/a10` auction.  
The bid price is `100uguu` and `50uguu` deposit.  
The lending interest rate on the deposit is `0.1%`.  
The bid is valid for `48` hours.

```sh
ununifid tx nftbackedloan place-bid a10 a10 100uguu 50uguu 0.1 48 --from user2 --chain-id test
```

### Cancel Bidding

Cancel Bid.

```sh
ununifid tx nftbackedloan cancel-bid [nft-class-id] [nft-id] --from --chain-id
```

::: details Example

Cancel Bid on `a10/a10` auction.

```sh
ununifid tx nftbackedloan cancel-bid a10 a10 --from user2 --chain-id test
```

### Borrow tokens

Borrow tokens from the auction you have listed.

```sh
ununifid tx nftbackedloan borrow [nft-class-id] [nft-id] [bidder] [amount] --from --chain-id test
```

::: details Example

Borrow 50uguu from bidder `ununifitest` in a10/a10 auction

```sh
ununifid tx nftbackedloan borrow a10 a10 ununifitest 50uguu --from user --chain-id test
```

### Repay tokens

Returns borrowed tokens.

```sh
ununifid tx nftbackedloan repay [nft-class-id] [nft-id] [bidder] [amount] --from --chain-id test
```

::: details Example

Repay 50uguu to bidder `ununifitest` in a10/a10 auction

```sh
ununifid tx nftbackedloan repay a10 a10 ununifitest 50uguu --from user --chain-id test
```

### Decide to sell

Offer to sell NFT to the highest bidder.

```sh
ununifid tx nftbackedloan selling-decision [nft-class-id] [nft-id] --from --chain-id test
```

::: details Example

Make an offer to the highest bidder in the a10/a10 auction

```sh
ununifid tx nftbackedloan selling-decision a10 a10 --from user --chain-id test
```

### Pay full bid price

Pay the difference between the bid price and the deposit amount.

```sh
ununifid tx nftbackedloan pay-remainder [nft-class-id] [nft-id] --from --chain-id test
```

::: details Example

Pay full bid price for the a10/a10 auction

```sh
ununifid tx nftbackedloan pay-remainder a10 a10 --from user --chain-id test
```

## queries

### bidder-bids

Show User's bids

```sh
ununifid query nftbackedloan bidder-bids [bidder_address]
```

### liquidation

Show Information on liquidation of listing NFT

```sh
ununifid query nftbackedloan liquidation [class-id] [nft-id]
```

### listed-class

Show NFT class on listed

```bash

ununifid query nftbackedloan listed-class
```

### listed-nfts

Show NFT on listed

```sh
ununifid query nftbackedloan listed-nfts
```

### loan

show loans of listing NFT

```sh
ununifid query nftbackedloan loan [class-id] [nft-id]
```

### loans

show loans in all auctions

```sh
ununifid query nftbackedloan loans
```

### nft-bids

show bid in auction

```sh
ununifid query nftbackedloan nft_bids [class_id] [nft_id]
```

### nft-listing

show list info

```sh
ununifid query nftbackedloan nft-listing [class_id] [nft_id]
```

### params

shows nftbackedloan params

```sh
ununifid q nftbackedloan params
```
