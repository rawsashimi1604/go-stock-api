package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func CreateConnection() *sql.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Something went wrong when loading the .env file.")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Println(err)
		log.Fatal("Something went wrong when connecting to the db.")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to postgres.")
	return db
}

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	msg := "Hello world from index! Path: /"
	fmt.Println(msg)

	response := response{
		Message: msg,
	}

	writer.Header().Set("Content-Type", "application.json")
	json.NewEncoder(writer).Encode(response)
}
