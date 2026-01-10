package e2e

import (
	"context"
	"fmt"
	env "github.com/devlibx/go-template-project"
	"github.com/devlibx/go-template-project/cmd/server/command"
	"github.com/devlibx/go-template-project/pkg/base"
	httpHelper "github.com/devlibx/gox-base/v2/http_helper"
	httpCommand "github.com/devlibx/gox-http/v4/command/http"
	"github.com/stretchr/testify/suite"
	"github.com/zeebo/assert"
	"gopkg.in/resty.v1"
	"os"
	"testing"
	"time"
)

type e2eTestSuite struct {
	suite.Suite

	restyClient        *resty.Client
	ctx                context.Context
	done               context.CancelFunc
	applicationContext *base.ApplicationContext
}

func (s *e2eTestSuite) SetupSuite() {
	env.SetupE2ETestEnv(map[string]string{}, env.DefaultEnvSetupFunc())

	// Setup random ports for testing - you can simulate TEST service on this port
	go func() {
		mapping, err := httpHelper.AllocateFreePortsAndAssignToEnvironmentVariables("TEST_SERVICE", "")
		assert.NoError(s.T(), err)
		_ = os.Setenv("TEST_SERVICE", "localhost")
		_ = os.Setenv("TEST_SERVICE", fmt.Sprintf("%d", mapping["TEST_SERVICE"]))
	}()

	s.applicationContext = &base.ApplicationContext{}

	// Setup resty client
	httpCommand.EnableRestyDebug = true
	s.restyClient = resty.New()
	s.restyClient.HostURL = fmt.Sprintf("http://localhost:%s/%s/api/v1", os.Getenv("HTTP_PORT"), os.Getenv("APP_NAME"))
	s.restyClient.Debug = true
	s.restyClient.SetHeader("x-client-id", os.Getenv("CLIENT_ID"))
	s.restyClient.SetHeader("x-access-token", os.Getenv("CLIENT_TOKEN"))

	// Run full main application
	s.ctx, s.done = context.WithTimeout(context.Background(), 30*time.Second)
	ch := make(chan bool, 1)
	go func() {
		command.FullMain(s.ctx, ch, s.applicationContext)
	}()
	<-ch
}

func (s *e2eTestSuite) TearDownSuite() {
}

func TestE2ETestSuite(t *testing.T) {
	if os.Getenv("E2E_TESTS_ENABLED") != "true" {
		t.Skip("skipping test; E2E_TESTS_ENABLED is set to true")
		return
	}
	suite.Run(t, new(e2eTestSuite))
}
