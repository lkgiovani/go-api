package user_controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"net/http"
)

type userController struct {
	db *sql.DB
}

func (uc *userController) SetUser(w http.ResponseWriter, r *http.Request, user user_model.User) error {
	userDB := user_repo.NewUserRepository(uc.db)

	err := userDB.SetUser(context.Background(), user)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	response := map[string]string{"message": "User created successfully!"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	return nil
}

func (uc *userController) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetAllUser(context.Background())
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	fmt.Println("salve")

	response := map[string][]user_model.User{"users": users}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	return nil
}

func (uc *userController) GetUserById(w http.ResponseWriter, r *http.Request, id string) error {

	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetUserById(id)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	fmt.Printf("Users: %+v\n", users)

	response := user_model.User{Id: id, Name: users.Name, Email: users.Email}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	return nil

}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &userController{
		db: db,
	}
}
