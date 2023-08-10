package router

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/go-stock-api/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", middleware.HandleIndex).Methods("GET")
	router.HandleFunc("/stock/all", middleware.HandleGetAllStocks).Methods("GET")
	router.HandleFunc("/stock/{id}", middleware.HandleGetStock).Methods("GET")
	router.HandleFunc("/stock", middleware.HandleCreateStock).Methods("POST")
	router.HandleFunc("/stock/{id}", middleware.HandleDeleteStock).Methods("DELETE")
	router.HandleFunc("/stock/{id}", middleware.HandleUpdateStock).Methods("PUT", "PATCH")
	return router
}
