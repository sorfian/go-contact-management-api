package test

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sorfian/go-contact-management-api/app"
	"github.com/sorfian/go-contact-management-api/controller"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/repository"
	"gorm.io/gorm"
)

// setupTestFiberApp creates and configures the Fiber app for testing
func setupTestFiberApp(
	userController controller.UserController,
	contactController controller.ContactController,
	addressController controller.AddressController,
	userRepository repository.UserRepository,
	db *gorm.DB,
) *fiber.App {
	testApp := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			status := "Internal Server Error"
			message := err.Error()

			// Handle custom errors
			switch e := err.(type) {
			case *fiber.Error:
				code = e.Code
				status = helper.GetStatusText(e.Code)
				message = e.Message
			case validator.ValidationErrors:
				code = fiber.StatusBadRequest
				status = "Bad Request"
				message = "Validation failed: " + e.Error()
			case helper.NotFoundError:
				code = fiber.StatusNotFound
				status = "Not Found"
				message = e.Err
			case helper.ResourceConflictError:
				code = fiber.StatusConflict
				status = "Resource Conflict"
				message = e.Err
			case helper.BadRequestError:
				code = fiber.StatusBadRequest
				status = "Bad Request"
				message = e.Err
			case helper.UnauthorizedError:
				code = fiber.StatusUnauthorized
				status = "Unauthorized"
				message = e.Err
			}

			return ctx.Status(code).JSON(web.Response{
				Code:   code,
				Status: status,
				Data:   message,
			})
		},
	})

	testApp.Use(recover.New(recover.Config{
		EnableStackTrace: false,
	}))

	app.Router(testApp, userController, contactController, addressController, userRepository, db)

	return testApp
}
