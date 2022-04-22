#!/bin/bash

date

docker run -it -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid:test ununifid collect-gentxs
sudo chown -c -R $USER:docker ~/.ununifi

date
