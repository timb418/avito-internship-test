package dbl

import (
	"database/sql"
	"fmt"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "habrpguser"
	password = "pgpwd4habr"
	dbname   = "habrdb"
)

func GetDBConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	return db, err
}

func AddMoneyToUser(b Balance) error {
	db, err := GetDBConnection()
	if err == nil {
		sqlStatement := `
			INSERT INTO balances (user_id, balance)
			VALUES ($1, $2)`
		_, err = db.Exec(sqlStatement, b.UserId, b.MoneyAmount)
		return err
	}
	return err
}
func GetUserBalanceFromDB(userId string) (Balance, error) {
	db, err := GetDBConnection()
	var b Balance
	if err == nil {
		sqlStatement := `select user_id, balance from balances where user_id = $1`
		_ = db.QueryRow(sqlStatement, userId).Scan(&b.UserId, &b.MoneyAmount)

	}
	return b, err
}
func UpdateUserBalanceInDB(userId, newBalance string) error {
	db, err := GetDBConnection()
	if err == nil {
		sqlStatement := `UPDATE balances set balance=$1 where user_id=$2`
		_, err = db.Exec(sqlStatement, newBalance, userId)
		return err
	}
	return err
}
func FreezeMoney(r Reserve) error {
	db, err := GetDBConnection()
	if err == nil {
		//TODO
		sqlStatement := `INSERT INTO reserved_money (user_id, service_id, order_id, reserve_amount)
							VALUES($1, $2, $3, $4) 
							ON CONFLICT (user_id) 
							DO 
						   	UPDATE SET reserve_amount = ` + r.MoneyAmount + `;`
		_, err = db.Exec(sqlStatement, r.UserId, r.ServiceId, r.OrderId, r.MoneyAmount)
		return err
	}
	return err
}
func GetFrozenMoneyForUser(userId string) (Reserve, error) {
	db, err := GetDBConnection()
	var r Reserve
	if err == nil {
		sqlStatement := `select user_id, service_id, order_id, reserve_amount from reserved_money where user_id = $1`
		_ = db.QueryRow(sqlStatement, userId).Scan(&r.UserId, &r.ServiceId, &r.OrderId, &r.MoneyAmount)

	}
	return r, err
}
