#!/bin/bash

# load variables
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../../setup/variables.sh
conf="--home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test --gas 300000 -y --broadcast-mode=block"

ununifid tx derivatives mint-lpt 100000ubtc --from=debug --chain-id=test --broadcast-mode block -y;

# error show command
# ununifid tx derivatives open-position perpetual-options 100ubtc 10ubtc 10usd  --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
ununifid tx derivatives open-position perpetual-options 100ubtc 10ubtc 10usd  --from=$USER1 $conf | jq .