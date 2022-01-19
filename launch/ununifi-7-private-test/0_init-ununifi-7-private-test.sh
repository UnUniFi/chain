#!/bin/bash
docker pull ghcr.io/ununifi/ununifid:test
docker run -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid init cauchye-a-private-test --chain-id ununifi-7-private-test
sudo chown -c -R $USER:docker ~/.ununifi
