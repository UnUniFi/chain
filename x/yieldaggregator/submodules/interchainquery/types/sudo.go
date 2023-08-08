package types

// MessageKVQueryResult is passed to a contract's sudo() entrypoint when a result
// was submitted for a kv-query.
type MessageKVQueryResult struct {
	KVQueryResult struct {
		ConnectionId string `json:"connection_id"`
		ChainId      string `json:"chain_id"`
		QueryPrefix  string `json:"query_prefix"`
		QueryKey     []byte `json:"query_key"`
		Data         []byte `json:"data"`
	} `json:"kv_query_result"`
}
