// +build ignore

package main

import (
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"syscall"
)

func main() {
	setLimit()

	listener, err := net.Listen("tcp", ":8087")
	if err != nil {
		panic(err)
	}

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()

	for {
		conn, e := listener.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		go handleConn(conn)
		connections = append(connections, conn)
		if (len(connections) % 100) == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}
}

func handleConn(conn net.Conn) {
	log.Printf("LocalAddr:%s, RemoteAddr:%s\n", conn.LocalAddr(), conn.RemoteAddr())
	//io.Copy(ioutil.Discard, conn)
	io.Copy(os.Stdout, conn)
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	log.Printf("rLimit %+v\n", rLimit)

	//TODO rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}
