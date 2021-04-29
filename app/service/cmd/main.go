package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "gf/app/service/grpc"
	"gf/app/service/internal/conf"
	"gf/app/service/internal/server/grpc"
	"gf/app/service/internal/service"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	serv := service.New(conf.Conf)
	rpcServ := pb.New(serv)
	go func() {
		grpc.Init(conf.Conf, rpcServ)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Println("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}