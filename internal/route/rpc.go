package route

import (
	"go-starter/internal/handler/rpc"
	"go-starter/internal/pb"

	"google.golang.org/grpc"
)

func SetupRpc(srv grpc.ServiceRegistrar) {
	greeter := new(rpc.HelloService)
	pb.RegisterGreeterServer(srv, greeter)
}
