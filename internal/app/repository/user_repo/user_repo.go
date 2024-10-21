package user_repo

import (
	"context"
	"go-api/internal/app/api/model/user_model"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id string) user_model.User
	SetUser(ctx context.Context, user user_model.User) error
}
