package ste_server

// wait for request from web & worker
// listen on ports: 1234, 1235(config?)
// handle requests on 

import (
	"net/rpc"
	"log"
	"net/http"
	"net"
)

type Args struct {
	Count int
}

var queues = new(Queues)

func Run() {
	exQueues := new(ExQueues)
	rpc.Register(exQueues)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}