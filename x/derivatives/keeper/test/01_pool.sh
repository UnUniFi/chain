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

user1_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')
user1_ubtc_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom ubtc -o json | jq .amount | tr -d '"')

expected_user1_udlp_balance=$(($init_udlp_balance + 166388890))
expected_user1_ubtc_balance=$(($init_ubtc_balance - 100000000))

if [ "$user1_udlp_balance" = "$expected_user1_udlp_balance" ]; then
  echo "pass: udlp balance is correct: $user1_udlp_balance"
else
  echo "error: udlp balance is incorrect:"
  echo "expected: $expected_user1_udlp_balance actual: $user1_udlp_balance"
fi

if [ "$user1_ubtc_balance" = "$expected_user1_ubtc_balance" ]; then
  echo "pass: ubtc balance is correct: $user1_ubtc_balance"
else
  echo "error: ubtc balance is incorrect:"
  echo "expected: $expected_user1_ubtc_balance actual: $user1_ubtc_balance"
fi

echo "------------deposit to pool 2nd------------"
ununifid tx derivatives deposit-to-pool 100000000ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

user1_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')
expected_user1_udlp_balance=$(($expected_user1_udlp_balance + 83194445))

if [ "$user1_udlp_balance" = "$expected_user1_udlp_balance" ]; then
  echo "pass: udlp balance is correct: $user1_udlp_balance"
else
  echo "error: udlp balance is incorrect:"
  echo "expected: $expected_user1_udlp_balance actual: $user1_udlp_balance"
fi

echo "------------withdraw from pool------------"
ununifid tx derivatives withdraw-from-pool 249583334 ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

user1_ubtc_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom ubtc -o json | jq .amount | tr -d '"')

if [ "$user1_ubtc_balance" > "$init_ubtc_balance" ]; then
  echo "pass: ubtc balance is correct: $user1_ubtc_balance"
else
  echo "error: ubtc balance is incorrect:"
  echo "initial: $init_ubtc_balance actual: $user1_ubtc_balance"
fi
