#!/bin/bash
# READ ME
# This script is a test script to check basic operation of upgrade-v1.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.
set -e
ununifid query bank balances ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7
ununifid query bank balances ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz

ununifid tx gov submit-proposal software-upgrade upgrade_v1 \
--title upgrade-test-v1 \
--description upgrade \
--upgrade-info '{"binaries":{"linux/amd64":"https://github.com/hikaruNagamine/shared/releases/download/v1/ununifid_v1?checksum=md5:e337c3b632474de117aceff5c8cd381d"}}' \
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


ununifid query bank balances ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7
ununifid query bank balances ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz