package main

import (
	"database/sql"
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func initDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATASOURCE"))
	handleErr(err)
	return db
}
