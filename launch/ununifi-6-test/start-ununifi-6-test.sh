#!/bin/bash

date

cd ~/ununifi
docker-compose down
docker system prune -a -f

docker pull ghcr.io/ununifi/ununifid:latest

rm ~/.ununifi/config/genesis.json
curl -L https://raw.githubusercontent.com/UnUniFi/chain/main/launch/ununifi-6-test/genesis.json -o ~/.ununifi/config/genesis.json

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid unsafe-reset-all
sudo chown -c -R $USER:docker ~/.ununifi

cd ~/ununifi
curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/docker-compose.yml

docker-compose up -d

date
