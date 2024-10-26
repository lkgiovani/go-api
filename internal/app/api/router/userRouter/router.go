package userRouter

import (
	"database/sql"
	"net/http"
)

type Router struct {
	w       http.ResponseWriter
	Request *http.Request
	db      *sql.DB
}

func newRouter(w http.ResponseWriter, r *http.Request, db *sql.DB) *Router {
	return &Router{
		w:       w,
		Request: r,
		db:      db,
	}
}

func (router *Router) initializeRoutes() {

	switch router.Request.URL.Path {

	case "/user/getUser":
		router.getUser()

	case "/user/setUser":
		router.setUser(router.db)

	default:
		http.NotFound(router.w, router.Request)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	router := newRouter(w, r, db)

	router.initializeRoutes()
}
