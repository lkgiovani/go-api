package userRouter

import (
	"database/sql"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/pkg/projectError"
	"net/http"
)

func getAllUser(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return &projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Failed get all user",
			Path:    "internal/app/api/router/userRouter/getAllUser.go",
		}
	}

	controller := user_controller.NewUserController(db)

	err := controller.GetAllUser(w, r)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "Failed get all user",
			PrevError: err,
			Path:      "internal/app/api/router/userRouter/getAllUser.go",
		}
	}

	return nil
}
