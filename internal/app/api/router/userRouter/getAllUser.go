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

func (router *Router) getAllUser(db *sql.DB) {
	if router.Request.Method != "GET" {
		http.Error(router.w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	controller := user_controller.NewUserController(db)

	users, err := controller.GetAllUser()
	if err != nil {
		http.Error(router.w, "Failed to set user", http.StatusInternalServerError)
		fmt.Println(&projectError.Error{Code: projectError.EINTERNAL, Message: "Failed to set user", PrevError: err})

		return
	}
	fmt.Printf("Users: %+v\n", users)

	response := map[string][]user_model.User{"users": users}
	jsonResponse, _ := json.Marshal(response)
	router.w.Header().Set("Content-Type", "application/json")
	router.w.Write(jsonResponse)
}
