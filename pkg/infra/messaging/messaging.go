package consumers

import (
	"context"
	"github.com/devlibx/gox-base/v2"
	"github.com/devlibx/gox-base/v2/errors"
	goxMessaging "github.com/devlibx/gox-messaging/v2"
	"github.com/devlibx/gox-messaging/v2/factory"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MessagingFactory goxMessaging.Factory

type messagingServiceImpl struct {
	gox.CrossFunction
	logger *zap.Logger
	goxMessaging.Factory
}

func NewMessagingFactoryLifecycle(lifecycle fx.Lifecycle, cf gox.CrossFunction, configuration *goxMessaging.Configuration, service MessagingFactory) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			// Start messaging
			err := service.Start(*configuration)
			if err != nil {
				return errors.Wrap(err, "failed to start messaging service")
			}

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
}

func NewMessagingFactory(cf gox.CrossFunction) (MessagingFactory, error) {
	service := messagingServiceImpl{
		CrossFunction: cf,
		logger:        cf.Logger(),
		Factory:       factory.NewMessagingFactory(cf),
	}
	return &service, nil
}
