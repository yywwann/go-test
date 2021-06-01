package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example/load_balance_test/lb_demo2/generator/config"
	"example/load_balance_test/lb_demo2/generator/server/grpc"
	"example/load_balance_test/lb_demo2/generator/service"
	"example/load_balance_test/lb_demo2/naming"
	"github.com/Terry-Mao/goim/pkg/ip"
	"github.com/inconshreveable/log15"
)

const srvName = "generator"

var log = log15.New("cmd", srvName)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	log.Info("config info:",
		"Node", config.Conf.Node,
		"Weight", config.Conf.Weight)
	// service init
	svc := service.New(config.Conf)
	rpc := grpc.New(config.Conf.GRPCServer, svc)

	// register server
	_, port, _ := net.SplitHostPort(config.Conf.GRPCServer.Addr)
	addr := fmt.Sprintf("%s:%s", ip.InternalIP(), port)
	if err := naming.Register(config.Conf.Reg.RegAddrs, config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema, 15, config.Conf.Weight); err != nil {
		panic(err)
	}
	fmt.Println("register ok")

	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			time.Sleep(time.Second * 2)
			rpc.Shutdown(ctx)
			naming.UnRegister(config.Conf.Reg.SrvName, addr, config.Conf.Reg.Schema)
			log.Info(srvName + " server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
