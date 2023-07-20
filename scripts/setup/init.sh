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
rm -rf $NODE_HOME &> /dev/null

# Add directories for both chains, exit if an error occurs
if ! mkdir -p $NODE_HOME 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

echo "Initializing $CHAINID_1..."
$BINARY init test --home $NODE_HOME --chain-id=$CHAINID_1

# change main token
sed -i -e 's/\bstake\b/'$BINARY_MAIN_TOKEN'/g' $NODE_HOME/config/genesis.json

echo "Adding genesis accounts..."
echo $VAL_MNEMONIC_1    | $BINARY keys add $VAL1 --home $NODE_HOME --recover --keyring-backend=test
echo $FAUCET_MNEMONIC_1 | $BINARY keys add $FAUCET --home $NODE_HOME --recover --keyring-backend=test
echo $USER_MNEMONIC_1 | $BINARY keys add $USER1 --home $NODE_HOME --recover --keyring-backend=test
echo $USER_MNEMONIC_2 | $BINARY keys add $USER2 --home $NODE_HOME --recover --keyring-backend=test
echo $USER_MNEMONIC_3 | $BINARY keys add $USER3 --home $NODE_HOME --recover --keyring-backend=test
echo $USER_MNEMONIC_4 | $BINARY keys add $USER4 --home $NODE_HOME --recover --keyring-backend=test
echo $PRICEFEED_MNEMONIC | $BINARY keys add $PRICEFEED --home $NODE_HOME --recover --keyring-backend=test

