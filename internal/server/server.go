package server

import (
	"context"
	"log/slog"
	"net/http"
)

type Server struct {
	// client *kubernetes.Clientset

	httpServer *http.Server
}

func NewServer() (*Server, error) {
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	return nil, err
	// }

	// client, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	return nil, err
	// }

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})

	return &Server{
		// client: client,
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}, nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("shutting down server")

	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Run() error {
	slog.Info("starting server")

	return s.httpServer.ListenAndServe()
}
