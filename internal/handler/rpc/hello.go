package rpc

import (
	"context"
	"go-starter/internal/pb"
)

type HelloService struct {
	*pb.UnimplementedGreeterServer
}

var _ pb.GreeterServer = HelloService{}

func (HelloService) SayHello(ctx context.Context, req *pb.GreeterRequest) (*pb.GreeterResponse, error) {
	return &pb.GreeterResponse{Hello: "hello world"}, nil
}
