package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id   int
	name string
)

func main() {
	db, err := sql.Open("mysql", "test:password@tcp(192.168.0.112:3306)/world")
	if err != nil {
		log.Println("Could not open SQL")
	}
	defer db.Close()
}
