package user_controller

import (
	"context"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"net/http"
	"strings"
)

func (uc *userController) DeleteUserById(w http.ResponseWriter, r *http.Request) error {

	id := strings.TrimPrefix(r.URL.Path, "/user/")
	if id == "" {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"id": "Missing id"}})
		return newProjectError(projectError.EINTERNAL, "Missing id", nil, "internal/app/api/controller/user_controller/deleteUserById.go")
	}

	userDB := user_repo.NewUserRepository(uc.db)
	err := userDB.DeleteUserById(context.Background(), id)
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"error": "Failed to delete user"}})
		return newProjectError(projectError.EINTERNAL, "Failed to delete user", nil, "internal/app/api/controller/user_controller/deleteUserById.go")
	}

	jsonSuccessResponse(w, http.StatusOK, "User deleted successfully!")
	return nil
}
