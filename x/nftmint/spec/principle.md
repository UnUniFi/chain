# Principle

Since SDK's nft module has very large blank space, a lot of principle rules have to be made before implementing the actual x/nftmint module outside of it.   

### Id string generation logic

- `Class.Id`
- `NFT.Id`

### Name and Symbol and URI rule

- `Class.Name`
- `Class.Symbol`
- `NFT.Name`
- `Class.Uri`
- `NFT.Uri`
- `Class.Description`

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
