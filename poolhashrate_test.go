package mpos

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPoolHashrate(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getpoolhashrate":{
			    "version":"1.0.0",
			    "runtime":4.1780471801758,
			    "data":260903.18506667
			  }
			}`

	expectedItem := 260903.18506667

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getpoolhashrate", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	poolhashrate, err := mposClient.GetPoolHashrate()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, poolhashrate)
}
