#!/bin/bash

CHAIN_ID="jpyx-1"
KEY_PASSPHRASE=''
VALIDATOR_NAME=""
VALIDATOR_ADDRESS=""
VALIDATOR_MNEMONIC=""

function add_key() {
  expect -c "
  set timeout 5
  spawn docker run -v $HOME/.jpyxd:/root/.jpyxd -v $HOME/.jpyxcli:/root/.jpyxcli -it jpyx jpyxcli keys add $1 --recover
  expect \"> Enter your bip39 mnemonic\"
  send \"$2\n\"
  expect \"Enter keyring passphrase:\"
  send \"$3\n\"
  expect \"Re-enter keyring passphrase:\"
  send \"$3\n\"
  interact
  "
}

function gen_tx() {
  sudo mkdir -p ~/.jpyxd/config/gentx
  expect -c "
  set timeout 5
  spawn docker run -v $HOME/.jpyxd:/root/.jpyxd -v $HOME/.jpyxcli:/root/.jpyxcli -it jpyx jpyxd gentx --amount $1 --name $2 --output-document "$3"
  expect \"Enter keyring passphrase:\"
  send \"$4\n\"
  expect \"Enter keyring passphrase:\"
  send \"$4\n\"
  expect \"Enter keyring passphrase:\"
  send \"$4\n\"
  interact
  "
}

sudo rm -rf ~/.jpyxd ~/.jpyxcli

docker build -t jpyx ../

docker run -v ~/.jpyxd:/root/.jpyxd -v ~/.jpyxcli:/root/.jpyxcli -it jpyx jpyxd init jpyx --chain-id "$CHAIN_ID"
docker run -v ~/.jpyxd:/root/.jpyxd -v ~/.jpyxcli:/root/.jpyxcli -it jpyx jpyxcli config chain-id "$CHAIN_ID"
docker run -v ~/.jpyxd:/root/.jpyxd -v ~/.jpyxcli:/root/.jpyxcli -it jpyx jpyxcli config trust-node true
add_key "$VALIDATOR_NAME" "$VALIDATOR_MNEMONIC" "$KEY_PASSPHRASE"

KEY_EXISTS=$(jq ".app_state.pricefeed.params.markets[0].oracles | contains([\"$VALIDATOR_ADDRESS\"])" ./genesis.json)
if [ $KEY_EXISTS = "false" ]; then
  jq ".app_state.pricefeed.params.markets[].oracles += [\"$VALIDATOR_ADDRESS\"]" ./genesis.json > ./genesis.json.tmp
  sudo mv ./genesis.json.tmp ~/.jpyxd/config/genesis.json
else
  sudo cp ./genesis.json ~/.jpyxd/config/genesis.json
fi

docker run -v ~/.jpyxd:/root/.jpyxd -v ~/.jpyxcli:/root/.jpyxcli -it jpyx jpyxd add-genesis-account $VALIDATOR_ADDRESS "500000000000ujsmn,500000000000token"
gen_tx "500000000000ujsmn" "$VALIDATOR_NAME" "/root/.jpyxd/config/gentx/gentx-validator.json"  "$KEY_PASSPHRASE"
docker run -v ~/.jpyxd:/root/.jpyxd -v ~/.jpyxcli:/root/.jpyxcli -it jpyx jpyxd collect-gentxs
