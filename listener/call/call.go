package call

import (
	"context"
	"log"
)

type Call struct{}

func (c *Call) Talk(ctx context.Context, req *Request) (*Response, error) {
	log.Printf("received message: %s", req.GetIncoming())
	return &Response{Outgoing: "Can you repeat that"}, nil
}

func (c *Call) Hello(ctx context.Context, req *Request) (*Response, error) {
	log.Printf("received message: %s", req.GetIncoming())
	return &Response{Outgoing: "Hey, Hi"}, nil
}

func (c *Call) Bye(ctx context.Context, req *Request) (*Response, error) {
	log.Printf("received message: %s", req.GetIncoming())
	return &Response{Outgoing: "Take Care Mate!"}, nil
}
