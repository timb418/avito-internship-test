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
	// Parse the request
	var req Reserve
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	oldBalance, err := GetUserBalanceFromDB(req.UserId)
	oldAmount, _ := currency.NewAmount(oldBalance.MoneyAmount, "RUB")
	moneyToSub, err := currency.NewAmount(req.MoneyAmount, "RUB")

	am, _ := oldAmount.Sub(moneyToSub)

	err = UpdateUserBalanceInDB(req.UserId, am.Number())
	if err != nil {
		return
	}

	// get frozen money, add to it, update
	newReserve, err := GetFrozenMoneyForUser(req.UserId)
	freezeAmount, _ := currency.NewAmount(newReserve.MoneyAmount, "RUB")
	am, _ = freezeAmount.Add(moneyToSub)

	//push it to db
	err = FreezeMoney(Reserve{UserId: req.UserId, ServiceId: req.ServiceId,
		OrderId: req.OrderId, MoneyAmount: am.Number(),
	})

}
func Acknowledge(w http.ResponseWriter, r *http.Request) {
	// Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму
	var req Reserve
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// списать деньги со старого замороженного счета
	oldBalance, err := GetUserFrozenBalanceFromDB(req.UserId)
	oldAmount, _ := currency.NewAmount(oldBalance.MoneyAmount, "RUB")
	moneyToSub, err := currency.NewAmount(req.MoneyAmount, "RUB")
	am, _ := oldAmount.Sub(moneyToSub)

	err = UpdateFrozenUserBalanceInDB(req.UserId, am.Number())
	if err != nil {
		return
	}

	// добавить в бух
	err = AddAccountingRecord(req)
	if err != nil {
		return
	}
}
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
