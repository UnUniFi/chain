package bindings

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
)

type UnunifiMsg struct {
	SubmitICQRequest *SubmitICQRequest        `json:"submit_i_c_q_request,omitempty"`
	IBCTransfer      *wasmvmtypes.TransferMsg `json:"ibc_transfer,omitempty"`
	DeputyListNft    *DeputyListNft           `json:"deputy_list_nft,omitempty"`
}

type SubmitICQRequest struct {
	ConnectionId string `json:"connection_id"`
	ChainId      string `json:"chain_id"`
	QueryPrefix  string `json:"query_prefix"`
	QueryKey     []byte `json:"query_key"`
}

type DeputyListNft struct {
	Lister         string `json:"lister"`
	ClassId        string `json:"class_id"`
	TokenId        string `json:"token_id"`
	BidDenom       string `json:"bid_denom"`
	MinDepositRate string `json:"min_deposit_rate"` // TODO: type
	MinBidPeriod   string `json:"min_bid_period"`   // TODO: type
}
