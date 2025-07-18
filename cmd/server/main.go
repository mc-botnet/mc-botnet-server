package main

import (
	"context"
	"errors"
	"github.com/mc-botnet/mc-botnet-server/internal/bot"
	"github.com/mc-botnet/mc-botnet-server/internal/rpc"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mc-botnet/mc-botnet-server/internal/server"
)

func main() {
	// Create the bot runner
	runner, err := bot.NewKubernetesRunner()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Create the gRPC acceptor
	acceptor := rpc.NewAcceptor()

	// Create the bot manager
	manager := bot.NewManager(runner, acceptor)

	// Create the HTTP server
	s, err := server.NewServer(manager)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	go func() {
		err := acceptor.Run(":8081")
		if err != nil && !errors.Is(err, net.ErrClosed) {
			slog.Error(err.Error())
		}
		stop()
	}()

	go func() {
		err := s.Run(":8080")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
		}
		stop()
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := acceptor.Shutdown(ctx)
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.Shutdown(ctx)
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	wg.Wait()
}
