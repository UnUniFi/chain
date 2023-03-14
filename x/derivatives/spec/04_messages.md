# Messages　And Queries

## Messages

### MintLiquidityProviderToken

[MintLiquidityProviderToken](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L22-L32)

Mint liquidity provider token `DLP` by providing the acceptable asset into the pool.    
The token's price is determined by the worth of all tokens within the pool and factoring in the profits and losses of all currently opened positions.    
Hence, the `DLP` amount of being minted is determined at the time of minting.   
Fee is charged based on the defined static param.

### BurnLiquidityProviderToken

[BurnLiquidityProviderToken](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L36-L50)

Burn liquidity provider token `DLP` to the arbitrary acceptable token.    
Fee is charged based on the defined static param.

### OpenPosition

[OpenPosition](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L54-L72)

Open a perpetual futures position.    
User defines the trading pair, long/short, position size and levarage rate.   
The maximum position size is limited by the amount of the corresponding token in the pool. User cannot take a position that is larger than the pool size.

### ClosePosition

[ClosePosition](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L76-L85)

Close a whole position by defining a unique position id.   
Only the owner of the position can close it. If the position has profit, the profit will be distributed in the same token of the position margin. The fee is taken at this time. The fee rate is the defined static number in the params.

### ReportLiquidation

[ReportLiquidation](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L89-L103)

This Msg reports a position that needs to be liquidated.   
This architecture make the chain avoidable to be aware of liquidation logic in EndBlock handler to enhance the scalability.

### ReportLevyPeriod

[ReportLevyPeriod](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L107-L121)

Report a position that needs to be levied for [imaginary funding rate](todo). The reporter gets the reward based on the fee rate of the levy period report reward rate in the params.

## Queries

### Params

### Pool

[Pool](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L88-L106)

### LiquidityProviderTokenRealAPY

[LiquidityProviderTokenRealAPY](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L108-L122)

### LiquidityProviderTokenNominalAPY

[LiquidityProviderTokenNominalAPY](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L124-L138)

### PerpetualFutures

[PerpetualFutures](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L140-L162)

### PerpetualFuturesMarket

[PerpetualFuturesMarket](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L164-L197)

### AllPositions

[AllPositions](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L214-L224)

### Position

[Position](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L226-L249)

### PerpetualFuturesPositionSize

[PerpetualFuturesPositionSize](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L251-L265)

### AddressPositions

[AddressPositions](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L267-L280)

### DLPTokenRates

[DLPTokenRates](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L283-L292)

### EstimateDLPTokenAmount

[EstimateDLPTokenAmount](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L294-L312)

### EstimateRedeemAmount

[EstimateRedeemAmount](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L314-L332)
