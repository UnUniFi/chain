package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewRecipientContainer(id string, subjectsInfos []WeightedAddress) RecipientContainer {
	// check if the addresses in subjectInfo are valid AccAddress
	for _, subjectInfo := range subjectsInfos {
		if _, err := sdk.AccAddressFromBech32(subjectInfo.Address); err != nil {
			panic(err)
		}
	}

	return RecipientContainer{
		Id:                id,
		WeightedAddresses: subjectsInfos,
	}
}

func NewSubjectInfo(subjectAddr string, weight sdk.Dec) WeightedAddress {
	// check if the address in subjectInfo are valid AccAddress
	if _, err := sdk.AccAddressFromBech32(subjectAddr); err != nil {
		panic(err)
	}

	return WeightedAddress{
		Address: subjectAddr,
		Weight:  weight,
	}
}

func NewRecipientContainerIdsByAddr(address string, recipientContainerId string) BelongingRecipientContainers {
	// check if the address are valid AccAddress
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		panic(err)
	}

	var recipientContainerIds []string
	recipientContainerIds = append(recipientContainerIds, recipientContainerId)

	return BelongingRecipientContainers{
		Address:               address,
		RecipientContainerIds: recipientContainerIds,
	}
}

func (m BelongingRecipientContainers) AddRecipientContainerId(recipientContainerId string) []string {
	return append(m.RecipientContainerIds, recipientContainerId)
}

func (m BelongingRecipientContainers) CreateOrUpdate(address string, recipientContainerId string) BelongingRecipientContainers {
	// check if the address are valid AccAddress
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		panic(err)
	}

	if len(m.Address) == 0 {
		m = NewRecipientContainerIdsByAddr(address, recipientContainerId)
	} else {
		m.RecipientContainerIds = m.AddRecipientContainerId(recipientContainerId)
	}

	return m
}
