package user_controller

import (
	"github.com/go-playground/validator/v10"
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
	// Extrai o UUID da URL
	uuid := strings.TrimPrefix(r.URL.Path, "/user/")
	if uuid == "" {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"id": "Missing id"}})
		return newProjectError(projectError.EINTERNAL, "Missing id", nil, "internal/app/api/controller/user_controller/getUserById.go")
	}

	// Valida o UUID
	if err := validateUUID(uuid); err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"id": "Invalid id"}})
		return newProjectError(projectError.EINTERNAL, "Error validating id", err, "internal/app/api/controller/user_controller/getUserById.go")
	}

	// Busca o usuário no banco de dados
	userDB := user_repo.NewUserRepository(uc.db)
	user, err := userDB.GetUserById(uuid)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"error": "User not found"}})
		return newProjectError(projectError.EINTERNAL, "Failed to get user", err, "internal/app/api/controller/user_controller/getUserById.go")
	}

	// Responde com o usuário encontrado
	jsonSuccessResponse(w, http.StatusOK, user)
	return nil
}
