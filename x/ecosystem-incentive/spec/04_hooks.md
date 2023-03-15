# Hooks

**NOTE: This is early draft.**

All rewards accumulation are executed when the according hooks function is called.   

The example hooks functions interfaces in x/nftmarket module:

```go
type NftmarketHooks interface {
	AfterNftListed(ctx sdk.Context, nftIdentifier NftIdentifier, txMemo string)
	AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier NftIdentifier, fee sdk.Coin)
	AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier NftIdentifier)
}
```

## AfterNftListed

This hook function is called for the resistration for the `ecosystem-incentive` with the `txMemo` and `nftIdentifiler`.   
To pass the `txMemo` from the memo data of `MsgListNft` requires a method to get memo data in the process of `MsgListNft` in `x/nftmarket` module.

### Location to be inserted

- `ListNft(ctx sdk.Context, msg *types.MsgListNft)` from x/nftmarket in nft_listing.go

## AfterNftPaymentWithCommission

This hook function is called for the accumulation of the reward for the subjects which are connected with the `nftIdentifiler` in the argument.
The calculation of the actual reward amount is executed in methods which this hook function calls in this module.

### Location to be inserted

- `ProcessPaymentWithCommissionFee(ctx sdk.Context, listingOwner sdk.AccAddress, denom string, amount sdk.Int)`  from x/nftmarket in nft_listing.go

## AfterNftUnlistedWituoutPayment

This hook function is called when a nft is unlisted for some reason like liquidation.   
The purpose is to remove the unlisted nft information from `NftmarketFrontendIncentiveIdTable` KVStore to keep the data consystent.

### Location to be inserted

- `CancelNftListing(ctx sdk.Context, msg *types.MsgCancelNftListing)` from x/nftmarket in nft_listing.go
- Case which bid's length for the listing is 0 in `EndNftListing(ctx sdk.Context, msg *types.MsgEndNftListing)` from x/nftmarket in nft_listing.go
