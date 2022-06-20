# Principle

Since SDK's nft module has very large blank space, a lot of principle rules have to be made before implementing the actual x/nftmint module outside of it.   

### Id string generation logic

Please refer to the other pages under x/nftmint/spec to know what below words mean.

- `Class.Id`

As described in term_list, `Class.Id` is the unique identifier for each collective NFT. 
So we have to decide how we make it keep unique.   
Options:

1. User enter it first. When the operation starts, if that `Class.Id` is conflicted with the other `Class`'s Id in the chain, revert and stop operation.
1. The protocol generate `Class.Id` somehow and check the duplication and return `Class.Id` to let message sender know what it is.

### Id, Name, Symbol and URI rule

Please refer to the other pages under x/nftmint/spec to know what below words mean.

- `Class.Id` - the unique identifier of collective NFT
- `NFT.Id` - the unique identifier in collective NFT
- `Class.Name` - the collective name
- `Class.Symbol` - the collective NFT's ticker
- `NFT.Name` - the specific NFT name
- `Class.Uri` - the meta content location of collective NFT
- `NFT.Uri` - the metadata location of each NFT
- `Class.Description` - the description of collective NFT

### Content of `NFT.Uri` (metadsata structure)

The standard of `NFT.Uri`(metadata structure) on UnUniFi.   
The thing we have to pay attention is the compotibility to the NFTs on the other network, especially in big market size like Ethereum, solana.   
The possible elements which we should apply from the OpenSea document:   

- image - This is the URL to the image of the item.
- external_url - This is the URL that will appear below the asset's image on OpenSea and will allow users to leave OpenSea and view the item on your site.
- description - A human readable description of the item. 
- name - Name of the item.
- attribute - These are the attributes for the item, which will show up on the OpenSea page for the item. 
- background_color - Background color of the item on OpenSea. Must be a six-character hexadecimal without a pre-pended #.
- animation_url - A URL to a multi-media attachment for the item. The file extensions GLTF, GLB, WEBM, MP4, M4V, OGV, and OGG are supported, along with the audio-only extensions MP3, WAV, and OGA.
- youtube_url - A URL to a YouTube video.

Possible reference: https://docs.opensea.io/docs/metadata-standards
