package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/model/web"
	"github.com/sorfian/go-todo-list/repository"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewAuthMiddleware(userRepository repository.UserRepository, db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{UserRepository: userRepository, DB: db}
}

func (middleware *AuthMiddleware) Authenticate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get token from Authorization header
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(web.Response{
				Code:   401,
				Status: "Unauthorized",
				Data:   "Missing authorization header",
			})
		}

		// Check if token has Bearer prefix
		token := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(web.Response{
				Code:   401,
				Status: "Unauthorized",
				Data:   "Invalid token format",
			})
		}

		// Begin transaction
		tx := middleware.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}
		}()

		// Find a user by token using a repository
		user, err := middleware.UserRepository.FindByToken(ctx, tx, token)
		if err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.Status(fiber.StatusUnauthorized).JSON(web.Response{
					Code:   401,
					Status: "Unauthorized",
					Data:   "Invalid token",
				})
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(web.Response{
				Code:   500,
				Status: "Internal Server Error",
				Data:   "Failed to authenticate",
			})
		}

		// Check if the token is expired
		currentTime := time.Now().Unix()
		if user.TokenExp < currentTime {
			return ctx.Status(fiber.StatusUnauthorized).JSON(web.Response{
				Code:   401,
				Status: "Unauthorized",
				Data:   "Token expired",
			})
		}

		// Set user to context
		ctx.Locals("user", user)

		// Continue to the next handler
		return ctx.Next()
	}
}
