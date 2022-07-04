package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/stretchr/testify/require"
)

const (
	testAddr = "ununifi1ghaguquuytpdgdhfthmtva0vdjmf4q99k2jhd2"
)

func TestCreateId(t *testing.T) {
	var seq uint64 = 0
	addr := testAddr
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	classIdSeq0 := createClassId(seq, accAddr)
	err := nfttypes.ValidateClassID(classIdSeq0)
	require.NoError(t, err)

	seq += 1
	classIdSeq1 := createClassId(seq, accAddr)
	require.NoError(t, err)
	require.NotEqual(t, classIdSeq0, classIdSeq1)
}
