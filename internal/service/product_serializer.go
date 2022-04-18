package service

import (
	"BasketProjectGolang/internal/entity"
	"BasketProjectGolang/internal/model/api"
)

func requestToProduct(request *api.CreateProductRequest) entity.Product {
	return entity.Product{
		Name:  *request.Name,
		Code:  *request.Code,
		Price: *request.Price,
	}
}

func productToResponse(product *entity.Product) api.CreateProductResponse {
	return api.CreateProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Code:  product.Code,
		Price: product.Price,
	}
}
