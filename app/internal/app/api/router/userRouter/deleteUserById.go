package userRouter

import (
	"database/sql"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/pkg/projectError"
	"net/http"
)

func deleteUserById(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Failed to delete user",
			Path:    "internal/app/api/router/userRouter/deleteUserById.go",
		}
	}

	controller := user_controller.NewUserController(db)

	err := controller.DeleteUserById(w, r)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "Failed to delete user",
			PrevError: err,
			Path:      "internal/app/api/router/userRouter/deleteUserById.go",
		}
	}
	return nil
}
