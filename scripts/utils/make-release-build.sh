#!/bin/bash
sudo rm ./docker-build/ununifid
docker build . -t ununifid-dev
docker run -it -v $PWD:/root ununifid-dev ash -c "cp /usr/bin/ununifid /root/docker-build"
md5sum ./docker-build/ununifid
./docker-build/ununifid version
