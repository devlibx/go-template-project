package post

import (
	"context"
	"github.com/devlibx/gox-base"
)

type Service interface {
	GetPost(ctx context.Context, postId string) (gox.StringObjectMap, error)
}
