package mpos

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestGetUserTransactions(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			  "getusertransactions":{
			    "version":"1.0.0",
			    "runtime":14.245986938477,
			    "data":{
			      "transactions":[
			        {
			          "id":12135254,
			          "username":"test",
			          "type":"Bonus",
			          "amount":0.6,
			          "coin_address":null,
			          "timestamp":"2015-05-28 19:00:12",
			          "txid":null,
			          "height":283126,
			          "blockhash":"1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014",
			          "confirmations":120
			        },
			        {
			          "id":12135253,
			          "username":"test",
			          "type":"Fee",
			          "amount":0.00045155,
			          "coin_address":null,
			          "timestamp":"2015-05-28 19:00:12",
			          "txid":null,
			          "height":283126,
			          "blockhash":"1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014",
			          "confirmations":120
			        },
			        {
			          "id":12135252,
			          "username":"test",
			          "type":"Credit",
			          "amount":0.02257736,
			          "coin_address":null,
			          "timestamp":"2015-05-28 19:00:12",
			          "txid":null,
			          "height":283126,
			          "blockhash":"1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014",
			          "confirmations":120
			        },
			        {
			          "id":553639,
			          "username":"test",
			          "type":"TXFee",
			          "amount":0.1,
			          "coin_address":"TESTADDRESS",
			          "timestamp":"2014-07-27 21:03:55",
			          "txid":null,
			          "height":null,
			          "blockhash":null,
			          "confirmations":null
			        },
			        {
			          "id":553638,
			          "username":"test",
			          "type":"Debit_AP",
			          "amount":4575.1699014,
			          "coin_address":"TESTADDRESS",
			          "timestamp":"2014-07-27 21:03:55",
			          "txid":"60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752",
			          "height":null,
			          "blockhash":null,
			          "confirmations":null
			        }
			      ],
			      "transactionsummary":{
			        "Bonus":1.8000015,
			        "Credit":4669.39744585,
			        "Debit_AP":4575.1699014,
			        "Fee":93.38794898,
			        "TXFee":0.1
			      }
			    }
			  }
			}`

	expectedItem := UserTransactions{
		Transactions:[]UserTransaction{
			UserTransaction{
				Id:12135254,
				Username:"test",
				Type:"Bonus",
				Amount:0.6,
				CoinAddress:"",
				Timestamp:TransactionTimestamp(time.Date(2015, 5, 28, 19, 00, 12, 0, time.UTC)),
				TxId:"",
				Height:283126,
				BlockHash:"1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014",
				Confirmations:120,
			},
			UserTransaction{
				Id:12135253,
				Username:"test",
				Type:"Fee",
				Amount:0.00045155,
				CoinAddress:"",
				Timestamp:TransactionTimestamp(time.Date(2015, 5, 28, 19, 00, 12, 0, time.UTC)),
				TxId:"",
				Height:283126,
				BlockHash:"1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014",
				Confirmations:120,
			},
			UserTransaction{
				Id:12135252,
				Username:"test",
				Type:"Credit",
				Amount:0.02257736,
				CoinAddress:"",
				Timestamp:TransactionTimestamp(time.Date(2015, 5, 28, 19, 0, 12, 0, time.UTC)),
				TxId:"",
				Height:283126,
				BlockHash:"1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014",
				Confirmations:120},
			UserTransaction{
				Id:553639,
				Username:"test",
				Type:"TXFee",
				Amount:0.1,
				CoinAddress:"TESTADDRESS",
				Timestamp:TransactionTimestamp(time.Date(2014, 7, 27, 21, 3, 55, 0, time.UTC)),
				TxId:"",
				Height:0,
				BlockHash:"",
				Confirmations:0},
			UserTransaction{
				Id:553638,
				Username:"test",
				Type:"Debit_AP",
				Amount:4575.1699014,
				CoinAddress:"TESTADDRESS",
				Timestamp:TransactionTimestamp(time.Date(2014, 7, 27, 21, 3, 55, 0, time.UTC)),
				TxId:"60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752",
				Height:0,
				BlockHash:"",
				Confirmations:0},
		},
		TransactionSummary:UserTransactionSummary{
			Bonus:1.8000015,
			Credit:4669.39744585,
			DebitAP:4575.1699014,
			Fee:93.38794898,
			TXFee:0.1,
		}}

	mux.HandleFunc("/index.php", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "api", r.URL.Query().Get("page"))
		assert.Equal(t, "getusertransactions", r.URL.Query().Get("action"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("api_key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	mposClient := NewMposClient(httpClient, "http://dummy.com/", "FAKEKEY")
	usertransactions, err := mposClient.GetUserTransactions()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, usertransactions)
}