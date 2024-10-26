package userRouter

import (
	"database/sql"
	"net/http"
)

func initializeRoutes(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	switch r.URL.Path {

	case "/user/getUser":

		getUserById(w, r, db)

	case "/user/getAllUser":

		getAllUser(w, r, db)

	case "/user/setUser":

		setUser(w, r, db)

	default:
		http.NotFound(w, r)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	initializeRoutes(w, r, db)
}
