package mpos

import (
	"encoding/json"
	"strconv"
)

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

func (s *PoolStatus) UnmarshalJSON(data []byte) error {
	type Alias PoolStatus
	aux := &struct {
		Hashrate interface{} `json:"hashrate,string"`
		Workers  interface{} `json:"workers,string"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch v := aux.Hashrate.(type) {
	case string:
		h, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		s.Hashrate = h
	case uint32:
	case uint64:
		s.Hashrate = float64(v)
	case float32:
	case float64:
		s.Hashrate = v
	case bool:
		s.Hashrate = 0.0
	default:
		s.Hashrate = 0.0
	}
	switch v := aux.Workers.(type) {
	case string:
		w, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		s.Workers = uint32(w)
	case uint32:
		s.Workers = v
	case float32:
	case float64:
		s.Workers = uint32(v)
	case bool:
		s.Workers = 0
	default:
		s.Workers = 0
	}
	return nil
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

