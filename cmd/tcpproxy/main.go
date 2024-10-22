package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"waffle/internal/proxy"
)

func main() {
	ctx := context.Background()

	go func() {
		if err := tcpDummy(); err != nil {
			log.Panicf("dummy: %s", err.Error())
		}
	}()

	// just for testing
	time.Sleep(time.Second * 1)

	sender := proxy.NewTCPSender("127.0.0.1:8083", "127.0.0.1:8081")
	receiver := proxy.NewTCPReceiver("127.0.0.1:8080", sender)

	if err := sender.Start(ctx); err != nil {
		log.Panicf("start sender: %s", err.Error())
	}

	if err := receiver.Run(); err != nil {
		log.Panicf("run receiver: %s", err.Error())
	}
}

func tcpDummy() error {
	addr, err := net.ResolveTCPAddr("tcp", ":8081")
	if err != nil {
		return fmt.Errorf("resolve tcp addr: %w", err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen tcp: %w", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("listener accept: %w", err)
		}

		log.Println(conn)
	}
}
