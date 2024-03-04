package types

import (
	"cosmossdk.io/errors"
)

// x/irs module sentinel errors
var (
	ErrParsingParams           = errors.Register(ModuleName, 1, "failed to marshal or unmarshal module params")
	ErrInvalidVaultName        = errors.Register(ModuleName, 2, "invalid vault name")
	ErrInvalidVaultDescription = errors.Register(ModuleName, 3, "invalid vault description")
	ErrTrancheNotMatured       = errors.Register(ModuleName, 4, "the vault is not matured")
	ErrTrancheAlreadyMatured   = errors.Register(ModuleName, 5, "tranche already matured")
	ErrInvalidPtDenom          = errors.Register(ModuleName, 6, "invalid pt denom")
	ErrInvalidYtDenom          = errors.Register(ModuleName, 7, "invalid yt denom")
	ErrInSufficientTokenInMaxs = errors.Register(ModuleName, 8, "insufficient max tokens amount")
	ErrInvalidTotalShares      = errors.Register(ModuleName, 9, "invalid total shares amount")
	ErrInsufficientExitCoins   = errors.Register(ModuleName, 10, "insufficient exit coins amount")
	ErrTrancheNotFound         = errors.Register(ModuleName, 11, "tranche not found")
	ErrDenomNotFoundInPool     = errors.Register(ModuleName, 12, "denom not found on the pool")
	ErrInvalidMathApprox       = errors.Register(ModuleName, 13, "invalid math approximation")
	ErrLimitMinAmount          = errors.Register(ModuleName, 14, "calculated amount is lower than min amount")
	ErrZeroAmount              = errors.Register(ModuleName, 15, "zero amount")
	ErrSupplyNotFound          = errors.Register(ModuleName, 16, "supply not found")
	ErrInvalidTrancheType      = errors.Register(ModuleName, 17, "invalid tranche type")
	ErrInvalidTrancheStartTime = errors.Register(ModuleName, 18, "invalid tranche start time")
	ErrInsufficientFunds       = errors.Register(ModuleName, 19, "insufficient funds")
	ErrInvalidDepositDenom     = errors.Register(ModuleName, 20, "invalid deposit denom")
	ErrVaultNotFound           = errors.Register(ModuleName, 22, "vault not found")
	ErrInvalidAmount           = errors.Register(ModuleName, 23, "invalid amount")
	ErrNoPtDenomExists         = errors.Register(ModuleName, 24, "no pt denom exists")
	ErrNoUtDenomExists         = errors.Register(ModuleName, 25, "no ut denom exists")
	ErrInvalidPoolAssets       = errors.Register(ModuleName, 26, "invalid pool assets")
	ErrZeroDepositRate         = errors.Register(ModuleName, 27, "zero deposit rate")
	ErrNoDepositDenomExists    = errors.Register(ModuleName, 28, "no deposit denom exists")
	ErrNoDepositRequired       = errors.Register(ModuleName, 29, "no deposit required")
)
