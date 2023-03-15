# state

The `x/nftmarket` module keeps state of n primary objects:

## basic lithing abstract flow

```mermaid
flowchart TD
    listing --> bidding
    bidding --> pay_fee
    pay_fee --> swap_token_and_nft
    swap_token_and_nft --> end_liting
```

## late shipping lithing abstract flow

```mermaid
flowchart TD
    listing --> bidding
    bidding --> pay_fee
    pay_fee --> lister_send_things_to_win_bidder
    lister_send_things_to_win_bidder --> win_bidder_receive_things
    win_bidder_receive_things --> swap_token_and_nft
    swap_token_and_nft --> end_listing

```

# basic listing

## listing state

|No |state                |Description.                                                                                                                                         |
|---|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
|1  |unsold_state         |NFT not listed for listing                                                                                                                           |
|2  |listing_state        |It's the state of listing.                                                                                                                           |
|3  |bidding_state        |It's state that there are bids in the listing                                                                                                        |
|4  |SellingDecision_state|It's state that the lister has decided to sell                                                                                                       |
|5  |liquidation_state    |The value of the Denom used for bidding has dropped. Therefore, the collateral rate for the stabled coins issued by lister has exceeded the threshold|
|6  |end_listing_state    |It's just the state of the listing period has ended.                                                                                                 |
|7  |successful_bid_state |The lister has ended and the candidate bidder has paid for the item. The successful bidder and lister can exchange NFTs and tokens.                  |

### msg list

| ID  | Name                 |
| --- | -------------------- |
| 1   | sell Msg             |
| 2   | cancel sell Msg      |
| 3   | ------------         |
| 4   | SellingDecision Msg  |
| 5   | end listing Msg      |
| 6   | bid Msg              |
| 7   | mint stable coin Msg |
| 8   | burn stable coin Msg |
| 9   | bid cancellation Msg |
| 10  | extend Msg           |
| 11  | pay listing fee Msg  |

### listing flow

```mermaid
flowchart TD
    unsold_state -->|1.sell_msg| listing_state
　  bidding_state   -->|time out|Extend_listing_period{Extend listing period?}
　  bidding_state   -->|Exceeds available \n loan amount|liquidation{liquidation?}
    bidding_state   --->|collateral falls below \n the loan value|?{?}
　  Extend_listing_period-->|Yes_10.extend_msg|bidding_state
　  listing_state   -->|6.bid_Msg| over_minimum_bid{over minimum bid?}
    over_minimum_bid   -->|Yes| bidding_state
    over_minimum_bid   -->|No_reject_6.bit_msg| listing_state
    bidding_state   -->|6.bit_Msg| bidding_state
    bidding_state   -->|7.Mint_Msg| bidding_state
    bidding_state   -->|8.Burn_Msg| bidding_state
　  bidding_state   -->|4.SellingDecisions_Msg|　SellingDecisions_state
    SellingDecisions_state --> check_pay_fee{pay_fee?}
    check_pay_fee --yes--> successful_bid_state
    check_pay_fee --no--> Deposit_collection_process_top_bidder
    Deposit_collection_process_top_bidder --> bidding_state
　  bidding_state   -->|2.cancel_sell_msg| unsold_state
    listing_state   -->|2.cancel_sell_msg| unsold_state
　  Extend_listing_period-->|No_or_NonAction|　end_listing_state
　  liquidation -->|Yes_5.end_listing_Msg|　end_listing_state
　  liquidation -->|No_8.BurnMsg|　bidding_state
　  liquidation -->|NonAction|　penalty_process
　  penalty_process -->|5.end_listing_Msg|　end_listing_state
　  end_listing_state   -->　pay_check{pay fee?}
　  pay_check   -->|yes|　successful_bid_state
　  pay_check   -->|no|　Deposit_collection_process
　  Deposit_collection_process   -->　check_wining_bidder_candidates{any_wining_bidder_candidates?}
　  check_wining_bidder_candidates   -->|yes|　pay_check
　  check_wining_bidder_candidates   -->|no|　pay_collected_deposit_process
　  pay_collected_deposit_process   -->　unsold_state
```

### listing Token and NFT flow

#### case. start listing

