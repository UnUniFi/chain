package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewIncentiveUnit(id string, subjectsInfo []SubjectInfo) IncentiveUnit {
	return IncentiveUnit{
		Id:              id,
		SubjectInfoList: subjectsInfo,
	}
}

func NewSubjectInfo(subjectAddr sdk.AccAddress, weight sdk.Dec) SubjectInfo {
	return SubjectInfo{
		Address: subjectAddr.Bytes(),
		Weight:  weight,
	}
}
