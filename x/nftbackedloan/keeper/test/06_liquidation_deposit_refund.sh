#!/bin/sh

# block speed
sleep=5

init_user1_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu -o json | jq .amount | tr -d '"')
init_user2_balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu -o json | jq .amount | tr -d '"')
init_user3_balance=$(ununifid q bank balances ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6 --denom uguu -o json | jq .amount | tr -d '"')

# mint nft
echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# list nft
echo "------------list nft------------"
ununifid tx nftbackedloan list \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06 \
--bid-token uguu --min-deposit-rate 0.1 --min-bidding-period-hours 0.005 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# bid nft from user2 (auto payment)
echo "------------bid nft from user2------------"
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06 \
200uguu 40uguu 0.1 36 --from user2 --keyring-backend test --chain-id test --yes

sleep $sleep
# bid nft from user3 (auto payment)
echo "------------bid nft from user3------------"
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06 \
140uguu 50uguu 0.2 30 --from user3 --keyring-backend test --chain-id test --yes

sleep $sleep

# borrow
echo "------------borrow------------"
ununifid tx nftbackedloan borrow \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06 \
ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6 60uguu \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
ununifid q nftbackedloan nft-bids ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06

echo "wait.......... Bid expire"
sleep 30

echo "============check nft status liquidation ============"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06
ununifid q nftbackedloan nft-bids ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06

# not pay remainder
echo "wait.......... NftListingFullPaymentPeriod: 30s"
sleep 30

echo "============check nft status successful_bid ============"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06

echo "wait.......... NftListingNftDeliveryPeriod: 30s"
sleep 30

# check refund bidder balance
user2_balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu -o json | jq .amount | tr -d '"')
# +0 (interest)
expected_user2_balance=$((init_user2_balance + 0))

if [ "$user2_balance" = "$expected_user2_balance" ]; then
  echo "pass: refund bidder balance is correct: $user2_balance"
else
  echo "error: refund bidder balance is incorrect"
  echo "expected: $expected_user2_balance actual: $user2_balance
fi

# check win bidder balance
user3_balance=$(ununifid q bank balances ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6 --denom uguu -o json | jq .amount | tr -d '"')
# -140 (bid price)
expected_user3_balance=$((init_user3_balance - 140))

if [ "$user3_balance" = "$expected_user3_balance" ]; then
  echo "pass: win bidder balance is correct: $user3_balance"
else
  echo "error: win bidder balance is incorrect"
  echo "expected: $expected_user3_balance actual: $user3_balance
fi

user1_balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu -o json | jq .amount | tr -d '"')
# + 50 (borrow) + 140 - 50 - 4(fee) = 136
expected_user1_balance=$((init_user1_balance + 136))

if [ "$user1_balance" = "$expected_user1_balance" ]; then
  echo "pass: owner balance is correct: $user1_balance"
else
  echo "error: owner balance is incorrect"
  echo "expected: $expected_user1_balance actual: $user1_balance"
fi

owner=$(ununifid q nft owner ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a06 -o json | jq .owner | tr -d '"')
# user3
expected_owner="ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6"

if [ "$owner" = "$expected_owner" ]; then
  echo "pass: Owner is changed to $owner"
else
  echo "error: Owner is not changed from $owner"
  echo "expected: $expected_owner actual: $owner"
fi
