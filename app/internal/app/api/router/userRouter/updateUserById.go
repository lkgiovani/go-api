package userRouter

import (
	"database/sql"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/pkg/projectError"
	"net/http"
)

func updateUserById(w http.ResponseWriter, r *http.Request, db *sql.DB) error {

	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Failed to update user",
			Path:    "internal/app/api/router/userRouter/updateUserById.go",
		}
	}

	controller := user_controller.NewUserController(db)
	// Inserindo o usuário
	err := controller.UpdateUserById(w, r)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "Failed to update user",
			PrevError: err,
			Path:      "internal/app/api/router/userRouter/updateUserById.go",
		}
	}

	return nil
}
