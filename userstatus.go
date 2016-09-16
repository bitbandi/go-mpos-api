package mpos

type UserStatus struct {
	Username  string     `json:"username"`
	Shares    UserShares `json:"shares"`
	Hashrate  float64        `json:"hashrate"`
	Sharerate float64        `json:"sharerate"`
}

type UserShares struct {
	Valid         float64 `json:"valid"`
	Invalid       float64 `json:"invalid"`
	DonatePercent float64 `json:"donate_percent"`
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
