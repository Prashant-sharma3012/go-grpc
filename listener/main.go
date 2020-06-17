package main

import (
	"log"
	"net"

	"github.com/basic-grpc/listener/call"
	grpc "google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":9000")

	c := call.Call{}
	grpcServer := grpc.NewServer()

	call.RegisterCallingServiceServer(grpcServer, &c)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc %v", err)
	}
}
