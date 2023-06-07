package types_test

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

func TestAddIncentiveUnitid(t *testing.T) {
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

	incentiveUnitIdsByAddr := types.NewIncentiveUnitIdsByAddr(addr.String(), "test")
	addingId := "added_id"
	newIncentiveUnitIdsByAddr := incentiveUnitIdsByAddr.AddIncentiveUnitId(addingId)
	require.Equal(t, 2, len(newIncentiveUnitIdsByAddr))
	require.Contains(t, newIncentiveUnitIdsByAddr, addingId)
}

func TestCreateOrUpdate(t *testing.T) {
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

	// in case for creating new one
	incentiveUnitIdsByAddr := types.IncentiveUnitIdsByAddr{}
	incentiveUnitIdsByAddr = incentiveUnitIdsByAddr.CreateOrUpdate(addr.String(), "test")
	require.Equal(t, 1, len(incentiveUnitIdsByAddr.IncentiveUnitIds))

	// in case for the update by adding new id
	addingId := "added_id"
	incentiveUnitIdsByAddr = incentiveUnitIdsByAddr.CreateOrUpdate(addr.String(), addingId)
	require.Equal(t, 2, len(incentiveUnitIdsByAddr.IncentiveUnitIds))
	require.Contains(t, incentiveUnitIdsByAddr.IncentiveUnitIds, addingId)
}
