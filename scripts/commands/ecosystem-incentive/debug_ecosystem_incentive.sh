#!/bin/bash

cd ./build

set -uxe

# delegate from debug account to check the reward for the stakers in ecosystem-incentive
VAL_ADDR=$(./ununifid q staking validators -o=json | jq -r '.validators[0].operator_address')
./ununifid tx staking delegate $VAL_ADDR 1000000stake --chain-id=test --from=debug --keyring-backend=test --gas=300000 -y;

TX_HASH=$(./ununifid tx nftmint create-class Test ipfs://testcid/ 1000 0  --from validator --chain-id test -y --keyring-backend=test -o json | jq -r '.txhash')
sleep 5;
export CLASS_ID=$(./ununifid q tx $TX_HASH -o=json | jq -r '.logs[0].events[1].attributes[].value' | grep "ununifi-" | sed 's/^.*"\(.*\)".*$/\1/');

./ununifid tx nftmint mint-nft $CLASS_ID a10 $(./ununifid keys show -a validator --keyring-backend=test)  --from validator --chain-id test -y --keyring-backend=test;
sleep 5;

./ununifid tx ecosystem-incentive register --register-file ../scripts/commands/ecosystem-incentive/register.json --from validator --chain-id test -y --keyring-backend=test;
sleep 5;
# # to see if AfterNftUnlistedWithoutPayment hook method will be trigger
# ./ununifid tx nftbackedloan listing $CLASS_ID a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block \
# --note="";
# sleep 15;
# ./ununifid tx nftbackedloan cancel_listing $CLASS_ID a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block;

# to see if AfterNftPaymentWithCommission hook method will be triggered
./ununifid tx nftbackedloan listing $CLASS_ID a10 --min-deposit-rate=0.001 --bid-token=uguu --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y \
--note='{"version":"v1","incentive_unit_id":"incentive-unit-1"}';
sleep 5;
./ununifid tx nftbackedloan placebid $CLASS_ID a10 100000000uguu 100000uguu 0.01 100000 --automatic-payment=true --chain-id=test --from=debug --keyring-backend=test --gas=300000 -y;
sleep 5;

./ununifid tx nftbackedloan selling_decision $CLASS_ID a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y;
sleep 5;

# expect accumulating rewards
INCENTIVE_RECEIVER=$(cat ../scripts/commands/ecosystem-incentive/register.json | jq -r '.["subject_addrs"][0]')
./ununifid q ecosystem-incentive  all-rewards $INCENTIVE_RECEIVER

# queries
./ununifid q ecosystem-incentive recorded-incentive-unit-id $CLASS_ID a10

# withdraw reward txs
# ./ununifid tx ecosystem-incentive withdraw-all-rewards --from debug --chain-id=test -y;
# ./ununifid tx ecosystem-incentive withdraw-reward uguu --from debug --chain-id=test -y;

./ununifid q distribution rewards $(./ununifid keys show -a validator --keyring-backend=test) $VAL_ADDR
./ununifid q distribution rewards $(./ununifid keys show -a debug --keyring-backend=test) $VAL_ADDR

# ./ununifid tx distribution withdraw-all-rewards --from=validator --chain-id=test -y --keyring-backend=test;
# ./ununifid tx distribution withdraw-all-rewards --from=debug --chain-id=test -y --keyring-backend=test;
# ./ununifid q bank balances $(./ununifid keys show -a validator --keyring-backend=test)
