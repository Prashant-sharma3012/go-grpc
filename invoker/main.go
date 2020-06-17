package main

import (
	"context"
	"log"

	"github.com/basic-grpc/listener/call"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, _ = grpc.Dial(":9000", grpc.WithInsecure())

	defer conn.Close()

	c := call.NewCallingServiceClient(conn)

	req := call.Request{
		Incoming: "Hey You, hows life bro",
	}

	res, _ := c.Talk(context.Background(), &req)
	log.Printf("Response: %s", res.Outgoing)
}
