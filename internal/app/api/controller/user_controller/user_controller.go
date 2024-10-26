package user_controller

import (
	"context"
	"database/sql"
	"fmt"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
)

type userController struct {
	db *sql.DB
}

func (uc *userController) SetUser(user user_model.User) error {

	fmt.Println(user)

	userDB := user_repo.NewUserRepository(uc.db)

	userDB.SetUser(context.Background(), user_model.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})

	return nil
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &userController{
		db: db,
	}
}
