# Requirement

**NOTES: This is just for development. Once the other legitimate pages are finalized, remove this.**

# Basic

**_The requirements for collective NFT mainly._**

## What this feature mainly provide and doesn't

### Possible

- Anyone can create Collective NFT by creating `Class` and its belonging `NFT`s while the owner of `Class` controls the minting permission. (Collective NFT means the NFT project like CryptoPunks.)
- The NFT standard generally follows ERC721 by default.
- The content of `NFT.Uri` (metadata strucure) follows the other major market place standards.
- The owner of `NFT` can Burn its `NFT`.
- The owner of `Class` can update `Class` data with each specific message.

- The `Class.Id` can be queried by `Class.Name` if the full name is matched. 
- ( If the module which performs the data transion between cosmos SDK's x/nft module and wasmd module is implemented, the `NFT` can be extended by using CosmWasm. )
- ( For the sale, listing and selling can be done in x/nftmarket mudule. )

### Impossible

- The NFT data field is limited. So the addition of data field by the creators is impossible.
- The flexible `Class.Id` is not supported.
- The flexible `NFT.Id` is not supported.
- The flexible `NFT.Uri` is not supported.
- The addition of function to `NFT` behavior by the creators and third-party developers is impossible by the features on the chain.
- This module doesn't support simple NFT sale with fixed price at the moment.

## Creating Class

1. Anyone can create `Class` to mint NFT.
1. The owner of `Class` is recorded with `Class.Id`.
1. The initial owner of `Class` is the creator of `Class`.
1. The owner of `Class` can transfer the ownership to any recipient by sending a message.
1. The `BaseTokenUri` must be registered when to be created `Class`.
1. The `TotalSupplyCap` must be registered when to be created `Class`.
1. The `MintingPermission` must be determined when to be created `Class`.

## Class and Relating Attributes

#### Class.Id

1. `Class.Id` is generated automatically in a protocol to be unique. (the way has not determined yet.)

#### BaseTokenUri

1. `BaseTokenUri` is recorded with `Class.Id` in this module.

#### TokenSupplyCap

1. `TokenSupplyCap` is recorded with `Class.Id` in this module.
1. The `TotalSupply` of `NFT`s in a `Class` can't be exceeded over `TokenSupplyCap`.

#### MintingPermission

1. The `MintingPermission` is recored with `Class.Id` in boolean type as a flag option.
1. The default value is `True`.
1. If the `MintingPermission` is `True`, only the owner of `Class` can mint `NFT`s under that `Class`.
1. If the `MintingPermission` is `False`, anyone can mint `NFT` under that `Class`.

## NFT and Relating Attributes

#### NFT.Id

1. `NFT.Id` is the number counted from 1 by one automatically.
1. `NFT.Id` can use SDK's module storing token supply data.
1. `NFT.Id` mustn't exceed the `TokenSupplyCap`.

#### NFT.Uri

1. `NFT.Uri` represents the content identifier of the `NFT`. 
1. `NFT.Uri` is generated automatically, following this rule: `NFT.Uri` = `BaseTokenUri` + `NFT.Id`.
1. Each `NFT.Uri` can be updated at once by the owner of `Class` sending `MsgUpdateBaseTokenUri`.

## Mint

1. The owner of `Class` can choose the permission to mint `NFT` when to create `Class`.
1. If minted, the total supply of `NFT`s in `Class` increases in SDk's x/nft module.
1. The original minter address of `NFT` should be recorded with `Class.Id` and `NFT.Id`.

## Burn

1. The NFTs can be burned by the owner of `NFT`.
1. The total supply of `NFT`s in `Class` deceases.

As we handle NFTs on NFTFi protocol on UnUniFi, the burn permission belongs only the owner of the `NFT`.

## Update

We handle updating methods by managing messages for high demand elements individually.   
e.g. To update `Class.Name`, we create UpdateClassName-like message.

#### Class

1. The owner of `Class` can update data relating `Class` with each specific message.

#### TokenSupplyCap

1. `TokenSupplyCap` can be updated by the owner of that `Class`.

#### BaseTokenUri

1. `BaseTokenUri` can be updated by the owner of that `Class`.
1. `MsgUpdateBaseTokenUri` performs the change of all belonging `NFT.Uri` at once.

#### NFT

1. There's no things to be updated for `NFT`.

## Transfer

1. The owner of `NFT` can transfer that `NFT` to any recipient.
1. (In the future,) the owner of `NFT` may be able to transfer that `NFT` onto the other blockchain using IBC as long as the `NFT` follows sdk's nft module standard.

## Validation

1. There's no duplicating `Class.Id` on the chain.
1. There's no duplicating `NFT.Id` in a `Class`.
1. TokenURI must be legal by length (idea: len(tokenURI) > 0).
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition.
1. `TokenSupplyCap` must be within the limitation of `NFT.Id` restriction by SDK's x/nft module. (and overflow)

There're many other rules to be set. Please refer to these:   
https://github.com/irisnet/irismod/blob/master/modules/nft/types/validation.go   
https://github.com/cosmos/cosmos-sdk/blob/v0.46.0-rc1/x/nft/validation.go   

## Message

NOTE: There're many optional data fields. So need to decide whether we set them as the flags or as arguments.   
There's more close info [here](https://github.com/UnUniFi/chain/blob/design/spec/x/nftmint/spec/03_messages.md).

- CreateClass(class_id, ..)
- SendClass(class_id, recipient)
- UpdateBaseTokenUri(class_id, ...)
- UpdateTokenSupplyCap(class_id, ...)
- MintNFT(class_id, nft_id, ...)
- BurnNFT(class_id, nft_id)

From cosmosd SDK's x/nft module:

- Send(class_id, nft_id, sender, receiver)

## Query

The cosmos SDK's nft module has many queries.   
Please refer them since I don't write duplicated queries here. ([Query service methods](https://github.com/cosmos/cosmos-sdk/blob/aba9bdc24cb6a7b9a85e6cad617f7b55d6dcdcec/docs/architecture/adr-043-nft-module.md?plain=1#L175))

1. The owner address of `Class` can be queried by `Class.Id`.
1. The minter address of `NFT` can be queried by `Class.Id` and `NFT.Id`.
1. The `update_status_level` of `Class` can be queried by `Class.Id`.
1. The `update_status_level` of `NFT` can be queried by `Class.Id`.
1. The list of `Class.Id` can be queried by `Class.Name`. **NOTE: return object has not been decided yet.** Possible choises are just `Class.Id`.

#### Query Services

1. ClassOwner(class_id) owner
1. NFTMinter(class_id, nft_id) minter
1. ClassByName(class_name) []class_id
1. ClassBaseTokenUri() base_token_uri
1. ClassTokenSupplyCap() token_supply_cap

From cosmos SDk's x/nft module:

- Balance(class_id, owner) amount
- Owner(class_id, nft_id) owner
- Supply(class_id) amount
- NFTs(class_id, owner, pagination) nfts, pagination
- NFT(class_id, nft_id) nft
- Class(class_id) class
- Classes(pagination) classes, pagination

## Constant

**NOTE: This sectino might be removed.**

There're some flexible variables that sdk's nft module has like minimum and maximum `Class.Id` string length.   
Write down those variable which we must define constantly to validate. (var name is not fixed, just momentary)   

`MinClassIdLen`   
`MinClassNameLen`   
`MaxClassNameLen`   
`MinNFTIDLen`   
`MaxNFTIDLen`   
`MaxNameLen`   
`MaxDescriptionLen`   
`MaxURILen`    

Or, some of those could be solved by fixing and generating automatically in a protocol.   
The way to handle those elements have not been determined yet.

# Non-transferable NFT (ntNFT)

**Maybe this must be performed by Smartcontract in CosmWasm to be easy**

The requirements for Non-transferable NFT minting.

## Logic Concept

Current one idea is to create non-transferable NFT with same flow as normal NFTs,but store those NFTs in our `nftmint` module’s KVStore, not to be transfered by calling native nft module’s message.   
This can be achieved by putting `Transferable` key in `ClassAttributes` and only separate minting method from the normal minting method to store that `NFT` in our `nftmint` module.   
In this way, we can use same message outside of module implementation only if we set the conditions in each method that message calls.   

**The requirements for ntNFT have many common in normal NFT's. So I don't write obvious ones.**

## Creating Class

Same as normal NFT's.

## Burn

Same as normal NFT's.

## Transfer

#### Class

1. The owner of `Class` can transfer the ownership to any recipient.

#### ntNFT

1. Any `ntNFT` can't be transferred.
1. This can be achieved by not implementing transfer method to our `nftmint` module.

## Update

Same as normal NFT's.

## Validation

1. Same validation can be applied because many designs can be brought from normal `NFT`'s.

## Query

Almost same as normal's.

1. The `ClassAttributes.Transferable` can be queried by the `Class.Id`.