$BINARY genesis add-genesis-account $($BINARY --home $NODE_HOME keys show $VAL1 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN --home $NODE_HOME
$BINARY genesis add-genesis-account $($BINARY --home $NODE_HOME keys show $FAUCET --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN,100000000000000ubtc,100000000000000uusdc --home $NODE_HOME
$BINARY genesis add-genesis-account $($BINARY --home $NODE_HOME keys show $USER1 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN,100000000000000ubtc,100000000000000uusdc --home $NODE_HOME
$BINARY genesis add-genesis-account $($BINARY --home $NODE_HOME keys show $USER2 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN,100000000000000ubtc,100000000000000uusdc --home $NODE_HOME
$BINARY genesis add-genesis-account $($BINARY --home $NODE_HOME keys show $USER3 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN,100000000000000ubtc,100000000000000uusdc --home $NODE_HOME
$BINARY genesis add-genesis-account $($BINARY --home $NODE_HOME keys show $USER4 --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN,100000000000000ubtc,100000000000000uusdc  --home $NODE_HOME
$BINARY genesis add-genesis-account $($BINARY --home $NODE_HOME keys show $PRICEFEED --keyring-backend test -a) 100000000000$BINARY_MAIN_TOKEN,100000000000000ubtc,100000000000000uusdc  --home $NODE_HOME

echo "Creating and collecting gentx..."
$BINARY genesis gentx $VAL1 7000000000$BINARY_MAIN_TOKEN --home $NODE_HOME --chain-id $CHAINID_1 --keyring-backend test
$BINARY genesis collect-gentxs --home $NODE_HOME

echo "Changing defaults config files..."
OS=$(uname -s)
if [ "$OS" == "Darwin" ]; then
  echo $OS
  sleep 1
  sed_i="sed -i '' "
elif [ "$OS" == "Linux" ]; then
  echo $OS
  sleep 1
  sed_i="sed -i"
fi

$sed_i '/\[api\]/,+3 s/enable = false/enable = true/' $NODE_HOME/config/app.toml;
$sed_i 's/mode = "full"/mode = "validator"/' $NODE_HOME/config/config.toml;
$sed_i "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/" $NODE_HOME/config/app.toml;
$sed_i 's/minimum-gas-prices = ""/minimum-gas-prices = "0uguu"/' $NODE_HOME/config/app.toml;
$sed_i 's/stake/uguu/g' $NODE_HOME/config/genesis.json;


# PRICEFEED=$(ununifid keys show pricefeed --address)
# for derivativs
jq '.app_state.pricefeed.params.markets = [
  { "market_id": "ubtc:usd", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED_ADDRESS'" ], "active": true },
  { "market_id": "ubtc:usd:30", "base_asset": "ubtc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED_ADDRESS'" ], "active": true },
  { "market_id": "uusdc:usd", "base_asset": "uusdc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED_ADDRESS'" ], "active": true },
  { "market_id": "uusdc:usd:30", "base_asset": "uusdc", "quote_asset": "usd", "oracles": [ "'$PRICEFEED_ADDRESS'" ], "active": true },
  { "market_id": "ubtc:uusdc", "base_asset":"ubtc", "quote_asset":"uusdc", "oracles": ["'$PRICEFEED_ADDRESS'"], "active": true},
  { "market_id": "ubtc:uusdc:30", "base_asset":"ubtc", "quote_asset":"uusdc", "oracles": ["'$PRICEFEED_ADDRESS'"], "active": true}
  ]'  $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.pool_params.quote_ticker = "usd"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.pool_params.base_lpt_mint_fee = "0.001"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.pool_params.base_lpt_redeem_fee = "0.001"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.pool_params.report_liquidation_reward_rate = "0.3"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.pool_params.report_levy_period_reward_rate = "0.3"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.pool_params.accepted_assets_conf = [{"denom":"ubtc", "target_weight": "0.6"}, {"denom":"uusdc", "target_weight":"0.4"}]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.commission_rate = "0.001"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.margin_maintenance_rate = "0.5"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.imaginary_funding_rate_proportional_coefficient = "0.0005"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.derivatives.params.perpetual_futures.markets = [{"base_denom": "ubtc", "quote_denom": "uusdc" }]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.bank.denom_metadata = [
  {"base" : "ubtc" , "symbol": "BTC"},
  {"base" : "uusdc", "symbol": "USDC"}
  ]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.pricefeed.posted_prices = [
  {"expiry": "2024-02-20T12:02:01Z","market_id": "ubtc:usd","oracle_address": "ununifi1h7ulktk5p2gt7tnxwhqzlq0yegq47hum0fahcr","price": "0.024508410211260500"},
  {"expiry": "2024-02-20T12:02:47Z","market_id": "ubtc:usd:30","oracle_address": "ununifi1h7ulktk5p2gt7tnxwhqzlq0yegq47hum0fahcr","price": "0.024508410211260500"},
  {"expiry": "2024-02-20T12:03:30Z","market_id": "uusdc:usd","oracle_address": "ununifi1h7ulktk5p2gt7tnxwhqzlq0yegq47hum0fahcr","price": "0.000001001479651825"},
  {"expiry": "2024-02-20T12:04:11Z","market_id": "uusdc:usd:30","oracle_address": "ununifi1h7ulktk5p2gt7tnxwhqzlq0yegq47hum0fahcr","price": "0.000001002011358752"},
  {"expiry": "2024-02-20T12:00:38Z","market_id": "ubtc:uusdc","oracle_address": "ununifi1h7ulktk5p2gt7tnxwhqzlq0yegq47hum0fahcr","price": "24472.1998760521"},
  {"expiry": "2024-02-20T12:00:38Z","market_id": "ubtc:uusdc:30","oracle_address": "ununifi1h7ulktk5p2gt7tnxwhqzlq0yegq47hum0fahcr","price": "24459.2139572006"}
]'  $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;

# ununifid start --home=$NODE_HOME

# for nftmint
jq '.app_state.nftfactory.class_attributes_list = [
  {
    "base_token_uri": "ipfs://testcid/",
    "class_id": "ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3",
    "minting_permission": "Anyone",
    "owner": "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
    "token_supply_cap": "100000"
  },
  {
    "base_token_uri": "ipfs://testcid/",
    "class_id": "ununifi-D4AC8DBC54261BB1B6ACBBF721A60D131A048F83",
    "minting_permission": "OnlyOwner",
    "owner": "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
    "token_supply_cap": "100000"
  }
]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.nft.classes = [
  {
    "data": null,
    "description": "",
    "id": "ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3",
    "name": "Test",
    "symbol": "",
    "uri": "",
    "uri_hash": ""
  },
  {
    "data": null,
    "description": "",
    "id": "ununifi-D4AC8DBC54261BB1B6ACBBF721A60D131A048F83",
    "name": "Test",
    "symbol": "",
    "uri": "",
    "uri_hash": ""
  }
]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;

# for stakeibc
jq '.app_state.stakeibc.params.deposit_interval = "1"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.delegate_interval = "1"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.rewards_interval = "1"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.redemption_rate_interval = "1"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.ununifi_commission = "1"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.reinvest_interval = "1"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.validator_rebalancing_threshold = "100"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.ica_timeout_nanos = "600000000000"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.buffer_size = "5"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.ibc_timeout_blocks = "300"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.fee_transfer_timeout_nanos = "1800000000000"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.max_stake_ica_calls_per_epoch = "100"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.safety_min_redemption_rate_threshold = "90"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.safety_max_redemption_rate_threshold = "150"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
jq '.app_state.stakeibc.params.ibc_transfer_timeout_nanos = "1800000000000"' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;

# for iya
jq '.app_state.wasm.codes = [
  {
    "code_bytes": "",
    "code_id": "1",
    "code_info": {
      "code_hash": "VfURcnyxjkug9BOEddgx5PiUhefx7qp+G87ihceWjaw=",
      "creator": "'$VAL_ADDRESS_1'",
      "instantiate_config": {
        "address": "",
        "addresses": [],
        "permission": "Everybody"
      }
    },
    "pinned": false
  }
]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
python3 $SCRIPT_DIR/rewrite.py $NODE_HOME/config/genesis.json $SCRIPT_DIR/tmp_code_bytes.txt $NODE_HOME/config/genesis.json

jq '.app_state.wasm.contracts = [
  {
    "contract_address": "'$CONTRACT_ADDRESS1'",
    "contract_code_history": [
      {
        "code_id": "1",
        "msg": {
          "deposit_denom": "uguu",
          "owner": "'$VAL_ADDRESS_1'",
          "unbond_period": 1
        },
        "operation": "CONTRACT_CODE_HISTORY_OPERATION_TYPE_INIT",
        "updated": {
          "block_height": "6",
          "tx_index": "0"
        }
      }
    ],
    "contract_info": {
      "admin": "",
      "code_id": "1",
      "created": {
        "block_height": "6",
        "tx_index": "0"
      },
      "creator": "'$VAL_ADDRESS_1'",
      "extension": null,
      "ibc_port_id": "",
      "label": "BaseStrategy"
    },
    "contract_state": [
      {
        "key": "636F6E666967",
        "value": "eyJvd25lciI6InVudW5pZmkxYThqY3NtbGE2aGV1OTlsZHRhemMyN2RuYTRxY2Q0anlnc3RoeDYiLCJ1bmJvbmRfcGVyaW9kIjoxLCJkZXBvc2l0X2Rlbm9tIjoidWd1dSIsInJlZGVtcHRpb25fcmF0ZSI6IjEwMDAwMDAiLCJ0b3RhbF9kZXBvc2l0IjoiMCJ9"
      }
    ]
  }

]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;


jq '.app_state.wasm.sequences = [
  {
    "id_key": "BGxhc3RDb2RlSWQ=",
    "value": "2"
  },
  {
    "id_key": "BGxhc3RDb250cmFjdElk",
    "value": "2"
  }
]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;

jq ' .app_state.yieldaggregator.strategies = [
  {
    "contract_address": "'$CONTRACT_ADDRESS1'",
    "denom": "uguu",
    "git_url": "",
    "id": "0",
    "name": "Contract Staking"
  }
]' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;
# for disable bank
# jq '.app_state.bank.params.default_send_enabled = false' $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;

# jq ".app_state.gov.params.voting_period = \"20s\"" $NODE_HOME/config/genesis.json > temp.json ; mv temp.json $NODE_HOME/config/genesis.json;

