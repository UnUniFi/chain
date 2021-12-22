#!/bin/bash

date

# validators (vesting account)
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account cauchye-a-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account cauchye-b-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account tokyo-0-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account tokyo-1-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account genio01-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account zofuku-japan-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account zofuku-tokyo-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account chikako0903-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account kurata0211-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account keyplayers01-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account keyplayers02-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account tokyo-2-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account tokyo-3-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account tokyo-4-test-temp 7943000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account toko1631-test-temp 6170000000uguu --vesting-amount="5000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"

# other accounts (vesting account)
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account dev-team-temp 200000000000uguu --vesting-amount="200000000000uguu" --vesting-end-time="1651147074" --vesting-start-time="1640133940"

# other accounts
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account other-validators-temp 182628000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account airdrop-temp 30000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account cauchye-a-test-pricefeed-temp 1000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account cauchye-b-test-pricefeed-temp 1000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account other-oracles-temp 98000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account ecosystem-development-temp 170000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account marketing-temp 90000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account advisor-temp 10000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account partners-temp 100000000000uguu
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid add-genesis-account test-faucet-temp 1000000000000uguu,100000000ubtc

date
