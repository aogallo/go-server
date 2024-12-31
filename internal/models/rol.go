package models

import (
	"time"
)

type Role struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" form:"name" binding:"required" gorm:"unique"`
	Users     []User    `json:"users" gorm:"many2many:user_roles"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Tabler interface {
	TableName() string
}

func (Role) TableName() string {
	return "roles"
}

type RolAPI struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (rol *Role) mapRolToApi() RolAPI {
	return RolAPI{
		ID:   rol.ID,
		Name: rol.Name,
	}
}
