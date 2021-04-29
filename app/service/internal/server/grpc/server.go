package grpc

import (
	"net"

	pb "gf/app/service/grpc"
	"gf/app/service/internal/conf"
	"google.golang.org/grpc"
)

func Init(c *conf.Config, serv *pb.GrpcServ)  {
	lis, err := net.Listen("tcp", c.Rpc.Addr)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDemoGrpcServer(grpcServer, serv)
	grpcServer.Serve(lis)
}