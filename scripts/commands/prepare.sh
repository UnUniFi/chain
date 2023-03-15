#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../setup/variables.sh

echo "execute tx"

MATCH_PATTERN="^.*hoge.*$"
MISMATCH_PATTERN="^.*Error.*$"
until  ununifid q block 2>&1 |grep "last_block_id" >/dev/null 2>&1 ; do
    printf 'waitting...'
    sleep 1
done
# $SCRIPT_DIR/nftmarket_v2.sh
