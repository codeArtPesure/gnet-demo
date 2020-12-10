package main

import (
	"fmt"
	"gnet-demo/example"
	"log"
	"sync"
	"time"

	"github.com/panjf2000/gnet"
)

type pushServer struct {
	*gnet.EventServer
	tick             time.Duration
	connectedSockets sync.Map
}

func (ps *pushServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Push server is listening on %s (multi-cores: %t, loops: %d), "+
		"pushing data every %s ...\n", srv.Addr.String(), srv.Multicore, srv.NumEventLoop, ps.tick.String())
	return
}

func (ps *pushServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("Socket with addr: %s has been opened...\n", c.RemoteAddr().String())
	ps.connectedSockets.Store(c.RemoteAddr().String(), c)
	return
}

func (ps *pushServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("Socket with addr: %s is closing...\n", c.RemoteAddr().String())
	ps.connectedSockets.Delete(c.RemoteAddr().String())
	return
}

func (ps *pushServer) Tick() (delay time.Duration, action gnet.Action) {
	log.Println("It's time to push data to clients!!!")
	ps.connectedSockets.Range(func(key, value interface{}) bool {
		addr := key.(string)
		c := value.(gnet.Conn)
		c.AsyncWrite([]byte(fmt.Sprintf("heart beating to %s\n", addr)))
		return true
	})
	delay = ps.tick
	return
}

func (ps *pushServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	out = frame
	return
}

func main() {
	push := &pushServer{tick: time.Second}
	go example.StartEchoClient("0.0.0.0:9000", "hello")
	log.Fatal(gnet.Serve(push, fmt.Sprintf("tcp://:%d", 9000), gnet.WithMulticore(true), gnet.WithTicker(true)))
}
