# NFTMARKET

The ``NFTMARKET`` module provides the logic to create and interact with auction on the UnUniFi for NFT liquidity

## Contents

1. **[Concepts](#concepts)**
2. **[Liquidation](#Liquidation)**
4. **[Parameters](#network-parameters)**
5. **[Messages](#messages)**
6. **[Transactions](#transactions)**
7. **[Queries and Transactions](#queries-and-transactions)**

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
<!-- todo: write bidding formula -->

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

todo: write spec
## messages

todo: write spec

## transactions

todo: write spec

## queries-and-transactions

todo: write spec