# READ ME
# This script is a test script to check basic operation of upgrade-v1.
# Note: It is intended to be run in a clean environment.
# Please be careful not to run it in a production environment or 
# in an environment where ununifid has already been set up.

# Wait until the height reaches 20 or more.
# Verify that BankSend is running correctly.
# height : .ununifi/data/priv_validator_state.json

echo "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7 : amount:4999999699994"
echo "ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz : amount:125000000300006"
ununifid query bank balances ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7
ununifid query bank balances ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz