package database

import (
	"context"
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
		var id, name, email string
		var balance int32
		var rating float64
		var birth_date, created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id, "Name:", name, "Email:", email, "Balance:", balance, "Rating:", rating, "Birth Date:", birth_date, "Married:", married, "Created At:", created_at)
	}
}
