package main

import (
	"context"
	"log"

	"github.com/you/pirate/internal/pkg/transport/client/grpc"

	gogrpc "google.golang.org/grpc"
)

func main() {
	log.Println("client")

	conn, err := gogrpc.Dial(":50051", gogrpc.WithInsecure(), gogrpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	portClient := grpc.NewPort(conn)
	if err := portClient.Create(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
