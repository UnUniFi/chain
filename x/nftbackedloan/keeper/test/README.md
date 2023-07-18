# NFT Backed Loan Keeper Test

Before run debug `build then launch`
When re-run same script, chain needs to be initialized.

Each `.sh` uses functions of `nft_listing.go` & test the following value.

## Selling NFT

### 01_selling_decision

Successful sale of NFT

- The lister get Bid price - fee
- The bidder lose Bid price
- The NFT Transfer

### 02_selling_decision_no_paid

Decided to sell, but the bidder did not pay.

- The Deposit forfeiture from 1st bidder
- The collect_amount add forfeited deposit
- The collect_amount is reflected in the next selling_decision
- The NFT Transfer (2nd selling_decision)

### 03_Liquidation

- The bidder lose Bid price
- The lister get (Bid price - borrow) - fee
- The NFT Transfer

### 04_Liquidation_no_paid

- The bidder lose deposit
- The Lister get (forfeited deposit - borrow)
- The NFT return to lister

### 05_liquidation_deposit_forfeiture

- The bidder(user2) lose deposit.
- The bidder(user3) pay bid price & get NFT.
- The owner get price & forfeited deposit.

### 06_liquidation_deposit_refund

- The bidder(user2) get refund.
- The bidder(user3) pay bid price & get NFT.
- The owner get price.
