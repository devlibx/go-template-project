package service

import (
	"github.com/devlibx/go-template-project/pkg/service/post"
	"go.uber.org/fx"
)

var Provider = fx.Options(
	fx.Provide(post.NewPostService),
)
