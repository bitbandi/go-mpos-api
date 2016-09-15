package mpos

type poolHashrateResponse struct {
	Result struct {
		       Version string     `json:"version"`
		       Runtime float32    `json:"runtime"`
		       Data    float64 `json:"data"`
	       } `json:"getpoolhashrate"`
}

func (client *MposClient) GetPoolHashrate() (float64, error) {
	poolhashrate := poolHashrateResponse{}
	req := &MposRequest{Page: "api", Action:"getpoolhashrate", Apikey:client.apikey}
	_, err := client.sling.New().Get("index.php").QueryStruct(req).ReceiveSuccess(&poolhashrate)
	if err != nil {
		return poolhashrate.Result.Data, err
	}

	return poolhashrate.Result.Data, err
}
