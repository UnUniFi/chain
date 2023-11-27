package v4_0_1

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BankSendList(ctx sdk.Context) (ResultList, error) {
	ctx.Logger().Info(fmt.Sprintf("bank send list:%s", UpgradeName))

	// Read file and get list
	var result ResultList
	if err := json.Unmarshal([]byte(BANK_SEND_LIST), &result); err != nil {
		panic(err)
	}
	return result, nil
}

const BANK_SEND_LIST string = `{
  "campaign": [
    {
			"number": 1,
      "toAddress": "ununifi10jath6g7kn8ly6jkthdrpu37dd72s565wvmrvx",
      "amount": 9407000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 2,
      "toAddress": "ununifi10t97e5z8qk0xgkcnkmg79k46c6s0thgvc5mfxy",
      "amount": 31852000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 3,
      "toAddress": "ununifi10vz7wnnhapuhwyyzcfq3ync4tjgvy0xk3jpf49",
      "amount": 5960000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 4,
      "toAddress": "ununifi1208sjd9egkgazc09vzmwma2f5wtam2xvxp8ef8",
      "amount": 9799000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 5,
      "toAddress": "ununifi12ndyfyxhzctjx6ygfgguzmam8j8jskhlrk3gqr",
      "amount": 75882000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 6,
      "toAddress": "ununifi12sqnxw4n8kqlchey5qj3z7tdlztej56u22dat3",
      "amount": 15156000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 7,
      "toAddress": "ununifi138dr9pkqljnztl5ax4hkaq3k2xwxcgcpjvlr39",
      "amount": 195559000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 8,
      "toAddress": "ununifi13ys7lplhpemeujxl89ageks3rnap9rmpq6nsr0",
      "amount": 95893000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 9,
      "toAddress": "ununifi167u7spephht5kt7347m7vrra6jpjmarh7vheh7",
      "amount": 33756000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 10,
      "toAddress": "ununifi16truj42f58f0zupeyas7gh8vn0f22kv78crn33",
      "amount": 329928000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 11,
      "toAddress": "ununifi16znhj0k8fufx2cqdl7e5mk9yxqaqa4vz80mypz",
      "amount": 4039000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 12,
      "toAddress": "ununifi179xh34gna498hdtglq7mv2hwrxqwzw8tfg43ec",
      "amount": 5775000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 13,
      "toAddress": "ununifi17c67lkjttj3kt4684he8qldre20996f4093hd2",
      "amount": 29613000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 14,
      "toAddress": "ununifi17c7l66wgut2trha0qv996tqtg4scwusp4wm649",
      "amount": 19558000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 15,
      "toAddress": "ununifi186zlv24r7wjwchp2eq04emk0r9qzctx6k5aqm8",
      "amount": 53628000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 16,
      "toAddress": "ununifi1a2uaj2tagfx9w3pken36pvdhfys35zmucax3we",
      "amount": 46393000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 17,
      "toAddress": "ununifi1d5csc65hy3l8yk9s8rges7zxl9fwka88f6waec",
      "amount": 22179000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 18,
      "toAddress": "ununifi1dgz844y8flmzlp6rgl5y2044sdyywm5q8sqngw",
      "amount": 20254000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 19,
      "toAddress": "ununifi1dm2ayct3tznq8rmclztjj9xyhwkx7p0gghkwzj",
      "amount": 16835000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 20,
      "toAddress": "ununifi1dq2fg39njz44f69cpu8khxtsg5rxlkszvjdjr9",
      "amount": 3126000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 21,
      "toAddress": "ununifi1f6dchhs6xqg05tjrscffrz48qcsmqmf9rm4ate",
      "amount": 19907000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 22,
      "toAddress": "ununifi1fxt80g5tf9tze0uxfgpkm5w0v8frr2dfcj4vav",
      "amount": 4016000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 23,
      "toAddress": "ununifi1jak3a0m3q2tgk7y23n7efsrr52d0m6h2ycv4zr",
      "amount": 20381000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 24,
      "toAddress": "ununifi1jesw4hglck9twv63nyup5yn6e0j4wmzfzrcc72",
      "amount": 3332000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 25,
      "toAddress": "ununifi1jgu55t8rnjpneytsrlf6ga2t7kz6f57zkjt98u",
      "amount": 3517000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 26,
      "toAddress": "ununifi1jmg6q5p5ply88y5cuqprj2nv7rh5m9yjwatgul",
      "amount": 5275000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 27,
      "toAddress": "ununifi1kuhp2ywypwg2ekuz05zw4dlvtm40zswfawf32v",
      "amount": 36006000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 28,
      "toAddress": "ununifi1lq2rnqqzph42gunckmhhy06ppnu9uxefpfvzgh",
      "amount": 35487000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 29,
      "toAddress": "ununifi1mlxpzqmut2freclfh3dxfds6c9xvvv2hmttr87",
      "amount": 188861000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 30,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 31,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 32,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 33,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 34,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 35,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 36,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 37,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 38,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 39,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 40,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 41,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 42,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 43,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 44,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 45,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 46,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 47,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 48,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 49,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 50,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    },
		{
			"number": 51,
      "toAddress": "",
      "amount": 000000,
      "denom": "uguu",
      "vesting_starts": 1671548400,
      "vesting_ends": 1702998000
    }
  ]
}`
