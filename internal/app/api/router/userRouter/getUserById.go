package userRouter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-api/internal/app/api/controller/user_controller"
	"go-api/internal/app/api/model/user_model"
	"go-api/pkg/projectError"
	"net/http"
)

func (router *Router) getUserById(db *sql.DB) {
	if router.Request.Method != "GET" {
		http.Error(router.w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	id := router.Request.URL.Query().Get("id")

	if id == "" {
		http.Error(router.w, "Missing id", http.StatusBadRequest)
		fmt.Println(&projectError.Error{
			Code:    projectError.EINTERNAL,
			Message: "Missing id",
		})
		return
	}

	fmt.Println("ID:", id)
	controller := user_controller.NewUserController(db)

	user, err := controller.GetUserById(id)
	if err != nil {
		http.Error(router.w, "Failed to set user", http.StatusInternalServerError)
		fmt.Println(&projectError.Error{Code: projectError.EINTERNAL, Message: "Failed to set user", PrevError: err})
		return
	}

	fmt.Println("salve")

	response := user_model.User{Id: user.Id, Name: user.Name, Email: user.Email}
	jsonResponse, _ := json.Marshal(response)
	router.w.Header().Set("Content-Type", "application/json")
	router.w.Write(jsonResponse)
}
