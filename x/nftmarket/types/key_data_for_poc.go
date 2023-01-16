package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// key data for poc2 by nft_id and started time
	KeyPrefixKeyDataForPoC2 = "key_data_for_poc2"
)

func KeyDataForPoC2Key(nftIdBytes []byte, started_time time.Time) []byte {
	return append(append([]byte(KeyPrefixKeyDataForPoC2), nftIdBytes...), []byte(started_time.String())...)
}

// this is for the first creation moment
func NewKeyDataForPoC2(
	nftId NftIdentifier,
	startedAt time.Time,
) KeyDataForPoC2 {
	return KeyDataForPoC2{
		NftId:         nftId,
		StartedAt:     startedAt,
		TotalBidCount: 0,
	}
}

// calcualte the duration in seconds from startedAt to endAt and return KeyDataForPoC2
// with the added duration value
func (m KeyDataForPoC2) CalculateDuration() KeyDataForPoC2 {
	// calculate duration from startedAt to endAt
	m.ListedDurationInSeconds = int64(m.EndAt.Sub(m.StartedAt).Seconds())
	return m
}

func (m KeyDataForPoC2) UpdateMaxBorrowableAmount(currentMaxBorrowableAmount sdk.Coin) KeyDataForPoC2 {
	if m.MaxBorrowableAmountInListingPeriod.IsLT(currentMaxBorrowableAmount) {
		m.MaxBorrowableAmountInListingPeriod = currentMaxBorrowableAmount
	}
	return m
}

// update totalBorrowedAmount by adding new bid amount and return KeyDataForPoC2
func (m KeyDataForPoC2) UpdateTotalBorrowedAmount(newBidAmount sdk.Coin) KeyDataForPoC2 {
	m.TotalBorrowedAmount = m.TotalBorrowedAmount.Add(newBidAmount)
	return m
}

// update totalBidCount by adding 1 and return KeyDataForPoC2
func (m KeyDataForPoC2) TotalBidCountUp() KeyDataForPoC2 {
	m.TotalBidCount++
	return m
}
