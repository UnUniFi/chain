package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validateParams(t *testing.T) {
	params := DefaultParams()

	require.NoError(t, params.Validate())
}
