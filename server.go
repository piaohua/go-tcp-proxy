// +build ignore

package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8087")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn.(*net.TCPConn))
	}
}

func handle(server *net.TCPConn) {
	client, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Printf("dial err: %v\n", err)
		return
	}

	go func() {
		buf := make([]byte, 2048)
		_, err := io.CopyBuffer(server, client, buf)
		if err != nil {
			log.Printf("server err %v\n", err)
		}
		server.Close()
	}()

	buf := make([]byte, 2048)
	_, err = io.CopyBuffer(client, server, buf)
	if err != nil {
		log.Printf("client err %v\n", err)
	}
	client.Close()
}
