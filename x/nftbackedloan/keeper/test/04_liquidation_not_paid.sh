#!/bin/sh

# block speed
sleep=5

init_user1_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu -o json | jq .amount | tr -d '"')
init_user2_balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu -o json | jq .amount | tr -d '"')

# mint nft
echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a04 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# list nft
echo "------------list nft------------"
ununifid tx nftbackedloan list \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a04 \
--bid-token uguu --min-deposit-rate 0.1 --min-bidding-period-hours 0.005 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# bid nft from user2
echo "------------bid nft from user2------------"
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a04 \
200uguu 50uguu 0.1 30 --automatic-payment=false --from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

# borrow
echo "------------borrow------------"
ununifid tx nftbackedloan borrow \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a04 \
ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla 40uguu \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "wait.......... Bid expire"
sleep 30

echo "============check nft status liquidation ============"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a04

# not pay remainder
echo "wait.......... NftListingFullPaymentPeriod: 30s"
sleep 30

# check bidder balance
user2_balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu -o json | jq .amount | tr -d '"')
# -50
expected_user2_balance=$((init_user2_balance - 50))

if [ "$user2_balance" = "$expected_user2_balance" ]; then
  echo "pass: bidder balance is correct: $user2_balance"
else
  echo "error: bidder balance is incorrect"
  echo "expected: $expected_user2_balance actual: $user2_balance
fi

user1_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu -o json | jq .amount | tr -d '"')
# + 40 (borrow) + 50 - 40 = 50
expected_user1_balance=$((init_user1_balance + 50))

if [ "$user1_balance" = "$expected_user1_balance" ]; then
  echo "pass: owner balance is correct: $user1_balance"
else
  echo "error: owner balance is incorrect"
  echo "expected: $expected_user1_balance actual: $user1_balance"
fi

owner=$(ununifid q nft owner ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a04 -o json | jq .owner | tr -d '"')
# user1
expected_owner="ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w"

if [ "$owner" = "$expected_owner" ]; then
  echo "pass: Owner is changed to $owner"
else
  echo "error: Owner is not changed from $owner"
  echo "expected: $expected_owner actual: $owner"
fi
