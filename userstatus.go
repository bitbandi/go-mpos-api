package mpos

import (
	"strconv"
	"encoding/json"
)

type strFloat64 float64

func (s *strFloat64) UnmarshalJSON(data []byte) error {
	var v float64
	var err error
	if len(data) == 0 {
		*s = 0
		return nil
	}

	if data[0] != '"' || data[len(data) - 1] != '"' {
		v, err = strconv.ParseFloat(string(data), 64)
	} else {
		v, err = strconv.ParseFloat(string(data[1:len(data) - 1]), 64)
	}
	if err != nil {
		return err
	}

	*s = strFloat64(v)
	return nil
}

func (s *strFloat64) Float64() float64 {
	return float64(*s)
}

type UserStatus struct {
	Username  string     `json:"username"`
	Shares    UserShares `json:"shares"`
	Hashrate  float64        `json:"hashrate"`
	Sharerate float64        `json:"sharerate"`
}

// Used to avoid recursion in UnmarshalJSON below.
type userStatus struct {
	Username  string     `json:"username"`
	Shares    UserShares `json:"shares"`
	Hashrate  float64        `json:"hashrate"`
	Sharerate strFloat64        `json:"sharerate"`
}

func (s *UserStatus) UnmarshalJSON(data []byte) error {
	var us userStatus
	if err := json.Unmarshal(data, &us); err != nil {
		return err
	}
	s.Username = us.Username
	s.Shares = us.Shares
	s.Hashrate = us.Hashrate
	s.Sharerate = float64(us.Sharerate)
	return nil
}

type UserShares struct {
	Valid         float64 `json:"valid"`
	Invalid       float64 `json:"invalid"`
	DonatePercent strFloat64 `json:"donate_percent"`
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
	req := &MposRequest{Page: "api", Action:"getuserstatus", Apikey:client.apikey, UserId:client.userid}
	_, err := client.sling.New().Get("").QueryStruct(req).ReceiveSuccess(&userstatus)
	if err != nil {
		return userstatus.Result.Data, err
	}

	return userstatus.Result.Data, err
}
