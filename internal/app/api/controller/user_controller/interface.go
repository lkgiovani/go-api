package user_controller

import (
	"go-api/internal/app/api/model/user_model"
)

type UserControllerInterface interface {
	SetUser(user_model.User) error
	GetAllUser() ([]user_model.User, error)
	GetUserById(id string) (user_model.User, error)
}
