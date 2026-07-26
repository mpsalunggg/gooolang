package repository

import (
	"context"
	"database/sql"
	"restapi/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, tx sql.Tx, category domain.Category) {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, tx sql.Tx, categoryId int) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, tx sql.Tx) []domain.Category {
	panic("not implemented") // TODO: Implement
}
