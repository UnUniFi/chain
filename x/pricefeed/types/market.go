package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMarket returns a new Market
func NewMarket(id, base, quote string, oracles []string, active bool) Market {
	return Market{
		MarketId:   id,
		BaseAsset:  base,
		QuoteAsset: quote,
		Oracles:    oracles,
		Active:     active,
	}
}

// Validate performs a basic validation of the market params
func (m Market) Validate() error {
	if strings.TrimSpace(m.MarketId) == "" {
		return errors.New("market id cannot be blank")
	}
	if err := sdk.ValidateDenom(m.BaseAsset); err != nil {
		return fmt.Errorf("invalid base asset: %w", err)
	}
	if err := sdk.ValidateDenom(m.QuoteAsset); err != nil {
		return fmt.Errorf("invalid quote asset: %w", err)
	}
	seenOracles := make(map[string]bool)
	for _, oracle := range m.Oracles {
		_, err := sdk.AccAddressFromBech32(oracle)
		if err != nil {
			return fmt.Errorf("invalid oracle address: %w", err)
		}

		if seenOracles[oracle] {
			return fmt.Errorf("duplicated oracle %s", oracle)
		}
		seenOracles[oracle] = true
	}
	return nil
}

// Markets array type for oracle
type Markets []Market

// Validate checks if all the markets are valid and there are no duplicated
// entries.
func (ms Markets) Validate() error {
	seenMarkets := make(map[string]bool)
	for _, m := range ms {
		if seenMarkets[m.MarketId] {
			return fmt.Errorf("duplicated market %s", m.MarketId)
		}
		if err := m.Validate(); err != nil {
			return err
		}
		seenMarkets[m.MarketId] = true
	}
	return nil
}

// String implements fmt.Stringer
func (ms Markets) String() string {
	out := "Markets:\n"
	for _, m := range ms {
		out += fmt.Sprintf("%s\n", m.String())
	}
	return strings.TrimSpace(out)
}

// NewCurrentPrice returns an instance of CurrentPrice
func NewCurrentPrice(marketID string, price sdk.Dec) CurrentPrice {
	return CurrentPrice{MarketId: marketID, Price: price}
}

// CurrentPrices type for an array of CurrentPrice
type CurrentPrices []CurrentPrice

// NewPostedPrice returns a new PostedPrice
func NewPostedPrice(marketID string, oracle string, price sdk.Dec, expiry time.Time) PostedPrice {
	return PostedPrice{
		MarketId:      marketID,
		OracleAddress: oracle,
		Price:         price,
		Expiry:        expiry,
	}
}

// Validate performs a basic check of a PostedPrice params.
func (pp PostedPrice) Validate() error {
	if strings.TrimSpace(pp.MarketId) == "" {
		return errors.New("market id cannot be blank")
	}
	_, err := sdk.AccAddressFromBech32(pp.OracleAddress)
	if err != nil {
		return fmt.Errorf("invalid oracle address: %w", err)
	}

	if pp.Price.IsNegative() {
		return fmt.Errorf("posted price cannot be negative %s", pp.Price)
	}
	if pp.Expiry.Unix() <= 0 {
		return errors.New("expiry time cannot be zero")
	}
	return nil
}

// PostedPrices type for an array of PostedPrice
type PostedPrices []PostedPrice

// Validate checks if all the posted prices are valid and there are no duplicated
// entries.
func (pps PostedPrices) Validate() error {
	seenPrices := make(map[string]bool)
	for _, pp := range pps {
		_, err := sdk.AccAddressFromBech32(pp.OracleAddress)
		if err != nil {
			return fmt.Errorf("invalid oracle address: %w", err)
		}
		if seenPrices[pp.MarketId+pp.OracleAddress] {
			return fmt.Errorf("duplicated posted price for marked id %s and oracle address %s", pp.MarketId, pp.OracleAddress)
		}

		if err := pp.Validate(); err != nil {
			return err
		}
		seenPrices[pp.MarketId+pp.OracleAddress] = true
	}

	return nil
}

// String implements fmt.Stringer
func (ps PostedPrices) String() string {
	out := "Posted Prices:\n"
	for _, p := range ps {
		out += fmt.Sprintf("%s\n", p.String())
	}
	return strings.TrimSpace(out)
}

// SortDecs provides the interface needed to sort sdk.Dec slices
type SortDecs []sdk.Dec

func (a SortDecs) Len() int           { return len(a) }
func (a SortDecs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortDecs) Less(i, j int) bool { return a[i].LT(a[j]) }
