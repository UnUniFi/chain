package types

import (
	"encoding/json"
	"unsafe"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StringAccAddress sdk.AccAddress

func StringAccAddresses(addrs []sdk.AccAddress) []StringAccAddress {
	return *(*[]StringAccAddress)(unsafe.Pointer(&addrs))
}

func (aa StringAccAddress) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(aa)
}

func AccAddresses(addrs []StringAccAddress) []sdk.AccAddress {
	return *(*[]sdk.AccAddress)(unsafe.Pointer(&addrs))
}

func (aa StringAccAddress) Marshal() ([]byte, error) {
	str := aa.AccAddress().String()

	return json.Marshal(str)
}

// MarshalTo implements the gogo proto custom type interface.
func (aa *StringAccAddress) MarshalTo(data []byte) (n int, err error) {
	bz, err := aa.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}

// Unmarshal implements the gogo proto custom type interface.
func (aa *StringAccAddress) Unmarshal(data []byte) error {
	if len(data) == 0 {
		aa = nil
		return nil
	}

	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	_aa, err := sdk.AccAddressFromBech32(str)
	if err != nil {
		return err
	}
	*aa = StringAccAddress(_aa)

	return nil
}

// Size implements the gogo proto custom type interface.
func (aa *StringAccAddress) Size() int {
	bz, _ := aa.Marshal()
	return len(bz)
}
