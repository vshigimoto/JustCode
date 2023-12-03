package main

import (
	"log"
	"net"

	"github.com/you/pirate/internal/pkg/transport/server/grpc"

	gogrpc "google.golang.org/grpc"

	protoport "github.com/you/pirate/pkg/protobuf/port"
)

func main() {
	log.Println("server")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := gogrpc.NewServer()
	protoServer := grpc.Port{}

	protoport.RegisterPortServiceServer(server, protoServer)

	log.Fatalln(server.Serve(listener))
}
