package types

func (m Market) InMarketSet(marketSet []Market) bool {
	for _, market := range marketSet {
		if m == market {
			return true
		}
	}
	return false
}

func GetMarketsOutOfPerpetualFuturesNetPositionOfMarket(netPositionOfMarket []PerpetualFuturesNetPositionOfMarket) []Market {
	markets := []Market{}
	for _, netPosition := range netPositionOfMarket {
		markets = append(markets, netPosition.Market)
	}
	return markets
}
