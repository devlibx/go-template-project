package consumers

import (
	"context"
	"github.com/devlibx/gox-base"
	goxMessaging "github.com/devlibx/gox-messaging"
	"github.com/devlibx/gox-messaging/factory"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MessagingFactory goxMessaging.Factory

type messagingServiceImpl struct {
	gox.CrossFunction
	logger *zap.Logger
	goxMessaging.Factory
}

func NewMessagingFactory(lifecycle fx.Lifecycle, cf gox.CrossFunction, configuration *goxMessaging.Configuration) (MessagingFactory, error) {
	service := messagingServiceImpl{
		CrossFunction: cf,
		logger:        cf.Logger(),
		Factory:       factory.NewMessagingFactory(cf),
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			// Start messaging
			err := service.Start(*configuration)

			// Metric publisher depends on this so we need to start it here when messaging is running
			if p, ok := cf.Publisher().(*Publisher); ok {
				err = p.Start(ctx, service)
			}

			return err
		},
		OnStop: func(ctx context.Context) error {
			return service.Stop()
		},
	})
	return &service, nil
}
