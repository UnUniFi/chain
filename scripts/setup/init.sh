#!/bin/bash

# Load shell variables
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/variables.sh

# Stop if it is already running 
if pgrep -x "$BINARY" >/dev/null; then
    echo "Terminating $BINARY..."
    killall $BINARY
fi

echo "Removing previous data..."
rm -rf $CHAIN_DIR/$CHAINID_1 &> /dev/null

# Add directories for both chains, exit if an error occurs
if ! mkdir -p $CHAIN_DIR/$CHAINID_1 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

echo "Initializing $CHAINID_1..."
$BINARY init test --home $CHAIN_DIR/$CHAINID_1 --chain-id=$CHAINID_1

# change main token
sed -i -e 's/\bstake\b/'$BINARY_MAIN_TOKEN'/g' $CHAIN_DIR/$CHAINID_1/config/genesis.json

echo "Adding genesis accounts..."
echo $VAL_MNEMONIC_1    | $BINARY keys add $VAL1 --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $USER_MNEMONIC_1 | $BINARY keys add $USER1 --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $USER_MNEMONIC_2 | $BINARY keys add $USER2 --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $USER_MNEMONIC_3 | $BINARY keys add $USER3 --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $USER_MNEMONIC_4 | $BINARY keys add $USER4 --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $PRICEFEED_MNEMONIC | $BINARY keys add $PRICEFEED --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $VAL1 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER1 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER2 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER3 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER4 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $PRICEFEED --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN,100000000000000ubtc,100000000000000uusd  --home $CHAIN_DIR/$CHAINID_1

echo "Creating and collecting gentx..."
$BINARY gentx $VAL1 7000000000$BINARY_MAIN_TOKEN --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID_1

echo "Changing defaults config files..."
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

sed -i '/\[api\]/,+3 s/enable = false/enable = true/' $CHAIN_DIR/$CHAINID_1/config/app.toml;
sed -i -e 's/minimum-gas-prices = ""/minimum-gas-prices = '\"0$BINARY_MAIN_TOKEN\"/'' $CHAIN_DIR/$CHAINID_1/config/app.toml;
$sed_i 's/mode = "full"/mode = "validator"/' $CHAIN_DIR/$CHAINID_1/config/config.toml;
$sed_i "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/" $CHAIN_DIR/$CHAINID_1/config/app.toml;
$sed_i 's/mode = "full"/mode = "validator"/' $CHAIN_DIR/$CHAINID_1/config/config.toml;
$sed_i 's/minimum-gas-prices = ""/minimum-gas-prices = "0uguu"/' $CHAIN_DIR/$CHAINID_1/config/app.toml;
$sed_i 's/stake/uguu/g' $CHAIN_DIR/$CHAINID_1/config/genesis.json;


# PRICEFEED=$(ununifid keys show pricefeed --address)
jq '.app_state.pricefeed.params.markets = [{ "market_id": "ubtc:usd", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED_ADDRESS'" ], "active": true }, { "market_id": "ubtc:usd:30", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED_ADDRESS'" ], "active": true }]'  $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.derivatives.params.pool.base_lpt_mint_fee = "0.001"' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.derivatives.params.pool.base_lpt_redeem_fee = "0.001"' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.derivatives.params.pool.accepted_assets = [{"denom":"ubtc", "target_weight": "0.6"}, {"denom":"uusd", "target_weight":"0.4"}]' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.commission_rate = "0.001"' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.margin_maintenance_rate = "0.5"' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.imaginary_funding_rate_proportional_coefficient = "0.0005"' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.markets = [{"base_denom": "ubtc", "quote_denom": "uusd" }]' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;
jq '.app_state.bank.denom_metadata = [{"base" : "ubtc" , "symbol": "ubtc"}, {"base" : "uusd", "symbol": "usd"}]' $CHAIN_DIR/$CHAINID_1/config/genesis.json > temp.json ; mv temp.json $CHAIN_DIR/$CHAINID_1/config/genesis.json;