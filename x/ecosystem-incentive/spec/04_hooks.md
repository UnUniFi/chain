# Hooks

**NOTE: This is early draft.**

All rewards accumulation are executed when the according hooks function is called.   

The example hooks functions interfaces in x/nftmarket module:

```go
type NftmarketHooks interface {
   AfterNftListed(ctx sdk.Context, nft_id types.NftIdentifier, incentive_id string)
   AfterNftPaid(ctx sdk.Context, nft_id types.NftIdentifier, fee_amount mathInt, fee_denom string)
   AfterNftUnlisted(ctx sdk.Context, nft_id types.NftIdentifier)
}
```

## AfterNftListed

This hooks function is called for the resistration for the `ecosystem-incentive` with the `incentive_id` and `NftIdentifiler` if the `incentive_id` is already registered on `ecosystem-incentive` module by sending `MsgRegister` message.   
To pass the `incentive_id` from the memo data of `MsgListNft` requires a method to get memo data in the process of `MsgListNft` in `x/nftmarket` module.

## AfterNftPaid

This hooks function is called for the accumulation of the reward for the subjects which are connected with the `nft_id` in the argument.
The calculation of the actual reward amount is executed in methods which this hook function calls in this module.

## AfterNftUnlisted

This hook function is called when a nft is unlisted for some reason like liquidation.   
The purpose is to remove the unlisted nft information from `IncentiveIdTable` KVStore to keep the data consystent.
