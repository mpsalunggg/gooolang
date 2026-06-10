package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func RunConnection() *sql.DB {
	dsn := "root:rootpass123@tcp(localhost:3306)/belajargolang"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
