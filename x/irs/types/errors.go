package types

import (
	"cosmossdk.io/errors"
)

// x/irs module sentinel errors
var (
	ErrParsingParams                = errors.Register(ModuleName, 1, "failed to marshal or unmarshal module params")
	ErrInvalidFeeDenom              = errors.Register(ModuleName, 2, "invalid fee denom")
	ErrInsufficientFee              = errors.Register(ModuleName, 3, "insufficient fee")
	ErrInvalidDepositDenom          = errors.Register(ModuleName, 4, "invalid deposit denom")
	ErrInsufficientDeposit          = errors.Register(ModuleName, 5, "insufficient deposit")
	ErrInvalidVaultId               = errors.Register(ModuleName, 6, "invalid vault id")
	ErrNotVaultOwner                = errors.Register(ModuleName, 7, "not a vault owner")
	ErrVaultHasPositiveBalance      = errors.Register(ModuleName, 8, "vault has positive balance")
	ErrInvalidCommissionRate        = errors.Register(ModuleName, 9, "invalid commission rate")
	ErrDuplicatedStrategy           = errors.Register(ModuleName, 10, "duplicated strategy")
	ErrInvalidStrategyWeightSum     = errors.Register(ModuleName, 11, "invalid strategy weight sum")
	ErrInvalidStrategyInvolved      = errors.Register(ModuleName, 12, "invalid strategy id involved")
	ErrInvalidWithdrawReserveRate   = errors.Register(ModuleName, 13, "invalid withdraw reserve rate")
	ErrInvalidAmount                = errors.Register(ModuleName, 14, "invalid amount")
	ErrStrategyNotFound             = errors.Register(ModuleName, 15, "strategy not found")
	ErrVaultNotFound                = errors.Register(ModuleName, 16, "vault not found")
	ErrInvalidVaultName             = errors.Register(ModuleName, 17, "invalid vault name")
	ErrInvalidVaultDescription      = errors.Register(ModuleName, 18, "invalid vault description")
	ErrDenomDoesNotMatchVaultSymbol = errors.Register(ModuleName, 19, "denom does not match vault symbol")
)
