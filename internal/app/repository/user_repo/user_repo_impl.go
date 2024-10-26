package user_repo

import (
	"context"
	"database/sql"
	"fmt"
	"go-api/internal/app/api/model/user_model"
	"go-api/pkg/projectError"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) GetUserById(ctx context.Context, id string) (user_model.User, error) {
	user := user_model.User{}

	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return user_model.User{}, &projectError.Error{Code: projectError.EINTERNAL, Message: err.Error()}
	}

	return user, nil
}

func (r *UserRepositoryImpl) SetUser(ctx context.Context, user user_model.User) error {
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Id, user.Name, user.Email)
	if err != nil {
		return &projectError.Error{Code: projectError.EINTERNAL, Message: err.Error()}
	}

	fmt.Println("salve1")

	return nil
}
