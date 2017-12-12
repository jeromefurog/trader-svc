package rest

import (
	"fmt"
	"net/http"
	"github.com/jeromefurog/trader-svc/poc"
)

func HandleGetUpdates(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, poc.GetUpdate())
}
