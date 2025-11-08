package client

import (
	"net/http"
	"sync"

	polymarketgamma "github.com/ivanzzeth/polymarket-go-gamma-client"
)

var (
	gammaClient     *polymarketgamma.Client
	gammaClientOnce sync.Once
)

func GetGammaClient() *polymarketgamma.Client {
	gammaClientOnce.Do(func() {
		gammaClient = polymarketgamma.NewClient(http.DefaultClient)
	})
	return gammaClient
}
