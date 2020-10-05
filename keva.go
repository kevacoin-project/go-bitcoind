package bitcoind

// NamespaceEntry - namespace info
type NamespaceEntry struct {
	NamespaceID string `json:"namespaceId"`
	DisplayName string `json:"displayName"`
}

// ValueResult - result returned from keva_get
type ValueResult struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	TxID   string `json:"txid"`
	Vout   int    `json:"vout"`
	Height int    `json:"height"`
}

// PutResult - result returned from keva_put
type PutResult struct {
	TxID string `json:"txid"`
}
