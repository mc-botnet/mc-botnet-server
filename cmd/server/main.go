package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mc-botnet/mc-botnet-server/internal/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		err = s.Run()
		if err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
		}
		stop()
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.Shutdown(ctx)
	if err != nil {
		slog.Error(err.Error())
	}
}
