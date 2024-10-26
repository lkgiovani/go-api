package user_controller

import (
	"context"
	"database/sql"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
	"log"
)

type userController struct {
	db *sql.DB
}

func (uc *userController) SetUser(user user_model.User) error {
	userDB := user_repo.NewUserRepository(uc.db)

	err := userDB.SetUser(context.Background(), user)
	if err != nil {
		return &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	log.Println("User created successfully") // Usando log em vez de Println
	return nil
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &userController{
		db: db,
	}
}
