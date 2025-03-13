package service

import (
	"errors"
	"go-unit-test/5-mock/entity"
	"go-unit-test/5-mock/repository"
)

type CategoryServiceInterface interface {
	Get(id string) (*entity.Category, error)
}

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}
}
