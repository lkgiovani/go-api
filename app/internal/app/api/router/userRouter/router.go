package userRouter

import (
	"database/sql"
	"go-api/pkg/projectError"
	"log"
	"net/http"
)

func handeError(next func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			log.Println(&projectError.Error{
				Code:      projectError.EINTERNAL,
				Message:   "Error",
				PrevError: err,
				Path:      "internal/app/api/router/userRouter/router.go",
			})

		}
		next(w, r)
	}
}

func handeWithDb(db *sql.DB, next func(w http.ResponseWriter, r *http.Request, db *sql.DB) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return next(w, r, db)
	}
}

func NewUserRouter(mux *http.ServeMux, db *sql.DB) error {

	mux.HandleFunc("GET /user", handeError(handeWithDb(db, getAllUser)))

	mux.HandleFunc("GET /user/{id}", handeError(handeWithDb(db, getUserById)))

	mux.HandleFunc("POST /user", handeError(handeWithDb(db, setUser)))

	mux.HandleFunc("DELETE /user/{id}", handeError(handeWithDb(db, deleteUserById)))

	mux.HandleFunc("PUT /user/{id}", handeError(handeWithDb(db, updateUserById)))

	return nil
}
