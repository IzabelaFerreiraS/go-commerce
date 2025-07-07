package services

import (
	"errors"
	"fmt"

	"go-commerce/dtos/request"
	"go-commerce/repositories"
	"go-commerce/schemas"
)

var ErrSaleNotFound = errors.New("sale not found")

type SaleService interface {
    CreateSale(req request.CreateSaleRequest) (schemas.Sale, error)
    DeleteSale(id string, currentUserRole string) error
    ListSales() ([]schemas.Sale, error)
    ShowSale(id string) (schemas.Sale, error)
}

type saleService struct {
    saleRepo    repositories.SaleRepository
    productRepo repositories.ProductRepository
}

func NewSaleService(saleRepo repositories.SaleRepository, productRepo repositories.ProductRepository) SaleService {
    return &saleService{saleRepo: saleRepo, productRepo: productRepo}
}

func (s *saleService) CreateSale(req request.CreateSaleRequest) (schemas.Sale, error) {
    product, err := s.productRepo.FindByID(fmt.Sprintf("%d", req.ProductID))
    if err != nil {
        return schemas.Sale{}, fmt.Errorf("product not found")
    }

    totalPrice := float64(product.Price) * float64(req.Quantity)

    sale := schemas.Sale{
        UserID:     req.UserID,
        ProductID:  req.ProductID,
        Quantity:   req.Quantity,
        TotalPrice: totalPrice,
    }

    err = s.saleRepo.Create(&sale)
    return sale, err
}

func (s *saleService) DeleteSale(id string, currentUserRole string) error {
    if currentUserRole != "admin" {
        return errors.New("only admin can delete sales")
    }

    sale, err := s.saleRepo.FindByID(id)
    if err != nil {
        return err
    }
    if sale.ID == 0 {
        return ErrSaleNotFound
    }

    return s.saleRepo.Delete(&sale)
}

func (s *saleService) ListSales() ([]schemas.Sale, error) {
    return s.saleRepo.List()
}

func (s *saleService) ShowSale(id string) (schemas.Sale, error) {
    sale, err := s.saleRepo.FindByID(id)
    if err != nil {
        return schemas.Sale{}, err
    }
    if sale.ID == 0 {
        return schemas.Sale{}, ErrSaleNotFound
    }
    return sale, nil
}
