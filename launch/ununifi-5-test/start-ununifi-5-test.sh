#!/bin/bash

date

docker-compose down

docker pull ghcr.io/ununifi/ununifid:latest

rm ~/.ununifi/config/genesis.json
curl -L https://raw.githubusercontent.com/UnUniFi/chain/main/launch/ununifi-5-test/genesis.json -o ~/.ununifi/config/genesis.json

cd ~/ununifi
curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/docker-compose.yml

docker-compose up -d

date
