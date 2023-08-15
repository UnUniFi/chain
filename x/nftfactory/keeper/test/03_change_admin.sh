#!/bin/sh

# block speed
sleep=5

user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w
user2_address=ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla

echo "------------create class------------"
ununifid tx nftfactory create-class test \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

class_id=$(ununifid q nftfactory classes-from-creator $user1_address -o json | jq .classes[0] | tr -d '"')

echo "------------change admin------------"
ununifid tx nftfactory change-admin $class_id $user2_address \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

admin=$(ununifid q nftfactory class-authority-metadata $class_id -o json | jq .authority_metadata.Admin | tr -d '"')

echo "------------check admin------------"
if [ "$user2_address" = "$admin" ]; then
  echo "pass: admin is correct: $admin"
else
  echo "error: admin is incorrect:"
  echo "expected: $user2_address actual: $admin"
fi

echo "------------mint nft------------"
ununifid tx nftfactory mint-nft $class_id test02 $user2_address \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

owner=$(ununifid q nft owner $class_id test02 -o json | jq .owner | tr -d '"')

echo "------------check owner------------"
if [ "$user2_address" = "$owner" ]; then
  echo "pass: owner is correct: $owner"
else
  echo "error: owner is incorrect:"
  echo "expected: $user2_address actual: $owner"
fi
