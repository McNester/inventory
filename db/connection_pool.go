package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"inventory/utils"
	"sync"
)

var (
	once       sync.Once
	connection *sqlx.DB
)

func init() {
	once.Do(connectToDB)
}

func connectToDB() {
	dbc := utils.GetEnvVariable("DBC")
	driverName := utils.GetEnvVariable("driver")

	db, err := sqlx.Connect(driverName, dbc)

	if err != nil {
		panic("Failed to connect to the db: " + err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic("Failed to ping the db: " + err.Error())
	}

	connection = db
}

func GetConnection() *sqlx.DB {
	return connection
}

func CloseConnection() {
	err := connection.Close()

	if err != nil {
		panic("Failed to close DB connection pool: " + err.Error())
	}
}
