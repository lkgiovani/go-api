package userRouter

import (
	"net/http"
)

type Router struct {
	w http.ResponseWriter
	r *http.Request
}

func newRouter(w http.ResponseWriter, r *http.Request) *Router {
	return &Router{
		w: w,
		r: r,
	}
}

func (router *Router) initializeRoutes() {
	switch router.r.URL.Path {

	case "/user/getUser":
		router.getUser()

	case "/user/setUser":
		router.setUser()

	default:
		http.NotFound(router.w, router.r)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	router := newRouter(w, r)

	router.initializeRoutes()
}
