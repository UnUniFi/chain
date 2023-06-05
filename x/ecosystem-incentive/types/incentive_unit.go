package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewIncentiveUnit(id string, subjectsInfos []SubjectInfo) IncentiveUnit {
	// check if the addresses in subjectInfo are valid AccAddress
	for _, subjectInfo := range subjectsInfos {
		if _, err := sdk.AccAddressFromBech32(subjectInfo.SubjectAddr); err != nil {
			panic(err)
		}
	}

	return IncentiveUnit{
		Id:               id,
		SubjectInfoLists: subjectsInfos,
	}
}

func NewSubjectInfo(subjectAddr string, weight sdk.Dec) SubjectInfo {
	// check if the address in subjectInfo are valid AccAddress
	if _, err := sdk.AccAddressFromBech32(subjectAddr); err != nil {
		panic(err)
	}

	return SubjectInfo{
		SubjectAddr: subjectAddr,
		Weight:      weight,
	}
}

func NewIncentiveUnitIdsByAddr(address string, incentiveUnitId string) IncentiveUnitIdsByAddr {
	// check if the address are valid AccAddress
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		panic(err)
	}

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

func (m IncentiveUnitIdsByAddr) CreateOrUpdate(address string, incentiveUnitId string) IncentiveUnitIdsByAddr {
	// check if the address are valid AccAddress
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		panic(err)
	}

	if len(m.Address) == 0 {
		m = NewIncentiveUnitIdsByAddr(address, incentiveUnitId)
	} else {
		m.IncentiveUnitIds = m.AddIncentiveUnitId(incentiveUnitId)
	}

	return m
}
