package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Base response wrapper for consistent API responses
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// User domain model - represents the database entity
type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Never sent in responses
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateUserRequest - specific type for user creation endpoint
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

// UserResponse - stripped down user data for API responses
type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Example handler showing how these types work together
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create user in database...
	user := User{
		Email:     req.Email,
		Password:  hashPassword(req.Password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	// Convert to response type
	response := UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// Send response
	respond(w, APIResponse[UserResponse]{
		Success: true,
		Data:    response,
	})
}
