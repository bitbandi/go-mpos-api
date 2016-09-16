package mpos

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserStatus(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getuserstatus":{
			    "version":"1.0.0",
			    "runtime":4.3489933013916,
			    "data":{
			      "username":"test",
			      "shares":{
			        "valid":0,
			        "invalid":0,
			        "donate_percent":0,
			        "is_anonymous":0
			      },
			      "hashrate":0,
			      "sharerate":0
			    }
			  }
			}`

	expectedItem := UserStatus{
		Username:"test",
		Shares:UserShares{
			Valid:0,
			Invalid:0,
			DonatePercent:0,
			IsAnonymous:0,
		},
		Hashrate:0,
		Sharerate:0}

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getuserstatus", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	userstatus, err := mposClient.GetUserStatus()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, userstatus)
}