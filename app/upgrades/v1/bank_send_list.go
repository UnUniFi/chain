package v1

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BankSendList(ctx sdk.Context) (ResultList, error) {
	ctx.Logger().Info(fmt.Sprintf("bank send list:%s", UpgradeName))

	// Read file and get list
	var result ResultList
	json.Unmarshal([]byte(BANK_SEND_LIST), &result)

	return result, nil
}

const BANK_SEND_LIST string = `{
	"validator": [
		{
			"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
			"toAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
			"amount": "uguu",
			"denom": 100001,
			"vesting_starts": 1660521600, 
			"vesting_ends": 1660696500
		}
	],
	"airdropCommunityRewardModerator": [
			{
					"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
					"toAddress": "ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz",
					"amount": "uguu",
					"denom": 100002,
					"vesting_starts": 1660521600, 
					"vesting_ends": 1660696500
			},
			{
					"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
					"toAddress": "ununifi1mtvjd2rsyll8kps6qqkxd6p78mr8qkjx27tn2p",
					"amount": "uguu",
					"denom": 100003,
					"vesting_starts": 1660521600, 
					"vesting_ends": 1660696500
			},
			{
					"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
					"toAddress": "ununifi14x04hcu8gmku53ll04v96tdgae84h2ylmkal9k",
					"amount": "uguu",
					"denom": 100004,
					"vesting_starts": 1660521600, 
					"vesting_ends": 1660696500
			},
			{
					"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
					"toAddress": "ununifi16ayyysehst594k98a7leym6l5jrrhgf9yk9hn5",
					"amount": "uguu",
					"denom": 100005,
					"vesting_starts": 0, 
					"vesting_ends": 1660696500
			}
	]
}`
