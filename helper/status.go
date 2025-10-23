package helper

import "github.com/gofiber/fiber/v2"

func GetStatusText(code int) string {
	switch code {
	case fiber.StatusOK:
		return "OK"
	case fiber.StatusCreated:
		return "Created"
	case fiber.StatusBadRequest:
		return "Bad Request"
	case fiber.StatusUnauthorized:
		return "Unauthorized"
	case fiber.StatusForbidden:
		return "Forbidden"
	case fiber.StatusNotFound:
		return "Not Found"
	case fiber.StatusInternalServerError:
		return "Internal Server Error"
	default:
		return "Error"
	}
}
