#!/bin/sh

# block speed
sleep=5

user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w
user2_address=ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla
init_ubtc_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom ubtc -o json | jq .amount | tr -d '"')
init_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')
init_user2_ubtc_balance=$(ununifid q bank balances $user2_address --denom ubtc -o json | jq .amount | tr -d '"')

echo "------------deposit to pool 1st------------"
ununifid tx derivatives deposit-to-pool 100000000ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

user1_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')

echo "------------open position------------"
ununifid tx derivatives open-position perpetual-futures 10000000ubtc ubtc uusdc long 40 4 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------opened position------------"
ununifid q derivatives positions $user2_address

echo "------------price change------------"
ununifid tx pricefeed postprice ubtc:usd 0.012508410211260500 1200 \
--from=pricefeed --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------report liquidation------------"
ununifid tx derivatives report-liquidation 1 $user1_address \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------liquidated position------------"
ununifid q derivatives positions $user2_address

echo "------------withdraw from pool------------"
ununifid tx derivatives withdraw-from-pool $user1_udlp_balance ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

user1_ubtc_balance=$(ununifid q bank balances $user1_address --denom ubtc -o json | jq .amount | tr -d '"')
user2_ubtc_balance=$(ununifid q bank balances $user2_address --denom ubtc -o json | jq .amount | tr -d '"')
user1_profit=$(($user1_ubtc_balance - $init_ubtc_balance))
user2_loss=$(($init_user2_ubtc_balance - $user2_ubtc_balance))

if [ "$user1_ubtc_balance" -gt "$init_ubtc_balance" ]; then
  echo "pass: ubtc balance is correct: $user1_ubtc_balance"
  echo "profit $user1_profit trader's loss $user2_loss"
else
  echo "error: ubtc balance is incorrect:"
  echo "initial: $init_ubtc_balance actual: $user1_ubtc_balance"
fi