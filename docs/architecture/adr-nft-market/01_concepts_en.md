# Concepts

stablecoins can be minted with NFT as collateral

# requirement

## basic

## listing system

1. There are two types of listing
   normal listing
   late shipping nft listing.
1. normal listing is an listing where NFTs are passed and received over the protocol
1. late shipping nft listing are listing with delivery outside of protocol

### listing

1. You can list the NFTs you own on the marketplace.
1. the token used for bidding is determined by the lister based on BD.(global_option)
1. the NFT to be listed on the marketplace is locked
1. if no bids are received, The listing will close in X days. (global_option)
1. Lister can display authenticated NFTs
1. Lister can decide a minimum deposit rate at the time of listing (global_option)
1. Lister can choose between normal listing and late shipping nft listing at the time of listing
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### listing cancel

1. if no one has bid on the item, the lister can cancel the listing
1. if a bid has been placed, the lister may cancel the listing by paying a cancellation fee
1. cancellation fee is X% of the bid deposit(global_option)
1. Commission paid by the lister will be divided all bidder in proportion to their percentage of the deposit amount
1. The listing of items can only be cancelled after N seconds have elapsed from the time it was placed on the marketplace (global_option)
1. the NFT to be listed will be unlocked and returned to the lister when the listing is cancelled
1. bid deposits will be refunded
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### expand listing period

1. the lister can pay BD tokens to extend the period of the listing (global_option)
1. Commission paid by the lister will be divided wining bidder candidate in proportion to their percentage of the deposit amount
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### bid

1. you can bid on the NFTs on the marketplace
1. tokens to be bid on must meet BD criteria
1. you cannot bid unless you meet the bidding criteria.
   see 10_collateralized_liquidation_deposit_auction.md
1. bid deposit can be set freely as long as it meets the criteria
1. bid deposit lending rates can be freely set
1. bidding period can be freely set
1. bidders can activate the automatic payment feature at the time of bidding
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### bid cancel

1. bids cannot be cancelled at will
1. only bids received after the bidding period will be automatically and free of charge cancelled by the protocol.

### SellingDecision

1. the lister can decide the successful bidder when there are bids
1. the winning bidder must pay the bid amount - deposit amount within N hours (global_option)
1. if the successful bidder fails to pay on time, the amount of the successful bidder's deposit will be collected
1. if the winning bidder does not pay by the due date, the listing period will be extended for X days and the listing will be restarted (global_option)
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### end listing

1. the auction ends with SellingDecision or liquidation
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### pay listing fee

1. bidders must pay the bid amount minus the deposit amount by N time (global_option)
1. if automatic payment is enabled, bidder's balance will be paid automatically
1. after N hours, the protocol checks whether the bidders have paid their bids, starting with the highest bidder
1. the deposit amount of the bidders who has not paid at the time of confirmation will be collected
1. upon confirmation of payment by the bidder, it shall be the successful bidder
1. the deposit amount of the bidders below the successful bidder will be returned
1. the winning bid price paid to the lister will be the amount of the `（deposit_collected + (bidder price - bidder deposit)) * (1.00 - fee_rate)` Note: fee_rate variable name could be changed
1. if all wining bidder candidates do not pay, the amount of the collected deposit plus NFT to be listed will be given to the lister
1. When an listing is successful, tokens are handed over to the lister and NFTs are handed over to the successful bidder.
1. delivery of the NFT will be made X days after the successful bid (global_option)
1. the Token will be delivered X days after the successful bid. (global_option)
1. When an listing is successful, the price information must be recorded and pulled up in query
1. the price information shall be recorded as NFT, lister, successful bidder, successful bid price, successful bid date and time, successful bid type, and the number of bids cumulatively.

### late shipping nft

1. late shipping nft listting are delivered by the lister outside of protocol after payment by the winning bidder
1. the listing will end after the winning bidder receives the item from the lister
1. Token will be paid to the lister after the listing ends.
1. the Token will be delivered X days after the successful bid. (global_option)

### boost staking reward

1. if the BD is a Direct Borrowed Asset type, When a bid is placed, the bidder's staking GUUs are increased up to the limit of (GUUs staked by bidder x 2 or N)(global_option)
1. staking GUUs will increase for a period of time until the listing ends or you cancel your bid

### borrow

1. in the case of a direct borrow asset type listing, the lister can borrow the Total deposit amount above bid_active_rank rank from the protocol
1. the lister can return the borrowed tokens to the protocol
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### CDP

1. in the case of a synthetic asset creation type listings, the lister can issue stablecoins with the total deposit of bid_active_rank rank or higher as collateral
1. the lister can return the issued stave tokens to the protocol
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### liquidation

1. if the amount of stabled tokens issued exceeds 50% of the bid amount due to a decline in the bid token value, you must liquidate or return the stabled tokens to less than 50% of the bid amount.
1. When liquidation is occurring, if you don't return the stave token by N hours or decide to win the bid, you will be penalized. (global_option)
1. the penalty is that no tokens will be paid to the lister. Other than that, the process will be the same as a successful bid.
1. the NFT and the collection deposit will be held by the module if all wining bidder candidates not pay at the time of penalty
1. tx will not be accepted except for the Lister's Sign. And keep a log.

### BD(Bidding Token)

1. BD is the token standard used for listings
1. There are two types of BDs, and they change depending on the listing type and whether the lister issues stablecoins.
1. for direct borrowing asset type listings, BD is only for tokens specified by the lister (global_option)
1. for synthetic asset creation type listings, BD can use any token supported by UnUniFi (global_option)

### Trading Fee

1. There's fee for the trading NFT except for trasaction fee
1. Specifically, that's extracted when user actually pay listing fee to buy NFT
1. The fee rate is defined as global_option between 0-100%
1. The protocol earned fee goes to the other module account (specific module has not decieded yet)
1. The fee is contained in bidding price

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
1. X can stop the listing function (global_option)

※global_option is defined by genesisjson
