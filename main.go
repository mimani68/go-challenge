package main

import (
	"flag"
	"fmt"
	"net"

	"app.ir/config"
	"app.ir/internal/data/db"
	"app.ir/internal/job"
	app "app.ir/internal/transport/grpc"
	"app.ir/pkg/logHandler"
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

	client, connect, disconnect := db.Open(fmt.Sprintf("%s/%s", cfg.Database.DbUri, cfg.Database.DbName))
	defer db.Close(client, connect, disconnect)

	go job.RunJobs(cfg.Server.CronJobTime)

	app.RunServer(listener)
}
