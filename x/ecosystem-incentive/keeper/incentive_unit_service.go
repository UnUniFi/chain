package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

// Register method record subjects info in IncentiveUnit type
func (k Keeper) Register(ctx sdk.Context, msg *types.MsgRegister) (*[]types.SubjectInfo, error) {
	if _, exists := k.GetIncentiveUnit(ctx, msg.IncentiveUnitId); exists {
		return nil, sdkerrors.Wrap(types.ErrRegisteredIncentiveId, msg.IncentiveUnitId)
	}

	var subjectInfoList []types.SubjectInfo
	for i := 0; i < len(msg.SubjectAddrs); i++ {
		subjectInfo := types.NewSubjectInfo(msg.SubjectAddrs[i].AccAddress(), msg.Weights[i])
		subjectInfoList = append(subjectInfoList, subjectInfo)
	}

	incentiveUnit := types.NewIncentiveUnit(msg.IncentiveUnitId, subjectInfoList)
	if err := k.SetIncentiveUnit(ctx, incentiveUnit); err != nil {
		return nil, err
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

func (k Keeper) DeleteIncentiveUnit(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnit))

	prefixStore.Delete([]byte(id))
}
