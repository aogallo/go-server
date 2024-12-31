package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" binding:"required" gorm:"unique"`
	Description string  `json:"description" `
	Price       float64 `json:"price" binding:"required"`
	Quantity    int32   `json:"quantity" binding:"required"`
	gorm.Model
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int32     `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (product *Product) ConvertToResponse() (p ProductResponse) {
	return ProductResponse{
		ID: product.ID,

		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
