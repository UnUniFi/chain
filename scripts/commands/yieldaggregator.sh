#!/usr/bin/env bash

# queries
ununifid query yieldaggregator all-asset-management-accounts
ununifid query yieldaggregator all-farming-units
ununifid query yieldaggregator asset-management-account OsmosisFarm
ununifid query yieldaggregator params
ununifid query yieldaggregator user-info $(ununifid keys show -a validator --keyring-backend=test)

# farming order txs
ununifid tx yieldaggregator add-farming-order --farming-order-id="order1" --strategy-type="ManualStrategy" --whitelisted-target-ids="OsmosisFarmTarget1" --blacklisted-target-ids="" --max-unbonding-seconds=10 --overall-ratio=10 --min=1 --max=10 --date=1661967198 --active=true --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator inactivate-farming-order order1  --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator activate-farming-order order1 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator delete-farming-order order1 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator execute-farming-orders order1 --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# deposit/withdraw txs
ununifid tx yieldaggregator deposit 1000000uguu --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block
ununifid tx yieldaggregator withdraw 1000000uguu --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

# proposal txs
ununifid tx yieldaggregator proposal-add-yieldfarm --title="title" --description="description" --deposit=1000000uguu --assetmanagement-account-id="OsmosisFarm" --assetmanagement-account-name="Osmosis Farm" --enabled=true --chain-id=test --from=validator --keyring-backend=test --gas=300000 -y --broadcast-mode=block

ununifid tx yieldaggregator proposal-add-yieldfarmtarget    Submit a proposal to add a yield farm target
ununifid tx yieldaggregator proposal-remove-yieldfarm       Submit a proposal to remove a yield farm
ununifid tx yieldaggregator proposal-stop-yieldfarm         Submit a proposal to stop a yield farm
ununifid tx yieldaggregator proposal-stop-yieldfarmtarget   Submit a proposal to stop a yield farm target
ununifid tx yieldaggregator proposal-update-yieldfarm       Submit a proposal to update a yield farm
ununifid tx yieldaggregator proposal-update-yieldfarmtarget Submit a proposal to update a yield farm target
