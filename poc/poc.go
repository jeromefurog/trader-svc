package poc

import (
	"fmt"
	"reflect"
	"time"
)

var marginPercentage float64
var usdFund float64
var iotaFund float64
var bidPrice float64
var update string

var orderIdCount int64
var currentOrder Order
var loc *time.Location

type Order struct {
	orderID int64
	price float64
	iotaAmount float64
	usdAmount float64
	isBuy bool
}

func init() {
	usdFund = 100.0 // $100
	marginPercentage = 0.011
	bidPrice = 4.2

	loc, _ = time.LoadLocation("Asia/Shanghai")
}

func GetUpdate() string {
	return update
}

func ProcessOrder(currentPrice float64) {
	var orderPrice float64
	// first order
	if reflect.DeepEqual(currentOrder, Order{}) {
		orderPrice = bidPrice - (bidPrice * marginPercentage)
		currentOrder = buyOrder(usdFund, orderPrice)
	}

	if orderStatus(currentOrder, currentPrice) {

		if currentOrder.isBuy {
			if currentPrice > currentOrder.price {
				orderPrice = currentPrice + (currentPrice * marginPercentage)
			} else {
				orderPrice = currentOrder.price + (currentOrder.price * marginPercentage)
			}

			currentOrder = sellOrder(iotaFund, orderPrice)
		} else {

			if currentPrice < currentOrder.price {
				orderPrice = currentPrice - (currentPrice * marginPercentage)
			} else {
				orderPrice = currentOrder.price - (currentOrder.price * marginPercentage)
			}

			currentOrder = buyOrder(usdFund, orderPrice)
		}
	}

	update = fmt.Sprintf("IOTA Price: %v, USD FUND: %v, IOTA FUND: %v, Current Order: %v, IsBuy: %v, Price: %v, Time: %v", currentPrice, usdFund, iotaFund, currentOrder.orderID, currentOrder.isBuy, currentOrder.price, time.Now().In(loc))
	fmt.Println(update)
}

func orderStatus(ord Order, currPrice float64) bool {

	if ord.isBuy {
		if currPrice <= ord.price {
			iotaFund += ord.iotaAmount
			usdFund += ord.usdAmount
			fmt.Println(fmt.Sprintf("BUY ORDER EXECUTED - ID: %v, Price: %v, USD: %v, IOTA: %v", ord.orderID, ord.price, ord.usdAmount, ord.iotaAmount))
			return true
		}
	} else {
		if currPrice >= ord.price {
			iotaFund += ord.iotaAmount
			usdFund += ord.usdAmount
			fmt.Println(fmt.Sprintf("SELL ORDER EXECUTED - ID: %v, Price: %v, USD: %v, IOTA: %v", ord.orderID, ord.price, ord.usdAmount, ord.iotaAmount))
			return true
		}
	}

	return false
}

func buyOrder(usdAmount, price float64) (ord Order) {
	price = round(price, 0.005) // limit to 3 decimal places

	orderIdCount += 1
	ord.orderID = orderIdCount
	ord.iotaAmount = usdAmount / price
	ord.usdAmount = -1.0 * usdAmount
	ord.isBuy = true
	ord.price = price

	fmt.Println(fmt.Sprintf("BUY ORDER CREATED - ID: %v, Price: %v, USD: %v, IOTA: %v", ord.orderID, ord.price, ord.usdAmount, ord.iotaAmount))
	return
}

func sellOrder(iotaAmount, price float64) (ord Order) {
	price = round(price, 0.005) // limit to 3 decimal places

	orderIdCount += 1
	ord.orderID = orderIdCount
	ord.iotaAmount = -1.0 * iotaAmount
	ord.usdAmount = iotaAmount * price
	ord.isBuy = false
	ord.price = price

	fmt.Println(fmt.Sprintf("SELL ORDER CREATED - ID: %v, Price: %v, USD: %v, IOTA: %v", ord.orderID, ord.price, ord.usdAmount, ord.iotaAmount))
	return
}

func round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}
