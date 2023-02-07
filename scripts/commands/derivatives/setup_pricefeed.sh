#!/bin/bash

mkdir pricefeed
cd pricefeed
curl -O https://raw.githubusercontent.com/ununifi/utils/main/projects/pricefeed/docker-compose.yml
cp ../.env.pricefeed .env

docker-compose up -d
