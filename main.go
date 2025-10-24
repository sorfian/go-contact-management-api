package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sorfian/go-todo-list/app"
	"github.com/sorfian/go-todo-list/controller"
	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/web"
	"github.com/sorfian/go-todo-list/repository"
	"github.com/sorfian/go-todo-list/service"
)

func main() {
	db := app.Connect()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	fiberApp := fiber.New(fiber.Config{
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

	// Middleware
	fiberApp.Use(recover.New())
	fiberApp.Use(logger.New())

	// Setup routes
	app.Router(fiberApp, userController, userRepository, db)

	// Start server
	log.Fatal(fiberApp.Listen(":3000"))
}
