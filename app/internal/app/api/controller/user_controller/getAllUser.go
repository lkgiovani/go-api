package user_controller

import (
	"context"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"net/http"
)

func (uc *userController) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetAllUser(context.Background())
	if err != nil {
		jsonErrorResponse(w, http.StatusBadRequest, []map[string]string{{"error": "failed to get user from database"}})
		return newProjectError(projectError.EINTERNAL, "failed to get user from database", nil, "internal/app/api/controller/user_controller/getAllUser.go")

	}

	response := map[string][]user_model.User{"users": users}
	jsonSuccessResponse(w, http.StatusOK, response)
	return nil
}
