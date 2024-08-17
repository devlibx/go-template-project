package post

import (
	"context"
	"github.com/devlibx/gox-base/v2"
	jsonplaceholderClient "go-template-project/pkg/clients/jsonplaceholder"
)

type postServiceImpl struct {
	postClient jsonplaceholderClient.Client
}

func (p *postServiceImpl) GetPost(ctx context.Context, postId string) (gox.StringObjectMap, error) {
	if post, err := p.postClient.GetPosts(ctx, postId); err == nil {
		return gox.StringObjectMap{"id": post.Id}, nil
	} else {
		return nil, err
	}
}

func NewPostService(postClient jsonplaceholderClient.Client) Service {
	return &postServiceImpl{postClient: postClient}
}
