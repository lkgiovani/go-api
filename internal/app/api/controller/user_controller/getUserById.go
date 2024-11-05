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
	uuid string `validate:"required,min=36,max=36"`
}

func (uc *userController) GetUserById(w http.ResponseWriter, r *http.Request) error {

	validate := validator.New()

	uuid := strings.TrimPrefix(r.URL.Path, "/user/")
	if uuid == "" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Missing id"}`, http.StatusBadRequest)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Missing id",
			Path:    "internal/app/api/controller/user_controller/getUserById.go",
		}
	}

	idInstance := id{
		uuid: uuid,
	}

	err := validate.Struct(idInstance)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Invalid id"}`, http.StatusBadRequest)
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "Error validating id",
			PrevError: err,
			Path:      "internal/app/api/controller/user_controller/getUserById.go",
		}
	}

	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetUserById(uuid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Invalid id"}`, http.StatusBadRequest)
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to get user",
			PrevError: err,
			Path:      "internal/app/api/controller/user_controller/getUserById.go",
		}
	}

	response := user_model.User{Id: uuid, Name: users.Name, Email: users.Email}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil

}
