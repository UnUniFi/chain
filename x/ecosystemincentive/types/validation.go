package types

import (
	"regexp"
)

func ValidateRecipientContainerId(recipientContainerId string) error {
	if !regexp.MustCompile(RecipientContainerIdPattern).MatchString(recipientContainerId) {
		return ErrInvalidRecipientContainerId
	}

	return nil
}
