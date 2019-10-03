package main

import (
	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(c gnet.Conn) (out []byte, action gnet.Action) {
	top, tail := c.ReadPair()
	out = top
	if tail != nil {
		out = append(top, tail...)
	}
	c.ResetBuffer()
	return
}

func main() {
	echo := new(echoServer)
	err := gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true))
	if err != nil {
		panic(err)
	}
}
