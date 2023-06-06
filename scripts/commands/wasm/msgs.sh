#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../../setup/variables.sh


# $BINARY tx wasm store $WASM_A_LOCATION --from $VAL1 $conf --gas=auto --gas-adjustment=1.3
# $BINARY tx wasm instantiate 1 '{"owner":"'$VALIDATOR'", "unbond_period":1, "deposit_denom": "stake"}' --from $VAL1 --label "BaseStrategy"   --no-admin $conf


# option
# | grep txhash | awk '{ print $2 }'| xargs -I {} sh -c 'sleep 5; $0 q tx {}' $BINARY

contract_denom=uguu
$BINARY tx wasm store $WASM_A_LOCATION --from $VAL1 $conf --gas=auto --gas-adjustment=1.3 | grep txhash | awk '{ print $2 }'| xargs -I {} sh -c 'sleep 5; $0 q tx {}' $BINARY
$BINARY tx wasm instantiate 1 '{"owner":"'$VAL_ADDRESS_1'", "unbond_period":1, "deposit_denom": "'$contract_denom'"}' --from $VAL1 --label "BaseStrategy"   --no-admin $conf | grep txhash | awk '{ print $2 }'| xargs -I {} sh -c 'sleep 5; $0 q tx {}' $BINARY
$BINARY tx wasm execute $CONTRACT '{"stake":{}}' --amount=1000uguu --from $VAL1 --gas=9223372036854775807 $conf | grep txhash | awk '{ print $2 }'| xargs -I {} sh -c 'sleep 5; $0 q tx {}' $BINARY
$BINARY tx wasm execute $CONTRACT '{"add_rewards":{}}' --amount=1000uguu --from $VAL1 --gas=9223372036854775807 $conf | grep txhash | awk '{ print $2 }'| xargs -I {} sh -c 'sleep 5; $0 q tx {}' $BINARY
# $BINARY tx wasm execute $CONTRACT '{"stake":{}}' --amount=1000stake --from $VAL1 --gas=auto --gas-adjustment=1.3 $conf | grep txhash | awk '{ print $2 }'| xargs -I {} sh -c 'sleep 5; $0 q tx {}' $BINARY
$BINARY export --home=./data/test --modules-to-export wasm


CONTRACT=ununifi14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sm5z28e

$BINARY query wasm list-code 
$BINARY query wasm list-contract-by-code 1
# $BINARY tx wasm execute $CONTRACT '{"stake":{}}' --from $VAL1 $conf --gas=auto --gas-adjustment=1.3 

$BINARY tx wasm execute $CONTRACT '{"stake":{}}' --amount=1000stake --from $VAL1 --gas=auto --gas-adjustment=1.3 $conf


$BINARY tx wasm execute $CONTRACT '{"stake":{}}' --from $VAL1 $conf --gas=9223372036854775807
# --gas=18446744073709551615
# $BINARY tx wasm execute $CONTRACT '{"add_rewards":{}}' --from $VAL1 $conf --gas=auto --gas-adjustment=1.3 
$BINARY tx wasm execute $CONTRACT '{"add_rewards":{}}' --from $VAL1 $conf --gas=auto --gas-adjustment=1.3 
$BINARY tx wasm execute $CONTRACT '{"unstake":{"amount":"1000"}}' --from=validator --gas=auto --gas-adjustment=1.3 --chain-id=test -y --keyring-backend=test
# $BINARY tx wasm execute $CONTRACT '{"update_config":{"owner":"'$VALIDATOR'","unbond_period":0,"deposit_denom":"stake"}}' --from=validator --gas=auto --gas-adjustment=1.3 --chain-id=test -y --keyring-backend=test 

