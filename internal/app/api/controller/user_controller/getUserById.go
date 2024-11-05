package user_controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"net/http"
	"strings"
)

type id struct {
	UUID string `validate:"required,uuid"`
}

func validateUUID(uuid string) error {
	validate := validator.New()
	idInstance := id{UUID: uuid}
	return validate.Struct(idInstance)
}

func (uc *userController) GetUserById(w http.ResponseWriter, r *http.Request) error {
	uuid := strings.TrimPrefix(r.URL.Path, "/user/")
	if uuid == "" {
		jsonError(w, http.StatusBadRequest, "Missing id")
		return newProjectError(projectError.EINTERNAL, "Missing id", nil, "internal/app/api/controller/user_controller/getUserById.go")
	}

	if err := validateUUID(uuid); err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid id")
		return newProjectError(projectError.EINTERNAL, "Error validating id", err, "internal/app/api/controller/user_controller/getUserById.go")
	}

	userDB := user_repo.NewUserRepository(uc.db)
	users, err := userDB.GetUserById(uuid)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "User not found")
		return newProjectError(projectError.EINTERNAL, "Failed to get user", err, "internal/app/api/controller/user_controller/getUserById.go")
	}

	response := user_model.User{Id: uuid, Name: users.Name, Email: users.Email}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Failed to encode response")
		return newProjectError(projectError.EINTERNAL, "Failed to encode response", err, "internal/app/api/controller/user_controller/getUserById.go")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}
