#!/usr/bin/env bash

# queries
ununifid query yieldaggregator all-asset-management-accounts
ununifid query yieldaggregator all-farming-units
ununifid query yieldaggregator asset-management-account OsmosisFarm
ununifid query yieldaggregator params
ununifid query yieldaggregator user-info $(ununifid keys show -a validator --keyring-backend=test)
ununifid query yieldaggregator daily-reward-percents

# farming order txs
ununifid tx yieldaggregator add-farming-order --farming-order-id="order1" --strategy-type="ManualStrategy" --whitelisted-target-ids="OsmosisFarmTarget1" --blacklisted-target-ids="" --max-unbonding-seconds=10 --overall-ratio=10 --min=1 --max=10 --date=1661967198 --active=true --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator inactivate-farming-order order1  --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator activate-farming-order order1 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator delete-farming-order order1 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator execute-farming-orders order1 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator set-daily-reward-percent OsmosisFarm OsmosisGUUFarm 0.1 1662429412 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# deposit/withdraw txs
ununifid tx yieldaggregator deposit 1000000uguu --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator begin-withdraw-all --chain-id=test --from=validator --keyring-backend=test --gas=500000 -y --broadcast-mode=block
ununifid tx yieldaggregator withdraw 1000000uguu --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# proposal txs
ununifid tx yieldaggregator proposal-add-yieldfarm --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --assetmanagement-account-name="Osmosis Farm" --enabled=true --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator proposal-remove-yieldfarm --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator proposal-stop-yieldfarm --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator proposal-update-yieldfarm --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --assetmanagement-account-name="Osmosis Farm" --enabled=true --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

ununifid query gov proposals
ununifid tx gov vote 1 yes  --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# cosmwasm yieldfarm target
CONTRACT=ununifi14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sm5z28e
ununifid tx yieldaggregator proposal-add-yieldfarmtarget --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --assetmanagement-account-address="$CONTRACT" --assetmanagement-target-id="OsmosisGUUFarm" --unbonding-seconds=10 --asset-conditions="uguu:1:2,stake:10:2" --integration-type="COSMWASM" --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

ununifid tx yieldaggregator proposal-add-yieldfarmtarget --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --assetmanagement-account-address="" --assetmanagement-target-id="OsmosisGUUFarm" --unbonding-seconds=10 --asset-conditions="uguu:1:2,stake:10:2" --integration-type="GOLANG_MOD" --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator proposal-update-yieldfarmtarget --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --assetmanagement-account-address="" --assetmanagement-target-id="OsmosisGUUFarm" --unbonding-seconds=10 --asset-conditions="uguu:1:2,stake:10:2" --integration-type="GOLANG_MOD" --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator proposal-remove-yieldfarmtarget --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --assetmanagement-target-id="OsmosisGUUFarm" --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator proposal-stop-yieldfarmtarget --title="title" --description="description" --deposit=10000000stake --assetmanagement-account-id="OsmosisFarm" --assetmanagement-target-id="OsmosisGUUFarm" --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

ununifid query gov proposals
ununifid tx gov vote 2 yes  --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block