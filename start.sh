


#!/bin/sh

rm -rf ~/.ununifi

# Build genesis
ununifid init --chain-id=test test
ununifid keys add validator --keyring-backend="test"
ununifid add-genesis-account $(ununifid keys show validator -a --keyring-backend="test") 100000000000000uguu,100000000000000stake
ununifid gentx validator 100000000stake --keyring-backend="test" --chain-id=test
ununifid collect-gentxs

# Start node
ununifid start --pruning=nothing
