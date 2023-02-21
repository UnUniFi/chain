#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../setup/variables.sh

echo "execute tx"
# SCRIPT_DIR/nftmarket.sh
MATCH_PATTERN="^.*hoge.*$"
MISMATCH_PATTERN="^.*Error.*$"
until  ununifid q block 2>&1 |grep "last_block_id" >/dev/null 2>&1 ; do
    printf 'waitting...'
    sleep 1
done
$SCRIPT_DIR/derivatives/msgs.sh