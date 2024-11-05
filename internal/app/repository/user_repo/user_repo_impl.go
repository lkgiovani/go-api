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

func (r *UserRepositoryImpl) GetAllUser(ctx context.Context) ([]user_model.User, error) {
	users := []user_model.User{}

	query := "SELECT id, name, email FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed query all user",
			PrevError: err,
			Path:      "repository/user_repo/user_repo_impl.go",
		}
	}
	defer rows.Close()

	for rows.Next() {
		var user user_model.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, &projectError.Error{
				Code:      projectError.EINTERNAL,
				Message:   "failed to get all user",
				PrevError: err,
				Path:      "repository/user_repo/user_repo_impl.go",
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to get all user",
			PrevError: err,
			Path:      "repository/user_repo/user_repo_impl.go",
		}
	}

	return users, nil
}

func (r *UserRepositoryImpl) SetUser(ctx context.Context, user user_model.User) error {
	query := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Id, user.Name, user.Email)
	if err != nil {
		return &projectError.Error{Code: projectError.EINTERNAL, Message: err.Error()}
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserById(id string) (user_model.User, error) {
	var user user_model.User

	query := `SELECT id, name, email FROM users WHERE id = ?`
	rows := r.db.QueryRowContext(context.Background(), query, id)

	if err := rows.Err(); err != nil {
		return user, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to get user",
			PrevError: err,
			Path:      "repository/user_repo/user_repo_impl.go",
		}
	}

	err := rows.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return user, &projectError.Error{
			Code:      projectError.EINTERNAL,
			Message:   "failed to get user",
			PrevError: err,
			Path:      "repository/user_repo/user_repo_impl.go",
		}
	}

	return user, nil

}

func (r *UserRepositoryImpl) DeleteUserById(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = ?"
	response, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return &projectError.Error{Code: projectError.EINTERNAL, Message: err.Error()}
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		return &projectError.Error{Code: projectError.EINTERNAL, Message: err.Error()}
	}

	if rowsAffected == 0 {
		return &projectError.Error{
			Code:      projectError.ENOTFOUND,
			Message:   fmt.Sprintf("user with id %s not found", id),
			PrevError: err,
		}
	}

	return nil
}

type UpdateUserByIdDTO struct {
	Id    string
	Name  *string
	Email *string
}

func (r *UserRepositoryImpl) UpdateUserById(ctx context.Context, user UpdateUserByIdDTO) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Id)
	if err != nil {
		return &projectError.Error{Code: projectError.EINTERNAL, Message: err.Error()}
	}

	return nil
}
