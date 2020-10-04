package main

import (
	"example.com/test/pkg/adder"
	"example.com/test/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalf("Error Listen tcp %v", err)
	}

	if err = s.Serve(l); err != nil{
		log.Fatalf("Error Serve gprc NewServer %v", err)
	}
}