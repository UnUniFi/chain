package bindings

type UnunifiMsg struct {
	/// Contracts can create denoms, namespaced under the contract's address.
	/// A contract may create any number of independent sub-denoms.
	SubmitICQRequest *SubmitICQRequest `json:"create_denom,omitempty"`
}

// CreateDenom creates a new factory denom, of denomination:
// factory/{creating contract address}/{Subdenom}
// Subdenom can be of length at most 44 characters, in [0-9a-zA-Z./]
// The (creating contract address, subdenom) pair must be unique.
// The created denom's admin is the creating contract address,
// but this admin can be changed using the ChangeAdmin binding.
type SubmitICQRequest struct {
	ConnectionId string `json:"connection_id"`
	ChainId      string `json:"chain_id"`
	QueryPrefix  string `json:"query_prefix"`
	QueryKey     []byte `json:"query_key"`
}
