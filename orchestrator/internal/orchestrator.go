package internal

import (
	"context"
	pb "github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	roundrobin "github.com/hlts2/round-robin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
	"time"
)

const (
	some = "NAW !"
)

type Orchestrator struct {
	pb.UnimplementedOrchestratorServiceServer
	Manager                *grpc.Server
	ServerPort             string
	Listener               net.Listener
	Wg                     sync.WaitGroup
	WorkerPool             map[string]*Worker
	WorkerPoolMtx          sync.RWMutex
	AcquirableWorkerIDs    []string
	rr                     roundrobin.RoundRobin
	AcquirableWorkerIDsMtx sync.Mutex
	MaxHealthCheckOverlook uint8
	DataStorePool          *pgxpool.Pool
	HealthCheckTTL         time.Duration
	Ctx                    context.Context
	CtxCancel              context.CancelFunc
}

type Worker struct {
	HealthCheckMiss uint8
	Address         string
	Manager         *grpc.ClientConn
	Recipient       pb.BackgroundWorkerServiceClient
}

func NewOrchestrator(server *grpc.Server, ln net.Listener, port string, ctx context.Context) *Orchestrator {
	newCtx, cancelCtxFunc := context.WithCancel(ctx)
	return &Orchestrator{
		Manager:                server,
		ServerPort:             port,
		Listener:               ln,
		WorkerPool:             make(map[string]*Worker),
		MaxHealthCheckOverlook: 3,
		DataStorePool:          &pgxpool.Pool{},
		HealthCheckTTL:         2,
		Ctx:                    newCtx,
		CtxCancel:              cancelCtxFunc,
	}
}

func (s *Orchestrator) Serve() {
	err := s.Manager.Serve(s.Listener)
	if err != nil {
		logrus.Error(err)
	}
}

func (s *Orchestrator) PerformReflection() {
	logrus.WithFields(logrus.Fields{"message": "using reflection!"}).Info("Task-Zephyr")
	reflection.Register(s.Manager)
}
