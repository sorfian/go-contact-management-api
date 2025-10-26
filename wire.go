//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/sorfian/go-contact-management-api/app"
	"github.com/sorfian/go-contact-management-api/controller"
	"github.com/sorfian/go-contact-management-api/repository"
	"github.com/sorfian/go-contact-management-api/service"
	"gorm.io/gorm"
)

// InitializeApp initializes the application with all dependencies
func InitializeApp() *fiber.App {
	wire.Build(
		// App dependencies (database, validator)
		app.Set,

		// Repositories
		repository.Set,

		// Services
		service.Set,

		// Controllers
		controller.Set,

		// Fiber app setup
		ProvideFiberApp,
	)
	return nil
}

// ProvideFiberApp creates and configures the Fiber app
func ProvideFiberApp(
	userController controller.UserController,
	contactController controller.ContactController,
	addressController controller.AddressController,
	userRepository repository.UserRepository,
	db *gorm.DB,
) *fiber.App {
	return setupFiberApp(userController, contactController, addressController, userRepository, db)
}
