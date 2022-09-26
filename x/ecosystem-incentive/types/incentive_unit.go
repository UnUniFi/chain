package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewIncentiveUnit(incentiveId string, subjectsInfo []SubjectInfo) IncentiveUnit {
	return IncentiveUnit{
		IncentiveId:     incentiveId,
		SubjectInfoList: subjectsInfo,
	}
}

func NewSubjectInfo(subjectAddr sdk.AccAddress, weight sdk.Dec) SubjectInfo {
	return SubjectInfo{
		Address: subjectAddr.Bytes(),
		Weight:  weight,
	}
}
