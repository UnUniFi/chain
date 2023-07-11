package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v7/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	tendermint "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	ics23 "github.com/cosmos/ics23/go"
	"github.com/spf13/cast"

	"github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/types"
)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: &keeper}
}

var _ types.MsgServer = msgServer{}

// check if the query requires proving; if it does, verify it!
func (k Keeper) VerifyKeyProof(ctx sdk.Context, msg *types.MsgSubmitQueryResponse, query types.Query) error {
	pathParts := strings.Split(query.QueryType, "/")

	// the query does NOT have an associated proof, so no need to verify it.
	if pathParts[len(pathParts)-1] != "key" {
		return nil
	}

	// If the query is a "key" proof query, verify the results are valid by checking the poof
	if msg.ProofOps == nil {
		return errorsmod.Wrapf(types.ErrInvalidICQProof, "Unable to validate proof. No proof submitted")
	}

	// Get the client consensus state at the height 1 block above the message height
	msgHeight, err := cast.ToUint64E(msg.Height)
	if err != nil {
		return err
	}
	height := clienttypes.NewHeight(clienttypes.ParseChainID(query.ChainId), msgHeight+1)

	// Get the client state and consensus state from the connection Id
	connection, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, query.ConnectionId)
	if !found {
		return errorsmod.Wrapf(types.ErrInvalidICQProof, "ConnectionId %s does not exist", query.ConnectionId)
	}

	clientState, found := k.IBCKeeper.ClientKeeper.GetClientState(ctx, connection.ClientId)
	if !found {
		return errorsmod.Wrapf(types.ErrInvalidICQProof, "Unable to fetch client state for client %s", connection.ClientId)
	}

	consensusState, found := k.IBCKeeper.ClientKeeper.GetClientConsensusState(ctx, connection.ClientId, height)
	if !found {
		return errorsmod.Wrapf(types.ErrInvalidICQProof, "Consensus state not found for client %s and height %d", connection.ClientId, height)
	}
	var stateRoot exported.Root
	var clientStateProof []*ics23.ProofSpec

	switch clientState.ClientType() {
	case exported.Tendermint:
		tendermintConsensusState, ok := consensusState.(*tendermint.ConsensusState)
		if !ok {
			errMsg := fmt.Sprintf("[ICQ Resp] for query %s, error unmarshaling client state %v", query.Id, clientState)
			return sdkerrors.Wrapf(types.ErrInvalidICQProof, errMsg)
		}
		stateRoot = tendermintConsensusState.GetRoot()
	// case exported.Wasm:
	// 	wasmConsensusState, ok := consensusState.(*wasm.ConsensusState)
	// 	if !ok {
	// 		return errorsmod.Wrapf(types.ErrInvalidConsensusState, "Error casting consensus state: %s", err.Error())
	// 	}
	// 	tmClientState, ok := clientState.(*tmclienttypes.ClientState)
	// 	if !ok {
	// 		return errorsmod.Wrapf(types.ErrInvalidICQProof, "Client state is not tendermint")
	// 	}
	// 	clientStateProof = tmClientState.ProofSpecs
	// 	stateRoot = wasmConsensusState.GetRoot()
	default:
		panic("not implemented")
	}

	// Get the merkle path and merkle proof
	path := commitmenttypes.NewMerklePath([]string{pathParts[1], url.PathEscape(string(query.Request))}...)
	merkleProof, err := commitmenttypes.ConvertProofs(msg.ProofOps)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidICQProof, "Error converting proofs: %s", err.Error())
	}

	// If we got a non-nil response, verify inclusion proof
	if len(msg.Result) != 0 {
		if err := merkleProof.VerifyMembership(clientStateProof, stateRoot, path, msg.Result); err != nil {
			return errorsmod.Wrapf(types.ErrInvalidICQProof, "Unable to verify membership proof: %s", err.Error())
		}
		k.Logger(ctx).Info(fmt.Sprintf("Proof validated! module: %s, queryId %s", types.ModuleName, query.Id))

	} else {
		// if we got a nil query response, verify non inclusion proof.
		if err := merkleProof.VerifyNonMembership(clientStateProof, stateRoot, path); err != nil {
			return errorsmod.Wrapf(types.ErrInvalidICQProof, "Unable to verify non-membership proof: %s", err.Error())
		}
		k.Logger(ctx).Info(fmt.Sprintf("Non-inclusion Proof validated, stopping here! module: %s, queryId %s", types.ModuleName, query.Id))
	}

	return nil
}

