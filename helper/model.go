package helper

import (
	"restful-api/model/domain"
	"restful-api/model/web"
)

func TocategoryResponse(category domain.Category) web.CategoryResponse {

	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

}

func TocategoryResponses(categories []domain.Category) []web.CategoryResponse {

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, web.CategoryResponse(category))
	}

	return categoryResponses

}
