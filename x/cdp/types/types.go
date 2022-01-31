package types

import "encoding/json"

type DebtDenomMap map[string]string

func NewDebtDenomMap(dps DebtParams) DebtDenomMap {
	if len(dps) == 0 {
		panic("is empty DebtParams")
	}
	new_map := make(DebtDenomMap)
	for _, dp := range dps {
		if dp.DebtDenom == "" {
			panic("not exists DebtDenom")
		}
		new_map[dp.Denom] = dp.DebtDenom
	}
	return new_map
}

func NewDebtDenomMapFromByte(dps_bytes []byte) DebtDenomMap {
	if len(dps_bytes) == 0 {
		panic("is empty DebtParams")
	}
	var debutDenomMap DebtDenomMap
	err := json.Unmarshal(dps_bytes, &debutDenomMap)
	if err != nil {
		panic("is not DebtParams byte")
	}
	return debutDenomMap
}

func (d DebtDenomMap) Byte() []byte {
	bytes, err := json.Marshal(d)
	if err != nil {
		return []byte{}
	}
	return bytes
}
