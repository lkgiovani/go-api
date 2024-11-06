package user_controller

import (
	"database/sql"
)

type userController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &userController{
		db: db,
	}
}
