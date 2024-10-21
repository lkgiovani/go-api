package router

import (
	"fmt"
	"net/http"
)

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("funciona")
	})

	mux.HandleFunc("/salve", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("salve")
	})

	return mux
}
