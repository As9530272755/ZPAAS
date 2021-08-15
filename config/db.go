package config

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func LinkDb(driverName, dsn string) error {
	var err error
	DB, err = sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if DB.Ping(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CloseDb() {
	DB.Close()
}
