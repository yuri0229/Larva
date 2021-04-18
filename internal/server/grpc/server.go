package grpc

import (
	"gf/api"
	"gf/internal/conf"
	"google.golang.org/grpc"
	"net"
)

func Init(c *conf.Config, serv *api.GrpcServ)  {
	lis, err := net.Listen("tcp", c.Rpc.Addr)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDemoGrpcServer(grpcServer, serv)
	grpcServer.Serve(lis)
}