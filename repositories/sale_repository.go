package repositories

import (
	"go-commerce/schemas"

	"gorm.io/gorm"
)

type SaleRepository interface {
    Create(sale *schemas.Sale) error
    Delete(sale *schemas.Sale) error
    List() ([]schemas.Sale, error)
    FindByID(id string) (schemas.Sale, error)
}

type saleRepository struct {
    db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) SaleRepository {
    return &saleRepository{db: db}
}

func (r *saleRepository) Create(sale *schemas.Sale) error {
    return r.db.Create(sale).Error
}

func (r *saleRepository) Delete(sale *schemas.Sale) error {
    return r.db.Delete(sale).Error
}

func (r *saleRepository) List() ([]schemas.Sale, error) {
    var sales []schemas.Sale
    err := r.db.Find(&sales).Error
    return sales, err
}

func (r *saleRepository) FindByID(id string) (schemas.Sale, error) {
    var sale schemas.Sale
    err := r.db.First(&sale, id).Error
    return sale, err
}
