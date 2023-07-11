#!/bin/sh

# block speed
sleep=5

# mint nft
echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# list nft
echo "------------list nft------------"
ununifid tx nftbackedloan list \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03 \
--bid-token uguu --min-deposit-rate 0.1 --min-bidding-period-hours 0.005 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# bid nft from user2
echo "------------bid nft from user2------------"
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03 \
200uguu 50uguu 0.1 20 --from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

# borrow
echo "------------borrow------------"
ununifid tx nftbackedloan borrow \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03 \
ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla 40uguu \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "wait.......... Bid expire"
sleep 20

echo "============check nft status liquidation ============"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03

# pay remainder
echo "------------pay remainder------------"
ununifid tx nftbackedloan pay-remainder \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep
echo "wait.......... NftListingFullPaymentPeriod: 30s"
sleep 30

echo "============check nft status successful_bid ============"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03

echo "wait.......... NftListingNftDeliveryPeriod: 30s"
sleep 30

# check bidder balance
amount=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu -o json | jq .amount | tr -d '"')
# -200
expected_amount="99999999800"

if [ "$amount" = "$expected_amount" ]; then
  echo "pass: bidder amount is correct: $amount"
else
  echo "error: bidder amount is incorrect: $amount"
fi

amount=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu -o json | jq .amount | tr -d '"')
# + 40 (borrow) + 160 - 8(fee) = 192
expected_amount="100000000192"

if [ "$amount" = "$expected_amount" ]; then
  echo "pass: owner amount is correct: $amount"
else
  echo "error: owner amount is incorrect: $amount"
fi

owner=$(ununifid q nft owner ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a03 -o json | jq .owner | tr -d '"')
# user2
expected_owner="ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla"

if [ "$owner" = "$expected_owner" ]; then
  echo "pass: Owner is changed to $owner"
else
  echo "error: Owner is not changed from $owner"
fi
