package main

import (
	"database/sql"
	"fmt"

	dbPack "example_mock/internal/db"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err.Error())
	}
	defer db.Close()

	dbService := dbPack.New(db)

	names, _ := dbService.GetNames()

	for _, name := range names {
		fmt.Println(name)
	}
}
