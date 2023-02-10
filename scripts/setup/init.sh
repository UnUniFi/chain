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

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $VAL1 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER1 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER2 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER3 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show $USER4 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN  --home $CHAIN_DIR/$CHAINID_1

echo "Creating and collecting gentx..."
$BINARY gentx $VAL1 7000000000$BINARY_MAIN_TOKEN --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID_1

echo "Changing defaults config files..."
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' $CHAIN_DIR/$CHAINID_1/config/app.toml;
sed -i -e 's/minimum-gas-prices = ""/minimum-gas-prices = '\"0$BINARY_MAIN_TOKEN\"/'' $CHAIN_DIR/$CHAINID_1/config/app.toml;
sed -i 's/mode = "full"/mode = "validator"/' $CHAIN_DIR/$CHAINID_1/config/config.toml;
sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $CHAIN_DIR/$CHAINID_1/config/app.toml;