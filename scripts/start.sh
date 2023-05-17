#!/bin/sh

rm -rf ~/.ununifi

set -o errexit -o nounset

# Build genesis
ununifid init --chain-id=test test
ununifid keys mnemonic >& validator.txt
ununifid keys mnemonic >& debug.txt
ununifid keys add validator --recover < validator.txt
ununifid keys add debug --recover < debug.txt
ununifid add-genesis-account $(ununifid keys show validator --address) 100000000000000uguu,100000000000000stake
ununifid add-genesis-account $(ununifid keys show debug --address) 100000000000000uguu,100000000000000stake
ununifid gentx validator 100000000stake --chain-id=test
ununifid collect-gentxs

# Edit app.toml to enable unsafe-cors and set pruning everything to reduce disk usage.
sed -E -i "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/" ~/.ununifi/config/app.toml;
sed -E -i "s/pruning = \".*\"/pruning = \"everything\"/" ~/.ununifi/config/app.toml;

# Start node
ununifid start --pruning=nothing --minimum-gas-prices=0stake
