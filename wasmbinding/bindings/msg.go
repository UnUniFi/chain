package bindings

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
)

type UnunifiMsg struct {
	// v0 messages
	SubmitICQRequest *SubmitICQRequest `json:"submit_i_c_q_request,omitempty"`
	// v1 messages
	IBCTransfer            *wasmvmtypes.TransferMsg `json:"ibc_transfer,omitempty"`
	RequestKvIcq           *SubmitICQRequest        `json:"request_kv_icq,omitempty"`
	DeputyDepositToVault   *DeputyDepositToVault    `json:"deputy_deposit_to_vault,omitempty"`
	DeputyDepositToTranche *DeputyDepositToTranche  `json:"deputy_deposit_to_tranche,omitempty"`
}

type SubmitICQRequest struct {
	ConnectionId string `json:"connection_id"`
	ChainId      string `json:"chain_id"`
	QueryPrefix  string `json:"query_prefix"`
	QueryKey     []byte `json:"query_key"`
}

type DeputyDepositToVault struct {
	Depositor string           `json:"depositor"`
	VaultId   string           `json:"vault_id"`
	Amount    wasmvmtypes.Coin `json:"amount"`
}

type DeputyDepositToTranche struct {
	Depositor   string            `json:"depositor"`
	TrancheId   string            `json:"tranche_id"`
	TrancheType string            `json:"tranche_type"`
	Amount      wasmvmtypes.Coins `json:"amount"`
}
