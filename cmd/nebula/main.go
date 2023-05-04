package main

import (
	"errors"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/synchthia/nebula-api/database"
	"github.com/synchthia/nebula-api/logger"
	"github.com/synchthia/nebula-api/server"
	"github.com/synchthia/nebula-api/service"
	"github.com/synchthia/nebula-api/stream"
)

func startGRPC(port string, svc *server.Services) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	return server.NewGRPCServer(svc).Serve(lis)
}

func main() {
	// Init Logger
	logger.Init()

	// Init
	logrus.Printf("[NEBULA] Starting Nebula Server...")

	svc := &server.Services{}

	// IP Filter
	isIPFilter := os.Getenv("ENABLE_IP_FILTER")
	if isIPFilter == "true" {
		dbipToken := os.Getenv("DB_IP_TOKEN")
		if len(dbipToken) == 0 {
			logrus.Errorf("[IPFilter] IPFilter enabled, but DB_IP_TOKEN does not provided!!")
			panic(errors.New("DB_IP_TOKEN is null"))
		}

		ipfw, err := service.NewIPFilter(&service.IPFilterConfig{
			DBIPToken: dbipToken,
		})
		if err != nil {
			panic(err)
		}
		svc.IPFilter = ipfw
	}

	// Redis
	go func() {
		redisAddr := os.Getenv("REDIS_ADDRESS")
		if len(redisAddr) == 0 {
			redisAddr = "localhost:6379"
		}
		stream.NewRedisPool(redisAddr)
	}()

	// Connect to MySQL
	mysqlConStr := os.Getenv("MYSQL_CONNECTION_STRING")
	if len(mysqlConStr) == 0 {
		mysqlConStr = "root:docker@tcp(localhost:3306)/nebula?charset=utf8mb4&parseTime=True&loc=Local"
	}
	mysqlClient := database.NewMysqlClient(mysqlConStr, "nebula")
    svc.MySQL = mysqlClient

	// gRPC
	wait := make(chan struct{})
	go func() {
		defer close(wait)
		port := os.Getenv("GRPC_LISTEN_PORT")
		if len(port) == 0 {
			port = ":17200"
		}

		msg := logrus.WithField("listen", port)
		msg.Infof("[GRPC] Listening %s", port)

		if err := startGRPC(port, svc); err != nil {
			logrus.Fatalf("[GRPC] gRPC Error: %s", err)
		}
	}()
	<-wait
}
