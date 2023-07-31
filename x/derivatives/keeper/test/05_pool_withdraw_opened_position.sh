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

echo "------------withdraw from pool (error)------------"
ununifid tx derivatives withdraw-from-pool $user1_udlp_balance ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------ubtc balance check------------"
user1_udlp_balance_after=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')

if [ "$user1_udlp_balance_after" = "$user1_udlp_balance" ]; then
  echo "pass: depositer udlp balance is correct: $user1_udlp_balance_after"
else
  echo "error: depositer udlp balance is incorrect"
  echo "expexted: $user1_udlp_balance actual: $user1_udlp_balance_after"
fi
