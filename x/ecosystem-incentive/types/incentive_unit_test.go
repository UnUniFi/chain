package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

func TestAddIncentiveUnitid(t *testing.T) {
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	incentiveUnitIdsByAddr := types.NewIncentiveUnitIdsByAddr(addr.Bytes(), "test")
	addingId := "added_id"
	newIncentiveUnitIdsByAddr := incentiveUnitIdsByAddr.AddIncentiveUnitId(addingId)
	require.Equal(t, 2, len(newIncentiveUnitIdsByAddr))
	require.Contains(t, newIncentiveUnitIdsByAddr, addingId)
}

func TestCreateOrUpdate(t *testing.T) {
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// in case for creating new one
	incentiveUnitIdsByAddr := types.IncentiveUnitIdsByAddr{}
	incentiveUnitIdsByAddr = incentiveUnitIdsByAddr.CreateOrUpdate(addr.Bytes(), "test")
	require.Equal(t, 1, len(incentiveUnitIdsByAddr.IncentiveUnitIds))

	// in case for the update by adding new id
	addingId := "added_id"
	incentiveUnitIdsByAddr = incentiveUnitIdsByAddr.CreateOrUpdate(addr.Bytes(), addingId)
	require.Equal(t, 2, len(incentiveUnitIdsByAddr.IncentiveUnitIds))
	require.Contains(t, incentiveUnitIdsByAddr.IncentiveUnitIds, addingId)
}
