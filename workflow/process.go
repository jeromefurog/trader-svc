package workflow

import (
	"fmt"
	"github.com/jeromefurog/trader-svc/order"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"reflect"
	"strconv"
)

var orderID int64

func Process(price float64) {
	fmt.Println(fmt.Sprintf("IOTUSD Price: $%v", price))
	ord, err := order.GetLiveIOTUSDOrder()
	if err != nil {
		return
	}

	isOrderEmpty := reflect.DeepEqual(ord, bitfinex.Order{})
	if isOrderEmpty && orderID == 0 {
		fmt.Println("ERROR: Process() - No initial order to process.")
		return
	}

	bal, _ := order.GetBalance()
	fmt.Println(bal)
	return

	ord, err = order.GetOrderByID(6012389870)
	if err != nil {
		return
	}

	if isOrderEmpty && orderID > 0 {
		ord, err = order.GetOrderByID(6012389870)
		if err != nil {
			return
		}

		orderPricePrev, err := strconv.ParseFloat(ord.Price, 64)
		if err != nil {
			fmt.Println(fmt.Sprintf("ERROR: Process() - error converting price string. %v", err.Error()))
			return
		}

		var orderPrice float64
		if ord.Side == "sell" {
			if price < orderPricePrev {
				orderPrice = price - (price * 0.015)
			} else {
				orderPrice = orderPricePrev - (orderPricePrev * 0.015)
			}

			ords, _ := order.CreateBuyOrder(0.0, orderPrice)
			orderID = ords.ID
		} else {
			if price > orderPricePrev {
				orderPrice = price + (price * 0.015)
			} else {
				orderPrice = orderPricePrev + (orderPricePrev * 0.015)
			}

			ords, _ := order.CreateSellOrder(0.0, orderPrice)
			orderID = ords.ID
		}

	} else {
		orderID = ord.ID
	}

	fmt.Println(ord)
}
