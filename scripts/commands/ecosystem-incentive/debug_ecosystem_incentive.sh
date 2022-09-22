#!/bin/bash

set -uxe

CLASS_ID=$(ununifid tx nftmint create-class Test ipfs://testcid/ 1000 0  --from validator --chain-id test -y  -o json |jq -r '.logs[0].events[1].attributes[].value' | grep "ununifi-" | sed 's/^.*"\(.*\)".*$/\1/');
export CLASS_ID=$CLASS_ID
ununifid tx nftmint mint-nft $CLASS_ID a10 $(ununifid keys show -a validator)  --from validator --chain-id test -y;

# to see if AfterNftUnlistedWithoutPayment hook method will be trigger
ununifid tx nftmarket listing $CLASS_ID a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block \
--note="";
sleep 15;
ununifid tx nftmarket cancel_listing $CLASS_ID a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block;

# to see if AfterNftPaymentWithCommission hook method will be triggered
ununifid tx nftmarket listing $CLASS_ID a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block \
--note="";
ununifid tx nftmarket placebid $CLASS_ID a10 100uguu --automatic-payment=true --chain-id=test --from=debug --keyring-backend=test --gas=300000 -y --broadcast-mode=block;
