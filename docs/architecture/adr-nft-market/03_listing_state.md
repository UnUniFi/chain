# state

The `x/nftmarket` module keeps state of n primary objects:

## basic lithing abstract flow

```mermaid
flowchart TD
    listing_state --> bidding_state
    bidding_state --> end_liting_state
    end_liting_state --> successful_bid_state
```

# basic listing

## listing state

|No |state                |Description.                                                                                                                                         |
|---|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
|1  |unsold_state         |NFT not listed for listing                                                                                                                           |
|2  |listing_state        |It's the state of listing.                                                                                                                           |
|3  |bidding_state        |It's state that there are bids in the listing                                                                                                        |
|4  |SellingDecision_state|It's state that the lister has decided to sell                                                                                                       |
|5  |end_listing_state    |The borrowing term of the DEPOSIT has been exceeded.                                                                                                 |
|6  |successful_bid_state |The lister has ended and the candidate bidder has paid for the item. The successful bidder and lister can exchange NFTs and tokens.                  |

### state change msg list

| ID  | Name                 |
| --- | -------------------- |
| 1   | listing Msg          |
| 2   | cancel list Msg      |
| 3   | bid Msg              |
| 4   | cancel bid Msg       |
| 5   | SellingDecision Msg  |
| 6   | pay listing fee Msg  |

### state flow
Yellow lines are automatically checked by protocol
```mermaid
flowchart TD

    unsold_state
    listing_state

    unsold_state -->|1.listing_msg| listing_state
    listing_state   -->|2.cancel_list_msg| unsold_state
    listing_state   -->|no bid and x days passed| unsold_state
    listing_state   -->|3.bid_Msg| bidding_state
　  bidding_state   -->|4.cancel_bid_msg| listing_state
　  bidding_state   -->|2.cancel_list_msg| unsold_state
　  bidding_state   -->|no borrow and x days passed| unsold_state

　  bidding_state   -->|borrow Deposit past deadline|liquidation{liquidation?}
　  liquidation -->|No_auto borrow|　bidding_state
　  bidding_state   -->|5.SellingDecisions_Msg|　SellingDecisions_state
    SellingDecisions_state --> check_pay_fee{pass SellingDecisions pay check?}
    check_pay_fee --yes--> successful_bid_state
    check_pay_fee --no--> bidding_state
　  liquidation -->|Yes|　end_listing_state
　  end_listing_state   -->　pay_check{pass liquidation pay check?}
　  pay_check   -->|yes|　successful_bid_state
　  pay_check   -->|no|　unsold_state


    linkStyle 2 stroke:#ff3,stroke-width:4px,color:red;
    linkStyle 6 stroke:#ff3,stroke-width:4px,color:red;
    linkStyle 7 stroke:#ff3,stroke-width:4px,color:red;
```
### paycheck flow
#### SellingDecisions pay check flow

```mermaid
flowchart TD 
    bidding_state>change bidding_state]
    successful_bid_state>change successful_bid_state]

    SellingDecisions_state --> check_pay_fee{highest bidder pay fee?}
    check_pay_fee --no--> Deposit_collection_process_top_bidder
    Deposit_collection_process_top_bidder --> bidding_state
    check_pay_fee --yes--> successful_bid_state
```

#### liquidation pay check flow
```mermaid
flowchart TD 
    back_unsold_state>change unsold_state]
    successful_bid_state>change successful_bid_state]

　  end_listing_state   -->　pay_check{current bidder pay fee?}
　  pay_check   -->|no|　Deposit_collection_process
　  Deposit_collection_process   -->　check_any_bidder{any_bidder?}
　  check_any_bidder   -->|yes|　pay_check
　  check_any_bidder   -->|no|　pay_collected_deposit_process
　  pay_collected_deposit_process   -->　back_unsold_state
　  pay_check   --yes-->　successful_bid_state
```


### listing Token and NFT flow

#### case. start listing and bidding

```mermaid
flowchart TD
subgraph start listing nft
  lister1[lister]
  lister1--NFT--> nftmarket_mod
  subgraph nftmarket_mod state when listing
    listing_state[[listing_state]]
    unsold_state[[unsold_state]]
    unsold_state--> listing_state
  end
end
subgraph start bidding
  bidder1[bidder]
  nftmarket_mod1[nftmarket_mod]
  bidder1 --deposit--> nftmarket_mod1
  subgraph nftmarket_mod state when bidding
    listing_state1[[listing_state]]
    bidding_state1[[bidding_state]]
    listing_state1--> bidding_state1
  end
end
```

#### case. bidding

```mermaid
flowchart TD
subgraph borrow_token
  lister[[lister]]
  nftmarket_mod[[nftmarket_mod]]

  nftmarket_mod --token--> lister
end

subgraph repay_token
  lister2[[lister]]
  nftmarket_mod1[[nftmarket_mod]]
  lister2 --borrow_token--> nftmarket_mod1
end
```

```mermaid
flowchart TD
subgraph liquidation_occurred
  nftmarket_mod[[nftmarket_mod]]

  nftmarket_mod --> check_deposit_deadline{borrowed deposit deadline has passed?}
  check_deposit_deadline --yes--> check_amount{check borrowed amount < borrowable amount?}
  check_deposit_deadline --no--> do_nothing
  check_amount --yes--> auto_borrow
  check_amount --no--> liquidation_pay_check
end

```

### bidding cancel

```mermaid
flowchart TD
subgraph bidding_cancel_bid
  cancel_bidder[[cancel_bidder]]
  check_borrow{lister borrow token?}
  check_your_deposit_use{cancel bidder deposit used?}

  check_borrow.-NO.-> bidding_state
  check_your_deposit_use--No--> bidding_state
  bidding_state----> delay_time_process
  check_borrow--Yes--> check_your_deposit_use
  check_your_deposit_use--Yes--> not_accept_cancel_bid
  delay_time_process --token--> cancel_bidder

end
```

### end listing

```mermaid
flowchart TD
subgraph successful bid
  bidder[[bidder]]
  lister[[lister]]
  protocol_pool[[protocol_pool]]

  end_listing_state3[end_listing_state]
  end_listing_state3  --token--> successful_bid_state
  end_listing_state3  --NFT--> successful_bid_state
  successful_bid_state  --token--> lister
  successful_bid_state  --token--> protocol_pool
  successful_bid_state  --NFT--> bidder
end
```

# late shipping nft listing

TODO: Describe the detailed flow
