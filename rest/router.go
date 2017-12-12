package rest

import (
	"time"
	"net/http/pprof"
	"log"
	"fmt"

	sw "github.com/werpas/ui-swagger"
	"github.com/gorilla/mux"
)

const base_path = "/api/v1/"

func attachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}

func getRouter() (router *mux.Router) {
	startTime = time.Now().UTC()

	log.Println(fmt.Sprintf("Configuring Base REST Server Interface at %v", startTime.Format(time.RFC850)))

	// Register a handler for each route pattern
	router = mux.NewRouter()

	// Add a trivial handler for INFO
	attachProfiler(router)

	// attach swagger documentation api
	sw.AttachSwaggerUI(router, base_path)

	//  standard endpoints
	api := router.PathPrefix(base_path).Subrouter()

	//  these should not require authentication to get results
	api.Path("/info").Methods("GET").HandlerFunc(HandleInfo)
	api.Path("/ping").Methods("GET").HandlerFunc(HandlePing)

	// Custom REST handlers
	for _, route := range routes {
		api.Path(route.Pattern).Methods(route.Method).Handler(route.HandlerFunc)
	}

	return
}

