package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// SQLiteデータベースに接続
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// クエリを実行
	name := "Curl"
	age := 33
	row := db.QueryRow("SELECT id FROM users WHERE name = ? AND age = ?", name, age)

	// 結果を取得
	var id int
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("指定された条件に一致する行は存在しませんでした")
	case nil:
		fmt.Printf("id: %d\n", id)
	default:
		fmt.Println(err)
	}
}
