//go:build wireinject
// +build wireinject

package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/sorfian/go-contact-management-api/app"
	"github.com/sorfian/go-contact-management-api/controller"
	"github.com/sorfian/go-contact-management-api/repository"
	"github.com/sorfian/go-contact-management-api/service"
	"gorm.io/gorm"
)

// TestDependencies holds all test dependencies
type TestDependencies struct {
	App            *fiber.App
	DB             *gorm.DB
	UserRepository repository.UserRepository
}

// InitializeTestApp initializes the test application with all dependencies
func InitializeTestApp() *TestDependencies {
	wire.Build(
		// App dependencies
		app.Set,

		// Repositories
		repository.Set,

		// Services
		service.Set,

		// Controllers
		controller.Set,

		// Test app setup
		ProvideTestDependencies,
	)
	return nil
}

// ProvideTestDependencies creates and configures all test dependencies
func ProvideTestDependencies(
	userController controller.UserController,
	contactController controller.ContactController,
	addressController controller.AddressController,
	userRepository repository.UserRepository,
	db *gorm.DB,
) *TestDependencies {
	app := setupTestFiberApp(userController, contactController, addressController, userRepository, db)
	return &TestDependencies{
		App:            app,
		DB:             db,
		UserRepository: userRepository,
	}
}
