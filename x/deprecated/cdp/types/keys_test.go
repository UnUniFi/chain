package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/crypto/ed25519"
)

var addr = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

func TestKeys(t *testing.T) {
	key := CdpKeySuffix(0x01, 2)
	db, id := SplitCdpKey(key)
	require.Equal(t, int(id), 2)
	require.Equal(t, byte(0x01), db)

	denomKey := DenomIterKey(0x01)
	db = SplitDenomIterKey(denomKey)
	require.Equal(t, byte(0x01), db)

	depositKey := DepositKeySuffix(2, addr)
	id, a := SplitDepositKey(depositKey)
	require.Equal(t, 2, int(id))
	require.Equal(t, a, addr)

	depositIterKey := DepositIterKey(2)
	id = SplitDepositIterKey(depositIterKey)
	require.Equal(t, 2, int(id))

	require.Panics(t, func() { SplitDepositIterKey([]byte{0x03}) })

	collateralKey := CollateralRatioKey(0x01, 2, sdk.MustNewDecFromStr("1.50"))
	db, id, ratio := SplitCollateralRatioKey(collateralKey)
	require.Equal(t, byte(0x01), db)
	require.Equal(t, int(id), 2)
	require.Equal(t, ratio, sdk.MustNewDecFromStr("1.50"))

	bigRatio := sdk.OneDec().Quo(sdk.SmallestDec()).Mul(sdk.OneDec().Add(sdk.OneDec()))
	collateralKey = CollateralRatioKey(0x01, 2, bigRatio)
	db, id, ratio = SplitCollateralRatioKey(collateralKey)
	require.Equal(t, ratio, MaxSortableDec)

	collateralIterKey := CollateralRatioIterKey(0x01, sdk.MustNewDecFromStr("1.50"))
	db, ratio = SplitCollateralRatioIterKey(collateralIterKey)
	require.Equal(t, byte(0x01), db)
	require.Equal(t, ratio, sdk.MustNewDecFromStr("1.50"))

	require.Panics(t, func() { SplitCollateralRatioKey(badRatioKey()) })
	require.Panics(t, func() { SplitCollateralRatioIterKey(badRatioIterKey()) })

}

func badRatioKey() []byte {
	r := append(append(append(append([]byte{0x01}, sep...), []byte("nonsense")...), sep...), []byte{0xff}...)
	return r
}

func badRatioIterKey() []byte {
	r := append(append([]byte{0x01}, sep...), []byte("nonsense")...)
	return r
}
