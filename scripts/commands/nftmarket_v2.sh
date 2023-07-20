#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/../setup/variables.sh

$BINARY tx nftbackedloan mint a10 a10 uri 888838 --home $CHAIN_DIR/$CHAINID_1 --from=$VAL1 $conf| jq .;
$BINARY tx nftbackedloan mint a10 a11 uri 888838 --home $CHAIN_DIR/$CHAINID_1 --from=$VAL1 $conf| jq .;

# auto-refinancing
$BINARY tx nftbackedloan listing a10 a10 --min-deposit-rate 0.01 --bid-token uguu -r --from $VAL1 $conf| jq . ;

# not auto-refinancing
# ununifid tx nftbackedloan listing a10 a10 --min-deposit-rate 0.1 --bid-token uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $VAL1 --keyring-backend test --gas 300000 -y| jq . ;
$BINARY tx nftbackedloan listing a10 a11 --min-deposit-rate 0.1 --bid-token uguu --from $VAL1 $conf| jq . ;

# ununifid tx nftbackedloan placebid a10 a10 100uguu 80uguu 0.1 20 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER1 --keyring-backend test --gas 300000  -y --broadcast-mode=block --automatic-payment| jq .;
# ununifid tx nftbackedloan placebid a10 a10 100uguu 1uguu 0.1 20 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER1 --keyring-backend test --gas 300000  -y --broadcast-mode=block --automatic-payment| jq .;

# ununifid tx nftbackedloan placebid a10 a10 100uguu 60uguu 0.1 20 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER1 --keyring-backend test --gas 300000  -y --broadcast-mode=block --automatic-payment| jq .;
$BINARY tx nftbackedloan placebid a10 a10 100uguu 50uguu 0.1 10000000000 --from $USER1 $conf| jq .;
# ununifid tx nftbackedloan placebid a10 a10 102uguu 20uguu 0.3 12000000 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER2 --keyring-backend test --gas 300000 -y --broadcast-mode=block --automatic-payment| jq .;
ununifid tx nftbackedloan placebid a10 a10 103uguu 2uguu 0.2 12000000 --from $USER3 $conf| jq .;
# not automatic payment
ununifid tx nftbackedloan placebid a10 a10 102uguu 20uguu 0.3 12000000 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER2 --keyring-backend test --gas 300000 -y --broadcast-mode=block | jq .;


# ununifid tx nftbackedloan placebid a10 a11 100uguu 50uguu 0.1 20 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER1 --keyring-backend test --gas 300000  -y --broadcast-mode=block --automatic-payment| jq .;
# ununifid tx nftbackedloan placebid a10 a11 100uguu 52uguu 0.2 120 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER2 --keyring-backend test --gas 300000 -y --broadcast-mode=block --automatic-payment| jq .;
# ununifid tx nftbackedloan placebid a10 a11 100uguu 51uguu 0.2 120 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER3 --keyring-backend test --gas 300000 -y --broadcast-mode=block --automatic-payment| jq .;

# normal borrowing
ununifid tx nftbackedloan borrow a10 a10 72uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;
# ununifid tx nftbackedloan borrow a10 a10 150uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;
# ununifid tx nftbackedloan borrow a10 a10 130uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;
# borrowing twice
# ununifid tx nftbackedloan borrow a10 a10 10uguu  --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;

# try over borrowing
# ununifid tx nftbackedloan borrow a10 a10 230uguu --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from=$VAL1 --keyring-backend test --gas 300000 -y --broadcast-mode=block| jq .;

# ununifid q nftbackedloan liquidation a10 a10


# cancel bid

## can cancelbid
# ununifid tx nftbackedloan cancelbid a10 a10 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER2 --keyring-backend test --gas 300000 -y --broadcast-mode=block | jq .;
## can't cancelbid
# ununifid tx nftbackedloan cancelbid a10 a10 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $USER3 --keyring-backend test --gas 300000 -y --broadcast-mode=block | jq .;

# ununifid tx nftbackedloan selling_decision a10 a10 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $VAL1 --keyring-backend test --gas 300000 -y| jq . ;

# cancel listing
# ununifid tx nftbackedloan cancel_listing a10 a11 --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test --from=$VAL1 --gas=300000 -y --broadcast-mode=block| jq .;
# repay
# ununifid tx nftbackedloan repay a10 a10 10uguu  --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $VAL1 --keyring-backend test --gas 300000 -y| jq . ;


watch ununifid q nftbackedloan nft_bids a10 a10

# sleep 5
# ununifid tx nftbackedloan repay a10 a10 50uguu  --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --from $VAL1 --keyring-backend test --gas 300000 -y| jq . ;

# watch ununifid q bank balances $USER_ADDRESS_1 