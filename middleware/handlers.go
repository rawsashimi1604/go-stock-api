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
	"github.com/rawsashimi1604/go-stock-api/models"
)

type response struct {
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func createConnection() *sql.DB {
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
		Code:    http.StatusOK,
	}

	writer.Header().Set("Content-Type", "application.json")
	json.NewEncoder(writer).Encode(response)
}

// func HandleGetAllStocks(writer http.ResponseWriter, request http.Request) {
// 	stocks := make([]models.Stock, 0)

// }

func HandleCreateStock(writer http.ResponseWriter, request *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(request.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	stock.Id, err = insertStock(stock)
	if err != nil {
		log.Print("Something went wrong.")
	}

	response := response{
		Message: fmt.Sprintf("Created stock with the ID: %v", stock.Id),
		Code:    http.StatusCreated,
		Data:    stock,
	}

	writer.Header().Set("Content-Type", "application.json")
	json.NewEncoder(writer).Encode(response)
}

// ------------------------- handler functions ----------------
func insertStock(stock models.Stock) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stock (name, price, company) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Print("Unable to execute the query. %v", err)
		return 0, err
	}

	fmt.Printf("Inserted a single record with id: %v", id)
	return id, nil
}
