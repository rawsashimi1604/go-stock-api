package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rawsashimi1604/go-stock-api/router"
)

func main() {

	port := 8080
	fmt.Printf("Starting the server on port %v.\n", port)
	router := router.NewRouter()

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		panic(err)
	}
}
