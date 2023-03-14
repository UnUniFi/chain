# Concepts

## Pool

- `QuoteTicker` is `USD` by default.
- Tickers `USD` and `USDC` should be discriminated.
- Be aware of the difference between `denom` and `ticker`. The former is `uguu` and the latter is `GUU`
- Each token has to have `DenomMetadata` to be indexed Symbol in the implementation.

## Liquidity Provider Token

Read whitepaper.

## Perpetual Futures

### Position

The created position cannot be modified except for closing a whole at this moment. And, the liquidation is occured against each position. But, the margin of the position also cannot be added afterward.

### Imaginary Funding Rates

- Imaginary Funding Rates are levied at every 8 hours.

## Price Feed

Prices for the accepting token are provided through our pricefeed module.   
Our pricefeed module takes the price data from the restricted addresses which are defined in the protocol in advance.   
The token price is calculated like below:
  The price of BTC in the pair of USDC,   
  `price_BTC = price_BTC_USD / price_USDC_USD`   
So, the pricefeed module has the data of BTC price based on USD and USDC price based on USD in this case.
