package wasmbinding

import (
	"encoding/json"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/UnUniFi/chain/wasmbinding/bindings"
	icqkeeper "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/keeper"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/types"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(bank *bankkeeper.BaseKeeper, icqKeeper *icqkeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:   old,
			bank:      bank,
			icqKeeper: icqKeeper,
		}
	}
}

type CustomMessenger struct {
	wrapped   wasmkeeper.Messenger
	bank      *bankkeeper.BaseKeeper
	icqKeeper *icqkeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path where this is really creating / minting / swapping ...
		// leave everything else for the wrapped version
		var contractMsg bindings.UnunifiMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "ununifi msg")
		}
		if contractMsg.SubmitICQRequest != nil {
			return m.submitICQRequest(ctx, contractAddr, contractMsg.SubmitICQRequest)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) submitICQRequest(ctx sdk.Context, contractAddr sdk.AccAddress, submitICQRequest *bindings.SubmitICQRequest) ([]sdk.Event, [][]byte, error) {
	err := PerformSubmitICQRequest(m.icqKeeper, m.bank, ctx, contractAddr, submitICQRequest)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "perform icq request submission")
	}
	return nil, nil, nil
}

func PerformSubmitICQRequest(f *icqkeeper.Keeper, b *bankkeeper.BaseKeeper, ctx sdk.Context, contractAddr sdk.AccAddress, submitICQRequest *bindings.SubmitICQRequest) error {
	if submitICQRequest == nil {
		return wasmvmtypes.InvalidRequest{Err: "icq request empty"}
	}

	ttl := ctx.BlockTime().Add(time.Hour * 504).Nanosecond() // 3 weeks
	err := f.MakeRequest(
		ctx,
		submitICQRequest.ConnectionId,
		submitICQRequest.ChainId,
		// use "bank" store to access acct balances which live in the bank module
		// use "key" suffix to retrieve a proof alongside the query result
		submitICQRequest.QueryPrefix,
		submitICQRequest.QueryKey,
		sdk.NewInt(-1),
		interchainquerytypes.ModuleName,
		contractAddr.String(), // set contract address on callback id
		uint64(ttl),           // ttl
		0,                     // height always 0 (which means current height)
	)
	if err != nil {
		return sdkerrors.Wrap(err, "creating icq request")
	}
	return nil
}

// parseAddress parses address from bech32 string and verifies its format.
func parseAddress(addr string) (sdk.AccAddress, error) {
	parsed, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "address from bech32")
	}
	err = sdk.VerifyAddressFormat(parsed)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "verify address format")
	}
	return parsed, nil
}
