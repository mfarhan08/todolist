package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func Connect() error {

	database, _ = sql.Open("mysql", "root:neosteam@tcp(127.0.0.1:3306)/todolist")
	statement, err := database.Prepare(TodoItemTable)
	fmt.Println(err)
	statement.Exec()
	return err
}
