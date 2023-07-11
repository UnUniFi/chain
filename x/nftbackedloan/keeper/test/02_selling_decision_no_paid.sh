#!/bin/sh

# block speed
sleep=5

# mint nft
echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a02 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# list nft
echo "------------list nft------------"
ununifid tx nftbackedloan listing \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a02 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep
# bid nft from user2
echo "------------bid nft from user2------------"
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a02 \
200uguu 50uguu 0.1 7200 --from user2 --keyring-backend test --chain-id test --yes

sleep $sleep
# selling decision
echo "------------selling decision------------"
ununifid tx nftbackedloan selling-decision \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a02 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------check nft status selling_decision ------------"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a02

# no pay full bid

# NftListingFullPaymentPeriod: 30
sleep 30

echo "------------check nft status listing ------------"
ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a02

# check bidder balance
balance=$(ununifid q bank balances ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla --denom uguu)
amount=$(echo "$balance" | awk -F': ' '/amount/ { print $2 }' | tr -d '"')
# -50
expected_amount="99999999950"

if [ "$amount" = "$expected_amount" ]; then
  echo "pass: bidder amount is correct: $amount"
else
  echo "error: bidder amount is incorrect: $amount"
fi

# check collected amount
response=$(ununifid q nftbackedloan nft-listing ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a02 -o json)
amount=$(echo "$response" | jq .listing.collected_amount.amount | tr -d '"')
negative=$(echo "$response" | jq .listing.collected_amount_negative | tr -d '"')
expected_amount="50"
if [ "$amount" = "$expected_amount" ]; then
  echo "pass: collect amount is correct: $amount"
else
  echo "error: collect amount is incorrect: $amount"
fi

if [ "$negative" = "false" ]; then
  echo "pass: negative bool is correct: $negative"
else
  echo "error: negative bool is incorrect: $negative"
fi
