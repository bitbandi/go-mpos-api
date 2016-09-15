package mpos

type PoolInfo struct {
	Currency  string `json:"currency"`
	CoinName  string `json:"coinname"`
	CoinTarget  string `json:"cointarget"`
	CoinDiffChangeTarget int `json:"coindiffchangetarget"`
	Algorithm  string `json:"algorithm"`
	StratumPort  int `json:"stratumport,string"`
	PayoutSystem  string `json:"payout_system"`
	Confirmations int `json:"confirmations"`
	MinAPThreshold int `json:"min_ap_threshold"`
	MaxAPThreshold int `json:"max_ap_threshold"`
	RewardType  string `json:"reward_type"`
	Reward int `json:"reward"`
	TxFee float64 `json:"txfee"`
	TxFeeManual float64 `json:"txfee_manual"`
	TxFeeAuto float64 `json:"txfee_auto"`
	Fees float64 `json:"fees"`
}

type poolInfoResponse struct {
	Result struct {
		       Version string     `json:"version"`
		       Runtime float32    `json:"runtime"`
		       Data    PoolInfo `json:"data"`
	       } `json:"getpoolinfo"`
}

func (client *MposClient) GetPoolInfo() (PoolInfo, error) {
	poolinfo := poolInfoResponse{}
	req := &MposRequest{Page: "api", Action:"getpoolinfo", Apikey:client.apikey}
	_, err := client.sling.New().Get("index.php").QueryStruct(req).ReceiveSuccess(&poolinfo)
	if err != nil {
		return poolinfo.Result.Data, err
	}

	return poolinfo.Result.Data, err
}

