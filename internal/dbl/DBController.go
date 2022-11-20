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
func UpdateUserBalanceInDB(userId string) (Balance, error) {
	db, err := GetDBConnection()
	if err == nil {
		sqlStatement := `
			INSERT INTO balances (user_id, balance)
			VALUES ($1, $2)`
		_, err = db.Exec(sqlStatement, b.UserId, b.MoneyAmount)
	}
	return err
}
