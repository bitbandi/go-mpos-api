package mpos

type PoolStatus struct {
	Poolname            string `json:"pool_name"`
	Hashrate            float64 `json:"hashrate"`
	Efficiency          float64 `json:"efficiency"`
	Progress            float64 `json:"progress"`
	Workers             uint32 `json:"workers"`
	CurrentNetworkBlock uint32 `json:"currentnetworkblock"`
	NextNetworkBlock    uint32 `json:"nextnetworkblock"`
	LastBlock           uint32 `json:"lastblock"`
	NetworkDiff         float64 `json:"networkdiff"`
	EstTime             float64 `json:"esttime"`
	EstShares           float64 `json:"estshares"`
	TimeSinceLast       uint32 `json:"timesincelast"`
	NetHashRate         float64 `json:"nethashrate"`
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
	_, err := client.sling.New().Get("").QueryStruct(req).ReceiveSuccess(&poolstatus)
	if err != nil {
		return poolstatus.Result.Data, err
	}

	return poolstatus.Result.Data, err
}

