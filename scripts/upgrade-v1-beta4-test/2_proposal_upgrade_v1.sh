#!/bin/bash
# READ ME
# This script is a test script to check basic operation of upgrade-v1-beta4.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.
set -e

# --upgrade-info '{"binaries":{"linux/amd64":"https://github.com/hikaruNagamine/shared/releases/download/v2/ununifid_v2?checksum=md5:2754c2ea536b8ad71ef8d7de475557ce"}}' \
ununifid tx gov submit-proposal software-upgrade v1-beta.4 \
--title upgrade-test-v1-beta4 \
--description upgrade \
--upgrade-height 20 \
--from validator-a  \
--deposit 20000000000uguu \
--yes \
--chain-id ununifi-upgrade-test-v1-beta4 | jq .;

# sleep 10;

# ununifid tx gov deposit 1 \
# 20000000000uguu --from validator-a --yes \
# --chain-id ununifi-upgrade-test-v1-beta4 | jq .;

ununifid tx gov vote 1 \
yes --from validator-a \
--yes --chain-id ununifi-upgrade-test-v1-beta4 | jq .;

ununifid query gov proposals;

mkdir -p ~/.ununifi/cosmovisor/upgrades/v1-beta.4/bin
# Prepare binary files for updates in advance and place them in the update folder 
cp ~/ununifid_v1_beta4 ~/.ununifi/cosmovisor/upgrades/v1-beta.4/bin/ununifid
