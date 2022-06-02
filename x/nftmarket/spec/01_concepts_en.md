# Concepts
stablecoins can be minted with NFT as collateral

# requirement

## basic

## auction system
1. listers decide “bid hook”
1. bid hook affects bid deposit amount, bid cancellation amount, collateral amount, and bidder determination logic  
  see 10_Generalized_auction_deposit.md

### listing 
1. You can list the NFTs you own on the marketplace.
1. The lister can decide the bid hook from a number between 1 to 100 at the time of listing. (default:x, default is global_option)
1. the token used for bidding is determined by the lister based on BT.(global_option)
1. the NFT to be listed on the marketplace is locked
1. if no bids are received, the item will automatically be re-listed up to x times (global_option)
1. Lister can display authenticated NFTs
1. Lister can decide a minimum bid at the time of listing (global_option)
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### listing cancel
1. if no one has bid on the item, the lister can cancel the listing
1. if a bid has been placed, the lister may cancel the listing by paying a cancellation fee
1. cancellation fee is X% of the bid deposit(global_option)
1. Commission paid by the lister will be divided wining bidder candidate in proportion to their percentage of the deposit amount
1. The listing of items can only be cancelled after N seconds have elapsed from the time it was placed on the marketplace (global_option)
1. the NFT to be listed will be unlocked and returned to the lister when the listing is cancelled
1. bid deposits will be refunded
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### expand auction period
1. the lister can pay BT tokens to extend the period of the auction (global_option)
1. Commission paid by the lister will be divided wining bidder candidate in proportion to their percentage of the deposit amount
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### bid
1. you can bid on the NFTs on the marketplace
1. tokens to be bid on must meet BT criteria
1. you cannot bid unless you exceed the minimum bid
1. bidding with "p" amount of tokens will deposit "d" amount(Calculation Formula: $d=\frac{1}{bid hook}\times p$)
1. if you are the highest bidder and you want to make a higher bid, bid again
1. if the bidder has N hours remaining in the auction when the bidding takes place, the auction time will automatically be extended by n' minutes.  (global_option)
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### bid cancel
1. you can cancel your bid
1. if you are the only bidder yourself, you cannot cancel
1. Bidder can cancel bids free of charge if the bidder's bid rank is below the bid hook.
1. if the bid rank is bid hook or higher, the bid can be cancelled by paying a cancellation fee.
1. Cancellation Fee Formula: ```MAX{canceling_bidder's_deposit - (total_deposit - borrowed_lister_amount), 0}```
1. bids can only be cancelled X days after bidding (global_option)
1. tokens will be reimbursed X days after the bid cancellation is approved (global_option)
1. Liquidation may occur for sellers whose bids are cancelled.
1. the bidder and the bid canceller's Sign must match, otherwise the bid will not be accepted and a log will be kept.

### SellingDecision 
1. the lister can decide the successful bidder when there are bids
1. the winning bidder must pay the bid amount - deposit amount within N hours (global_option)
1. if the successful bidder fails to pay on time, the amount of the successful bidder's deposit will be collected
1. if the winning bidder does not pay by the due date, the auction period will be extended for X days and the auction will be restarted (global_option)
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### end auction
1. the auction will end after a certain amount of time
1. deposits below the bid hook will be returned at the end of the auction
1. at the end of the auction, the bidder with the bid hook position or higher will be considered as a wining bidder candidate
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### pay auction fee
1. the wining bidder candidates must pay the bid amount minus the deposit amount by N time (global_option)
1. after N hours, the protocol checks whether the wining bidder candidates  have paid their bids, starting with the highest bidder
1. the deposit amount of the wining bidder candidates who has not paid at the time of confirmation will be collected
1. upon confirmation of payment by the wining bidder candidates, it shall be the successful bidder
1. the deposit amount of the wining bidder candidates below the successful bidder will be returned
1. the winning bid price paid to the lister will be the amount of the ```deposit_collected + (bidder price - bidder deposit)```
1. if all wining bidder candidates do not pay, the amount of the collected deposit plus NFT to be listed will be given to the lister
1. When an auction is successful, tokens are handed over to the lister and NFTs are handed over to the successful bidder.
1. delivery of the NFT will be made X days after the successful bid (global_option)
1. the Token will be delivered X days after the successful bid.  (global_option)
1. When an auction is successful, the price information must be recorded and pulled up in query
1. the price information shall be recorded as NFT, lister, successful bidder, successful bid price, successful bid date and time, successful bid type, and the number of bids cumulatively.

### boost staking reward
1. if the BT is a Direct Borrowed Asset type, When a bid is placed, the bidder's staking GUUs are increased up to the limit of (GUUs staked by bidder x 2 or N)(global_option)
1. staking GUUs will increase for a period of time until the auction ends or you cancel your bid

### borrow
1. in the case of a direct borrow asset type auction, the lister can borrow the Total deposit amount above bid hook rank from the protocol
1. the lister can return the borrowed tokens to the protocol
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### CDP
1. in the case of a synthetic asset creation type auctions, the lister can issue stablecoins with the total deposit of bid hook rank or higher as collateral
1. the lister can return the issued stave tokens to the protocol
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### liquidation
1. if the amount of stabled tokens issued exceeds 50% of the bid amount due to a decline in the bid token value, you must liquidate or return the stabled tokens to less than 50% of the bid amount.
1.  When liquidation is occurring, if you don't return the stave token by N hours or decide to win the bid, you will be penalized. (global_option)
1. the penalty is that no tokens will be paid to the lister. Other than that, the process will be the same as a successful bid.
1. the NFT and the collection deposit will be held by the module if all wining bidder candidates not pay at the time of penalty
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### BT(Bidding Token)
1. BT is the token standard used for auctions
1. There are two types of BTs, and they change depending on the auction type and whether the lister issues stablecoins.
1. for direct borrowing asset type auctions, BT is only for tokens specified by the lister (global_option)
1. for synthetic asset creation type auctions, BT can use any token supported by UnUniFi (global_option)


## expanded

### incentive system
1. when a lister repurchases an NFT, the NFT author, the external front-end engineer, and the ecosystem pool will each receive an N% incentive (global_option)
1. upon winning the NFT bid, the NFT author, the external front-end engineer, and the ecosystem pool will each receive an N% incentive (global_option)

### vault system
1. the NFTs you submitted are stored in Vault
1. tokens used for bidding are stored in Vault
1. When you create CDP using NFT as collateral, do so via vault
1. NFT disbursement at the time of bidding is done via vault
1. Token disbursement at the time of bidding is done via vault
1. vault can be closed by X in the event of a security incident (global_option)

### saving system
1. vote allows you to regain tokens stolen from an attacker
1. vote allows you to regain NFTs stolen from an attacker
1. x can close the IBC channel (global_option)
1. X may cancel the use of a particular bid token (global_option)
1. X can stop the auction function (global_option)


※global_option is defined by genesisjson 