package user_repo

import (
	"context"
	"go-api/internal/app/api/model/user_model"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) []user_model.User
	GetUserById(ctx context.Context, id string) user_model.User
	SetUser(ctx context.Context, user user_model.User) error
}
