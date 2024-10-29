package user_controller

import (
	"context"
	"encoding/json"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"io"
	"net/http"
	"strings"
)

func (uc *userController) UpdateUserById(w http.ResponseWriter, r *http.Request) error {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Missing id"}`, http.StatusBadRequest)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Missing id",
		}

	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)

	}

	var user PostUpDateUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)

	}

	userDB := user_repo.NewUserRepository(uc.db)

	err = userDB.UpdateUserById(context.Background(), user_repo.UpdateUserByIdDTO{
		Id:    id,
		Name:  &user.Name,
		Email: &user.Email,
	})

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
