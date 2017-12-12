package ticker

import (
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	//"github.com/jeromefurog/trader-svc/workflow"
	"github.com/jeromefurog/trader-svc/poc"
	"fmt"
)

var client *bitfinex.Client

func Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Run():", r)
			go Run()
		}
	}()
	client = bitfinex.NewClient()
	// in case your proxy is using a non valid certificate set to TRUE
	client.WebSocketTLSSkipVerify = false

	err := client.WebSocket.Connect()
	if err != nil {
		//log.Fatal("Error connecting to web socket : ", err)
		panic(err)
	}
	defer client.WebSocket.Close()

	ticker_chan := make(chan []float64)
	client.WebSocket.AddSubscribe(bitfinex.ChanTicker, bitfinex.IOTUSD, ticker_chan)

	go listen(ticker_chan)

	err = client.WebSocket.Subscribe()
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}

}

func listen(in chan []float64) {
	for {
		msg := <-in
		price := msg[0]
		//workflow.Process(price)
		poc.ProcessOrder(price)
	}
}