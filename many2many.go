package main

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint   `gorm:"primarykey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Name      string
	Roles     []Role `gorm:"many2many:user_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Role model
type Role struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
	Users       []User `gorm:"many2many:user_roles;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// UserRole is the join table
type UserRole struct {
	UserID    uint `gorm:"primaryKey"`
	RoleID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	// Initialize DB connection
	db, err := gorm.Open( /* your database configuration */ )
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate the schemas
	db.AutoMigrate(&User{}, &Role{}, &UserRole{})

	// Create a role
	adminRole := Role{
		Name:        "admin",
		Description: "Administrator role",
	}

	// Create the role first
	result := db.Create(&adminRole)
	if result.Error != nil {
		panic("failed to create role: " + result.Error.Error())
	}

	// Create a user
	user := User{
		Email: "admin@example.com",
		Name:  "Admin User",
	}

	// Create the user
	result = db.Create(&user)
	if result.Error != nil {
		panic("failed to create user: " + result.Error.Error())
	}

	// Assign role to user - Method 1: Using Association
	db.Model(&user).Association("Roles").Append(&adminRole)

	// Method 2: Creating multiple roles and assigning them
	roles := []Role{
		{Name: "editor", Description: "Editor role"},
		{Name: "viewer", Description: "Viewer role"},
	}

	// Create roles in batch
	db.Create(&roles)

	// Assign multiple roles to user
	db.Model(&user).Association("Roles").Append(roles)

	// Query user with roles
	var userWithRoles User
	db.Preload("Roles").First(&userWithRoles, user.ID)

	// Remove a role from user
	db.Model(&user).Association("Roles").Delete(&adminRole)

	// Clear all roles from user
	db.Model(&user).Association("Roles").Clear()

	// Find users with specific role
	var usersWithAdminRole []User
	db.Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Where("roles.name = ?", "admin").
		Find(&usersWithAdminRole)
}

// Helper function to check if user has role
func HasRole(db *gorm.DB, userID uint, roleName string) bool {
	var count int64
	db.Table("user_roles").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ? AND roles.name = ?", userID, roleName).
		Count(&count)
	return count > 0
}
