package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProvideDatabase provides a database connection
func ProvideDatabase() *gorm.DB {
	return Connect()
}

// ProvideValidator provides a validator instance
func ProvideValidator() *validator.Validate {
	return validator.New()
}

// Set AppSet is a Wire provider set for app dependencies
var Set = wire.NewSet(
	ProvideDatabase,
	ProvideValidator,
)
