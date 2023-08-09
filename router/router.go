package router

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/go-stock-api/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", middleware.HandleIndex).Methods("GET")
	router.HandleFunc("/stock", middleware.HandleCreateStock).Methods("POST")
	return router
}
