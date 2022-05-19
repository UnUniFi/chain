# state

The `x/nftmarket` module keeps state of n primary objects:

1. auction state.
1. bid balances.
1. NFT Ownership.
1. CDP balance using NFT.


## auction state
1. unsold_state
1. selling_state
1. bidding_state
1. liquidation_state
1. successful_bid_state


auction flow 
```mermaid
flowchart TD
    unsold_state -->|1.sell_msg| selling_state
　  bidding_state   -->|time out|Extend_auction_period{Extend auction period?}
　  bidding_state   -->|Exceeds available \n loan amount|liquidation{liquidation?}
		bidding_state   --->|collateral falls below \n the loan value|?{?}
　  Extend_auction_period-->|Yes_10.extend_msg|bidding_state
　  selling_state   -->|6.bid_Msg| over_minimum_bid{over minimum bid?}
　  over_minimum_bid   -->|Yes|	bidding_state
　  over_minimum_bid   -->|No_reject_6.bit_msg|	selling_state
　  bidding_state   -->|6.bit_Msg|	bidding_state
　  bidding_state   -->|7.Mint_Msg|	bidding_state
　  bidding_state   -->|8.Burn_Msg|	bidding_state
　  bidding_state   -->|4.accept_Msg|　successful_bid_state
　  bidding_state   -->|3.buy_back_Msg| unsold_state
		bidding_state   -->|9.bid_cancellation_Msg| cancelling_check_bidder{other bidder?}
		cancelling_check_bidder--->|No_faield_bid_cancellation_Msg| bidding_state
		cancelling_check_bidder-->|Yes| cancelling_check_limit{over limitr?}
		cancelling_check_limit-->|Yes| liquidation
		cancelling_check_limit-->|No| bidding_state
    selling_state   -->|2.cancel_sell_msg| unsold_state
　  Extend_auction_period-->|No_or_NonAction|　successful_bid_state
　  liquidation -->|Yes_5.end_auction_Msg|　successful_bid_state
　  liquidation -->|No_8.BurnMsg|　bidding_state
　  liquidation -->|NonAction|　penalty_process
　  penalty_process -->|5.end_auction_Msg|　successful_bid_state
```
