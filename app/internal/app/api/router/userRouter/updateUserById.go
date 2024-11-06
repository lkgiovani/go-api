package userRouter

import (
	"database/sql"
	"go-api/internal/app/api/controller/user_controller"
	"net/http"
)

func updateUserById(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	controller := user_controller.NewUserController(db)
	// Inserindo o usuário
	err := controller.UpdateUserById(w, r)
	if err != nil {
		return
	}
	return
}
