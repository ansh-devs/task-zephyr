package internal

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	pb "github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	roundrobin "github.com/hlts2/round-robin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	//viper.SetConfigType("env")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../.")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(err)
	}

	dbHost := viper.GetString("DB_HOST")
	dbUser := viper.GetString("DB_USER")
	dbPwd := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbString := fmt.Sprintf("postgres://%s:%s@%s/%s", dbUser, dbPwd, dbHost, dbName)
	conn, err := ConnectToDatabase(dbString, ctx)
	if err != nil {
		logrus.Errorf("error while connecting to database : %s", err)
	} else {
		err := conn.Ping(ctx)
		if err == nil {
			logrus.Info("ping to the database successful")
		} else {
			logrus.Error("ping to the database failed :", err)
		}
	}
	return &Orchestrator{
		Manager:                server,
		ServerPort:             port,
		Listener:               ln,
		WorkerPool:             make(map[string]*Worker),
		MaxHealthCheckOverlook: 3,
		DataStorePool:          conn,
		HealthCheckTTL:         time.Second * 2,
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
