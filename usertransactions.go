package mpos

type UserTransactions struct {
	Transactions       []UserTransaction `json:"transactions"`
	TransactionSummary UserTransactionSummary `json:"transactionsummary"`
}

type UserTransaction struct {
	Id            int `json:"id"`
	Username      string `json:"username"`
	Type          string `json:"type"`
	Amount        float64 `json:"amount"`
	CoinAddress   string `json:"coin_address"`
	Timestamp     TransactionTimestamp `json:"timestamp"`
	TxId          string `json:"txid"`
	Height        uint32 `json:"height"`
	BlockHash     string `json:"blockhash"`
	Confirmations uint16 `json:"confirmations"`
}

type UserTransactionSummary struct {
	Bonus   float64 `json:"Bonus"`
	Credit  float64 `json:"Credit"`
	DebitAP float64 `json:"Debit_AP"`
	Fee     float64 `json:"Fee"`
	TXFee   float64 `json:"TXFee"`
}

type userTransactionsResponse struct {
	Result struct {
		       Version string     `json:"version"`
		       Runtime float32    `json:"runtime"`
		       Data    UserTransactions `json:"data"`
	       } `json:"getusertransactions"`
}

func (client *MposClient) GetUserTransactions() (UserTransactions, error) {
	usertransactions := userTransactionsResponse{}
	req := &MposRequest{Page: "api", Action:"getusertransactions", Apikey:client.apikey}
	_, err := client.sling.New().Get("index.php").QueryStruct(req).ReceiveSuccess(&usertransactions)
	if err != nil {
		return usertransactions.Result.Data, err
	}

	return usertransactions.Result.Data, err
}

