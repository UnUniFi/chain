package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
)

func NewIncentiveUnit(id string, subjectsInfos []SubjectInfo) IncentiveUnit {
	return IncentiveUnit{
		Id:               id,
		SubjectInfoLists: subjectsInfos,
	}
}

func NewSubjectInfo(subjectAddr ununifitypes.StringAccAddress, weight sdk.Dec) SubjectInfo {
	return SubjectInfo{
		SubjectAddr: subjectAddr,
		Weight:      weight,
	}
}

func NewIncentiveUnitIdsByAddr(address ununifitypes.StringAccAddress, incentiveUnitId string) IncentiveUnitIdsByAddr {
	var incentiveUnitIds []string
	incentiveUnitIds = append(incentiveUnitIds, incentiveUnitId)

	return IncentiveUnitIdsByAddr{
		Address:          address,
		IncentiveUnitIds: incentiveUnitIds,
	}
}

func (m IncentiveUnitIdsByAddr) AddIncentiveUnitId(incentiveUnitId string) []string {
	return append(m.IncentiveUnitIds, incentiveUnitId)
}

func (m IncentiveUnitIdsByAddr) CreateOrUpdate(address ununifitypes.StringAccAddress, incentiveUnitId string) IncentiveUnitIdsByAddr {
	if m.Address.AccAddress().Empty() {
		m = NewIncentiveUnitIdsByAddr(address, incentiveUnitId)
	} else {
		m.IncentiveUnitIds = m.AddIncentiveUnitId(incentiveUnitId)
	}

	return m
}