```mermaid
flowchart TD
subgraph start sell nft
  lister1[[lister]]
  listing_state[[listing_state]]
  lister1--NFT--> listing_state
end
subgraph start listing
  bidder1[[bidder]]
  bidding_state1[[bidding_state]]
  bidder1 --BD--> bidding_state1
end
```

#### case. bidding

```mermaid
flowchart TD
subgraph bidding_buy_back
  bidder_buy_back[[bidder]]
  lister_buy_back[[lister]]
  NFT_author_buy_back[[NFT author]]
  UI_author_buy_back[[UI author]]

  lister_buy_back --BD--> buy_back_process
  buy_back_process --BD--> bidder_buy_back
  buy_back_process --BD--> NFT_author_buy_back
  buy_back_process --BD--> UI_author_buy_back
  buy_back_process --NFT--> lister_buy_back
end

subgraph bidding_expand_period
  top_bidder_expand_period[[top_bidder]]
  second_bidder_expand_period[[second_bidder]]
  lister_expand_period[[lister]]

  lister_expand_period--BD--> expand_period_process
  expand_period_process--BD--> top_bidder_expand_period
  expand_period_process--BD--> second_bidder_expand_period
end
```

```mermaid
flowchart TD
subgraph bidding_mint_stable_token
  lister[[lister]]

  bidding_state --BD--> cdp_process
  cdp_process--stable_token--> lister
end

subgraph bidding_prevent_liquidation
  lister2[[lister]]

  lister2 --stable_token--> liquidation_process
  liquidation_process--BD--> change_state_to_bidding_state
end
```

```mermaid
flowchart TD
subgraph bidding_accept_liquidation
  lister3[[lister]]

  lister3 --> accept_liquidation_process
  accept_liquidation_process--Locked_BD--> end_listing_process
end

subgraph bidding_ignore_liquidation
  bidder_state--BD--> force_liquidation_process
  force_liquidation_process --BD--> NFT_author
  force_liquidation_process --BD--> UI_author
  force_liquidation_process --BD--> system
  force_liquidation_process --NFT--> bidder
end
```

### bidding cancel

Bid cancellations will have a delay before the tokens are returned  
The time to return tokens for canceled bids can be determined at the time of the listing

```mermaid
flowchart TD
subgraph bidding_cancel_bid
  cancel_bidder[[cancel_bidder]]
  check_mint{lister mint stable coins?}
  check_limit{all bid deposit - cancel bid deposit < borrowed token amount?}

  check_mint.-NO.-> bidding_state
  check_limit--No--> bidding_state
  bidding_state--BD--> delay_time_process
  check_mint--Yes--> check_limit
  check_limit--Yes--> cancel_process
  cancel_process --BD--> bidding_state
  delay_time_process .-BD.-> cancel_bidder
  delay_time_process --decrease_BD_or_empty--> cancel_bidder

end
```

### end listing

```mermaid
flowchart TD

subgraph disigion
  SellingDecision_state --> check_pay_only_top_bidder{payed?}
  check_pay_only_top_bidder  --yes--> successful_bid_state1
  check_pay_only_top_bidder  --no--> collected_deposit_only_top_bidder
  collected_deposit_only_top_bidder --deposit_BD--> bidding_state
end

subgraph end listing
  lister2[[lister]]
  successful_bid_state1[successful_bid_state1]
  end_listing_state  --> check_pay{payed?}
  check_pay  --yes--> successful_bid_state1
  check_pay  --no--> collected_deposit
  collected_deposit --> check_wining_bidder{any_wining_bidder?}
  check_wining_bidder --yes--> check_pay
  check_wining_bidder --no--> return_process
  return_process --NFT_and_BD--> lister2
  return_process --> unsold_state
end

subgraph successful bid
  bidder3[[bidder]]
  lister3[[lister]]
  NFT_author[[NFT author]]
  UI_author[[UI author]]
  end_listing_state3[end_listing_state]
  end_listing_state3  --BD--> successful_bid_state
  end_listing_state3  --NFT--> successful_bid_state
  successful_bid_state  --BD_and_Locked_BD--> lister3
  successful_bid_state  --BD--> UI_author
  successful_bid_state  --BD--> NFT_author
  successful_bid_state  --NFT--> bidder3
end
```

# late shipping nft listing

TODO: Describe the detailed flow
