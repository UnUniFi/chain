package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
)

func NewIncentiveUnit(id string, subjectsInfo []SubjectInfo) IncentiveUnit {
	return IncentiveUnit{
		Id:              id,
		SubjectInfoList: subjectsInfo,
	}
}

func NewSubjectInfo(subjectAddr ununifitypes.StringAccAddress, weight sdk.Dec) SubjectInfo {
	return SubjectInfo{
		Address: subjectAddr,
		Weight:  weight,
	}
}
