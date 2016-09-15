package mpos

import (
	"github.com/dghubble/sling"
	"net/http"
	"net/http/httputil"
	"log"
)

type MposClient struct {
	sling  *sling.Sling
	apikey string
}

// mpos send the api response with text/html content type
// we fix this: change content type to json
type fixedHttpClient struct {
	client *http.Client
}

func (d fixedHttpClient) Do(req *http.Request) (*http.Response, error) {
	//d.dumpRequest(req)
	resp, err := func() (*http.Response, error) { if d.client != nil { return d.client.Do(req) } else { return http.DefaultClient.Do(req) } }()
	//d.dumpResponse(resp)
	if err == nil {
		if resp.Header.Get("Content-Type") == "text/html" {
			resp.Header.Set("Content-Type", "application/json")
		}
	}
	return resp, err
}

func (d fixedHttpClient) dumpRequest(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (d fixedHttpClient) dumpResponse(r *http.Response) {
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

func NewMposClient(client *http.Client, BaseURL string, ApiToken string) *MposClient {
	fixedclient := &fixedHttpClient{client:client}
	return &MposClient{
		sling: sling.New().Doer(fixedclient).Base(BaseURL),
		apikey: ApiToken,
	}
}

type MposRequest struct {
	Page   string `url:"page"`
	Action string `url:"action"`
	Apikey string `url:"api_key"`
}
