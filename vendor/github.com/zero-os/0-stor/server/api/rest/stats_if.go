package rest

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

// StatsInterface is interface for /stats root endpoint
type StatsInterface interface { // GetStoreStats is the handler for GET /stats
	// Return usage statistics about the whole store
	GetStoreStats(http.ResponseWriter, *http.Request)
}

// StatsInterfaceRoutes is routing for /stats root endpoint
func StatsInterfaceRoutes(r *mux.Router, i StatsInterface) {
	r.Handle("/stats", alice.New().Then(http.HandlerFunc(i.GetStoreStats))).Methods("GET")
}
