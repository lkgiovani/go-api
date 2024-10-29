package user_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"net/http"
)

func (uc *userController) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetAllUser(context.Background())
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	fmt.Println("salve")

	response := map[string][]user_model.User{"users": users}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	return nil
}
