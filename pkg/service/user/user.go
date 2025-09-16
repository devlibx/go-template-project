package user

import (
	"context"
	userModels "github.com/devlibx/go-template-project/pkg/database/user"
	"github.com/devlibx/gox-base/v2"
)

type userServiceImpl struct {
	gox.CrossFunction
	userDataStore userModels.UserDataStore
}

func (u *userServiceImpl) CreateUser(ctx context.Context, req userModels.CreateUserRequest) error {
	return u.userDataStore.CreateUser(ctx, req)
}

func (u *userServiceImpl) GetUserByID(ctx context.Context, userID string) (*userModels.User, error) {
	return u.userDataStore.GetUserByID(ctx, userID)
}

func (u *userServiceImpl) GetAllUsers(ctx context.Context) ([]*userModels.User, error) {
	return u.userDataStore.GetAllUsers(ctx)
}

func (u *userServiceImpl) UpdateUser(ctx context.Context, userID string, req userModels.UpdateUserRequest) error {
	return u.userDataStore.UpdateUser(ctx, userID, req)
}

func (u *userServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	return u.userDataStore.DeleteUser(ctx, userID)
}

func NewUserService(cf gox.CrossFunction, userDataStore userModels.UserDataStore) Service {
	return &userServiceImpl{
		CrossFunction: cf,
		userDataStore: userDataStore,
	}
}