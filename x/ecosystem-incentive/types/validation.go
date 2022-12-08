package types

func ValidateMaxSubjectInfoNumInUnit(maxSubjectInfoNumInUnit uint64, incentiveUnit IncentiveUnit) error {
	if len(incentiveUnit.SubjectInfoLists) > int(maxSubjectInfoNumInUnit) {
		return ErrInvalidSubjectInfoNumInUnit
	}

	return nil
}
