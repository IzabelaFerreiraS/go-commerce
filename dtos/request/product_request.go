package request

import (
	"fmt"
	"go-commerce/utils"
)

type CreateProductRequest struct {
	Name     string `json:"name"`
	Category  string `json:"category"`
	Description string `json:"description"`
	Price   int64  `json:"price"`
}

func (r *CreateProductRequest) Validate() error {
	if r.Name == "" && r.Category == "" && r.Description == "" && r.Price <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return utils.ErrParamIsRequired("name", "string")
	}
	if r.Category == "" {
		return utils.ErrParamIsRequired("category", "string")
	}
	if r.Description == "" {
		return utils.ErrParamIsRequired("description", "string")
	}
	if r.Price <= 0 {
		return utils.ErrParamIsRequired("price", "int64")
	}
	return nil
}

type UpdatedProductRequest struct {
	Name     string `json:"name"`
	Category  string `json:"category"`
	Description string `json:"description"`
	Price   int64  `json:"price"`
}

func (r *UpdatedProductRequest) Validate() error {
	if r.Name != "" || r.Category != "" || r.Description != "" || r.Price > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}