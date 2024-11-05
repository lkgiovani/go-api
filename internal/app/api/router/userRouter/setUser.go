package userRouter

import (
	"database/sql"
	"go-api/internal/app/api/controller/user_controller"
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
		return
	}
	return
}
