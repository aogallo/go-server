package models

import "time"

type Rol struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uint
}

type Tabler interface {
	TableName() string
}

func (Rol) TableName() string {
	return "roles"
}
