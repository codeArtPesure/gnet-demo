package main

import (
	"gnet-demo/example"
	"time"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	time.Sleep(1 * time.Second)
	out = frame
	return
}

func main() {
	echo := new(echoServer)
	go gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true))
	go example.StartEchoClient("0.0.0.0:9000", "hello")
	time.Sleep(5 * time.Minute)
}
