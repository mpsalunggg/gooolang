package service

import (
	"context"
	"database/sql"
	"restapi/helper"
	"restapi/model/domain"
	"restapi/model/web"
	"restapi/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = s.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = s.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	s.CategoryRepository.Delete(ctx, tx, category)
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := s.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
