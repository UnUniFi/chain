package bindings

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
)

type UnunifiMsg struct {
	SubmitICQRequest *SubmitICQRequest        `json:"submit_i_c_q_request,omitempty"`
	IBCTransfer      *wasmvmtypes.TransferMsg `json:"ibc_transfer,omitempty"`
}

type SubmitICQRequest struct {
	ConnectionId string `json:"connection_id"`
	ChainId      string `json:"chain_id"`
	QueryPrefix  string `json:"query_prefix"`
	QueryKey     []byte `json:"query_key"`
}
