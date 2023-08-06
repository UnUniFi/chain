#!/bin/sh

# block speed
sleep=5

user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w
user2_address=ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla
user3_address=ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6
init_ubtc_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom ubtc -o json | jq .amount | tr -d '"')
init_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')
init_user2_ubtc_balance=$(ununifid q bank balances $user2_address --denom ubtc -o json | jq .amount | tr -d '"')
init_user3_ubtc_balance=$(ununifid q bank balances $user3_address --denom ubtc -o json | jq .amount | tr -d '"')

echo "------------deposit to pool 1st------------"
ununifid tx derivatives deposit-to-pool 100000000ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

user1_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')

echo "------------open position------------"
ununifid tx derivatives open-position perpetual-futures 10000000ubtc ubtc uusdc long 40 4 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------position opened user------------"
ununifid q derivatives positions $user2_address

echo "------------nft send------------"
ununifid tx nft send derivatives/perpetual_futures/positions 1 $user3_address \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------position nft owner------------"
ununifid q derivatives positions $user3_address

sleep $sleep

echo "------------close position------------"
ununifid tx derivatives close-position 1 \
--from user3 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------closed position nft owner------------"
ununifid q derivatives positions $user3_address

echo "------------withdraw from pool------------"
ununifid tx derivatives withdraw-from-pool $user1_udlp_balance ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------ubtc balance check------------"
user1_ubtc_balance=$(ununifid q bank balances $user1_address --denom ubtc -o json | jq .amount | tr -d '"')
user2_ubtc_balance=$(ununifid q bank balances $user2_address --denom ubtc -o json | jq .amount | tr -d '"')
user3_ubtc_balance=$(ununifid q bank balances $user3_address --denom ubtc -o json | jq .amount | tr -d '"')
user1_profit=$(($user1_ubtc_balance - $init_ubtc_balance))
user2_loss=$(($init_user2_ubtc_balance - $user2_ubtc_balance))
user3_profit=$(($user3_ubtc_balance - $init_ubtc_balance))

if [ "$user1_ubtc_balance" -gt "$init_ubtc_balance" ]; then
  echo "pass: depositer ubtc balance is correct: $user1_ubtc_balance"
  echo "loss $user1_loss"
else
  echo "error: depositor ubtc balance is incorrect"
  echo "initial: $init_ubtc_balance actual: $user1_ubtc_balance"
fi

if [ "$user2_ubtc_balance" -lt "$init_user2_ubtc_balance" ]; then
  echo "pass: trader ubtc balance is correct: $user2_ubtc_balance"
  echo "profit $user2_profit"
else
  echo "error: trader ubtc balance is incorrect"
  echo "initial: $init_ubtc_balance actual: $user2_ubtc_balance"
fi

if [ "$user3_ubtc_balance" -gt "$init_user3_ubtc_balance" ]; then
  echo "pass: nft owner ubtc balance is correct: $user3_ubtc_balance"
  echo "profit $user3_profit"
else
  echo "error: nft owner ubtc balance is incorrect"
  echo "initial: $init_ubtc_balance actual: $user3_ubtc_balance"
fi