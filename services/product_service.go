package services

import (
	"errors"

	"go-commerce/dtos/request"
	"go-commerce/repositories"
	"go-commerce/schemas"
)

var ErrProductNotFound = errors.New("product not found")

type ProductService interface {
    CreateProduct(request.CreateProductRequest) (schemas.Product, error)
    DeleteProduct(id string) error
    ListProducts() ([]schemas.Product, error)
    ShowProduct(id string) (schemas.Product, error)
    UpdateProduct(id string, req request.UpdatedProductRequest) (schemas.Product, error)
}

type productService struct {
    repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
    return &productService{repo: repo}
}

func (s *productService) CreateProduct(req request.CreateProductRequest) (schemas.Product, error) {
    product := schemas.Product{
        Name:        req.Name,
        Category:    req.Category,
        Description: req.Description,
        Price:       req.Price,
    }
    err := s.repo.Create(&product)
    return product, err
}

func (s *productService) DeleteProduct(id string) error {
    product, err := s.repo.FindByID(id)
    if err != nil {
        return err
    }
    if product.ID == 0 {
        return ErrProductNotFound
    }
    return s.repo.Delete(&product)
}

func (s *productService) ListProducts() ([]schemas.Product, error) {
    return s.repo.List()
}

func (s *productService) ShowProduct(id string) (schemas.Product, error) {
    product, err := s.repo.FindByID(id)
    if err != nil {
        return schemas.Product{}, err
    }
    if product.ID == 0 {
        return schemas.Product{}, ErrProductNotFound
    }
    return product, nil
}

func (s *productService) UpdateProduct(id string, req request.UpdatedProductRequest) (schemas.Product, error) {
    product, err := s.repo.FindByID(id)
    if err != nil {
        return schemas.Product{}, err
    }
    if product.ID == 0 {
        return schemas.Product{}, ErrProductNotFound
    }

    if req.Name != "" {
        product.Name = req.Name
    }
    if req.Category != "" {
        product.Category = req.Category
    }
    if req.Description != "" {
        product.Description = req.Description
    }
    if req.Price > 0 {
        product.Price = req.Price
    }

    err = s.repo.Update(&product)
    return product, err
}
