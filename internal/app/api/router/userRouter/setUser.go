package userRouter

import (
	"database/sql"
	"fmt"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/pkg/projectError"
	"net/http"
)

func setUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	controller := user_controller.NewUserController(db)
	// Inserindo o usuário
	err := controller.SetUser(w, r)
	if err != nil {
		http.Error(w, "Failed to set user", http.StatusInternalServerError)
		fmt.Println(&projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "Failed to set user",
			PrevError: err,
		})
		return
	}
	return
}
