# Keepers

## WrappedNftKeeper

```go
type Keeper interface {
  MintWrapedNft(ctx sdk.Context, account_id string, nftId string, name string)
  BurnWrapedNft(ctx sdk.Context, nftId string)
  DepositNft(ctx sdk.Context, nftId string)
  WithdrawNft(ctx sdk.Context, account_id string, nftId string)
}
