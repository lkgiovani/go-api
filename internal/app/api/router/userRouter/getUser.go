package userRouter

import (
	"fmt"
	"net/http"
)

func (router *Router) getUser() {
	if router.Request.Method != "GET" {
		http.Error(router.w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(router.w, "Rota /user/getUser acessada!")
}
