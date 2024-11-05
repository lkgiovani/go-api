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
		jsonError(w, http.StatusBadRequest, "Missing id")
		return newProjectError(projectError.EINTERNAL, "Missing id", nil, "internal/app/api/controller/user_controller/deleteUserById.go")
	}

	userDB := user_repo.NewUserRepository(uc.db)
	err := userDB.DeleteUserById(context.Background(), id)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Failed to delete user")
		return newProjectError(projectError.EINTERNAL, "Failed to delete user", nil, "internal/app/api/controller/user_controller/deleteUserById.go")
	}

	response := map[string]string{"message": "User deleted successfully!"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}
