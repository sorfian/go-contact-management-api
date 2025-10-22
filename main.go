package go_todo_list_app

import (
	"github.com/go-playground/validator/v10"
	"github.com/sorfian/go-todo-list/app"
	"github.com/sorfian/go-todo-list/repository"
	"github.com/sorfian/go-todo-list/service"
)

func main() {
	db := app.Connect()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
}
