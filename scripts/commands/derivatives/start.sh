#!/bin/bash

# Script to start network with pricefeed
# just run ./start.sh
SCRIPT_DIR=$(cd $(dirname $0); pwd)

rm -rf ~/.ununifi

set -o errexit -o nounset

# Build genesis
ununifid init --chain-id=test test
ununifid keys mnemonic 2> validator.txt
ununifid keys mnemonic 2> debug.txt
ununifid keys add validator --recover < validator.txt
ununifid keys add debug --recover < debug.txt
ununifid keys add pricefeed --recover < $SCRIPT_DIR/pricefeed.txt
ununifid add-genesis-account $(ununifid keys show validator --address) 100000000000000uguu,100000000000000ubtc,100000000000000uusd
ununifid add-genesis-account $(ununifid keys show debug --address) 100000000000000uguu,100000000000000ubtc,100000000000000uusd
ununifid add-genesis-account $(ununifid keys show pricefeed --address) 100000000000000uguu,100000000000000ubtc,100000000000000uusd
ununifid gentx validator 100000000uguu --chain-id=test
ununifid collect-gentxs

# Edit app.toml to enable unsafe-cors and set pruning everything to reduce disk usage.
OS=$(uname -s)
if [ "$OS" == "Darwin" ]; then
  echo $OS
  sleep 1
  sed_i="sed -i ''"
elif [ "$OS" == "Linux" ]; then
  echo $OS
  sleep 1
  sed_i="sed -i"
fi

$sed_i "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/" ~/.ununifi/config/app.toml;
$sed_i "s/pruning = \".*\"/pruning = \"everything\"/" ~/.ununifi/config/app.toml;
$sed_i 's/mode = "full"/mode = "validator"/' ~/.ununifi/config/config.toml;
$sed_i 's/minimum-gas-prices = ""/minimum-gas-prices = "0uguu"/' ~/.ununifi/config/app.toml;
$sed_i 's/stake/uguu/g' ~/.ununifi/config/genesis.json;

# modify genesis.json
PRICEFEED=$(ununifid keys show pricefeed --address)
jq '.app_state.pricefeed.params.markets = [{ "market_id": "ubtc:usd", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED'" ], "active": true }, { "market_id": "ubtc:usd:30", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED'" ], "active": true }]'  ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.pool.base_lpt_mint_fee = "0.001"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.pool.base_lpt_redeem_fee = "0.001"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.pool.accepted_assets = [{"denom":"ubtc", "target_weight": "0.6"}, {"denom":"uusd", "target_weight":"0.4"}]' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.commission_rate = "0.001"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.margin_maintenance_rate = "0.5"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.imaginary_funding_rate_proportional_coefficient = "0.0005"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.markets = [{"base_denom": "ubtc", "quote_denom": "uusd" }]' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;

# run pricefeed
$SCRIPT_DIR/setup_pricefeed.sh

# Start node
ununifid start
