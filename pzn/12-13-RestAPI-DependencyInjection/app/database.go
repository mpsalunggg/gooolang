package app

import (
	"database/sql"
	"restapi/helper"
	"time"
)

func NewDB() *sql.DB {
	dsn := "golang:golang123@tcp(localhost:3307)/belajargolang?parseTime=true"
	db, err := sql.Open("mysql", dsn)

	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
