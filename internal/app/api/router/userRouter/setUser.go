package userRouter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PostUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (router *Router) setUser() {
	if router.r.Method != "POST" {
		http.Error(router.w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	body, err := io.ReadAll(router.r.Body)
	if err != nil {
		http.Error(router.w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	var user PostUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(router.w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Println(user)

	fmt.Fprintf(router.w, "Rota /user/setUser acessada!")
}
