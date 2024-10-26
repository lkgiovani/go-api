package userRouter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/internal/app/api/model/user_model"
	"go-api/pkg/projectError"
	"io"
	"net/http"
)

type PostUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (router *Router) setUser(db *sql.DB) {
	if router.Request.Method != "POST" {
		http.Error(router.w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(router.Request.Body)
	if err != nil {
		http.Error(router.w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var user PostUserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(router.w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	controller := user_controller.NewUserController(db)

	// Gerando o UUID para o usuário
	id, err := uuid.NewV7()
	if err != nil {
		http.Error(router.w, "Failed to generate UUID", http.StatusInternalServerError)
		return
	}

	// Inserindo o usuário
	err = controller.SetUser(user_model.User{
		Id:    id.String(),
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		http.Error(router.w, "Failed to set user", http.StatusInternalServerError)
		fmt.Println(&projectError.Error{Code: projectError.EINTERNAL, Message: err.Error()})

		return
	}

	response := map[string]string{"message": "User created successfully!"}
	jsonResponse, _ := json.Marshal(response)
	router.w.Header().Set("Content-Type", "application/json")
	router.w.Write(jsonResponse)
}
