package repository

import (
	"context"
	"go-crud/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, category domain.Category) domain.Category
	Update(ctx context.Context, category domain.Category) domain.Category
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context) []domain.Category
}
