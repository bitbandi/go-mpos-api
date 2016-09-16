package mpos

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPoolStatus(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getpoolstatus":{
			    "version":"1.0.0",
			    "runtime":21.305799484253,
			    "data":{
			      "pool_name":"Test Pool",
			      "hashrate":259634.40810667,
			      "efficiency":99.45,
			      "progress":291.97,
			      "workers":5,
			      "currentnetworkblock":623687,
			      "nextnetworkblock":623688,
			      "lastblock":623684,
			      "networkdiff":10.67583674,
			      "esttime":176.60359422353,
			      "estshares":699652,
			      "timesincelast":543,
			      "nethashrate":418239673
			    }
			  }
			}`

	expectedItem := PoolStatus{
		Poolname:"Test Pool",
		Hashrate:259634.40810667,
		Efficiency:99.45,
		Progress:291.97,
		Workers:5,
		CurrentNetworkBlock:623687,
		NextNetworkBlock:623688,
		LastBlock:623684,
		NetworkDiff:10.67583674,
		EstTime:176.60359422353,
		EstShares:699652,
		TimeSinceLast:543,
		NetHashRate:418239673,
	}

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getpoolstatus", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	poolstatus, err := mposClient.GetPoolStatus()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, poolstatus)
}
