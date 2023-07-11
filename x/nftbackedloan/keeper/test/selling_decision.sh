#!/bin/sh

# block speed
sleep=5

# mint nft
echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# list nft
echo "------------list nft------------"
ununifid tx nftbackedloan listing \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# bid nft from user2
echo "------------bid nft from user2------------"
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
200uguu 50uguu 0.1 7200 --from user2 --keyring-backend test --chain-id test --yes

sleep $sleep
# selling decision
echo "------------selling decision------------"
ununifid tx nftbackedloan selling-decision \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01

# pay full bid
echo "------------pay full bid------------"
ununifid tx nftbackedloan pay-full-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01

# Check nft transfer & balance
balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu)
amount=$(echo "$balance" | awk -F': ' '/amount/ { print $2 }' | tr -d '"')
# -200
expected_amount="99999999800"

if [ "$amount" = "$expected_amount" ]; then
  echo "pass: bidder amount is correct: $amount"
else
  echo "error: bidder amount is incorrect: $amount"
fi

sleep $sleep
# NftListingFullPaymentPeriod: 30
sleep 30

ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01
# NftListingNftDeliveryPeriod: 30
sleep 30

balance=$(ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w --denom uguu)
amount=$(echo "$balance" | awk -F': ' '/amount/ { print $2 }' | tr -d '"')
# + 200 - 10 (5% fee)
expected_amount="100000000190"

if [ "$amount" = "$expected_amount" ]; then
  echo "pass: owner amount is correct: $amount"
else
  echo "error: owner amount is incorrect: $amount"
fi

response=$(ununifid q nft owner ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01)
owner=$(echo "$response" | awk -F': ' '/owner/ { print $2 }' | tr -d '"')
# user2
expected_owner="ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla"

if [ "$owner" = "$expected_owner" ]; then
  echo "pass: Owner is changed to $owner"
else
  echo "error: Owner is not changed from $owner"
fi
