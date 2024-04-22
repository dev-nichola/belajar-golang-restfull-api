package service

import (
	"belajar_belajar_golang_restfull_api/exception"
	"belajar_belajar_golang_restfull_api/helper"
	"belajar_belajar_golang_restfull_api/model/domain"
	"belajar_belajar_golang_restfull_api/model/web"
	"belajar_belajar_golang_restfull_api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: categoryRepository,
		DB:         DB,
		Validate:   validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbakc(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.Repository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbakc(tx)

	category, err := service.Repository.FindById(ctx, tx, request.Id)

	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.Repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbakc(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)

	helper.PanicIfError(err)

	service.Repository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbakc(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollbakc(tx)

	categories := service.Repository.FindAll(ctx, tx)

	var categoryReponses []web.CategoryResponse

	for _, category := range categories {
		categoryReponses = append(categoryReponses, helper.ToCategoryResponse(category))
	}

	return categoryReponses
}
