package user_controller

import (
	"context"
	"encoding/json"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"net/http"
	"strings"
)

func (uc *userController) DeleteUserById(w http.ResponseWriter, r *http.Request) error {

	id := strings.TrimPrefix(r.URL.Path, "/user/")

	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Missing id"}`, http.StatusBadRequest)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Missing id",
		}

	}

	userDB := user_repo.NewUserRepository(uc.db)

	err := userDB.DeleteUserById(context.Background(), id)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	response := map[string]string{"message": "User deleted successfully!"}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}
