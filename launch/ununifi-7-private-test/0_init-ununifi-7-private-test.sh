#!/bin/bash
docker pull ghcr.io/ununifi/ununifid:test
# Note: You need to change from cauchye-a-private-test to your moniker and change from ununifi-7-private-test to your chain id.
docker run -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid init cauchye-a-private-test --chain-id ununifi-7-private-test
sudo chown -c -R $USER:docker ~/.ununifi
