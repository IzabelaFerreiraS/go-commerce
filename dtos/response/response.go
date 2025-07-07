package response

import "time"

type ProductResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
}

type SaleResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	ProductID  uint      `json:"productId"`
	UserID   uint      `json:"userId"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"totalPrice"`
}

type UserResponse struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"` 
	Role string `json:"role"`
}