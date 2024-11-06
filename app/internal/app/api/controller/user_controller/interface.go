package user_controller

import (
	"net/http"
)

type UserControllerInterface interface {
	SetUser(w http.ResponseWriter, r *http.Request) error
	GetAllUser(w http.ResponseWriter, r *http.Request) error
	GetUserById(w http.ResponseWriter, r *http.Request) error
	DeleteUserById(w http.ResponseWriter, r *http.Request) error
	UpdateUserById(w http.ResponseWriter, r *http.Request) error
}
