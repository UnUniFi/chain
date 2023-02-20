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

# error show command
# ununifid tx derivatives open-position perpetual-futures 100ubtc 10ubtc 10usd  --from=$USER1 $conf | jq .raw_log | sed 's/\\n/\n/g'
# ununifid tx derivatives open-position perpetual-futures 100ubtc ubtc uusd  --from=$USER1 $conf | jq .
