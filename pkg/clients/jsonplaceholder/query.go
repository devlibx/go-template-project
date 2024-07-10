package jsonplaceholderClient

import (
	"context"
	goxHttpApi "github.com/devlibx/gox-http/v2/api"
	"github.com/devlibx/gox-http/v2/command"
)

func (c *client) GetPosts(ctx context.Context, postId string) (*PostDto, error) {
	type responseObj struct {
		Id     int    `json:"id"`
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}

	httpRequest := command.NewGoxRequestBuilder("getPosts").
		WithContentTypeJson().
		WithPathParam("postId", postId).
		Build()
	httpResponse, err := goxHttpApi.ExecuteHttp[responseObj, any](ctx, c, httpRequest)
	if err == nil {
		return &PostDto{
			Id:     httpResponse.Response.Id,
			UserId: httpResponse.Response.UserId,
			Title:  httpResponse.Response.Title,
		}, nil
	}
	return nil, err
}
