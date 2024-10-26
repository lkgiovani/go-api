package router

import (
	"database/sql"
	"fmt"
	"go-api/internal/app/api/router/userRouter"
	"net/http"
)

func InitializeRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		userRouter.UserHandler(w, r, db)
	})

	return mux
}
