package userRouter

import (
	"database/sql"
	"fmt"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/pkg/projectError"
	"net/http"
)

func getAllUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	controller := user_controller.NewUserController(db)

	err := controller.GetAllUser(w, r)
	if err != nil {
		http.Error(w, "Failed to set user", http.StatusInternalServerError)
		fmt.Println(&projectError.Error{Code: projectError.EINTERNAL, Message: "Failed to set user", PrevError: err})

		return
	}
	return
}
