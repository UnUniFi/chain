# Concepts

The model of Perpetual Futures feature of this module totally follows GMX perpetual futures model.   
(ref: https://gmxio.gitbook.io/gmx/)

Briefly saying, the tradings on the perpetual futures market are supported by a unique multi-asset pool that earns liquidity prodicers fees from market making, swap fees and levrage trading.   
Because pool asset will take the counter part of the arbitral trade by a user, users can trade with high leverage and no slippage to the price from oracle with low fee.

## Pool

### Liquidity Provider Token

Our liquidity provider token's ticker is `DLP`.   
In the backend, it has `udlp` denom, which is the micro unit of `DLP`.   

DLP consists of an index of assets used for leverage trading. It can be minted using the assets which the protocol supports and burnt to redeem any index asset. The price for minting and redemption is calculated based on the formulas in the WhitePaper in the section "3.1.1".   
WhitePaper: https://ununifi.io/assets/download/UnUniFi-Whitepaper.pdf

Fees earned on the platform are directly added to the pool.　Therefore, DLP holders can benefit from them as a reward through the increasement of the DLP price.   

There's dynamic change of the minting and redemption fee rate at this moment. There's the static rate which is defined in the protocol. And, the actual fee rate also consider the difference of asset proportion between target and actual proportion.  The static base fee rate can be modified through the governace voting.

One thing to be noted is that the Liquidity Pool will take the counterpart position
of a trader’s order, so, if traders get profit, the pool and at the same time liquidity providers get loss.

## Perpetual Futures

### Position

User can open a perpetual futures position by sending `MsgOpenPosition` tx. There's no fee for opening a position.   
The position can be covered by two types of asset as margin, which are the tokens of the trading pair. If you trade 'BTC/USDC' pair, you can deposit BTC or USDC as margin. The profit will be distributed in the same token as the margin if there's some.    
The created position cannot be modified except for closing a whole in the current implementation.    
And, the liquidation is triggered against each position. The margin of the position cannot be added afterward now. But, this will be supported in the near future.   
The max leverage rate is defined in the params  of the protocol for all trading pairs equially. This can be modified through the governance voting.

When a position is created, the corresponding amount of token in the pool will be allocated to the position to assure the distribution of the profit for the position. (This could be considered as lending)    
There's no fee as the default settings for borrowing at this moment. But, it can be modified through the governance voting.

### Liquidation

A position can be liquidated if the losses and fees reduces the margin to the point where:    
`remaining_margin / (position_size / leverage) <= MarginMaintenanceRate`
MaintenanceMarginRate is defined as a parameter in the protocol. The default value of it is `0.5`.   
The values are all based on `QuoteTicker` of the protocol, which the default value is `USD`.

This is achieved through any user seding a `MsgReportLiquidation` tx. The reporter gets the fee based on the remaining margin and ReportLiquidationRewardRate by the protocol. And the remaining amount of token will be sent back the position owner.   
There's no penalty for reporting the position that is not needed to be liquidated.

### Imaginary Funding Rates

To mitigate the effect of the feature of our perpetual futures model which the liquidity provider takes the counterpart of the trader, Imaginary Funding rate exists.   
If the net position of traders lean to one side, the imaginary funding rate work to make the net position of traders neutral. The neutral net position of traders means the neutral position of the pool an at the same time liquidity providers. In the perspective of economics, it can be expressed that this model unifies the conventional funding rate and the time cost of waiting for matchmaking to the imaginary funding rate.

Imaginary Funding are levied at every 8 hours by a reporter who send the `MsgReportLevyPeriod`. The reporter gets the reward based on the imaginary funding and ReportLevyPeriodRewardRate.

## Price Feed

Prices for the accepting token are provided through our pricefeed module.   
Our pricefeed module takes the price data from the restricted addresses which are defined in the protocol in advance.   
The token price is calculated like below:
  The price of BTC in the pair of USDC,   
  `price_BTC = price_BTC_USD / price_USDC_USD`   
So, the pricefeed module has the data of BTC price based on USD and USDC price based on USD in this case.   
Price is calculated the meadian price of major exchanges for each token. Price will ideally be updated at every block. USDC or other stablecoins are not hard-coded in the protocol.

Price data is treated in this form:

```go
type CurrentPrice struct {
	  MarketId string                                 `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty" yaml:"market_id"`
	  Price    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
}
```
