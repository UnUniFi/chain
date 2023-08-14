#!/bin/sh

# block speed
sleep=5

user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w

echo "------------create class------------"
ununifid tx nftfactory create-class test \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

class_id=$(ununifid q nftfactory classes-from-creator $user1_address -o json | jq .classes[0] | tr -d '"')

echo "------------mint nft------------"
ununifid tx nftfactory mint-nft $class_id test01 $user1_address \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

owner=$(ununifid q nft owner $class_id test01 -o json | jq .owner | tr -d '"')

echo "------------check owner------------"
if [ "$user1_address" = "$owner" ]; then
  echo "pass: owner is correct: $owner"
else
  echo "error: owner is incorrect:"
  echo "expected: $user1_address actual: $owner"
fi

echo "------------not owner cannot burn nft------------"
ununifid tx nftfactory burn-nft $class_id test01 \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep
owner=$(ununifid q nft owner $class_id test01 -o json | jq .owner | tr -d '"')

echo "------------check owner------------"
if [ "$user1_address" = "$owner" ]; then
  echo "pass: owner is correct: $owner"
else
  echo "error: owner is incorrect:"
  echo "expected: $user1_address actual: $owner"
fi


echo "------------burn nft------------"
ununifid tx nftfactory burn-nft $class_id test01 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

owner=$(ununifid q nft owner $class_id test01 -o json | jq .owner | tr -d '"')

echo "------------check owner------------"
if [ "" = "$owner" ]; then
  echo "pass: owner is correct: $owner"
else
  echo "error: owner is incorrect:"
  echo "expected: "" actual: $owner"
fi
