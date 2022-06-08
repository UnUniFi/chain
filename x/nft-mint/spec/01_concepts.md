# Concepts

The `x/nft-mint` module implements the feature to mint NFTs on UnUniFi.

(This is very early draft. Just writing certification part as following the old planded nft-certification module)    

      
          
---     
### About Status
There is certification-status data field in class(not decided yet) data type.
The status have two stage, which are `none` and `certificated`.

**None**   
The `none` status is put when the NFTs are minted on UnUniFi.

**certificated**   
The `certificated` status is put when the specific governance proposal is passed.

## Logic   
   
`governance-certificated` requires to post the specific governance proposal. If the proposal passes, the collective NFT gets that status on class data type.


# requirement

## basic
()

## certification

### Add status

- Governance-certificated status can be brought by making the specific governance proposal   
    

### Remove status

- Governance-certificated status can be removed by making the specific governance proposal

### Query

- Callable to check cetification-status from any NFT