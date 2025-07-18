package server

import (
	"context"
	"github.com/mc-botnet/mc-botnet-server/internal/bot"
	"log/slog"
	"net/http"
)

type Server struct {
	manager *bot.Manager

	httpServer *http.Server
}

func NewServer(manager *bot.Manager) (*Server, error) {
	s := new(Server)

	mux := registerRoutes(s)

	s.manager = manager
	s.httpServer = &http.Server{
		Handler: mux,
	}

	return s, nil
}

func registerRoutes(s *Server) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Pong!"))
	})

	mux.HandleFunc("POST /bot", s.createBot)

	return mux
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("shutting down server")

	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Run(addr string) error {
	slog.Info("starting server", "addr", addr)
	s.httpServer.Addr = addr
	return s.httpServer.ListenAndServe()
}
