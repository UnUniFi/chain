# state

The `x/nftmarket` module keeps state of n primary objects:


## auction state
1. unsold_state
    NFT not listed for auction
1. selling_state
1. bidding_state
1. liquidation_state
1. end_auction_state
1. successful_bid_state

### msg list 

|ID |Name                         |
|---|-----------------------------|
|1  |sell Msg                     |
|2  |cancel sell Msg              |
|3  |buy back Msg                 |
|4  |end auction Msg              |
|5  |end auction Msg              |
|6  |bid Msg                      |
|7  |mint stable coin Msg         |
|8  |burn stable coin Msg         |
|9  |bid cancellation Msg         |
|10 |extend Msg                   |
|11 |pay auction fee Msg          |

### auction flow 
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
　  bidding_state   -->|4.end_auction_Msg|　end_auction_state
　  bidding_state   -->|3.buy_back_Msg| unsold_state
		bidding_state   -->|9.bid_cancellation_Msg| cancelling_check_bidder{other bidder?}
		cancelling_check_bidder--->|No_faield_bid_cancellation_Msg| bidding_state
		cancelling_check_bidder-->|Yes| cancelling_check_limit{over limitr?}
		cancelling_check_limit-->|Yes| liquidation
		cancelling_check_limit-->|No| bidding_state
    selling_state   -->|2.cancel_sell_msg| unsold_state
　  Extend_auction_period-->|No_or_NonAction|　end_auction_state
　  liquidation -->|Yes_5.end_auction_Msg|　end_auction_state
　  liquidation -->|No_8.BurnMsg|　bidding_state
　  liquidation -->|NonAction|　penalty_process
　  penalty_process -->|5.end_auction_Msg|　end_auction_state
　  end_auction_state   -->　pay_check{pay fee?}
　  pay_check   -->|yes|　successful_bid_state
　  pay_check   -->|no|　Deposit_collection_process
　  Deposit_collection_process   -->　check_wining_bidder_candidates{any_wining_bidder_candidates?}
　  check_wining_bidder_candidates   -->|yes|　pay_check
　  check_wining_bidder_candidates   -->|no|　pay_collected_deposit_process
　  pay_collected_deposit_process   -->　unsold_state
```

### auction Token and NFT flow 
#### case. start acution
```mermaid
flowchart TD
subgraph start sell nft
	seller1[[seller]]
	selling_state[[selling_state]]
	seller1--NFT--> selling_state
end
subgraph start auction
	bidder1[[bidder]]
	bidding_state1[[bidding_state]]
	bidder1 --BT--> bidding_state1
end
```

#### case. bidding 
```mermaid
flowchart TD
subgraph bidding_buy_back
	bidder_buy_back[[bidder]]
	seller_buy_back[[seller]]
	NFT_author_buy_back[[NFT author]]
	UI_author_buy_back[[UI author]]

	seller_buy_back --BT_And_GUU--> buy_back_process
	buy_back_process --BT_And_GUU--> bidder_buy_back
	buy_back_process --GUU--> NFT_author_buy_back
	buy_back_process --GUU--> UI_author_buy_back
	buy_back_process --NFT--> seller_buy_back
end

subgraph bidding_expand_period
	top_bidder_expand_period[[top_bidder]]
	second_bidder_expand_period[[second_bidder]]
	seller_expand_period[[seller]]

	seller_expand_period--GUU--> expand_period_process
	expand_period_process--GUU--> top_bidder_expand_period
	expand_period_process--GUU--> second_bidder_expand_period
end
```

```mermaid
flowchart TD
subgraph bidding_mint_stable_token
	seller[[seller]]

	bidding_state --BT--> cdp_process
	cdp_process--stable_token--> seller
end

subgraph bidding_prevent_liquidation
	seller2[[seller]]

	seller2 --stable_token--> liquidation_process
	liquidation_process--BT--> change_state_to_bidding_state
end
```

```mermaid
flowchart TD
subgraph bidding_accept_liquidation
	seller3[[seller]]

	seller3 --> accept_liquidation_process
	accept_liquidation_process--Locked_BT--> end_auction_process
end

subgraph bidding_ignore_liquidation
	bidder_state--BT--> force_liquidation_process 
	force_liquidation_process --GUU--> NFT_author
	force_liquidation_process --GUU--> UI_author
	force_liquidation_process --BT--> system
	force_liquidation_process --NFT--> bidder
end
```

### bidding cancel 
Bid cancellations will have a delay before the tokens are returned  
The time to return tokens for canceled bids can be determined at the time of the auction  
```mermaid
flowchart TD
subgraph bidding_cancel_bid
	cancel_bidder[[cancel_bidder]]
	check_mint{mint stable coins?}
	check_limit{over limit?}

	check_mint--NO--> bidding_state
	check_limit--No--> bidding_state
	bidding_state--BT--> delay_time_process
	check_mint--Yes--> check_limit
	check_limit--Yes--> cancel_process
	cancel_bidder--BT--> cancel_process 
	liquidation_state --BT--> delay_time_process
	delay_time_process --BT--> cancel_bidder
	cancel_process --BT--> liquidation_state
	
end
```

### end auction 
```mermaid
flowchart TD

subgraph end auction
	seller2[[seller]]
	successful_bid_state1[successful_bid_state1]
	end_auction_state  --> check_pay{payed?}
	check_pay  --yes--> successful_bid_state1
	check_pay  --no--> collected_deposit
	collected_deposit --> check_wining_bidder{any_wining_bidder?}
	check_wining_bidder --yes--> check_pay
	check_wining_bidder --no--> return_process
	return_process --NFT_and_BT--> seller2
	return_process --> unsold_state
end

subgraph successful bid
	bidder3[[bidder]]
	seller3[[seller]]
	NFT_author[[NFT author]]
	UI_author[[UI author]]
	end_auction_state3[end_auction_state]
	end_auction_state3  --BT--> successful_bid_state
	end_auction_state3  --NFT--> successful_bid_state
	successful_bid_state  --BT_and_Locked_BT--> seller3
  successful_bid_state  --BT--> UI_author
	successful_bid_state  --BT--> NFT_author
	successful_bid_state  --NFT--> bidder3
end
```
