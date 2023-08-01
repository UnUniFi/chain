#!/bin/sh

# block speed
sleep=5

# mint nft
echo "------------mint nft------------"
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------check nft------------"
ununifid q nft nft ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01

# list nft
echo "------------list nft------------"
ununifid tx nftbackedloan list \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user1 --keyring-backend test --chain-id test --yes

sleep $sleep

echo "------------check nft------------"
ununifid q nft nft ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01
