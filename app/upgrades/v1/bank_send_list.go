package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BankSendList(ctx sdk.Context) (ResultList, error) {
	ctx.Logger().Info(fmt.Sprintf("bank send list:%s", UpgradeName))

	// Specify the JSON file path of the bank send list in the "upgradeBankSendListJson" environment variable.
	jsonBankSendList := os.Getenv("upgradeBankSendListJson")
	ctx.Logger().Info(fmt.Sprintf("jsonBankSendList:%s", jsonBankSendList))
	if jsonBankSendList == "" {
		ctx.Logger().Info(fmt.Sprintf("The environment variable for upgradeBankSendListJson is not set. It will use the default path [%s].", DefaultBankSendJsonFilePath))
		jsonBankSendList = DefaultBankSendJsonFilePath
	}

	// Read file and get list
	raw, err := ioutil.ReadFile(jsonBankSendList)
	var result ResultList
	if err != nil {
		return ResultList{}, err
	}

	json.Unmarshal(raw, &result)

	return result, nil
}
