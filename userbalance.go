package mpos

type UserBalance struct {
	Confirmed   float64 `json:"confirmed"`
	Unconfirmed float64 `json:"unconfirmed"`
	Orphaned    float64 `json:"orphaned"`
}

type userBalanceResponse struct {
	Result struct {
		       Version string     `json:"version"`
		       Runtime float32    `json:"runtime"`
		       Data    UserBalance `json:"data"`
	       } `json:"getuserbalance"`
}

func (client *MposClient) GetUserBalance() (UserBalance, error) {
	userbalance := userBalanceResponse{}
	req := &MposRequest{Page: "api", Action:"getuserbalance", Apikey:client.apikey}
	_, err := client.sling.New().Get("index.php").QueryStruct(req).ReceiveSuccess(&userbalance)
	if err != nil {
		return userbalance.Result.Data, err
	}

	return userbalance.Result.Data, err
}

