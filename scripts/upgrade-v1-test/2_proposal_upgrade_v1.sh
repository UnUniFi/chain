#!/bin/bash
# READ ME
# This script is a test script to check basic operation of upgrade-v1.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.
set -e

# --upgrade-info '{"binaries":{"linux/amd64":"https://github.com/hikaruNagamine/shared/releases/download/v1/ununifid_v1?checksum=md5:d8852a87392511f8b31bfadcb35a536f"}}' \
ununifid tx gov submit-proposal software-upgrade v1-beta.2 \
--title upgrade-test-v1 \
--description upgrade \
--upgrade-height 20 \
--from validator-a  \
--deposit 20000000000uguu \
--yes \
--chain-id ununifi-upgrade-test-v1 | jq .;

# sleep 10;

# ununifid tx gov deposit 1 \
# 20000000000uguu --from validator-a --yes \
# --chain-id ununifi-upgrade-test-v1 | jq .;

ununifid tx gov vote 1 \
yes --from validator-a \
--yes --chain-id ununifi-upgrade-test-v1 | jq .;

ununifid query gov proposals;
