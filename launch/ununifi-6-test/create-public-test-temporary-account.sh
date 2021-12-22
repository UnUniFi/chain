#!/bin/bash

date

# validators
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-a-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-b-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-0-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-1-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add genio01-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add zofuku-japan-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add zofuku-tokyo-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add chikako0903-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add kurata0211-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add keyplayers01-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add keyplayers02-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-2-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-3-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add tokyo-4-test-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add toko1631-test-temp

# other accounts
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add dev-team-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add other-validators-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add airdrop-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-a-test-pricefeed-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add cauchye-b-test-pricefeed-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add other-oracles-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add ecosystem-development-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add marketing-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add advisor-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add partners-temp
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid keys add test-faucet-temp

date
