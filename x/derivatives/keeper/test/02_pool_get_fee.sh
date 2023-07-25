#!/bin/sh

# block speed
sleep=5

user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w
user2_address=ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla
init_ubtc_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom ubtc -o json | jq .amount | tr -d '"')
init_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')

echo "------------deposit to pool 1st------------"
ununifid tx derivatives deposit-to-pool 100000000ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------open position------------"
ununifid tx derivatives open-position perpetual-futures 10000000ubtc ubtc uusdc long 40 4 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------price change------------"
ununifid tx pricefeed postprice ubtc:usd 0.022508410211260500 1200 \
--from=pricefeed --keyring-backend test --chain-id test --yes

sleep $sleep

ununifid tx derivatives report-liquidation 1 $user1_address \
--from user1 --keyring-backend test --chain-id test --yes