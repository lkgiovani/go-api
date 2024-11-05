package userRouter

import (
	"database/sql"
	"net/http"
)

func NewUserRouter(mux *http.ServeMux, db *sql.DB) {

	mux.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {

		getAllUser(w, r, db)
	})

	mux.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {

		getUserById(w, r, db)
	})

	mux.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {

		setUser(w, r, db)
	})

	mux.HandleFunc("DELETE /user/{id}", func(w http.ResponseWriter, r *http.Request) {

		deleteUserById(w, r, db)
	})

	mux.HandleFunc("PUT /user/{id}", func(w http.ResponseWriter, r *http.Request) {

		updateUserById(w, r, db)
	})

}
