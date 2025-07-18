package main

import (
	"context"
	"errors"
	"github.com/mc-botnet/mc-botnet-server/internal/bot"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mc-botnet/mc-botnet-server/internal/server"
)

func main() {
	mgr, err := bot.NewKubernetesManager()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer mgr.Shutdown()

	s, err := server.NewServer(mgr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		err = s.Run()
		if err != nil && !errors.Is(http.ErrServerClosed, err) {
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
