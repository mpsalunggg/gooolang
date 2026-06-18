package repository

import (
	"context"
	"database/sql"
	"errors"
	"pzn-databases/entity"
	"strconv"
)

type CommentRepository interface {
	Create(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindByID(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{
		DB: db,
	}
}

func (repo *commentRepositoryImpl) Create(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comment(email, comment) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	comment.ID = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindByID(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)

	comment := entity.Comment{}

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment"

	rows, err := repo.DB.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}

		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
