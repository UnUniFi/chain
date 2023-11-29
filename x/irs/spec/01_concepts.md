# Concepts

## Introduction

IRS (InterestRateSwap) module is to tokenize underlying yield asset (UT) into principal tokens and yield tokens.
Principal tokens (PT) for fixed yield searchers and yield tokens (YT) for leverage yield users.

## InterestRateSwap Vault

Vault is periodically creating Tranche pools based on cycle.

## Tranche pool

Tranche pool is the place where users can mint/burn PT/YT tokens.

## Liquidity pool for PT/UT

Liquidity for PT/UT is provided by the users who would like to get swap fees by providing liquidity for PT/UT pair.
The swap fee income is stable and it has low Impermanent Loss.

PT's price will automatically fluctuate following the maturity because PT is zero-coupon bond (fixed yield).
Hence if it is vanilla AMM, it will force LPers to have impermanent loss inevitably.

## Module Params

```go
type Params struct {
	Authority    string
	TradeFeeRate sdk.Dec
}
```

## Module Queries

```proto
// Query defines the gRPC querier service.
service Query {
  // Parameters queries the parameters of the module.
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse) {
    option (google.api.http).get = "/ununifi/irs/params";
  }
  // Vaults queries the InterestRateSwapVaults
  rpc Vaults(QueryVaultsRequest) returns (QueryVaultsResponse) {
    option (google.api.http).get = "/ununifi/irs/vaults";
  }
  // Vault queries a single InterestRateSwapVault
  rpc Vault(QueryVaultRequest) returns (QueryVaultResponse) {
    option (google.api.http).get = "/ununifi/irs/vault/{strategy_contract}";
  }
  // VaultDetails queries the details of the vault
  rpc VaultDetails(QueryVaultDetailsRequest) returns (QueryVaultDetailsResponse) {
    option (google.api.http).get = "/ununifi/irs/vault/{strategy_contract}/maturities/{maturity}";
  }
  // Tranches by Strategy
  rpc Tranches(QueryTranchesRequest) returns (QueryTranchesResponse) {
    option (google.api.http).get = "/ununifi/irs/tranches/{strategy_contract}";
  }
  // Tranche by id
  rpc Tranche(QueryTrancheRequest) returns (QueryTrancheResponse) {
    option (google.api.http).get = "/ununifi/irs/tranche/{id}";
  }
}
```

## Module Messages

```proto
// Msg defines the Msg service.
service Msg {
  rpc UpdateParams(MsgUpdateParams) returns (MsgUpdateParamsResponse);
  rpc RegisterInterestRateSwapVault(MsgRegisterInterestRateSwapVault)
      returns (MsgRegisterInterestRateSwapVaultResponse);
  rpc DepositLiquidity(MsgDepositLiquidity) returns (MsgDepositLiquidityResponse);
  rpc WithdrawLiquidity(MsgWithdrawLiquidity) returns (MsgWithdrawLiquidityResponse);
  rpc DepositToTranche(MsgDepositToTranche) returns (MsgDepositToTrancheResponse);
  rpc WithdrawFromTranche(MsgWithdrawFromTranche) returns (MsgWithdrawFromTrancheResponse);
}
```
