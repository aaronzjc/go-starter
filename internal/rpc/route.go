package rpc

import (
	"go-starter/internal/pb"
	"go-starter/internal/rpc/handler"

	"google.golang.org/grpc"
)

func SetupRoute(srv grpc.ServiceRegistrar) {
	greeter := new(handler.HelloService)
	pb.RegisterGreeterServer(srv, greeter)
}
