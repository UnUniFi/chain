package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/auction/types"
)

// RegisterInvariants registers all staking invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {

	ir.RegisterRoute(types.ModuleName, "module-account",
		ModuleAccountInvariants(k))
	ir.RegisterRoute(types.ModuleName, "valid-auctions",
		ValidAuctionInvariant(k))
	ir.RegisterRoute(types.ModuleName, "valid-index",
		ValidIndexInvariant(k))
}

// ModuleAccountInvariants checks that the module account's coins matches those stored in auctions
func ModuleAccountInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {

		totalAuctionCoins := sdk.NewCoins()
		k.IterateAuctions(ctx, func(auction types.Auction) bool {
			a, ok := auction.(types.GenesisAuction)
			if !ok {
				panic("stored auction type does not fulfill GenesisAuction interface")
			}
			totalAuctionCoins = totalAuctionCoins.Add(a.GetModuleAccountCoins()...)
			return false
		})

		moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
		moduleAccCoins := k.bankKeeper.GetAllBalances(ctx, moduleAcc.GetAddress())
		broken := !moduleAccCoins.IsEqual(totalAuctionCoins)

		invariantMessage := sdk.FormatInvariant(
			types.ModuleName,
			"module account",
			fmt.Sprintf(
				"\texpected ModuleAccount coins: %s\n"+
					"\tactual ModuleAccount coins:   %s\n",
				totalAuctionCoins, moduleAccCoins),
		)
		return invariantMessage, broken
	}
}

// ValidAuctionInvariant verifies that all auctions in the store are independently valid
func ValidAuctionInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var validationErr error
		var invalidAuction types.Auction
		k.IterateAuctions(ctx, func(auction types.Auction) bool {
			a, ok := auction.(types.GenesisAuction)
			if !ok {
				panic("stored auction type does not fulfill GenesisAuction interface")
			}

			if err := a.Validate(); err != nil {
				validationErr = err
				invalidAuction = a
				return true
			}
			return false
		})

		broken := validationErr != nil
		invariantMessage := sdk.FormatInvariant(
			types.ModuleName,
			"valid auctions",
			fmt.Sprintf(
				"\tfound invalid auction, reason: %s\n"+
					"\tauction:\n\t%s\n",
				validationErr, invalidAuction),
		)
		return invariantMessage, broken
	}
}

// ValidIndexInvariant checks that all auctions in the store are also in the index and vice versa.
func ValidIndexInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		/* Method:
		- check all the auction IDs in the index have a corresponding auction in the store
		- index is now valid but there could be extra auction in the store
		- check for these extra auctions by checking num items in the store equals that of index (store keys are always unique)
		- doesn't check the IDs in the auction structs match the IDs in the keys
		*/

		// Check all auction IDs in the index are in the auction store
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextAuctionIDKey))

		indexIterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionByTimeKey))
		defer indexIterator.Close()

		var indexLength int
		for ; indexIterator.Valid(); indexIterator.Next() {
			indexLength++

			idBytes := indexIterator.Value()
			auctionBytes := store.Get(idBytes)
			if auctionBytes == nil {
				invariantMessage := sdk.FormatInvariant(
					types.ModuleName,
					"valid index",
					fmt.Sprintf("\tauction with ID '%d' found in index but not in store", types.Uint64FromBytes(idBytes)))
				return invariantMessage, true
			}
		}

		// Check length of auction store matches the length of the index
		storeIterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextAuctionIDKey))
		defer storeIterator.Close()
		var storeLength int
		for ; storeIterator.Valid(); storeIterator.Next() {
			storeLength++
		}

		if storeLength != indexLength {
			invariantMessage := sdk.FormatInvariant(
				types.ModuleName,
				"valid index",
				fmt.Sprintf("\tmismatched number of items in auction store (%d) and index (%d)", storeLength, indexLength))
			return invariantMessage, true
		}

		return "", false
	}
}
