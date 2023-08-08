package middleware

import (
	"fmt"
	"net/http"
)

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello world from index! Path: /")
}
