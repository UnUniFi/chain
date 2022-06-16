# Requirement

## Basic

### Mint

### Burn


### Validation

1. There's no duplicating `Class.Id` on the chain
1. There's no duplicating `NFT.Id` in a `Class`
1. TokenURI must be legal by length (idea: len(tokenURI) > 0)
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition
## Constant

There're some flexible variable that sdk's nft module has like minimum and maximum `Class.Id` string length.   
Write down those variable which we must define constantly to validate. (var name is not fixed, just momentary)   

`MinClassIdLen`   
`MinClassNameLen`   
`MaxClassNameLen`   
`MinNFTIDLen`   
`MaxNFTIDLen`   
`MaxNameLen`   
`MaxDescriptionLen`   
`MaxURILen`   
`DoNotModify`   
`IDPrefix`   
`DenomPrefix`   

## Non-transferable NFT

### Mint

### Burn

### Validation
