package types

func ValidateIncentiveUnitIdLen(maxIncentiveUnitIdLen uint64, incentiveUnitId string) error {
	if len(incentiveUnitId) > int(maxIncentiveUnitIdLen) || len(incentiveUnitId) == 0 {
		return ErrInvalidIncentiveUnitIdLen
	}

	return nil
}

func ValidateSubjectInfoNumInUnit(maxSubjectInfoNumInUnit uint64, incentiveUnit IncentiveUnit) error {
	if len(incentiveUnit.SubjectInfoLists) > int(maxSubjectInfoNumInUnit) || len(incentiveUnit.SubjectInfoLists) == 0 {
		return ErrInvalidSubjectInfoNumInUnit
	}

	return nil
}
