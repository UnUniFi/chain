#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
mkdir pricefeed
cd pricefeed
curl -O https://raw.githubusercontent.com/ununifi/utils/main/projects/pricefeed/docker-compose.yml
cp $SCRIPT_DIR/.env.pricefeed .env

docker-compose up -d
