package userRouter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/internal/app/api/model/user_model"
	"io"
	"net/http"
)

type PostUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (router *Router) setUser(db *sql.DB) {
	if router.Request.Method != "POST" {
		http.Error(router.w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	body, err := io.ReadAll(router.Request.Body)
	if err != nil {
		http.Error(router.w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	var user PostUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(router.w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	controller := user_controller.NewUserController(db)

	id, err := uuid.NewV7()
	if err != nil {
		fmt.Println("Erro ao gerar UUID v7:", err)
		return
	}

	err = controller.SetUser(user_model.User{
		Id:    id.String(),
		Name:  user.Name,
		Email: user.Email,
	})

	if err != nil {
		fmt.Println("Erro ao setar usuário:", err)
		return
	}

	fmt.Fprintf(router.w, "Rota /user/setUser acessada!")
}
