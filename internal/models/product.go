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

type ProductToUpdate struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int32   `json:"quantity"`
}

func (product *ProductToUpdate) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Name") {
		tx.Statement.SetColumn("Name", product.Name)
	}

	if tx.Statement.Changed("Price") {
		tx.Statement.SetColumn("Price", product.Price)
	}

	if tx.Statement.Changed("Description") {
		tx.Statement.SetColumn("Description", product.Description)
	}

	if tx.Statement.Changed("Quantity") {
		tx.Statement.SetColumn("Quantity", product.Quantity)
	}

	return nil
}

func (product *Product) ConvertToResponse() ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
