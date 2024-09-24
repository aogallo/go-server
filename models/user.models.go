package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string //`json:"firstName"`
	LastName  string //`json:"lastName"`
	Password  string `json:"-"`
	Username  string //`json:"username"`
}
