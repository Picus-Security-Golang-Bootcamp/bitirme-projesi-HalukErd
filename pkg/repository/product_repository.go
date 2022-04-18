package repository

import (
	"BasketProjectGolang/internal/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	repository := &productRepository{db: db}
	zap.L().Info("ProductRepository has been initialized.")
	//if err := repository.Migration(); err != nil {
	//	log.Fatalln("Product Migration Failed")
	//}

	return repository
}

func (r *productRepository) Create(b *entity.Product) (*entity.Product, error) {
	zap.L().Debug("product.repo.Create", zap.Reflect("product", b))
	if err := r.db.Create(b).Error; err != nil {
		zap.L().Error("product.repo.Create failed toΩ Create product", zap.Error(err))
		return nil, err
	}
	return b, nil
}

func (r *productRepository) GetAll(pageIndex, pageSize int) (*[]entity.Product, int) {
	zap.L().Debug("product.repo.GetAll")

	var bs = &[]entity.Product{}
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Preload("Author").Find(&bs).Count(&count)

	return bs, int(count)
}

func (r *productRepository) GetByID(id string) (*entity.Product, error) {
	zap.L().Debug("product.repo.GetByID", zap.Reflect("id", id))

	var product = &entity.Product{}
	if result := r.db.Preload("Author").First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *productRepository) Update(a *entity.Product) (*entity.Product, error) {
	zap.L().Debug("product.repo.Update", zap.Reflect("product", a))

	if result := r.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (r *productRepository) Delete(id string) error {
	zap.L().Debug("product.repo.Delete", zap.Reflect("id", id))

	product, err := r.GetByID(id)
	if err != nil {
		return err
	}

	if result := r.db.Delete(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *productRepository) Migration() error {
	return r.db.AutoMigrate(&entity.Product{})
}
