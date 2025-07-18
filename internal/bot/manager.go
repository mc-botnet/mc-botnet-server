package bot

import (
	"context"
	"github.com/google/uuid"
	"github.com/mc-botnet/mc-botnet-server/internal/rpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"net"
	"strconv"
	"sync"
)

type StartOptions struct {
	Host     string
	Port     int
	Username string
	Auth     string
	Token    string
}

func toPod(opts *StartOptions, id uuid.UUID, image string) *corev1.Pod {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bot-" + id.String(),
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				Name:  "bot",
				Image: image,
				Env: []corev1.EnvVar{
					{
						Name:  "BOT_HOST",
						Value: opts.Host,
					},
					{
						Name:  "BOT_PORT",
						Value: strconv.Itoa(opts.Port),
					},
					{
						Name:  "BOT_USERNAME",
						Value: opts.Username,
					},
					{
						Name:  "BOT_AUTH",
						Value: opts.Auth,
					},
				},
			}},
		},
	}

	if opts.Token != "" {
		pod.Spec.Containers[0].Env = append(pod.Spec.Containers[0].Env, corev1.EnvVar{
			Name:  "BOT_TOKEN",
			Value: opts.Token,
		})
	}

	return pod
}

type Handle struct {
	ID     uuid.UUID
	conn   *grpc.ClientConn
	client rpc.BotClient
}

type GRPCAcceptor struct {
	rpc.UnimplementedAcceptorServer
	mu      sync.RWMutex
	clients map[uuid.UUID]rpc.BotClient
}

func (G *GRPCAcceptor) Ready(ctx context.Context, request *rpc.ReadyRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

type Manager interface {
	Run(ctx context.Context) error
	Shutdown() error

	StartBot(ctx context.Context, opts *StartOptions) (*Handle, error)
	StopBot(ctx context.Context, handle *Handle) error
}

type KubernetesManager struct {
	client   *kubernetes.Clientset
	listener net.Listener
}

func NewKubernetesManager() (*KubernetesManager, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubernetesManager{client, nil}, nil
}

func (k *KubernetesManager) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}
	k.listener = listener

	grpcServer := grpc.NewServer()

	acceptor := &GRPCAcceptor{
		clients: make(map[uuid.UUID]rpc.BotClient),
	}
	rpc.RegisterAcceptorServer(grpcServer, acceptor)

	return grpcServer.Serve(listener)
}

func (k *KubernetesManager) Shutdown() error {
	return k.listener.Close()
}

func (k *KubernetesManager) StartBot(ctx context.Context, opts *StartOptions) (*Handle, error) {
	id := uuid.New()

	// TODO move image name to config
	pod := toPod(opts, id, "mc-botnet-bot:latest")

	pod, err := k.pods().Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	pod.UID = pod.ObjectMeta.GetUID()

	return &Handle{ID: id}, nil
}

func (k *KubernetesManager) StopBot(ctx context.Context, handle *Handle) error {
	//TODO implement me
	panic("implement me")
}

func (k *KubernetesManager) pods() typedv1.PodInterface {
	return k.client.CoreV1().Pods("default")
}
