package bot

import (
	"github.com/google/uuid"
	"github.com/mc-botnet/mc-botnet-server/internal/rpc"
	"github.com/mc-botnet/mc-botnet-server/internal/rpc/pb"
	"google.golang.org/grpc"
)

type StartOptions struct {
	McHost     string
	McPort     int
	McUsername string
	McAuth     string
	McToken    string

	GRPCHost string
	GRPCPort int
}

type Bot struct {
	ID     uuid.UUID
	conn   *grpc.ClientConn
	client pb.BotClient
}

type Manager struct {
	runner   Runner
	acceptor *rpc.Acceptor
}

func NewManager(runner Runner, acceptor *rpc.Acceptor) *Manager {
	return &Manager{runner, acceptor}
}
