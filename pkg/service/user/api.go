package user

import (
	"context"
	userModels "github.com/devlibx/go-template-project/pkg/database/user"
)

type Service interface {
	CreateUser(ctx context.Context, req userModels.CreateUserRequest) error
	GetUserByID(ctx context.Context, userID string) (*userModels.User, error)
	GetAllUsers(ctx context.Context) ([]*userModels.User, error)
	UpdateUser(ctx context.Context, userID string, req userModels.UpdateUserRequest) error
	DeleteUser(ctx context.Context, userID string) error
}