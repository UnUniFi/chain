package keeper

// TODO: add test for
// VaultAmountInStrategies(ctx sdk.Context, vault types.Vault) sdk.Int {
// VaultUnbondingAmountInStrategies(ctx sdk.Context, vault types.Vault) sdk.Int {
// VaultWithdrawalAmount(ctx sdk.Context, vault types.Vault) sdk.Int {
// VaultAmountTotal(ctx sdk.Context, vault types.Vault) sdk.Int {
// EstimateMintAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, principalAmount sdk.Int) sdk.Coin {
// EstimateRedeemAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, lpAmount sdk.Int) sdk.Coin {
// DepositAndMintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, principalAmount sdk.Int) error {
// TODO: add test for BurnLPTokenAndRedeem
// Imagine
// withdraw reserve 100
// staked 900
// User A execute MsgWithdraw with amount 10.
// maintenance rate is max(0, 100 - 10) / (100 + 0) = 90 / 100 = 0.9
// withdraw reserve 90
// staked 900
// bonding 890
// unbonding 10
// Then User B execute MsgWithdraw with amount 10.
// maintenance rate is max(0, 90 - 10) / (90 + 10) = 80 / 100 = 0.8
// withdraw reserve 80
// staked 900
// bonding 880
// unbonding 20
// after the unbonding period of user A withdrawal, unbonded token will go to withdraw reserve. ( for simplification, don't think the rebalancing now)
// withdraw reserve 90
// staked 890
// bonding 880
// unbonding 10
// after the unbonding period of user B withdrawal, unbonded token will go to withdraw reserve. ( for simplification, don't think the rebalancing now)
// withdraw reserve 100
// staked 880
// bonding 880
// unbonding 0
// Then User C execute MsgWithdraw with amount 10.
// maintenance rate is max(0, 100 - 10) / (100 + 0) = 90 / 100 = 0.9
