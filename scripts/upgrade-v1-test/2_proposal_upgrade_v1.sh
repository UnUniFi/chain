#!/bin/bash
# READ ME
# This script is a test script to check basic operation of upgrade-v1.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.
set -e
ununifid tx vesting create-vesting-account ununifi16ayyysehst594k98a7leym6l5jrrhgf9yk9hn5 10000000uguu 1660665600 --from=validator-a --delayed=true --yes --chain-id ununifi-upgrade-test-v1
# jq '.app_state.bank.params.default_send_enabled = false'  ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;

echo "token sendor"
ununifid query bank balances ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7
echo "faucet"
ununifid query bank balances ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz
ununifid query bank balances ununifi14x04hcu8gmku53ll04v96tdgae84h2ylmkal9k
ununifid query bank balances ununifi1mtvjd2rsyll8kps6qqkxd6p78mr8qkjx27tn2p
ununifid query bank balances ununifi16ayyysehst594k98a7leym6l5jrrhgf9yk9hn5

ununifid tx gov submit-proposal software-upgrade upgrade_v1 \
--title upgrade-test-v1 \
--description upgrade \
--upgrade-info '{"binaries":{"linux/amd64":"https://github.com/hikaruNagamine/shared/releases/download/v1/ununifid_v1.1?checksum=md5:d765ca3db2c6bbbf29c33632bc94fa48"}}' \
--upgrade-height 20 \
--from validator-a  \
--yes \
--chain-id ununifi-upgrade-test-v1 | jq .;

sleep 10;

ununifid tx gov deposit 1 \
10000000uguu --from validator-a --yes \
--chain-id ununifi-upgrade-test-v1 | jq .;

ununifid tx gov vote 1 \
yes --from validator-a \
--yes --chain-id ununifi-upgrade-test-v1 | jq .;

ununifid query gov proposals;


echo "validator"
ununifid query bank balances ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7
echo "faucet"
ununifid query bank balances ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz
ununifid query bank balances ununifi14x04hcu8gmku53ll04v96tdgae84h2ylmkal9k
ununifid query bank balances ununifi1mtvjd2rsyll8kps6qqkxd6p78mr8qkjx27tn2p
ununifid query bank balances ununifi16ayyysehst594k98a7leym6l5jrrhgf9yk9hn5
