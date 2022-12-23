package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

// Register method record subjects info in IncentiveUnit type
func (k Keeper) Register(ctx sdk.Context, msg *types.MsgRegister) (*[]types.SubjectInfo, error) {
	// check if the IncentiveUnitId is already registered
	if _, exists := k.GetIncentiveUnit(ctx, msg.IncentiveUnitId); exists {
		return nil, sdkerrors.Wrap(types.ErrRegisteredIncentiveId, msg.IncentiveUnitId)
	}

	// check the length of the IncentiveUnitId by referring MaxInentiveUnitIdLen in the Params
	if err := types.ValidateIncentiveUnitIdLen(k.GetMaxIncentiveUnitIdLen(ctx), msg.IncentiveUnitId); err != nil {
		return nil, err
	}

	var subjectInfoList []types.SubjectInfo
	for i := 0; i < len(msg.SubjectAddrs); i++ {
		subjectInfo := types.NewSubjectInfo(msg.SubjectAddrs[i], msg.Weights[i])
		subjectInfoList = append(subjectInfoList, subjectInfo)
	}

	incentiveUnit := types.NewIncentiveUnit(msg.IncentiveUnitId, subjectInfoList)

	// checks if the number of the subject info is vaid
	if err := types.ValidateSubjectInfoNumInUnit(k.GetMaxSubjectInfoNumInUnitParam(ctx), incentiveUnit); err != nil {
		return nil, err
	}

	if err := k.SetIncentiveUnit(ctx, incentiveUnit); err != nil {
		return nil, err
	}

	// operation related to IncentiveUnitIdsByAddr
	// if exists already, add incentuve unit id in msg into data
	// if not, newly create and set
	for _, addr := range msg.SubjectAddrs {
		incentiveUnitIdsByAddr := k.GetIncentiveUnitIdsByAddr(ctx, addr)
		incentiveUnitIdsByAddr = incentiveUnitIdsByAddr.CreateOrUpdate(addr, msg.IncentiveUnitId)

		if err := k.SetIncentiveUnitIdsByAddr(ctx, incentiveUnitIdsByAddr); err != nil {
			return nil, err
		}
	}

	return &subjectInfoList, nil
}

func (k Keeper) SetIncentiveUnit(ctx sdk.Context, incentiveUnit types.IncentiveUnit) error {
	bz, err := k.cdc.Marshal(&incentiveUnit)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnit))
	prefixStore.Set([]byte(incentiveUnit.Id), bz)

	return nil
}

func (k Keeper) SetIncentiveUnitIdsByAddr(ctx sdk.Context, incentiveUnitIdsByAddr types.IncentiveUnitIdsByAddr) error {
	bz, err := k.cdc.Marshal(&incentiveUnitIdsByAddr)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnitIdsByAddr))
	prefixStore.Set(incentiveUnitIdsByAddr.Address, bz)

	return nil
}

func (k Keeper) GetIncentiveUnit(ctx sdk.Context, id string) (types.IncentiveUnit, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnit))

	bz := prefixStore.Get([]byte(id))
	if len(bz) == 0 {
		return types.IncentiveUnit{}, false
	}

	var incentiveUnit types.IncentiveUnit
	k.cdc.MustUnmarshal(bz, &incentiveUnit)
	return incentiveUnit, true
}

func (k Keeper) GetIncentiveUnitIdsByAddr(ctx sdk.Context, address ununifitypes.StringAccAddress) types.IncentiveUnitIdsByAddr {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnitIdsByAddr))

	bz := prefixStore.Get(address)
	if len(bz) == 0 {
		return types.IncentiveUnitIdsByAddr{}
	}

	var incentiveUnitIdsByAddr types.IncentiveUnitIdsByAddr
	k.cdc.MustUnmarshal(bz, &incentiveUnitIdsByAddr)
	return incentiveUnitIdsByAddr
}

func (k Keeper) DeleteIncentiveUnit(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnit))

	prefixStore.Delete([]byte(id))
}

func (k Keeper) DeleteIncentiveUnitIdsByAddr(ctx sdk.Context, address ununifitypes.StringAccAddress) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnitIdsByAddr))

	prefixStore.Delete(address)
}
