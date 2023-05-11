package types

func (m Market) InMarketSet(marketSet []Market) bool {
	for _, market := range marketSet {
		if m == market {
			return true
		}
	}
	return false
}

func GetMarketsOutOfPerpetualFuturesGrossPositionOfMarket(grossPositionOfMarket []PerpetualFuturesGrossPositionOfMarket) []Market {
	markets := []Market{}
	for _, grossPosition := range grossPositionOfMarket {
		markets = append(markets, grossPosition.Market)
	}
	return markets
}
