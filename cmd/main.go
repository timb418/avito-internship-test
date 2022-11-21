package main

import (
	"avito-internship-test/internal/dbl"

	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	dbl.InitDB()

	http.HandleFunc("/profit", dbl.Profit)
	http.HandleFunc("/user-balance", dbl.GetUserBalance)
	http.HandleFunc("/reserve", dbl.ReserveMoney)
	http.HandleFunc("/acknowledge", dbl.Acknowledge)

	http.ListenAndServe(":8090", nil)
}
