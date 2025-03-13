package repository

import "go-unit-test/5-mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}