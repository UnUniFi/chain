#!/bin/sh

# rm -rf ~/.ununifi

set -o errexit -o nounset

## Build genesis file incl account for passed address
# ununifid init --chain-id=test test
# ununifid keys mnemonic >& validator.txt
# ununifid keys mnemonic >& debug.txt
# ununifid keys add validator --recover < validator.txt
# ununifid keys add debug --recover < debug.txt
# ununifid add-genesis-account $(ununifid keys show validator --address) 100000000000000uguu,100000000000000stake
# ununifid add-genesis-account $(ununifid keys show debug --address) 100000000000000uguu,100000000000000stake
# ununifid gentx validator 100000000stake --chain-id=test
# ununifid collect-gentxs

## Start 
# ununifid start --pruning=nothing --mode=validator --minimum-gas-prices=0stake

# create class and get the class id
# ununifid tx nftmint create-class Test[ClassName] ipfs://testcid/[BaseTokenUri] 1000[TokenSupplyCap] 0[MintingPermission]  --from debug --chain-id test"
CLASS_ID=$(ununifid tx nftmint create-class Test ipfs://testcid/ 1000 0  --from debug --chain-id test -y --output json | jq -r '.events[-1].attributes[1].value')

## NOTE: $CLASS_ID returns "class_id" that returns error against below messages. 
##       Just try redefine CLASS_ID with simple text or just replace once you get the class_id.
##       If you know the solution for this, please let me know or just commit and push.
# mint nft
ununifid tx nftmint mint-nft $CLASS_ID a00 $(ununifid keys show -a debug)  --from debug --chain-id test -y

# burn nft
ununifid tx nftmint burn-nft $CLASS_ID a00 --from debug --chain-id test -y

# update token supply cap
ununifid tx nftmint update-token-supply-cap $CLASS_ID 2 --from debug --chain-id test -y                     

# update base token uri
ununifid tx nftmint update-base-token-uri $CLASS_ID  ipfs://testcid-latest/ --from debug --chain-id test -y

# send class ownership
ununifid tx nftmint send-class $CLASS_ID $(ununifid keys show -a validator)  --from debug --chain-id test -y

# queries
ununifid q nftmint class-attributes $CLASS_ID
ununifid q nftmint class-ids-by-owner $(ununifid keys show -a debug)
ununifid q nftmint class-ids-by-name Test
ununifid q nftmint nft-minter $CLASS_ID a00

ununifid q nft classes
ununifid q nft class $CLASS_ID
ununifid q nft nfts --class-id $CLASS_ID
ununifid q nft supply $CLASS_ID
