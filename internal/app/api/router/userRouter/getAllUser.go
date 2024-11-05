package userRouter

import (
	"database/sql"
	"go-api/internal/app/api/controller/user_controller"
	"net/http"
)

func getAllUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	controller := user_controller.NewUserController(db)

	err := controller.GetAllUser(w, r)
	if err != nil {
		return
	}
	return
}
