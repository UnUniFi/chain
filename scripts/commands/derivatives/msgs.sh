#!/bin/bash

# load variables
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../../setup/variables.sh

# if arg 0 debug
if [ "$1" = "debug" ]; then
    set -x
    $BINARY tx derivatives mint-lpt 100ubtc --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g';
    $BINARY tx derivatives open-position perpetual-futures 0uusdc ubtc uusdc long 1 1  --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
    $BINARY tx derivatives open-position perpetual-futures 2500000uusdc ubtc uusdc long 1 1  --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
    # $BINARY tx derivatives open-position perpetual-futures 1ubtc ubtc uusdc long 5 5  --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
    # $BINARY tx derivatives open-position perpetual-futures 10000ubtc ubtc uusdc long 5 5  --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
    $BINARY tx derivatives open-position perpetual-futures 2ubtc ubtc uusdc long 1 1  --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
    # $BINARY tx derivatives close-position 0 --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
    # end shell
    exit 0
fi

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
$BINARY q derivatives positions $USER_ADDRESS_1

# close potision
$BINARY tx derivatives close-position 1 --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'