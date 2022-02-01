package main

import (
	"flag"
	"fmt"
	"net"

	"app.ir/internal/handler"
	"app.ir/pkg/logHandler"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	flag.Parse()

	server, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logHandler.LogError(err.Error())
	}

	handler.RunHandler(server)
	logHandler.Log(fmt.Sprintf("server listening at %v", server.Addr()))

}
