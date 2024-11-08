package user_controller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"io"
	"net/http"
	"strings"
)

type UpdateUser struct {
	Name  string `validate:"omitempty,min=7,max=40"` // "omitempty" permite que o campo seja opcional
	Email string `validate:"omitempty,email"`
}

func validateUpdateUser(email string, name string) error {
	validate := validator.New()
	idInstance := UpdateUser{Email: email, Name: name}
	return validate.Struct(idInstance)
}

func (uc *userController) UpdateUserById(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	if id == "" {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"id": "Missing id"}})
		return newProjectError(projectError.EINTERNAL, "Missing id", nil, "internal/app/api/controller/user_controller/updateUserById.go")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"error": "Failed to read request body"}})
		return newProjectError(projectError.EINTERNAL, "Failed to read request body", nil, "internal/app/api/controller/user_controller/updateUserById.go")
	}

	var user PostUpDateUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"error": "Invalid JSON payload"}})
		return newProjectError(projectError.EINTERNAL, "Invalid JSON payload", nil, "internal/app/api/controller/user_controller/updateUserById.go")
	}
	// Validação do usuário
	if err := validateUpdateUser(user.Email, user.Name); err != nil {
		errorList := formatValidationErrors(err)
		jsonErrorResponse(w, http.StatusBadRequest, errorList)
		return newProjectError(projectError.EINTERNAL, "Invalid data", err, "internal/app/api/controller/user_controller/updateUserById.go")
	}

	userDB := user_repo.NewUserRepository(uc.db)
	err = userDB.UpdateUserById(context.Background(), user_repo.UpdateUserByIdDTO{
		Id:    id,
		Name:  &user.Name,
		Email: &user.Email,
	})

	if err != nil {
		jsonErrorResponse(w, http.StatusInternalServerError, []map[string]string{{"error": "Failed to update user in database"}})
		return newProjectError(projectError.EINTERNAL, "Failed to update user in database", err, "internal/app/api/controller/user_controller/updateUserById.go")

	}

	jsonSuccessResponse(w, http.StatusOK, "User updated successfully!")
	return nil
}
