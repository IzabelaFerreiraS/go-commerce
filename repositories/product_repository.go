package repositories

import (
	"go-commerce/schemas"

	"gorm.io/gorm"
)

type ProductRepository interface {
    Create(product *schemas.Product) error
    Delete(product *schemas.Product) error
    List() ([]schemas.Product, error)
    FindByID(id string) (schemas.Product, error)
    Update(product *schemas.Product) error
}

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) Create(product *schemas.Product) error {
    return r.db.Create(product).Error
}

func (r *productRepository) Delete(product *schemas.Product) error {
    return r.db.Delete(product).Error
}

func (r *productRepository) List() ([]schemas.Product, error) {
    var products []schemas.Product
    err := r.db.Find(&products).Error
    return products, err
}

func (r *productRepository) FindByID(id string) (schemas.Product, error) {
    var product schemas.Product
    err := r.db.First(&product, id).Error
    return product, err
}

func (r *productRepository) Update(product *schemas.Product) error {
    return r.db.Save(product).Error
}
