package main

import (
	"context"
	httpCommand "github.com/devlibx/gox-http/v2/command/http"
	"go-template-project/cmd/server/command"
	consumers "go-template-project/pkg/infra/messaging"
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
