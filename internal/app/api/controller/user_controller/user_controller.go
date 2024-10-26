package user_controller

import (
	"context"
	"database/sql"
	"go-api/internal/app/api/model/user_model"
	"go-api/internal/app/repository/user_repo"
	"go-api/pkg/projectError"
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

	return nil
}

func (uc *userController) GetAllUser() ([]user_model.User, error) {
	userDB := user_repo.NewUserRepository(uc.db)

	users, err := userDB.GetAllUser(context.Background())
	if err != nil {
		return users, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	return users, nil
}

func (uc *userController) GetUserById(id string) (user_model.User, error) {
	userDB := user_repo.NewUserRepository(uc.db)

	user, err := userDB.GetUserById(id)
	if err != nil {
		return user, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to set user in database",
			PrevError: err,
		}
	}

	return user, nil

}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &userController{
		db: db,
	}
}
