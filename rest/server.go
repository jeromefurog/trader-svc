package rest

import (
	"net/http"
	"time"
	"log"
	"fmt"

	"github.com/gorilla/handlers"
)

const default_port = "8080"
var startTime time.Time

func StartServer() {

	startTime = time.Now().UTC()

	// Get Cloud Foundry assigned port
	port := default_port
	if port == "" {
		port = default_port
		log.Println(fmt.Sprintf("Port number is not set using default: %s\n", port))
	}

	// get router object
	router := getRouter()

	// Start listening on the configured port.
	// ListenAndServe is not expected to return, so we wrap it in a log.Fatal
	// also include CORS handling
	err:= http.ListenAndServe(":"+port, handlers.CORS()(router))
	if(err != nil){
		log.Fatal(err.Error())
	}
}
