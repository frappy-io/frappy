package types

type TransactionType struct {
	TransactionId string `json:"transaction_id"`
	Payer         string `json:"payer"`
	Proposer      string `json:"proposer"`
	Authorizers   string `json:"authorizers"`
}

type TransactionResultType struct {
	Status string `json:"status"`
	Error  error  `json:"error"`
}

type AccountType struct {
	Address   string `json:"address"`
	Balance   int    `json:"balance"`
	Contracts int    `json:"contracts"`
	Keys      int    `json:"keys"`
}

type BlockType struct {
	ID        string `json:"id"`
	Height    int    `json:"height"`
	Timestamp string `json:"timestamp"`
}

type CollectionType struct {
	ID           string `json:"id"`
	Transactions string `json:"transaction"`
}
