package router

import (
	"fmt"
	"go-api/internal/app/api/router/userRouter"
	"net/http"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	fmt.Println(123123)

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	mux.HandleFunc("/user/", userRouter.UserHandler)

	return mux
}
