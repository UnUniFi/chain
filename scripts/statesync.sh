#!/bin/bash

# print every command
set -ux

rm -fr ~/.ununifi

# make sure ununifid version is over v1.0.0-beta.3.1
# git clone https://github.com/UnUniFi/chain.git ununifi
# cd ununifi
# git checkout v1.0.0-beta.3.1
# make install

ununifid init test

wget -O ~/.ununifi/config/genesis.json https://raw.githubusercontent.com/UnUniFi/network/main/launch/ununifi-beta-v1/genesis.json

INTERVAL=1000

# get trust hash and trust height

LATEST_HEIGHT=$(ununifid q block --node https://a.lcd.ununifi.cauchye.net:443 | jq -r .block.header.height)
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL))
TRUST_HASH=$(ununifid q block $BLOCK_HEIGHT --node https://a.lcd.ununifi.cauchye.net:443 | jq -r .block_id.hash)

# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"

# export state sync vars
export UNUNIFID_P2P_MAX_NUM_OUTBOUND_PEERS=200
export UNUNIFID_STATESYNC_ENABLE=true
export UNUNIFID_STATESYNC_RPC_SERVERS="https://a.lcd.ununifi.cauchye.net:443,https://rpc.ununifi.nodestake.top:443"
export UNUNIFID_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export UNUNIFID_STATESYNC_TRUST_HASH=$TRUST_HASH

peers="80587494080597cd89fb24d345d4569dc5a09b3b@57.128.19.124:26656,b1888101a21ed188a4741c332ef6799275f612d4@94.130.148.251:26656,553d7226aaee5a043b234300f57f99e74c81f10c@88.99.69.190:26656,c23c1909cf536048d03795511acbd6e2557a3f4a@94.130.108.145:26656,2d5d596edfa7e8f5d241c58677cb65b629fb0a56@157.90.116.92:26656"
sed -i.bak -e "s/^persistent_peers *=.*/persistent_peers = \"$peers\"/" ~/.ununifi/config/config.toml

ununifid start --minimum-gas-prices=0uguu
