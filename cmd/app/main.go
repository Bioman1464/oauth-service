package main

import (
	"context"
	"os/signal"
	"syscall"

	"auth-service/internal/app"
)

func main() {
	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, cancel := signal.NotifyContext(appCtx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	go app.Start(ctx)

	<-ctx.Done()

	app.Stop(ctx)
}
