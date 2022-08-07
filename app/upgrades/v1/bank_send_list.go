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
	"response": [
			{
					"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
					"toAddress": "ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz",
					"amount": "uguu",
					"denom": 100001
			},
			{
					"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
					"toAddress": "ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz",
					"amount": "uguu",
					"denom": 100002
			},
			{
					"fromAddress": "ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7",
					"toAddress": "ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz",
					"amount": "uguu",
					"denom": 100003
			}
	]
}`
