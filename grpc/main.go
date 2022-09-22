package main

import (
	"github.com/abassGarane/currency/protos"
	"github.com/abassGarane/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	log := hclog.Default()
	g := grpc.NewServer()
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		panic(err)
	}
	reflection.Register(g)
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(g, cs)
	if err = g.Serve(l); err != nil {
		panic(err)
	}
}
