#!/bin/bash

# load variables
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../../setup/variables.sh

# mint lpt
$BINARY tx derivatives mint-lpt 100000ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g';
$BINARY tx derivatives mint-lpt 100000ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g';
# burn lpt
$BINARY tx derivatives burn-lpt 1 ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
$BINARY tx derivatives burn-lpt 1 ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'

# open-position perpetual-futures
$BINARY tx derivatives open-position perpetual-futures 100ubtc ubtc uusd long --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
$BINARY tx derivatives open-position perpetual-futures 100uusdc ubtc uusd short --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'

# query positions
$BINARY q derivatives positions $USER_ADDRESS_1

# close potision
$BINARY tx derivatives close-position 1 --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'