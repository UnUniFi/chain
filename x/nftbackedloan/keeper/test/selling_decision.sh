#!/bin/sh

# mint nft
ununifid tx nftfactory mint-nft \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w \
--from user1 --keyring-backend test --chain-id test --yes
# list nft
ununifid tx nftbackedloan listing \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user1 --keyring-backend test --chain-id test --yes
# bid nft from user2
ununifid tx nftbackedloan place-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
200uguu 50uguu 0.1 7200 --from user2 --keyring-backend test --chain-id test --yes

# selling decision
ununifid tx nftbackedloan selling-decision \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user1 --keyring-backend test --chain-id test --yes

ununifid tx nftbackedloan pay-full-bid \
ununifi-1AFC3C85B52311F13161F724B284EF900458E3B3 a01 \
--from user2 --keyring-backend test --chain-id test --yes

ununifid q nftbackedloan listed-nfts
ununifid q bank balances ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w