# READ ME
# This script is a test script to check basic operation of upgrade-v1.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.

# Wait until the height reaches 20 or more.
# Verify that BankSend is running correctly.
# height : .ununifi/data/priv_validator_state.json
echo "validator"
ununifid query bank balances ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7
echo "faucet"
# new ContinuousVestingAccount
ununifid query bank balances ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz
# ContinuousVestingAccount
ununifid query bank balances ununifi14x04hcu8gmku53ll04v96tdgae84h2ylmkal9k
# DelayedVestingAccount
ununifid query bank balances ununifi1mtvjd2rsyll8kps6qqkxd6p78mr8qkjx27tn2p
# DelayedVestingAccount
ununifid query bank balances ununifi16ayyysehst594k98a7leym6l5jrrhgf9yk9hn5
