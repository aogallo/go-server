package models

import "gorm.io/gorm"

type Product struct {
	Name     string  `json:"name" form:"name" binding:"required" gorm:"unique"`
	Price    float64 `json:"price" form:"price" binding:"required"`
	Quantity int32   `json:"quantity" form:"quantity" bindinng:"required"`
	gorm.Model
}
