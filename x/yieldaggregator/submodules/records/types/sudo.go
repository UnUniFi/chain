package types

// MessageTransferCallback is passed to a contract's sudo() entrypoint for ibc transfer callback.
type MessageTransferCallback struct {
	TransferCallback struct {
		Denom    string `json:"denom"`
		Amount   string `json:"amount"`
		Sender   string `json:"sender"`
		Receiver string `json:"receiver"`
		Memo     string `json:"memo"`
		Success  bool   `json:"success"`
	} `json:"transfer_callback"`
}
