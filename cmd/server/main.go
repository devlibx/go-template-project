package main

import (
	"context"
	"github.com/devlibx/go-template-project/cmd/server/command"
	consumers "github.com/devlibx/go-template-project/pkg/infra/messaging"
	httpCommand "github.com/devlibx/gox-http/v4/command/http"
	"log/slog"
	"os"
)

func main() {
	consumers.DumpPinotMetricOnConsoleToDebug = true
	httpCommand.EnableRestyDebug = true

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)
	slog.SetDefault(logger)

	ctx := context.Background()
	command.FullMain(ctx, make(chan bool, 10))
	<-ctx.Done()
}
