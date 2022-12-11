#!/bin/bash
# READ ME
# This script is a test script to check basic operation of upgrade-v1.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.
set -e

cd $HOME

# apt update & upgrade
sudo apt update && sudo apt upgrade -y;

# tools install
sudo apt install jq git build-essential curl -y;

# go install
wget https://go.dev/dl/go1.18.linux-amd64.tar.gz;
sudo rm -rf /usr/local/go;
sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz;
export PATH=$PATH:/usr/local/go/bin;

# get ununifid(v1.0.0-beta.1)
sudo rm -rf ~/chain_repo
git clone https://github.com/UnUniFi/chain chain_repo;
cd chain_repo;
git checkout main;
git pull;
make install;
# sudo cp ~/go/pkg/mod/github.com/\!cosm\!wasm/wasmvm@v1.0.0-beta10/api/libwasmvm.so /usr/local/lib/
# sudo ldconfig
~/go/bin/ununifid version;

# setup path 
cd $HOME;
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile;
echo "export PATH=\$PATH:\$HOME/go/bin" >> ~/.profile;
echo "export TARGET=ununifid" >> ~/.profile;
echo "export TARGET_HOME=.ununifi" >> ~/.profile;
echo "export MONIKER=ununifi-upgrade-test-v2" >> ~/.profile;
echo "export CHAIN_ID=ununifi-upgrade-test-v2" >> ~/.profile;
echo "export DAEMON_NAME=\$TARGET" >> ~/.profile;
echo "export DAEMON_HOME=\$HOME/\$TARGET_HOME" >> ~/.profile;
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=true" >> ~/.profile;
echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile;
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile;

# ufw
sudo ufw allow 26656;
sudo ufw allow 1317;
sudo ufw reload;
sudo ufw enable;
sudo ufw status;

# dammy user
echo "harsh lunar orient canal chalk pupil pupil duck scorpion mandate crack artwork token smart elevator eternal end change cup thought yellow trust brass busy" > ~/backup-validator-a-mnemonic.txt
# echo "cruel melt pulse sniff margin inspire frequent ostrich credit shop real ankle anger human vintage ribbon make cricket miracle pole thought rebel fame skull" > ~/backup-faucet-a-mnemonic.txt
# echo "idle mass away toward habit renew month awkward border drill identify wrong recall true before announce exist satoshi large mountain enough note eager move" > ~/backup-faucet-b-mnemonic.txt
# echo "jewel feel valley thank prize okay vehicle recycle test asthma affair flush air slab cry exercise tide park girl need kiss garden thank step" > ~/backup-faucet-c-mnemonic.txt
# echo "seat top exile tank wrestle aunt love debris nasty source refuse cave zone engine cloud scene figure fashion pass bench tourist pulse useful ankle" > ~/backup-faucet-d-mnemonic.txt

# source ~/.profile;