package request

import (
	"fmt"
	"go-commerce/utils"
)

type CreateSaleRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (r *CreateSaleRequest) Validate() error {
	if r.UserID == 0 && r.ProductID == 0 && r.Quantity == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.UserID == 0 {
		return utils.ErrParamIsRequired("user_id", "uint")
	}
	if r.ProductID == 0 {
		return utils.ErrParamIsRequired("product_id", "uint")
	}
	if r.Quantity <= 0 {
		return utils.ErrParamIsRequired("quantity", "positive integer")
	}
	return nil
}

type UpdatedSaleRequest struct {
	Quantity int `json:"quantity"`
}

func (r *UpdatedSaleRequest) Validate() error {
	if r.Quantity <= 0 {
		return fmt.Errorf("at least one valid field must be provided (quantity must be > 0)")
	}
	return nil
}

type DeletedSaleRequest struct {
	Role string `json:"role"`
}

func (r *DeletedSaleRequest) Validate() error {
	if r.Role == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	return nil
}
