package database

import (
	"context"
	"fmt"
	"testing"
)


func TestExecSql(t *testing.T){
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