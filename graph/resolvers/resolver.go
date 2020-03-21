package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/phamstack/godek/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver -> root resolver
type Resolver struct {
	// db *gorm.DB
	Services *models.Services
}
