package handler

import (
	"github.com/devlibx/gox-base"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go-template-project/pkg/service/post"
	"go.uber.org/fx"
)

type PostHandler struct {
	fx.In
	gox.CrossFunction
	PostService post.Service
}

func (h *PostHandler) GetPost(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c, "postHandler.GetPost")
	defer span.Finish()

	if p, err := h.PostService.GetPost(ctx, c.Param("postId")); err == nil {
		c.JSON(200, p)
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}
