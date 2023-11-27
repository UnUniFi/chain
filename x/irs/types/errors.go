package types

import (
	"cosmossdk.io/errors"
)

// x/irs module sentinel errors
var (
	ErrParsingParams           = errors.Register(ModuleName, 1, "failed to marshal or unmarshal module params")
	ErrInvalidVaultName        = errors.Register(ModuleName, 2, "invalid vault name")
	ErrInvalidVaultDescription = errors.Register(ModuleName, 3, "invalid vault description")
	ErrVaultNotMatured         = errors.Register(ModuleName, 4, "the vault is not matured")
	ErrInvalidPtDenom          = errors.Register(ModuleName, 5, "invalid pt denom")
	ErrInvalidYtDenom          = errors.Register(ModuleName, 6, "invalid yt denom")
	ErrInSufficientTokenInMaxs = errors.Register(ModuleName, 7, "insufficient max tokens amount")
	ErrInvalidTotalShares      = errors.Register(ModuleName, 8, "invalid total shares amount")
	ErrInsufficientExitCoins   = errors.Register(ModuleName, 9, "insufficient exit coins amount")
	ErrTrancheNotFound         = errors.Register(ModuleName, 10, "tranche not found")
	ErrDenomNotFoundInPool     = errors.Register(ModuleName, 11, "denom not found on the pool")
	ErrInvalidMathApprox       = errors.Register(ModuleName, 12, "invalid math approximation")
	ErrLimitMinAmount          = errors.Register(ModuleName, 13, "calculated amount is lower than min amount")
)
