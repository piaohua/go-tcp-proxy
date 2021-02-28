# go-tcp-proxy

A example for small TCP proxy written in Go.

### How to test

go get github.com/piaohua/go-tcp-proxy

- 启动一个tcp服务，监听`8088`端口
```
$ nc -lk 127.0.0.1 8088
hello
```

- build `server`

```
$ GO111MODULE=off go build server.go

```

- run command `./server`

```
$ ./server

```

- 创建tcp连接，并发送消息
```
$ nc -v 127.0.0.1 8087
found 0 associations
found 1 connections:
     1: flags=82<CONNECTED,PREFERRED>
        outif lo0
        src 127.0.0.1 port 58804
        dst 127.0.0.1 port 8088
        rank info not available
        TCP aux info available

Connection to 127.0.0.1 port 8088 [tcp/radan-http] succeeded!
hello
```

### Reference

- [go-tcp-proxy](https://github.com/jpillora/go-tcp-proxy)
- [用Go实现TCP连接的双向拷贝](https://zhuanlan.zhihu.com/p/29657180)
- [1m-go-tcp-server](https://github.com/smallnest/1m-go-tcp-server)
