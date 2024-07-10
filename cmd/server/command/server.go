package command

import (
	"github.com/devlibx/gox-base"
	goxBaseConfig "github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-base/server/common"
	goxHttpApi "github.com/devlibx/gox-http/v2/api"
	stats "github.com/devlibx/gox-metrics/common"
	"github.com/gin-gonic/gin"
	"go-template-project/internal/handler"
	"go.uber.org/fx"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"net/http"
)

type ServerSignal struct {
	StartedCh chan error
}

type ServerImpl struct {
	fx.In
	common.Server
	gox.CrossFunction
	App                           *goxBaseConfig.App
	ServerSignal                  *ServerSignal
	RequestResponseSecurityConfig *goxHttpApi.RequestResponseSecurityConfig
	MetricHandler                 *stats.MetricHandler

	PostHandler handler.PostHandler
}

func (s *ServerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.GetRouter().ServeHTTP(w, r)
}

func (s *ServerImpl) routes() {

	router := s.GetRouter()

	// Setup prom path for stats
	router.GET("/metrics", func(context *gin.Context) {
		s.MetricHandler.ServeHTTP(context.Writer, context.Request)
	})
	router.GET("/health", s.healthCheck())

	// APIs which are exposed to other systems
	publicRouter := router.Group(s.App.AppName)
	publicRouter.Use(gintrace.Middleware(s.App.AppName))

	// V1 APIs - Protected
	v1Apis := publicRouter.Group("/api/v1")

	// Test Api
	v1UserApis := v1Apis.Group("/post")
	{
		v1UserApis.GET("/:postId", s.PostHandler.GetPost)
	}
}

func (s *ServerImpl) healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
	}
}
