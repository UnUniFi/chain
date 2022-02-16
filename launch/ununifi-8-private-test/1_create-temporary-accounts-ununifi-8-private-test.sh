#!/bin/bash

# Note: Execute this script in the same directory where the mnemonic backup files exist.

date

# # Renew all keys and Create backup files.
# # Note: If you want to renew keys, execute this section.
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-a-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-b-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-c-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-d-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-pricefeed-granter-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-a-pricefeed-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-b-pricefeed-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-c-pricefeed-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-d-pricefeed-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-a-faucet-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-b-faucet-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-c-faucet-private-test-temp.txt
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys mnemonic > cauchye-d-faucet-private-test-temp.txt

# Recover all keys from text file.
# Note: -it is NG. -i is OK.
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-a-private-test-temp --recover < cauchye-a-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-b-private-test-temp --recover < cauchye-b-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-c-private-test-temp --recover < cauchye-c-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-d-private-test-temp --recover < cauchye-d-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-pricefeed-granter-private-test-temp --recover < cauchye-pricefeed-granter-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-a-pricefeed-private-test-temp --recover < cauchye-a-pricefeed-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-b-pricefeed-private-test-temp --recover < cauchye-b-pricefeed-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-c-pricefeed-private-test-temp --recover < cauchye-c-pricefeed-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-d-pricefeed-private-test-temp --recover < cauchye-d-pricefeed-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add faucet --recover < cauchye-a-faucet-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-b-faucet-private-test-temp --recover < cauchye-b-faucet-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-c-faucet-private-test-temp --recover < cauchye-c-faucet-private-test-temp.txt
docker run -i -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid keys add cauchye-d-faucet-private-test-temp --recover < cauchye-d-faucet-private-test-temp.txt

sudo chown -c -R $USER:docker ~/.ununifi

date
