package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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

func HandleGetStock(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application.json")

	// Get the id as query param from http request.
	id := mux.Vars(request)["id"]
	fmt.Println("Id: ", id)

	// Id converted to int64
	convertedId, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println("ConvertedId: ", convertedId)

	stock, err := getStock(convertedId)

	if err != nil {
		response := response{
			Message: "Something went wrong on the server.",
			Code:    http.StatusBadGateway,
		}
		json.NewEncoder(writer).Encode(response)
		return
	}

	// If no stock was found
	if stock == (models.Stock{}) {
		response := response{
			Message: "Unable to find stock with the id " + id,
			Code:    http.StatusNotFound,
		}
		json.NewEncoder(writer).Encode(response)
		return
	}

	// Success flow
	response := response{
		Message: fmt.Sprintf("Successfully got stock of id: %v", stock.Id),
		Code:    http.StatusOK,
		Data:    stock,
	}
	json.NewEncoder(writer).Encode(response)
}

func HandleGetAllStocks(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application.json")

	stocks, err := getAllStocks()
	if err != nil {
		response := response{
			Message: "Something went wrong on the server.",
			Code:    http.StatusBadGateway,
		}
		json.NewEncoder(writer).Encode(response)
		return
	}

	// Success flow.
	response := response{
		Message: "Successfully got all stocks",
		Code:    http.StatusOK,
		Data:    stocks,
	}
	json.NewEncoder(writer).Encode(response)
}

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
		log.Printf("Unable to execute the query. %v", err)
		return 0, err
	}

	fmt.Printf("Inserted a single record with id: %v", id)
	return id, nil
}

func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM stock WHERE stock.id=$1`

	stock := models.Stock{}

	err := db.QueryRow(sqlStatement, id).Scan(
		&stock.Id,
		&stock.Name,
		&stock.Price,
		&stock.Company,
	)

	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return models.Stock{}, nil
	}

	fmt.Printf("Successfully got the stock with id: %v", id)
	return stock, nil
}

func getAllStocks() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stocks = make([]models.Stock, 0)

	sqlStatement := `SELECT * FROM stock`

	// Send the sql statement, return *Rows
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Printf("Unable to execute the query, %v", err)
		return stocks, err
	}

	defer rows.Close()

	// Iterate over each row
	for rows.Next() {
		stockFromDb := models.Stock{}

		err = rows.Scan(
			&stockFromDb.Id,
			&stockFromDb.Name,
			&stockFromDb.Price,
			&stockFromDb.Company,
		)

		if err != nil {
			log.Printf("Unable to scan the row. %v", err)
			return []models.Stock{}, err
		}

		stocks = append(stocks, stockFromDb)
	}

	return stocks, nil
}
