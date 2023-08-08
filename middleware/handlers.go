package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	msg := "Hello world from index! Path: /"
	fmt.Println(msg)

	response := Response{
		Message: msg,
		Code:    200,
	}

	writer.Header().Set("Content-Type", "application.json")
	json.NewEncoder(writer).Encode(response)
}
