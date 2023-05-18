#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../setup/variables.sh

# make backup data in alpha node
curl http://$ALPHA_NODE_URL:3030/custom-make-backup
sleep 1
# if ununifi.tar.gz exists, remove it
if [ -f ununifi.tar.gz ]; then
  rm ununifi.tar.gz
fi
# transfer ununifi.tar.gz from alpha node to local
scp -i $SSH_PREV_KEY_LOCATION root@$ALPHA_NODE_URL:/root/ununifi.tar.gz ununifi.tar.gz
echo "copy alpha node data"
# if data dir exists, remove it
if [ -d data ]; then
  rm -rf data
fi
# copy debug dir
mkdir -p data/test
tar -xvf ununifi.tar.gz --strip-components=2 -C data/test
rm -rf data/test/cosmovisor
