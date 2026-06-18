package repository

import (
	"context"
	"fmt"
	database "pzn-databases"
	"pzn-databases/entity"
	"testing"
)

func TestCreate(t *testing.T) {
	comment := entity.Comment{
		Email:   "mps123@gmail.com",
		Comment: "Hellow",
	}

	commentRepo := NewCommentRepository(database.RunConnection())

	ctx := context.Background()

	result, err := commentRepo.Create(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByID(t *testing.T) {
	var ID int32 = 17

	commentRepo := NewCommentRepository(database.RunConnection())

	ctx := context.Background()

	result, err := commentRepo.FindByID(ctx, ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("ID : ", result.ID)
	fmt.Println("Email : ", result.Email)
	fmt.Println("Comment : ", result.Comment)
}

func TestFindAll(t *testing.T) {
	commentRepo := NewCommentRepository(database.RunConnection())

	ctx := context.Background()

	result, err := commentRepo.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println("ID : ", comment.ID)
		fmt.Println("Email : ", comment.Email)
		fmt.Println("Comment : ", comment.Comment)
		fmt.Println("============================")
	}
}
