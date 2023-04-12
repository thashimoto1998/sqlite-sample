package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// テーブルの作成
	createTableStmt := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            age INTEGER
        );
    `
	_, err = db.Exec(createTableStmt)
	if err != nil {
		panic(err)
	}

	// データの挿入
	insertStmt := `
        INSERT INTO users (name, age) VALUES (?, ?);
    `
	_, err = db.Exec(insertStmt, "Alice", 28)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(insertStmt, "Curl", 33)
	if err != nil {
		panic(err)
	}

	// データのクエリ
	//rows, err := db.Query("SELECT * FROM users")
	rows, err := db.Query("SELECT id, name, age FROM users WHERE name = ? AND age = ?", "Curl", 33)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//for rows.Next() {
	rows.Next()
	var id int
	var name string
	var age int
	err = rows.Scan(&id, &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	//}
}
