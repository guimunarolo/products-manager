package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"github.com/go-pg/pg/v10"
	
	"github.com/guimunarolo/products-manage/calculator-service/calculator"
)

func main() {
	logger := hclog.Default()

	var db *pg.DB
	{
		opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil{
			logger.Error("Unable to parse db url", "error", err)
			os.Exit(-1)
		}

		db = pg.Connect(opt)
	}

	var calc *calculator.Calculator
	{
		urep := calculator.NewUserRepo(db, logger)
		prep := calculator.NewProductRepo(db, logger)

		calc = calculator.NewCalculator(logger, urep, prep)
	}

	server := grpc.NewServer()
	calculator.RegisterCalculatorServer(server, calc)

	port := 9000
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("Unable to create listener", "error", err)
		os.Exit(-1)
	}

	logger.Info("gRPC Server running", "port", port)
	server.Serve(l)
}
