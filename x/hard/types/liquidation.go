package types

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValuationMap holds the JPY value of various coin types
type ValuationMap struct {
	Jpy map[string]sdk.Dec
}

// NewValuationMap returns a new instance of ValuationMap
func NewValuationMap() ValuationMap {
	return ValuationMap{
		Jpy: make(map[string]sdk.Dec),
	}
}

// Get returns the JPY value for a specific denom
func (m ValuationMap) Get(denom string) sdk.Dec {
	return m.Jpy[denom]
}

// SetZero sets the JPY value for a specific denom to 0
func (m ValuationMap) SetZero(denom string) {
	m.Jpy[denom] = sdk.ZeroDec()
}

// Increment increments the JPY value of a denom
func (m ValuationMap) Increment(denom string, amount sdk.Dec) {
	_, ok := m.Jpy[denom]
	if !ok {
		m.Jpy[denom] = amount
		return
	}
	m.Jpy[denom] = m.Jpy[denom].Add(amount)
}

// Decrement decrements the JPY value of a denom
func (m ValuationMap) Decrement(denom string, amount sdk.Dec) {
	_, ok := m.Jpy[denom]
	if !ok {
		m.Jpy[denom] = amount
		return
	}
	m.Jpy[denom] = m.Jpy[denom].Sub(amount)
}

// Sum returns the total JPY value of all coins in the map
func (m ValuationMap) Sum() sdk.Dec {
	sum := sdk.ZeroDec()
	for _, v := range m.Jpy {
		sum = sum.Add(v)
	}
	return sum
}

// GetSortedKeys returns an array of the map's keys in alphabetical order
func (m ValuationMap) GetSortedKeys() []string {
	keys := make([]string, len(m.Jpy))
	i := 0
	for k := range m.Jpy {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
