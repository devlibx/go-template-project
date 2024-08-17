package post

import (
	"context"
	"github.com/devlibx/gox-base/v2"
)

type Service interface {
	GetPost(ctx context.Context, postId string) (gox.StringObjectMap, error)
}
