#!/bin/sh

# block speed
sleep=5

user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w
user2_address=ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla
init_user2_ubtc_balance=$(ununifid q bank balances $user2_address --denom ubtc -o json | jq .amount | tr -d '"')

echo "------------deposit to pool 1st------------"
ununifid tx derivatives deposit-to-pool 100000000ubtc \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

user1_udlp_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom udlp -o json | jq .amount | tr -d '"')

ununifid tx pricefeed postprice ubtc:usd 0.024508410211260500 1200 \
--from=pricefeed --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------open position------------"
ununifid tx derivatives open-position perpetual-futures 10000000ubtc ubtc uusdc long 40 4 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------position opened user------------"
ununifid q derivatives positions $user2_address

echo "------------nft listing------------"
ununifid tx nftbackedloan list derivatives/perpetual_futures/positions 1 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------check nft send disabled------------"
ununifid q nft nft derivatives/perpetual_futures/positions 1

echo "------------price change------------"
ununifid tx pricefeed postprice ubtc:usd 0.020508410211260500 1200 \
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

echo "------------ubtc balance check------------"
user2_ubtc_balance=$(ununifid q bank balances $user2_address --denom ubtc -o json | jq .amount | tr -d '"')
expected_user2_ubtc_balance=$(($init_user2_ubtc_balance - 10000000))

if [ "$user2_ubtc_balance" = "$expected_user2_ubtc_balance" ]; then
  echo "pass: trader ubtc balance is correct: $user2_ubtc_balance"
else
  echo "error: trader ubtc balance is incorrect"
  echo "expected: $expected_user2_ubtc_balance actual: $user2_ubtc_balance"
fi

echo "------------nft cancel listing------------"
ununifid tx nftbackedloan cancel-listing derivatives/perpetual_futures/positions 1 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------check nft send enabled------------"
ununifid q nft nft derivatives/perpetual_futures/positions 1

echo "------------close position------------"
ununifid tx derivatives close-position 1 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

user2_closed_ubtc_balance=$(ununifid q bank balances $user2_address --denom ubtc -o json | jq .amount | tr -d '"')

echo "------------ubtc balance check------------"
if [ "$user2_closed_ubtc_balance" -gt "$user2_ubtc_balance" ]; then
  echo "pass: trader ubtc balance is correct: $user2_closed_ubtc_balance"
else
  echo "error: trader ubtc balance is incorrect"
  echo "before closed: $user2_ubtc_balance after closed: $user2_closed_ubtc_balance"
fi

echo "------------check nft not exist------------"
ununifid q nft nft derivatives/perpetual_futures/positions 1
