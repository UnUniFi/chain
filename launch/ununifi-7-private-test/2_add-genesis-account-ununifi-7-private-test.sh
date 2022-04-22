#!/bin/bash

date

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-a-private-test-temp 240000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-b-private-test-temp 240000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-c-private-test-temp 240000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-d-private-test-temp 240000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-pricefeed-granter-private-test-temp 20000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-a-pricefeed-private-test-temp 5000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-b-pricefeed-private-test-temp 5000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-c-pricefeed-private-test-temp 5000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-d-pricefeed-private-test-temp 5000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-a-faucet-private-test-temp 500000000000uguu,5000000000ubtc
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid add-genesis-account cauchye-b-faucet-private-test-temp 500000000000uguu,5000000000ubtc

sudo chown -c -R $USER:docker ~/.ununifi

date
