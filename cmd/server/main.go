package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/mc-botnet/mc-botnet-server/internal/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer s.Shutdown(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		err = s.Run()
		if err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
		}
		stop()
	}()

	<-ctx.Done()
}
