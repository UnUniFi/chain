#!/bin/bash

# load variables
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../../setup/variables.sh

# postPrice
$BINARY tx pricefeed postprice uusdc:ubtc 24528.185864015486004064 60 --from=$PRICEFEED $conf | jq .

# mint lpt
$BINARY tx derivatives mint-lpt 100000ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g';
$BINARY tx derivatives mint-lpt 100000ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g';
# burn lpt
$BINARY tx derivatives burn-lpt 1 ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
$BINARY tx derivatives burn-lpt 1 ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'

# open-position perpetual-futures
$BINARY tx derivatives open-position perpetual-futures 100ubtc ubtc uusdc long --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
$BINARY tx derivatives open-position perpetual-futures 100uusdc ubtc uusdc short --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'

# query positions
$BINARY q derivatives positions $(ununifid keys show -a $USER1) --home=$NOME_HOME
