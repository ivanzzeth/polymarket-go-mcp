package client

import (
	"net/http"
	"sync"

	polymarketdata "github.com/ivanzzeth/polymarket-go-data-client"
)

var (
	dataClient     *polymarketdata.DataClient
	dataClientOnce sync.Once
)

func GetDataClient() *polymarketdata.DataClient {
	dataClientOnce.Do(func() {
		client, err := polymarketdata.NewDataClient(http.DefaultClient)
		if err != nil {
			panic(err)
		}

		dataClient = client
	})
	return dataClient
}
