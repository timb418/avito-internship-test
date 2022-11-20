package dbl

import (
	"encoding/json"
	"fmt"
	"github.com/bojanz/currency"
	"log"
	"net/http"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	db, err := GetDBConnection()
	writer.Write([]byte("after connection"))
	if err == nil {
		writer.Write([]byte("inside err nil"))

		var version string
		_ = db.QueryRow("select version();").Scan(&version)
		fmt.Println(version)

		writer.Write([]byte(version))
	}
	writer.Write([]byte("after if"))
	InitDB()
}

type Balance struct {
	UserId      string
	MoneyAmount string
}
type Reserve struct {
	UserId      string
	ServiceId   string
	OrderId     string
	MoneyAmount string
}

func Profit(w http.ResponseWriter, r *http.Request) {
	var profit Balance
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&profit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = AddMoneyToUser(profit)
	if err != nil {
		log.Fatalln("Could not add money to a user")
		return
	}
	w.WriteHeader(http.StatusOK)
}
func ReserveMoney(w http.ResponseWriter, r *http.Request) {
	// Parse the reqauest
	var reserve Reserve
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&reserve)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO
	// Subtract money from the old balance
	oldBalance, err := GetUserBalanceFromDB(reserve.UserId)
	oldAmount, _ := currency.NewAmount(oldBalance.MoneyAmount, "RUB")
	moneyToSub, err := currency.NewAmount(reserve.MoneyAmount, "RUB")

	oldAmount.Sub(moneyToSub)

	//TODO
	// Add the money to a reserve

}
func Acknowledge(writer http.ResponseWriter, request *http.Request) {}
func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	var userId Balance
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	foundUserBalance, err := GetUserBalanceFromDB(userId.UserId)

	json.NewEncoder(w).Encode(foundUserBalance)
}

//func GetUserBalanceTest(w http.ResponseWriter, r *http.Request) {
//
//	w.Header().Set("Content-Type", "application/json")
//	res, _ := GetUserBalanceFromDB("timur")
//	w.Write([]byte(res.MoneyAmount))
//}
