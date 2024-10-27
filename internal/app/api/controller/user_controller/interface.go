package user_controller

import (
	"go-api/internal/app/api/model/user_model"
	"net/http"
)

type UserControllerInterface interface {
	SetUser(w http.ResponseWriter, r *http.Request, user user_model.User) error
	GetAllUser(w http.ResponseWriter, r *http.Request) error
	GetUserById(w http.ResponseWriter, r *http.Request) error
	DeleteUserById(w http.ResponseWriter, r *http.Request) error
}
