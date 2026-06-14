package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := RunConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('mps', 'mps')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}

func TestQuerySql(t *testing.T) {
	db := RunConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := RunConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var created_at time.Time
		var birth_date sql.NullTime
		var married sql.NullBool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email)
		}
		if balance.Valid {
			fmt.Println("Balance:", balance)
		}
		if rating.Valid {
			fmt.Println("Rating:", rating)
		}
		if birth_date.Valid {
			fmt.Println("Birth Date:", birth_date)
		}
		if married.Valid {
			fmt.Println("Married:", married)
		}
		fmt.Println("Created At:", created_at)
	}
}

func TestSqlInjection(t *testing.T) {
	db := RunConnection()
	defer db.Close()

	ctx := context.Background()

	username := "putra'; #"
	password := "test"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string

		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login success")
	} else {
		fmt.Println("Login Failed")
	}
}
