# Concepts
Stable tokens can be minted with NFT as collateral

# requirement

## basic

## auction system
1. sellers can list "bid hook"
1. bid hook affects bid deposit amount, bid cancellation amount, collateral amount, and bidder determination logic  
  see 10_Generalized_auction_deposit.md
### selling 
1. You can list the NFTs you own on the market.
1. xxxx
1. You must stake a GUU in order to list your NFTs on the market.  (global_option)
1. you can choose which token to use for bidding.  (global_option)
1. the NFT to be listed on the market is locked
1. the auction will end after a certain amount of time
1. if no bids are received, the item will automatically be re-listed up to x times (global_option)
1. can display authenticated NFTs
1. there will be a Delay of X seconds to be listed NFT to be listed.(global_option)
1. commission fee will be charged from the Nth listing (global_option)
1. can list a minimum bid at the time of listing (global_option)
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### selling cancel
1. if no one has bid on the item, the seller can cancel the listing
1. if a bid has been submitted, it is not possible to cancel the listing
1. an item can only be cancelled after N seconds have elapsed from the time it was placed on the market (global_option)
1. the NFT to be listed will be unlocked and returned to the seller when the listing is cancelled
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### buy back
1. you can buy back the NFTs you have listed
1. in the case of a buy-back, the bid * (100 + n)% must be paid (global_option)
1. if the seller buys back the item, the highest bidder can get a commission
1. if the seller buys back the item, the 2. the second and lower bidders are not entitled to a commission
1. all locked bid tokens will be refunded
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### expand auction period
1. the listee can pay X tokens to extend the duration of the auction (global_option)
1. the commission paid by the seller will be divided equally by 50% to the first and second place bidders
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### bid
1. you can bid on the NFTs on the market
1. tokens to be bid on must meet BT criteria
1. you cannot bid unless you exceed the minimum bid
1. bidding with "p" amount of tokens will deposit "d" amount(Calculation Formula: $d=\frac{1}{bid hook}\times p$)
1. if you are the highest bidder and you want to make a higher bid, bid again
1. a high bid will result in a tick-bid from the second highest price(like ebay)
1. if the bidder has N hours remaining in the auction when the bidding takes place, the auction time will automatically be extended by n' minutes.  (global_option)
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### bid cancel
1. you may cancel your bid
1. if you are the only bidder yourself, you cannot cancel
1. free bid cancellation if bid rank is below bid hook
1. if the bid rank is bid hook or higher, the bid can be cancelled by paying the deposit amount
1. bids can only be cancelled X days after bidding (global_option)
1. tokens will be disbursed X days after the bid cancellation is approved (global_option)
1. a seller whose bid is cancelled may experience a liquidation
1. the bidder and the bid canceller's Sign must match, otherwise the bid will not be accepted and a log will be kept.

### end auction
1. sellers can end the auction
1. deposits below the bid hook will be returned at the end of the auction
1. at the end of the auction, the bidder with the bid hook position or higher will be considered as a wining bidder candidate
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### pay auction fee
1. the wining bidder candidates must pay the bid amount minus the deposit amount by N time (global_option)
1. after N hours, the protocol checks whether the wining bidder candidates  have paid their bids, starting with the highest bidder
1. the deposit amount ofthe wining bidder candidates who has not paid at the time of confirmation will be collected
1. upon confirmation of payment by the wining bidder candidates, it shall be the successful bidder
1. the deposit amount of the wining bidder candidates below the successful bidder will be returned
1. the winning bid price will be the amount of the deposit collected + the amount of the bid
1. if all wining bidder candidates do not pay, the amount of the collected deposit plus NFT to be listed will be given to the seller
1. When an auction is successful, tokens are handed over to the seller and NFTs are handed over to the successful bidder.
1. delivery of the NFT will be made X days after the successful bid (global_option)
1. the Token will be delivered X days after the successful bid.  (global_option)
1. price information must be recorded and pulled up in query
1. the price information shall be recorded as NFT, seller, successful bidder, successful bid price, successful bid date and time, successful bid type, and the number of bids cumulatively.

### bidding reward
1. if the BT is a Direct Borrowed Asset type, the bidder will receive bGUUs up to (GUUs staked by the bidder x 2 or N) (global_option)
1. you must return the bGUU when you win a bid or cancel a bid to receive the NFT
1. you must return the bGUU at the time of bid cancellation to receive the token

### borrow
1. in the case of a direct borrow asset type auction, the seller can borrow 50% of the bid tokens directly from the PROTOCOL
1. the seller can return the borrowed tokens to the protocol
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### CDP
1. in the case of a synthetic asset creation type auctions, the seller may issue stable tokens for up to 50% of the bid amount, but not more than
1. the seller can return the issued stave tokens to the protocol
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### liquidation
1. stave tokens issued by CDP and borrowed tokens are considered TemporaryToken
1. if TemporaryToken exceeds 50% of the bid amount due to bid cancellations or a drop in bid token value, it must be liquidated or TemporaryToken returned to less than 50%.Failing that, the NFT will be given to the bidder.
1. the seller must return the TemporaryToken before the liquidation or decide to win the bid. Failure to do so will result in penalties.  (global_option)
1. the penalty is that no tokens will be paid to the seller. Other than that, the process will be the same as a successful bid.
1. tx will not be accepted except for the Seller's Sign. And keep a log.

### BT(Bidding Token)
1. BT is the token standard used for auctions
1. There are two types of BTs, and they change depending on the auction type and whether the seller issues stable tokens.
1. for direct borrowing asset type auctions, BT is only for tokens specified by the seller (global_option)
1. for synthetic asset creation type auctions, BT can use any token supported by UnUniFi (global_option)


## expanded

### incentive system
1. when a seller repurchases an NFT, the NFT author, the external front-end engineer, and the ecosystem pool will each receive an N% incentive (global_option)
1. upon winning the NFT bid, the NFT author, the external front-end engineer, and the ecosystem pool will each receive an N% incentive (global_option)

### valut system
1. the NFTs you submitted are stored in Valut
1. tokens used for bidding are stored in Valut
1. when CDPing NFT as collateral, do so via vault
1. the way to pass the NFT to the successful bidder is via vault
1. the way to give tokens to sellers after the auction is over is via vault
1. valut can be closed by X in the event of a security incident (global_option)

### saving system
1. vote allows you to regain tokens stolen from an attacker
1. vote allows you to regain NFTs stolen from an attacker
1. x can close the IBC channel (global_option)
1. X may cancel the use of a particular bid token (global_option)
1. X can stop the auction function (global_option)


※global_option is defined by genesisjson 