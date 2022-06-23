#!/bin/sh

rm -rf ~/.ununifi

set -o errexit -o nounset

# Build genesis file incl account for passed address
ununifid init --chain-id test test
ununifid keys add validator --keyring-backend="test"
ununifid add-genesis-account $(ununifid keys show validator -a --keyring-backend="test") 100000000000000uguu,100000000000stake
ununifid gentx validator 100000000stake --keyring-backend="test" --chain-id test
ununifid collect-gentxs

# Start bitsong
ununifid start --pruning=nothing --mode=validator --minimum-gas-prices=0stake
