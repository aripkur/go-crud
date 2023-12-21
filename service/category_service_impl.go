package service

import (
	"context"
	"go-crud/exception"
	"go-crud/helper"
	"go-crud/model/domain"
	"go-crud/model/web"
	"go-crud/repository"
	"go-crud/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := validator.CategoryValidateUpdate(request)
	helper.PanicIfError(err)

	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}
	category = service.CategoryRepository.Update(ctx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	service.CategoryRepository.Delete(ctx, categoryId)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	categories := service.CategoryRepository.FindAll(ctx)

	return helper.ToCategoryResponses(categories)
}
