package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/controller"
	"github.com/sorfian/go-todo-list/middleware"
	"github.com/sorfian/go-todo-list/repository"
	"gorm.io/gorm"
)

func Router(app *fiber.App, userController controller.UserController, userRepository repository.UserRepository, db *gorm.DB) {
	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(userRepository, db)

	// API v1 group
	api := app.Group("/api")

	// User routes
	users := api.Group("/users")

	users.Post("/register", userController.Register)
	users.Post("/login", userController.Login)

	users.Get("/current", authMiddleware.Authenticate(), userController.Get)
	users.Patch("/current", authMiddleware.Authenticate(), userController.Update)
	users.Delete("/current", authMiddleware.Authenticate(), userController.Logout)
}
