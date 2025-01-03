package models

import (
	"time"

	"gorm.io/gorm"
)

// User domain model - represents the database entity
type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName"`
	Password  string    `binding:"required"`
	Username  string    `json:"username" binding:"required" gorm:"unique"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Roles     []Role    `json:"roles" gorm:"many2many:user_roles"`
}

// UserResponse - stripped down user data for API responses
type UserResponse struct {
	ID        uint           `json:"id"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Roles     []RoleResponse `json:"roles" `
}

type UserToUpdate struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Roles     []Role `json:"roles" `
}

// UserRole is the join table
type UserRole struct {
	UserID    uint `gorm:"primaryKey"`
	RoleID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (uc *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        uc.ID,
		FirstName: uc.FirstName,
		LastName:  uc.LastName,
		Email:     uc.Email,
		CreatedAt: uc.CreatedAt,
		UpdatedAt: uc.UpdatedAt,
		Roles:     convertRoles(uc.Roles),
	}
}

func convertRoles(roles []Role) []RoleResponse {
	rolesApi := make([]RoleResponse, len(roles))
	for i, role := range roles {
		rolesApi[i] = role.mapRolToApi()
	}
	return rolesApi
}

func (user *UserToUpdate) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("FirstName") {
		tx.Statement.SetColumn("FirstName", user.FirstName)
	}

	if tx.Statement.Changed("LastName") {
		println("last name")
		tx.Statement.SetColumn("LastName", user.LastName)
	}

	if tx.Statement.Changed("Email") {
		tx.Statement.SetColumn("Email", user.Email)
	}

	if tx.Statement.Changed("Roles") {
		tx.Statement.SetColumn("Roles", user.Roles)
	}

	return nil
}
