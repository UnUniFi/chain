#!/bin/sh

## before run this script, change params scripts/setup/init.sh
## Line 88 levy_period_required_seconds under block speed

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

sleep $sleep

echo "------------levy tx------------"
ununifid tx derivatives report-levy-period 1 $user1_address \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------levied position------------"
ununifid q derivatives positions $user2_address

echo "------------udlp rate check------------"
rate=$(ununifid q derivatives delp-token-rate -o json | jq .rates[0].amount | tr -d '"')

if [ "$rate" -lt "1000000" ]; then
  echo "pass: delp token rate is correct: $rate"
else
  echo "error: delp token rate is incorrect:"
  echo "expected: 0 actual: $rate"
fi

echo "------------pool amount check------------"
pool_amount=$(ununifid q derivatives pool -o json | jq .pool_market_cap.asset_info[0].amount | tr -d '"')

if [ "$pool_amount" -lt "100000000" ]; then
  echo "pass: pool amount is correct: $pool_amount"
else
  echo "error: pool market cap is incorrect:"
  echo "expected: lower 100000000, actual: $pool_amount"
fi
