#!/bin/bash
# READ ME
# This script is a test script to check basic operation of upgrade-v1.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.
set -e
cd $HOME
sudo rm -rf ~/.ununifi
ununifid init ununifi-upgrade-test-v2 --chain-id ununifi-upgrade-test-v2
wget https://raw.githubusercontent.com/UnUniFi/network/test/upgrade-v1/test/upgrade-v2/genesis-mainnet-mock-added-users-upgrade-v2.json -O  ~/.ununifi/config/genesis.json
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' ~/.ununifi/config/app.toml;
sed -i 's/minimum-gas-prices = ".*"/minimum-gas-prices = "0uguu"/' ~/.ununifi/config/app.toml;
sed -i 's/stake/uguu/g' ~/.ununifi/config/genesis.json;
# jq '.app_state.bank.params.default_send_enabled = false'  ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
# jq '.app_state.gov.voting_params.voting_period = "20s"'  ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
ununifid keys add validator-a --recover < ~/backup-validator-a-mnemonic.txt;
ununifid add-genesis-account validator-a 125000000000000uguu;
ununifid gentx validator-a 120000000000000uguu --chain-id ununifi-upgrade-test-v2 --keyring-backend test;
ununifid collect-gentxs;

go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0
mkdir -p $DAEMON_HOME/cosmovisor
mkdir -p $DAEMON_HOME/cosmovisor/genesis
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
mkdir -p $DAEMON_HOME/cosmovisor/upgrades
cp ~/go/bin/ununifid $DAEMON_HOME/cosmovisor/genesis/bin
# ~/go/bin/cosmovisor start

sudo touch /lib/systemd/system/cosmovisor.service
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment=\"DAEMON_NAME=ununifid\"
Environment=\"DAEMON_HOME=$HOME/.ununifi\"
Environment=\"DAEMON_RESTART_AFTER_UPGRADE=true\"
Environment=\"DAEMON_ALLOW_DOWNLOAD_BINARIES=false\"
Environment=\"DAEMON_LOG_BUFFER_SIZE=512\"
Environment=\"UNSAFE_SKIP_BACKUP=true\"
User=$USER
ExecStart=$HOME/go/bin/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target" | sudo tee /lib/systemd/system/cosmovisor.service

sudo systemctl daemon-reload
sudo systemctl restart systemd-journald
sudo systemctl enable cosmovisor
sudo systemctl start cosmovisor
sudo systemctl status cosmovisor
