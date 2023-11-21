package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AmountsResp struct {
	TotalDeposited string `json:"total_deposited"`
	BondingStandby string `json:"bonding_standby"`
	Bonded         string `json:"bonded"`
	Unbonding      string `json:"unbonding"`
}

func (k Keeper) GetAmountFromStrategy(ctx sdk.Context, sender sdk.AccAddress, strategyContract string) (sdk.Int, error) {
	wasmQuery := fmt.Sprintf(`{"amounts":{"addr": "%s"}}`, sender.String())
	contractAddr := sdk.MustAccAddressFromBech32(strategyContract)
	resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	parsedAmounts := AmountsResp{}
	err = json.Unmarshal(resp, &parsedAmounts)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	amount, ok := sdk.NewIntFromString(parsedAmounts.Bonded)
	if !ok {
		return sdk.ZeroInt(), nil
	}
	return amount, err

}

func (k Keeper) GetUnbondingAmountFromStrategy(ctx sdk.Context, sender sdk.AccAddress, strategyContract string) (sdk.Int, error) {
	wasmQuery := fmt.Sprintf(`{"amounts":{"addr": "%s"}}`, sender.String())
	contractAddr := sdk.MustAccAddressFromBech32(strategyContract)
	resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	parsedAmounts := AmountsResp{}
	err = json.Unmarshal(resp, &parsedAmounts)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	amount, ok := sdk.NewIntFromString(parsedAmounts.Unbonding)
	if !ok {
		return sdk.ZeroInt(), nil
	}
	return amount, err
}

func (k Keeper) GetStrategyVersion(ctx sdk.Context, strategyContract string) uint8 {
	wasmQuery := fmt.Sprintf(`{"version":{}}`)
	contractAddr := sdk.MustAccAddressFromBech32(strategyContract)
	result, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
	if err != nil {
		return 0
	}

	jsonMap := make(map[string]uint8)
	err = json.Unmarshal(result, &jsonMap)
	if err != nil {
		return 0
	}

	return jsonMap["version"]
}

type DenomInfo struct {
	Denom            string `json:"denom"`
	TargetChainId    string `json:"target_chain_id"`
	TargetChainDenom string `json:"target_chain_denom"`
	TargetChainAddr  string `json:"target_chain_addr"`
}

func (k Keeper) GetStrategyDepositInfo(ctx sdk.Context, strategyContract string) (info DenomInfo) {
	wasmQuery := fmt.Sprintf(`{"deposit_denom":{}}`)
	contractAddr := sdk.MustAccAddressFromBech32(strategyContract)
	result, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &info)
	if err != nil {
		return DenomInfo{}
	}

	return
}
