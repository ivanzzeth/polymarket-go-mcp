package client

import (
	"net/http"
	"sync"

	polymarketdata "github.com/ivanzzeth/polymarket-go-data-client"
	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
)

var (
	gammaClient     *polymarketgamma.Client
	gammaClientOnce sync.Once

	dataClient     *polymarketdata.DataClient
	dataClientOnce sync.Once
)

func GetGammaClient() *polymarketgamma.Client {
	gammaClientOnce.Do(func() {
		gammaClient = polymarketgamma.NewClient(http.DefaultClient)
	})
	return gammaClient
}

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
