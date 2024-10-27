package router

import (
	"database/sql"
	"encoding/json"
	"go-api/internal/app/api/router/userRouter"
	"net/http"
)

type Router struct {
	W       http.ResponseWriter
	Request *http.Request
	db      *sql.DB
}

func InitializeRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		jsonResponse, _ := json.Marshal(map[string]string{"message": "pong"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		userRouter.UserHandler(w, r, db)
	})

	return mux
}
