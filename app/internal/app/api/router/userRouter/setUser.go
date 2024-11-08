package userRouter

import (
	"database/sql"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/pkg/projectError"
	"net/http"
)

func setUser(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Failed to set user",
			Path:    "internal/app/api/router/userRouter/getAllUser.go",
		}
	}

	controller := user_controller.NewUserController(db)
	// Inserindo o usuário
	err := controller.SetUser(w, r)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "Failed to set user",
			PrevError: err,
			Path:      "app/internal/app/api/router/userRouter/setUser.go",
		}
	}
	return nil
}
