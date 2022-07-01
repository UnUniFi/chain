#!/usr/bin/env bash
rm -rf ~/.ununifi;
~/go/bin/ununifid init alpha-node1 --chain-id ununifi-test-private-m1;

sed -i '/\[api\]/,+3 s/enable = false/enable = true/' ~/.ununifi/config/app.toml;
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0uguu"/' ~/.ununifi/config/app.toml;
sed -i 's/mode = "full"/mode = "validator"/' ~/.ununifi/config/config.toml;
sed -i 's/stake/uguu/g' ~/.ununifi/config/genesis.json;
ununifid keys add my_validator --keyring-backend test
ununifid keys add my_receiver --keyring-backend test
ununifid keys add faucet --keyring-backend test
ununifid add-genesis-account my_validator 100000000000uguu,100000000000ubtc;
ununifid add-genesis-account my_receiver 100000000000uguu,100000000000ubtc;
ununifid add-genesis-account faucet 100000000000uguu,100000000000ubtc,100000000000jpu;
ununifid gentx my_validator 100000000uguu --chain-id ununifi-test-private-m1 --keyring-backend test;
ununifid collect-gentxs;

# query nft
ununifid query nft nft a10 a10
ununifid query nft owner a10 a10
ununifid query nftmarket listed_nfts

# mint an nft
ununifid tx nftmarket mint a10 a10 uri 888838 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# listing nft
ununifid tx nftmarket listing a10 a10 --chain-id  ununifi-test-private-m1 --from my_validator --keyring-backend test --gas 300000 -y |jq .
ununifid tx nftmarket listing a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# cancel listing
ununifid tx nftmarket cancel_listing a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# end listing 
ununifid tx nftmarket endlisting a10 a10 --chain-id  ununifi-test-private-m1 --from my_validator --keyring-backend test --gas 300000 -y|jq .;
ununifid tx nftmarket endlisting a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# expand nft listing
ununifid tx nftmarket expand_nft_listing a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# selling decision
ununifid tx nftmarket selling_decision a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# place bid
ununifid tx nftmarket placebid a10 a10 100uguu --chain-id  ununifi-test-private-m1 --from my_receiver --keyring-backend test --gas 300000 -y|jq .
ununifid tx nftmarket placebid a10 a10 100uguu --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# cancel bid
ununifid tx nftmarket cancelbid a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# pay full bid
ununifid tx nftmarket pay_fullbid a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# borrow uguu
ununifid tx nftmarket borrow a10 a10 1uguu --chain-id  ununifi-test-private-m1 --from my_validator --keyring-backend test --gas 300000 -y|jq .
# repay uguu
ununifid tx nftmarket repay a10 a10 1uguu --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# mint stable coin
ununifid tx nftmarket mint_stablecoin a10 a10 100usd --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# burn stable coin
ununifid tx nftmarket burn_stablecoin a10 a10 100usd --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# liquidate position
ununifid tx nftmarket liquidate a10 a10 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
