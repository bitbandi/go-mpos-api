package mpos

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserWorkers(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getuserworkers":{
			    "version":"1.0.0",
			    "runtime":4.0218830108643,
			      "data":[
			      {
			        "id":138,
			        "username":"test.1",
			        "password":"x",
			        "monitor":1,
			        "count_all":0,
			        "count_all_archive":0,
			        "hashrate":0,
			        "difficulty":0
			      },
			      {
			        "id":139,
			        "username":"test.2",
			        "password":"x",
			        "monitor":1,
			        "count_all":0,
			        "count_all_archive":0,
			        "hashrate":0,
			        "difficulty":0
			      }
			    ]
			  }
			}`

	expectedItem := []UserWorkers{
		UserWorkers{
			Id:138,
			Username:"test.1",
			Password:"x",
			Monitor:1,
			CountAll:0,
			CountAllArchive:0,
			Hashrate:0,
			Difficulty:0},
		UserWorkers{
			Id:139,
			Username:"test.2",
			Password:"x",
			Monitor:1,
			CountAll:0,
			CountAllArchive:0,
			Hashrate:0,
			Difficulty:0},
	}

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getuserworkers", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY", "")
	userworkers, err := mposClient.GetUserWorkers()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, userworkers)
}