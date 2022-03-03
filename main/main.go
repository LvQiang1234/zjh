package main

import (
	"os"
	"os/signal"
	"syscall"
	"zjh/log"
	"zjh/logic/account"
	"zjh/network"
)

func init() {
	account.Init()
}

func main() {
	server := network.ServerSocket{}
	server.Init("127.0.0.1", 6000)
	if !server.Start() {
		log.Debug("服务器启动失败")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c

	log.Debug("server exit because get signal: %v", s)
}
