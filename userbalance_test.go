package mpos

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserBalance(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getuserbalance":{
			    "version":"1.0.0",
			    "runtime":2.8622150421143,
			    "data":{
			      "confirmed":2.53959697,
			      "unconfirmed":1,
			      "orphaned":0
			    }
			  }
			}`

	expectedItem := UserBalance{
		Confirmed:2.53959697,
		Unconfirmed:1,
		Orphaned:0}

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getuserbalance", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY", 0, "")
	userbalance, err := mposClient.GetUserBalance()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, userbalance)
}

func TestGetUserBalanceStr(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getuserbalance":{
			    "version":"1.0.0",
			    "runtime":2.8622150421143,
			    "data":{
			      "confirmed": "2.53959697",
			      "unconfirmed": "1",
			      "orphaned":0
			    }
			  }
			}`

	expectedItem := UserBalance{
		Confirmed:2.53959697,
		Unconfirmed:1,
		Orphaned:0}

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getuserbalance", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY", 0, "")
	userbalance, err := mposClient.GetUserBalance()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, userbalance)
}
