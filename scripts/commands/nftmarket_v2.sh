#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../setup/variables.sh

ununifid tx nftmarket mint a10 a10 uri 888838 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test --from=$VAL1 --gas=300000 -y --broadcast-mode=block| jq .;
ununifid tx nftmarket mint a10 a11 uri 888838 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test --from=$VAL1 --gas=300000 -y --broadcast-mode=block| jq .;

# auto-refinancing
# ununifid tx nftmarket listing a10 a10 --min-minimum-deposit-rate 0.1 --bid-token uguu -r  --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $VAL1 --keyring-backend test --gas 300000 -y| jq . ;

# not auto-refinancing
ununifid tx nftmarket listing a10 a10 --min-minimum-deposit-rate 0.1 --bid-token uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $VAL1 --keyring-backend test --gas 300000 -y| jq . ;
ununifid tx nftmarket listing a10 a11 --min-minimum-deposit-rate 0.1 --bid-token uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $VAL1 --keyring-backend test --gas 300000 -y| jq . ;

ununifid tx nftmarket placebid a10 a10 100uguu 50uguu 0.1 50 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER1 --keyring-backend test --gas 300000  -y --broadcast-mode=block --automatic-payment| jq .;
ununifid tx nftmarket placebid a10 a10 100uguu 52uguu 0.2 120 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER2 --keyring-backend test --gas 300000 -y --broadcast-mode=block --automatic-payment| jq .;
ununifid tx nftmarket placebid a10 a10 100uguu 51uguu 0.2 120 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER3 --keyring-backend test --gas 300000 -y --broadcast-mode=block --automatic-payment| jq .;

# ununifid tx nftmarket placebid a10 a11 100uguu 50uguu 0.1 20 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER1 --keyring-backend test --gas 300000  -y --broadcast-mode=block --automatic-payment| jq .;
# ununifid tx nftmarket placebid a10 a11 100uguu 52uguu 0.2 120 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER2 --keyring-backend test --gas 300000 -y --broadcast-mode=block --automatic-payment| jq .;
# ununifid tx nftmarket placebid a10 a11 100uguu 51uguu 0.2 120 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER3 --keyring-backend test --gas 300000 -y --broadcast-mode=block --automatic-payment| jq .;

# normal borrowing
ununifid tx nftmarket borrow a10 a10 130uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;
# borrowing twice
# ununifid tx nftmarket borrow a10 a10 10uguu  --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;

# try over borrowing
# ununifid tx nftmarket borrow a10 a10 230uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;
