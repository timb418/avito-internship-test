package dbl

import "log"

const (
	balances = `CREATE TABLE IF NOT EXISTS balances (
	user_id varchar(30) NOT NULL,
	balance text,
	PRIMARY KEY (user_id)
);`
	reservedMoney = `CREATE TABLE IF NOT EXISTS reserved_money (
	user_id varchar(30) UNIQUE,
	service_id text,
	order_id text NOT NULL,
	reserve_amount text,
	CONSTRAINT fk_user
      		FOREIGN KEY(user_id) 
	  	REFERENCES balances(user_id)
);`
	accountingLogbook = `CREATE TABLE IF NOT EXISTS accounting_logbook (
	user_id varchar(30),
	service_id text,
	order_id text NOT NULL,
	money_spent text,
	CONSTRAINT fk_user
      		FOREIGN KEY(user_id) 
	  	REFERENCES balances(user_id));`
)

func InitDB() {
	db, err := GetDBConnection()

	if err == nil {
		_, _ = db.Exec(balances)
		_, _ = db.Exec(reservedMoney)
		_, _ = db.Exec(accountingLogbook)
	} else {
		log.Fatalln("Could not initialize database")
	}
}
