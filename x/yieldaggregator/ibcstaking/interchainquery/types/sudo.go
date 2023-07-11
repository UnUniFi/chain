package types

// MessageKVQueryResult is passed to a contract's sudo() entrypoint when a result
// was submitted for a kv-query.
type MessageKVQueryResult struct {
	KVQueryResult struct {
		QueryType string `json:"query_id"`
		Request   []byte `json:"request"`
		Data      []byte `json:"data"`
	} `json:"kv_query_result"`
}
