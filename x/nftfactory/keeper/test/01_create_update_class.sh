#!/bin/sh

# block speed
sleep=5

user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w

echo "------------create class id------------"
ununifid tx nftfactory create-class test \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

class_id=$(ununifid q nftfactory classes-from-creator $user1_address -o json | jq .classes[0] | tr -d '"')
expected_class_id=factory/$user1_address/test

echo "------------check class id------------"
if [ "$expected_class_id" = "$class_id" ]; then
  echo "pass: class_id is correct: $class_id"
else
  echo "error: class_id is incorrect:"
  echo "expected: $expected_class_id actual: $class_id"
fi

echo "------------update class------------"
ununifid tx nftfactory update-class $class_id \
--name "updated" --symbol "OK" --description "after class changed" --uri "change url" --uri-hash "changed" \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

class_name=$(ununifid q nft class $class_id -o json | jq .class.name | tr -d '"')
expected_class_name=updated

echo "------------check class name------------"
if [ "$expected_class_name" = "$class_name" ]; then
  echo "pass: class_name is correct: $class_name"
else
  echo "error: class_name is incorrect:"
  echo "expected: $expected_class_name actual: $class_name"
fi
