package user_controller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"io"
	"net/http"
)

type SetUser struct {
	Name  string `validate:"required,min=7,max=40"`
	Email string `validate:"required,email"`
}

func validateSetUser(email string, name string) error {
	validate := validator.New()
	idInstance := SetUser{Email: email, Name: name}
	return validate.Struct(idInstance)
}

func (uc *userController) SetUser(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"error": "Failed to read request body"}})
		return newProjectError(projectError.EINTERNAL, "Failed to read request body", nil, "internal/app/api/controller/user_controller/getUserById.go")
	}

	var user PostSetUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"error": "Invalid JSON payload"}})
		return newProjectError(projectError.EINTERNAL, "Invalid JSON payload", nil, "internal/app/api/controller/user_controller/getUserById.go")
	}

	// Validação do usuário
	if err := validateSetUser(user.Email, user.Name); err != nil {
		errorList := formatValidationErrors(err)
		jsonErrorResponse(w, http.StatusBadRequest, errorList)
		return newProjectError(projectError.EINTERNAL, "Invalid data", err, "internal/app/api/controller/user_controller/getUserById.go")
	}

	// Gerando o UUID para o usuário
	id, err := uuid.NewV7()
	if err != nil {
		jsonErrorResponse(w, http.StatusInternalServerError, []map[string]string{{"error": "Failed to generate ID"}})
		return newProjectError(projectError.EINTERNAL, "Failed to generate ID", nil, "internal/app/api/controller/user_controller/getUserById.go")
	}

	userDB := user_repo.NewUserRepository(uc.db)
	err = userDB.SetUser(context.Background(), user_model.User{Id: id.String(), Name: user.Name, Email: user.Email})
	if err != nil {
		jsonErrorResponse(w, http.StatusInternalServerError, []map[string]string{{"error": "Failed to set user in database"}})
		return newProjectError(projectError.EINTERNAL, "Failed to set user in database", nil, "internal/app/api/controller/user_controller/getUserById.go")
	}

	jsonSuccessResponse(w, http.StatusOK, "User created successfully!")
	return nil
}
