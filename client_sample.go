// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"syscall"
	"time"
)

var (
	ip          = flag.String("ip", "127.0.0.1", "server IP")
	connections = flag.Int("conn", 10, "number of tcp connections")
)

func main() {
	flag.Parse()

	setLimit()

	addr := *ip + ":8087"
	log.Println("service addr", addr)
	var conns []net.Conn
	for i := 0; i < *connections; i++ {
		c, err := net.DialTimeout("tcp", addr, 10*time.Second)
		if err != nil {
			log.Println("failed to connect", i, err)
			i--
			continue
		}
		conns = append(conns, c)
		time.Sleep(200 * time.Millisecond)
	}

	defer func() {
		for _, c := range conns {
			log.Printf("LocalAddr: %s", c.LocalAddr())
			c.Close()
		}
	}()

	log.Printf("init conn %d ", len(conns))

	tts := 3 * time.Second
	for range time.Tick(tts) {
		for i := 0; i < len(conns); i++ {
			conn := conns[i]
			log.Printf("send data LocalAddr: %s, index %d", conn.LocalAddr(), i)
			msg := fmt.Sprintf("hello world %d->%s\r\n", i, conn.LocalAddr())
			conn.Write([]byte(msg))
		}
	}
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
