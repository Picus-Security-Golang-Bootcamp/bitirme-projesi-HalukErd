package repository

import (
	"BasketProjectGolang/internal/entity"
)

type ProductRepository interface {
	Create(product *entity.Product) (*entity.Product, error)
	GetAll(pageIndex, pageSize int) (*[]entity.Product, int)
	GetByID(id string) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id string) error
	Migration() error
}
