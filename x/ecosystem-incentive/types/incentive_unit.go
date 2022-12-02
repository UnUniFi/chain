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
		Address: subjectAddr,
		Weight:  weight,
	}
}
