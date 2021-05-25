package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"gf/app/api/internal/conf"
	"gf/app/api/internal/server/http"
	"gf/app/api/internal/service"
	"gf/pkg/log"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	serv := service.New(conf.Conf)
	go func() {
		http.Init(conf.Conf, serv)
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