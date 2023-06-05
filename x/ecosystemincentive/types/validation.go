package types

func ValidateIncentiveUnitIdLen(maxIncentiveUnitIdLen uint64, incentiveUnitId string) error {
	if len(incentiveUnitId) > int(maxIncentiveUnitIdLen) || len(incentiveUnitId) == 0 {
		return ErrInvalidIncentiveUnitIdLen
	}

	return nil
}

func ValidateSubjectInfoNumInUnit(maxSubjectInfoNumInUnit uint64, incentiveUnit RecipientContainer) error {
	if len(incentiveUnit.WeightedAddresses) > int(maxSubjectInfoNumInUnit) || len(incentiveUnit.WeightedAddresses) == 0 {
		return ErrInvalidSubjectInfoNumInUnit
	}

	return nil
}
