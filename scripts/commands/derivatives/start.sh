#!/bin/sh

# Script to start network with pricefeed
# just run ./start.sh

rm -rf ~/.ununifi

set -o errexit -o nounset

# Build genesis
ununifid init --chain-id=test test
ununifid keys mnemonic >& validator.txt
ununifid keys mnemonic >& debug.txt
ununifid keys add validator --recover < validator.txt
ununifid keys add debug --recover < debug.txt
ununifid keys add pricefeed --recover < pricefeed.txt
ununifid add-genesis-account $(ununifid keys show validator --address) 100000000000000uguu,100000000000000ubtc,100000000000000uusd
ununifid add-genesis-account $(ununifid keys show debug --address) 100000000000000uguu,100000000000000ubtc,100000000000000uusd
ununifid add-genesis-account $(ununifid keys show pricefeed --address) 100000000000000uguu,100000000000000ubtc,100000000000000uusd
ununifid gentx validator 100000000uguu --chain-id=test
ununifid collect-gentxs

# Edit app.toml to enable unsafe-cors and set pruning everything to reduce disk usage.
sed -E -i '' "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/" ~/.ununifi/config/app.toml;
sed -E -i '' "s/pruning = \".*\"/pruning = \"everything\"/" ~/.ununifi/config/app.toml;
sed -i '' 's/mode = "full"/mode = "validator"/' ~/.ununifi/config/config.toml;
sed -i '' 's/minimum-gas-prices = ""/minimum-gas-prices = "0uguu"/' ~/.ununifi/config/app.toml;
sed -i '' 's/stake/uguu/g' ~/.ununifi/config/genesis.json;

# modify genesis.json
PRICEFEED=$(ununifid keys show pricefeed --address)
jq '.app_state.pricefeed.params.markets = [{ "market_id": "ubtc:usd", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED'" ], "active": true }, { "market_id": "ubtc:usd:30", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED'" ], "active": true }]'  ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.pool.base_lpt_mint_fee = "0.001"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.pool.base_lpt_redeem_fee = "0.001"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.pool.accepted_assets = [{"denom":"ubtc", "target_weight": "0.6"}, {"denom":"uusd", "target_weight":"0.4"}]' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.commission_rate = "0.001"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.margin_maintenance_rate = "0.5"' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.markets = [{"base_denom": "ubtc", "quote_denom": "uusd" }]' ~/.ununifi/config/genesis.json > temp.json ; mv temp.json ~/.ununifi/config/genesis.json;

# run pricefeed
./setup_pricefeed.sh

# Start node
ununifid start
