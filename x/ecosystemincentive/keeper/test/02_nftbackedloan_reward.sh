#!/bin/sh

# block speed
sleep=5
class_id=test
token_id=a01
user1_address=ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w
user2_address=ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla

echo "------------create class------------"
ununifid tx nftfactory create-class $class_id \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

class_id=$(ununifid q nftfactory classes-from-creator $user1_address -o json | jq .classes[0] | tr -d '"')

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

echo "------------bid nft from user2------------"
ununifid tx nftbackedloan place-bid \
$class_id $token_id \
200uguu 50uguu 0.1 7200 --automatic-payment=false --from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------selling decision------------"
ununifid tx nftbackedloan selling-decision \
$class_id $token_id \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------pay remainder------------"
ununifid tx nftbackedloan pay-remainder \
$class_id $token_id \
--from user2 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "============check nft status selling_decision ============"
ununifid q nftbackedloan nft-listing $class_id $token_id

echo "wait.......... NftListingFullPaymentPeriod: 30s"
sleep 30

echo "============check nft status successful_bid ============"
ununifid q nftbackedloan nft-listing $class_id $token_id

echo "wait.......... NftListingNftDeliveryPeriod: 30s"
sleep 30

reward=$(ununifid q ecosystem-incentive all-rewards $user1_address -o json | jq .reward_record.rewards[0].amount | tr -d '"')
denom=$(ununifid q ecosystem-incentive all-rewards $user1_address -o json | jq .reward_record.rewards[0].denom | tr -d '"')

if [ "$reward" -gt "0" ]; then
  echo "pass: reward exist: $reward $denom"
else
  echo "error: reward not exist: $reward"
fi
