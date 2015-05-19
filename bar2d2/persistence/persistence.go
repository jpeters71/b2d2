package persistence

import (
	"fmt"
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

var bInitialized bool = false

func InitPers() {
	if !bInitialized {
		fmt.Printf("Registering driver.\n")

		var DB_DRIVER string
		sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{})		
		//sql.Register("sqlite3", &sqlite3.SQLiteDriver{})
		fmt.Printf("DB driver %s.\n", DB_DRIVER)
		bInitialized = true
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
