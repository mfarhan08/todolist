package model

import "fmt"

type Item struct {
	Id   int
	Name string
}

const TodoItemTable = "CREATE TABLE IF NOT EXISTS todo(id INTEGER PRIMARY KEY AUTO_INCREMENT, name CHAR(200))"

func CreateTodo(item string) Item {
	query := fmt.Sprintf("INSERT INTO todo(name) values('%s')", item)
	statement, _ := database.Prepare(query)
	res, _ := statement.Exec()
	id, _ := res.LastInsertId()
	return GetTodo(int(id))
}

func GetTodo(id int) Item {
	query := fmt.Sprintf("SELECT * FROM todo WHERE id=%d", id)
	rows, _ := database.Query(query)
	defer rows.Close()
	result := Item{}
	rows.Next()
	rows.Scan(&(result.Id), &(result.Name))
	return result
}

func GetAllTodo() []Item {
	query := fmt.Sprintf("SELECT * FROM todo")
	rows, _ := database.Query(query)
	result := make([]Item, 0)
	defer rows.Close()
	for rows.Next() {
		new_row := Item{}
		rows.Scan(&(new_row.Id), &(new_row.Name))
		result = append(result, new_row)
	}
	return result
}

func DeleteTodo(id int) {
	query := fmt.Sprintf("DELETE FROM todo WHERE id=%d", id)
	delstate, _ := database.Prepare(query)
	delstate.Exec()
}
