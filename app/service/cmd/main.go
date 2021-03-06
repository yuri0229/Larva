package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "gf/app/service/grpc"
	"gf/app/service/internal/conf"
	"gf/app/service/internal/server/grpc"
	"gf/app/service/internal/service"
	"gf/pkg/log"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	serv := service.New(conf.Conf)
	rpcServ := pb.New(serv)
	go func() {
		grpc.Init(conf.Conf, rpcServ)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}