package handler

import (
	"github.com/devlibx/gox-base/v2"
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

type name struct {
	name string `json:"name" validate:"required"`
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
