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


#### Auto Re-refinancing

#### Auto Payment
