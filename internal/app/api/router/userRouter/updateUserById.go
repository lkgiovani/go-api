package userRouter

import (
	"database/sql"
	"net/http"
)

func updateUserById(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
