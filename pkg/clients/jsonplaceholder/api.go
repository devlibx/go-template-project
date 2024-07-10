package jsonplaceholderClient

import (
	"context"
	"github.com/devlibx/gox-base"
	goxHttpApi "github.com/devlibx/gox-http/v2/api"
)

type Client interface {
	GetPosts(ctx context.Context, postId string) (*PostDto, error)
}

type PostDto struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

type client struct {
	gox.CrossFunction
	goxHttpApi.GoxHttpContext
}

func NewClient(cf gox.CrossFunction, goxHttpCtx goxHttpApi.GoxHttpContext) Client {
	return &client{
		CrossFunction:  cf,
		GoxHttpContext: goxHttpCtx,
	}
}
