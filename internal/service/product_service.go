package service

import (
	"BasketProjectGolang/internal/entity"
	"BasketProjectGolang/internal/model/api"
	"BasketProjectGolang/internal/repository"
	"context"
	"go.uber.org/zap"
)

type ProductService interface {
	Create(ctx context.Context, req *api.CreateProductRequest) (*api.CreateProductResponse, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService {
	p := &productService{repo: repo}
	zap.L().Info("ProductService has been initialized.")
	return p
}

func (s *productService) Create(ctx context.Context, request *api.CreateProductRequest) (*api.CreateProductResponse, error) {
	var product entity.Product
	product = requestToProduct(request)
	createdProduct, err := s.repo.Create(&product)
	if err != nil {
		return nil, err
	}
	response := productToResponse(createdProduct)
	return &response, nil
}
