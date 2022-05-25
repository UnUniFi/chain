# class diagram

```mermaid
 classDiagram
      Sell <|-- collateral
      Sell <|-- Bit
      Bit <|-- collateral
      collateral <|-- stateOperation
      Bit <|-- stateOperation
      Sell <|-- stateOperation
      class Sell{
          +enum tokenType
          +enum auctionType
          +int minimumBit
          +any sellerInfo
          sell()
          sold()
          cancel()
          canCancel()
          buyBuck(Bit instance)
      }
      class Bit{
          bit()
          cancel()
          getTopBidder()
          getBidderList()
      }
      class collateral {
          mint()
          burn()
      }
      class stateOperation {
          +state
          +time
          nextState(string msgName)
      }
      class keeper{
          + collateralRateListByCollateral
          + timeList
      }
```


The collateralRateListByCollaterall is used for liquidation checks.  
timeList is used for auction closeout checks.   