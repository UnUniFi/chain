#!/bin/bash
# mkdir ~/faucet
# cd ~/faucet
# curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/faucet/deploy.sh
docker-compose down
docker system prune -a -f
curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/faucet/docker-compose.yml
curl -O https://raw.githubusercontent.com/UnUniFi/chain/main/faucet/nginx.conf
docker cp $(docker ps -qf "name=ununifid"):/usr/bin/ununifid ~/faucet/ununifid
docker-compose up -d
