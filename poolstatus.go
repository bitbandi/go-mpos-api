package mpos

type PoolStatus struct {
	Poolname            string `json:"pool_name"`
	Hashrate            float64 `json:"hashrate"`
	Efficiency          float64 `json:"efficiency"`
	Progress            float64 `json:"progress"`
	Workers             int `json:"workers"`
	CurrentNetworkBlock int `json:"currentnetworkblock"`
	NextNetworkBlock    int `json:"nextnetworkblock"`
	LastBlock           int `json:"lastblock"`
	NetworkDiff         float64 `json:"networkdiff"`
	EstTime             float64 `json:"esttime"`
	EstShares           int `json:"estshares"`
	TimeSinceLast       int `json:"timesincelast"`
	NetHashRate         int `json:"nethashrate"`
}

type poolStatusResponse struct {
	Result struct {
		       Version string     `json:"version"`
		       Runtime float32    `json:"runtime"`
		       Data    PoolStatus `json:"data"`
	       } `json:"getpoolstatus"`
}

func (client *MposClient) GetPoolStatus() (PoolStatus, error) {
	poolstatus := poolStatusResponse{}
	req := &MposRequest{Page: "api", Action:"getpoolstatus", Apikey:client.apikey}
	_, err := client.sling.New().Get("index.php").QueryStruct(req).ReceiveSuccess(&poolstatus)
	if err != nil {
		return poolstatus.Result.Data, err
	}

	return poolstatus.Result.Data, err
}

