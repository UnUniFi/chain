#!/bin/sh

# rm -rf ~/.ununifi

set -o errexit -o nounset
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../setup/variables.sh

## Build genesis file incl account for passed address
# ununifid init --chain-id=test test
# ununifid keys mnemonic >& validator.txt
# ununifid keys mnemonic >& debug.txt
# ununifid keys add validator --recover < validator.txt
# ununifid keys add debug --recover < debug.txt
# ununifid genesis add-genesis-account $(ununifid keys show validator --address) 100000000000000uguu,100000000000000stake
# ununifid genesis add-genesis-account $(ununifid keys show debug --address) 100000000000000uguu,100000000000000stake
# ununifid genesis gentx validator 100000000stake --chain-id=test
# ununifid genesis collect-gentxs

## Start 
# ununifid start --pruning=nothing --mode=validator --minimum-gas-prices=0stake

# create class and get the class id
# ununifid tx nftmint create-class Test[ClassName] ipfs://testcid/[BaseTokenUri] 1000[TokenSupplyCap] 0[MintingPermission]  --from debug --chain-id test"
# onlyOwner can create nft
$BINARY tx nftmint create-class Test ipfs://testcid/ 100000 0 --from $USER1 $conf 
# everyone can create nft
$BINARY tx nftmint create-class Test ipfs://testcid/ 100000 1 --from $USER1 $conf 
CLASS_ID_ONLYOWNER=$(ununifid q nftmint class-ids-by-owner $USER_ADDRESS_1 --output json |jq .owning_class_id_list.class_id[0] |sed -e 's/\"//g')
CLASS_ID_EVERYONE=$(ununifid q nftmint class-ids-by-owner $USER_ADDRESS_1 --output json |jq .owning_class_id_list.class_id[1] |sed -e 's/\"//g')
## NOTE: $CLASS_ID returns "class_id" that returns error against below messages. 
##       Just try redefine CLASS_ID with simple text or just replace once you get the class_id.
##       If you know the solution for this, please let me know or just commit and push.
# mint nft
$BINARY tx nftmint mint-nft $CLASS_ID_ONLYOWNER a00 $USER_ADDRESS_1  --from $USER1 --chain-id test -y
$BINARY tx nftmint mint-nft $CLASS_ID_EVERYONE a00 $USER_ADDRESS_1  --from $USER1 --chain-id test -y

# # burn nft
# $BINARY tx nftmint burn-nft $CLASS_ID a00 --from debug --chain-id test -y

# # update token supply cap
# $BINARY tx nftmint update-token-supply-cap $CLASS_ID 2 --from debug --chain-id test -y                     

# # update base token uri
# $BINARY tx nftmint update-base-token-uri $CLASS_ID  ipfs://testcid-latest/ --from debug --chain-id test -y

# # send class ownership
# $BINARY tx nftmint send-class $CLASS_ID $(ununifid keys show -a validator)  --from debug --chain-id test -y

# # queries
# $BINARY q nftmint class-attributes $CLASS_ID
# $BINARY q nftmint class-ids-by-owner $(ununifid keys show -a debug)
# $BINARY q nftmint class-ids-by-name Test
# $BINARY q nftmint nft-minter $CLASS_ID a00

# $BINARY q nft classes
# $BINARY q nft class $CLASS_ID
# $BINARY q nft nfts --class-id $CLASS_ID
# $BINARY q nft supply $CLASS_ID
