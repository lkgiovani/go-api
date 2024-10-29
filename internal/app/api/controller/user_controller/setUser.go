package user_controller

import (
	"context"
	"encoding/json"
	"github.com/gofrs/uuid"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"io"
	"net/http"
)

func (uc *userController) SetUser(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)

	}

	var user PostSetUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)

	}

	// Gerando o UUID para o usuário
	id, err := uuid.NewV7()
	if err != nil {
		http.Error(w, "Failed to generate UUID", http.StatusInternalServerError)

	}

	userDB := user_repo.NewUserRepository(uc.db)

	err = userDB.SetUser(context.Background(), user_model.User{Id: id.String(), Name: user.Name, Email: user.Email})
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	response := map[string]string{"message": "User created successfully!"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}
