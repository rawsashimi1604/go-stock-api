package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rawsashimi1604/go-stock-api/middleware"
	"github.com/rawsashimi1604/go-stock-api/router"
)

func main() {
	fmt.Println("Starting the server.")
	port := 8080
	router := router.NewRouter()

	db := middleware.CreateConnection()
	db.Ping()

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		panic(err)
	}
}
