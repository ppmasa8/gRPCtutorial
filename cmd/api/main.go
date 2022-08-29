package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ppmasa8/grpctutorial/pb"
	"github.com/ppmasa8/grpctutorial/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRockScissorsPaperServiceServer(server, service.NewRockScissorsPaperService())

	reflection.Register(server)
	server.Serve(listenPort)
}