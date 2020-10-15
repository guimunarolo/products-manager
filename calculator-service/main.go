package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	protos "github.com/guimunarolo/products-manage/calculator-service/protos/calculator"
	"github.com/guimunarolo/products-manage/calculator-service/server"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	c := server.NewCalculator(log)

	protos.RegisterCalculatorServer(gs, c)
	reflection.Register(gs)

	port := 9000
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	log.Info("gRPC Server running", "port", port)
	gs.Serve(l)
}
