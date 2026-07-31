package service

import (
	"context"
	"database/sql"
	"restapi/model/web"
	"restapi/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (r *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (r *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	panic("not implemented") // TODO: Implement
}
