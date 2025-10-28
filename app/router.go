package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sorfian/go-contact-management-api/controller"
	"github.com/sorfian/go-contact-management-api/middleware"
	"github.com/sorfian/go-contact-management-api/repository"
	"gorm.io/gorm"
)

func Router(app *fiber.App, userController controller.UserController, contactController controller.ContactController, addressController controller.AddressController, userRepository repository.UserRepository, db *gorm.DB) {
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
	users.Delete("/logout", authMiddleware.Authenticate(), userController.Logout)

	// Contact routes
	contacts := api.Group("/contacts", authMiddleware.Authenticate())
	contacts.Post("/", contactController.Create)
	contacts.Get("/", contactController.GetAll)
	contacts.Get("/:contactId", contactController.Get)
	contacts.Patch("/:contactId", contactController.Update)
	contacts.Delete("/:contactId", contactController.Delete)

	// Address routes (nested under contacts)
	addresses := contacts.Group("/:contactId/addresses")
	addresses.Post("/", addressController.Create)
	addresses.Get("/", addressController.GetAll)
	addresses.Get("/:addressId", addressController.Get)
	addresses.Patch("/:addressId", addressController.Update)
	addresses.Delete("/:addressId", addressController.Delete)

	// Serve OpenAPI spec file
	app.Get("/apispec.yaml", func(c *fiber.Ctx) error {
		file, err := os.ReadFile("./docs/apispec.yaml")
		if err != nil {
			return c.Status(404).SendString("OpenAPI spec not found")
		}
		c.Set("Content-Type", "application/x-yaml")
		return c.Send(file)
	})

	// Swagger UI
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:                      "/apispec.yaml",
		DeepLinking:              false,
		DocExpansion:             "list",
		DefaultModelsExpandDepth: 1,
		InstanceName:             "swagger",
		Title:                    "Contact Management API Documentation",
	}))
}
