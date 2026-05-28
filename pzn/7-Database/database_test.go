package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	dsn := "root:rootpass123@tcp(localhost:3306)/belajargolang"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
