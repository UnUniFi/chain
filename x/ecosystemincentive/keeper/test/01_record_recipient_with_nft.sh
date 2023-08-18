#!/bin/sh

# block speed
sleep=5
class_id=test
token_id=a01
user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w

echo "------------create class------------"
ununifid tx nftfactory create-class $class_id \
--from user1 --keyring-backend test --chain-id test --yes

class_id=$(ununifid q nftfactory classes-from-creator $user1_address -o json | jq .classes[0] | tr -d '"')

sleep $sleep

echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
$class_id $token_id $user1_address \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------list nft------------"
ununifid tx nftbackedloan list \
$class_id $token_id \
--note "{\"frontend\":{\"version\": 1, \"recipient\": \"$user1_address\"}}" \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------check recipient registered------------"
recipient_address=$(ununifid q ecosystem-incentive recipient-with-nft $class_id $token_id -o json | jq .address | tr -d '"')

if [ "$user1_address" = "$recipient_address" ]; then
  echo "pass: registered recipient is correct: $recipient_address"
else
  echo "error: registered recipient is incorrect:"
  echo "initial: $user1_address actual: $recipient_address"
fi