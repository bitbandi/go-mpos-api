package mpos

import (
	"encoding/json"
)

type UserBalance struct {
	Confirmed   float64 `json:"confirmed"`
	Unconfirmed float64 `json:"unconfirmed"`
	Orphaned    float64 `json:"orphaned"`
}

type userBalance struct {
	Confirmed   strFloat64 `json:"confirmed"`
	Unconfirmed strFloat64 `json:"unconfirmed"`
	Orphaned    strFloat64 `json:"orphaned"`
}

func (s *UserBalance) UnmarshalJSON(data []byte) error {
	var us userBalance
	if err := json.Unmarshal(data, &us); err != nil {
		return err
	}
	s.Confirmed = float64(us.Confirmed)
	s.Unconfirmed = float64(us.Unconfirmed)
	s.Orphaned = float64(us.Orphaned)
	return nil
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
	req := &MposRequest{Page: "api", Action:"getuserbalance", Apikey:client.apikey, UserId:client.userid}
	_, err := client.sling.New().Get("").QueryStruct(req).ReceiveSuccess(&userbalance)
	if err != nil {
		return userbalance.Result.Data, err
	}

	return userbalance.Result.Data, err
}

