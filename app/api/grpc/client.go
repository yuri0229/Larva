package grpc

import (
	"gf/app/api/internal/conf"
	"google.golang.org/grpc"
)

func NewClient(c *conf.Config) DemoGrpcClient {
	conn, err := grpc.Dial(c.Rpc.Addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return NewDemoGrpcClient(conn)
}
