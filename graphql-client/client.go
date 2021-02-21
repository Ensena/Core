package graphql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Ensena/core/middleware"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmhttp"
)

type payload struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

var url, auth string

func init() {
	url = os.Getenv("graphql_url")
	auth = os.Getenv("graphql_auth")
	if url == "" {
		log.Fatalln("You need env: graphql_url to connect a server")
	}
	if auth == "" {
		log.Fatalln("You need env: graphql_auth to connect a server")
	}

}

func Query(ctx *gin.Context, graph string) ([]byte, error) {
	//client := &http.Client{}

	client := apmhttp.WrapClient(http.DefaultClient)
	// If “ctx” contains a reference to an in-progress transaction, then the call below will start a new span.

	p := payload{}
	p.Query = graph
	byteData, _ := json.Marshal(p)
	req, err := http.NewRequest("POST", url, bytes.NewReader(byteData))
	if err != nil {
		log.Fatal(err)
	}

	tx := middleware.GetTransaction(ctx.Request)
	middleware.SetTransaction(req, tx)
	req.Header.Set("authorization", "Basic "+auth)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	//resp, err := client.Do(req)ctx
	resp, err := client.Do(req.WithContext(ctx.Request.Context()))

	if err != nil {
		log.Println("Error Connect to GraphqlServer", err)

	}
	return ioutil.ReadAll(resp.Body)
}
