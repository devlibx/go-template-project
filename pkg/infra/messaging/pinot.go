package consumers

import (
	"context"
	"github.com/devlibx/gox-base/v2/metrics"
	messaging "github.com/devlibx/gox-messaging/v2"
	"go.uber.org/zap"
)

// DumpPinotMetricOnConsoleToDebug flag to be enabled to debug what is going into Pinot.
// In prod setup we will disable it
var DumpPinotMetricOnConsoleToDebug = false

type Publisher struct {
	MessagingFactory
}

func NewMetricPublisher(configuration *messaging.Configuration, logger *zap.Logger) (metrics.Publisher, error) {
	return metrics.NewNoOpPublisher(), nil
}

func (p *Publisher) Start(ctx context.Context, mf MessagingFactory) error {
	p.MessagingFactory = mf
	return nil
}

func (n *Publisher) Publish(ctx context.Context, p metrics.Publishable) error {
	return nil
}

func (n *Publisher) SilentPublish(ctx context.Context, p metrics.Publishable) {
}
