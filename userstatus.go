package mpos

type UserStatus struct {
	Username  string     `json:"username"`
	Shares    UserShares `json:"shares"`
	Hashrate  int        `json:"hashrate"`
	Sharerate int        `json:"sharerate"`
}

type UserShares struct {
	Valid         int `json:"valid"`
	Invalid       int `json:"invalid"`
	DonatePercent int `json:"donate_percent"`
	IsAnonymous   int `json:"is_anonymous"`
}

type userStatusResponse struct {
	Result struct {
		       Version string     `json:"version"`
		       Runtime float32    `json:"runtime"`
		       Data    UserStatus `json:"data"`
	       } `json:"getuserstatus"`
}

func (client *MposClient) GetUserStatus() (UserStatus, error) {
	userstatus := userStatusResponse{}
	req := &MposRequest{Page: "api", Action:"getuserstatus", Apikey:client.apikey}
	_, err := client.sling.New().Get("index.php").QueryStruct(req).ReceiveSuccess(&userstatus)
	if err != nil {
		return userstatus.Result.Data, err
	}

	return userstatus.Result.Data, err
}
