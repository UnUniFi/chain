#!/bin/bash

# This script will work on a.test.ununifi.cauchye.com

date

# Stop node
cd ~/ununifi
docker-compose down
sudo chown -c -R $USER:docker ~/.ununifi

# Delete or Edit old geneesis.json
# rm ~/.ununifi/config.genesis.json
# or edit something

# pull new image
docker pull ghcr.io/ununifi/ununifid:latest

# Note: Only first time on the node
# Execute initialization
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid init cauchye-c-test --chain-id ununifi-5-test
# sudo chown -c -R $USER:docker ~/.ununifi

# Show node info and record them
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid tendermint show-node-id
# docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid tendermint show-validator
# sudo chown -c -R $USER:docker ~/.ununifi

# a.test.ununifi.cauchye.net
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx private-test-cauchye-a-test 200000000000uguu --chain-id="ununifi-5-test" --from="private-test-cauchye-a-test" --ip="a.test.ununifi.cauchye.net" --moniker="cauchye-a-test" --identity="cauchye-a-test" --website="https://cauchye.com" --node-id="de321a535a5334268890940fa41429e3a0f9586a" --pubkey="ununifivalconspub1zcjduepq0ykrm3duk9dvjp3j38ey3zz39mz6nalcj0qunenmmdt920njq9ns92gept"
sudo chown -c -R $USER:docker ~/.ununifi

# b.test.ununifi.cauchye.net
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx private-test-cauchye-b-test 200000000000uguu --chain-id="ununifi-5-test" --from="private-test-cauchye-b-test" --ip="b.test.ununifi.cauchye.net" --moniker="cauchye-b-test" --identity="cauchye-b-test" --website="https://cauchye.com" --node-id="b13e4681fd890b5ea7f3621cb4b6fe73ee0e5c34" --pubkey="ununifivalconspub1zcjduepqw8kkjmyzwj32sjmasd0xlryy0ucwfur90x2uz0p07m053jrqtujqkeflw0"
sudo chown -c -R $USER:docker ~/.ununifi

# c.test.ununifi.cauchye.net
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx private-test-cauchye-c-test 200000000000uguu --chain-id="ununifi-5-test" --from="private-test-cauchye-c-test" --ip="c.test.ununifi.cauchye.net" --moniker="cauchye-c-test" --identity="cauchye-c-test" --website="https://cauchye.com" --node-id="46a4580084634cf24cf7ecc589a9f1eb47291074" --pubkey="ununifivalconspub1zcjduepq93nvwmkrfakhjnch7xpxa68a6wfdr9zsp7xekfgyygjnxcv6udcqcpmrlk"
sudo chown -c -R $USER:docker ~/.ununifi

# d.test.ununifi.cauchye.net
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid gentx private-test-cauchye-d-test 200000000000uguu --chain-id="ununifi-5-test" --from="private-test-cauchye-d-test" --ip="d.test.ununifi.cauchye.net" --moniker="cauchye-d-test" --identity="cauchye-d-test" --website="https://cauchye.com" --node-id="192ddcaea7789f65aa9b1f840ea963d4d0a61d8f" --pubkey="ununifivalconspub1zcjduepq7srlsqdwmwfn2rjalvhz0p5z29qww30l7jz448rrsjrsfrw394gq6ykyyu"
sudo chown -c -R $USER:docker ~/.ununifi

# collect-gentxs
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid collect-gentxs
sudo chown -c -R $USER:docker ~/.ununifi

# unsafe-reset-all
docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid unsafe-reset-all
sudo chown -c -R $USER:docker ~/.ununifi

date
