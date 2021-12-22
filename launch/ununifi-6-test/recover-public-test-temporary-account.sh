#!/bin/bash

date

# validators
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-a-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-b-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-0-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-1-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add genio01-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add zofuku-japan-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add zofuku-tokyo-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add chikako0903-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add kurata0211-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add keyplayers01-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add keyplayers02-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-2-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-3-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-4-test-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add toko1631-test-temp --interactive --recover

# other accounts
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add dev-team-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add other-validators-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add airdrop-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-a-test-pricefeed-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-b-test-pricefeed-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add other-oracles-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add ecosystem-development-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add marketing-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add advisor-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add partners-temp --interactive --recover
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add test-faucet-temp --interactive --recover

date
