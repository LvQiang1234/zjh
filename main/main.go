package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"zjh/log"
	"zjh/network"
)

func main() {
	args := os.Args

	server := network.ServerSocket{}
	server.Init("127.0.0.1", 6000)
	if !server.Start() {
		log.Debug("服务器启动失败")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c

	fmt.Printf("server【%s】 exit ------- signal:[%v]", args[1], s)
}
