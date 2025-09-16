package service

import (
	"github.com/devlibx/go-template-project/pkg/service/post"
	"github.com/devlibx/go-template-project/pkg/service/user"
	userModels "github.com/devlibx/go-template-project/pkg/database/user"
	"go.uber.org/fx"
)

var Provider = fx.Options(
	fx.Provide(post.NewPostService),
	fx.Provide(user.NewUserService),
	fx.Provide(userModels.NewUserDataStore),
	fx.Provide(userModels.NewOrderDataStore),
)
