package schemas

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string
	Category string
	Description string
	Price int64
}