# Hook

**NOTE: This is early draft.**

All rewards accumulation are executed when the according hook function is called.   

The hook functions interfaces in x/nftmarket module:

```go
type NftmarketHooks interface {
   AfterNftListed(ctx sdk.Context, nft_id types.NftIdentifier, incentive_id string)
   AfterNftPaid(ctx sdk.Context, nft_id types.NftIdentifier, fee_amount mathInt, fee_denom string)
}
```

## AfterNftListed

This hook function is called for the resistration for the `nftmarket-incentive` with the `incentive_id` and `NftIdentifiler` if the `incentive_id` is already registered on `nftmarket-incentive` module by sending `MsgRegister` message.   

## AfterNftPaid

This hook function is called for the accumulation of the reward for the subjects which are connected with the `nft_id` in the argument.
The calculation of the actual reward amount is executed in methods which this hook function calls in this module.
