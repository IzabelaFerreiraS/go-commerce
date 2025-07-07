package schemas

import (
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ProductID  uint
	UserID   uint
	Quantity   int
	TotalPrice float64
}

