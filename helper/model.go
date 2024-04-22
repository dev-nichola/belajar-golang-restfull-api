package helper

import (
	"belajar_belajar_golang_restfull_api/model/domain"
	"belajar_belajar_golang_restfull_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
