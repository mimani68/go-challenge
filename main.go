package main

import (
	"flag"
	"fmt"
	"net"

	"app.ir/config"
	"app.ir/internal/data/db"
	"app.ir/internal/data/repository"
	"app.ir/internal/handler"
	"app.ir/internal/job"
	"app.ir/internal/service"
	"app.ir/internal/transport/grpc"
	"app.ir/pkg/logHandler"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	port = flag.Int("port", 3000, "The server port")
)

func main() {

	flag.Parse()

	cfg := config.NewConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		logHandler.LogError(err.Error())
	}

	dsn := fmt.Sprintf("%s/%s?authSource=admin&connect=direct", cfg.Database.DbUri, cfg.Database.DbName)
	client, connect, disconnect := db.Open(dsn)
	if client.Ping(connect, &readpref.ReadPref{}) == nil {
		logHandler.Log("[Db] Ping")
	}
	defer db.Close(client, connect, disconnect)
	dbInstance := db.NewDatabase(*cfg, client.Database(cfg.Database.DbName))

	segmentRepo := repository.CreateSegmentRepository(*cfg, dbInstance)
	segmentService := service.NewSegmentService(*segmentRepo)
	segmentHandler := handler.NewSegmentHandler(segmentService)

	go job.RunJobs(cfg.Server.CronJobTime)

	grpc.RunServer(listener, segmentHandler)
}
