package main

import (
	"github.com/jeromefurog/trader-svc/ticker"
	"github.com/jeromefurog/trader-svc/rest"
	"fmt"
)


func main() {

	fmt.Println("Starting trading bot...")
	go ticker.Run()
	rest.StartServer()
}



