package resolvers

import "github.com/jinzhu/gorm"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver -> root resolver
type Resolver struct {
	db *gorm.DB
}
