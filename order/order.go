package order

import (
	"fmt"
	"strings"

	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/jeromefurog/trader-svc/config"
)

func GetLiveIOTUSDOrder() (order bitfinex.Order, err error) {

	ordersAll, err := getLiveOrders()
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: GetLiveIOTUSDOrder() - %v", err.Error()))
		return
	}

	for _, o := range ordersAll {
		if strings.EqualFold(bitfinex.IOTUSD, o.Symbol) && o.IsLive && strings.EqualFold(o.Type, bitfinex.OrderTypeExchangeLimit) {
			order = o
			return
		}
	}

	return
}

func GetOrderByID(id int64) (order bitfinex.Order, err error) {
	client := bitfinex.NewClient().Auth(config.API_KEY, config.API_SECRET)

	order, err = client.Orders.Status(id)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: GetOrderByID() - %v", err.Error()))
	}
	return
}

func getLiveOrders() (orders []bitfinex.Order, err error) {
	client := bitfinex.NewClient().Auth(config.API_KEY, config.API_SECRET)

	orders, err = client.Orders.All()
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: getLiveOrders() - %v", err.Error()))
	}
	return
}

func CreateSellOrder(amount, price float64) (order *bitfinex.Order, err error) {
	client := bitfinex.NewClient().Auth(config.API_KEY, config.API_SECRET)

	// Sell 0.01BTC at $12.000
	order, err = client.Orders.Create(bitfinex.IOTUSD, -1 * amount, price, bitfinex.OrderTypeExchangeLimit)

	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: CreateSellOrder() - %v", err.Error()))
	} else {
		fmt.Println("Created sell order: ", order)
	}

	return
}

func CreateBuyOrder(amount, price float64) (order *bitfinex.Order, err error) {
	client := bitfinex.NewClient().Auth(config.API_KEY, config.API_SECRET)

	// Sell 0.01BTC at $12.000
	order, err = client.Orders.Create(bitfinex.IOTUSD, amount, price, bitfinex.OrderTypeExchangeLimit)

	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: CreateBuyOrder() - %v", err.Error()))
	} else {
		fmt.Println("Created buy order: ", order)
	}

	return
}


func GetBalance() ([]bitfinex.WalletBalance, error) {
	client := bitfinex.NewClient().Auth(config.API_KEY, config.API_SECRET)

	return client.Balances.All()
}