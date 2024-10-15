package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName"`
	Password  string    `binding:"required"`
	Username  string    `json:"username" binding:"required" gorm:"unique"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Roles     []Rol     `json:"roles" gorm:"many2many:user_roles;"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Roles     []Rol     `json:"roles" `
}

type UserUpdate struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Roles     []Rol  `json:"roles" `
}

func (uc *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        uc.ID,
		FirstName: uc.FirstName,
		LastName:  uc.LastName,
		Email:     uc.Email,
		Roles:     uc.Roles,
	}
}
