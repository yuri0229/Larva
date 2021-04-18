package main

import (
	"gf/api"
	"gf/internal/conf"
	"gf/internal/service"
	"gf/internal/server/grpc"
	"gf/internal/server/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	serv := service.New(conf.Conf)
	httpServ := api.NewHttp(conf.Conf, serv)
	go func() {
		http.Init(conf.Conf, httpServ)
	}()
	grpcServ := api.NewGrpc(serv)
	go func() {
		grpc.Init(conf.Conf, grpcServ)
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