// call the query's associated callback function
func (k Keeper) InvokeCallback(ctx sdk.Context, msg *types.MsgSubmitQueryResponse, q types.Query) error {
	// get all the stored queries and sort them for determinism
	moduleNames := []string{}
	for moduleName := range k.callbacks {
		moduleNames = append(moduleNames, moduleName)
	}
	sort.Strings(moduleNames)

	for _, moduleName := range moduleNames {
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] executing callback for queryId (%s), module (%s)", q.Id, moduleName))
		moduleCallbackHandler := k.callbacks[moduleName]

		if moduleName == types.ModuleName { // if callback module is icq module, call contract
			contractAddress := sdk.MustAccAddressFromBech32(q.CallbackId)

			x := types.MessageKVQueryResult{}
			x.KVQueryResult.ConnectionId = q.ConnectionId
			x.KVQueryResult.ChainId = q.ChainId
			x.KVQueryResult.QueryPrefix = q.QueryType
			x.KVQueryResult.QueryKey = q.Request
			x.KVQueryResult.Data = msg.Result

			m, err := json.Marshal(x)
			if err != nil {
				return fmt.Errorf("failed to marshal MessageKVQueryResult: %v", err)
			}

			_, err = k.wasmKeeper.Sudo(ctx, contractAddress, m)
			if err != nil {
				k.Logger(ctx).Debug("SudoTxQueryResult: failed to Sudo",
					"error", err, "contract_address", contractAddress)
				return fmt.Errorf("failed to Sudo: %v", err)
			}
		} else if moduleCallbackHandler.Has(q.CallbackId) {
			k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] callback (%s) found for module (%s)", q.CallbackId, moduleName))
			// call the correct callback function
			err := moduleCallbackHandler.Call(ctx, q.CallbackId, msg.Result, q)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("[ICQ Resp] error in ICQ callback, error: %s, msg: %s, result: %v, type: %s, params: %v", err.Error(), msg.QueryId, msg.Result, q.QueryType, q.Request))
				return err
			}
		} else {
			k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] callback not found for module (%s)", moduleName))
		}
	}
	return nil
}

// verify the query has not exceeded its ttl
func (k Keeper) HasQueryExceededTtl(ctx sdk.Context, msg *types.MsgSubmitQueryResponse, query types.Query) (bool, error) {
	k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] query %s with ttl: %d, resp time: %d.", msg.QueryId, query.Ttl, ctx.BlockHeader().Time.UnixNano()))
	currBlockTime, err := cast.ToUint64E(ctx.BlockTime().UnixNano())
	if err != nil {
		return false, err
	}

	if query.Ttl < currBlockTime {
		errMsg := fmt.Sprintf("[ICQ Resp] aborting query callback due to ttl expiry! ttl is %d, time now %d for query of type %s with id %s, on chain %s",
			query.Ttl, ctx.BlockTime().UnixNano(), query.QueryType, query.ChainId, msg.QueryId)
		fmt.Println(errMsg)
		k.Logger(ctx).Error(errMsg)
		return true, nil
	}
	return false, nil
}

func (k msgServer) SubmitQueryResponse(goCtx context.Context, msg *types.MsgSubmitQueryResponse) (*types.MsgSubmitQueryResponseResponse, error) {
	fmt.Println("DEBUG SubmitQueryResponse", string(k.cdc.MustMarshalJSON(msg)))
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the response has an associated query stored on stride
	q, found := k.GetQuery(ctx, msg.QueryId)
	if !found {
		k.Logger(ctx).Info("[ICQ Resp] ignoring non-existent query response (note: duplicate responses are nonexistent)")
		return &types.MsgSubmitQueryResponseResponse{}, nil // technically this is an error, but will cause the entire tx to fail if we have one 'bad' message, so we can just no-op here.
	}

	defer ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyQueryId, q.Id),
		),
		sdk.NewEvent(
			"query_response",
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyQueryId, q.Id),
			sdk.NewAttribute(types.AttributeKeyChainId, q.ChainId),
		),
	})

	// 1. verify the response's proof, if one exists
	err := k.VerifyKeyProof(ctx, msg, q)
	if err != nil {
		return nil, err
	}
	// 2. immediately delete the query so it cannot process again
	k.DeleteQuery(ctx, q.Id)

	// 3. verify the query's ttl is unexpired
	ttlExceeded, err := k.HasQueryExceededTtl(ctx, msg, q)
	if err != nil {
		return nil, err
	}
	if ttlExceeded {
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] %s's ttl exceeded: %d < %d.", msg.QueryId, q.Ttl, ctx.BlockHeader().Time.UnixNano()))
		return &types.MsgSubmitQueryResponseResponse{}, nil
	}

	// 4. if the query is contentless, end
	if len(msg.Result) == 0 {
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] query %s is contentless, removing from store.", msg.QueryId))
		return &types.MsgSubmitQueryResponseResponse{}, nil
	}

	// 5. call the query's associated callback function
	err = k.InvokeCallback(ctx, msg, q)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitQueryResponseResponse{}, nil
}
