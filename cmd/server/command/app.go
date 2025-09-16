package command

import (
	"context"
	"fmt"
	"github.com/devlibx/go-template-project/pkg/base"
	jsonplaceholderClient "github.com/devlibx/go-template-project/pkg/clients/jsonplaceholder"
	"github.com/devlibx/go-template-project/pkg/infra/database"
	consumers "github.com/devlibx/go-template-project/pkg/infra/messaging"
	"github.com/devlibx/go-template-project/pkg/service"
	"github.com/devlibx/gox-base/v2"
	"github.com/devlibx/gox-base/v2/errors"
	"github.com/devlibx/gox-base/v2/metrics"
	goxServer "github.com/devlibx/gox-base/v2/server"
	"github.com/devlibx/gox-base/v2/server/common"
	goxHttpApi "github.com/devlibx/gox-http/v4/api"
	statsCommon "github.com/devlibx/gox-metrics/v2/common"
	goxCadence "github.com/devlibx/gox-workfkow/workflow/framework/cadence"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log/slog"
	"time"
)

func AppMain(ctx context.Context, appConfig *ApplicationConfig, applicationContext *base.ApplicationContext) error {
	appConfig.SetDefaults()

	var sh goxServer.ServerShutdownHook
	var serverSignal *ServerSignal
	app := fx.New(
		// Supplied arguments
		fx.Supply(appConfig),
		fx.Supply(appConfig.App),
		fx.Supply(appConfig.HttpConfig),
		fx.Supply(appConfig.MetricConfig),
		fx.Supply(appConfig.MessagingConfig),
		fx.Supply(appConfig.RequestResponseSecurityConfig),
		fx.Supply(appConfig.CadenceConfig),
		fx.Supply(appConfig.OrdersRoMysqlConfig, appConfig.OrdersMysqlConfig),

		// Common generics dependencies
		fx.Provide(newCrossFunctionProvider),
		fx.Provide(statsCommon.NewMetricService),
		fx.Provide(common.NewServer),
		fx.Provide(goxServer.NewFxServerShutdownHook),
		fx.Provide(goxHttpApi.NewGoxHttpContext),
		fx.Provide(goxCadence.NewCadenceClient),
		fx.Provide(consumers.NewMessagingFactory),

		// Services
		service.Provider,
		database.Provider,

		// Clients
		jsonplaceholderClient.Provider,

		// Invoke - these will execute before app starts
		fx.Invoke(newApplicationEntryPoint),
		fx.Invoke(postApplicationSeverStart),
		fx.Invoke(consumers.NewMessagingFactoryLifecycle),
		fx.Invoke(goxCadence.NewCadenceWorkflowApiInvokerAtBoot),

		// This is a server signal which is sent when server is started
		fx.Provide(func() *ServerSignal { return &ServerSignal{StartedCh: make(chan error, 5)} }),
		fx.Populate(&sh, &serverSignal),

		fx.Populate(
			&applicationContext.GoxHttpContext,
			&applicationContext.OrdersDataStore,
		),
	)

	err := app.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to start server - app.Start() failed")
	}

	serverSignalError := <-serverSignal.StartedCh
	if serverSignalError != nil {
		return errors.Wrap(serverSignalError, "failed to start server")
	}

	sh.Setup(app)
	return nil
}

func newCrossFunctionProvider(appConfig *ApplicationConfig, metric metrics.Scope) (gox.CrossFunction, metrics.Publisher, error) {
	env := appConfig.App.Environment
	var loggerConfig zap.Config
	if env == "prod" {
		loggerConfig = zap.NewProductionConfig()
	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ := loggerConfig.Build()

	// Build metric publisher
	publisher, err := consumers.NewMetricPublisher(appConfig.MessagingConfig, logger)
	if err != nil {
		return nil, nil, err
	}

	return gox.NewCrossFunction(logger, metric, publisher), publisher, nil
}

type None string

// newApplicationEntryPoint is the main entry point function - which will start the server
func newApplicationEntryPoint(lc fx.Lifecycle, serverImpl ServerImpl, serverSignal *ServerSignal) None {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				// Setup routes
				serverImpl.routes()

				// Start server
				ch := make(chan error, 2)
				go func() {
					err := serverImpl.Start()
					ch <- err
				}()

				// Wait for server for 2 sec - in case of error we will stop
				select {
				case <-time.After(2 * time.Second):
					serverSignal.StartedCh <- nil
					close(serverSignal.StartedCh)
					slog.Info("Server started...")
					return nil
				case e := <-ch:
					if e != nil {
						serverSignal.StartedCh <- e
						close(serverSignal.StartedCh)
						slog.Error("Server failed to start...", e)
						return e
					} else {
						serverSignal.StartedCh <- nil
						close(serverSignal.StartedCh)
						slog.Info("Server started...")
						return nil
					}
				}
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Calling stop")
				<-serverImpl.Stop()
				return nil
			},
		},
	)
	return ""
}

func postApplicationSeverStart(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Calling server stop")
			return nil
		},
	})
}
