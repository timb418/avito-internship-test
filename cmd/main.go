package main

import (
	"avito-internship-test/internal/dbl"

	//_ "github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "habrpguser"
	password = "pgpwd4habr"
	dbname   = "habrdb"
)

func main() {
	dbl.InitDB()

	//http.HandleFunc("/hello", dbl.Hello)
	http.HandleFunc("/profit", dbl.Profit)
	http.HandleFunc("/user-balance", dbl.GetUserBalance)
	http.HandleFunc("/reserve", dbl.ReserveMoney)
	//http.HandleFunc("/user-balance-test", dbl.GetUserBalanceTest)

	http.ListenAndServe(":8090", nil)
}
