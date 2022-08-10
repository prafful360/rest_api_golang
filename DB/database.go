package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Initialize() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:PR15@aug@97@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func RowExists(query string, args ...interface{}) bool {
	db := Initialize()
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}


