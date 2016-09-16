package mpos

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPoolInfo(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getpoolinfo":{
			    "version":"1.0.0",
			    "runtime":2.830982208252,
			    "data":{
			      "currency":"TEST",
			      "coinname":"TestCoin",
			      "cointarget":"150",
			      "coindiffchangetarget":2016,
			      "algorithm":"scrypt",
			      "stratumport":"3333",
			      "payout_system":"pplns",
			      "confirmations":120,
			      "min_ap_threshold":1,
			      "max_ap_threshold":500000,
			      "reward_type":"block",
			      "reward":50,
			      "txfee":0.1,
			      "txfee_manual":0.1,
			      "txfee_auto":0.1,
			      "fees":2
			    }
			  }
			}`

	expectedItem := PoolInfo{
		Currency:"TEST",
		CoinName:"TestCoin",
		CoinTarget:"150",
		CoinDiffChangeTarget:2016,
		Algorithm:"scrypt",
		StratumPort:3333,
		PayoutSystem:"pplns",
		Confirmations:120,
		MinAPThreshold:1,
		MaxAPThreshold:500000,
		RewardType:"block",
		Reward:50,
		TxFee:0.1,
		TxFeeManual:0.1,
		TxFeeAuto:0.1,
		Fees:2}

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getpoolinfo", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	poolinfo, err := mposClient.GetPoolInfo()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, poolinfo)
}
