package command

import (
	"context"
	"github.com/devlibx/gox-base/errors"
	"github.com/devlibx/gox-base/serialization"
	"go-template-project/config"
	"log/slog"
)

func FullMain(ctx context.Context, started chan bool) {

	// Read merged configs
	fullConfig, err := config.GetEnvExpandedMergedYamlApplicationConfig()
	if err != nil {
		panic(errors.Wrap(err, "something is wrong, failed to generate merged application config"))
	}

	// Build application config
	appConfig := ApplicationConfig{}
	err = serialization.ReadParameterizedYaml(fullConfig, &appConfig, "env")
	if err != nil {
		panic(errors.Wrap(err, "something is wrong, failed to build application config"))
	}

	slog.Info("Http Port", slog.Int("port", appConfig.App.HttpPort))

	// Start server
	if err = AppMain(ctx, &appConfig); err != nil {
		panic(err)
	}
	started <- true
	<-ctx.Done()
}
